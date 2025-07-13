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
		OutOrderNo:   "111",
		Amount:       "1000000",
		BankName:     "ACB",
		BankBranch:   "aa",
		BankUserName: "cy",
		BankCard:     "107719719971",
	}
}
