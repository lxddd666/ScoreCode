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
  createdAt: string;
  updatedAt: string;
  deletedAt: string;
  initiator: number;
  sender: number;
  receiver: number;
  reqId: string;
  sendMsg: string;
  translatedMsg: string;
  msgType: number;
  sendTime: string;
  read: number;
  comment: string;
  sendStatus: number;
}

export const defaultState = {
  id: 0,
  createdAt: '',
  updatedAt: '',
  deletedAt: '',
  initiator: 0,
  sender: 0,
  receiver: 0,
  reqId: '',
  sendMsg: '',
  translatedMsg: '',
  msgType: 1,
  sendTime: '',
  read: 0,
  comment: '',
  sendStatus: 0,
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}

export const options = ref<Options>({
  read_status: [],
  send_status: [],
});

export const rules = {};

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
  {
    field: 'initiator',
    component: 'NInputNumber',
    label: '聊天发起人',
    componentProps: {
      placeholder: '请输入聊天发起人',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'sender',
    component: 'NInputNumber',
    label: '发送人',
    componentProps: {
      placeholder: '请输入发送人',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'receiver',
    component: 'NInputNumber',
    label: '接收人',
    componentProps: {
      placeholder: '请输入接收人',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'reqId',
    component: 'NInput',
    label: '请求id',
    componentProps: {
      placeholder: '请输入请求id',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'read',
    component: 'NSelect',
    label: '是否已读',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择是否已读',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'sendStatus',
    component: 'NSelect',
    label: '发送状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择发送状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export const columns = [

  {
    title: '创建时间',
    key: 'createdAt',
  },
  {
    title: '更新时间',
    key: 'updatedAt',
  },
  {
    title: '聊天发起人',
    key: 'initiator',
  },
  {
    title: '发送人',
    key: 'sender',
  },
  {
    title: '接收人',
    key: 'receiver',
  },
  {
    title: '请求id',
    key: 'reqId',
  },
  {
    title: '消息类型',
    key: 'msgType',
  },
  {
    title: '发送时间',
    key: 'sendTime',
  },
  {
    title: '是否已读',
    key: 'read',
    render(row) {
      if (isNullObject(row.read)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.read_status, row.read),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.read_status, row.read),
        }
      );
    },
  },
  {
    title: '备注',
    key: 'comment',
  },
  {
    title: '发送状态',
    key: 'sendStatus',
    render(row) {
      if (isNullObject(row.sendStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.send_status, row.sendStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.send_status, row.sendStatus),
        }
      );
    },
  },
];

async function loadOptions() {
  options.value = await Dicts({
    types: [
      'read_status',
      'send_status',
    ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'read':
        item.componentProps.options = options.value.read_status;
        break;
      case 'sendStatus':
        item.componentProps.options = options.value.send_status;
        break;
    }
  }
}

await loadOptions();
