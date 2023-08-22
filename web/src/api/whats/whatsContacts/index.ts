import { http, jumpExport } from '@/utils/http/axios';

// 获取联系人管理列表
export function List(params) {
  return http.request({
    url: '/whatsContacts/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除联系人管理
export function Delete(params) {
  return http.request({
    url: '/whatsContacts/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑联系人管理
export function Edit(params) {
  return http.request({
    url: '/whatsContacts/edit',
    method: 'POST',
    params,
  });
}




// 获取联系人管理指定详情
export function View(params) {
  return http.request({
    url: '/whatsContacts/view',
    method: 'GET',
    params,
  });
}



// 导出联系人管理
export function Export(params) {
  jumpExport('/whatsContacts/export', params);
}