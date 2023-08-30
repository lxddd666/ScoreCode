import { http } from '@/utils/http/axios';

export function getDeptList(params?) {
  return http.request({
    url: '/dept/list',
    method: 'GET',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/dept/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/dept/status',
    method: 'POST',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/dept/delete',
    method: 'POST',
    params,
  });
}

export function getDeptOption(param?) {
  const params = { pageSize: 100, orgId: param };
  return getDeptOptionList(params);
}

export function getDeptOptionList(params?) {
  return http.request({
    url: '/dept/option',
    method: 'GET',
    params,
  });
}

export function getOrgOption(params?) {
  return getOrgList(params);
}

export function getOrgList(params?) {
  return http.request({
    url: '/dept/deptOrgOption',
    method: 'GET',
    params,
  });
}
