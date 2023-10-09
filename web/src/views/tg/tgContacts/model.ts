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
  tgId: number;
  username: string;
  firstName: string;
  lastName: string;
  phone: string;
  photo: string;
  orgId: number;
  comment: string;
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
  phone: '',
  photo: '',
  orgId: 0,
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


export const rules = {
  orgId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入organization id',
  },
};

export const schemas = ref<FormSchema[]>([
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
    title: 'tg id',
    key: 'tgId',
  },
  {
    title: 'username',
    key: 'username',
  },
  {
    title: 'First Name',
    key: 'firstName',
  },
  {
    title: 'Last Name',
    key: 'lastName',
  },
  {
    title: 'phone',
    key: 'phone',
  },
  {
    title: 'organization id',
    key: 'orgId',
  },
  {
    title: 'comment',
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