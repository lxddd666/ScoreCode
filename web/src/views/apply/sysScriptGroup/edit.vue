<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
        :title="params?.id > 0 ? '编辑 #' + params?.id : '添加'"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-form
          :model="params"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="组织ID" path="orgId">
            <n-input-number placeholder="请输入组织ID" v-model:value="params.orgId" />
          </n-form-item>

          <n-form-item label="部门ID" path="deptId">
            <n-input-number placeholder="请输入部门ID" v-model:value="params.deptId" />
          </n-form-item>

          <n-form-item label="用户ID" path="memberId">
            <n-input-number placeholder="请输入用户ID" v-model:value="params.memberId" />
          </n-form-item>

          <n-form-item label="分组类型" path="type">
            <n-select v-model:value="params.type" :options="options.script_type" />
          </n-form-item>

          <n-form-item label="自定义组名" path="name">
            <n-input placeholder="请输入自定义组名" v-model:value="params.name" />
          </n-form-item>

          <n-form-item label="话术数量" path="scriptCount">
            <n-input-number placeholder="请输入话术数量" v-model:value="params.scriptCount" />
          </n-form-item>


        </n-form>
        <template #action>
          <n-space>
            <n-button @click="closeForm">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, computed, watch } from 'vue';
import { Edit, View } from '@/api/apply/sysScriptGroup';
import { rules, options, State, newState } from './model';
import { useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';

const emit = defineEmits(['reloadTable', 'updateShowModal']);

interface Props {
  showModal: boolean;
  formParams?: State;
}

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
    emit('updateShowModal', value);
  },
});

const loading = ref(false);
const params = ref<State>(props.formParams);
const message = useMessage();
const formRef = ref<any>({});
const dialogWidth = ref('75%');
const formBtnLoading = ref(false);

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(params.value).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          isShowModal.value = false;
          emit('reloadTable');
        });
      });
    } else {
      message.error('请填写完整信息');
    }
    formBtnLoading.value = false;
  });
}

onMounted(async () => {
  adaModalWidth(dialogWidth);
});

function closeForm() {
  isShowModal.value = false;
}

function loadForm(value) {
  // 新增
  if (value.id < 1) {
    params.value = newState(value);
    loading.value = false;
    return;
  }

  loading.value = true;
  // 编辑
  View({ id: value.id })
    .then((res) => {
      params.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

watch(
  () => props.formParams,
  (value) => {
    loadForm(value);
  }
);
</script>

<style lang="less"></style>
