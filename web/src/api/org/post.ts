import { http } from '@/utils/http/axios';

export function getPostList(params?) {
  return http.request({
    url: '/post/list',
    method: 'GET',
    params,
  });
}

export function getPostOption(param?) {
  const params = { pageSize: 100, orgId: param };
  return getPostList(params);
}

export function Edit(params) {
  return http.request({
    url: '/post/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/post/status',
    method: 'POST',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/post/delete',
    method: 'POST',
    params,
  });
}
