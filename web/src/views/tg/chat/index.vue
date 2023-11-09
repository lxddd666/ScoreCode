<template>
  <div class="chat-root">
    <n-grid cols="24 300:1 600:24" :x-gap="0" style="height: 100%">
      <n-grid-item span="6" style="overflow: hidden">
        <n-card :bordered="false" class="proCard aside">
          <div class="search">
            <div class="search-left">
              <n-dropdown
                v-if="!showBackIcon"
                trigger="click"
                placement="bottom-end"
                :options="chatOptions"
              >
                <n-icon size="20" class="search-left-icon">
                  <MenuOutlined/>
                </n-icon>
              </n-dropdown>
              <n-icon v-else size="20" class="search-left-icon" @click="onBackClick">
                <ArrowLeftOutlined/>
              </n-icon>
            </div>
            <div class="search-right">
              <n-input round placeholder="搜索" @click="onSearchClick">
                <template #prefix>
                  <n-icon :component="SearchOutlined"/>
                </template>
              </n-input>
            </div>
          </div>
          <div class="chat-list">
            <n-scrollbar trigger="hover">
              <ChatItem
                v-for="item in chatList"
                :key="item.phone"
                :is-active="activeItem.phone === item.phone"
                :data="item"
                @click="onChatItemClick(item)"
              />
            </n-scrollbar>
          </div>
        </n-card>
      </n-grid-item>
      <n-grid-item span="18" style="overflow: hidden">
        <ChatArea
          :data="activeItem"
          :phone="phone"
        >
        </ChatArea>
      </n-grid-item>
    </n-grid>

  </div>
</template>
<script lang="ts" setup>
import {onMounted, ref} from 'vue';
import {ArrowLeftOutlined, MenuOutlined, SearchOutlined} from '@vicons/antd';
import {DropdownMixedOption} from 'naive-ui/lib/dropdown/src/interface';
import ChatItem from './components/ChatItem.vue';
import ChatArea from './components/ChatArea.vue';
import router from "@/router";
import {TgGetDialogs, TgLogin} from "@/api/tg/tgUser";
import {defaultState, TChatItemParam} from "@/views/tg/chat/components/model";

const chatOptions = ref<DropdownMixedOption[]>([
  {
    label: '联系人',
    key: 'contacts',
  },
  {
    label: '设置',
    key: 'settings',
  },
]);
const showBackIcon = ref(false);
const chatList = ref<TChatItemParam[]>([]);
const activeItem = ref<TChatItemParam>(defaultState);
const phone = ref<number>(0);
const id = Number(router.currentRoute.value.params.id);
const onSearchClick = () => {
  showBackIcon.value = true;
};
const onBackClick = () => {
  showBackIcon.value = false;
};
const onChatItemClick = (item: TChatItemParam) => {
  activeItem.value = item;
};
const getChatList = async (account: number) => {
  const res = await TgGetDialogs({account: account});
  console.log(res)
  chatList.value = res.list;
  activeItem.value = res.list[0];
};

const load = async (id: number) => {
  const logInfo = await TgLogin({id: id});
  phone.value = logInfo.phone;
  await getChatList(logInfo.phone);
}

onMounted(() => {
  load(id);
});
</script>
<style lang="less" scoped>
.chat-root {
  height: calc(100vh - 120px);

  .aside {
    border-radius: 0px;
    height: 100%;
    overflow: hidden;

    :deep(.n-card__content) {
      display: flex;
      flex-direction: column;
      overflow: hidden;
      padding: 0px;
    }

    .search {
      display: flex;
      align-items: center;
      gap: 8px;
      padding: 8px;

      &-left {
        &-icon {
          width: 40px;
          height: 40px;
          display: flex;
          align-items: center;
          justify-content: center;
          border-radius: 50%;
          color: #707579;
          cursor: pointer;

          &:hover {
            background-color: #f4f4f4;
          }
        }
      }

      &-right {
        flex: 1;
      }
    }

    .chat-list {
      flex: 1;
      overflow: hidden;

      :deep(.n-scrollbar-rail.n-scrollbar-rail--vertical) {
        right: 0;
      }
    }
  }

  .chat-area {
    height: 100%;
  }
}
</style>
