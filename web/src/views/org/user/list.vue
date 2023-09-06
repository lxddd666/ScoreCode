<template>
  <div>
    <BasicForm
      @register="register"
      @submit="handleSubmit"
      @reset="handleReset"
      @keyup.enter="handleSubmit"
      ref="searchFormRef"
    >
      <template #statusSlot="{ model, field }">
        <n-input v-model:value="model[field]" />
      </template>
    </BasicForm>

    <BasicTable
      :openChecked="true"
      :columns="columns"
      :request="loadDataTable"
      :row-key="(row) => row.id"
      ref="actionRef"
      :actionColumn="actionColumn"
      @update:checked-row-keys="onCheckedRow"
      :scroll-x="1800"
    >
      <template #tableTitle>
        <n-button
          type="primary"
          @click="addTable"
          class="min-left-space"
          v-if="hasPermission(['/member/edit'])"
        >
          <template #icon>
            <n-icon>
              <PlusOutlined />
            </n-icon>
          </template>
          添加用户
        </n-button>
        <n-button
          type="error"
          @click="batchDelete"
          :disabled="batchDeleteDisabled"
          class="min-left-space"
          v-if="hasPermission(['/member/delete'])"
        >
          <template #icon>
            <n-icon>
              <DeleteOutlined />
            </n-icon>
          </template>
          批量删除
        </n-button>

        <n-button
          type="success"
          @click="handleInviteQR(userStore.info?.inviteCode)"
          class="min-left-space"
          v-if="userStore.loginConfig?.loginRegisterSwitch === 1"
        >
          <template #icon>
            <n-icon>
              <QrCodeOutline />
            </n-icon>
          </template>
          邀请注册
        </n-button>
      </template>
    </BasicTable>

    <Edit
      @reloadTable="reloadTable"
      @updateShowModal="updateShowModal"
      :showModal="showModal"
      :formParams="formParams"
    />

    <n-modal v-model:show="showQrModal" :show-icon="false" preset="dialog" title="邀请注册二维码">
      <n-form class="py-4">
        <div class="text-center">
          <qrcode-vue :value="qrParams.qrUrl" :size="220" class="canvas" style="margin: 0 auto" />
        </div>
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="() => (showQrModal = false)">关闭</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { h, onMounted, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { ActionItem, BasicTable, TableAction } from '@/components/Table';
  import { BasicForm } from '@/components/Form/index';
  import { Delete, List, ResetPwd, Status } from '@/api/org/user';
  import { columns } from './columns';
  import { DeleteOutlined, PlusOutlined } from '@vicons/antd';
  import { QrCodeOutline } from '@vicons/ionicons5';
  import { adaModalWidth } from '@/utils/hotgo';
  import { getRandomString } from '@/utils/charset';
  import { cloneDeep } from 'lodash-es';
  import QrcodeVue from 'qrcode.vue';
  import { addNewState, loadOptions, register } from './model';
  import { usePermission } from '@/hooks/web/usePermission';
  import { useUserStore } from '@/store/modules/user';
  import { LoginRoute } from '@/router';
  import { getNowUrl } from '@/utils/urlUtils';
  import Edit from './edit.vue';

  interface Props {
    type?: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    type: '-1',
  });

  const rules = {
    username: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入用户名',
    },
  };

  const { hasPermission } = usePermission();
  const userStore = useUserStore();
  const message = useMessage();
  const actionRef = ref();
  const dialog = useDialog();
  const showModal = ref(false);
  const searchFormRef = ref<any>({});
  const formRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const dialogWidth = ref('50%');
  const formParams = ref<any>();
  const showQrModal = ref(false);
  const qrParams = ref({
    name: '',
    qrUrl: '',
  });

  const actionColumn = reactive({
    width: 220,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      const downActions = getDropDownActions(record);
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '已启用',
            onClick: handleStatus.bind(null, record, 2),
            ifShow: () => {
              return record.status === 1 && record.id !== 1;
            },
            auth: ['/member/status'],
          },
          {
            label: '已禁用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 2 && record.id !== 1;
            },
            auth: ['/member/status'],
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
            auth: ['/member/edit'],
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
            auth: ['/member/delete'],
          },
        ],
        dropDownActions: downActions,
        select: (key) => {
          if (key === 0) {
            return handleResetPwd(record);
          }
          if (key === 102) {
            if (userStore.loginConfig?.loginRegisterSwitch !== 1) {
              message.error('管理员暂未开启此功能');
              return;
            }
            return handleInviteQR(record.inviteCode);
          }
        },
      });
    },
  });

  function getDropDownActions(record: Recordable): ActionItem[] {
    if (record.id === 1) {
      return [];
    }

    let list = [
      {
        label: '重置密码',
        key: 0,
      },
    ];

    if (userStore.loginConfig?.loginRegisterSwitch === 1) {
      list.push({
        label: 'TA的邀请码',
        key: 102,
      });
    }

    return list;
  }

  function addTable() {
    showModal.value = true;
    formParams.value = addNewState(null);
  }

  const loadDataTable = async (res) => {
    adaModalWidth(dialogWidth);
    return await List({ ...res, ...searchFormRef.value?.formModel, ...{ roleId: props.type } });
  };

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleEdit(record: Recordable) {
    showModal.value = true;
    formParams.value = cloneDeep(record);
  }

  function updateShowModal(value) {
    showModal.value = value;
  }

  function handleResetPwd(record: Recordable) {
    record.password = getRandomString(12);
    dialog.warning({
      title: '警告',
      content: '你确定要重置密码？\r\n重置成功后密码为：' + record.password + '\r\n 请先保存',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        ResetPwd(record).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
    });
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
    });
  }

  function batchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
    });
  }

  function handleSubmit(_values: Recordable) {
    reloadTable();
  }

  function handleReset(_values: Recordable) {
    reloadTable();
  }

  function handleStatus(record: Recordable, status) {
    Status({ id: record.id, status: status }).then((_res) => {
      message.success('操作成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }

  function handleInviteQR(code: any) {
    const domain = getNowUrl() + '#';
    qrParams.value.qrUrl = domain + LoginRoute.path + '?scope=register&inviteCode=' + code;
    showQrModal.value = true;
  }

  onMounted(async () => {
    await loadOptions();
  });
</script>

<style lang="less" scoped></style>
