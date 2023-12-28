<!--suppress ALL -->
<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
       :title="params?.id > 0 ? '编辑用户 #' + params?.id : '添加用户'"
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
          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="姓名" path="realName">
                <n-input placeholder="请输入姓名" v-model:value="params.realName" />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="用户名" path="username">
                <n-input placeholder="请输入登录用户名" v-model:value="params.username" />
              </n-form-item>
            </n-gi>
          </n-grid>

          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="firstName" path="firstName">
                <n-input placeholder="请输入firstName" v-model:value="params.firstName" />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="lastName" path="lastName">
                <n-input placeholder="请输入登录用户名" v-model:value="params.lastName" />
              </n-form-item>
            </n-gi>
          </n-grid>

          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="所属公司(组织)" path="orgId">
                <n-select
                  key-field="id"
                  :options="options.org"
                  v-model:value="params.orgId"
                  @update:value="handleUpdateOrgValue"
                />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="绑定角色" path="roleId">
                <n-tree-select
                  key-field="id"
                  :options="options.role"
                  v-model:value="params.roleId"
                  :default-expand-all="true"
                  @update:value="handleUpdateRoleValue"
                />
              </n-form-item>
            </n-gi>
          </n-grid>
          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="密码" path="password">
                <n-input
                  type="password"
                  :placeholder="params.id === 0 ? '请输入' : '不填则不修改'"
                  v-model:value="params.password"
                />
              </n-form-item>
            </n-gi>
          </n-grid>
          <n-divider title-placement="left">填写更多信息</n-divider>
          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="手机号" path="mobile">
                <n-input placeholder="请输入" v-model:value="params.mobile" />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="邮箱" path="email">
                <n-input placeholder="请输入" v-model:value="params.email" />
              </n-form-item>
            </n-gi>
          </n-grid>

          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="性别" path="sex">
                <n-radio-group v-model:value="params.sex" name="sex">
                  <n-radio-button
                    v-for="status in sexOptions"
                    :key="status.value"
                    :value="status.value"
                    :label="status.label"
                  />
                </n-radio-group>
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="状态" path="status">
                <n-radio-group v-model:value="params.status" name="status">
                  <n-radio-button
                    v-for="status in statusOptions"
                    :key="status.value"
                    :value="status.value"
                    :label="status.label"
                  />
                </n-radio-group>
              </n-form-item>
            </n-gi>
          </n-grid>

          <n-form-item label="备注" path="remark">
            <n-input type="textarea" placeholder="请输入备注" v-model:value="params.remark" />
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
  import { Edit,GetMemberView } from '@/api/org/user';
  import {
    rules,
    options,
    State,
    defaultState,
    addNewState,
    addState
  } from './model';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';
  import {cloneDeep} from "lodash-es";
  import {sexOptions, statusOptions} from "@/enums/optionsiEnum";
  import {getOrgOption} from "@/api/org/dept";

  const emit = defineEmits(['reloadTable', 'updateShowModal']);

  interface Props {
    showModal: boolean;
    formParams?: State;
  }

  const props = withDefaults(defineProps<Props>(), {
    showModal: false,
    formParams: () => {
      return defaultState;
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
  const placeholderSelect = ref('请选择公司(组织)');
  const showOptionsSelect = ref(true);

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
      loading.value = false;
    });
  }

  onMounted(async () => {
    adaModalWidth(dialogWidth);
  });

  function closeForm() {
    loading.value = false;
    emit('updateShowModal', false);
  }

  function loadForm(value) {
    // 新增
    if (value.id < 1) {
      params.value = cloneDeep(value);
      loading.value = false;
      return;
    }

    // 编辑
    GetMemberView({ id: value.id })
      .then((res) => {
        params.value = res;
      });
    loading.value = false;
  }

  function handleUpdateDeptValue(value) {
    params.value.deptId = Number(value);
  }
  function handleUpdateOrgValue(value) {
    if (value) {
      showOptionsSelect.value = false;
      placeholderSelect.value = '请选择';
    }
    params.value.orgId = Number(value);
  }

  function handleUpdateRoleValue(value) {
    params.value.roleId = Number(value);
  }

  function handleUpdatePostValue(value) {
    params.value.postIds = value;
  }

  watch(
    () => props.formParams,
    (value) => {
      loadForm(value);
    }
  );
</script>

<style lang="less"></style>
