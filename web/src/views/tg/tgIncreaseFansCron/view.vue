<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="TG频道涨粉任务详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>组织ID</template>
          {{ formValue.orgId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>发起任务的用户ID</template>
          {{ formValue.memberId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>频道地址</template>
          {{ formValue.channel }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>持续天数</template>
          {{ formValue.dayCount }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>涨粉数量</template>
          {{ formValue.fansCount }}
        </n-descriptions-item>

        <n-descriptions-item label="任务状态">
          <n-tag
            :type="getOptionTag(options.sys_normal_disable, formValue?.cronStatus)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.sys_normal_disable, formValue?.cronStatus) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>备注</template>
          <span v-html="formValue.comment"></span></n-descriptions-item>

        <n-descriptions-item>
          <template #label>已执行天数</template>
          {{ formValue.executedDays }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>已添加粉丝数</template>
          {{ formValue.increasedFans }}
        </n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/org/tgIncreaseFansCron';
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
