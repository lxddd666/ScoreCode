<template>
  <div class="chat-area">
    <n-card :bordered="false" class="proCard chat-area-head-card">
      <div class="chat-area-head">
        <div class="chat-area-head-left">
          <n-avatar :size="54" round
                    color="transparent"
                    :src="data.avatar==0?'https://gw.alipayobjects.com/zos/antfincdn/aPkFc8Sj7n/method-draw-image.svg':
            tgPrefix+'/arts/user/getUserAvatar?authorization='+token+'&account='+me.phone+'&getUser='+data.tgId+'&photoId='+data.avatar"
          ></n-avatar>
          <div class="chat-area-head-left-info">
            <div class="chat-area-head-left-info-name">{{
                data.type == 1 ? data.firstName + " " + data.lastName : data.title
              }}
            </div>
          </div>
        </div>
        <div class="chat-area-head-right">
          <n-space>
            <n-icon size="24" class="chat-area-head-right-icon">
              <SearchOutlined/>
            </n-icon>
            <n-icon size="24" class="chat-area-head-right-icon">
              <MoreOutlined/>
            </n-icon>
          </n-space>
        </div>
      </div>
    </n-card>
    <MessageArea :data="data" :me="me" @updateTChatItem="updateTChatItem"/>
    <div class="bg"></div>
  </div>
</template>
<script lang="ts" setup>
import {MoreOutlined, SearchOutlined} from '@vicons/antd';
import MessageArea from './MessageArea.vue';
import {TChatItemParam} from "@/views/tg/chat/components/model";
import {useGlobSetting} from "@/hooks/setting";
import {storage} from "@/utils/Storage";
import {ACCESS_TOKEN} from "@/store/mutation-types";

interface IChatItemProps {
  data: TChatItemParam;
  me: TChatItemParam;
}

const emit = defineEmits(['updateTChatItem']);

function updateTChatItem(item: TChatItemParam) {
  emit('updateTChatItem', item);
}

const globSetting = useGlobSetting();
const tgPrefix = globSetting.tgPrefix || '';

const token = storage.get(ACCESS_TOKEN);

defineProps<IChatItemProps>();

</script>
<style lang="less" scoped>
.chat-area {
  height: 100%;
  display: flex;
  flex-direction: column;
  position: relative;
  z-index: 1;

  &-head-card {
    z-index: 2;

    :deep(.n-card__content) {
      border-left: 1px solid #dadce0;
      padding: 8px 12px 8px 24px !important;
    }
  }

  &-head {
    display: flex;
    align-items: center;
    justify-content: space-between;

    &-left {
      display: flex;
      align-items: center;
      gap: 8px;

      &-info {
        &-name {
          font-size: 18px;
          color: #000;
          line-height: 22px;
        }

        &-status {
          font-size: 14px;
          color: #707579;
          line-height: 18px;
          display: flex;
          align-items: center;
          gap: 4px;
        }
      }
    }

    &-right {
      &-icon {
        color: #707579;
        cursor: pointer;
      }
    }
  }

  .bg {
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    z-index: -1;
    overflow: hidden;

    &::before {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      bottom: 0;
      right: 0;
      background-position: center;
      background-repeat: no-repeat;
      background-size: cover;
      background-image: url(../../../../assets/images/chat-bg.png);
    }

    &::after {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      bottom: 0;
      right: 0;
      background-image: url(../../../../assets/images/chat-bg-pattern.png);
      background-position: top right;
      background-size: 510px auto;
      background-repeat: repeat;
      mix-blend-mode: overlay;
    }
  }
}
</style>
