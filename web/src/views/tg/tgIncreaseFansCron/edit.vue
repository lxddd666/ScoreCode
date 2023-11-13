<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
        :title="params?.id > 0 ? '编辑 #' + params?.id : '添加'"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-form
          :model="params"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="200"
          class="py-4"
        >
          <n-form-item label="频道地址" path="channel">
            <n-input
              placeholder="请输入有效的频道地址，例如 https://t.me/xxxx"
              v-model:value="params.channel"
            />
            <n-space>
              <n-button type="info" :loading="formBtnLoading" @click="confirmChannelForm">
                校验频道
              </n-button>
            </n-space>
          </n-form-item>
          <n-form-item label="任务名称" path="taskName">
            <n-input placeholder="唯一的任务名称" v-model:value="params.taskName" />
          </n-form-item>
          <n-form-item label="频道当前粉丝数" path="channelInitFansCount">
            <n-input-number
              :disabled="true"
              placeholder="频道当前粉丝数"
              v-model:value="params.channelMemberCount"
            />
          </n-form-item>
          <n-form-item label="持续天数" path="dayCount">
            <n-input-number placeholder="请输入持续天数" v-model:value="params.dayCount" />
            推荐天数-><n-input-number
              :disabled="true"
              placeholder="推荐天数"
              v-model:value="params.recommendedDays"
            />
          </n-form-item>

          <n-form-item label="涨粉数量" path="fansCount">
            <n-input-number placeholder="请输入涨粉数量" v-model:value="params.fansCount" />
            <n-space>
              <n-button type="info" :loading="formBtnLoading" @click="ChannelDailyIncrease">
                增长数量计算
              </n-button>
            </n-space>
          </n-form-item>
        </n-form>
        <template #action>
          <n-space>
            <n-button @click="closeForm">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref, computed, watch } from 'vue';
  import { CheckChannel, DailyIncrease, Edit, View } from '@/api/org/tgIncreaseFansCron';
  import { rules, options, State, newState } from './model';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable', 'updateShowModal']);

  interface Props {
    showModal: boolean;
    formParams?: State;
  }

  const props = withDefaults(defineProps<Props>(), {
    showModal: false,
    formParams: () => {
      return newState(null);
    },
  });

  const isShowModal = computed({
    get: () => {
      return props.showModal;
    },
    set: (value) => {
      emit('updateShowModal', value);
    },
  });

  const loading = ref(false);
  const params = ref<State>(props.formParams);
  const message = useMessage();
  const formRef = ref<any>({});
  const dialogWidth = ref('75%');
  const formBtnLoading = ref(false);

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(params.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            isShowModal.value = false;
            emit('reloadTable');
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }
  function ChannelDailyIncrease(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        DailyIncrease(params.value).then((_res) => {
          debugger;
          params.value.recommendedDays = _res.totalDay;
          var str = '';
          _res.dailyIncreaseFan.forEach(function (fanIncrease, index) {
            str += '第' + (index + 1) + '天：' + fanIncrease + '粉丝\n';
          });
          console.info(str);
          message.success('操作成功' + '每日增长数量约为' + str);
          setTimeout(() => {
            isShowModal.value = true;
            emit('reloadTable');
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function confirmChannelForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        CheckChannel(params.value).then((_res) => {
          params.value.channelMemberCount = _res.channelMsg.channelMemberCount;
          debugger;
          message.success(
            '频道有效,频道Title:' +
              _res.channelMsg.channelTitle +
              ',频道人数:' +
              _res.channelMsg.channelMemberCount
          );
          setTimeout(() => {
            isShowModal.value = true;
            emit('reloadTable');
          });
        });
      } else {
        message.error('请填写有效的频道地址');
      }
      formBtnLoading.value = false;
    });
  }

  onMounted(async () => {
    adaModalWidth(dialogWidth);
  });

  function closeForm() {
    isShowModal.value = false;
  }

  function loadForm(value) {
    // 新增
    if (value.id < 1) {
      params.value = newState(value);
      loading.value = false;
      return;
    }

    loading.value = true;
    // 编辑
    View({ id: value.id })
      .then((res) => {
        params.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  watch(
    () => props.formParams,
    (value) => {
      loadForm(value);
    }
  );
</script>

<style lang="less"></style>