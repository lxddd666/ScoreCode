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
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// QiNiuDrive 七牛云对象存储驱动
type QiNiuDrive struct {
}

// Upload 上传到七牛云对象存储
func (d *QiNiuDrive) Upload(ctx context.Context, file *FileMeta) (fullPath string, err error) {
	if config.QiNiuPath == "" {
		err = gerror.New(g.I18n().T(ctx, "{#QiNiuCloud}"))
		return
	}

	putPolicy := storage.PutPolicy{
		Scope: config.QiNiuBucket,
	}
	token := putPolicy.UploadToken(qbox.NewMac(config.QiNiuAccessKey, config.QiNiuSecretKey))

	cfg := storage.Config{}

	// 是否使用https域名
	cfg.UseHTTPS = true

	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 空间对应的机房
	cfg.Region, err = storage.GetRegion(config.QiNiuAccessKey, config.QiNiuBucket)
	if err != nil {
		return
	}

	fullPath = GenFullPath(config.QiNiuPath, gfile.Ext(file.Filename))
	err = storage.NewFormUploader(&cfg).Put(ctx, &storage.PutRet{}, token, fullPath, bytes.NewReader(file.Content), file.Size, &storage.PutExtra{})
	return
}
