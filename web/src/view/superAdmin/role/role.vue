<template>
  <div class="role">
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button icon="plus" type="primary" @click="addRole(0)">新增角色</el-button>
      </div>
      <el-table
          :data="tableData"
          :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
          row-key="roleId"
          style="width: 100%"
      >
        <el-table-column label="角色ID" min-width="180" prop="roleId"/>
        <el-table-column align="left" label="角色名称" min-width="180" prop="roleName"/>
        <el-table-column align="left" label="操作" width="350">
          <template #default="scope">
            <el-button icon="setting" link type="primary" @click="openDrawer(scope.row)">设置权限</el-button>
            <el-button icon="plus" link type="primary" @click="addRole(scope.row.roleId)">新增子角色</el-button>
            <el-button icon="edit" link type="primary" @click="editRole(scope.row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dialogFormVisible" :title="dialogTitle">
      <el-form ref="roleForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="父级角色" prop="parentId">
          <el-cascader
              v-model="form.parentId"
              :disabled="dialogType==='add'"
              :options="RoleOption"
              :props="{ checkStrictly: true,label:'roleName',value:'roleId',disabled:'disabled',emitPath:false}"
              :show-all-levels="false"
              filterable
              style="width:100%"
          />
        </el-form-item>
        <el-form-item label="角色ID" prop="roleId">
          <el-input v-model="form.roleId" :disabled="dialogType==='edit'" autocomplete="off"/>
        </el-form-item>
        <el-form-item label="角色姓名" prop="roleName">
          <el-input v-model="form.roleName" autocomplete="off"/>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-drawer v-if="drawer" v-model="drawer" :with-header="false" class="auth-drawer" size="25%" title="角色配置">
      <el-tabs :before-leave="autoEnter" type="border-card">
        <el-tab-pane label="角色菜单">
          <Menus ref="menus" :row="activeRow" @changeRow="changeRow"/>
        </el-tab-pane>
        <el-tab-pane label="角色api">
          <Apis ref="apis" :row="activeRow" @changeRow="changeRow"/>
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  getRoleList,
  deleteRole,
  createRole,
  updateRole,
  copyRole
} from '@/api/swaRole'

import Menus from '@/view/superAdmin/role/components/menus.vue'
import Apis from '@/view/superAdmin/role/components/apis.vue'
import Datas from '@/view/superAdmin/role/components/datas.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'

import {ref} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'

const mustUint = (rule, value, callback) => {
  if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
    return callback(new Error('请输入正整数'))
  }
  return callback()
}

const RoleOption = ref([
  {
    roleId: 0,
    roleName: '根角色'
  }
])
const drawer = ref(false)
const dialogType = ref('add')
const activeRow = ref({})

const dialogTitle = ref('新增角色')
const dialogFormVisible = ref(false)
const apiDialogFlag = ref(false)
const copyForm = ref({})

const form = ref({
  roleId: 0,
  roleName: '',
  parentId: 0
})
const rules = ref({
  roleId: [
    {required: true, message: '请输入角色ID', trigger: 'blur'},
    {validator: mustUint, trigger: 'blur', message: '必须为正整数'}
  ],
  roleName: [
    {required: true, message: '请输入角色名', trigger: 'blur'}
  ],
  parentId: [
    {required: true, message: '请选择父角色', trigger: 'blur'},
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})

const getTableData = async () => {
  const table = await getRoleList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
  console.log('获取角色列表:', table)
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const changeRow = (key, value) => {
  activeRow.value[key] = value
}
const menus = ref(null)
const apis = ref(null)
const datas = ref(null)
const autoEnter = (activeName, oldActiveName) => {
  const paneArr = [menus, apis, datas]
  if (oldActiveName) {
    if (paneArr[oldActiveName].value.needConfirm) {
      paneArr[oldActiveName].value.enterAndNext()
      paneArr[oldActiveName].value.needConfirm = false
    }
  }
}
const copyRoleFunc = (row) => {
  setOptions()
  dialogTitle.value = '拷贝角色'
  dialogType.value = 'copy'
  for (const k in form.value) {
    form.value[k] = row[k]
  }
  copyForm.value = row
  dialogFormVisible.value = true
}
const openDrawer = (row) => {
  console.log('设置权限：', row)
  drawer.value = true
  activeRow.value = row
}
const deleteAuth = (row) => {
  ElMessageBox.confirm('此操作将永久删除该角色, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
      .then(async () => {
        const res = await deleteRole({roleId: row.roleId})
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '删除成功!'
          })
          if (tableData.value.length === 1 && page.value > 1) {
            page.value--
          }
          await getTableData()
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '已取消删除'
        })
      })
}
const roleForm = ref(null)
const initForm = () => {
  if (roleForm.value) {
    roleForm.value.resetFields()
  }
  form.value = {
    roleId: 0,
    roleName: '',
    parentId: 0
  }
}
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
  apiDialogFlag.value = false
}

