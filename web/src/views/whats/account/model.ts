import {h, ref} from 'vue';
import {NTag} from 'naive-ui';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {Dicts} from '@/api/dict/dict';

import {isNullObject} from '@/utils/is';
import {getOptionLabel, getOptionTag, Options} from '@/utils/hotgo';


export interface State {
  id: number;
  account: string;
  nickName: string;
  avatar: string;
  accountStatus: number;
  isOnline: number;
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

export const rules = {
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'account',
    component: 'NInput',
    label: '账号号码',
    componentProps: {
      placeholder: '请输入账号',
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
]);

export const columns = [
  {
    title: 'ID',
    key: 'id',
    sorter: true, // 单列排序
  },
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
