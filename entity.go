package go_exglobal

type ExglobalInitParams struct {
	MerchantId int64  `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"` // merchantId
	AccessKey  string `json:"accessKey" mapstructure:"accessKey" config:"accessKey"  yaml:"accessKey"`     //接入秘钥
	BackKey    string `json:"backKey" mapstructure:"backKey" config:"backKey"  yaml:"backKey"`             //回调秘钥

	DepositUrl  string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`
	WithdrawUrl string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
	//DepositBackUrl  string `json:"depositBackUrl" mapstructure:"depositBackUrl" config:"depositBackUrl"  yaml:"depositBackUrl"`
	//WithdrawBackUrl string `json:"WithdrawBackUrl" mapstructure:"WithdrawBackUrl" config:"WithdrawBackUrl"  yaml:"WithdrawBackUrl"`
}

// ----------pre order-------------------------

// 5.2. Create a collection order
type ExglobalDepositReq struct {
	MerchantOrderNo  string  `json:"merchantOrderNo" mapstructure:"merchantOrderNo"`   //商户订单号CurrencyCoinName string `json:"currencyCoinName" mapstructure:"currencyCoinName"` //支持VND
	CurrencyCoinName string  `json:"currencyCoinName" mapstructure:"currencyCoinName"` //支持 VND
	Amount           float64 `json:"amount" mapstructure:"amount"`                     //不支持小数
	PaymentMethod    int     `json:"paymentMethod" mapstructure:"paymentMethod"`       //枚举: 1->JSON,3->平台收银台
	//SDK帮计算
	//Signature string `json:"signature" mapstructure:"signature"` //签名
	//UID       int64  `json:"uid" mapstructure:"uid"`             //商户号
	//BankCode string `json:"bankCode" mapstructure:"bankCode"` //option, 如果用 ScanQRCode, 则这里填 AllBanksSupported
	//ChannelCode      string  `json:"channelCode" mapstructure:"channelCode"`           //网银扫码:ScanQRCode, 银行直连:BankDirect
}

type ExglobalDepositResponse struct {
	Code    int    `json:"code" mapstructure:"code"`       // 1是成功
	Success bool   `json:"success" mapstructure:"success"` //true
	Message string `json:"message" mapstructure:"message"`
	Data    struct {
		UID             int64   `json:"uid" mapstructure:"uid"`                         //商户号
		MerchantOrderNo string  `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //商户订单号
		RecordId        int64   `json:"recordId" mapstructure:"recordId"`               //平台订单号
		Amount          float64 `json:"amount" mapstructure:"amount"`                   //订单金额
		URL             string  `json:"url" mapstructure:"url"`                         //收银台地址
		QRCode          string  `json:"qrcode" mapstructure:"qrcode"`                   //二维码内容
		MemoCode        string  `json:"memoCode" mapstructure:"memoCode"`               //附言码
		//以下字段不会返回
		BankId        *string `json:"bankId" mapstructure:"bankId"`
		AccountNumber *string `json:"accountNumber" mapstructure:"accountNumber"`
		BankOwner     *string `json:"bankOwner" mapstructure:"bankOwner"`
		Signature     string  `json:"signature" mapstructure:"signature"`
	} `json:"data" mapstructure:"data"`
}

// ------------------------------------------------------------
type ExglobalDepositBackReq struct {
	RecordId        int64  `json:"recordId" mapstructure:"recordId"`               //平台订单号
	UID             int64  `json:"uid" mapstructure:"uid"`                         //商户号
	OrderAmount     string `json:"orderAmount" mapstructure:"orderAmount"`         //订单金额与下单实际金额一致
	MerchantOrderNo string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //商户订单号
	TradeStatus     int    `json:"tradeStatus" mapstructure:"tradeStatus"`         //2支付中, 3支付完成, 4支付失败
	Signature       string `json:"signature" mapstructure:"signature"`
}

type ExglobalDepositBackResp struct {
	Code    int    `json:"code" mapstructure:"code"` // 1是成功
	Success bool   `json:"success" mapstructure:"success"`
	Message string `json:"message" mapstructure:"message"`
	Data    string `json:"data" mapstructure:"data"`
}

//===========withdraw===================================

type ExglobalWithdrawReq struct {
	MerchantOrderNo  string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"`   //商户订单号
	CurrencyCoinName string `json:"currencyCoinName" mapstructure:"currencyCoinName"` //VND
	Amount           string `json:"amount" mapstructure:"amount"`                     //不支持小数
	BankCode         string `json:"bankCode" mapstructure:"bankCode"`
	BankName         string `json:"bankName" mapstructure:"bankName"`
	BankBranchName   string `json:"bankBranchName" mapstructure:"bankBranchName"`
	BankUserName     string `json:"bankUserName" mapstructure:"bankUserName"`
	BankAccount      string `json:"bankAccount" mapstructure:"bankAccount"`
	//以下sdk帮搞
	//UID       int64  `json:"uid" mapstructure:"uid"`             //商户编码
	//ChannelCode      string `json:"channelCode" mapstructure:"channelCode"` //写死 BankDirect
	//Signature string `json:"signature" mapstructure:"signature"` //签名
}

type ExglobalWithdrawResponse struct {
	Code    int    `json:"code" mapstructure:"code"`       // 1是成功
	Success bool   `json:"success" mapstructure:"success"` //true
	Message string `json:"message" mapstructure:"message"`
	Data    struct {
		Uid              int64  `json:"uid" mapstructure:"channelCode"`                 //商户编号
		MerchantOrderNo  string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //商户订单号
		RecordId         int64  `json:"recordId" mapstructure:"recordId"`               //平台订单号
		Amount           string `json:"amount" mapstructure:"amount"`
		CurrencyCoinName string `json:"currencyCoinName" mapstructure:"currencyCoinName"` //VND
		BankName         string `json:"bankName" mapstructure:"bankName"`
		BankBranchName   string `json:"bankBranchName" mapstructure:"bankBranchName"`
		BankUserName     string `json:"bankUserName" mapstructure:"bankUserName"`
		BankAccount      string `json:"bankAccount" mapstructure:"bankAccount"`
	} `json:"data" mapstructure:"data"`
}

type ExglobalWithdrawBackReq struct {
	RecordId        int64  `json:"recordId" mapstructure:"recordId"`               //平台订单号
	UID             int64  `json:"uid" mapstructure:"uid"`                         //商户号
	OrderAmount     string `json:"orderAmount" mapstructure:"orderAmount"`         //订单金额与下单实际金额一致
	MerchantOrderNo string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //商户订单号
	TradeStatus     int    `json:"tradeStatus" mapstructure:"tradeStatus"`         //2支付中, 3支付完成, 4支付失败
	Signature       string `json:"signature" mapstructure:"signature"`
}

type ExglobalWithdrawBackResp struct {
	Code    int    `json:"code" mapstructure:"code"` // 1是成功
	Success bool   `json:"success" mapstructure:"success"`
	Message string `json:"message" mapstructure:"message"`
	Data    string `json:"data" mapstructure:"data"`
}