const enterDialog = () => {
  form.value.roleId = Number(form.value.roleId)
  if (form.value.roleId === 0) {
    ElMessage({
      type: 'error',
      message: '角色id不能为0'
    })
    return false
  }
  roleForm.value.validate(async valid => {
    if (valid) {
      switch (dialogType.value) {
        case 'add': {
          const res = await createRole(form.value)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '添加成功!'
            })
            await getTableData()
            closeDialog()
          }
        }
          break
        case 'edit': {
          const res = await updateRole(form.value)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '添加成功!'
            })
            await getTableData()
            closeDialog()
          }
        }
          break
        case 'copy': {
          const data = {
            role: {
              roleId: 0,
              roleName: '',
              subRoles: [],
              parentId: 0
            },
            oldRoleId: 0
          }
          data.role.roleId = form.value.roleId
          data.role.roleName = form.value.roleName
          data.role.parentId = form.value.parentId
          data.role.subRoles = copyForm.value.subRoles
          data.oldRoleId = copyForm.value.roleId
          const res = await copyRole(data)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '复制成功！'
            })
            await getTableData()
          }
        }
      }

      initForm()
      dialogFormVisible.value = false
    }
  })
}
const setOptions = () => {
  RoleOption.value = [
    {
      roleId: 0,
      roleName: '根角色'
    }
  ]
  setRoleOptions(tableData.value, RoleOption.value, false)
}
const setRoleOptions = (RoleData, optionsData, disabled) => {
  form.value.roleId = String(form.value.roleId)
  RoleData &&
  RoleData.forEach(item => {
    if (item.children && item.children.length) {
      const option = {
        roleId: item.roleId,
        roleName: item.roleName,
        disabled: disabled || item.roleId === form.value.roleId,
        children: []
      }
      setRoleOptions(
          item.children,
          option.children,
          disabled || item.roleId === form.value.roleId
      )
      optionsData.push(option)
    } else {
      const option = {
        roleId: item.roleId,
        roleName: item.roleName,
        disabled: disabled || item.roleId === form.value.roleId
      }
      optionsData.push(option)
    }
  })
}
const addRole = (parentId) => {
  initForm()
  dialogTitle.value = '新增角色'
  dialogType.value = 'add'
  form.value.parentId = parentId
  setOptions()
  dialogFormVisible.value = true
}
const editRole = (row) => {
  setOptions()
  dialogTitle.value = '编辑角色'
  dialogType.value = 'edit'
  for (const key in form.value) {
    form.value[key] = row[key]
  }
  setOptions()
  dialogFormVisible.value = true
}

</script>

<script>

export default {
  name: 'Role'
}
</script>

<style lang="scss">
.role {
  .el-input-number {
    margin-left: 15px;

    span {
      display: none;
    }
  }
}

.tree-content {
  margin-top: 10px;
  height: calc(100vh - 148px);
  overflow: auto;
}

</style>
