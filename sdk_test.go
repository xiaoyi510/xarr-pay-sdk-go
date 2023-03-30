package xarr_pay_sdk

import (
	"log"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	xarr := New("http://127.0.0.1:24017", 10000, "")

	orderReq := &OrderCreateReq{
		PayType:    PAY_TYPE_ALIPAY,
		OutTradeNo: "123456222",
		NotifyUrl:  "",
		ReturnUrl:  "",
		Name:       "测试支付",
		Money:      "1.2",
		ClientIp:   "",
		Device:     "",
		Param:      "哈哈",
	}
	order, err := xarr.CreateOrder(orderReq)
	if err != nil {
		log.Println("支付错误", err.Error())
		return
	}
	log.Println("支付成功", order)
}
