<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="TG账号详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>用户名</template>
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
          <template #label>手机号</template>
          {{ formValue.phone }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>账号头像</template>
          <span v-html="formValue.photo"></span></n-descriptions-item>

        <n-descriptions-item label="账号状态">
          <n-tag
            :type="getOptionTag(options.account_status, formValue?.accountStatus)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.account_status, formValue?.accountStatus) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>是否在线</template>
          {{ formValue.isOnline }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>代理地址</template>
          {{ formValue.proxyAddress }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>上次登录时间</template>
          {{ formValue.lastLoginTime }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>备注</template>
          <span v-html="formValue.comment"></span></n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref} from 'vue';
import {useRouter} from 'vue-router';
import {useMessage} from 'naive-ui';
import {View} from '@/api/tg/tgUser';
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
      message.error('id不正确，请检查！');
      return;
    }
    formValue.value = await View({ id: id });
  });
</script>

<style lang="less" scoped></style>
