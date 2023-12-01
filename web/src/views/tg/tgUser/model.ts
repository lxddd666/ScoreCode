import { h, ref } from 'vue';
import { NImage, NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';

import { isNullObject } from '@/utils/is';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { errorImg, getOptionLabel, getOptionTag, Options } from '@/utils/hotgo';
import { getPhoto } from '@/utils/tgUtils';
import { List } from '@/api/tg/tgFolders';

export interface State {
  id: number;
  tgId: number;
  username: string;
  firstName: string;
  folders: number;
  lastName: string;
  phone: string;
  photo: number;
  bio: string;
  accountStatus: number;
  isOnline: number;
  proxyAddress: string;
  lastLoginTime: string;
  comment: string;
  session: string;
  deletedAt: string;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  tgId: 0,
  username: '',
  firstName: '',
  lastName: '',
  folders: 0,
  phone: '',
  photo: 0,
  bio: '',
  accountStatus: 0,
  isOnline: -1,
  proxyAddress: '',
  lastLoginTime: '',
  comment: '',
  session: '',
  deletedAt: '',
  createdAt: '',
  updatedAt: '',
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}

export const options = ref<Options>({
  account_status: [],
  login_status: [],
  folderId: [],
});

export const rules = {};

export const schemas = ref<FormSchema[]>([
  {
    field: 'username',
    component: 'NInput',
    label: '用户名',
    componentProps: {
      placeholder: '请输入用户名',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'firstName',
    component: 'NInput',
    label: '名字',
    componentProps: {
      placeholder: '请输入First Name',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'folderId',
    component: 'NSelect',
    label: '分组选择',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择分组',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'lastName',
    component: 'NInput',
    label: '姓氏',
    componentProps: {
      placeholder: '请输入Last Name',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'phone',
    component: 'NInput',
    label: '手机号',
    componentProps: {
      placeholder: '请输入手机号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'accountStatus',
    component: 'NSelect',
    label: '账号状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择账号状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'isOnline',
    component: 'NSelect',
    label: '是否在线',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择在线状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'proxyAddress',
    component: 'NInput',
    label: '代理地址',
    componentProps: {
      placeholder: '请输入代理地址',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export const columns = [
  {
    title: '所属用户',
    key: 'memberUsername',
  },
  {
    title: '用户名',
    key: 'username',
  },
  {
    title: '名字',
    key: 'firstName',
  },
  {
    title: '姓氏',
    key: 'lastName',
  },
  {
    title: '手机号',
    key: 'phone',
  },
  {
    title: '账号头像',
    key: 'photo',
    render(row) {
      // @ts-ignore
      return h(NImage, {
        width: 40,
        height: 40,
        src: getPhoto(row.phone, row.tgId, row.photo),
        fallbackSrc: errorImg,
        style: {
          width: '40px',
          height: '40px',
          'max-width': '100%',
          'max-height': '100%',
        },
      });
    },
  },
  {
    title: '账号状态',
    key: 'accountStatus',
    render(row) {
      if (isNullObject(row.accountStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.account_status, row.accountStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.account_status, row.accountStatus),
        }
      );
    },
  },
  {
    title: '是否在线',
    key: 'isOnline',
    render(row) {
      if (isNullObject(row.isOnline)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.login_status, row.isOnline),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.login_status, row.isOnline),
        }
      );
    },
  },
  {
    title: '代理地址',
    key: 'proxyAddress',
  },
  {
    title: '上次登录时间',
    key: 'lastLoginTime',
  },
  {
    title: '备注',
    key: 'comment',
  },
  {
    title: '创建时间',
    key: 'createdAt',
  },
  {
    title: '更新时间',
    key: 'updatedAt',
  },
];

export const uploadColumns = [
  {
    title: '代理地址',
    key: 'address',
  },
  {
    title: '代理类型',
    key: 'type',
  },
  {
    title: '最大连接数',
    key: 'maxConnections',
  },
  {
    title: '地区',
    key: 'region',
  },
  {
    title: '备注',
    key: 'comment',
  },
];

async function loadOptions() {
  options.value = await Dicts({
    types: ['account_status', 'login_status'],
  });
  const folderId = await List({ page: 1, pageSize: 9999 });
  if (folderId.list) {
    options.value.folderId = folderId.list;
    for (let i = 0; i < folderId.list.length; i++) {
      folderId.list[i].label = folderId.list[i].folderName;
      folderId.list[i].value = folderId.list[i].id;
    }
  } else {
    options.value.folderId = [];
  }
  debugger;
  for (const item of schemas.value) {
    switch (item.field) {
      case 'accountStatus':
        item.componentProps.options = options.value.account_status;
        break;
      case 'isOnline':
        item.componentProps.options = options.value.login_status;
        break;
      case 'folderId':
        item.componentProps.options = options.value.folderId;
        break;
    }
  }
}

await loadOptions();
