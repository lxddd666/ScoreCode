import { whats } from '@/utils/http/axios';

// 获取账号管理列表
export function List(params: any) {
  return whats.request({
    url: '/whatsAccount/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除账号管理
export function Delete(params: any) {
  return whats.request({
    url: '/whatsAccount/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑账号管理
export function Edit(params: any) {
  return whats.request({
    url: '/whatsAccount/edit',
    method: 'POST',
    params,
  });
}

// 获取账号管理指定详情
export function View(params: any) {
  return whats.request({
    url: '/whatsAccount/view',
    method: 'GET',
    params,
  });
}

// 上传账号
export function Upload(params: any) {
  return whats.request({
    url: '/whatsAccount/upload',
    method: 'POST',
    params,
  });
}

// 登录/批量登录账号
export function Login(params: any) {
  return whats.request({
    url: '/whats/login',
    method: 'POST',
    params,
  });
}

// 退出登录/批量退出登录
export function Logout(params: any) {
  return whats.request({
    url: '/whats/logout',
    method: 'POST',
    params,
  });
}

// 发送消息
export function SendMsg(params: any) {
  return whats.request({
    url: '/whats/sendMsg',
    method: 'POST',
    params,
  });
}

// 发送消息
export function SendFile(params: any) {
  return whats.request({
    url: '/whats/sendFile',
    method: 'POST',
    params,
  });
}
// 发送名片
export function SendVcardMsg(params: any) {
  return whats.request({
    url: '/whats/sendVcardMsg',
    method: 'POST',
    params,
  });
}

// 同步联系人
export function SyncContact(params: any) {
  return whats.request({
    url: '/whats/syncContact',
    method: 'POST',
    params,
  });
}

// 解除绑定
export function UnBind(params: any) {
  return whats.request({
    url: '/whatsAccount/unBind',
    method: 'POST',
    params,
  });
}
//绑定账号
export function Bind(params: any) {
  return whats.request({
    url: '/whatsAccount/bind',
    method: 'POST',
    params,
  });
}
