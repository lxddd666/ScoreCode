<template>
  <n-form ref="formRef" label-placement="left" size="large" :model="formInline" :rules="rules">
    <n-form-item path="username">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.username"
        placeholder="请输入用户名"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <PersonOutline />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>
    <n-form-item path="pass">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.pass"
        type="password"
        placeholder="请输入密码"
        show-password-on="click"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <LockClosedOutline />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>

    <n-form-item path="confirmPwd">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.confirmPwd"
        type="password"
        placeholder="再次输入密码"
        show-password-on="click"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <LockClosedOutline />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>

    <n-form-item path="mobile">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.mobile"
        placeholder="请输入手机号码"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <MobileOutlined />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>
    <n-form-item path="firstName">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.firstName"
        placeholder="请输入FirstName"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <MobileOutlined />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>
    <n-form-item path="lastName">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.lastName"
        placeholder="请输入LastName"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <MobileOutlined />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>
    <n-form-item path="email">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.email"
        placeholder="请输入邮箱"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <MobileOutlined />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>

    <n-form-item path="code">
      <n-input-group>
        <n-input
          @keyup.enter="handleSubmit"
          v-model:value="formInline.code"
          placeholder="请输入验证码"
        >
          <template #prefix>
            <n-icon size="18" color="#808695" :component="SafetyCertificateOutlined" />
          </template>
        </n-input>
        <n-button
          type="primary"
          ghost
          @click="sendMobileCode"
          :disabled="isCounting"
          :loading="sendLoading"
        >
          {{ sendLabel }}
        </n-button>
      </n-input-group>
    </n-form-item>

    <n-form-item path="inviteCode">
      <n-input
        :style="{ width: '100%' }"
        placeholder="邀请码(选填)"
        @keyup.enter="handleSubmit"
        v-model:value="formInline.inviteCode"
        :disabled="inviteCodeDisabled"
      >
        <template #prefix>
          <n-icon size="18" color="#808695" :component="TagOutlined" />
        </template>
      </n-input>
    </n-form-item>
    <n-form-item path="companyName">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.companyName"
        placeholder="请输入公司名称"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <PersonOutline />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>
    <n-form-item path="companyCode">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.companyCode"
        placeholder="请输入公司编码"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <PersonOutline />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>
    <n-form-item path="companyPhone">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.companyPhone"
        placeholder="请输入公司联系电话"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <PersonOutline />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>
    <n-form-item path="leader">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.leader"
        placeholder="请输入公司负责人"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <PersonOutline />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>
    <n-form-item path="email">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.companyEmail"
        placeholder="请输入公司邮箱"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <PersonOutline />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>
    <n-form-item class="default-color">
      <Agreement
        v-model:value="agreement"
        @clickProtocol="handleClickProtocol"
        @clickPolicy="handleClickPolicy"
      />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="handleSubmit" size="large" :loading="loading" block>
        注册
      </n-button>
    </n-form-item>

    <FormOther moduleKey="login" tag="登录账号" @updateActiveModule="updateActiveModule" />
  </n-form>

  <n-modal
    v-model:show="showModal"
    :show-icon="false"
    :mask-closable="false"
    preset="dialog"
    :closable="false"
    :style="{
      width: dialogWidth,
      position: 'top',
      bottom: '15vw',
    }"
  >
    <n-space justify="center">
      <div class="agree-title">《{{ agreeTitle }}》</div>
    </n-space>

    <div v-html="modalContent"></div>

    <n-divider />
    <n-space justify="center">
      <n-button type="info" ghost strong @click="handleAgreement(true)">我已知晓并接受</n-button>
      <n-button type="error" ghost strong @click="handleAgreement(false)">我拒绝</n-button>
    </n-space>
  </n-modal>
</template>

