<template>
  <n-modal
      v-model:show="isShowModal"
      :mask-closable="false"
      preset="dialog"
      title="绑定帐号"
      content=""
      positive-text="确认"
      negative-text="取消"
      @positive-click="onPositiveClick"
      @negative-click="onNegativeClick"
      :style="{
          width: dialogWidth,
        }"
  />
</template>

<script lang="ts" setup>
import { onMounted, ref, computed, watch } from 'vue';
import {useMessage} from 'naive-ui'
import {newState, State} from "@/views/whats/whatsProxy/model";
import {adaModalWidth} from "@/utils/hotgo";

const emit = defineEmits([ 'updateBindShowModal']);

interface Props {
  showModal: boolean;
  formParams?: State;
}

const message = useMessage()
const dialogWidth = ref('75%');


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
    emit('updateBindShowModal', value);
  },
});

function onNegativeClick(e) {
  message.success('Cancel')
}

function onPositiveClick(e) {
  message.success('Submit')
}


onMounted(async () => {
  adaModalWidth(dialogWidth);
});

function closeForm() {
  isShowModal.value = false;
}

function loadForm(value) {
  // 新增
}

watch(
    () => props.formParams,
    (value) => {
      loadForm(value);
    }
);
</script>

<style scoped lang="less">

</style>
