package xarr_pay_sdk

import (
	"log"
	"testing"
)

func getSdk() *XArrPay {
	return New("http://127.0.0.1:24017", 10000, "fc5778SmmANqnfoKosgz")

}

func TestCreateOrder(t *testing.T) {
	xarr := getSdk()

	orderReq := &OrderCreateReq{
		PayType:    PAY_TYPE_ALIPAY,
		OutTradeNo: "123456222yr",
		NotifyUrl:  "",
		ReturnUrl:  "",
		Name:       "测试支付2",
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

func TestGetOrderInfo(t *testing.T) {
	xarr := getSdk()

	status, err := xarr.GetOrderStatus("123456222")
	if err != nil {
		return
	}
	log.Println(status)
}
