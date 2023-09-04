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
          {{ '上传' + typeTag }}
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
        accept=".xlsx,.xlx"
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
          <n-text style="font-size: 16px"> 点击或者拖动{{ typeTag }}到该区域来上传</n-text>
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
import {Upload} from '@/api/whats/whatsAccount'

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
  a.href = "./static/账号模板.xlsx";
  a.download = "账号模板.xlsx";
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
  if (extension != 'xlsx' && extension != 'xls') {
    message.error('只能上传xlsx格式的excel文件，请重新上传');
    return;
  }

  const fileReader = new FileReader();
  fileReader.onload = (e) => {
    const data = new Uint8Array(e.target.result);
    const workbook = XLSX.read(data, {type: 'array'});
    const weldmachine = XLSX.utils.sheet_to_json(workbook.Sheets[workbook.SheetNames[0]]);
    debugger

    //中英文映射
    var columnMapping = {
      '账号': 'account',
      '号码ID': 'identify',
      '公钥': 'publicKey',
      '私钥': 'privateKey',
      '消息公钥': 'publicMsgKey',
      '消息私钥': 'privateMsgKey',
    };
    var newrows = [];
    //中英文转换
    for (let row of weldmachine) {
      var newRow = {};
      // 使用for...in循环遍历对象的属性
      for (let key in row) {
        debugger
        //要是能在colmnMapping[key]中找到对应的英文，则变成对应的数据，否则还是用原来的
        if (columnMapping[key] === undefined) {
          message.error('解析失败, 请检查文件格式');
          uploadRef.value.clear();
          return;
        }
        newRow[columnMapping[key] || key] = row[key];
      }
      newrows.push(newRow);
    }
    //上传数据
    showTableModal.value = true;
    tableData.value = newrows;

  }
  fileReader.onerror = (e) => {
    message.error('解析失败, 请检查文件格式');
    showLoading.value = false;
  }
  fileReader.readAsArrayBuffer(file.file as File);
}


function handleSumbit() {
  if (tableData.value.length > 0) {
    showLoading.value = true;
    // 编辑
    Upload({"list": tableData.value})
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
