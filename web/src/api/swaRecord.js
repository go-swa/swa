import service from '@/utils/request'
export const deleteSwaRecord = (data) => {
  return service({
    url: '/swaRecord/deleteSwaRecord',
    method: 'delete',
    data
  })
}

export const deleteSwaRecordByIds = (data) => {
  return service({
    url: '/swaRecord/deleteSwaRecordByIds',
    method: 'delete',
    data
  })
}

export const getSwaRecordList = (params) => {
  return service({
    url: '/swaRecord/getSwaRecordList',
    method: 'get',
    params
  })
}
