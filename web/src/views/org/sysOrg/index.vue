<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <div class="n-layout-page-header">
        <n-card :bordered="false" title="公司信息">
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
            type="primary"
            @click="addTable"
            class="min-left-space"
            v-if="hasPermission(['/org/edit'])"
          >
            <template #icon>
              <n-icon>
                <PlusOutlined/>
              </n-icon>
            </template>
            添加
          </n-button>
          <n-button
            type="error"
            @click="handleBatchDelete"
            :disabled="batchDeleteDisabled"
            class="min-left-space"
            v-if="hasPermission(['/org/delete'])"
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
            v-if="hasPermission(['/org/export'])"
          >
            <template #icon>
              <n-icon>
                <ExportOutlined/>
              </n-icon>
            </template>
            导出
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
    <Ports
      @reloadTable="reloadTable"
      @updateShowModal="updateShowPortsModal"
      :showModal="showPortsModal"
      :formParams="formParams"
    />
  </div>
</template>

<script lang="ts" setup>
import {h, reactive, ref} from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import {BasicTable, TableAction} from '@/components/Table';
import {BasicForm, useForm} from '@/components/Form/index';
import {usePermission} from '@/hooks/web/usePermission';
import {Delete, Export, List, Status} from '@/api/tg/sysOrg';
import {columns, newState, options, schemas, State} from './model';
import {DeleteOutlined, ExportOutlined, PlusOutlined} from '@vicons/antd';
import {useRouter} from 'vue-router';
import {getOptionLabel} from '@/utils/hotgo';
import Edit from './edit.vue';
import Ports from './ports.vue';

const {hasPermission} = usePermission();
const router = useRouter();
const actionRef = ref();
const dialog = useDialog();
const message = useMessage();
const searchFormRef = ref<any>({});
const batchDeleteDisabled = ref(true);
const checkedIds = ref([]);
const showModal = ref(false);
const formParams = ref<State>();

const showPortsModal = ref(false);

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
          auth: ['/org/edit'],
        },
        {
          label: '禁用',
          onClick: handleStatus.bind(null, record, 2),
          ifShow: () => {
            return record.status === 1;
          },
          auth: ['/org/status'],
        },
        {
          label: '启用',
          onClick: handleStatus.bind(null, record, 1),
          ifShow: () => {
            return record.status === 2;
          },
          auth: ['/org/status'],
        },
        {
          label: '修改端口数',
          onClick: handlePorts.bind(null, record),
          auth: ['/org/ports'],
        },
      ],
      dropDownActions: [
        {
          label: '查看详情',
          key: 'view',
          auth: ['/org/view'],
        },
        {
          label: '删除',
          key: 'delete',
          auth: ['/org/delete'],
        },
      ],
      select: (key: string) => {
        if (key === 'view') {
          return handleView(record);
        } else if (key === 'delete') {
          return handleDelete(record);
        }
      },
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

function updateShowModal(value: boolean) {
  showModal.value = value;
}

function updateShowPortsModal(value: boolean) {
  showPortsModal.value = value;
}

function onCheckedRow(rowKeys) {
  batchDeleteDisabled.value = rowKeys.length <= 0;
  checkedIds.value = rowKeys;
}

function reloadTable() {
  actionRef.value.reload();
}

function handleView(record: Recordable) {
  router.push({name: 'sysOrgView', params: {id: record.id}});
}

function handleEdit(record: Recordable) {
  showModal.value = true;
  formParams.value = newState(record as State);
}

function handlePorts(record: Recordable) {
  showPortsModal.value = true;
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

function handleStatus(record: Recordable, status: number) {
  Status({id: record.id, status: status}).then((_res) => {
    message.success('设为' + getOptionLabel(options.value.sys_normal_disable, status) + '成功');
    setTimeout(() => {
      reloadTable();
    });
  });
}
</script>

<style lang="less" scoped></style>
