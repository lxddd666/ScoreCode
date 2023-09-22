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
          <n-form-item label="聊天发起人" path="initiator">
            <n-input-number placeholder="请输入聊天发起人" v-model:value="params.initiator" />
          </n-form-item>

          <n-form-item label="发送人" path="sender">
            <n-input-number placeholder="请输入发送人" v-model:value="params.sender" />
          </n-form-item>

          <n-form-item label="接收人" path="receiver">
            <n-input-number placeholder="请输入接收人" v-model:value="params.receiver" />
          </n-form-item>

          <n-form-item label="请求id" path="reqId">
          <n-input placeholder="请输入请求id" v-model:value="params.reqId" />
          </n-form-item>

          <n-form-item label="发送消息原文(加密)" path="sendMsg">
          <n-input placeholder="请输入发送消息原文(加密)" v-model:value="params.sendMsg" />
          </n-form-item>

          <n-form-item label="发送消息译文(加密)" path="translatedMsg">
          <n-input placeholder="请输入发送消息译文(加密)" v-model:value="params.translatedMsg" />
          </n-form-item>

          <n-form-item label="消息类型" path="msgType">
            <n-select v-model:value="params.msgType" :options="options.msg_type" />
          </n-form-item>

          <n-form-item label="发送时间" path="sendTime">
            <DatePicker v-model:formValue="params.sendTime" type="datetime" />
          </n-form-item>

          <n-form-item label="是否已读" path="read">
            <n-select v-model:value="params.read" :options="options.read_status" />
          </n-form-item>
          <n-form-item label="发送状态" path="sendStatus">
            <n-select v-model:value="params.sendStatus" :options="options.send_status" />
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
  import { Edit, View } from '@/api/whats/whatsMsg';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
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
