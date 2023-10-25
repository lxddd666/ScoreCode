<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
        title="修改端口数"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-form
          :model="params"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="200"
          class="py-4"
        >
          <n-form-item label="端口数" path="ports">
          <n-input-number placeholder="请输入公司名称" v-model:value="params.ports" />
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
  import {Ports, View} from '@/api/tg/sysOrg';
  import { rules, State, newState } from './model';
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

  function confirmForm(e: { preventDefault: () => void; }) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors: any) => {
      if (!errors) {
        Ports(params.value).then((_res) => {
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
