package go_blizzard

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &BlizzardInitParams{MERCHANT_ID, ACCESS_SECRET, BACK_SECRET, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_BACK_URL, WITHDRAW_BACK_URL, DEPOSIT_FE_BACK_URL})

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() BlizzardWithdrawReq {
	return BlizzardWithdrawReq{
		OutOrderNo:   "202512110937738",
		Amount:       "600",
		BankName:     "ACB", //印度使用ifsc code
		BankBranch:   "ACB",
		BankUserName: "jane",
		BankCard:     "107719719971",
		Currency:     "INR",
	}
}

func TestWithdrawBack(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &BlizzardInitParams{MERCHANT_ID, ACCESS_SECRET, BACK_SECRET, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_BACK_URL, WITHDRAW_BACK_URL, DEPOSIT_FE_BACK_URL})

	//发请求
	err := cli.WithdrawCallBack(GenWithdrawBackRequestDemo(), withdrawProcessor)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
}

func withdrawProcessor(BlizzardWithdrawBackReq) error {
	fmt.Println("withdraw back success!")
	return nil
}

// {"appId":"10511","currency":"INR","orderNo":"T10511251211155443448600","orderStatus":"1","outOrderNo":"202512110937738","sign":"471b3db0847ada7b7d01431be0446310"}
func GenWithdrawBackRequestDemo() BlizzardWithdrawBackReq {
	return BlizzardWithdrawBackReq{
		AppId:       "10511",
		Currency:    "INR",
		OrderNo:     "T10511251211155443448600",
		OrderStatus: "1",
		OutOrderNo:  "202512110937738",
		Sign:        "471b3db0847ada7b7d01431be0446310",
	}
}
