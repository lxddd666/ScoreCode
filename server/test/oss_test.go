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

	b, err := getFileFromOSSAndConvertToBytes("http://tgcloud/attachment/2023-11-02/cwo9cw61cspc5x0cmf.jpeg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}

func getFileFromOSSAndConvertToBytes(url string) ([]byte, error) {

	client, err := oss.New("http://oss-ap-southeast-1.aliyuncs.com", "LTAI5t7aFWwdbZpsP5JWFVty", "wtH8LIVdNsymsuirE3wgXgcFqC3y4s")
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket("tgcloud")
	if err != nil {
		return nil, err
	}
	obj, err := bucket.GetObject(url)
	if obj != nil {
		defer obj.Close()
	}
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	_, err = io.Copy(&buf, obj)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func TestAvatar(t *testing.T) {
	content := g.Client().GetBytes(ctx, "https://api.vvhan.com/api/avatar")
	mime := mimetype.Detect(content)
	fmt.Println(mime)
	gfile.PutBytes(grand.S(12)+mime.Extension(), content)
}
