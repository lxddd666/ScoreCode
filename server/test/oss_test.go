package test

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"testing"
)

func TestOss(t *testing.T) {
	client, err := oss.New("http://oss-ap-southeast-1.aliyuncs.com", "LTAI5t5dNpLRuMRxnaQRQRCR", "KAadavtt86IKUFRRKuiLHpwvhksmZ8")
	if err != nil {
		return
	}
	bucket, err := client.Bucket("tgcloud")
	if err != nil {
		return
	}
	obj, err := bucket.GetObject("tgcloud/attachment/2023-11-01/cwnbg79moj4gblcd1p.ico")
	defer obj.Close()
	buf := bytes.Buffer{}
	_, err = io.Copy(&buf, obj)
	fmt.Println(buf)
}
