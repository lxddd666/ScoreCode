<template>
  <n-card :bordered="false" class="proCard">
    <div :class="['chat-item', isActive ? 'active' : '']">
      <div class="chat-item-left">
        <n-avatar
          round
          lazy
          :size="54"
          color="transparent"
          :src="data.avatar"
        ></n-avatar>
      </div>
      <div class="chat-item-right">
        <div class="chat-item-right-info">
          <div class="chat-item-right-info-name">{{
              data.type == 1 ? data.firstName + " " + data.lastName : data.title
            }}
          </div>
          <div class="chat-item-right-info-meta">
            <n-space :size="4">
              <span class="chat-item-right-info-meta-read">{{
                  data.unreadCount < 1 ? '已读' : '未读'
                }}</span>
              <span class="chat-item-right-info-meta-date">{{ data.last.date }}</span>
            </n-space>
          </div>
        </div>
        <p class="chat-item-right-message">{{ data.last.message }}</p>
      </div>
    </div>
  </n-card>
</template>
<script lang="ts" setup>


import {TChatItemParam} from "@/views/tg/chat/components/model";
import {TgGetUserAvatar} from "@/api/tg/tgUser";
import {watch} from "vue";

interface IChatItemProps {
  isActive?: boolean;
  data: TChatItemParam;
  me: TChatItemParam;
}

const props = defineProps<IChatItemProps>();




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
