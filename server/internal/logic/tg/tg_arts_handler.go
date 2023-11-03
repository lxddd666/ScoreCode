package tg

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/container/array"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
)

func (s *sTgArts) handlerRandomProxy(ctx context.Context, notAccounts *array.Array[*entity.TgUser]) (err error, result *array.Array[*entity.TgUser]) {

	proxy, err := s.getRandomProxy(ctx)

	forNum := 0
	num := gconv.Int(proxy.MaxConnections - proxy.AssignedCount)
	tgUserIds := make([]uint64, 0)
	if num > notAccounts.Len() {
		// 可分配数量超过 登录数量
		forNum = notAccounts.Len()
		for i := 0; i < forNum; i++ {
			v, _ := notAccounts.Get(i)
			v.ProxyAddress = proxy.Address
			tgUserIds = append(tgUserIds, v.Id)
		}
		result = notAccounts
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			_, err = dao.SysProxy.Ctx(ctx).WherePri(proxy.Id).Data(do.SysProxy{
				AssignedCount: gdb.Raw(fmt.Sprintf("%s + %d", dao.SysProxy.Columns().AssignedCount, forNum)),
			}).Update()
			if err != nil {
				return
			}
			_, err = dao.TgUser.Ctx(ctx).WhereIn(dao.TgUser.Columns().Id, tgUserIds).Data(do.TgUser{
				ProxyAddress: proxy.Address,
				PublicProxy:  1,
			}).Update()
			return
		})

		if err != nil {
			return
		}
	} else if num == notAccounts.Len() {
		// 可分配数量超过 登录数量
		forNum = num
		for i := 0; i < forNum; i++ {
			v, _ := notAccounts.Get(i)
			v.ProxyAddress = proxy.Address
			tgUserIds = append(tgUserIds, v.Id)
		}
		result = notAccounts
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			_, err = dao.SysProxy.Ctx(ctx).WherePri(proxy.Id).Data(do.SysProxy{
				AssignedCount: gdb.Raw(fmt.Sprintf("%s + %d", dao.SysProxy.Columns().AssignedCount, forNum)),
			}).Update()
			if err != nil {
				return
			}
			_, err = dao.TgUser.Ctx(ctx).WhereIn(dao.TgUser.Columns().Id, tgUserIds).Data(do.TgUser{
				ProxyAddress: proxy.Address,
				PublicProxy:  1,
			}).Update()
			return
		})

		if err != nil {
			return
		}
	} else {
		//此时需要拆分数组
		forNum = num
		for i := 0; i < forNum; i++ {
			v, _ := notAccounts.Get(i)
			v.ProxyAddress = proxy.Address
			tgUserIds = append(tgUserIds, v.Id)
		}
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			_, err = dao.SysProxy.Ctx(ctx).WherePri(proxy.Id).Data(do.SysProxy{
				AssignedCount: gdb.Raw(fmt.Sprintf("%s + %d", dao.SysProxy.Columns().AssignedCount, forNum)),
			}).Update()
			if err != nil {
				return
			}
			_, err = dao.TgUser.Ctx(ctx).WhereIn(dao.TgUser.Columns().Id, tgUserIds).Data(do.TgUser{
				ProxyAddress: proxy.Address,
				PublicProxy:  1,
			}).Update()
			return
		})
		if err != nil {
			return
		}
		newArray := array.NewFrom(notAccounts.Slice()[:forNum], true)
		err, a := s.handlerRandomProxy(ctx, array.NewArrayFrom(notAccounts.Slice()[forNum:]))
		if err != nil {
			return err, nil
		}
		newArray.Merge(a.Slice())
		result = newArray
	}

	return
}

// 获取代理
func (s *sTgArts) getRandomProxy(ctx context.Context) (result entity.SysProxy, err error) {
	err = dao.SysProxy.Ctx(ctx).Where(do.SysProxy{
		OrgId:  1,
		Status: 1,
		Type:   "socks5",
	}).Where("max_connections - assigned_count > 0").Scan(&result)

	if err != nil || g.IsEmpty(result) {
		err = gerror.New(g.I18n().T(ctx, "{#NoProxyContactAdministrator}"))
		return
	}
	return

}
