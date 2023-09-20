import {h, ref} from 'vue';
import {NTag} from 'naive-ui';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {Dicts} from '@/api/dict/dict';

import {isNullObject} from '@/utils/is';
import {defRangeShortcuts} from '@/utils/dateUtil';
import {getOptionLabel, getOptionTag, Options} from '@/utils/hotgo';
import {getOrgOption} from "@/api/org/dept";

export interface State {
  orgId: number;
  id: number;
  address: string;
  connectedCount: number;
  maxConnections: number;
  region: string;
  comment: string;
  status: number;
  deletedAt: string;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  orgId: 0,
  id: 0,
  address: '',
  connectedCount: 0,
  maxConnections: 0,
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
  sys_normal_disable: [],
  org: [],
});

export const rules = {};

export const schemas = ref<FormSchema[]>([
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

export const uploadColumns = [
  {
    title: '代理地址',
    key: 'address',
  },
  {
    title: '最大连接数',
    key: 'max_connections',
  },
  {
    title: '已连接数',
    key: 'connected_count',
  },
  {
    title: '已分配账号数量',
    key: 'assigned_count',
  },
  {
    title: '长期未登录数量',
    key: 'long_term_count',
  },
  {
    title: '地区',
    key: 'region',
  },
  {
    title: '备注',
    key: 'comment',
  },
  {
    title: '状态(1正常, 2停用)',
    key: 'status',
  },
];

export const columns = [
  {
    title: '代理地址',
    key: 'address',
  },
  {
    title: '已连接数',
    key: 'connectedCount',
  },
  {
    title: '最大连接',
    key: 'maxConnections',
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

async function loadOptions() {
  options.value = await Dicts({
    types: ['sys_normal_disable'],
  });
  const org = await getOrgOption();
  if (org.list) {
    options.value.org = org.list;
    for (let i = 0; i < org.list.length; i++) {
      org.list[i].label = org.list[i].name;
      org.list[i].value = org.list[i].id;
    }
  } else {
    options.value.org = [];
  }
  for (const item of schemas.value) {
    switch (item.field) {
      case 'status':
        item.componentProps.options = options.value.sys_normal_disable;
        break;
    }
  }
}

await loadOptions();
