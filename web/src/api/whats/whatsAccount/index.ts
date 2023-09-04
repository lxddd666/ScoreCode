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

// 发送消息
export function SendMsg(params: any) {
  return whats.request({
    url: '/whats/sendMsg',
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


// 解除绑定
export function UnBind(params: any) {
  return whats.request({
    url: '/whatsAccount/unBind',
    method: 'POST',
    params,
  });
}
