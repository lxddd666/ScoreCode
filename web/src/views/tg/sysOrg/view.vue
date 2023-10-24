<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="客户公司详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>公司名称</template>
          {{ formValue.name }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>公司编码</template>
          <span v-html="formValue.code"></span></n-descriptions-item>

        <n-descriptions-item>
          <template #label>负责人</template>
          {{ formValue.leader }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>联系电话</template>
          {{ formValue.phone }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>邮箱</template>
          {{ formValue.email }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>排序</template>
          {{ formValue.sort }}
        </n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/tg/sysOrg';
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
      message.error('公司ID不正确，请检查！');
      return;
    }
    formValue.value = await View({ id: id });
  });
</script>

<style lang="less" scoped></style>