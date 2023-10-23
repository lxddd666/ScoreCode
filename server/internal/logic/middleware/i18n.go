package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	i18nList = []string{"zh-CN", "en"}
)

// I18n 国际化
func (s *sMiddleware) I18n(r *ghttp.Request) {
	language := r.GetHeader("Accept-Language")
	fmt.Println(language)
	for _, item := range i18nList {
		if gstr.Contains(language, item) {
			g.I18n().SetLanguage(item)
			r.Middleware.Next()
			return
		}
	}
	g.I18n().SetLanguage(i18nList[0])
	r.Middleware.Next()
}
