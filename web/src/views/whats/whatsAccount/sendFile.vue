<!--suppress ALL -->
<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
        title="发送文件"
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
          <n-form-item label="发送人" path="sender">
            <n-input placeholder="请输入发送人" v-model:value="params.sender" disabled/>
          </n-form-item>

          <n-form-item label="接收人" path="receiver">
            <n-input placeholder="请输入接收人" v-model:value="params.receiver"/>
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
import {computed, onMounted, ref, watch} from 'vue';
import {SendFile} from '@/api/whats/whatsAccount';
import {useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';

const emit = defineEmits(['reloadTable', 'sendMsgShowModal', 'sendVcardMsgShowModal']);

interface Props {
  showModal: boolean;
  sender?: string;
}

interface MsgReq {
  sender: string;
  receiver: string;
}

const props = withDefaults(defineProps<Props>(), {
  showModal: false,
  sender: () => {
    return '';
  },
});

const isShowModal = computed({
  get: () => {
    return props.showModal;
  },
  set: (value) => {
    emit('sendMsgShowModal', value);
  },
});

const rules = {
  sender: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入发送人',
  },
  receiver: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入接收人',
  },
};


const loading = ref(false);
const params = ref<MsgReq>({
  sender: '',
  receiver: ''
});
const message = useMessage();
const formRef = ref<any>({});
const dialogWidth = ref('75%');
const formBtnLoading = ref(false);

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      let req = {
        'sender': params.value.sender,
        'receiver': params.value.receiver
      }
      SendFile(req).then((_res) => {
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
  params.value.receiver = '';
  params.value.textMsg = '';
}

function loadForm(value) {
  // 发送信息
  params.value.sender = value;
  loading.value = false;
}

watch(
  () => props.sender,
  (value) => {
    loadForm(value);
  }
);
</script>

<style lang="less"></style>
