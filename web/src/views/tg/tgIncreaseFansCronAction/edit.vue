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
          :label-width="200"
          class="py-4"
        >
          <n-form-item label="任务ID" path="cronId">
            <n-input-number placeholder="请输入任务ID" v-model:value="params.cronId" />
          </n-form-item>

          <n-form-item label="加入频道的userId" path="tgUserId">
            <n-input-number placeholder="请输入加入频道的userId" v-model:value="params.tgUserId" />
          </n-form-item>

          <n-form-item label="加入状态：0失败，1成功，2完成" path="joinStatus">
            <n-select v-model:value="params.joinStatus" :options="options.sys_normal_disable" />
          </n-form-item>

          <n-form-item label="备注" path="comment">
            <n-input type="textarea" placeholder="备注" v-model:value="params.comment" />
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
  import { Edit, View } from '@/api/org/tgIncreaseFansCronAction';
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
