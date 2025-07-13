package go_blizzard

import (
	"errors"
	"github.com/asaka1234/go-blizzard/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallback(req BlizzardDepositBackReq, processor func(BlizzardDepositBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	paramsMap := ConvertToStringMap(params)

	verifyResult := utils.VerifySign(paramsMap, cli.Params.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if cast.ToString(req.AppId) != cli.Params.MerchantId {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}

//==========================================

// 充值的回调处理(传入一个处理函数)
func (cli *Client) WithdrawCallBack(req BlizzardWithdrawBackReq, processor func(BlizzardWithdrawBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	paramsMap := ConvertToStringMap(params)

	verifyResult := utils.VerifySign(paramsMap, cli.Params.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if cast.ToString(req.AppId) != cli.Params.MerchantId {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}

//---------------------------------

func ConvertToStringMap(m map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		value := cast.ToString(v)
		result[k] = value
	}
	return result
}
