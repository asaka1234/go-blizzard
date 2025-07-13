package go_blizzard

import (
	"crypto/tls"
	"github.com/asaka1234/go-blizzard/utils"
	"github.com/mitchellh/mapstructure"
)

// dai
func (cli *Client) Deposit(req BlizzardDepositReq) (*BlizzardDepositResponse, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["uid"] = cli.Params.MerchantId
	params["channelCode"] = "ScanQRCode"     //写死
	params["bankCode"] = "AllBanksSupported" //写死

	//签名
	signStr := utils.Sign(params, cli.Params.AccessKey)
	params["signature"] = signStr

	//返回值会放到这里
	var result BlizzardDepositResponse

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetBody(params).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
