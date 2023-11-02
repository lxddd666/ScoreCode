import {http, jumpExport, tg} from '@/utils/http/axios';

// 获取养号任务列表
export function List(params: any) {
  return tg.request({
    url: '/tgKeepTask/list',
    method: 'get',
    params,
  });
}

export function getTgUserOption(params?: any) {
  return tg.request({
    url: '/tgUser/list',
    method: 'GET',
    params,
  });
}

export function getScriptGroupOption(params?: any) {
  return http.request({
    url: '/2/scriptGroup/list',
    method: 'GET',
    params,
  });
}

// 删除/批量删除养号任务
export function Delete(params: any) {
  return tg.request({
    url: '/tgKeepTask/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑养号任务
export function Edit(params: any) {
  return tg.request({
    url: '/tgKeepTask/edit',
    method: 'POST',
    params,
  });
}


// 获取养号任务指定详情
export function View(params: any) {
  return tg.request({
    url: '/tgKeepTask/view',
    method: 'GET',
    params,
  });
}

// 修改养号任务状态
export function Status(params: any) {
  return tg.request({
    url: '/tgKeepTask/status',
    method: 'POST',
    params,
  });
}

// 执行一次
export function Once(params: any) {
  return tg.request({
    url: '/tgKeepTask/once',
    method: 'POST',
    params,
  });
}

// 导出养号任务
export function Export(params: any) {
  jumpExport('/tgKeepTask/export', params);
}
