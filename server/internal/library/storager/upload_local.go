// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"strings"
)

// LocalDrive 本地驱动
type LocalDrive struct {
}

// Upload 上传到本地
func (d *LocalDrive) Upload(ctx context.Context, file *FileMeta) (fullPath string, err error) {
	var (
		sp      = g.Cfg().MustGet(ctx, "server.serverRoot")
		nowDate = gtime.Date()
	)

	if sp.IsEmpty() {
		err = gerror.New(g.I18n().T(ctx, "{#LocalUploadStaticPath}"))
		return
	}

	if config.LocalPath == "" {
		err = gerror.New(g.I18n().T(ctx, "{#LocalUploadStoragePath}"))
		return
	}

	// 包含静态文件夹的路径
	fullDirPath := strings.Trim(sp.String(), "/") + "/" + config.LocalPath + nowDate + "/" + file.Filename
	err = gfile.PutBytes(fullDirPath, file.Content)
	if err != nil {
		return
	}
	// 不含静态文件夹的路径
	fullPath = config.LocalPath + nowDate + "/" + file.Filename
	return
}
