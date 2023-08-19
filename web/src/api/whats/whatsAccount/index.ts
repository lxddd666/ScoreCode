import { http, jumpExport } from '@/utils/http/axios';

// 获取小号管理列表
export function List(params) {
  return http.request({
    url: '/whatsAccount/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除小号管理
export function Delete(params) {
  return http.request({
    url: '/whatsAccount/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑小号管理
export function Edit(params) {
  return http.request({
    url: '/whatsAccount/edit',
    method: 'POST',
    params,
  });
}




// 获取小号管理指定详情
export function View(params) {
  return http.request({
    url: '/whatsAccount/view',
    method: 'GET',
    params,
  });
}

// 上传小号
export function Upload(params) {
  return http.request({
    url: '/whatsAccount/upload',
    method: 'POST',
    params,
  });
}
