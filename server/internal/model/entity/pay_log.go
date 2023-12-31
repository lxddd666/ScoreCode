// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayLog is the golang structure for table pay_log.
type PayLog struct {
	Id            int64       `json:"id"            description:"主键"`
	MemberId      int64       `json:"memberId"      description:"会员ID"`
	AppId         string      `json:"appId"         description:"应用ID"`
	AddonsName    string      `json:"addonsName"    description:"插件名称"`
	OrderSn       string      `json:"orderSn"       description:"关联订单号"`
	OrderGroup    string      `json:"orderGroup"    description:"组别[默认统一支付类型]"`
	Openid        string      `json:"openid"        description:"openid"`
	MchId         string      `json:"mchId"         description:"商户支付账户"`
	Subject       string      `json:"subject"       description:"订单标题"`
	Detail        *gjson.Json `json:"detail"        description:"支付商品详情"`
	AuthCode      string      `json:"authCode"      description:"刷卡码"`
	OutTradeNo    string      `json:"outTradeNo"    description:"商户订单号"`
	TransactionId string      `json:"transactionId" description:"交易号"`
	PayType       string      `json:"payType"       description:"支付类型"`
	PayAmount     float64     `json:"payAmount"     description:"支付金额"`
	ActualAmount  float64     `json:"actualAmount"  description:"实付金额"`
	PayStatus     int         `json:"payStatus"     description:"支付状态"`
	PayAt         *gtime.Time `json:"payAt"         description:"支付时间"`
	TradeType     string      `json:"tradeType"     description:"交易类型"`
	RefundSn      string      `json:"refundSn"      description:"退款单号"`
	IsRefund      int         `json:"isRefund"      description:"是否退款"`
	Custom        string      `json:"custom"        description:"自定义参数"`
	CreateIp      string      `json:"createIp"      description:"创建者IP"`
	PayIp         string      `json:"payIp"         description:"支付者IP"`
	NotifyUrl     string      `json:"notifyUrl"     description:"支付通知回调地址"`
	ReturnUrl     string      `json:"returnUrl"     description:"买家付款成功跳转地址"`
	TraceIds      *gjson.Json `json:"traceIds"      description:"链路ID集合"`
	Status        int         `json:"status"        description:"状态"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"修改时间"`
}
