import {h, ref} from 'vue';
import {NTag} from 'naive-ui';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {Dicts} from '@/api/dict/dict';

import {isNullObject} from '@/utils/is';
import {defRangeShortcuts} from '@/utils/dateUtil';
import {getOptionLabel, getOptionTag, Options} from '@/utils/hotgo';


export interface State {
  id: number;
  name: string;
}

export const defaultState = {
  id: 0,
  name: '',

};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}

export const options = ref<Options>({
  script_type: [],
});

export const rules = {
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入话术组名',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '话术组名',
    componentProps: {
      placeholder: '请输入话术组名',
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
    title: '组织',
    key: 'orgId',
  },
  {
    title: '用户',
    key: 'memberId',
  },
  {
    title: '分组类型',
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
          type: getOptionTag(options.value.script_type, row.type),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.script_type, row.type),
        }
      );
    },
  },
  {
    title: '话术组名',
    key: 'name',
  },
  {
    title: '话术数量',
    key: 'scriptCount',
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
   ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'type':
        item.componentProps.options = options.value.script_type;
        break;
     }
  }
}

await loadOptions();
