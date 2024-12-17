package conf

const (
	defaultUrl      = "https://a.stu.edu.cn/ac_portal/login.php"
	defaultUsername = "username"
	defaultPasswrod = "password"
	defaultAutoExit = "FALSE"
)

func setDefault() {
	Url = defaultUrl
	Username = defaultUsername
	Password = defaultPasswrod
	AutoExit = defaultAutoExit
}
