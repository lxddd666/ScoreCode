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
  address: string;
  type: string;
  maxConnections: number;
  connectedCount: number;
  assignedCount: number;
  longTermCount: number;
  region: string;
  comment: string;
  status: number;
  deletedAt: string;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  address: '',
  type: '',
  maxConnections: 0,
  connectedCount: 0,
  assignedCount: 0,
  longTermCount: 0,
  region: '',
  comment: '',
  status: 1,
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
  proxy_type: [],
  sys_normal_disable: [],
});

export const rules = {
  address: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入代理地址',
  },
  type: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入代理类型',
  },
  maxConnections: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入最大连接数',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'address',
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
    field: 'type',
    component: 'NSelect',
    label: '代理类型',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择代理类型',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态',
      options: [],
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
    title: 'id',
    key: 'id',
  },
  {
    title: '代理地址',
    key: 'address',
  },
  {
    title: '代理类型',
    key: 'type',
    render(row) {
      if (isNullObject(row.type)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.proxy_type, row.type),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.proxy_type, row.type),
        }
      );
    },
  },
  {
    title: '最大连接数',
    key: 'maxConnections',
  },
  {
    title: '已连接数',
    key: 'connectedCount',
  },
  {
    title: '已分配账号数量',
    key: 'assignedCount',
  },
  {
    title: '长期未登录数量',
    key: 'longTermCount',
  },
  {
    title: '地区',
    key: 'region',
  },
  {
    title: '延迟',
    key: 'delay',
  },
  {
    title: '备注',
    key: 'comment',
  },
  {
    title: '状态',
    key: 'status',
    render(row) {
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.status),
        }
      );
    },
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
    types: [
      'proxy_type',
      'sys_normal_disable',
    ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'type':
        item.componentProps.options = options.value.proxy_type;
        break;
      case 'status':
        item.componentProps.options = options.value.sys_normal_disable;
        break;
    }
  }
}

await loadOptions();
