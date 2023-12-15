<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <div class="n-layout-page-header">
        <n-card :bordered="false" title="TG账号">
          <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
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
          <n-input v-model:value="model[field]"/>
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
              type="error"
              @click="handleBatchDelete"
              :disabled="batchSelectDisabled"
              class="min-left-space"
              v-if="hasPermission(['/tgUser/delete'])"
          >
            <template #icon>
              <n-icon>
                <DeleteOutlined/>
              </n-icon>
            </template>
            批量删除
          </n-button>
          <n-button
              type="primary"
              @click="handleExport"
              class="min-left-space"
              v-if="hasPermission(['/tgUser/export'])"
          >
            <template #icon>
              <n-icon>
                <ExportOutlined/>
              </n-icon>
            </template>
            导出
          </n-button>
          <n-button
              type="success"
              @click="bindMemberClick"
              :disabled="batchSelectDisabled"
              class="min-left-space"
              v-if="hasPermission(['/tgUser/bindMember'])"
          >
            绑定员工
          </n-button>
          <n-button
              type="warning"
              @click="handleUnBindMember"
              :disabled="batchSelectDisabled"
              class="min-left-space"
              v-if="hasPermission(['/tgUser/unBindMember'])"
          >
            解绑员工
          </n-button>

          <n-button
              type="success"
              @click="bindProxyClick"
              :disabled="batchSelectDisabled"
              class="min-left-space"
              v-if="hasPermission(['/tgUser/bindProxy'])"
          >
            绑定代理
          </n-button>

          <n-button
              type="warning"
              @click="handleUnBindProxy"
              :disabled="batchSelectDisabled"
              class="min-left-space"
              v-if="hasPermission(['/tgUser/unBindProxy'])"
          >
            解绑代理
          </n-button>

          <n-button
              type="success"
              @click="handleBatchLogin"
              :disabled="batchSelectDisabled"
              class="min-left-space"
              v-if="hasPermission(['/arts/batchLogin'])"
          >
            <template #icon>
              <n-icon>
                <LoginOutlined/>
              </n-icon>
            </template>
            批量上线
          </n-button>

          <n-button
              type="error"
              @click="handleBatchLogout"
              :disabled="batchSelectDisabled"
              class="min-left-space"
              v-if="hasPermission(['/arts/batchLogout'])"
          >
            <template #icon>
              <n-icon>
                <LogoutOutlined/>
              </n-icon>
            </template>
            批量下线
          </n-button>

          <n-button
              type="error"
              @click="handleUpload"
              class="min-left-space"
              v-if="hasPermission(['/arts/batchLogout'])"
          >
            <n-icon>
              <UploadOutlined/>
            </n-icon>
            批量导入session
          </n-button>
          <n-button
              type="info"
              @click="handleCodeLogin"
              class="min-left-space"
              v-if="hasPermission(['/arts/codeLogin'])"
          >
            <template #icon>
              <n-icon>
                <LoginOutlined/>
              </n-icon>
            </template>
            手机验证码登录
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <Edit
        @reloadTable="reloadTable"
        @updateShowModal="updateShowModal"
        :showModal="showModal"
        :formParams="formParams"
    />
    <BindMember
        @reloadTable="reloadTable"
        @updateBindMemberShowModal="updateBindMemberShowModal"
        @handleBindMember="handleBindMember"
        :showModal="bindMemberShowModal"
    />
    <FolderMember
        @reloadTable="reloadTable"
        @updateFolderMemberShowModal="updateFolderMemberShowModal"
        @handleFolderMember="handleFolderMember"
        :showModal="folderMemberShowModal"
    />
    <BindProxy
        @reloadTable="reloadTable"
        @updateBindProxyShowModal="updateBindProxyShowModal"
        @handleBindProxy="handleBindProxy"
        :showModal="bindProxyShowModal"
    />
    <FileUpload @reloadTable="reloadTable" ref="fileUploadRef" :finish-call="handleFinishCall"/>
    <Login
        @reloadTable="reloadTable"
        @updateShowModal="updateLoginShowModal"
        :showModal="loginShowModal"
    />

  </div>
</template>

<script lang="ts" setup>
import {h, reactive, ref} from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import {BasicTable, TableAction} from '@/components/Table';
import {BasicForm, useForm} from '@/components/Form/index';
import {usePermission} from '@/hooks/web/usePermission';
import FileUpload from './uploadSession.vue';
import {
  Delete,
  Export,
  List,
  TgBathLogin,
  TgBathLogout,
  TgBindMember,
  TgBindProxy,
  TgUnBindMember,
  TgUnBindProxy,
} from '@/api/tg/tgUser';
import {columns, newState, schemas, State} from './model';
import {DeleteOutlined, ExportOutlined, LoginOutlined, LogoutOutlined, UploadOutlined,} from '@vicons/antd';
import {useRouter} from 'vue-router';
import Edit from './edit.vue';
import Login from './login.vue';
import BindMember from './bindMember.vue';
import BindProxy from './bindProxy.vue';
import FolderMember from './folderMember.vue';
import {Attachment} from '@/components/FileChooser/src/model';

