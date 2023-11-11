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
  cronId: number;
  tgUserId: number;
  joinStatus: number;
  comment: string;
  deletedAt: string;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  cronId: 0,
  tgUserId: 0,
  joinStatus: 0,
  comment: '',
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
});

export const rules = {
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'id',
    componentProps: {
      placeholder: '请输入id',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'joinStatus',
    component: 'NSelect',
    label: '加入状态：0失败，1成功，2完成',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择加入状态：0失败，1成功，2完成',
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
    title: '任务ID',
    key: 'cronId',
  },
  {
    title: '加入频道的userId',
    key: 'tgUserId',
  },
  {
    title: '加入状态：0失败，1成功，2完成',
    key: 'joinStatus',
    render(row) {
      if (isNullObject(row.joinStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.joinStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.joinStatus),
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
      'sys_normal_disable',
   ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'joinStatus':
        item.componentProps.options = options.value.sys_normal_disable;
        break;
     }
  }
}

await loadOptions();
