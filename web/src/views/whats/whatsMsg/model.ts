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
  sendStatus:number;
  comment: string;
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
  sendStatus: 0,
  comment: '',
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}

export const options = ref<Options>({
  msg_type: [],
  read_status: [],
  send_status:[],
});

export const rules = {
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'sender',
    component: 'NInput',
    label: '发送人',
    componentProps: {
      placeholder: '请输入发送人',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'sendTime',
    component: 'NDatePicker',
    label: '发送时间',
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
    render(row) {
      if (isNullObject(row.msgType)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.msg_type, row.msgType),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.msg_type, row.msgType),
        }
      );
    },
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
    title: '发送状态',
    key: 'sendStatus',
    render(row) {
      debugger
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

  {
    title: '备注',
    key: 'comment',
  },
];

async function loadOptions() {
  options.value = await Dicts({
    types: [
        'msg_type',
        'read_status',
        'send_status',
   ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'msgType':
        item.componentProps.options = options.value.msg_type;
        break;
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
