<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="消息记录详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>聊天发起人</template>
          {{ formValue.initiator }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>发送人</template>
          {{ formValue.sender }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>接收人</template>
          {{ formValue.receiver }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>请求id</template>
          {{ formValue.reqId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>发送消息原文(加密)</template>
          {{ formValue.sendMsg }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>发送消息译文(加密)</template>
          {{ formValue.translatedMsg }}
        </n-descriptions-item>

        <n-descriptions-item label="消息类型">
          <n-tag
            :type="getOptionTag(options.msg_type, formValue?.msgType)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.msg_type, formValue?.msgType) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>发送时间</template>
          {{ formValue.sendTime }}
        </n-descriptions-item>

        <n-descriptions-item label="是否已读">
          <n-tag
            :type="getOptionTag(options.read_status, formValue?.read)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.read_status, formValue?.read) }}</n-tag
          >
        </n-descriptions-item>
        <n-descriptions-item label="发送状态">
          <n-tag
            :type="getOptionTag(options.send_status, formValue?.sendStatus)"
            size="small"
            class="min-left-space"
          >{{ getOptionLabel(options.send_status, formValue?.sendStatus) }}</n-tag
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
  import { View } from '@/api/whats/whatsMsg';
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
