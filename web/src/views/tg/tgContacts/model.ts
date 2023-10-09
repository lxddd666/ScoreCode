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
  type: number;
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
  type: 0,
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

export const options = ref<Options>({
  contacts_type: [],
});

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
    field: 'phone',
    component: 'NInput',
    label: 'phone',
    componentProps: {
      placeholder: '请输入phone',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'type',
    component: 'NRadioGroup',
    label: 'type',
    giProps: {
      //span: 24,
    },
    componentProps: {
      options: [],
      onUpdateChecked: (e: any) => {
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
    title: 'type',
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
          type: getOptionTag(options.value.contacts_type, row.type),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.contacts_type, row.type),
        }
      );
    },
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

async function loadOptions() {
  options.value = await Dicts({
    types: [
      'contacts_type',
   ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'type':
        item.componentProps.options = options.value.contacts_type;
        break;
     }
  }
}

await loadOptions();