package go_blizzard

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-blizzard/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

// withdraw
func (cli *Client) Withdraw(req BlizzardWithdrawReq) (*BlizzardWithdrawResponse, error) {

	rawURL := cli.Params.WithdrawUrl

	var params map[string]string
	mapstructure.Decode(req, &params)
	params["appId"] = cli.Params.MerchantId
	params["currency"] = "THB" //写死

	//签名
	signStr := utils.Sign(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result BlizzardWithdrawResponse

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetFormData(params).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp))
	cli.logger.Infof("PSPResty#blizzad#withdraw->%+v", string(restLog))

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

	return &result, err
}
