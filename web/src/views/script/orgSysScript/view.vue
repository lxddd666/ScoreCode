<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="公司话术管理详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>分组</template>
          {{ formValue.groupId }}
        </n-descriptions-item>

        <n-descriptions-item label="话术分类">
          <n-tag
            :type="getOptionTag(options.script_class, formValue?.scriptClass)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.script_class, formValue?.scriptClass) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>快捷指令</template>
          {{ formValue.short }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>话术内容</template>
          <span v-html="formValue.content"></span></n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref} from 'vue';
import {useRouter} from 'vue-router';
import {useMessage} from 'naive-ui';
import {View} from '@/api/script/orgSysScript';
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
    formValue.value = await View({ id: id });
  });
</script>

<style lang="less" scoped></style>
