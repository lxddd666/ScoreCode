import { tg, jumpExport } from '@/utils/http/axios';

// 获取批量操作任务列表
export function List(params: any) {
  return tg.request({
    url: '/tgBatchExecutionTask/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除批量操作任务
export function Delete(params: any) {
  return tg.request({
    url: '/tgBatchExecutionTask/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑批量操作任务
export function Edit(params: any) {
  return tg.request({
    url: '/tgBatchExecutionTask/edit',
    method: 'POST',
    params,
  });
}

// 修改批量操作任务状态
export function Status(params: any) {
  return tg.request({
    url: '/tgBatchExecutionTask/status',
    method: 'POST',
    params,
  });
}

// 获取批量操作任务指定详情
export function View(params: any) {
  return tg.request({
    url: '/tgBatchExecutionTask/view',
    method: 'GET',
    params,
  });
}

// 导出批量操作任务
export function Export(params: any) {
  jumpExport('/tgBatchExecutionTask/export', params);
}
