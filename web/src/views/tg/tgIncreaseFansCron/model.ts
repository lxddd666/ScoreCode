import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';

import { isNullObject } from '@/utils/is';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { getOptionLabel, getOptionTag, Options } from '@/utils/hotgo';

export interface State {
  id: number;
  orgId: number;
  memberId: number;
  channel: string;
  channelId: string;
  dayCount: number;
  fansCount: number;
  cronStatus: number;
  comment: string;
  deletedAt: string;
  createdAt: string;
  updatedAt: string;
  executedDays: number;
  increasedFans: number;
  channelMemberCount: number;
  folderId: number;
  recommendedDays: number;
  taskName: string;
  executedPlan: any;
  executedPlanStr: string;
}

export const defaultState = {
  id: 0,
  orgId: 0,
  memberId: 0,
  channel: '',
  channelId: '',
  folderId: 0,
  dayCount: 0,
  fansCount: 0,
  cronStatus: 0,
  comment: '',
  deletedAt: '',
  createdAt: '',
  updatedAt: '',
  executedDays: 0,
  increasedFans: 0,
  channelMemberCount: 0,
  recommendedDays: 0,
  executedPlan: null,
  taskName: '',
  executedPlanStr: '',
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}

export const options = ref<Options>({
  increase_fans: [],
});

export const rules = {};

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
    field: 'cronStatus',
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
    title: 'id',
    key: 'id',
  },
  {
    title: '组织ID',
    key: 'orgId',
  },
  {
    title: '发起任务的用户ID',
    key: 'memberId',
  },
  {
    title: '频道地址',
    key: 'channel',
  },
  {
    title: '任务名称',
    key: 'taskName',
  },
  {
    title: '持续天数',
    key: 'dayCount',
  },
  {
    title: '涨粉数量',
    key: 'fansCount',
  },
  {
    title: '任务状态',
    key: 'cronStatus',
    render(row) {
      if (isNullObject(row.cronStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.increase_fans, row.cronStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.increase_fans, row.cronStatus),
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
  {
    title: '已执行天数',
    key: 'executedDays',
  },
  {
    title: '已添加粉丝数',
    key: 'increasedFans',
  },
];

async function loadOptions() {
  options.value = await Dicts({
    // types: ['sys_normal_disable'],
    types: ['increase_fans'],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'cronStatus':
        item.componentProps.options = options.value.increase_fans;
        break;
    }
  }
}

await loadOptions();
