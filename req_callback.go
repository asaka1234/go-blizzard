package go_exglobal

import (
	"errors"
	"github.com/asaka1234/go-exglobal/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallback(req ExglobalDepositBackReq, processor func(ExglobalDepositBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySign(params, cli.Params.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if req.UID != cli.Params.MerchantId {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}

//==========================================

// 充值的回调处理(传入一个处理函数)
func (cli *Client) WithdrawCallBack(req ExglobalWithdrawBackReq, processor func(ExglobalWithdrawBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySign(params, cli.Params.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if req.UID != cli.Params.MerchantId {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}
