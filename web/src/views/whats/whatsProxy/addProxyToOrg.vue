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
          <n-form-item label="所属公司(组织)" path="orgId">
            <n-select
              key-field="id"
              :options="options.org"
              v-model:value="params.orgId"
              @update:value="handleUpdateOrgValue"
            />
          </n-form-item>
          <n-form-item label="代理地址" path="address">
            <n-input placeholder="请输入代理地址" v-model:value="params.address" />
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
  import { AddProxyOrgReq, View } from '@/api/whats/whatsProxy';
  import { rules, options, State, newState } from './model';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';
  const emit = defineEmits(['reloadTable', 'updateShowModal', 'updateAddProxyToOrg']);

  interface Props {
    showModal: boolean;
    formParams?: State;
    address: [];
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
      emit('updateAddProxyToOrg', value);
    },
  });



  const loading = ref(false);
  const params = ref<State>(props.formParams);
  const message = useMessage();
  const formRef = ref<any>({});
  const dialogWidth = ref('75%');
  const formBtnLoading = ref(false);

  function confirmForm(e) {
    debugger;
    e.preventDefault();
    formBtnLoading.value = true;
    var postParam = {
      orgId: params.value.orgId,
      ProxyAddresses: [params.value.address],
    };
    formRef.value.validate((errors) => {
      if (!errors) {
        AddProxyOrgReq(postParam).then((_res) => {
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

  function handleUpdateOrgValue(value) {
    params.value.orgId = Number(value);
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
