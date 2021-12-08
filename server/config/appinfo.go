package config

//app返回信息结构体
type Appinfo struct {
	Logo           string `mapstructure:"logo" json:"Logo" yaml:"logo"`                               // logo
	Mobile         string `mapstructure:"mobile" json:"Mobile" yaml:"mobile"`                         // 官方电话
	Email          string `mapstructure:"email" json:"Email" yaml:"email"`                            // 官方邮箱
	Appname        string `mapstructure:"appname" json:"Appname" yaml:"appname"`                      // App名字
	Servicecontent string `mapstructure:"servicecontent" json:"Servicecontent" yaml:"servicecontent"` // App名字
	Version        string `mapstructure:"version" json:"Version" yaml:"version"`                      // 版本号
}
