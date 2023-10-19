<!--suppress ALL -->
<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
        title="迁移联系人"
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
          <n-form-item label="被修改账号" path="ModifiedAccount">
            <n-input placeholder=" 请输入被修改的账号" v-model:value="params.ModifiedAccount" />
          </n-form-item>
          <n-form-item label="修改账号" path="UpdateAccount">
            <n-input placeholder=" 请输入修改的账号" v-model:value="params.UpdateAccount"/>
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
import {MigrateContacts} from '@/api/whats/whatsAccount';
import {useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';

const emit = defineEmits(['reloadTable', 'migrateContactsShowModal']);

interface Props {
  showModal: boolean;
  ModifiedAccount?: string;
}

interface MsgReq {
   ModifiedAccount: string;
   UpdateAccount:string;
}

const props = withDefaults(defineProps<Props>(), {
  showModal: false,
  ModifiedAccount: () => {
    return '';
  },
});

const isShowModal = computed({
  get: () => {
    return props.showModal;
  },
  set: (value) => {
    emit('migrateContactsShowModal', value);
  },
});

const rules = {
  ModifiedAccount: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入账号',
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
        'ModifiedAccount':params.value.ModifiedAccount,
        'UpdateAccount': params.value.UpdateAccount,
      }
      MigrateContacts(req).then((_res) => {
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
  params.UpdateAccount = ''
}

function loadForm(value) {
  params.value.ModifiedAccount = value;
  loading.value = false;
}

watch(
  () => props.ModifiedAccount,
  (value) => {
    loadForm(value);
  }
);
</script>

<style lang="less"></style>
