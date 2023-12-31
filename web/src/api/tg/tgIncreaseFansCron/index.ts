import {http, jumpExport, tg} from '@/utils/http/axios';


// 获取TG频道涨粉任务列表
export function List(params: any) {
  return tg.request({
    url: '/tgIncreaseFansCron/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除TG频道涨粉任务
export function Delete(params: any) {
  return tg.request({
    url: '/tgIncreaseFansCron/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑TG频道涨粉任务
export function Edit(params: any) {
  return tg.request({
    url: '/tgIncreaseFansCron/edit',
    method: 'POST',
    params,
  });
}




// 获取TG频道涨粉任务指定详情
export function View(params: any) {
  return http.request({
    url: '/tgIncreaseFansCron/view',
    method: 'GET',
    params,
  });
}

// 导出TG频道涨粉任务
export function Export(params: any) {
  jumpExport('/tgIncreaseFansCron/export', params);
}
