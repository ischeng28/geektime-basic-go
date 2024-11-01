package repository

import (
	"context"
	"github.com/ischeng28/basic-go/webook/internal/domain"
	"github.com/ischeng28/basic-go/webook/internal/repository/cache"
	"github.com/ischeng28/basic-go/webook/internal/repository/dao"
	"log"
)

var (
	ErrDuplicateUser  = dao.ErrDuplicateEmail
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrUserNotFound
)

type CachedUserRepository struct {
	dao   dao.UserDAO
	cache cache.UserCache
}

type UserRepository interface {
	Create(ctx context.Context, u domain.User) error
	FindById(ctx context.Context, uid int64) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	toDomain(u dao.User) domain.User
	FindByWechat(ctx context.Context, openId string) (domain.User, error)
}

func NewUserRepository(dao dao.UserDAO, cache cache.UserCache) UserRepository {
	return &CachedUserRepository{
		dao:   dao,
		cache: cache,
	}
}

func (repo *CachedUserRepository) Create(ctx context.Context, u domain.User) error {
	return repo.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
	//	在这里操作缓存
}

func (r *CachedUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (repo *CachedUserRepository) FindByWechat(ctx context.Context, openId string) (domain.User, error) {
	ue, err := repo.dao.FindByWechat(ctx, openId)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(ue), nil
}

func (repo *CachedUserRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		WechatInfo: domain.WechatInfo{
			OpenId:  u.WechatOpenId.String,
			UnionId: u.WechatUnionId.String,
		},
	}
}

func (repo *CachedUserRepository) FindById(ctx context.Context, uid int64) (domain.User, error) {
	du, err := repo.cache.Get(ctx, uid)
	// 只要 err 为 nil，就返回
	if err == nil {
		return du, nil
	}

	// err 不为 nil，就要查询数据库
	// err 有两种可能
	// 1. key 不存在，说明 redis 是正常的
	// 2. 访问 redis 有问题。可能是网络有问题，也可能是 redis 本身就崩溃了

	u, err := repo.dao.FindById(ctx, uid)
	if err != nil {
		return domain.User{}, err
	}

	du = repo.toDomain(u)
	go func() {
		err = repo.cache.Set(ctx, du)
		if err != nil {
			// 网络崩了，也可能是 redis 崩了
			log.Println(err)
		}
	}()

	return du, nil
}
