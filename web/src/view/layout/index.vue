<template>
  <div class="bg-gray-50 text-slate-700 dark:text-slate-500 dark:bg-slate-800 w-screen h-screen">
    <el-watermark v-if="config.show_watermark" :font="font" :z-index="9999" :gap="[180,150]" class="absolute inset-0 pointer-events-none" :content="userStore.userInfo.nickName" />
    <gva-header />
    <div class="flex flex-row w-full gva-container  pt-16">
      <gva-aside />
      <div class="flex-1  w-0 h-full">
        <gva-tabs v-if="config.showTabs" />
        <div class="overflow-auto" :class="config.showTabs? 'gva-container2' :'gva-container pt-1'">
          <router-view v-if="reloadFlag" v-slot="{ Component }">
            <div id="gva-base-load-dom" class="gva-body-h bg-gray-50 dark:bg-slate-800">
              <transition mode="out-in" name="el-fade-in-linear">
                <keep-alive :include="routerStore.keepAliveRouters">
                  <component :is="Component" />
                </keep-alive>
              </transition>
            </div>
          </router-view>
          <BottomInfo />
        </div>
<<<<<<< HEAD
        <Aside class="aside" />
      </el-aside>
      <!-- 分块滑动功能 -->
      <el-main class="main-cont main-right">
        <transition
          :duration="{ enter: 800, leave: 100 }"
          mode="out-in"
          name="el-fade-in-linear"
        >
          <div
            :style="{width: `calc(100% - ${getAsideWidth()})`}"
            class="fixed top-0 box-border z-50"
          >
            <el-row>
              <el-col>
                <el-header class="header-cont">
                  <el-row class="p-0 h-full">
                    <el-col
                      :xs="2"
                      :lg="1"
                      :md="1"
                      :sm="1"
                      :xl="1"
                      class="z-50 flex items-center pl-3"
                    >
                      <div
                        class="text-black cursor-pointer text-lg leading-5"
                        @click="totalCollapse"
                      >
                        <div
                          v-if="isCollapse"
                          class="gvaIcon gvaIcon-arrow-double-right"
                        />
                        <div
                          v-else
                          class="gvaIcon gvaIcon-arrow-double-left"
                        />
                      </div>
                    </el-col>
                    <el-col
                      :xs="10"
                      :lg="14"
                      :md="14"
                      :sm="9"
                      :xl="14"
                      :pull="1"
                      class="flex items-center"
                    >
                      <!-- 修改为手机端不显示顶部标签 -->
                      <el-breadcrumb
                        v-show="!isMobile"
                        class="breadcrumb"
                      >
                        <el-breadcrumb-item
                          v-for="item in matched.slice(1,matched.length)"
                          :key="item.path"
                        >{{ fmtTitle(item.meta.title,route) }}</el-breadcrumb-item>
                      </el-breadcrumb>
                    </el-col>
                    <el-col
                      :xs="12"
                      :lg="9"
                      :md="9"
                      :sm="14"
                      :xl="9"
                      class="flex items-center justify-end"
                    >
                      <div class="mr-1.5 flex items-center">
                        <Search />
                        <el-dropdown>
                          <div class="flex justify-center items-center h-full w-full">
                            <span class="cursor-pointer flex justify-center items-center">
                              <CustomPic />
                              <span
                                v-show="!isMobile"
                                style="margin-left: 5px"
                              >{{ userStore.userInfo.nickName }}</span>
                              <el-icon>
                                <arrow-down />
                              </el-icon>
                            </span>
                          </div>
                          <template #dropdown>
                            <el-dropdown-menu>
                              <el-dropdown-item>
                                <span class="font-bold">
                                  当前角色：{{ userStore.userInfo.authority.authorityName }}
                                </span>
                              </el-dropdown-item>
                              <template v-if="userStore.userInfo.authorities">
                                <el-dropdown-item
                                  v-for="item in userStore.userInfo.authorities.filter(i=>i.authorityId!==userStore.userInfo.authorityId)"
                                  :key="item.authorityId"
                                  @click="changeUserAuth(item.authorityId)"
                                >
                                  <span>
                                    切换为：{{ item.authorityName }}
                                  </span>
                                </el-dropdown-item>
                              </template>
                              <el-dropdown-item icon="avatar">
                                <div
                                  class="command-box"
                                  style="display: flex"
                                  @click="handleCommand"
                                >
                                  <div>指令菜单</div>
                                  <div style="margin-left: 8px">
                                    <span class="button">{{ first }}</span>
                                    +
                                    <span class="button">K</span>
                                  </div>
                                </div>
                              </el-dropdown-item>
                              <el-dropdown-item
                                icon="avatar"
                                @click="toPerson"
                              >个人信息</el-dropdown-item>
                              <el-dropdown-item
                                icon="reading-lamp"
                                @click="userStore.LoginOut"
                              >登 出</el-dropdown-item>
                            </el-dropdown-menu>
                          </template>
                        </el-dropdown>
                      </div>
                    </el-col>
                  </el-row>
                </el-header>
              </el-col>
            </el-row>
            <!-- 当前面包屑用路由自动生成可根据需求修改 -->
            <!--
            :to="{ path: item.path }" 暂时注释不用-->
            <HistoryComponent ref="layoutHistoryComponent" />
          </div>
        </transition>
        <router-view
          v-if="reloadFlag"
          v-slot="{ Component }"
          class="admin-box"
        >
          <div
            id="gva-base-load-dom"
          >
            <transition
              mode="out-in"
              name="el-fade-in-linear"
            >
              <keep-alive :include="routerStore.keepAliveRouters">
                <component :is="Component" />
              </keep-alive>
            </transition>
          </div>
        </router-view>
        <!-- <BottomInfo /> -->
        <setting />
        <CommandMenu ref="command" />
      </el-main>
    </el-container>

  </el-container>
