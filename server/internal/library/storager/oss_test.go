package storager

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"testing"
)

func TestOss(t *testing.T) {
	client, err := oss.New("http://oss-ap-southeast-1.aliyuncs.com", "LTAI5t7aFWwdbZpsP5JWFVty", "wtH8LIVdNsymsuirE3wgXgcFqC3y4s")
	if err != nil {
		return
	}

	bucket, err := client.Bucket("tgcloud")
	if err != nil {
		return
	}
	obj, err := bucket.GetObject("tgcloud/attachment/2023-10-08/cw2yu0gxl0rs23viro.png")
	defer obj.Close()
	buf := bytes.Buffer{}
	_, err = io.Copy(&buf, obj)
	fmt.Println(buf)
}
