import './styles/tailwind.css';
import {createApp} from 'vue';
import App from './App.vue';
import router, {setupRouter} from './router';
import {setupStore} from '@/store';
import {setupDirectives, setupNaive} from '@/plugins';
import {AppProvider} from '@/components/Application';
import Websocket from '@/utils/websocket';
// import * as buffer from "buffer";

async function bootstrap() {
  const appProvider = createApp(AppProvider);

  const app = createApp(App);

  // 注册全局常用的 naive-ui 组件
  setupNaive(app);


  // if (typeof (window as any).global === "undefined") {
  //   (window as any).global = window;
  // }
  // if (typeof (window as any).Buffer === "undefined") {
  //   (window as any).Buffer = buffer.Buffer;
  // }
  // 注册全局自定义组件
  //setupCustomComponents();

  // 注册全局自定义指令，如：v-permission权限指令
  setupDirectives(app);

  // 注册全局方法，如：app.config.globalProperties.$message = message
  //setupGlobalMethods(app);

  // 挂载状态管理
  setupStore(app);

  //优先挂载一下 Provider 解决路由守卫，Axios中可使用，Dialog，Message 等之类组件
  appProvider.mount('#appProvider', true);

  // 挂载路由
  await setupRouter(app);

  // 路由准备就绪后挂载APP实例
  await router.isReady();

  // 全局websocket
  const onMessageList: Array<Function> = [];
  app.provide('onMessageList', onMessageList);
  const onMessage = (event: any) => {
    onMessageList.forEach((f) => {
      f.call(null, event);
    });
  };
  Websocket(onMessage);

  app.mount('#app', true);
}

void bootstrap();
