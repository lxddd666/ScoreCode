import { http, jumpExport, tg } from '@/utils/http/axios';

// 获取TG频道涨粉任务执行情况列表
export function List(params: any) {
  return tg.request({
    url: '/tgIncreaseFansCronAction/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除TG频道涨粉任务执行情况
export function Delete(params: any) {
  return tg.request({
    url: '/tgIncreaseFansCronAction/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑TG频道涨粉任务执行情况
export function Edit(params: any) {
  return tg.request({
    url: '/tgIncreaseFansCronAction/edit',
    method: 'POST',
    params,
  });
}

// 获取TG频道涨粉任务执行情况指定详情
export function View(params: any) {
  return http.request({
    url: '/tgIncreaseFansCronAction/view',
    method: 'GET',
    params,
  });
}

// 导出TG频道涨粉任务执行情况
export function Export(params: any) {
  jumpExport('/tgIncreaseFansCronAction/export', params);
}
