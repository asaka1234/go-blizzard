package go_blizzard

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-blizzard/utils"
	"github.com/mitchellh/mapstructure"
)

func (cli *Client) Deposit(req BlizzardDepositReq) (*BlizzardDepositResponse, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]string
	mapstructure.Decode(req, &params)

	params["appId"] = cli.Params.MerchantId
	params["channelId"] = "104"
	params["callbackUrl"] = cli.Params.DepositBackUrl
	params["successUrl"] = cli.Params.DepositFeBackUrl

	//签名
	signStr := utils.Sign(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result BlizzardDepositResponse

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetFormData(params).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode())
	}

	if resp.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp.Error(), resp.Body())
	}

	fmt.Printf("==>%+v\n", string(resp.Body()))

	return &result, nil
}
