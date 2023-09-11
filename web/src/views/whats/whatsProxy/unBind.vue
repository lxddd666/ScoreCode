<template>
  <n-modal
    v-model:show="isShowModal"
    :mask-closable="false"
    preset="dialog"
    title="解绑账号"
    content=""
    positive-text="确认"
    negative-text="取消"
    @positive-click="onPositiveClick"
    @negative-click="onNegativeClick"
    :style="{
      width: dialogWidth,
    }"
  >
    <AccountTable @reloadView="reloadView" :proxyAddress="formParams?.address"></AccountTable>
  </n-modal>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref, watch} from 'vue';
import {useMessage} from 'naive-ui'
import {newState, State} from "@/views/whats/whatsProxy/model";
import {adaModalWidth} from "@/utils/hotgo";
import {View} from "@/api/whats/whatsProxy";
import {useRouter} from "vue-router";
import AccountTable from "@/views/whats/whatsProxy/account.vue";

const emit = defineEmits([ 'updateUnBindShowModal']);

interface Props {
  showModal: boolean;
  formParams?: State;
}

const message = useMessage()
const dialogWidth = ref('75%');
const router = useRouter();
const id = Number(router.currentRoute.value.query.id);
const formValue = ref(newState(null));
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
    emit('updateUnBindShowModal', value);
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
  console.log("unbind--",value)

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

watch(
  () => props.formParams,
  (value) => {
    loadForm(value);
  }
);

</script>

<style scoped lang="less">

</style>
