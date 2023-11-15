<template>
  <n-card :bordered="false" class="proCard">
    <div :class="['chat-item', isActive ? 'active' : '']">
      <div class="chat-item-left">
        <n-avatar
          round
          :size="54"
          color="transparent"
          :src="
            data.avatar ??
            'https://gw.alipayobjects.com/zos/antfincdn/aPkFc8Sj7n/method-draw-image.svg'
          "
        />
      </div>
      <div class="chat-item-right">
        <div class="chat-item-right-info">
          <div class="chat-item-right-info-name">{{ data.firstName + " " + data.lastName }}</div>
          <div class="chat-item-right-info-meta">
            <n-space :size="4">
              <span class="chat-item-right-info-meta-read">{{ data.last.read===1?'已读':'未读' }}</span>
              <span class="chat-item-right-info-meta-date">{{ data.last.sendTime }}</span>
            </n-space>
          </div>
        </div>
        <p class="chat-item-right-message">{{ data.last.sendMsg }}</p>
      </div>
    </div>
  </n-card>
</template>
<script lang="ts" setup>


import {TChatItemParam} from "@/views/tg/chat/components/model";
import CryptoJS from 'crypto-js';

interface IChatItemProps {
  isActive?: boolean;
  data: TChatItemParam;
}

defineProps<IChatItemProps>();

function base64Dec(base64Str: string) {
  let parsedWordArray = CryptoJS.enc.Base64.parse(base64Str);
  return parsedWordArray.toString(CryptoJS.enc.Utf8);
}

</script>
<style lang="less" scoped>
.chat-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px;
  margin: 0 8px;
  border-radius: 12px;
  cursor: pointer;

  &:hover {
    background-color: #f4f4f5;
  }

  &-left {
    display: flex;
  }

  &-right {
    flex: 1;
    overflow: hidden;
    line-height: 28px;

    &-info {
      display: flex;
      align-items: center;
      justify-content: space-between;

      &-name {
        font-size: 16px;
        color: #000;
      }

      &-meta {
        font-size: 12px;

        &-read {
          color: #3390ec;
        }

        &-date {
          color: #686c72;
        }
      }
    }

    &-message {
      max-width: 100%;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      font-size: 16px;
      color: #707579;
    }
  }

  &.active {
    background-color: #3390ec;

    .chat-item-right-info-name,
    .chat-item-right-info-meta-read,
    .chat-item-right-info-meta-date,
    .chat-item-right-message {
      color: #fff;
    }
  }
}
</style>
