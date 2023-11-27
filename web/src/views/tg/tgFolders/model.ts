import { ref } from 'vue';

import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Options } from '@/utils/hotgo';
import { getTgUserOption } from '@/api/tg/tgKeepTask';

export interface State {
  id: number;
  orgId: number;
  memberId: number;
  folderName: string;
  accounts: any;
  memberCount: number;
  comment: string;
  deletedAt: string;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  orgId: 0,
  memberId: 0,
  folderName: '',
  memberCount: 0,
  comment: '',
  deletedAt: '',
  createdAt: '',
  updatedAt: '',
  accounts: null,
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}

export const options = ref<Options>({
  accounts: [],
});

export const rules = {
  accounts: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'any',
    message: '请输入账号',
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
  {
    field: 'accounts',
    component: 'NSelect',
    label: '账号',
    defaultValue: null,
    componentProps: {
      multiple: true,
      placeholder: '请选择账号',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'folderName',
    component: 'NInput',
    label: '分组名称',
    componentProps: {
      placeholder: '请输入分组名称',
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
    title: '组织ID',
    key: 'orgId',
  },
  {
    title: '用户ID',
    key: 'memberId',
  },
  {
    title: '分组名称',
    key: 'folderName',
  },
  {
    title: '分组人数',
    key: 'memberCunt',
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
  const tgUser = await getTgUserOption({page: 1, pageSize: 9999});
  if (tgUser.list) {
    options.value.accounts = tgUser.list;
    for (let i = 0; i < tgUser.list.length; i++) {
      tgUser.list[i].label =
        tgUser.list[i].phone + '--' + tgUser.list[i].firstName + ' ' + tgUser.list[i].lastName;
      tgUser.list[i].value = tgUser.list[i].id;
    }
  } else {
    options.value.accounts = [];
  }

  for (const item of schemas.value) {
    switch (item.field) {
      case 'accounts':
        item.componentProps.options = options.value.accounts;
        debugger
        break;
    }
  }
}

await loadOptions()
