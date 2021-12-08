package request

// User register structure
type Register struct {
	Username     string   `json:"userName"`
	Password     string   `json:"passWord"`
	NickName     string   `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg    string   `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	AuthorityId  string   `json:"authorityId" gorm:"default:888"`
	AuthorityIds []string `json:"authorityIds"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// app User login structure
type AppLogin struct {
	Username string `json:"username" example:"手机号"` // 手机号
	Password string `json:"password" example:"密码"`  // 密码
}

// SMS User login structure
type SmsLogin struct {
	Mobile string `json:"mobile" example:"手机号"` // 手机号
	Code   string `json:"code" example:"验证码"`   // 验证码
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username" example:"用户名"`    // 用户名
	Password    string `json:"password" example:"旧密码"`    // 密码
	NewPassword string `json:"newPassword" example:"新密码"` // 新密码
}

//接收的忘记密码的传参结构体
type ForgetREQStruct struct {
	Mobile   string `json:"phone"`    // 手机号
	Code     string `json:"code"`     // 验证码
	Password string `json:"password"` // 新密码
}

// Forget password structure
type ForgetPasswordStruct struct {
	Mobile      string `json:"mobile"`      // 手机号
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId string `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}
