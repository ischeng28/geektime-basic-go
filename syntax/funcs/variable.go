package main

func YourName(name string, aliases ...string) {
	if len(aliases) > 0 {
		println(aliases[0])
	}
}

func YourNameInvoke() {
	YourName("cheng")
	YourName("cheng", "hui")
	YourName("cheng", "xin", "hui")
}
