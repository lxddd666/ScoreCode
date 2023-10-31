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
            type="primary"
            @click="addTable"
            class="min-left-space"
            v-if="hasPermission(['/tgUser/edit'])"
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
            type="primary"
            @click="bindProxyClick"
            :disabled="batchSelectDisabled"
            class="min-left-space"
            v-if="hasPermission(['/tgUser/bindProxy'])"
          >
            绑定代理
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
    <BindProxy
      @reloadTable="reloadTable"
      @updateBindProxyShowModal="updateBindProxyShowModal"
      @handleBindProxy="handleBindProxy"
      :showModal="bindProxyShowModal"
    />
  </div>
</template>

<script lang="ts" setup>
import {h, reactive, ref} from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import {BasicTable, TableAction} from '@/components/Table';
import {BasicForm, useForm} from '@/components/Form/index';
import {usePermission} from '@/hooks/web/usePermission';
import {Delete, Export, List, TgBindMember, TgBindProxy} from '@/api/tg/tgUser';
import {columns, newState, schemas, State} from './model';
import {DeleteOutlined, ExportOutlined, PlusOutlined} from '@vicons/antd';
import {useRouter} from 'vue-router';
import Edit from './edit.vue';
import BindMember from "./bindMember.vue";
import BindProxy from "./bindProxy.vue";

const {hasPermission} = usePermission();
const router = useRouter();
const actionRef = ref();
const dialog = useDialog();
const message = useMessage();
const searchFormRef = ref<any>({});
const batchSelectDisabled = ref(true);
const checkedIds = ref([]);
const showModal = ref(false);
const formParams = ref<State>();

const bindMemberShowModal = ref(false);
const bindProxyShowModal = ref(false);

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

function handleBindProxy(id: number) {
  TgBindProxy({proxyId: id, ids: checkedIds.value}).then((_res) => {
    message.success('绑定成功');
    reloadTable();
  });
}

</script>

<style lang="less" scoped></style>
