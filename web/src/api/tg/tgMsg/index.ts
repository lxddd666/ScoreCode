import { tg, jumpExport } from '@/utils/http/axios';

// 获取消息记录列表
export function List(params: any) {
  return tg.request({
    url: '/tgMsg/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除消息记录
export function Delete(params: any) {
  return tg.request({
    url: '/tgMsg/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑消息记录
export function Edit(params: any) {
  return tg.request({
    url: '/tgMsg/edit',
    method: 'POST',
    params,
  });
}




// 获取消息记录指定详情
export function View(params: any) {
  return tg.request({
    url: '/tgMsg/view',
    method: 'GET',
    params,
  });
}



// 导出消息记录
export function Export(params: any) {
  jumpExport('/tgMsg/export', params);
}
