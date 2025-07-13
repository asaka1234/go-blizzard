package go_blizzard

type BlizzardInitParams struct {
	MerchantId string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"` // merchantId
	AccessKey  string `json:"accessKey" mapstructure:"accessKey" config:"accessKey"  yaml:"accessKey"`     //接入秘钥
	BackKey    string `json:"backKey" mapstructure:"backKey" config:"backKey"  yaml:"backKey"`             //回调秘钥

	DepositUrl       string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`
	WithdrawUrl      string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
	DepositBackUrl   string `json:"depositBackUrl" mapstructure:"depositBackUrl" config:"depositBackUrl"  yaml:"depositBackUrl"`
	DepositFeBackUrl string `json:"depositFeBackUrl" mapstructure:"depositFeBackUrl" config:"depositFeBackUrl"  yaml:"depositFeBackUrl"`
	WithdrawBackUrl  string `json:"WithdrawBackUrl" mapstructure:"WithdrawBackUrl" config:"WithdrawBackUrl"  yaml:"WithdrawBackUrl"`
}

// ----------pre order-------------------------

// 5.2. Create a collection order
type BlizzardDepositReq struct {
	OutTradeNo string `json:"outTradeNo" mapstructure:"outTradeNo"` //商户订单号
	Amount     string `json:"amount" mapstructure:"amount"`         // 2位小数
	UserName   string `json:"userName" mapstructure:"userName"`     //付款人name
	UserPhone  string `json:"userPhone" mapstructure:"userPhone"`   //付款人mobile
	UserEmail  string `json:"userEmail" mapstructure:"userEmail"`   //付款人email
	//SDK帮填写
	//AppId       string  `json:"appId" mapstructure:"appId"`             //商户号
	//Sign        string `json:"sign" mapstructure:"sign"`               //签名
	//CallbackUrl string `json:"callbackUrl" mapstructure:"callbackUrl"` //结果回调
	//SuccessUrl  string `json:"successUrl" mapstructure:"successUrl"`   //前端跳转
	//ChannelId  string `json:"channelId" mapstructure:"channelId"`   //通道code, 通过这个指定的ccy (只支持THB 104 写死)
}

type BlizzardDepositResponse struct {
	Code int    `json:"code" mapstructure:"code"` //200成功，成功时没有msg数据
	Msg  string `json:"msg" mapstructure:"msg"`
	Data struct {
		AppId      int64   `json:"appId" mapstructure:"appId"`
		ChannelId  string  `json:"channelId" mapstructure:"channelId"`
		OrderNo    string  `json:"orderNo" mapstructure:"orderNo"`
		Amount     float64 `json:"amount" mapstructure:"amount"`
		OutTradeNo string  `json:"outTradeNo" mapstructure:"outTradeNo"`
		PayUrl     string  `json:"payUrl" mapstructure:"payUrl"`
	} `json:"data"`
}

// ------------------------------------------------------------
type BlizzardDepositBackReq struct {
	AppId      interface{} `json:"appId" mapstructure:"appId"`
	OutTradeNo string      `json:"outTradeNo" mapstructure:"outTradeNo"` //商户订单号
	OrderNo    string      `json:"orderNo" mapstructure:"orderNo"`       //平台订单号

	ChannelId  string      `json:"channelId" mapstructure:"channelId"`   //
	Amount     interface{} `json:"amount" mapstructure:"amount"`         //金额
	AmountTrue interface{} `json:"amountTrue" mapstructure:"amountTrue"` //金额
	PayStatus  string      `json:"payStatus" mapstructure:"payStatus"`   //SUCCESS
	Sign       string      `json:"sign" mapstructure:"sign"`
}

//回调返回 SUCCESS

//===========withdraw===================================

type BlizzardWithdrawReq struct {
	OutOrderNo   string      `json:"outOrderNo" mapstructure:"outOrderNo"` //商户订单号
	Amount       interface{} `json:"amount" mapstructure:"amount"`         //todo 2位小数
	BankName     string      `json:"bankName" mapstructure:"bankName"`
	BankBranch   string      `json:"bankBranch" mapstructure:"bankBranch"`
	BankUserName string      `json:"bankUserName" mapstructure:"bankUserName"`
	BankCard     string      `json:"bankCard" mapstructure:"bankCard"`
	//以下sdk帮搞
	//AppId int64  `json:"appId" mapstructure:"appId"` //商户编码
	//Sign  string `json:"sign" mapstructure:"sign"`   //签名
	//Currency string `json:"currency" mapstructure:"currency"` //THB..
}

type BlizzardWithdrawResponse struct {
	Code int    `json:"code" mapstructure:"code"` //200成功，成功时没有msg数据
	Msg  string `json:"msg" mapstructure:"msg"`
	Data struct {
		AppId      int64       `json:"appId" mapstructure:"appId"`
		OrderNo    string      `json:"orderNo" mapstructure:"orderNo"`
		Apply      string      `json:"apply" mapstructure:"apply"`
		OutTradeNo string      `json:"outTradeNo" mapstructure:"outTradeNo"`
		Fee        interface{} `json:"fee" mapstructure:"fee"`
	} `json:"data"`
}

type BlizzardWithdrawBackReq struct {
	AppId      interface{} `json:"appId" mapstructure:"appId"`
	OutTradeNo string      `json:"outTradeNo" mapstructure:"outTradeNo"` //商户订单号
	OrderNo    string      `json:"orderNo" mapstructure:"orderNo"`       //平台订单号

	Currency    string      `json:"currency" mapstructure:"currency"`       //ccy
	OrderStatus interface{} `json:"orderStatus" mapstructure:"orderStatus"` //Order Status: 0 Not Processed,  1 Paid, 2 Rejected
	Sign        string      `json:"sign" mapstructure:"sign"`
}

//After the callback succeeds, return the SUCCESS string
