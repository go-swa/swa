import {asyncRouterHandle} from '@/utils/asyncRouter'
import {emitter} from '@/utils/bus.js'
import {asyncMenu} from '@/api/menu'
import {defineStore} from 'pinia'
import {ref} from 'vue'
import {jsonInBlacklist} from '@/api/jwt'

const notLayoutRouterArr = []
const keepAliveRoutersArr = []
const nameMap = {}

const formatRouter = (asyncRouter, routeMap) => {
    asyncRouter && asyncRouter.forEach(item => {
        item.meta.btns = item.btns
        item.meta.hidden = item.hidden
        if (item.meta.defaultMenu === true) {
            notLayoutRouterArr.push({
                ...item,
                path: `/${item.path}`,
            })
        } else {
            routeMap[item.name] = item
            if (item.children && item.children.length > 0) {
                formatRouter(item.children, routeMap)
            }
        }
    })
}

const KeepAliveFilter = (routes) => {
    routes && routes.forEach(item => {
        if ((item.children && item.children.some(ch => ch.meta.keepAlive) || item.meta.keepAlive)) {
            item.component && item.component().then(val => {
                keepAliveRoutersArr.push(val.default.name)
                nameMap[item.name] = val.default.name
            })
        }
        if (item.children && item.children.length > 0) {
            KeepAliveFilter(item.children)
        }
    })
}

export const useRouterStore = defineStore('router', () => {
    const keepAliveRouters = ref([])
    const asyncRouterFlag = ref(0)
    const setKeepAliveRouters = (history) => {
        const keepArrTemp = []
        history.forEach(item => {
            if (nameMap[item.name]) {
                keepArrTemp.push(nameMap[item.name])
            }
        })
        keepAliveRouters.value = Array.from(new Set(keepArrTemp))
    }
    emitter.on('setKeepAlive', setKeepAliveRouters)

    const asyncRouters = ref([])
    const routeMap = ({})
    const routerMenus = ref([])
    const SetAsyncRouter = async () => {
        asyncRouterFlag.value++
        const baseRouter = [{
            path: '/swa',
            name: 'swa',
            component: 'view/swa/index.vue',
            meta: {
                title: '微服务管理平台'
            },
            children: []
        }]
        const asyncRouterRes = await asyncMenu()
        console.log('获取后台菜单getMenu', asyncRouterRes)
        const asyncRouter = asyncRouterRes.data.menus

        routerMenus.value = asyncRouter
        asyncRouter && asyncRouter.push({
            path: 'reload',
            name: 'Reload',
            hidden: true,
            meta: {
                title: '',
                closeTab: true,
            },
            component: 'view/error/reload.vue'
        })

        formatRouter(asyncRouter, routeMap)
        baseRouter[0].children = asyncRouter
        if (notLayoutRouterArr.length !== 0) {
            baseRouter.push(...notLayoutRouterArr)
        }
        asyncRouterHandle(baseRouter)
        KeepAliveFilter(asyncRouter)
        asyncRouters.value = baseRouter
        return true
    }
    return {
        routerMenus,
        asyncRouters,
        keepAliveRouters,
        asyncRouterFlag,
        SetAsyncRouter,
        routeMap
    }
})

