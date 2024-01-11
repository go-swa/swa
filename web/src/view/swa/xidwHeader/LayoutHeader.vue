<script>
export default {
  name: 'LayoutHeader',
}
</script>
<template>
  <div id="layoutHeader">
    <div class="box-header">
      <div class="logo-name">
        <img alt="" class="img-size" src="@/assets/swa-48.png">
        <span class="name">{{ CONFIG.appName }}</span>
      </div>
      <div class="top-menus">
        <el-menu
            :active-text-color="theme.active"
            :default-active="activeIndex"
            :ellipsis="false"
            mode="horizontal"
        >
          <template v-for="(item, idx) in menusTop">
            <template v-if="idx < menusTop.length-1">
              <el-menu-item :key="idx" :index="''+idx" @click="activeTop(idx)">
                {{ item.meta.title }}
              </el-menu-item>
            </template>
          </template>
        </el-menu>

        <div class="person-operation">
          <div class="search-component">
            <div class="user-box">
              <UserMessage/>
            </div>
            <div class="user-box">
              <Screenfull :style="{cursor:'pointer'}" class="search-icon"/>
            </div>
            <div class="user-box">
              <ColorSetting/>
            </div>
          </div>
          <div class="username">
            <el-dropdown>
              <div class="panel">
                <span>你好: {{ userStore.userInfo.nickName }}</span>
                <el-icon>
                  <arrow-down/>
                </el-icon>
              </div>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item><span style="font-weight: 600;">当前角色：{{ userStore.userInfo.swaRole.roleName }}</span></el-dropdown-item>
                  <el-dropdown-item icon="avatar" @click="toPerson">个人信息</el-dropdown-item>
                  <el-dropdown-item icon="reading-lamp" @click="userStore.LoginOut">退出登录</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import CONFIG from '@/core/config'
import {storeToRefs} from 'pinia'
import {useMenuStore} from '@/stores/menus'
import {useRouter} from 'vue-router'
import Screenfull from '@/components/screenfull/index.vue'
import ColorSetting from '@/components/setting/index.vue'
import UserMessage from '@/components/userMessage/userMessage.vue'

import {ref} from 'vue'

const menuStore = useMenuStore()
import {useUserStore} from '@/pinia/modules/user'
import {ArrowDown} from '@element-plus/icons-vue'

const userStore = useUserStore()

const activeIndex = ref('0')
const activeTop = (i) => {
  menuStore.doTopMenusActive(i)
}
const router = useRouter()
const theme = ref({})
const getTheme = () => {
  switch (userStore.sideMode) {
    case '#fff':
      theme.value = {
        background: '#fff',
        activeBackground: 'var(--el-color-primary)',
        activeText: '#c122a7',
        normalText: '#333',
        hoverBackground: 'rgba(64, 158, 255, 0.08)',
        hoverText: '#333',
      }
      break
    case '#191a23':
      theme.value = {
        background: '#191a23',
        activeBackground: 'var(--el-color-primary)',
        activeText: '#6ed056',
        normalText: '#fff',
        hoverBackground: 'rgba(64, 158, 255, 0.08)',
        hoverText: '#fff',
      }
      break
  }
}
getTheme()

const {
  menusTop,
} = storeToRefs(menuStore)

const toPerson = () => {
  router.push({name: 'person'})
}

</script>

<style lang="scss">
@import "@/view/swa/xidwHeader/LayoutHeader.scss";
</style>
