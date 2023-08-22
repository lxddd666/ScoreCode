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
  name: string;
  phone: string;
  avatar: string;
  email: string;
  address: string;
  orgId: number;
  deptId: number;
  comment: string;
  deletedAt: string;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  name: '',
  phone: '',
  avatar: '',
  email: '',
  address: '',
  orgId: 0,
  deptId: 0,
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
  phone: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入联系人电话',
  },
  email: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.email,
  },
  orgId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入组织id',
  },
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
    title: '联系人姓名',
    key: 'name',
  },
  {
    title: '联系人电话',
    key: 'phone',
  },
  {
    title: '联系人头像',
    key: 'avatar',
  },
  {
    title: '联系人邮箱',
    key: 'email',
  },
  {
    title: '联系人地址',
    key: 'address',
  },
  {
    title: '组织id',
    key: 'orgId',
  },
  {
    title: '部门id',
    key: 'deptId',
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
