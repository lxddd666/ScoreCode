<!-- eslint-disable vue/no-v-html -->
<template>
  <div class="message-area">
    <div class="message-area-list">
      <n-scrollbar trigger="hover" ref="scrollRef">
        <div class="message-area-list-wrapper">
          <div
            :class="{ 'message-area-list-wrapper-item': true, isMe: item.out===1 }"
            v-for="item in data.msgList"
            :key="item.id"
          >
            <div class="message-area-list-wrapper-item-content">
              <div>{{ item.message }}</div>
              <span class="message-area-list-wrapper-item-content-meta">
                <span class="message-area-list-wrapper-item-content-meta-date">{{
                    item.date
                  }}</span>
                <span class="message-area-list-wrapper-item-content-meta-read">{{
                    item.id > 1 ? '已读' : '未读'
                  }}</span>
              </span>
            </div>
          </div>
        </div>
      </n-scrollbar>
    </div>
    <div class="message-area-input">
      <div class="message-area-input-wrapper">
        <n-popover placement="top-start" trigger="hover" class="emoji-popover" :show-arrow="false">
          <template #trigger>
            <n-icon size="24" class="message-area-input-wrapper-icon">
              <SmileOutlined/>
            </n-icon>
          </template>
          <div class="large-text">这是表情</div>
        </n-popover>
        <div
          ref="contentRef"
          @input="onContentInput"
          @keydown.enter.exact="onContentKeydown"
          class="message-area-input-wrapper-content"
          contenteditable
        ></div>
        <n-popover placement="top-end" trigger="hover" class="file-popover" :show-arrow="false">
          <template #trigger>
            <n-icon size="24" class="message-area-input-wrapper-icon">
              <PaperClipOutlined/>
            </n-icon>
          </template>
          <div class="large-text">这是附件</div>
        </n-popover>
      </div>
      <div class="message-area-input-button" @click="handleSendMsg">
        <n-icon size="26" class="message-area-input-button-icon">
          <PaperPlaneSharp/>
        </n-icon>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {PaperClipOutlined, SmileOutlined} from '@vicons/antd';
import {PaperPlaneSharp} from '@vicons/ionicons5';
import {ScrollbarInst} from 'naive-ui';
import {nextTick, ref, watch} from 'vue';
import {newState, TChatItemParam, TMessage} from "@/views/tg/chat/components/model";
import {TgGetMsgHistory, TgSendMsg} from "@/api/tg/tgUser";
import {sendMsg} from '@/utils/websocket';

const inputText = ref('');
const scrollRef = ref<ScrollbarInst>();
const contentRef = ref<HTMLDivElement>();
const messageList = ref<TMessage[]>([]);
const scrollToBottom = () => {
  nextTick(() => {
    scrollRef.value?.scrollBy({top: 100000000});
  });
};
const handleSendMsg = () => {
  let textMsg = inputText.value
  inputText.value = '';
  contentRef.value!.innerHTML = '';
  TgSendMsg({
    "account": props.me.phone,
    "receiver": props.data.tgId,
    "textMsg": [textMsg]
  })
  scrollToBottom();
};
const onContentInput = () => {
  inputText.value = contentRef.value?.innerText ?? '';
  // console.log('e---', e);
  // if (e.data) {
  //   inputText.value += e.data;
  // }
};
const onContentKeydown = (e: KeyboardEvent) => {
  if (!e.shiftKey) {
    if (!inputText.value) {
      // message.warning(`请输入内容后再发送~~~`);
    } else {
      handleSendMsg();
    }
    e.preventDefault();
    return false;
  } else {
  }
};


const emit = defineEmits(['updateTChatItem']);
const loadForm = async (account: number, offsetId: number | undefined, contact: number) => {
  if (props.data.msgList != null && props.data.msgList.length > 0) {
    return;
  }
  if (offsetId === undefined) {
    offsetId = 100;
  }
  const res = await TgGetMsgHistory({
    "contact": contact,
    "account": account,
    "offsetId": Number(offsetId) + 1
  })
  props.data.msgList = res.list.reverse();
  emit('updateTChatItem', props.data);
  scrollToBottom();
}

interface Props {
  data: TChatItemParam;
  me: TChatItemParam;
}

const props = withDefaults(defineProps<Props>(), {
  me: () => {
    return newState(null);
  },
  data: () => {
    return newState(null);
  },
});
watch(
  () => props.data,
  (value) => {
    loadForm(props.me.phone, props.data.last?.reqId, value.tgId);
  }
);
</script>
<style lang="less" scoped>
.message-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding-bottom: 20px;
  overflow: hidden;

  &-list {
    flex: 1;
    padding-top: 24px;
    overflow: hidden;
  // overflow-y: auto;

    &-wrapper {
      width: 75%;
      margin: 0 auto;

      &-item {
        display: flex;
        justify-content: flex-start;
        margin-bottom: 8px;

        &-content {
          word-break: break-word;
          white-space: pre-wrap;
          background-color: #fff;
          border-radius: 12px;
          padding: 6px 8px;
          font-size: 16px;
          color: #000;
          display: flex;
          align-items: flex-end;

          &-meta {
            display: flex;
            margin-left: 8px;
            gap: 4px;
            align-items: center;
            font-size: 12px;
            color: #686c72bf;
          }
        }

        &.isMe {
          justify-content: flex-end;

          .message-area-list-wrapper-item-content {
            background-color: #eeffde;

            &-meta {
              color: #4fae4e;
            }
          }
        }
      }
    }
  }

  &-input {
    display: flex;
    justify-content: center;
    align-items: flex-end;
    gap: 8px;

    &-wrapper {
      display: flex;
      align-items: flex-end;
      gap: 12px;
      background-color: #fff;
      padding: 18px 12px;
      border-radius: 12px;

      &-icon {
        color: #707579;
      }

      &-content {
        width: 520px;
        font-size: 16px;
        color: #000;
        vertical-align: text-bottom;
        max-height: 300px;
        overflow-y: auto;

        &:focus {
          border: none;
          outline: none;
        }
      }
    }

    &-button {
      width: 56px;
      height: 56px;
      border-radius: 50%;
      background-color: #fff;
      display: flex;
      align-items: center;
      justify-content: center;
      color: #3390ec;
      cursor: pointer;

      &:hover {
        background-color: #3390ec;
        color: #fff;
      }
    }
  }
}
</style>
<style lang="less">
[v-placement='top-start'] > .n-popover-shared.emoji-popover {
  margin-bottom: 20px;
  margin-left: -12px;
}

[v-placement='top-end'] > .n-popover-shared.file-popover {
  margin-bottom: 20px;
  margin-right: -12px;
}

.message-area-input-wrapper-content {
  img {
    display: inline-block;
    vertical-align: text-bottom;
  }
}
</style>
