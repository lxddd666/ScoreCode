// Package excel
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package excel

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model"
	"net/url"
	"reflect"
	"time"
)

var (
	// 单元格表头
	char = []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	// 默认行样式
	defaultRowStyle = excelize.Style{
		Font: &excelize.Font{
			Family: "arial",
			Size:   13,
			Color:  "#666666",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}
)

// ExportByStructs 导出切片结构体到excel表格
func ExportByStructs(ctx context.Context, tags []string, list interface{}, fileName string) (err error) {
	f := excelize.NewFile()
	err = f.SetSheetName("Sheet1", "Sheet1")
	if err != nil {
		return err
	}
	_ = f.SetRowHeight("Sheet1", 1, 30)

	rowStyleID, _ := f.NewStyle(&defaultRowStyle)
	if err != nil {
		return
	}
	_ = f.SetSheetRow("Sheet1", "A1", &tags)

	var (
		length    = len(tags)
		headStyle = letter(length)
		lastRow   string
		widthRow  string
	)

	for k, v := range headStyle {
		if k == length-1 {
			lastRow = fmt.Sprintf("%s1", v)
			widthRow = v
		}
	}
	if err = f.SetColWidth("Sheet1", "A", widthRow, 30); err != nil {
		return err
	}

	var rowNum = 1
	for _, v := range gconv.Interfaces(list) {
		t := reflect.TypeOf(v)
		value := reflect.ValueOf(v)
		row := make([]interface{}, 0)
		for l := 0; l < t.NumField(); l++ {
			val := value.Field(l).Interface()
			row = append(row, val)
		}
		rowNum++
		if err = f.SetSheetRow("Sheet1", "A"+gconv.String(rowNum), &row); err != nil {
			return
		}
		if err = f.SetCellStyle("Sheet1", fmt.Sprintf("A%d", rowNum), lastRow, rowStyleID); err != nil {
			return
		}
	}

	// 强转类型
	writer := ghttp.RequestFromCtx(ctx).Response.Writer
	w, ok := interface{}(writer).(*ghttp.ResponseWriter)
	if !ok {
		return gerror.New("Response.Writer uninitialized!")
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", url.QueryEscape(fileName)))
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	if err = f.Write(w); err != nil {
		return
	}

	// 加入到上下文
	contexts.SetResponse(ctx, &model.Response{
		Code:      gcode.CodeOK.Code(),
		Message:   "export successfully!",
		Timestamp: time.Now().Unix(),
		TraceID:   gctx.CtxId(ctx),
	})
	return
}

// letter 生成完整的表头
func letter(length int) []string {
	var str []string
	for i := 0; i < length; i++ {
		str = append(str, numToChars(i))
	}
	return str
}

// numToChars 将数字转换为具体的表格表头名称
func numToChars(num int) string {
	var cols string
	v := num
	for v > 0 {
		k := v % 26
		if k == 0 {
			k = 26
		}
		v = (v - k) / 26
		cols = char[k] + cols
	}
	return cols
}
