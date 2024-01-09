import service from '@/utils/request'
export const getRoleList = (data) => {
  return service({
    url: '/swaRole/getRoleList',
    method: 'post',
    data
  })
}

export const deleteRole = (data) => {
  return service({
    url: '/swaRole/deleteRole',
    method: 'post',
    data
  })
}

export const createRole = (data) => {
  return service({
    url: '/swaRole/createRole',
    method: 'post',
    data
  })
}

export const copyRole = (data) => {
  return service({
    url: '/swaRole/copyRole',
    method: 'post',
    data
  })
}

export const setSubRoles = (data) => {
  return service({
    url: '/swaRole/setSubRoles',
    method: 'post',
    data
  })
}

export const updateRole = (data) => {
  return service({
    url: '/swaRole/updateRole',
    method: 'put',
    data
  })
}

export const setRoleMenus = (data) => {
  return service({
    url: '/swaRole/setRoleMenus',
    method: 'post',
    data
  })
}

export const getSubRoles = (data) => {
  return service({
    url: '/swaRole/getSubRoles',
    method: 'post',
    data
  })
}
