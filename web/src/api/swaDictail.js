import service from '@/utils/request'
export const createSwaDictail = (data) => {
    return service({
        url: '/swaDictail/createSwaDictail',
        method: 'post',
        data
    })
}

export const deleteSwaDictail = (data) => {
    return service({
        url: '/swaDictail/deleteSwaDictail',
        method: 'delete',
        data
    })
}

export const updateSwaDictail = (data) => {
    return service({
        url: '/swaDictail/updateSwaDictail',
        method: 'put',
        data
    })
}

export const findSwaDictail = (params) => {
    return service({
        url: '/swaDictail/findSwaDictail',
        method: 'get',
        params
    })
}

export const getSwaDictailList = (params) => {
    return service({
        url: '/swaDictail/getSwaDictailList',
        method: 'get',
        params
    })
}
