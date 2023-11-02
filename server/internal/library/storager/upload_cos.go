// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

import (
	"bytes"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

// CosDrive 腾讯云cos驱动
type CosDrive struct {
}

// Upload 上传到腾讯云cos对象存储
func (d *CosDrive) Upload(ctx context.Context, file *FileMeta) (fullPath string, err error) {
	if config.CosPath == "" {
		err = gerror.New(g.I18n().T(ctx, "{#CosStorageDriver}"))
		return
	}

	URL, _ := url.Parse(config.CosBucketURL)
	client := cos.NewClient(&cos.BaseURL{BucketURL: URL}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.CosSecretId,
			SecretKey: config.CosSecretKey,
		},
	})

	fullPath = GenFullPath(config.CosPath, gfile.Ext(file.Filename))
	reader := bytes.NewReader(file.Content)
	_, err = client.Object.Put(ctx, fullPath, reader, nil)
	return
}
