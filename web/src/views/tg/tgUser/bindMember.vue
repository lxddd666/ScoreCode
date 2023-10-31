<template>
  <n-modal
    v-model:show="isShowModal"
    :mask-closable="false"
    preset="dialog"
    title="绑定员工"
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

      :columns="columnList"
      :request="loadDataTable"
      :row-key="(row) => row.id"
      ref="actionRef"
      @update:checked-row-keys="onCheckedRow"
      :scroll-x="1090"
      :resizeHeightOffset="-10000"
      size="small"
    >
    </BasicTable>
  </n-modal>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref} from 'vue';
import {useDialog, useMessage} from 'naive-ui'
import {schemas} from '@/views/org/user/model';
import {columns} from '@/views/org/user/columns';
import {adaModalWidth} from "@/utils/hotgo";
import {useRouter} from "vue-router";
import {BasicTable} from "@/components/Table";
import {List} from "@/api/org/user";
import {usePermission} from "@/hooks/web/usePermission";
import {BasicForm, useForm} from "@/components/Form";

const emit = defineEmits(['reloadTable', 'updateBindMemberShowModal', 'handleBindMember']);

interface Props {
  showModal: boolean;
}

const message = useMessage();
const dialogWidth = ref('75%');
const router = useRouter();
const {hasPermission} = usePermission();
const searchFormRef = ref<any>({});
const checkedIds = ref(0);
const dialog = useDialog();
const actionRef = ref();
const loadDataTable = async (res) => {
  return await List({
      ...{status: 1},
      ...searchFormRef.value?.formModel, ...res
    }
  );
};
let thisColumns = [{
  type: 'selection',
  multiple: false,
}];


const columnList = thisColumns.concat(columns)


const props = withDefaults(defineProps<Props>(), {
  showModal: false,
});

const isShowModal = computed({
  get: () => {
    return props.showModal;
  },
  set: (value) => {
    emit('updateBindMemberShowModal', value);
  },
});

const [register, {}] = useForm({
  gridProps: {cols: '1 s:1 m:2 l:3 xl:4 2xl:4'},
  labelWidth: 80,
  schemas,
});

function onNegativeClick() {

}

function onPositiveClick() {
  if (checkedIds.value !== 0) {
    emit('handleBindMember', checkedIds.value);
  }

}


onMounted(async () => {
  adaModalWidth(dialogWidth);
});


function onCheckedRow(rowKeys) {
  checkedIds.value = rowKeys[0];
}

function reloadTable() {
  actionRef.value.reload();
}

</script>

<style scoped lang="less">

</style>
