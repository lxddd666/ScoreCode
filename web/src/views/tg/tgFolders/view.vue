<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="tg分组详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>组织ID</template>
          {{ formValue.orgId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>用户ID</template>
          {{ formValue.memberId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>分组名称</template>
          {{ formValue.folderName }}
        </n-descriptions-item>
        <n-descriptions-item>
          <template #label>分组人数</template>
          {{ formValue.memberCount }}
        </n-descriptions-item>
        <n-descriptions-item>
          <template #label>备注</template>
          <span v-html="formValue.comment"></span>
        </n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/tg/tgFolders';
  import { newState } from './model';

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
