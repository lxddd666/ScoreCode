import { tg, jumpExport } from '@/utils/http/axios';

// 获取TG账号列表
export function List(params: any) {
  return tg.request({
    url: '/tgUser/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除TG账号
export function Delete(params: any) {
  return tg.request({
    url: '/tgUser/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑TG账号
export function Edit(params: any) {
  return tg.request({
    url: '/tgUser/edit',
    method: 'POST',
    params,
  });
}




// 获取TG账号指定详情
export function View(params: any) {
  return tg.request({
    url: '/tgUser/view',
    method: 'GET',
    params,
  });
}

// 绑定员工
export function TgBindMember(params: any) {
  return tg.request({
    url: '/tgUser/bindMember',
    method: 'POST',
    params,
  });
}

// 绑定代理
export function TgBindProxy(params: any) {
  return tg.request({
    url: '/tgUser/bindProxy',
    method: 'POST',
    params,
  });
}


// 导出TG账号
export function Export(params: any) {
  jumpExport('/tgUser/export', params);
}
