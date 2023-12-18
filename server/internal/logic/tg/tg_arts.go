package tg

import (
	"context"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
	"hotgo/internal/consts"
	"hotgo/internal/core/prometheus"
	"hotgo/internal/dao"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
)

type sTgArts struct{}

func NewTgArts() *sTgArts {
	return &sTgArts{}
}

func init() {
	service.RegisterTgArts(NewTgArts())
}

// TgSyncContact 同步联系人
func (s *sTgArts) TgSyncContact(ctx context.Context, inp *artsin.SyncContactInp) (res string, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	res, err = service.Arts().SyncContact(ctx, inp, consts.TgSvc)
	if err == nil {
		prometheus.InitiateSyncContactCount.WithLabelValues(gconv.String(inp.Account)).Inc()
		for _, contact := range inp.Contacts {
			prometheus.PassiveSyncContactCount.WithLabelValues(gconv.String(contact)).Inc()
		}
	}
	return
}

// TgGetContacts 获取contacts
func (s *sTgArts) TgGetContacts(ctx context.Context, account uint64) (list []*tgin.TgContactsListModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, account); err != nil {
		return
	}

	msg := &protobuf.GetContactList{
		Account: account,
	}

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_CONTACT_LIST,
		Type:    consts.TgSvc,
		Account: account,
		ActionDetail: &protobuf.RequestMessage_GetContactList{
			GetContactList: msg,
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	if resp.ActionResult == protobuf.ActionResult_ALL_SUCCESS {
		prometheus.AccountGetContactsCount.WithLabelValues(gconv.String(account)).Inc()
		err = gjson.DecodeTo(resp.Data, &list)
		if err == nil {
			s.handlerSaveContacts(ctx, account, list)
		}
	}

	return
}

func (s *sTgArts) handlerSaveContacts(ctx context.Context, account uint64, list []*tgin.TgContactsListModel) {
	in := make(map[uint64][]*tgin.TgContactsListModel)
	in[account] = list
	_ = service.TgContacts().SyncContactCallback(ctx, in)
}

