import { http, jumpExport } from '@/utils/http/axios';

// 获取客户公司列表
export function List(params: any) {
  return http.request({
    url: '/org/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除客户公司
export function Delete(params: any) {
  return http.request({
    url: '/org/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑客户公司
export function Edit(params: any) {
  return http.request({
    url: '/org/edit',
    method: 'POST',
    params,
  });
}


// 修改客户公司状态
export function Status(params: any) {
  return http.request({
    url: '/org/status',
    method: 'POST',
    params,
  });
}



// 获取客户公司指定详情
export function View(params: any) {
  return http.request({
    url: '/org/view',
    method: 'GET',
    params,
  });
}


// 获取客户公司最大排序
export function MaxSort() {
  return http.request({
    url: '/org/maxSort',
    method: 'GET',
  });
}


// 导出客户公司
export function Export(params: any) {
  jumpExport('/org/export', params);
}
