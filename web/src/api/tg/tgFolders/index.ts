import { tg, jumpExport } from '@/utils/http/axios';

// 获取tg分组列表
export function List(params: any) {
  return tg.request({
    url: '/tgFolders/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除tg分组
export function Delete(params: any) {
  return tg.request({
    url: '/tgFolders/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑tg分组
export function Edit(params: any) {
  return tg.request({
    url: '/tgFolders/edit',
    method: 'POST',
    params,
  });
}

// 获取tg分组指定详情
export function View(params: any) {
  return tg.request({
    url: '/tgFolders/view',
    method: 'GET',
    params,
  });
}

// 导出tg分组
export function Export(params: any) {
  jumpExport('/tgFolders/export', params);
}