// TgDownloadFile 下载聊天文件
func (s *sTgArts) TgDownloadFile(ctx context.Context, inp *tgin.TgDownloadMsgInp) (res *tgin.TgDownloadMsgModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	msgMap := make(map[uint64]*protobuf.DownLoadFileMsg)
	msgMap[inp.Account] = &protobuf.DownLoadFileMsg{
		ChatId:    inp.ChatId,
		MessageId: inp.MsgId,
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_DOWNLOAD_FILE,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_GetDownLoadFileDetail{
			GetDownLoadFileDetail: &protobuf.GetDownLoadFileDetail{
				DownloadFile: msgMap,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	prometheus.AccountDownloadFileCount.WithLabelValues(gconv.String(inp.Account)).Inc()
	var data callback.MsgCallbackRes
	err = gjson.DecodeTo(resp.Data, &data)
	if err != nil {
		return
	}
	mime := mimetype.Detect(data.SendMsg)
	var meta = &storager.FileMeta{
		Filename: data.FileName,
		Size:     gconv.Int64(len(data.SendMsg)),
		MimeType: mime.String(),
		Ext:      storager.Ext(data.FileName),
		Md5:      gmd5.MustEncryptBytes(data.SendMsg),
		Content:  data.SendMsg,
	}
	meta.Kind = storager.GetFileKind(meta.Ext)
	result, err := service.CommonUpload().UploadFile(ctx, storager.KindOther, meta)
	if err != nil {
		return
	}
	res = new(tgin.TgDownloadMsgModel)
	res.AttachmentListModel = result
	res.Account = inp.Account
	res.MsgId = inp.MsgId
	res.ChatId = inp.ChatId
	return

}

// TgAddGroupMembers 添加群成员
func (s *sTgArts) TgAddGroupMembers(ctx context.Context, inp *tgin.TgGroupAddMembersInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_ADD_GROUP_MEMBER,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_AddGroupMemberDetail{
			AddGroupMemberDetail: &protobuf.AddGroupMemberDetail{
				Detail: &protobuf.UintkeyStringvalue{
					Key:    inp.Account,
					Values: inp.AddMembers,
				},
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err == nil {
		prometheus.AddMemberToGroupCount.WithLabelValues(inp.GroupId).Inc()
		for _, member := range inp.AddMembers {
			prometheus.AccountJoinGroupCount.WithLabelValues(member).Inc()
		}
	}
	return
}

// TgCreateGroup 创建群聊
func (s *sTgArts) TgCreateGroup(ctx context.Context, inp *tgin.TgCreateGroupInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_CREATE_GROUP,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_CreateGroupDetail{
			CreateGroupDetail: &protobuf.CreateGroupDetail{
				GroupName: inp.GroupTitle,
				Detail: &protobuf.UintkeyStringvalue{
					Key:    inp.Account,
					Values: inp.AddMembers,
				},
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err == nil {
		prometheus.CreateGroupCount.WithLabelValues(gconv.String(inp.Account)).Inc()
	}
	return
}

// TgGetGroupMembers 获取群成员
func (s *sTgArts) TgGetGroupMembers(ctx context.Context, inp *tgin.TgGetGroupMembersInp) (list []*tgin.TgContactsListModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_GET_GROUP_MEMBERS,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_GetGroupMembersDetail{
			GetGroupMembersDetail: &protobuf.GetGroupMembersDetail{
				Account: inp.Account,
				ChatId:  inp.GroupId,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	prometheus.AccountGetGroupMsgCount.WithLabelValues(gconv.String(inp.Account)).Inc()

	if err != nil {
		return
	}
	err = gjson.DecodeTo(resp.Data, &list)
	return
}

// TgCreateChannel 创建频道
func (s *sTgArts) TgCreateChannel(ctx context.Context, inp *tgin.TgChannelCreateInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_CREATE_CHANNEL,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_CreateChannelDetail{
			CreateChannelDetail: &protobuf.CreateChannelDetail{
				ChannelTitle:       inp.Title,
				ChannelUserName:    inp.UserName,
				ChannelDescription: inp.Description,
				Detail: &protobuf.UintkeyStringvalue{
					Key:    inp.Account,
					Values: inp.Members,
				},
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err == nil {
		prometheus.CreateChannelCount.WithLabelValues(gconv.String(inp.Account)).Inc()
	}
	return
}

// TgChannelAddMembers 添加频道成员
func (s *sTgArts) TgChannelAddMembers(ctx context.Context, inp *tgin.TgChannelAddMembersInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_INVITE_TO_CHANNEL,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_InviteToChannelDetail{
			InviteToChannelDetail: &protobuf.InviteToChannelDetail{
				Channel: inp.Channel,
				Detail: &protobuf.UintkeyStringvalue{
					Key:    inp.Account,
					Values: inp.Members,
				},
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err == nil {
		prometheus.AddMemberToChannelCount.WithLabelValues(gconv.String(inp.Channel)).Inc()
		for _, member := range inp.Members {
			prometheus.AccountJoinChannelCount.WithLabelValues(gconv.String(member)).Inc()
		}
	}
	return
}

// TgChannelJoinByLink 加入频道
func (s *sTgArts) TgChannelJoinByLink(ctx context.Context, inp *tgin.TgChannelJoinByLinkInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_JOIN_BY_LINK,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_JoinByLinkDetail{
			JoinByLinkDetail: &protobuf.JoinByLinkDetail{
				Detail: &protobuf.UintkeyStringvalue{
					Key:    inp.Account,
					Values: inp.Link,
				},
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err == nil {
		prometheus.AccountJoinChannelCount.WithLabelValues(gconv.String(inp.Account)).Inc()
		contactList := []*tgin.TgContactsListModel{}
		err = gjson.DecodeTo(resp.Data, &contactList)
		if err == nil {
			param := make(map[uint64][]*tgin.TgContactsListModel)
			param[inp.Account] = contactList
			_ = service.TgContacts().SyncContactCallback(ctx, param)
		}
	}
	return
}

// TgGetUserAvatar 获取用户头像
func (s *sTgArts) TgGetUserAvatar(ctx context.Context, inp *tgin.TgGetUserAvatarInp) (res *tgin.TgGetUserAvatarModel, err error) {
	var tgPhoto *entity.TgPhoto
	err = dao.TgPhoto.Ctx(ctx).Where(do.TgPhoto{PhotoId: inp.PhotoId, TgId: inp.GetUser}).Scan(&tgPhoto)
	if err != nil {
		return
	}
	if tgPhoto != nil {
		res = &tgin.TgGetUserAvatarModel{
			Avatar: tgPhoto.FileUrl,
		}
		return
	}
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_GET_USER_HEAD_IMAGE,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_DownUserHeadImageDetail{
			DownUserHeadImageDetail: &protobuf.DownUserHeadImageDetail{
				Account: inp.Account,
				GetUser: inp.GetUser,
				PhotoId: inp.PhotoId,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	prometheus.AccountGetUserHeadImageCount.WithLabelValues(gconv.String(inp.Account)).Inc()
	mime := mimetype.Detect(resp.Data)
	var meta = &storager.FileMeta{
		Filename: gconv.String(inp.PhotoId) + mime.Extension(),
		Size:     gconv.Int64(len(resp.Data)),
		MimeType: mime.String(),
		Ext:      mime.Extension()[1:],
		Md5:      gmd5.MustEncryptBytes(resp.Data),
		Content:  resp.Data,
	}
	meta.Kind = storager.GetFileKind(meta.Ext)
	result, err := service.CommonUpload().UploadFile(ctx, storager.KindOther, meta)
	if err != nil {
		return
	}
	_, err = dao.TgPhoto.Ctx(ctx).Save(entity.TgPhoto{
		TgId:         int64(inp.GetUser),
		PhotoId:      inp.PhotoId,
		AttachmentId: result.Id,
		Path:         result.Path,
		FileUrl:      result.FileUrl,
	})
	if err != nil {
		return
	}
	res = &tgin.TgGetUserAvatarModel{
		Avatar: result.FileUrl,
	}
	return
}

func (s *sTgArts) TgGetSearchInfo(ctx context.Context, inp *tgin.TgGetSearchInfoInp) (res []*tgin.TgGetSearchInfoModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Sender); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_SEARCH,
		Type:    consts.TgSvc,
		Account: inp.Sender,
		ActionDetail: &protobuf.RequestMessage_SearchDetail{
			SearchDetail: &protobuf.SearchDetail{
				Sender: inp.Sender,
				Search: inp.Search,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	prometheus.AccountSearchInfoCount.WithLabelValues(gconv.String(inp.Sender)).Inc()
	err = gjson.DecodeTo(resp.Data, &res)
	if err != nil {
		return
	}
	return
}

// TgReadPeerHistory 消息已读
func (s *sTgArts) TgReadPeerHistory(ctx context.Context, inp *tgin.TgReadPeerHistoryInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Sender); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Account: inp.Sender,
		Action:  protobuf.Action_READ_HISTORY,
		Type:    consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_ReadHistoryDetail{
			ReadHistoryDetail: &protobuf.ReadHistoryDetail{
				Account:  inp.Sender,
				Receiver: inp.Receiver,
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	prometheus.AccountReadMsgHistoryCount.WithLabelValues(gconv.String(inp.Sender)).Inc()
	prometheus.AccountMsgPassiveReadHistoryCount.WithLabelValues(gconv.String(inp.Receiver)).Inc()
	return
}

// TgReadChannelHistory channel消息已读
func (s *sTgArts) TgReadChannelHistory(ctx context.Context, inp *tgin.TgReadChannelHistoryInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Sender); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Account: inp.Sender,
		Action:  protobuf.Action_READ_CHANNEL_HISTORY,
		Type:    consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_ReadChannelHistoryAction{
			ReadChannelHistoryAction: &protobuf.ReadChannelHistoryDetail{
				Account:  inp.Sender,
				Receiver: inp.Receiver,
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	prometheus.AccountReadMsgHistoryCount.WithLabelValues(gconv.String(inp.Sender)).Inc()
	prometheus.AccountChannelReadHistoryCount.WithLabelValues(gconv.String(inp.Receiver)).Inc()

	return
}

// TgChannelReadAddView channel view++
func (s *sTgArts) TgChannelReadAddView(ctx context.Context, inp *tgin.ChannelReadAddViewInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Sender); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Account: inp.Sender,
		Action:  protobuf.Action_CHANNEL_READ_VIEW,
		Type:    consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_ChannelReadViewDetail{
			ChannelReadViewDetail: &protobuf.ChannelReadViewDetail{
				Account:  inp.Sender,
				Receiver: inp.Receiver,
				MsgIds:   inp.MsgIds,
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	return
}

// TgUpdateUserInfo 修改用户信息
func (s *sTgArts) TgUpdateUserInfo(ctx context.Context, inp *tgin.TgUpdateUserInfoInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	sendMsg := make(map[uint64]*protobuf.UpdateUserInfoMsg)
	sendMsg[inp.Account] = &protobuf.UpdateUserInfoMsg{
		UserName:  inp.Username,
		FirstName: inp.FirstName,
		LastName:  inp.LastName,
		Bio:       inp.Bio,
		Photo: &protobuf.FileDetailValue{
			SendType: consts.SendTypeByte,
			Path:     inp.Photo.Name,
			Name:     inp.Photo.Name,
			FileByte: inp.Photo.Data,
		},
	}

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_UPDATE_USER_INFO,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_UpdateUserInfoDetail{
			UpdateUserInfoDetail: &protobuf.UpdateUserInfoDetail{
				Account:  inp.Account,
				SendData: sendMsg,
			},
		},
	}

	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	prometheus.AccountUpdateUserInfoCount.WithLabelValues(gconv.String(inp.Account)).Inc()
	var userFull tg.UsersUserFull
	err = (&bin.Buffer{Buf: resp.Data}).Decode(&userFull)
	if err == nil {
		if user, b := userFull.Users[0].AsNotEmpty(); b {
			updateMap := do.TgUser{
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Username:  user.Username,
				Bio:       userFull.FullUser.About,
			}
			if user.Photo != nil {
				photo, pFlag := user.Photo.AsNotEmpty()
				if pFlag {
					updateMap.Photo = photo.PhotoID
				}
			}
			_, err = dao.TgUser.Ctx(ctx).Data(updateMap).OmitEmpty().Where(dao.TgUser.Columns().Phone, inp.Account).Update()
		}

	}
	return
}

// TgCheckUsername 校验username
func (s *sTgArts) TgCheckUsername(ctx context.Context, inp *tgin.TgCheckUsernameInp) (flag bool, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Account: inp.Account,
		Action:  protobuf.Action_CHECK_USERNAME,
		Type:    consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_CheckUsernameDetail{
			CheckUsernameDetail: &protobuf.CheckUserNameDetail{
				Account:  inp.Account,
				Username: inp.Username,
			},
		},
	}
	flag = false
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	if resp.RespondAccountStatus == protobuf.AccountStatus_SUCCESS {
		flag = true
	}
	return
}

func (s *sTgArts) TgLeaveGroup(ctx context.Context, inp *tgin.TgUserLeaveInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	detail := &protobuf.UintkeyStringvalue{}
	detail.Key = inp.Account

	detail.Values = append(detail.Values, inp.TgId)

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_LEAVE,
		Type:    consts.TgSvc,
		Account: detail.Key,
		ActionDetail: &protobuf.RequestMessage_LeaveDetail{
			LeaveDetail: &protobuf.LeaveDetail{
				Detail: detail,
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	prometheus.AccountLeaveGroupCount.WithLabelValues(gconv.String(inp.Account)).Inc()
	prometheus.GroupLeaveGroupCount.WithLabelValues(gconv.String(inp.TgId)).Inc()
	return
}

// ContactsGetLocated 获取附近的人
func (s *sTgArts) ContactsGetLocated(ctx context.Context, inp *tgin.ContactsGetLocatedInp) (err error) {
	if err = s.TgCheckLogin(ctx, inp.Sender); err != nil {
		return
	}

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_CONTACTS_GET_LOCATED,
		Type:    consts.TgSvc,
		Account: inp.Sender,
		ActionDetail: &protobuf.RequestMessage_ContactsGetLocatedDetail{
			ContactsGetLocatedDetail: &protobuf.ContactsGetLocatedDetail{
				Sender:         inp.Sender,
				Background:     inp.Background,
				SelfExpires:    uint64(inp.SelfExpires),
				Lat:            inp.Lat,
				Long:           inp.Long,
				AccuracyRadius: uint64(inp.AccuracyRadius),
			},
		},
	}
	res, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	prometheus.AccountClearMsgDraft.WithLabelValues(gconv.String(inp.Sender)).Inc()
	//err = gjson.DecodeTo(resp.Data, &res)
	//if err != nil {
	//	return
	//}
	model := tgin.TgContactsGetLocatedModel{}
	err = gjson.DecodeTo(res.Data, &model)
	if err != nil {
		return
	}
	var box tg.Updates
	err = (&bin.Buffer{Buf: model.ResultBuf}).Decode(&box)
	handlerLocateContact(box)
	return
}

func handlerLocateContact(box tg.Updates) (list []tgin.TgPeerModel) {
	return
}

// EditChannelInfo 修改频道信息
func (s *sTgArts) EditChannelInfo(ctx context.Context, inp *tgin.EditChannelInfoInp) (err error) {
	if err = s.TgCheckLogin(ctx, inp.Sender); err != nil {
		return
	}

	photo := &protobuf.FileDetailValue{
		FileType: inp.Photo.MIME,
		SendType: consts.SendTypeByte,
		Name:     inp.Photo.Name,
		FileByte: gbase64.MustDecode(inp.Photo.Data),
	}

	geo := &protobuf.GeoPointValue{
		Lat:            inp.GeoPointType.Lat,
		Long:           inp.GeoPointType.Long,
		AccuracyRadius: uint64(inp.GeoPointType.AccuracyRadius),
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_EDIT_CHANNEL_INFO,
		Type:    consts.TgSvc,
		Account: inp.Sender,
		ActionDetail: &protobuf.RequestMessage_EditChannelInfoDetail{
			EditChannelInfoDetail: &protobuf.EditChannelInfoDetail{
				Sender:   inp.Sender,
				Channel:  inp.Channel,
				Photo:    photo,
				Title:    inp.Title,
				GeoPoint: geo,
				Address:  inp.Address,
				Describe: inp.Describe,
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	prometheus.AccountEditChannelInfo.WithLabelValues(gconv.String(inp.Sender)).Inc()
	prometheus.ChannelEditInfo.WithLabelValues(gconv.String(inp.Channel)).Inc()

	return
}

// EditChannelInfo 获取附近的人
func (s *sTgArts) EditChannelBannedRight(ctx context.Context, inp *tgin.EditChannelBannedRightsInp) (err error) {
	if err = s.TgCheckLogin(ctx, inp.Sender); err != nil {
		return
	}

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_EDIT_CHAT_BANNED_RIGHTS,
		Type:    consts.TgSvc,
		Account: inp.Sender,
		ActionDetail: &protobuf.RequestMessage_EditChatBannedRightsDetail{
			EditChatBannedRightsDetail: &protobuf.EditChatBannedRightsDetail{
				Sender:  inp.Sender,
				Channel: inp.Channel,
				ChatBannedRights: &protobuf.ChatBannedRightsType{
					ViewMessages:    inp.BannedRights.ViewMessages,
					SendMessages:    inp.BannedRights.SendMessages,
					SendMedia:       inp.BannedRights.SendMedia,
					SendStickers:    inp.BannedRights.SendStickers,
					SendGifs:        inp.BannedRights.SendGifs,
					SendGames:       inp.BannedRights.SendGames,
					SendInline:      inp.BannedRights.SendInline,
					EmbedLinks:      inp.BannedRights.EmbedLinks,
					SendPolls:       inp.BannedRights.SendPolls,
					ChangeInfo:      inp.BannedRights.ChangeInfo,
					InviteUsers:     inp.BannedRights.InviteUsers,
					PinMessages:     inp.BannedRights.PinMessages,
					ManageTopics:    inp.BannedRights.ManageTopics,
					SendPhotos:      inp.BannedRights.SendPhotos,
					SendVideos:      inp.BannedRights.SendVideos,
					SendRoundVideos: inp.BannedRights.SendStickers,
					SendAudios:      inp.BannedRights.SendAudios,
					SendVoices:      inp.BannedRights.SendVoices,
					SendDocs:        inp.BannedRights.SendDocs,
					SendPlain:       inp.BannedRights.SendPlain,
					UntilDate:       int64(inp.BannedRights.UntilDate),
				},
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	return
}

// GetManageChannels 获取自己管理的频道和群
func (s *sTgArts) GetManageChannels(ctx context.Context, inp *tgin.GetManageChannelsInp) (err error) {
	if err = s.TgCheckLogin(ctx, inp.Sender); err != nil {
		return
	}

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_GET_MANAGED_CHANNELS,
		Type:    "telegram",
		Account: inp.Sender,
		ActionDetail: &protobuf.RequestMessage_GetManageChannelsDetail{
			GetManageChannelsDetail: &protobuf.GetManageChannelsDetail{
				Sender:     inp.Sender,
				ByLocation: inp.ByLocation,
				CheckLimit: false, // 默认为false就行了
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err != nil {
		return
	}

	return
}
