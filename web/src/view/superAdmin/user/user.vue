<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button icon="plus" type="primary" @click="addUser">新增用户</el-button>
      </div>
      <el-table
          :data="tableData"
          row-key="ID"
      >
        <el-table-column align="left" label="头像" min-width="75">
          <template #default="scope">
            <CustomPic :pic-src="scope.row.headerImg" style="margin-top:8px"/>
          </template>
        </el-table-column>
        <el-table-column align="left" label="ID" min-width="50" prop="ID"/>
        <el-table-column align="left" label="用户名" min-width="150" prop="userName"/>
        <el-table-column align="left" label="昵称" min-width="150" prop="nickName"/>
        <el-table-column align="left" label="手机号" min-width="180" prop="phone"/>
        <el-table-column align="left" label="邮箱" min-width="180" prop="email"/>
        <el-table-column align="left" label="用户角色" min-width="200">
          <template #default="scope">
            <el-cascader
                v-model="scope.row.roleIds"
                :clearable="false"
                :options="authOptions"
                :props="{ multiple:true,checkStrictly: true,label:'roleName',value:'roleId',disabled:'disabled',emitPath:false}"
                :show-all-levels="false"
                collapse-tags
                @visible-change="(flag)=>{changeRole(scope.row,flag,0)}"
                @remove-tag="(removeAuth)=>{changeRole(scope.row,false,removeAuth)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="启用" min-width="150">
          <template #default="scope">
            <el-switch
                v-model="scope.row.enable"
                :active-value="1"
                :inactive-value="2"
                inline-prompt
                @change="()=>{switchEnable(scope.row)}"
            />
          </template>
        </el-table-column>

        <el-table-column fixed="right" label="操作" min-width="250">
          <template #default="scope">
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除此用户吗</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button link type="primary" @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" @click="deleteUserFunc(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button icon="delete" link type="primary">删除</el-button>
              </template>
            </el-popover>
            <el-button icon="edit" link type="primary" @click="openEdit(scope.row)">编辑</el-button>
            <el-button icon="magic-stick" link type="primary" @click="resetPasswordFunc(scope.row)">重置密码</el-button>
          </template>
        </el-table-column>

      </el-table>
      <div class="gva-pagination">
        <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
        v-model="addUserDialog"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :show-close="false"
        class="user-dialog"
        title="用户"
    >
      <div style="height:60vh;overflow:auto;padding:0 12px;">
        <el-form ref="userForm" :model="userInfo" :rules="rules" label-width="80px">
          <el-form-item v-if="dialogFlag === 'add'" label="用户名" prop="userName">
            <el-input v-model="userInfo.userName"/>
          </el-form-item>
          <el-form-item v-if="dialogFlag === 'add'" label="密码" prop="password">
            <el-input v-model="userInfo.password"/>
          </el-form-item>
          <el-form-item label="昵称" prop="nickName">
            <el-input v-model="userInfo.nickName"/>
          </el-form-item>
          <el-form-item label="手机号" prop="phone">
            <el-input v-model="userInfo.phone"/>
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="userInfo.email"/>
          </el-form-item>
          <el-form-item label="用户角色" prop="roleId">
            <el-cascader
                v-model="userInfo.roleIds"
                :clearable="false"
                :options="authOptions"
                :props="{ multiple:true,checkStrictly: true,label:'roleName',value:'roleId',disabled:'disabled',emitPath:false}"
                :show-all-levels="false"
                style="width:100%"
            />
          </el-form-item>
          <el-form-item label="启用" prop="disabled">
            <el-switch
                v-model="userInfo.enable"
                :active-value="1"
                :inactive-value="2"
                inline-prompt
            />
          </el-form-item>
          <el-form-item label="头像" label-width="80px">
            <div style="display:inline-block" @click="openHeaderChange">
              <img v-if="userInfo.headerImg" :src="(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg" alt="头像" class="header-img-box">
              <div v-else class="header-img-box">从媒体库选择</div>
            </div>
          </el-form-item>

        </el-form>

      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeAddUserDialog">取 消</el-button>
          <el-button type="primary" @click="enterAddUserDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="userInfo" :target-key="`headerImg`"/>
  </div>
</template>

<script>
export default {
  name: 'User',
}
</script>

<script setup>

import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser
} from '@/api/user'

import {getRoleList} from '@/api/swaRole'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import {setUserInfo, resetPassword} from '@/api/user.js'

