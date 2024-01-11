import service from '@/utils/request'
export const UpdateCasbin = (data) => {
    return service({
        url: '/casbin/updateCasbin',
        method: 'post',
        data
    })
}

export const getCasbinByRoleId = (data) => {
    return service({
        url: '/casbin/getCasbinByRoleId',
        method: 'post',
        data
    })
}
