import config from './config'

import * as ElIconModules from '@element-plus/icons-vue'

export const register = (app) => {
    for (const iconName in ElIconModules) {
        app.component(iconName, ElIconModules[iconName])
    }
    app.config.globalProperties.$GIN_VUE_ADMIN = config
}
