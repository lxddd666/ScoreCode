<!--suppress ALL -->
<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
        title="发送消息"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-form
          :model="params"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="发送人" path="sender">
            <n-input placeholder="请输入发送人" v-model:value="params.sender" disabled/>
          </n-form-item>

          <n-form-item label="添加的联系人" path="receiver">
            <n-input placeholder="请输入联系人" v-model:value="params.contact"/>
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
import {SendMsg, SyncContact} from '@/api/whats/whatsAccount';
import {useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';

const emit = defineEmits(['reloadTable', 'sendMsgShowModal','syncContactShowModal','sendVcardMsgShowModal','sendSyncContactModel']);

interface Props {
  showModal: boolean;
  sender?: string;
}

interface MsgReq {
  sender: string;
  receiver: string;
  textMsg: string;
  contact: string;
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
    emit('sendSyncContactModel', value);
  },
});

const rules = {
  sender: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入发送人',
  },
};


const loading = ref(false);
const params = ref<MsgReq>({});
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
        'account': params.value.sender,
        'contacts': [params.value.contact]
      }
      SyncContact(req).then((_res) => {
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
  params.contact = ''
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
