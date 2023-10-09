import { tg, jumpExport } from '@/utils/http/axios';

// 获取代理管理列表
export function List(params: any) {
  return tg.request({
    url: '/tgProxy/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除代理管理
export function Delete(params: any) {
  return tg.request({
    url: '/tgProxy/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑代理管理
export function Edit(params: any) {
  return tg.request({
    url: '/tgProxy/edit',
    method: 'POST',
    params,
  });
}


// 修改代理管理状态
export function Status(params: any) {
  return tg.request({
    url: '/tgProxy/status',
    method: 'POST',
    params,
  });
}



// 获取代理管理指定详情
export function View(params: any) {
  return tg.request({
    url: '/tgProxy/view',
    method: 'GET',
    params,
  });
}



// 导出代理管理
export function Export(params: any) {
  jumpExport('/tgProxy/export', params);
}
