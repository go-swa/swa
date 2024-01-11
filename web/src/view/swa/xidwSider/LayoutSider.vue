<script>
export default {
  name: 'LayoutSider'
}
</script>

<style lang="scss">
@import '@/view/swa/xidwSider/LayoutSider.scss';
</style>

<template>
  <div id="layoutSider">
    <div :style="{ background: userStore.sideMode }" class="box-aside">
      <el-scrollbar>
        <el-menu
            :active-text-color="theme.active"
            :background-color="theme.background"
            :collapse="isCollapse"
            :collapse-transition="false"
            :default-active="active"
            class="el-menu-vertical"
            unique-opened
            @select="selectMenuItem"
        >
          <template v-for="item in menusAside">
            <aside-component
                v-if="!item.hidden"
                :key="item.name"
                :is-collapse="isCollapse"
                :router-info="item"
                :theme="theme"
            />
          </template>
        </el-menu>
      </el-scrollbar>
    </div>
  </div>
</template>

<script setup>
import {ref, watch, onUnmounted} from 'vue'
import {storeToRefs} from 'pinia'
import {useMenuStore, moreToOneArray} from '@/stores/menus'
import {useUserStore} from '@/pinia/modules/user'
import {useRouterStore} from '@/pinia/modules/router'
import {useRoute, useRouter} from 'vue-router'
import Icon from '@/view/superAdmin/menu/icon.vue'
import AsideComponent from '@/components/asideComponent/index.vue'
import {emitter} from '@/utils/bus'

const menuStore = useMenuStore()
const {menusAside, menusAsideActive} = storeToRefs(menuStore)
const routerStore = useRouterStore()
const route = useRoute()
const router = useRouter()
const disposeMenusAside = ref([])
const userStore = useUserStore()

watch(
    menusAside,
    () => {
      disposeMenusAside.value = moreToOneArray(menusAside.value)
    },
    {immediate: true}
)

const isCollapse = ref(false)

const theme = ref({})

const getTheme = () => {
  switch (userStore.sideMode) {
    case '#fff':
      theme.value = {
        background: '#fff',
        activeBackground: 'var(--el-color-primary)',
        activeText: '#fff',
        normalText: '#333',
        hoverBackground: 'rgba(64, 158, 255, 0.08)',
        hoverText: '#333',
      }
      break
    case '#191a23':
      theme.value = {
        background: '#2d3a4b',
        activeBackground: 'var(--el-color-primary)',
        activeText: '#fff',
        normalText: '#fff',
        hoverBackground: 'rgba(64, 158, 255, 0.08)',
        hoverText: '#fff',
      }
      break
  }
}
getTheme()

const active = ref('')

watch(() => route, () => {
  active.value = route.meta.activeName || route.name
}, {deep: true})

watch(() => userStore.sideMode, () => {
  getTheme()
})

const initPage = () => {
  active.value = route.meta.activeName || route.name
  const screenWidth = document.body.clientWidth
  if (screenWidth < 1000) {
    isCollapse.value = !isCollapse.value
  }

  emitter.on('collapse', (item) => {
    isCollapse.value = item
  })
}

initPage()

onUnmounted(() => {
  emitter.off('collapse')
})
const selectMenuItem = (routeName, _, ele, aaa) => {
  const query = {}
  const params = {}
  routerStore.routeMap[routeName]?.parameters &&
  routerStore.routeMap[routeName]?.parameters.forEach((item) => {
    if (item.type === 'query') {
      query[item.key] = item.value
    } else {
      params[item.key] = item.value
    }
  })
  if (routeName === route.name) return
  if (routeName.indexOf('http://') > -1 || routeName.indexOf('https://') > -1) {
    window.open(routeName)
  } else {
    router.push({name: routeName, query, params})
  }
}
</script>
