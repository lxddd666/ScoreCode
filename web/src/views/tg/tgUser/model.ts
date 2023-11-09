import {h, ref} from 'vue';
import {NTag} from 'naive-ui';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {Dicts} from '@/api/dict/dict';

import {isNullObject} from '@/utils/is';
import {defRangeShortcuts} from '@/utils/dateUtil';
import {getOptionLabel, getOptionTag, Options} from '@/utils/hotgo';

export interface State {
  id: number;
  username: string;
  firstName: string;
  lastName: string;
  phone: string;
  photo: string;
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
  username: '',
  firstName: '',
  lastName: '',
  phone: '',
  photo: '',
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
    title: "所属用户",
    key: "memberUsername",
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

async function loadOptions() {
  options.value = await Dicts({
    types: [
      'account_status',
      'login_status',
    ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'accountStatus':
        item.componentProps.options = options.value.account_status;
        break;
      case 'isOnline':
        item.componentProps.options = options.value.login_status;
        break;
    }
  }
}

await loadOptions();
