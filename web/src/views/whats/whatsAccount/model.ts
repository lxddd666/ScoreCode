import {h, ref} from 'vue';
import {NTag} from 'naive-ui';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {Dicts} from '@/api/dict/dict';

import {isNullObject} from '@/utils/is';
import {defRangeShortcuts, formatBefore} from '@/utils/dateUtil';
import {getOptionLabel, getOptionTag, Options} from '@/utils/hotgo';


export interface State {
  id: number;
  account: string;
  nickName: string;
  avatar: string;
  accountStatus: number;
  isOnline: number;
  proxyAddress: string;
  comment: string;
  encryption: string;
  deletedAt: string;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  account: '',
  nickName: '',
  avatar: '',
  accountStatus: 1,
  isOnline: 1,
  proxyAddress: '',
  comment: '',
  encryption: '',
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
    field: 'account',
    component: 'NInput',
    label: '账号号码',
    componentProps: {
      placeholder: '请输入账号号码',
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
      placeholder: '请选择是否在线',
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
    title: '账号号码',
    key: 'account',
  },
  {
    title: '账号昵称',
    key: 'nickName',
  },
  {
    title: '账号头像',
    key: 'avatar',
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
    title: '最近活跃',
    key: 'lastLoginTime',
    width: 100,
    render(row) {
      if (row.lastLoginTime === null) {
        return '从未登录';
      }
      return formatBefore(new Date(row.lastLoginTime));
    },
  },
  {
    title: '代理地址',
    key: 'proxyAddress',
  },
  {
    title: '备注',
    key: 'comment',
  },
  {
    title: '更新时间',
    key: 'updatedAt',
  },
];

export const uploadColumns = [
  {
    title: '账号号码',
    key: 'account',
  },
  {
    title: '号码ID',
    key: 'identify',
  },
  {
    title: '公钥',
    key: 'publicKey',
  },
  {
    title: '私钥',
    key: 'privateKey',
  },
  {
    title: '消息公钥',
    key: 'publicMsgKey',
  },
  {
    title: '消息私钥',
    key: 'privateMsgKey',
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
