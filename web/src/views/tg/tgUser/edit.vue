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
          <n-form-item label="用户名" path="username">
            <n-input placeholder="请输入用户名" v-model:value="params.username"/>
          </n-form-item>

          <n-form-item label="First Name" path="firstName">
            <n-input placeholder="请输入First Name" v-model:value="params.firstName"/>
          </n-form-item>

          <n-form-item label="Last Name" path="lastName">
            <n-input placeholder="请输入Last Name" v-model:value="params.lastName"/>
          </n-form-item>

          <n-form-item label="手机号" path="phone">
            <n-input placeholder="请输入手机号" v-model:value="params.phone" disabled/>
          </n-form-item>

          <n-form-item label="账号头像" path="photo">
            <n-avatar
              round
              :size="54"
              color="transparent"
              :src="getPhoto(params.phone,params.tgId,params.photo)"
            ></n-avatar>
          </n-form-item>
          <n-form-item label="签名" path="bio">
            <n-input type="textarea" placeholder="签名" v-model:value="params.bio"/>
          </n-form-item>

          <n-form-item label="备注" path="comment">
            <n-input type="textarea" placeholder="备注" v-model:value="params.comment"/>
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
import {Edit, View} from '@/api/tg/tgUser';
import {newState, rules, State} from './model';
import {useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';
import {useGlobSetting} from "@/hooks/setting";
import {storage} from "@/utils/Storage";
import {ACCESS_TOKEN} from "@/store/mutation-types";
import {getPhoto} from "@/utils/tgUtils";

const emit = defineEmits(['reloadTable', 'updateShowModal']);

const globSetting = useGlobSetting();
const tgPrefix = globSetting.tgPrefix || '';

const token = storage.get(ACCESS_TOKEN);

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
  View({id: value.id})
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
    console.log(111)
    loadForm(value);
  }
);
</script>

<style lang="less"></style>
