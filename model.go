package xarr_pay_sdk

// 创建订单
type OrderCreateReq struct {
	Pid        int32  `p:"pid" json:"pid" v:"required#必填商户ID"`                    // 商户ID
	PayType    string `p:"pay_type" json:"pay_type" v:"required#必填支付方式"`          // 支付方式
	Name       string `p:"name" json:"name" v:"required#必填商品名称"`                  // 商品名称
	Money      string `p:"money" json:"money" v:"required#必填商品金额"`                // 商品金额	单位元
	OutTradeNo string `p:"out_trade_no" json:"out_trade_no" v:"required#必填商户订单号"` // 商户订单号
	NotifyUrl  string `p:"notify_url" json:"notify_url" `                         // 异步通知地址
	ReturnUrl  string `p:"return_url" json:"return_url" `                         // 跳转通知地址
	ClientIp   string `p:"clientip" json:"clientip" `                             // 用户IP地址
	Device     string `p:"device" json:"device"`                                  // 设备类型 根据当前用户浏览器的UA判断，	传入用户所使用的浏览器	或设备类型，默认为pc
	Param      string `p:"param" json:"param"`                                    // 业务扩展参数
	Sign       string `p:"sign" json:"sign" v:"required#签名不能为空"`
}

// 创建订单响应
type OrderCreateRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		TradeNo    string `json:"trade_no"`     // 平台支付订单号
		OutTradeNo string `json:"out_trade_no"` // 商户订单号
		ExpireTime int    `json:"expire_time"`  // 订单支付超时时间

		PayType     string `json:"pay_type"`     // 支付方式
		Amount      string `json:"amount"`       // 金额
		TradeAmount string `json:"trade_amount"` // 真实订单需要支付金额

		Uri    string `json:"uri"`    // 支付跳转地址
		Qrcode string `json:"qrcode"` // 支付二维码内容
		Scheme string `json:"scheme"` // 唤起支付的地址
	} `json:"data"`
}

// 检查订单支付状态
type OrderStatusReq struct {
	OrderId    string `json:"order_id"`
	OutOrderId string `json:"out_order_id"`
	Pid        int32  `json:"pid"`
}

// 订单状态
type OrderStatusRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status     int32  `json:"status"`
		ExpireTime int    `json:"expire_time"`
		ReturnUri  string `json:"return_uri,omitempty"`
	} `json:"data"`
}
