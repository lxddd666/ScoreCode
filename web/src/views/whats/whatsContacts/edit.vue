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
          :label-width="100"
          class="py-4"
        >
          <n-form-item label="联系人姓名" path="name">
            <n-input type="textarea" placeholder="联系人姓名" v-model:value="params.name" />
          </n-form-item>

          <n-form-item label="联系人电话" path="phone">
            <n-input placeholder="请输入联系人电话" v-model:value="params.phone" />
          </n-form-item>
          <n-form-item label="头像" path="avatar">
            <FileChooser v-model:value="params.avatar" file-type="image" />
          </n-form-item>

          <n-form-item label="联系人邮箱" path="email">
            <n-input type="textarea" placeholder="联系人邮箱" v-model:value="params.email" />
          </n-form-item>

          <n-form-item label="联系人地址" path="address">
            <n-input type="textarea" placeholder="联系人地址" v-model:value="params.address" />
          </n-form-item>

          <n-form-item label="部门id" path="deptId">
            <!--            <n-input-number placeholder="请输入部门id" v-model:value="params.deptId" />-->
            <n-tree-select
              key-field="id"
              :options="deptList"
              :default-value="params.deptId"
              :default-expand-all="true"
              placeholder="选择部门"
              @update:value="handleUpdateDeptValue"
            />
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
  import { Edit, View } from '@/api/whats/whatsContacts';
  import { rules, options, State, newState } from './model';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';
  import {loadOptions} from "@/views/org/user/model";
  import {getDeptOption} from "@/api/org/dept";
  import FileChooser from "@/components/FileChooser/index.vue";

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
  const deptList = ref([]);

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

  function handleUpdateDeptValue(value) {
    params.value.deptId = Number(value);
  }
  async function loadDept() {
    const dept = await getDeptOption();
    if (dept.list) {
      deptList.value = dept.list;
    }
  }

  watch(
    () => props.formParams,
    (value) => {
      loadForm(value);
    }
  );
  onMounted(async () => {
    await loadDept();
  });
</script>

<style lang="less"></style>
