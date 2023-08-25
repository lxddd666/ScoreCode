<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="代理管理详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>代理地址</template>
          {{ formValue.address }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>已连接数</template>
          {{ formValue.connectedCount }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>最大连接</template>
          {{ formValue.maxConnections }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>地区</template>
          <span v-html="formValue.region"></span></n-descriptions-item>

        <n-descriptions-item>
          <template #label>备注</template>
          <span v-html="formValue.comment"></span></n-descriptions-item>

        <n-descriptions-item label="状态">
          <n-tag
            :type="getOptionTag(options.sys_normal_disable, formValue?.status)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.sys_normal_disable, formValue?.status) }}</n-tag
          >
        </n-descriptions-item>


      </n-descriptions>
    </n-card>

    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }" title="关联账号">
      <AccountTable @reloadView="reloadView"></AccountTable>
    </n-card>

  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref} from 'vue';
import {useRouter} from 'vue-router';
import {useMessage} from 'naive-ui';
import {View} from '@/api/whats/whatsProxy';
import {newState, options} from './model';
import {getOptionLabel, getOptionTag} from '@/utils/hotgo';
import AccountTable from './account.vue'

const message = useMessage();
  const router = useRouter();
  const id = Number(router.currentRoute.value.query.id);
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

  async function reloadView() {
    formValue.value = await View({id: id});
  }
</script>

<style lang="less" scoped></style>
