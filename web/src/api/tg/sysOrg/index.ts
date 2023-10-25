import { http, jumpExport } from '@/utils/http/axios';

// 获取公司信息列表
export function List(params: any) {
  return http.request({
    url: '/org/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除公司信息
export function Delete(params: any) {
  return http.request({
    url: '/org/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑公司信息
export function Edit(params: any) {
  return http.request({
    url: '/org/edit',
    method: 'POST',
    params,
  });
}


// 修改公司信息状态
export function Status(params: any) {
  return http.request({
    url: '/org/status',
    method: 'POST',
    params,
  });
}

// 修改端口数
export function Ports(params: any) {
  return http.request({
    url: '/org/ports',
    method: 'POST',
    params,
  });
}


// 获取公司信息指定详情
export function View(params: any) {
  return http.request({
    url: '/org/view',
    method: 'GET',
    params,
  });
}


// 获取公司信息最大排序
export function MaxSort() {
  return http.request({
    url: '/org/maxSort',
    method: 'GET',
  });
}


// 导出公司信息
export function Export(params: any) {
  jumpExport('/org/export', params);
}