<script lang="ts" setup>
  import '../components/style.less';
  import { onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { ResultEnum } from '@/enums/httpEnum';
  import { LockClosedOutline, PersonOutline } from '@vicons/ionicons5';
  import { MobileOutlined, SafetyCertificateOutlined, TagOutlined } from '@vicons/antd';
  import { aesEcb } from '@/utils/encrypt';
  import FormOther from '../components/form-other.vue';
  import { useSendCode } from '@/hooks/common';
  import { validate } from '@/utils/validateUtil';
  import { register, SendEms, SendSms } from '@/api/system/user';
  import { useUserStore } from '@/store/modules/user';
  import { adaModalWidth } from '@/utils/hotgo';
  import {any} from "vue-types";

  interface FormState {
    username: string;
    pass: string;
    confirmPwd: string;
    mobile: string;
    email: string;
    code: string;
    inviteCode: string;
    password: string;
    firstName: string;
    lastName: string;
    companyName: string;
    companyCode: string;
    companyPhone: string;
    leader: string;
    companyEmail: string;
  }

  const formRef = ref();
  const router = useRouter();
  const message = useMessage();
  const userStore = useUserStore();
  const loading = ref(false);
  const showModal = ref(false);
  const agreeTitle = ref('');
  const modalContent = ref('');
  const { sendLabel, isCounting, loading: sendLoading, activateSend } = useSendCode();
  const agreement = ref(true);
  const inviteCodeDisabled = ref(false);
  const dialogWidth = ref('85%');
  const emit = defineEmits(['updateActiveModule']);

  const formInline = ref<FormState>({
    username: '',
    pass: '',
    confirmPwd: '',
    mobile: '',
    email: '',
    code: '',
    inviteCode: '',
    password: '',
    firstName: '',
    lastName: '',
    companyName: '',
    companyCode: '',
    companyPhone: '',
    leader: '',
    companyEmail: '',
  });

  const rules = {
    username: { required: true, message: '请输入用户名', trigger: 'blur' },
    pass: { required: true, message: '请输入密码', trigger: 'blur' },
    mobile: { required: true, message: '请输入手机号码', trigger: 'blur' },
    email: { required: true, message: '请输入邮箱', trigger: 'blur' },
    code: { required: true, message: '请输入验证码', trigger: 'blur' },
  };

  const handleSubmit = (e) => {
    debugger;
    e.preventDefault();
    formRef.value.validate(async (errors) => {
      if (!errors) {
        if (formInline.value.pass !== formInline.value.confirmPwd) {
          message.info('两次输入的密码不一致，请检查');
          return;
        }

        if (!agreement.value) {
          message.info('请确认你已经仔细阅读并接受《用户协议》和《隐私权政策》并已勾选接受选项');
          return;
        }

        message.loading('注册中...');
        loading.value = true;

        try {
          var jsonMsg = {
            username: formInline.value.username,
            password: aesEcb.encrypt(formInline.value.pass),
            mobile: formInline.value.mobile,
            email: formInline.value.email,
            code: formInline.value.code,
            inviteCode: formInline.value.inviteCode,
            firstName: formInline.value.firstName,
            lastName: formInline.value.lastName,
          };
          if (formInline.value.inviteCode == null || formInline.value.inviteCode == '') {
            var orgDetail = {
              name: formInline.value.companyName,
              code: formInline.value.companyCode,
              leader: formInline.value.leader,
              email: formInline.value.companyEmail,
              phone: formInline.value.companyPhone,
            };
            jsonMsg['orgInfo'] = orgDetail;
          }
          const { code, message: msg } = await register(jsonMsg);
          message.destroyAll();
          if (code == ResultEnum.SUCCESS) {
            message.success('注册成功，请登录！');
            updateActiveModule('login');
          } else {
            message.info(msg || '注册失败');
          }
        } finally {
          loading.value = false;
        }
      } else {
        message.error('请填写完整信息，并且进行验证码校验');
      }
    });
  };

  onMounted(() => {
    const inviteCode = router.currentRoute.value.query?.inviteCode as string;
    if (inviteCode) {
      inviteCodeDisabled.value = true;
      formInline.value.inviteCode = inviteCode;
    }

    adaModalWidth(dialogWidth);
  });

  function updateActiveModule(key: string) {
    emit('updateActiveModule', key);
  }

  function sendMobileCode() {
    if (formInline.value.email !== '') {
      validate.email(rules.email, formInline.value.email, function (error?: Error) {
        if (error === undefined) {
          activateSend(SendEms({ email: formInline.value.email, event: 'register' }),undefined);
          return;
        }
        message.error(error.message);
      });
    } else {
      validate.phone(rules.mobile, formInline.value.mobile, function (error?: Error) {
        if (error === undefined) {
          activateSend(SendSms({ mobile: formInline.value.mobile, event: 'register' }),undefined);
          return;
        }
        message.error(error.message);
      });
    }
  }

  function handleClickProtocol() {
    showModal.value = true;
    agreeTitle.value = '用户协议';
    modalContent.value = userStore.loginConfig?.loginProtocol as string;
  }

  function handleClickPolicy() {
    showModal.value = true;
    agreeTitle.value = '隐私权政策';
    modalContent.value = userStore.loginConfig?.loginPolicy as string;
  }

  function handleAgreement(agree: boolean) {
    showModal.value = false;
    agreement.value = agree;
  }
</script>

<style lang="less" scoped>
  .agree-title {
    font-size: 18px;
    margin-bottom: 22px;
  }
</style>
