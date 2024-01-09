
import cookies from '@/plugin/cookies'
import { useRouter, useRoute } from 'vue-router'
import { defineStore } from 'pinia'
import { ref, watch } from 'vue'

import { useRouterStore } from '@/pinia/modules/router'
const menusData = useRouterStore().routerMenus

export const useMenuStore = defineStore('menus', () => {
  const router = useRouter()
  const route = useRoute()
  const menusTop = ref([])
  const menusTopActive = ref('')
  const menusAside = ref([])
  const menusAsideActive = ref('')
  const menusNav = ref([])
  const menusNavActive = ref('')

  watch(menusNavActive, (path) => {
    router.push(path || '/')
  })

  watch(menusNav, (newArr) => {
    cookies.set('menusNav', JSON.stringify(newArr))
  }, {
    deep: true
  })

  menusNavActive.value = route.path

  const doTopMenusActive = (index) => {
    menusTopActive.value = menusTop.value[index].path
    if (menusTop.value[index].children) {
      menusAside.value = menusTop.value[index].children
      if (getArrayInnerData(menusAside.value).length) {
        const asideFirstMemu = getArrayInnerData(menusAside.value)[0]
        menusAsideActive.value = asideFirstMemu.path
        doActiveMenusNav(asideFirstMemu)
      }
    } else {
      menusAside.value = []
    }
    return menusAsideActive.value
  }

  const doActiveMenusNav = (menuData) => {
    const menuArr = JSON.parse(JSON.stringify(menuData))
    if (menuArr.length > 0) {
      const menu = menuArr[0]
      const includeMenu = ref([])
      if (menusNav.value.length > 0) {
        includeMenu.value = menusNav.value.filter((i) => i.path === menu.path)
      }
      menusNavActive.value = menu.path
      menusAsideActive.value = menu.path
      if (includeMenu.value.length < 1 && menusNav.value.length <= 9) {
        menusNav.value.push(menu)
      } else if (includeMenu.value.length < 1) {
        menusNav.value.shift()
        menusNav.value.push(menu)
      }
    }
  }

  const initMenu = (menus) => {
    menusTop.value = menus
    if (menus[0] && menus[0].children && menus[0].children.length) {
      menusAside.value = menus[0].children
    }
  }
  initMenu(menusData)

  const doAllMenusConnect = (path) => {
    menusNavActive.value = path
    menusAsideActive.value = path
    if (path !== '/') {
      const m = menusTop.value.filter((e, i) => {
        const o = path.match(`${e.path}`)
        return o && o.index === 0
      })[0]
      if (m) {
        menusTopActive.value = m.path
      }
      if (m && m.children && m.children.length) {
        menusAside.value = m.children
      }
    }
  }

  (function doFirstIntoPage() {
    cookies.remove('menusNav')
    menusNav.value = cookies.get('menusNav') ? JSON.parse(cookies.get('menusNav')) : []
    doAllMenusConnect(route.path)
  })()

  return {
    menusTop,
    menusTopActive,
    menusAside,
    menusAsideActive,
    menusNav,
    menusNavActive,
    doTopMenusActive,
    doActiveMenusNav,
    doAllMenusConnect
  }
})

const getArrayInnerData = (data, list = []) => {
  data.forEach((item) => {
    if (item.children && item.children.length > 0) {
      getArrayInnerData(item.children, list)
    } else {
      list.push(item)
    }
  })
  return list
}

const getArraySpecifyData = (data = [], key, target) => {
  for (let i = 0; i < data.length; i++) {
    if (data[i][key] !== target) {
      if (data[i].children && data[i].children.length) {
        getArraySpecifyData(data[i].children, key, target)
      }
    } else {
      return data[i]
    }
  }
}

export const moreToOneArray = (arr) => {
  return [].concat(...arr.map(item => {
    return item.children
      ? [].concat(item, ...moreToOneArray(item.children))
      : [].concat(item)
  }))
}
