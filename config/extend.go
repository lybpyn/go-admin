package config

var ExtConfig Extend

// Extend 扩展配置
//
//	extend:
//	  demo:
//	    name: demo-name
//
// 使用方法： config.ExtConfig......即可！！
type Extend struct {
	AMap     AMap     // 这里配置对应配置文件的结构即可
	Upload   Upload   // 文件上传配置
	PandaPay PandaPay // PandaPay支付配置
}

type AMap struct {
	Key string
}

// Upload 文件上传配置
type Upload struct {
	Path string // 上传文件存储路径
}

// PandaPay 支付配置
type PandaPay struct {
	ApiUrl     string // API地址，如：https://xxx/api
	MerchantId int    // 商户ID（纯数字）
	AppSecret  string // 商户密钥
	NotifyUrl  string // 异步通知地址
	Currency   string // 默认币种，如：NGN（尼日利亚-奈拉）、KES（肯尼亚先令）
}