import {nextTick, ref, watch} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'

const path = ref(import.meta.env.VITE_BASE_API + '/')
const setRoleOptions = (RoleData, optionsData) => {
  RoleData &&
  RoleData.forEach(item => {
    if (item.children && item.children.length) {
      const option = {
        roleId: item.roleId,
        roleName: item.roleName,
        children: []
      }
      setRoleOptions(item.children, option.children)
      optionsData.push(option)
    } else {
      const option = {
        roleId: item.roleId,
        roleName: item.roleName
      }
      optionsData.push(option)
    }
  })
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const getTableData = async () => {
  const table = await getUserList({page: page.value, pageSize: pageSize.value})
  if (table.code === 0) {
    console.log('getUserList:', table)
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

watch(() => tableData.value, () => {
  setRoleIds()
})

const initPage = async () => {
  getTableData()
  const res = await getRoleList({page: 1, pageSize: 999})
  setOptions(res.data.list)
}

initPage()

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm(
      '是否将此用户密码重置为123456?',
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
  ).then(async () => {
    const res = await resetPassword({
      ID: row.ID,
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg,
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg,
      })
    }
  })
}
const setRoleIds = () => {
  tableData.value && tableData.value.forEach((user) => {
    user.roleIds = user.authorities && user.authorities.map(i => {
      return i.roleId
    })
  })
}

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

const authOptions = ref([])
const setOptions = (authData) => {
  authOptions.value = []
  setRoleOptions(authData, authOptions.value)
}

const deleteUserFunc = async (row) => {
  const res = await deleteUser({id: row.ID})
  if (res.code === 0) {
    ElMessage.success('删除成功')
    row.visible = false
    await getTableData()
  }
}

const userInfo = ref({
  username: '',
  password: '',
  nickName: '',
  headerImg: '',
  roleId: '',
  roleIds: [],
  enable: 1,
})

const rules = ref({
  userName: [
    {required: true, message: '请输入用户名', trigger: 'blur'},
    {min: 5, message: '最低5位字符', trigger: 'blur'}
  ],
  password: [
    {required: true, message: '请输入用户密码', trigger: 'blur'},
    {min: 6, message: '最低6位字符', trigger: 'blur'}
  ],
  nickName: [
    {required: true, message: '请输入用户昵称', trigger: 'blur'}
  ],
  phone: [
    {pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/, message: '请输入合法手机号', trigger: 'blur'},
  ],
  email: [
    {pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g, message: '请输入正确的邮箱', trigger: 'blur'},
  ],
  roleId: [
    {required: true, message: '请选择用户角色', trigger: 'blur'}
  ]
})
const userForm = ref(null)
const enterAddUserDialog = async () => {
  userInfo.value.roleId = userInfo.value.roleIds[0]
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await register(req)
        if (res.code === 0) {
          ElMessage({type: 'success', message: '创建成功'})
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setUserInfo(req)
        if (res.code === 0) {
          ElMessage({type: 'success', message: '编辑成功'})
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.roleIds = []
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  dialogFlag.value = 'add'
  addUserDialog.value = true
}

const tempAuth = {}
const changeRole = async (row, flag, removeAuth) => {
  if (flag) {
    if (!removeAuth) {
      tempAuth[row.ID] = [...row.roleIds]
    }
    return
  }
  await nextTick()
  const res = await setUserAuthorities({
    ID: row.ID,
    roleIds: row.roleIds
  })
  if (res.code === 0) {
    ElMessage({type: 'success', message: '角色设置成功'})
  } else {
    if (!removeAuth) {
      row.roleIds = [...tempAuth[row.ID]]
      delete tempAuth[row.ID]
    } else {
      row.roleIds = [removeAuth, ...row.roleIds]
    }
  }
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  addUserDialog.value = true
}

const switchEnable = async (row) => {
  userInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...userInfo.value
  }
  const res = await setUserInfo(req)
  if (res.code === 0) {
    ElMessage({type: 'success', message: `${req.enable === 2 ? '禁用' : '启用'}成功`})
    await getTableData()
    userInfo.value.headerImg = ''
    userInfo.value.roleIds = []
  }
}

</script>

<style lang="scss">
.user-dialog {
  .header-img-box {
    width: 200px;
    height: 200px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 200px;
    cursor: pointer;
  }

  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }

  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }

  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}

.nickName {
  display: flex;
  justify-content: flex-start;
  align-items: center;
}

.pointer {
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}
</style>
