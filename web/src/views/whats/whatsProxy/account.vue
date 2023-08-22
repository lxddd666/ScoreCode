<template>
  <div>
    <n-card :bordered="false" class="proCard">
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
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="1090"
        :resizeHeightOffset="-10000"
        size="small"
      >
        <template #tableTitle>
          <n-button
            type="error"
            @click="handleBatchUnbind"
            :disabled="batchUnBindDisabled"
            class="min-left-space"
            v-if="hasPermission(['/whatsAccount/delete'])"
          >
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            解除绑定
          </n-button>

        </template>
      </BasicTable>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import {ref} from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import {BasicTable} from '@/components/Table';
import {BasicForm, useForm} from '@/components/Form/index';
import {usePermission} from '@/hooks/web/usePermission';
import {List, UnBind} from '@/api/whats/whatsAccount';
import {schemas} from './accountModel';
import {columns, newState, State} from '../whatsAccount/model';
import {useRouter} from 'vue-router';
import {DeleteOutlined} from "@vicons/antd";

const emit = defineEmits(['reloadView']);

const { hasPermission } = usePermission();
  const router = useRouter();
  const actionRef = ref();
  const dialog = useDialog();
  const message = useMessage();
  const searchFormRef = ref<any>({});
  const batchUnBindDisabled = ref(true);
  const checkedIds = ref([]);
  const showModal = ref(false);
  const formParams = ref<State>();
  const address = String(router.currentRoute.value.query.address);
  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  const loadDataTable = async (res) => {
    return await List({
      ...{ proxyAddress: address },
      ...searchFormRef.value?.formModel, ...res}
    );
  };

  function addTable() {
    showModal.value = true;
    formParams.value = newState(null);
  }

  function updateShowModal(value) {
    showModal.value = value;
  }
  function onCheckedRow(rowKeys) {
    batchUnBindDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleBatchUnbind() {
    dialog.warning({
      title: '警告',
      content: '你确定要解绑吗？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        UnBind({ id: checkedIds.value,proxyAddress: address }).then((_res) => {
          message.success('解绑成功');
          emit('reloadView');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }
</script>

<style lang="less" scoped></style>
