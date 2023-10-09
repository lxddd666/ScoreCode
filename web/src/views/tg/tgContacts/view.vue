<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="联系人管理详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>tg id</template>
          {{ formValue.tgId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>username</template>
          {{ formValue.username }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>First Name</template>
          {{ formValue.firstName }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>Last Name</template>
          {{ formValue.lastName }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>phone</template>
          {{ formValue.phone }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>photo</template>
          <n-image style="margin-left: 10px; height: 100px; width: 100px" :src="formValue.photo"
        /></n-descriptions-item>

        <n-descriptions-item>
          <template #label>organization id</template>
          {{ formValue.orgId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>comment</template>
          <span v-html="formValue.comment"></span></n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/tg/tgContacts';
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