package test

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/grand"
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

func TestAvatar(t *testing.T) {
	content := g.Client().GetBytes(ctx, "https://api.vvhan.com/api/avatar")
	mime := mimetype.Detect(content)
	fmt.Println(mime)
	gfile.PutBytes(grand.S(12)+mime.Extension(), content)
}