=======
      </div>
    </div>
  </div>
>>>>>>> 613ba9ba201cdb036e66a5c54abbff1679ad2368
</template>

<script setup>
import GvaAside from "@/view/layout/aside/index.vue";
import GvaHeader from "@/view/layout/header/index.vue";
import useResponsive  from "@/hooks/responsive";
import GvaTabs from "./tabs/index.vue"
import BottomInfo from "@/components/bottomInfo/bottomInfo.vue";
import { emitter } from "@/utils/bus.js";
import { ref, onMounted, nextTick, reactive, watchEffect} from "vue";
import { useRouter, useRoute } from "vue-router";
import { useRouterStore } from "@/pinia/modules/router";
import { useUserStore } from "@/pinia/modules/user";
import { useAppStore } from '@/pinia'
import { storeToRefs } from 'pinia'
const appStore = useAppStore()
const { config, theme } = storeToRefs(appStore)

defineOptions({
  name: "GvaLayout",
});

useResponsive(true)
const font = reactive({
  color: 'rgba(0, 0, 0, .15)',
})

watchEffect(()=>{
  font.color = theme.value === 'dark' ? 'rgba(255,255,255, .15)'  : 'rgba(0, 0, 0, .15)'

})

const router = useRouter();
const route = useRoute();
const routerStore = useRouterStore();

onMounted(() => {
  // 挂载一些通用的事件
  emitter.on("reload", reload);
  if (userStore.loadingInstance) {
    userStore.loadingInstance.close();
  }
});

const userStore = useUserStore();

const reloadFlag = ref(true);
let reloadTimer = null;
const reload = async () => {
  if (reloadTimer) {
    window.clearTimeout(reloadTimer);
  }
  reloadTimer = window.setTimeout(async () => {
    if (route.meta.keepAlive) {
      reloadFlag.value = false;
      await nextTick();
      reloadFlag.value = true;
    } else {
      const title = route.meta.title;
      router.push({ name: "Reload", params: { title } });
    }
  }, 400);
};




</script>

<style lang="scss">


</style>
