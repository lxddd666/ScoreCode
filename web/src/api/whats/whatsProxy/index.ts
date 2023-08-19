import { http, jumpExport } from '@/utils/http/axios';

// 获取代理管理列表
export function List(params) {
  return http.request({
    url: '/whatsProxy/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除代理管理
export function Delete(params) {
  return http.request({
    url: '/whatsProxy/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑代理管理
export function Edit(params) {
  return http.request({
    url: '/whatsProxy/edit',
    method: 'POST',
    params,
  });
}


// 修改代理管理状态
export function Status(params) {
  return http.request({
    url: '/whatsProxy/status',
    method: 'POST',
    params,
  });
}



// 获取代理管理指定详情
export function View(params) {
  return http.request({
    url: '/whatsProxy/view',
    method: 'GET',
    params,
  });
}



// 导出代理管理
export function Export(params) {
  jumpExport('/whatsProxy/export', params);
}