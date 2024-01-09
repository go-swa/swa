import service from '@/utils/request'
export const createSwaDict = (data) => {
  return service({
    url: '/swaDict/createSwaDict',
    method: 'post',
    data
  })
}

export const deleteSwaDict = (data) => {
  return service({
    url: '/swaDict/deleteSwaDict',
    method: 'delete',
    data
  })
}

export const updateSwaDict = (data) => {
  return service({
    url: '/swaDict/updateSwaDict',
    method: 'put',
    data
  })
}

export const findSwaDict = (params) => {
  return service({
    url: '/swaDict/findSwaDict',
    method: 'get',
    params
  })
}

export const getSwaDictList = (params) => {
  return service({
    url: '/swaDict/getSwaDictList',
    method: 'get',
    params
  })
}
