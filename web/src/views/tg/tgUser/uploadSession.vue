<!--suppress ALL -->
<template>
    <n-modal
      v-model:show="showFileModal"
      :show-icon="false"
      preset="dialog"
      :style="{
      width: width,
    }"
    >
      <template #header>
        <n-tag checkable disabled>
          {{ '上传' + 'zip' }}
        </n-tag>
        <n-button @click="handleDownload" secondary type="info">
          <template #icon>
            <n-icon>
              <DownloadOutlined/>
            </n-icon>
          </template>
          下载模板
        </n-button>
      </template>
      <n-upload
        directory-dnd
        accept=".zip,"
        :data="{ type: 0 }"
        @finish="finish"
        :custom-request="customRequest"
        name="file"
        :max="maxUpload"
        ref="uploadRef"
      >
        <n-upload-dragger>
          <div style="margin-bottom: 12px">
            <n-icon size="48" :depth="3">
              <FileAddOutlined/>
            </n-icon>
          </div>
          <n-text style="font-size: 16px"> 点击或者拖动zip到该区域来上传</n-text>
        </n-upload-dragger>
      </n-upload>
    </n-modal>

    <n-modal
      v-model:show="showTableModal"
      preset="dialog"
      :bordered="false"
      :on-close="handleClose"
      :style="{
      width: width,
    }">
      <template #header>
        <n-tag checkable disabled>
          {{ '共' + tableData.length + '条数据' }}
        </n-tag>
        <n-button @click="handleSumbit" :loading="showLoading" secondary type="info">
          确认提交
        </n-button>
      </template>

      <n-data-table
        :columns="uploadColumns"
        :resizeHeightOffset="-10000"
        size="small"
        :data="tableData"
        :max-height="500"
        :scroll-x="1800"
        :pagination="20"
        striped
      >
        <template #toolbar>
          <n-button
            type="primary"
            class="min-left-space"
          >
            确认上传
          </n-button>
          <n-button
            class="min-left-space"
          >
            取消
          </n-button>

        </template>
      </n-data-table>
    </n-modal>



</template>

<script lang="ts" setup>
import {computed, reactive, ref, h} from 'vue';
import {
  DeleteOutlined,
  FileAddOutlined,
  UploadOutlined,
  DownloadOutlined,
  PlusOutlined
} from '@vicons/antd';
import {useUserStoreWidthOut} from '@/store/modules/user';
import {useGlobSetting} from '@/hooks/setting';
import {NModal, UploadFileInfo, useMessage} from 'naive-ui';
import * as XLSX from 'xlsx';
import componentSetting from '@/settings/componentSetting';
import {ResultEnum} from '@/enums/httpEnum';
import {
  Attachment,
  FileType,
  getFileType,
  getFileExtension
} from '@/components/FileChooser/src/model';
import {columns, uploadColumns} from './model';
import {ImportSession} from "@/api/tg/tgUser";

const emit = defineEmits(['reloadTable']);

export interface Props {
  width?: string;
  maxUpload?: number;
  finishCall?: Function | null;
  uploadType?: FileType;
}

const props = withDefaults(defineProps<Props>(), {
  width: '60%',
  maxUpload: 1,
  finishCall: null,
  uploadType: 'xlsx',
});

const showLoading = ref<boolean>(false);
const uploadRef = ref();
const showFileModal = ref(false);
const showTableModal = ref(false);
const message = useMessage();
const tableData = ref([]);


const typeTag = computed(() => {
  return getFileType(props.uploadType);
});

function handleDownload() {
  let a = document.createElement("a");
  a.href = "./static/session.zip";
  a.download = "session.zip";
  a.style.display = "none";
  document.body.appendChild(a);
  a.click();
  a.remove();

}


//自定义上传
function customRequest({
                         file,
                         data,
                         headers,
                         withCredentials,
                         action,
                         onFinish,
                         onError,
                         onProgress
                       }: UploadCustomRequestOptions) {
  var extension = getFileExtension(file.file.name);
  if (extension != 'zip') {
    message.error('只能上传zip格式的文件，请重新上传');
    return;
  }

  const fileReader = new FileReader();

}


function handleSumbit() {
  debugger
  if (tableData.value.length > 0) {
    showLoading.value = true;
    console.log("hhhh")
    debugger
    // 导入
    ImportSession({"File": tableData.value})
      .then((res) => {
        message.success(res.message)
      })
      .finally(() => {
        tableData.value = [];
        showTableModal.value = false;
        showFileModal.value = false;
        showLoading.value = false;
        emit('reloadTable');
      });
  } else {
    tableData.value = [];
    showTableModal.value = false;
    message.warning('excel格式不正确, 请参考模板')
  }
}

function handleClose() {
  tableData.value = [];
}

function openModal() {
  showFileModal.value = true;
}

defineExpose({
  openModal,
});
</script>
