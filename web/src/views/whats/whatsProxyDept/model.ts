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
  deptId: string;
  proxyAddress: string;
  comment: string;
}

export const defaultState = {
  id: 0,
  deptId: NULL,
  proxyAddress: NULL,
  comment: NULL,
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}


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
]);

export const columns = [
  {
    title: 'id',
    key: 'id',
  },
  {
    title: '公司部门id',
    key: 'deptId',
  },
  {
    title: '代理地址',
    key: 'proxyAddress',
  },
  {
    title: '备注',
    key: 'comment',
  },
];