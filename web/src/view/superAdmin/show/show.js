import service from '@/utils/request'

export const getSysTbs = () => {
    return service({
        url: '/sys/getSysTbs',
        method: 'get',
    })
}

export const getTbShow = (data) => {
    return service({
        url: '/sys/getTbShow',
        method: 'post',
        data
    })
}
