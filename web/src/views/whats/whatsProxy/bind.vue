<template>
  <n-modal
      v-model:show="isShowModal"
      :mask-closable="false"
      preset="dialog"
      title="绑定帐号"
      content=""
      positive-text="确认"
      negative-text="取消"
      @positive-click="onPositiveClick"
      @negative-click="onNegativeClick"
      :style="{
          width: dialogWidth,
        }"
  >
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
      @update:checked-row-keys="onCheckedRow"
      :scroll-x="1090"
      :resizeHeightOffset="-10000"
      size="small"
    >
      <template #tableTitle>
        <n-button
          type="error"
          @click="handleBatchBind"
          :disabled="batchBindDisabled"
          class="min-left-space"
          v-if="hasPermission(['/whatsAccount/delete'])"
        >
          <template #icon>
            <n-icon>
              <PlusOutlined/>
            </n-icon>
          </template>
          绑定账号
        </n-button>
      </template>
    </BasicTable>
  </n-modal>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref, watch} from 'vue';
import {useDialog, useMessage} from 'naive-ui'
import {schemas} from './accountModel';
import {columns, newState, State} from '../whatsAccount/model';
import {adaModalWidth} from "@/utils/hotgo";
import {View} from "@/api/whats/whatsProxy";
import {useRouter} from "vue-router";
import {PlusOutlined} from "@vicons/antd";
import {BasicTable} from "@/components/Table";
import {Bind, List} from "@/api/whats/whatsAccount";
import {usePermission} from "@/hooks/web/usePermission";
import {BasicForm, useForm} from "@/components/Form";

const emit = defineEmits([ 'updateBindShowModal']);

interface Props {
  showModal: boolean;
  formParams?: State;
}

const message = useMessage()
const dialogWidth = ref('75%');
const router = useRouter();
const id = Number(router.currentRoute.value.query.id);
const formValue = ref(newState(null));
const batchBindDisabled = ref(true);
const {hasPermission} = usePermission();
const searchFormRef = ref<any>({});
const checkedIds = ref([]);
const dialog = useDialog();
const actionRef = ref();
const loadDataTable = async (res) => {
  return await List({
      ...{unbind: true},
      ...searchFormRef.value?.formModel, ...res
    }
  );
};

const props = withDefaults(defineProps<Props>(), {
  showModal: false,
  formParams: () => {
    return newState(null);
  },
});

const isShowModal = computed({
  get: () => {
    return props.showModal;
  },
  set: (value) => {
    emit('updateBindShowModal', value);
  },
});

const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

function onNegativeClick(e) {
  message.success('Cancel')
}

function onPositiveClick(e) {
  message.success('Submit')
}


onMounted(async () => {
  adaModalWidth(dialogWidth);
});

function closeForm() {
  isShowModal.value = false;
}

function loadForm(value) {
  console.log("bind--",value)

}
function onCheckedRow(rowKeys) {
  batchBindDisabled.value = rowKeys.length <= 0;
  checkedIds.value = rowKeys;
}
function reloadTable() {
  actionRef.value.reload();
}
function handleBatchBind() {
  dialog.warning({
    title: '警告',
    content: '你确定要绑定账号吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Bind({id: checkedIds.value, proxyAddress: props.formParams.address}).then((_res) => {
        message.success('绑定成功');
        emit('reloadView');
        reloadTable();
      });
    },
    onNegativeClick: () => {
      //message.error('取消');
    },
  });
}
onMounted(async () => {
  if (id < 1) {
    message.error('id不正确，请检查！');
    return;
  }
  formValue.value = await View({ id: id });
});

async function reloadView() {
  formValue.value = await View({id: id});
}
watch(
    () => props.formParams,
    (value) => {
      loadForm(value);
    }
);
</script>

<style scoped lang="less">

</style>
