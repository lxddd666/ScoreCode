import { http, jumpExport } from '@/utils/http/axios';

// 获取代理管理列表
export function List(params: any) {
  return http.request({
    url: '/sysProxy/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除代理管理
export function Delete(params: any) {
  return http.request({
    url: '/sysProxy/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑代理管理
export function Edit(params: any) {
  return http.request({
    url: '/sysProxy/edit',
    method: 'POST',
    params,
  });
}


// 修改代理管理状态
export function Status(params: any) {
  return http.request({
    url: '/sysProxy/status',
    method: 'POST',
    params,
  });
}

// 修改代理管理状态
export function Upload(params: any) {
  return http.request({
    url: '/sysProxy/import',
    method: 'POST',
    params,
  });
}


// 获取代理管理指定详情
export function View(params: any) {
  return http.request({
    url: '/sysProxy/view',
    method: 'GET',
    params,
  });
}



// 导出代理管理
export function Export(params: any) {
  jumpExport('/sysProxy/export', params);
}
