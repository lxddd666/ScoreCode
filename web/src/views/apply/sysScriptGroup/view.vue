<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="话术分组详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>组织ID</template>
          {{ formValue.orgId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>部门ID</template>
          {{ formValue.deptId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>用户ID</template>
          {{ formValue.memberId }}
        </n-descriptions-item>

        <n-descriptions-item label="分组类型">
          <n-tag
            :type="getOptionTag(options.script_type, formValue?.type)"
            size="small"
            class="min-left-space"
          >{{ getOptionLabel(options.script_type, formValue?.type) }}
          </n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>自定义组名</template>
          {{ formValue.name }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>话术数量</template>
          {{ formValue.scriptCount }}
        </n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref} from 'vue';
import {useRouter} from 'vue-router';
import {useMessage} from 'naive-ui';
import {View} from '@/api/apply/sysScriptGroup';
import {newState, options} from './model';
import {getOptionLabel, getOptionTag} from '@/utils/hotgo';

const message = useMessage();
const router = useRouter();
const id = Number(router.currentRoute.value.params.id);
const formValue = ref(newState(null));
const fileAvatarCSS = computed(() => {
  return {
    '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
    '--n-font-size': `18px`,
  };
});

//下载
function download(url: string) {
  window.open(url);
}

onMounted(async () => {
  if (id < 1) {
    message.error('ID不正确，请检查！');
    return;
  }
  formValue.value = await View({id: id});
});
</script>

<style lang="less" scoped></style>
