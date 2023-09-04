import { whats, jumpExport } from '@/utils/http/axios';

// 获取代理管理列表
export function List(params) {
  return whats.request({
    url: '/whatsProxy/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除代理管理
export function Delete(params) {
  return whats.request({
    url: '/whatsProxy/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑代理管理
export function Edit(params) {
  return whats.request({
    url: '/whatsProxy/edit',
    method: 'POST',
    params,
  });
}


// 修改代理管理状态
export function Status(params) {
  return whats.request({
    url: '/whatsProxy/status',
    method: 'POST',
    params,
  });
}



// 获取代理管理指定详情
export function View(params) {
  return whats.request({
    url: '/whatsProxy/view',
    method: 'GET',
    params,
  });
}

export function Upload(params: any) {
  return whats.request({
    url: '/whatsProxy/upload',
    method: 'POST',
    params,
  });
}

// 导出代理管理
export function Export(params) {
  jumpExport('/whatsProxy/export', params);
}
