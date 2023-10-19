// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

import (
	"bytes"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
)

// OssDrive 阿里云oss驱动
type OssDrive struct {
}

// Upload 上传到阿里云oss
func (d *OssDrive) Upload(ctx context.Context, file *FileMeta) (fullPath string, err error) {
	if config.OssPath == "" {
		err = gerror.New("OSS存储驱动必须配置存储路径!")
		return
	}

	client, err := oss.New(config.OssEndpoint, config.OssSecretId, config.OssSecretKey)
	if err != nil {
		return
	}

	bucket, err := client.Bucket(config.OssBucket)
	if err != nil {
		return
	}

	fullPath = GenFullPath(config.OssPath, gfile.Ext(file.Filename))
	reader := bytes.NewReader(file.Content)
	err = bucket.PutObject(fullPath, reader)
	return
}
