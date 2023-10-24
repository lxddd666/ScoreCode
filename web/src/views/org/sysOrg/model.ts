import {h, ref} from 'vue';
import {NTag} from 'naive-ui';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {Dicts} from '@/api/dict/dict';

import {isNullObject} from '@/utils/is';
import {defRangeShortcuts} from '@/utils/dateUtil';
import {validate} from '@/utils/validateUtil';
import {getOptionLabel, getOptionTag, Options} from '@/utils/hotgo';


export interface State {
  id: number;
  name: string;
  code: string;
  leader: string;
  phone: string;
  email: string;
  sort: number;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  name: '',
  code: '',
  leader: '',
  phone: '',
  email: '',
  sort: 1,
  status: 1,
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
});

export const rules = {
  email: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.email,
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '公司名称',
    componentProps: {
      placeholder: '请输入公司名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '公司状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择公司状态',
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
    title: '公司ID',
    key: 'id',
  },
  {
    title: '公司名称',
    key: 'name',
  },
  {
    title: '公司编码',
    key: 'code',
  },
  {
    title: '负责人',
    key: 'leader',
  },
  {
    title: '联系电话',
    key: 'phone',
  },
  {
    title: '邮箱',
    key: 'email',
  },
  {
    title: '总端口数',
    key: 'portTotal',
  },
  {
    title: '已用端口数',
    key: 'usedPortNum',
  },
  {
    title: '剩余端口数',
    key: 'marginPortNum',
  },
  {
    title: '公司状态',
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
    types: [
      'sys_normal_disable',
    ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'status':
        item.componentProps.options = options.value.sys_normal_disable;
        break;
    }
  }
}

await loadOptions();
