package config

//阿里云短信的结构体
type Sms struct {
	AccessKeyId     string `mapstructure:"access-key-id" json:"AccessKeyId" yaml:"access-key-id"`             // key
	AccessKeySecret string `mapstructure:"access-key-secret" json:"AccessKeySecret" yaml:"access-key-secret"` // 加密秘钥
	SignName        string `mapstructure:"sign-name" json:"SignName" yaml:"sign-name"`                        // 签名
	TemplateCode    string `mapstructure:"template-code" json:"TemplateCode" yaml:"template-code"`            // 短信模板编号
}
