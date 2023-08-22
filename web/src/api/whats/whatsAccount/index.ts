import { http } from '@/utils/http/axios';

// 获取帐号管理列表
export function List(params: any) {
  return http.request({
    url: '/whatsAccount/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除帐号管理
export function Delete(params: any) {
  return http.request({
    url: '/whatsAccount/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑帐号管理
export function Edit(params: any) {
  return http.request({
    url: '/whatsAccount/edit',
    method: 'POST',
    params,
  });
}




// 获取帐号管理指定详情
export function View(params: any) {
  return http.request({
    url: '/whatsAccount/view',
    method: 'GET',
    params,
  });
}

// 上传帐号
export function Upload(params: any) {
  return http.request({
    url: '/whatsAccount/upload',
    method: 'POST',
    params,
  });
}

// 登录/批量登录帐号
export function Login(params: any) {
  return http.request({
    url: '/whats/login',
    method: 'POST',
    params,
  });
}

// 发送消息
export function SendMsg(params: any) {
  return http.request({
    url: '/whats/sendMsg',
    method: 'POST',
    params,
  });
}


// 解除绑定
export function UnBind(params: any) {
  return http.request({
    url: '/whatsAccount/unBind',
    method: 'POST',
    params,
  });
}
