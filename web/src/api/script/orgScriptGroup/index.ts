import { http, jumpExport } from '@/utils/http/axios';

// 获取话术分组列表
export function List(params: any) {
  return http.request({
    url: '/2/scriptGroup/list',
    method: 'get',
    params,
  });
}

export function getGroupOption(params?: any) {
  return http.request({
    url: '/2/scriptGroup/list',
    method: 'GET',
    params,
  });
}

// 删除/批量删除话术分组
export function Delete(params: any) {
  return http.request({
    url: '/2/scriptGroup/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑话术分组
export function Edit(params: any) {
  return http.request({
    url: '/2/scriptGroup/edit',
    method: 'POST',
    params,
  });
}




// 获取话术分组指定详情
export function View(params: any) {
  return http.request({
    url: '/2/scriptGroup/view',
    method: 'GET',
    params,
  });
}



// 导出话术分组
export function Export(params: any) {
  jumpExport('/2/scriptGroup/export', params);
}
