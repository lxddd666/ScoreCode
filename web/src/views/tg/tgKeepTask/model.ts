import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { getOptionLabel, getOptionTag, Options } from '@/utils/hotgo';
import { getScriptGroupOption, getTgUserOption } from '@/api/tg/tgKeepTask';
import { isNullObject } from '@/utils/is';
import { NTag } from 'naive-ui';

export interface State {
  id: number;
  orgId: number;
  taskName: string;
  cron: string;
  actions: any;
  accounts: any;
  scriptGroup: any;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  orgId: 0,
  taskName: '',
  cron: '',
  actions: null,
  accounts: null,
  scriptGroup: null,
  status: 2,
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
  keep_action: [],
  sys_job_status: [],
  accounts: [],
  scriptGroup: [],
});

export const rules = {
  taskName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入任务名称',
  },
  actions: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'any',
    message: '请输入养号动作',
  },
  accounts: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'any',
    message: '请输入账号',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'taskName',
    component: 'NInput',
    label: '任务名称',
    componentProps: {
      placeholder: '请输入任务名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'actions',
    component: 'NSelect',
    label: '养号动作',
    defaultValue: null,
    componentProps: {
      multiple: true,
      placeholder: '请选择养号动作',
      options: [],
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
    field: 'status',
    component: 'NSelect',
    label: '任务状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择任务状态',
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
    title: '表达式',
    key: 'cron',
  },
  {
    title: '任务状态',
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
          type: getOptionTag(options.value.sys_job_status, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_job_status, row.status),
        }
      );
    },
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
    types: ['keep_action', 'sys_job_status'],
  });

  const tgUser = await getTgUserOption({ page: 1, pageSize: 9999 });
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

  const group = await getScriptGroupOption();
  if (group.list) {
    options.value.scriptGroup = group.list;
    for (let i = 0; i < group.list.length; i++) {
      group.list[i].label = group.list[i].name;
      group.list[i].value = group.list[i].id;
    }
  } else {
    options.value.scriptGroup = [];
  }

  for (const item of schemas.value) {
    switch (item.field) {
      case 'actions':
        item.componentProps.options = options.value.keep_action;
        break;
      case 'status':
        item.componentProps.options = options.value.sys_job_status;
        break;
      case 'accounts':
        item.componentProps.options = options.value.accounts;
        break;
      case 'scriptGroup':
        item.componentProps.options = options.value.scriptGroup;
        break;
    }
  }
}

await loadOptions();
