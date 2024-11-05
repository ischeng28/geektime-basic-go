package web

import (
	"errors"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ischeng28/basic-go/webook/internal/domain"
	"github.com/ischeng28/basic-go/webook/internal/service"
	ijwt "github.com/ischeng28/basic-go/webook/internal/web/jwt"
	"net/http"
)

// UserHandler 在上面定义跟用户有关的路由
type UserHandler struct {
	ijwt.Handler
	emailRexExp *regexp.Regexp
	passwordExp *regexp.Regexp
	svc         service.UserService
}

func NewUserHandler(svc service.UserService, hdl ijwt.Handler, codeSvc service.CodeService) *UserHandler {
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		// 和上面比起来，用 ` 看起来就比较清爽
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailRexExp: emailExp,
		passwordExp: passwordExp,
		svc:         svc,
		Handler:     hdl,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.LoginJWT)
	ug.POST("/edit", u.Edit)
	ug.GET("/refresh_token", u.RefreshToken)
	ug.GET("/logout", u.LogoutJWT)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	var req SignUpReq
	// Bind方法会根据 Content-Type 来解析你的数据到req里面
	// 解析错了，就会返回一个400的错误
	if err := ctx.Bind(&req); err != nil {
		return
	}

	ok, err := u.emailRexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "你的邮箱格式不正确")
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字和特殊符号")
		return
	}

	err = u.svc.Signup(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	switch {
	case err == nil:
		ctx.String(http.StatusOK, "注册成功")
	case errors.Is(err, service.ErrDuplicateEmail):
		ctx.String(http.StatusOK, "邮箱冲突，请换一个")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}

}

func (h *UserHandler) LoginJWT(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	u, err := h.svc.Login(ctx, req.Email, req.Password)
	switch {
	case err == nil:
		err = h.SetLoginToken(ctx, u.Id)
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
			return
		}
		ctx.String(http.StatusOK, "登录成功")
	case errors.Is(err, service.ErrInvalidUserOrPassword):
		ctx.String(http.StatusOK, "用户名或者密码不正确")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	u, err := h.svc.Login(ctx, req.Email, req.Password)
	switch {
	case err == nil:
		//设置session
		sess := sessions.Default(ctx)
		//要放到session中的值
		sess.Set("userId", u.Id)
		sess.Options(sessions.Options{
			Secure:   true,
			HttpOnly: true,
			MaxAge:   60 * 30,
		})
		sess.Save()
		ctx.String(http.StatusOK, "登录成功")
	case errors.Is(err, service.ErrInvalidUserOrPassword):
		ctx.String(http.StatusOK, "用户名或者密码不正确")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (h *UserHandler) Profile(ctx *gin.Context) {
	uc, ok := ctx.MustGet("user").(ijwt.UserClaims)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	u, err := h.svc.FindById(ctx, uc.Uid)
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}
	type User struct {
		Nickname string `json:"nickname"`
		Email    string `json:"Email"`
		AboutMe  string `json:"aboutMe"`
		Birthday string `json:"birthday"`
	}
	ctx.JSON(http.StatusOK, User{
		Email: u.Email,
	})
}

func (h *UserHandler) RefreshToken(ctx *gin.Context) {
	// 约定，前端在 Authorization 里面带上这个 refresh_token
	tokenStr := h.ExtractToken(ctx)
	var rc ijwt.RefreshClaims
	token, err := jwt.ParseWithClaims(tokenStr, &rc, func(token *jwt.Token) (interface{}, error) {
		return ijwt.RCJWTKey, nil
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if token == nil || !token.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = h.CheckSession(ctx, rc.Ssid)
	if err != nil {
		// token 无效或者 redis 有问题
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = h.SetJWTToken(ctx, rc.Uid, rc.Ssid)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Msg: "OK",
	})
}

func (h *UserHandler) LogoutJWT(ctx *gin.Context) {
	err := h.ClearToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{Code: 5, Msg: "系统错误"})
		return
	}
	ctx.JSON(http.StatusOK, Result{Msg: "退出登录成功"})
}