const {hasPermission} = usePermission();
const router = useRouter();
const actionRef = ref();
const dialog = useDialog();
const message = useMessage();
const searchFormRef = ref<any>({});
const batchSelectDisabled = ref(true);
const checkedIds = ref([]);
const showModal = ref(false);
const loginShowModal = ref(false);
const formParams = ref<State>();

const bindMemberShowModal = ref(false);
const folderMemberShowModal = ref(false);
const bindProxyShowModal = ref(false);
const fileUploadRef = ref();

const actionColumn = reactive({
  width: 300,
  title: '操作',
  key: 'action',
  // fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '聊天室',
          onClick: handleChat.bind(null, record),
          auth: ['/arts/batchLogin'],
        },
        {
          label: '编辑',
          onClick: handleEdit.bind(null, record),
          auth: ['/tgUser/edit'],
        },

        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          auth: ['/tgUser/delete'],
        },
      ],
      // dropDownActions: [
      //   {
      //     label: '查看详情',
      //     key: 'view',
      //     auth: ['/tgUser/view'],
      //   },
      // ],
      // select: (key) => {
      //   if (key === 'view') {
      //     return handleView(record);
      //   }
      // },
    });
  },
});

const [register, {}] = useForm({
  gridProps: {cols: '1 s:1 m:2 l:3 xl:4 2xl:4'},
  labelWidth: 80,
  schemas,
});

const loadDataTable = async (res) => {
  return await List({...searchFormRef.value?.formModel, ...res});
};

function addTable() {
  showModal.value = true;
  formParams.value = newState(null);
}

function updateShowModal(value) {
  showModal.value = value;
}

const updateLoginShowModal = (value: boolean) => {
  loginShowModal.value = value;
}

function onCheckedRow(rowKeys) {
  batchSelectDisabled.value = rowKeys.length <= 0;
  checkedIds.value = rowKeys;
}

function reloadTable() {
  actionRef.value.reload();
}

function handleView(record: Recordable) {
  router.push({name: 'tgUserView', params: {id: record.id}});
}

function handleEdit(record: Recordable) {
  showModal.value = true;
  formParams.value = newState(record as State);
}

function handleChat(record: Recordable) {
  router.push({name: 'tgChat', params: {id: record.id}});
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

function handleBatchDelete() {
  dialog.warning({
    title: '警告',
    content: '你确定要批量删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete({id: checkedIds.value}).then((_res) => {
        message.success('删除成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {
      // message.error('取消');
    },
  });
}

function handleExport() {
  message.loading('正在导出列表...', {duration: 1200});
  Export(searchFormRef.value?.formModel);
}

function updateBindMemberShowModal(value: boolean) {
  bindMemberShowModal.value = value;
}

function updateFolderMemberShowModal(value: boolean) {
  folderMemberShowModal.value = value;
}

function updateBindProxyShowModal(value: boolean) {
  bindProxyShowModal.value = value;
}

function bindMemberClick() {
  bindMemberShowModal.value = true;
}

function bindProxyClick() {
  bindProxyShowModal.value = true;
}

function handleBindMember(memberId: number) {
  TgBindMember({memberId: memberId, ids: checkedIds.value}).then((_res) => {
    message.success('绑定成功');
    reloadTable();
  });
}

function handleFolderMember(memberId: number) {
  TgBindMember({memberId: memberId, ids: checkedIds.value}).then((_res) => {
    message.success('绑定成功');
    reloadTable();
  });
}

function handleUnBindMember() {
  dialog.warning({
    title: '警告',
    content: '你确定要解除绑定吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      TgUnBindMember({ids: checkedIds.value}).then((_res) => {
        message.success('解绑成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {
      // message.error('取消');
    },
  });
}

function handleBindProxy(id: number) {
  TgBindProxy({proxyId: id, ids: checkedIds.value}).then((_res) => {
    message.success('绑定成功');
    reloadTable();
  });
}

function handleUnBindProxy() {
  dialog.warning({
    title: '警告',
    content: '你确定要解除绑定吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      TgUnBindProxy({ids: checkedIds.value}).then((_res) => {
        message.success('解绑成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {
      // message.error('取消');
    },
  });
}

function handleBatchLogin() {
  TgBathLogin({ids: checkedIds.value}).then((_res) => {
    message.success('登录中，请等待......');
    reloadTable();
  });
}

function handleUpload() {
  fileUploadRef.value.openModal();
}

function handleBatchLogout() {
  dialog.warning({
    title: '警告',
    content: '你确定要退出吗，退出后将无法接收最新消息？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      TgBathLogout({ids: checkedIds.value}).then((_res) => {
        message.success('下线成功');
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

const handleCodeLogin = () => {
  loginShowModal.value = true;
}

</script>

<style lang="less" scoped></style>
