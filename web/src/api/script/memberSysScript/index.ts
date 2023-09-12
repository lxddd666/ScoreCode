import { http, jumpExport } from '@/utils/http/axios';

// 获取个人话术管理列表
export function List(params: any) {
  return http.request({
    url: '/1/sysScript/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除个人话术管理
export function Delete(params: any) {
  return http.request({
    url: '/1/sysScript/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑个人话术管理
export function Edit(params: any) {
  return http.request({
    url: '/1/sysScript/edit',
    method: 'POST',
    params,
  });
}




// 获取个人话术管理指定详情
export function View(params: any) {
  return http.request({
    url: '/1/sysScript/view',
    method: 'GET',
    params,
  });
}



// 导出个人话术管理
export function Export(params: any) {
  jumpExport('/1/sysScript/export', params);
}
