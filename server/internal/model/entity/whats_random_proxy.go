package entity

// WhatsRandomProxy is the golang structure for redis .
type WhatsRandomProxy struct {
	ProxyAddress   string `json:"proxyAddress" description:"代理地址"`
	AvailableCount int32  `json:"availableCount" description:"可用数量"`
}
