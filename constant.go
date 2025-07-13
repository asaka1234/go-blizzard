package go_blizzard

type BlizzardChannelCode struct {
	Currency string `json:"currency"`
	Code     string `json:"code"`
	Name     string `json:"name"`
}

var BlizzardChannelCodes = []BlizzardChannelCode{
	{"THB", "104", "Thailand - Bank Transfer"},
	//{"THB", "105", "Thailand - Promptpay"}, //不可用

	//{"IDR", "113", "Indonesia - Bank Card VA"}, //印尼
	//{"IDR", "114", "Thailand - Bank Transfer"}, //不可用

	//{"EGP", "119", "Egypt - e-wallet"},
}
