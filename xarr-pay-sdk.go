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
func (s *XArrPay) CreateOrder(req *OrderCreateReq) (*OrderCreateData, error) {
	req.Pid = s.pid
	req.Sign = GenerateSign(req, s.token)

	data, err := s.do(req, "order/create")
	res := &OrderCreateRes{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, errors.New(res.Message)
	}

	return res.Data, nil
}

func (s *XArrPay) do(req interface{}, api string) ([]byte, error) {

	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// 请求支付通道
	client := http.Client{}

	response, err := client.Post(s.uri+"/xpay/"+api, "text/json", bytes.NewReader(marshal))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(data))

	return data, nil
}

// 获取订单状态
// @orderId 为商户ID
func (s *XArrPay) GetOrderStatus(orderId string) (*OrderStatusData, error) {
	req := &OrderStatusReq{}
	req.Pid = s.pid
	req.OutOrderId = orderId
	req.Sign = GenerateSign(req, s.token)

	data, err := s.do(req, "order/status")
	if err != nil {
		return nil, err
	}
	res := &OrderStatusRes{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, errors.New(res.Message)
	}

	return res.Data, nil
}

// 处理回调事件
func (s *XArrPay) CallbackCheck(req *OrderNotifyCallback) error {
	if req.Sign != GenerateSign(req, s.token) {
		return errors.New("校验失败")
	}
	return nil
}
