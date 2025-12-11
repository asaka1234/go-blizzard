package go_blizzard

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &BlizzardInitParams{MERCHANT_ID, ACCESS_SECRET, BACK_SECRET, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_BACK_URL, WITHDRAW_BACK_URL, DEPOSIT_FE_BACK_URL})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() BlizzardDepositReq {
	return BlizzardDepositReq{
		OutTradeNo: "202512112323212414", //商户id
		Amount:     "1199.00",
		UserName:   "John",
		UserPhone:  "12345",
		UserEmail:  "111",
		ChannelId:  "108", // THB=104,INR=108
	}
}

func TestDepositBack(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &BlizzardInitParams{MERCHANT_ID, ACCESS_SECRET, BACK_SECRET, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_BACK_URL, WITHDRAW_BACK_URL, DEPOSIT_FE_BACK_URL})

	//发请求
	err := cli.DepositCallback(GenDepositBackRequestDemo(), depositProcessor)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
}

func depositProcessor(BlizzardDepositBackReq) error {
	fmt.Println("deposit back success!")
	return nil
}

func GenDepositBackRequestDemo() BlizzardDepositBackReq {
	return BlizzardDepositBackReq{
		Amount:     "1199.00",
		AmountTrue: "1199.00",
		AppId:      "10511",
		ChannelId:  "108", // THB=104,INR=108
		OrderNo:    "10511251211151243448101",
		OutTradeNo: "202512112323212414", //商户id
		PayStatus:  "SUCCESS",
		Sign:       "dea66ceb254a3057079b352fa642e12d",
	}
}
