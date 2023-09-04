import { whats, jumpExport } from '@/utils/http/axios';

// 获取消息记录列表
export function List(params) {
  return whats.request({
    url: '/whatsMsg/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除消息记录
export function Delete(params) {
  return whats.request({
    url: '/whatsMsg/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑消息记录
export function Edit(params) {
  return whats.request({
    url: '/whatsMsg/edit',
    method: 'POST',
    params,
  });
}




// 获取消息记录指定详情
export function View(params) {
  return whats.request({
    url: '/whatsMsg/view',
    method: 'GET',
    params,
  });
}



// 导出消息记录
export function Export(params) {
  jumpExport('/whatsMsg/export', params);
}
