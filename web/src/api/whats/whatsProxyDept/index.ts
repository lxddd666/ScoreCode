import { http, jumpExport } from '@/utils/http/axios';

// 获取代理关联公司列表
export function List(params: any) {
  return http.request({
    url: '/whatsProxyDept/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除代理关联公司
export function Delete(params: any) {
  return http.request({
    url: '/whatsProxyDept/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑代理关联公司
export function Edit(params: any) {
  return http.request({
    url: '/whatsProxyDept/edit',
    method: 'POST',
    params,
  });
}




// 获取代理关联公司指定详情
export function View(params: any) {
  return http.request({
    url: '/whatsProxyDept/view',
    method: 'GET',
    params,
  });
}



// 导出代理关联公司
export function Export(params: any) {
  jumpExport('/whatsProxyDept/export', params);
}