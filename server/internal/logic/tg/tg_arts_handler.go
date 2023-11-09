package tg

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/container/array"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
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
	}).Where("max_connections - assigned_count > 0").OrderAsc(dao.SysProxy.Columns().UpdatedAt).Scan(&result)

	if err != nil || g.IsEmpty(result) {
		err = gerror.New(g.I18n().T(ctx, "{#NoProxyContactAdministrator}"))
		return
	}
	return

}

func (s *sTgArts) handleProxy(ctx context.Context, tgUser *entity.TgUser) (err error) {
	// 查看账号是否有代理
	if tgUser.ProxyAddress != "" {
		return
	}
	//随机分配一个代理
	proxy, err := s.getRandomProxy(ctx)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.SysProxy.Ctx(ctx).WherePri(proxy.Id).Data(do.SysProxy{
			AssignedCount: gdb.Raw(fmt.Sprintf("%s + %d", dao.SysProxy.Columns().AssignedCount, 1)),
		}).Update()
		if err != nil {
			return
		}
		_, err = dao.TgUser.Ctx(ctx).WherePri(tgUser.Id).Data(do.TgUser{
			ProxyAddress: proxy.Address,
			PublicProxy:  1,
		}).Update()
		return
	})

	return
}

// 处理端口号
func (s *sTgArts) handlerPort(ctx context.Context, sysOrg entity.SysOrg, tgUser *entity.TgUser) (err error) {
	// 判断端口数是否足够
	if sysOrg.AssignedPorts+1 >= sysOrg.Ports {
		return gerror.New(g.I18n().T(ctx, "{#InsufficientNumber}"))
	}
	// 更新已使用端口数
	_, err = service.SysOrg().Model(ctx).
		WherePri(sysOrg.Id).
		Data(do.SysOrg{AssignedPorts: gdb.Raw(fmt.Sprintf("%s+%d", dao.SysOrg.Columns().AssignedPorts, 1))}).
		Update()
	// 记录占用端口的账号
	loginPorts := make(map[string]interface{})
	loginPorts[tgUser.Phone] = sysOrg.Id
	_, err = g.Redis().HSet(ctx, consts.TgLoginPorts, loginPorts)
	return
}

func (s *sTgArts) isLogin(ctx context.Context, tgUser *entity.TgUser) bool {
	req := &protobuf.RequestMessage{
		Action: protobuf.Action_GET_ONLINE_ACCOUNTS,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_GetOnlineAccountsDetail{
			GetOnlineAccountsDetail: &protobuf.GetOnlineAccountsDetail{
				Phone: []string{tgUser.Phone},
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return false
	}
	var onlineAccounts []tgin.OnlineAccountInp
	_ = gjson.DecodeTo(resp.Data, &onlineAccounts)
	if len(onlineAccounts) > 0 {
		return true
	}
	return false
}
