import { h, ref } from 'vue';
import { NAvatar, NImage, NTag, NSwitch, NRate } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';

import { isArray, isNullObject } from '@/utils/is';
import { getFileExt } from '@/utils/urlUtils';
import { defRangeShortcuts, defShortcuts, formatToDate } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { getOptionLabel, getOptionTag, Options, errorImg } from '@/utils/hotgo';


export interface State {
  id: number;
  orgId: number;
  taskName: string;
  action: number;
  accounts: any;
  parameters: any;
  status: number;
  comment: string;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  orgId: 0,
  taskName: '',
  action: 0,
  accounts: null,
  parameters: null,
  status: 1,
  comment: '',
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
  orgId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入组织ID',
  },
  taskName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入任务名称',
  },
  action: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入操作动作',
  },
  accounts: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'any',
    message: '请输入账号 id',
  },
  parameters: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'any',
    message: '请输入执行任务参数',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'ID',
    componentProps: {
      placeholder: '请输入ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '任务状态,1运行,2停止,3完成,4失败',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择任务状态,1运行,2停止,3完成,4失败',
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
    title: 'ID',
    key: 'id',
  },
  {
    title: '组织ID',
    key: 'orgId',
  },
  {
    title: '任务名称',
    key: 'taskName',
  },
  {
    title: '操作动作',
    key: 'action',
  },
  {
    title: '任务状态,1运行,2停止,3完成,4失败',
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
    title: '备注',
    key: 'comment',
  },
  {
    title: '创建时间',
    key: 'createdAt',
  },
  {
    title: '修改时间',
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