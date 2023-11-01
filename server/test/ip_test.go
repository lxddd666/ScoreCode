package test

import (
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
	"time"
)

var (
	ctx = gctx.New()
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestIp(t *testing.T) {
	httpCli := g.Client().Discovery(nil).Timeout(10 * time.Second)
	startTime := time.Now()
	resp, err := httpCli.Proxy("socks5://fans007:fans888@64.176.225.97:39000").Get(gctx.New(), "https://api.ip.sb/geoip")
	if err != nil {
		panic(err)
	}
	delay := uint16(time.Since(startTime) / time.Millisecond)
	fmt.Println(delay)
	g.DumpJson(resp.ReadAllString())

}

func TestSession(t *testing.T) {
	ss := "0x54B424480D5EEA4955CECCCDDE02B7E3263307722C2CEC175E9743D612575C24D34CD83A80696702B4FAE8BCFD35A1B76E97E7D7440FDA4A5B8C13F56431FE3ACCCFFAE3803D88FE4E5C23D343E1B0C27505D3B5E854FD01B0ABBC41152794C7CD5499D26E96EA581818F6390A30DC91727DA06616CD1BC87BCE59D680EBA7028D4B1CA091A06504E1BA9E0EC725D1C46D6B4B2CCCEE5394BCEE6457DB096C78526927F5DB09ED165F409EAA7BE29C265099E6C277705B1940F777BB91520DD8C9BF3D94FD14B8E1C92BEA0123EC84B364186900383CA9C602697DC0459E5ADB03D3A0ADBAF0D158213D902789F7ADBF5CDED94D9166946E65769E32D43D87B9"
	bytes, err := hex.DecodeString(ss)
	if err != nil {
		return
	}
	fmt.Println(bytes)
}
