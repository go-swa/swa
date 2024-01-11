import service from '@/utils/request'

export const getRoleBtnApi = (data) => {
    return service({
        url: '/roleBtn/getRoleBtn',
        method: 'post',
        data
    })
}

export const setRoleBtnApi = (data) => {
    return service({
        url: '/roleBtn/setRoleBtn',
        method: 'post',
        data
    })
}

export const canRemoveRoleBtnApi = (params) => {
    return service({
        url: '/roleBtn/canRemoveRoleBtn',
        method: 'post',
        params
    })
}

