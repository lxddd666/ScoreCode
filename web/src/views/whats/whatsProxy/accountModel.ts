import {ref} from 'vue';
import {FormSchema} from '@/components/Form';
import {Dicts} from '@/api/dict/dict';
import {Options} from '@/utils/hotgo';


export const options = ref<Options>({
  account_status: [],
  login_status: [],
});

export const rules = {};

export const schemas = ref<FormSchema[]>([
  // {
  //   field: 'createdAt',
  //   component: 'NDatePicker',
  //   label: '创建时间',
  //   componentProps: {
  //     type: 'datetimerange',
  //     clearable: true,
  //     shortcuts: defRangeShortcuts(),
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  {
    field: 'account',
    component: 'NInput',
    label: '账号号码',
    defaultValue: null,
    componentProps: {
      clearable: true,
      placeholder: '请输入账号号码',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'accountStatus',
    component: 'NSelect',
    label: '账号状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择账号状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

async function loadOptions() {
  options.value = await Dicts({
    types: ['account_status', 'login_status'],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'accountStatus':
        item.componentProps.options = options.value.account_status;
        break;
      case 'isOnline':
        item.componentProps.options = options.value.login_status;
        break;
    }
  }
}

await loadOptions();
