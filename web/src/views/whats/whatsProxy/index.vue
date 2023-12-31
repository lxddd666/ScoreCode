<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <div class="n-layout-page-header">
        <n-card :bordered="false" title="代理管理">
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
        :row-key="(row) => row"
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
            v-if="hasPermission(['/whatsProxy/edit'])"
          >
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
          <n-button
            type="error"
            @click="handleBatchDelete"
            :disabled="batchDeleteDisabled"
            class="min-left-space"
            v-if="hasPermission(['/whatsProxy/delete'])"
          >
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
          <n-button
            type="primary"
            @click="handleUpload"
            class="min-left-space"
            v-if="hasPermission(['/whatsProxy/view'])"
          >
            <template #icon>
              <n-icon>
                <UploadOutlined />
              </n-icon>
            </template>
            导入
          </n-button>
          <n-button
            type="primary"
            @click="handleExport"
            class="min-left-space"
            v-if="hasPermission(['/whatsProxy/view'])"
          >
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            导出
          </n-button>
          <n-button
            type="primary"
            @click="addOrgTable"
            class="min-left-space"
            v-if="hasPermission(['/whatsProxy/view'])"
          >
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            绑定公司
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
    <UnBind
      @updateUnBindShowModal="updateUnBindShowModal"
      :showModal="unBindShowModal"
      :formParams="formParams"
    />
    <Bind
      @updateBindShowModal="updateBindShowModal"
      :showModal="bindShowModal"
      :formparams="formParams"
    />
    <AddProxyToOrg
      @updateAddProxyToOrg="updateAddProxyToOrg"
      :showModal="showOrgModal"
      :formParams="formParams"
      :address="proxyUrls"
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
  import { List, Export, Delete, Status } from '@/api/whats/whatsProxy';
  import { State, columns, schemas, options, newState } from './model';
  import { PlusOutlined, ExportOutlined, DeleteOutlined, UploadOutlined } from '@vicons/antd';
  import { useRouter } from 'vue-router';
  import { getOptionLabel } from '@/utils/hotgo';
  import Edit from './edit.vue';
  import UnBind from '@/views/whats/whatsProxy/unBind.vue';
  import Bind from '@/views/whats/whatsProxy/bind.vue';
  import AddProxyToOrg from '@/views/whats/whatsProxy/addProxyToOrg.vue';
  import { getRandomString } from '@/utils/charset';
  import { ResetPwd } from '@/api/org/user';
  import * as url from "url";
  const { hasPermission } = usePermission();
  const router = useRouter();
  const actionRef = ref();
  const dialog = useDialog();
  const message = useMessage();
  const searchFormRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const proxyUrls = ref([]);
  const showModal = ref(false);
  const showOrgModal = ref(false);
  const unBindShowModal = ref(false);
  const bindShowModal = ref(false);
  const formParams = ref<State>();
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
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            auth: ['/whatsProxy/edit'],
          },
          {
            label: '禁用',
            onClick: handleStatus.bind(null, record, 2),
            ifShow: () => {
              return record.status === 1;
            },
            auth: ['/whatsProxy/status'],
          },
          {
            label: '启用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 2;
            },
            auth: ['/whatsProxy/status'],
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/whatsProxy/delete'],
          },
        ],
        dropDownActions: [
          {
            label: '绑定账号',
            key: 'bind',
            auth: ['/whatsProxy/bind'],
          },
          {
            label: '解绑账号',
            key: 'unBind',
            auth: ['/whatsProxy/unBind'],
          },
          {
            label: '查看详情',
            key: 'view',
            auth: ['/whatsProxy/view'],
          },
        ],
        select: (key) => {
          if (key === 'view') {
            return handleView(record);
          } else if (key === 'unBind') {
            return handleUnbind(record);
          } else if (key === 'bind') {
            return handleBind(record);
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
    showModal.value = true;
    formParams.value = newState(null);
  }

  function addOrgTable() {
    showOrgModal.value = true;
    formParams.value = newState(null);
  }

  function handleUpload() {
    fileUploadRef.value.openModal();
  }

  function updateShowModal(value) {
    showModal.value = value;
  }

  function updateUnBindShowModal(value) {
    unBindShowModal.value = value;
  }
  function updateBindShowModal(value) {
    bindShowModal.value = value;
  }

  function updateAddProxyToOrg(value) {
    showOrgModal.value = value;
  }

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    let ids = [];
    let urls = [];
    for (let i = 0; i < rowKeys.length; i++) {;
      ids.push(rowKeys[i].id);
      urls.push(rowKeys[i].address);
    }
    proxyUrls.value = urls;
    checkedIds.value = ids;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleView(record: Recordable) {
    router.push({ name: 'whatsProxyView', query: { id: record.id, address: record.address } });
  }

  function handleEdit(record: Recordable) {
    showModal.value = true;
    formParams.value = newState(record as State);
  }
  function handleUnbind(record: Recordable) {
    unBindShowModal.value = true;
    formParams.value = newState(record as State);
  }
  function handleBind(record: Recordable) {
    bindShowModal.value = true;
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

  function handleExport() {
    message.loading('正在导出列表...', { duration: 1200 });
    Export(searchFormRef.value?.formModel);
  }

  function handleStatus(record: Recordable, status: number) {
    Status({ id: record.id, status: status }).then((_res) => {
      message.success('设为' + getOptionLabel(options.value.sys_normal_disable, status) + '成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }

  function handleFinishCall(result: Attachment, success: boolean) {
    if (success) {
      reloadTable();
    }
  }
</script>

<style lang="less" scoped></style>
