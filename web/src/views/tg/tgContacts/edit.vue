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
          <n-form-item label="tg id" path="tgId">
            <n-input-number placeholder="请输入tg id" v-model:value="params.tgId" />
          </n-form-item>

          <n-form-item label="username" path="username">
          <n-input placeholder="请输入username" v-model:value="params.username" />
          </n-form-item>

          <n-form-item label="First Name" path="firstName">
          <n-input placeholder="请输入First Name" v-model:value="params.firstName" />
          </n-form-item>

          <n-form-item label="Last Name" path="lastName">
          <n-input placeholder="请输入Last Name" v-model:value="params.lastName" />
          </n-form-item>

          <n-form-item label="phone" path="phone">
          <n-input placeholder="请输入phone" v-model:value="params.phone" />
          </n-form-item>

          <n-form-item label="photo" path="photo">
            <UploadImage :maxNumber="1" v-model:value="params.photo" />
          </n-form-item>

          <n-form-item label="type" path="type">
            <n-radio-group v-model:value="params.type" name="type">
            <n-radio-button
              v-for="type in options.contacts_type"
              :key="type.value"
              :value="type.value"
              :label="type.label"
            />
          </n-radio-group>
          </n-form-item>

          <n-form-item label="organization id" path="orgId">
            <n-input-number placeholder="请输入organization id" v-model:value="params.orgId" />
          </n-form-item>

          <n-form-item label="comment" path="comment">
            <n-input type="textarea" placeholder="comment" v-model:value="params.comment" />
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
  import { Edit, View } from '@/api/tg/tgContacts';
  import UploadImage from '@/components/Upload/uploadImage.vue';
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