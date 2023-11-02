import {h, ref} from 'vue';
import {NTag} from 'naive-ui';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {Dicts} from '@/api/dict/dict';
import {getGroupOption} from '@/api/script/memberScriptGroup';

import {isNullObject} from '@/utils/is';
import {defRangeShortcuts} from '@/utils/dateUtil';
import {getOptionLabel, getOptionTag, Options} from '@/utils/hotgo';


export interface State {
  id: number;
  orgId: number;
  memberId: number;
  groupId: number|any;
  scriptClass: number|any;
  short: string;
  content: string;
  sendCount: number;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  orgId: 0,
  memberId: 0,
  groupId: '',
  scriptClass: '',
  short: '',
  content: '',
  sendCount: 0,
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
  group: [],
  script_type: [],
  script_class: [],
});

export const rules = {
  groupId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入分组',
  },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入话术内容',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'short',
    component: 'NInput',
    label: '快捷指令',
    componentProps: {
      placeholder: '请输入快捷指令',
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
    title: '分组',
    key: 'groupId',
  },
  {
    title: '话术分类',
    key: 'scriptClass',
    render(row) {
      if (isNullObject(row.scriptClass)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.script_class, row.scriptClass),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.script_class, row.scriptClass),
        }
      );
    },
  },
  {
    title: '快捷指令',
    key: 'short',
  },
  {
    title: '话术内容',
    key: 'content',
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
      'script_type',
      'script_class',
    ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'type':
        item.componentProps.options = options.value.script_type;
        break;
      case 'scriptClass':
        item.componentProps.options = options.value.script_class;
        break;
    }
  }
  const group = await getGroupOption();
  if (group.list) {
    options.value.group = group.list;
    for (let i = 0; i < group.list.length; i++) {
      group.list[i].label = group.list[i].name;
      group.list[i].value = group.list[i].id;
    }
  } else {
    options.value.org = [];
  }

}

await loadOptions();
