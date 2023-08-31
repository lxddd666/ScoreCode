import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { Dicts } from '@/api/dict/dict';
import { getOptionLabel, getOptionTag, Options } from '@/utils/hotgo';
import { isNullObject } from '@/utils/is';
import { NAvatar, NTag } from 'naive-ui';

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
export const options = ref<Options>({
  org_id: [],
  dept_id: [],
});

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
    field: 'phone',
    component: 'NInput',
    label: '联系人号码',
    componentProps: {
      placeholder: '请输入账号号码',
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
    title: '姓名',
    key: 'name',
  },
  {
    title: '手机号',
    key: 'phone',
  },
  {
    title: '头像',
    key: 'avatar',
  },
  {
    title: '邮箱',
    key: 'email',
  },
  {
    title: '地址',
    key: 'address',
  },
  {
    title: '组织',
    key: 'orgId',
  },
  {
    title: '部门',
    key: 'deptId',
  },
  {
    title: '备注',
    key: 'comment',
  },
];

export const columns = [
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
    render(row) {
      if (row.avatar !== '') {
        return h(NAvatar, {
          circle: true,
          size: 'small',
          src: row.avatar,
        });
      } else {
        return h(
          NAvatar,
          {
            circle: true,
            size: 'small',
          },
          {
            default: () => (row.name !== '' ? row.name.substring(0, 1) : row.name.substring(0, 2)),
          }
        );
      }
    },
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
    render(row) {
      if (isNullObject(row.orgId)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.org_id, row.orgId),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.org_id, row.orgId),
        }
      );
    },
  },
  {
    title: '部门id',
    key: 'deptId',
    render(row) {
      if (isNullObject(row.deptId)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.dept_id, row.deptId),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.dept_id, row.deptId),
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
    types: ['org_id', 'dept_id'],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'orgId':
        item.componentProps.options = options.value.org_id;
        break;
      case 'deptId':
        item.componentProps.options = options.value.dept_id;
        break;
    }
  }
}

await loadOptions();
