package xarr_pay_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

type XArrPay struct {
	uri   string
	pid   int32
	token string
}

func New(uri string, pid int32, token string) *XArrPay {
	return &XArrPay{
		uri:   strings.TrimRight(uri, "/"),
		pid:   pid,
		token: token,
	}
}

// 创建订单
func (s *XArrPay) CreateOrder(req *OrderCreateReq) (*OrderCreateRes, error) {
	req.Pid = s.pid
	req.Sign = GenerateSign(req, s.token)

	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// 请求支付通道
	client := http.Client{}

	response, err := client.Post(s.uri+"/api/order/create", "text/json", bytes.NewReader(marshal))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(data))
	res := &OrderCreateRes{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, errors.New(res.Message)
	}

	return res, nil
}

// 获取订单状态
func (s *XArrPay) GetOrderStatus() {

}
