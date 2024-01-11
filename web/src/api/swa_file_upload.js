import service from '@/utils/request'
export const getFileList = (data) => {
    return service({
        url: '/fileUpload/getFileList',
        method: 'post',
        data
    })
}

export const deleteFile = (data) => {
    return service({
        url: '/fileUpload/deleteFile',
        method: 'post',
        data
    })
}

/**
 * 编辑文件名或者备注
 * @param data
 * @returns {*}
 */
export const editFileName = (data) => {
    return service({
        url: '/fileUpload/editFileName',
        method: 'post',
        data
    })
}
