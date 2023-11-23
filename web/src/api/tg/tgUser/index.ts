import {jumpExport, tg} from '@/utils/http/axios';

// 获取TG账号列表
export function List(params: any) {
  return tg.request({
    url: '/tgUser/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除TG账号
export function Delete(params: any) {
  return tg.request({
    url: '/tgUser/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑TG账号
export function Edit(params: any) {
  return tg.request({
    url: '/tgUser/edit',
    method: 'POST',
    params,
  });
}

// 添加/编辑TG账号
export function ImportSession(params: any) {
  return tg.request({
    url: '/tgUser/importSession',
    method: 'POST',
    params,
  });
}

// 获取TG账号指定详情
export function View(params: any) {
  return tg.request({
    url: '/tgUser/view',
    method: 'GET',
    params,
  });
}

// 绑定员工
export function TgBindMember(params: any) {
  return tg.request({
    url: '/tgUser/bindMember',
    method: 'POST',
    params,
  });
}

// 解绑员工
export function TgUnBindMember(params: any) {
  return tg.request({
    url: '/tgUser/unBindMember',
    method: 'POST',
    params,
  });
}

// 绑定代理
export function TgBindProxy(params: any) {
  return tg.request({
    url: '/tgUser/bindProxy',
    method: 'POST',
    params,
  });
}

// 解绑代理
export function TgUnBindProxy(params: any) {
  return tg.request({
    url: '/tgUser/unBindProxy',
    method: 'POST',
    params,
  });
}

// 登录
export function TgLogin(params: any) {
  return tg.request({
    url: '/arts/login',
    method: 'POST',
    params,
  });
}

// 批量登录
export function TgBathLogin(params: any) {
  return tg.request({
    url: '/arts/batchLogin',
    method: 'POST',
    params,
  });
}

// 批量下线
export function TgBathLogout(params: any) {
  return tg.request({
    url: '/arts/batchLogout',
    method: 'POST',
    params,
  });
}


// 导出TG账号
export function Export(params: any) {
  jumpExport('/tgUser/export', params);
}

// 获取会话文件夹
export function TgGetFolders(params: any) {
  return tg.request({
    url: '/arts/folders',
    method: 'GET',
    params,
  });
}

// 获取会话列表
export function TgGetDialogs(params: any) {
  return tg.request({
    url: '/arts/getDialogs',
    method: 'POST',
    params,
  });
}

// 获取聊天记录
export function TgGetMsgHistory(params: any) {
  return tg.request({
    url: '/arts/getMsgHistory',
    method: 'POST',
    params,
  });
}

// 获取发送消息
export function TgSendMsg(params: any) {
  return tg.request({
    url: '/arts/sendMsg',
    method: 'POST',
    params,
  });
}

// 下载头像
export function TgGetUserAvatar(params: any) {
  return tg.request({
    url: '/arts/user/getUserAvatar',
    method: 'POST',
    params,
  })
}
