<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
          v-model:show="isShowModal"
          :show-icon="false"
          preset="dialog"
          title="验证码登录"
          :style="{
          width: dialogWidth,
        }"
      >
        <n-form
            :model="formInline"
            :rules="rules"
            ref="formRef"
            label-placement="left"
            :label-width="200"
            class="py-4"
        >


          <n-form-item label="手机号" path="phone">
            <n-input placeholder="请输入手机号" v-model:value="formInline.phone"/>
          </n-form-item>
          <n-form-item label="验证码" path="code">
            <n-input-group>
              <n-input
                  @keyup.enter="confirmForm"
                  v-model:value="formInline.code"
                  placeholder="请输入验证码"
              >
                <template #prefix>
                  <n-icon size="18" color="#808695" :component="SafetyCertificateOutlined"/>
                </template>
              </n-input>
              <n-button
                  type="primary"
                  ghost
                  @click="sendCode"
                  :disabled="isCounting"
                  :loading="sendLoading"
              >
                {{ sendLabel }}
              </n-button>
            </n-input-group>
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
import {computed, onMounted, ref} from 'vue';
import {TgCodeLogin, TgSendCode} from '@/api/tg/tgUser';
import {useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';
import {useGlobSetting} from "@/hooks/setting";
import {storage} from "@/utils/Storage";
import {ACCESS_TOKEN} from "@/store/mutation-types";
import {SafetyCertificateOutlined} from "@vicons/antd";
import {useSendCode} from "@/hooks/common";

const emit = defineEmits(['reloadTable', 'updateShowModal']);
const {sendLabel, isCounting, loading: sendLoading, activateSend} = useSendCode();
const globSetting = useGlobSetting();
const tgPrefix = globSetting.tgPrefix || '';

const token = storage.get(ACCESS_TOKEN);

const rules = {
  phone: {required: true, message: '请输入手机号码', trigger: 'blur'},
  code: {required: true, message: '请输入验证码', trigger: 'blur'},
};

interface FormState {
  phone: string;
  code: string;
  reqId: string;

}

const formInline = ref<FormState>({
  phone: '',
  code: '',
  reqId: '',
});

interface Props {
  showModal: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  showModal: false,
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
const message = useMessage();
const formRef = ref<any>({});
const dialogWidth = ref('75%');
const formBtnLoading = ref(false);

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors:Error) => {
    if (!errors) {
      TgCodeLogin(formInline.value).then((_res) => {
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

function sendCode() {
  if (formInline.value.phone === '') {
    message.error('请输入手机号');
  } else {
    activateSend(TgSendCode({phone: formInline.value.phone}), function (_res: any) {
      console.log(_res);
      formInline.value.reqId = _res.reqId;
    });

    return;
  }
}

</script>

<style lang="less"></style>
