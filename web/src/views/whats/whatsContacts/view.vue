<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="联系人管理详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>联系人姓名</template>
          <span v-html="formValue.name"></span></n-descriptions-item>

        <n-descriptions-item>
          <template #label>联系人电话</template>
          {{ formValue.phone }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>联系人头像</template>
          <span v-html="formValue.avatar"></span></n-descriptions-item>

        <n-descriptions-item>
          <template #label>联系人邮箱</template>
          <span v-html="formValue.email"></span></n-descriptions-item>

        <n-descriptions-item>
          <template #label>联系人地址</template>
          <span v-html="formValue.address"></span></n-descriptions-item>

        <n-descriptions-item>
          <template #label>组织id</template>
          {{ formValue.orgId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>部门id</template>
          {{ formValue.deptId }}
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
  import { View } from '@/api/whats/whatsContacts';
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