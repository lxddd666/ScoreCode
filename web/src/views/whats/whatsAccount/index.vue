<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <div class="n-layout-page-header">
        <n-card :bordered="false" title="账号管理">
          <!--  这是系统自动生成的CURD表格，你可以将此行注释改为表格的描述 -->
        </n-card>
      </div>

      <BasicForm
        @register="register"
        @submit="reloadTable"
        @reset="reloadTable"
        @keyup.enter="reloadTable"
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
        :scroll-x="1090"
        :resizeHeightOffset="-10000"
        size="small"
      >
        <template #tableTitle>
          <n-button
            type="primary"
            @click="addTable"
            class="min-left-space"
            v-if="hasPermission(['/whatsAccount/edit'])"
          >
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
          <n-button
            type="primary"
            @click="handleUpload"
            class="min-left-space"
            v-if="hasPermission(['/whatsAccount/edit'])"
          >
            <template #icon>
              <n-icon>
                <UploadOutlined />
              </n-icon>
            </template>
            导入
          </n-button>
          <n-button
            color="#49CC90"
            @click="handleBatchLogin"
            :disabled="batchSelectDisabled"
            class="min-left-space"
            v-if="hasPermission(['/whats/login'])"
          >
            <template #icon>
              <n-icon>
                <LoginOutlined />
              </n-icon>
            </template>
            登录
          </n-button>

          <n-button
            type="error"
            @click="handleBatchDelete"
            :disabled="batchSelectDisabled"
            class="min-left-space"
            v-if="hasPermission(['/whatsAccount/delete'])"
          >
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <Edit
      @reloadTable="reloadTable"
      @updateShowModal="updateShowModal"
      :showModal="showEditModal"
      :formParams="formParams"
    />
    <SendMsg
      @reloadTable="reloadTable"
      @sendMsgShowModal="sendMsgShowModal"
      :showModal="showSendModal"
      :sender="account"
    />
    <SendFile
      @reloadTable="reloadTable"
      @sendMsgShowModal="sendFileShowModal"
      :showModal="showSendFileModal"
      :sender="account"
    />
    <SendVcardMsg
      @reloadTable="reloadTable"
      @sendMsgShowModal="sendVcardMsgShowModal"
      :showModal="showSendVcardModel"
      :sender="account"
    />
    <SyncContact
      @reloadTable="reloadTable"
      @sendSyncContactModel="sendSyncContactModel"
      :showModal="showSyncContactModel"
      :sender="account"
    />

    <FileUpload @reloadTable="reloadTable" ref="fileUploadRef" :finish-call="handleFinishCall" />
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import {Delete, List, Login, Logout} from '@/api/whats/whatsAccount';
  import { columns, newState, schemas, State } from './model';
  import { DeleteOutlined, LoginOutlined, PlusOutlined, UploadOutlined } from '@vicons/antd';
  import { useRouter } from 'vue-router';
  import Edit from './edit.vue';
  import SendMsg from './sendMsg.vue';
  import SendVcardMsg from '@/views/whats/whatsAccount/sendVcardMsg.vue';
  import FileUpload from './upload.vue';
  import { Attachment } from '@/components/FileChooser/src/model';
  import SyncContact from '@/views/whats/whatsAccount/syncContact.vue';
  import SendFile from "@/views/whats/whatsAccount/sendFile.vue";

  const { hasPermission } = usePermission();
  const router = useRouter();
  const actionRef = ref();
  const dialog = useDialog();
  const message = useMessage();
  const searchFormRef = ref<any>({});
  const batchSelectDisabled = ref(true);
  const checkedIds = ref([]);
  const showEditModal = ref(false);
  const showSendModal = ref(false);
  const showSendFileModal = ref(false);
  const showSyncContactModel = ref(false);
  const showSendVcardModel = ref(false);
  const formParams = ref<State>();
  const account = ref<string>();

  const fileUploadRef = ref();

  const actionColumn = reactive({
    width: 350,
    title: '操作',
    key: 'action',
    // fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            auth: ['/whatsAccount/edit'],
          },

          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/whatsAccount/delete'],
          },
          {
            label: '发送消息',
            onClick: handleSendMsg.bind(null, record),
            auth: ['/whats/sendMsg'],
          },
          {
            label: '发送文件',
            onClick: handleSendFile.bind(null, record),
            auth: ['/whats/sendFile'],
          },
        ],
        dropDownActions: [
          {
            label: '查看详情',
            key: 'view',
            auth: ['/whatsMsg/view'],
          },
          {
            label: '发送名片',
            key: 'sendVcardMsg',
            auth: ['/whats/sendVcardMsg'],
          },
          {
            label: '同步联系人',
            key: 'syncContact',
            auth: ['/whatsMsg/view'],
          },
          {
            label: '登出',
            key: 'logout',
            auth: ['/whatsMsg/view'],
          },
        ],
        select: (key) => {
          if (key === 'view') {
            return handleView(record);
          } else if (key === 'syncContact') {
            return handleSyncContact(record);
          } else if (key === 'logout') {
            return handleLogout(record);
          }else if(key=== 'sendVcardMsg'){
            return handleSendVcardMsg(record);
          }
        },
      });
    },
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  const loadDataTable = async (res) => {
    return await List({ ...searchFormRef.value?.formModel, ...res });
  };

  function addTable() {
    showEditModal.value = true;
    formParams.value = newState(null);
  }

  function updateShowModal(value) {
    showEditModal.value = value;
  }

  function sendMsgShowModal(value) {
    showSendModal.value = value;
  }

  function sendFileShowModal(value) {
    showSendFileModal.value = value;
  }

  function sendVcardMsgShowModal(value) {
    showSendVcardModel.value = value;
  }
  function sendSyncContactModel(value) {
    showSyncContactModel.value = value;
  }

  function handleUpload() {
    fileUploadRef.value.openModal();
  }

  function onCheckedRow(rowKeys) {
    batchSelectDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleView(record: Recordable) {
    router.push({ name: 'whatsAccountView', params: { id: record.id } });
  }

  function handleEdit(record: Recordable) {
    showEditModal.value = true;
    formParams.value = newState(record as State);
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
          message.success('删除成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function handleSendMsg(record: Recordable) {
    showSendModal.value = true;
    account.value = newState(record as State).account;
  }

  function handleSendFile(record: Recordable) {
    showSendFileModal.value = true;
    account.value = newState(record as State).account;
  }

  function handleSyncContact(record: Recordable) {
    showSyncContactModel.value = true;
    account.value = newState(record as State).account;
  }
  function handleLogout(record: Recordable) {
    account.value = newState(record as State).account;
    debugger;
    dialog.warning({
      title: '退出登录',
      content: '你确定退出登录吗？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Logout({ logoutDetail: [{ account: record.account, proxy: record.proxyAddress }] }).then(
          (_res) => {
            message.success('退出成功');
            reloadTable();
          }
        );
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function handleSendVcardMsg(record: Recordable) {
    showSendVcardModel.value = true;
    account.value = newState(record as State).account;
  }

  function handleBatchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要批量删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          message.success('删除成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function handleBatchLogin() {
    dialog.info({
      title: '提示',
      content: '点击确定将执行批量登录，登录结果请刷新页面查看登录状态',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Login({ ids: checkedIds.value }).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function handleFinishCall(result: Attachment, success: boolean) {
    if (success) {
      reloadTable();
    }
  }
</script>

<style lang="less" scoped></style>
