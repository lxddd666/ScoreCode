<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="TG频道涨粉任务执行情况详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>任务ID</template>
          {{ formValue.cronId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>加入频道的userId</template>
          {{ formValue.tgUserId }}
        </n-descriptions-item>

        <n-descriptions-item label="加入状态：0失败，1成功，2完成">
          <n-tag
            :type="getOptionTag(options.sys_normal_disable, formValue?.joinStatus)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.sys_normal_disable, formValue?.joinStatus) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>备注</template>
          <span v-html="formValue.comment"></span></n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/org/tgIncreaseFansCronAction';
  import { newState, options } from './model';
  import { getOptionLabel, getOptionTag } from '@/utils/hotgo';
  import { getFileExt } from '@/utils/urlUtils';

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
      message.error('id不正确，请检查！');
      return;
    }
    formValue.value = await View({ id: id });
  });
</script>

<style lang="less" scoped></style>
