import { tg, jumpExport } from '@/utils/http/axios';

// 获取联系人管理列表
export function List(params: any) {
  return tg.request({
    url: '/tgContacts/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除联系人管理
export function Delete(params: any) {
  return tg.request({
    url: '/tgContacts/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑联系人管理
export function Edit(params: any) {
  return tg.request({
    url: '/tgContacts/edit',
    method: 'POST',
    params,
  });
}




// 获取联系人管理指定详情
export function View(params: any) {
  return tg.request({
    url: '/tgContacts/view',
    method: 'GET',
    params,
  });
}



// 导出联系人管理
export function Export(params: any) {
  jumpExport('/tgContacts/export', params);
}
