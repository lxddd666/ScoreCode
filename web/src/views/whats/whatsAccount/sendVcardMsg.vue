<!--suppress ALL -->
<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
        title="发送名片"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-form
          :model="params"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="120"
          class="py-4"
        >
          <n-form-item label="发送人" path="sender">
            <n-input placeholder="请输入发送人" v-model:value="params.sender" disabled/>
          </n-form-item>

          <n-form-item label="接收人" path="receiver">
            <n-input placeholder="请输入接收人" v-model:value="params.receiver"/>
          </n-form-item>


          <n-form-item label="版本" path="version">
            <n-input type="version" placeholder="请输入版本" v-model:value="params.version"/>
          </n-form-item>

          <n-form-item label="生成名片的软件" path="prodid">
            <n-input type="prodid" placeholder="请输入生成名片的软件" v-model:value="params.prodid"/>
          </n-form-item>

          <n-form-item label="名字" path="fn">
            <n-input type="fn" placeholder="请输入名字" v-model:value="params.fn"/>
          </n-form-item>

          <n-form-item label="手机" path="tel">
            <n-input type="tel" placeholder="请输入手机" v-model:value="params.tel"/>
          </n-form-item>

          <n-form-item label="工作单位" path="org">
            <n-input type="org" placeholder="请输入工作单位" v-model:value="params.org"/>
          </n-form-item>

          <n-form-item label="自定义名字(一般和名字一样)" path="xwabizname">
            <n-input type="xwabizname" placeholder="请输入自定义名字" v-model:value="params.xwabizname"/>
          </n-form-item>

          <n-form-item label="名片结束部分" path="end">
            <n-input type="end" placeholder="请输入名片结束部分" v-model:value="params.end"/>
          </n-form-item>

          <n-form-item label="展示名字(一般和名字一样)" path="displayname">
            <n-input type="displayname" placeholder="请输入展示名字" v-model:value="params.displayname"/>
          </n-form-item>

          <n-form-item label="家庭" path="family">
            <n-input type="family" placeholder="请输入家庭" v-model:value="params.family"/>
          </n-form-item>

          <n-form-item label="名字前缀(例如Mr.或Dr.)" path="prefixes">
            <n-input type="prefixes" placeholder="请输入前缀" v-model:value="params.prefixes"/>
          </n-form-item>

          <n-form-item label="语言" path="language">
            <n-input type="language" placeholder="请输语言" v-model:value="params.language"/>
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
import {SendMsg, SendVcardMsg} from '@/api/whats/whatsAccount';
import {useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';

const emit = defineEmits(['reloadTable', 'sendMsgShowModal']);

interface Props {
  showModal: boolean;
  sender?: string;
}

interface MsgReq {
  sender: string;
  receiver: string;
  version: string;
  prodid: string;
  fn: string;
  org: string;
  tel: string;
  xwabizname: string;
  end: string;
  displayname: string;
  family: string;
  given: string;
  prefixes: string;
  language: string;
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
  textMsg: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入发送内容',
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
      let vacrd = {
        'version': params.value.version,
        'prodid': params.value.prodid,
        'fn': params.value.fn,
        'org': params.value.org,
        'tel': params.value.tel,
        'xwabizname': params.value.xwabizname,
        'end': params.value.end,
        'family': params.value.family,
        'given': params.value.given,
        'prefixes': params.value.prefixes,
        'language': params.value.language,
      }
      let req = {
        'sender': params.value.sender,
        'receiver': params.value.receiver,
        'vcard': vacrd,
      }
      SendVcardMsg(req).then((_res) => {
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
