<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="字典中文名">
          <el-input v-model="searchInfo.name" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="字典英文名">
          <el-input v-model="searchInfo.type" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" clear placeholder="请选择">
            <el-option key="true" label="是" value="true"/>
            <el-option key="false" label="否" value="false"/>
          </el-select>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="searchInfo.detail" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item>
          <el-button icon="search" type="primary" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="x-table-box">
      <div class="x-btn-list">
        <el-button icon="plus" type="primary" @click="openDialog">新增</el-button>
      </div>
      <el-table
          ref="multipleTable"
          :data="tableData"
          row-key="ID"
          style="width: 100%"
          tooltip-effect="dark"
      >
        <el-table-column type="selection" width="55"/>
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
            }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="字典中文名" prop="name" width="160"/>
        <el-table-column align="left" label="字典英文名" prop="type" width="120"/>
        <el-table-column align="left" label="状态" prop="status" width="120">
          <template #default="scope">{{
            formatBoolean(scope.row.status)
            }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="描述" prop="detail"/>
        <el-table-column align="left" label="按钮组" width="210">
          <template #default="scope">
            <el-button icon="document" link type="primary" @click="toDetail(scope.row)">详情</el-button>
            <el-button icon="edit" link type="primary" @click="updateSwaDictFunc(scope.row)">变更</el-button>
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px">
                <el-button link type="primary" @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" @click="deleteSwaDictFunc(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button icon="delete" link style="margin-left: 10px" type="primary" @click="scope.row.visible = true">删除</el-button>
              </template>
            </el-popover>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form
          ref="dialogForm"
          :model="formData"
          :rules="rules"
          label-width="110px"
      >
        <el-form-item label="字典名（中）" prop="name">
          <el-input v-model="formData.name" :style="{ width: '100%' }" clearable placeholder="请输入字典名（中）"/>
        </el-form-item>
        <el-form-item label="字典名（英）" prop="type">
          <el-input v-model="formData.type" :style="{ width: '100%' }" clearable placeholder="请输入字典名（英）"/>
        </el-form-item>
        <el-form-item label="状态" prop="status" required>
          <el-switch v-model="formData.status" active-text="开启" inactive-text="停用"/>
        </el-form-item>
        <el-form-item label="描述" prop="detail">
          <el-input v-model="formData.detail" :style="{ width: '100%' }" clearable placeholder="请输入描述"/>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'SwaDict',
}
</script>

<script setup>
import {
  createSwaDict,
  deleteSwaDict,
  updateSwaDict,
  findSwaDict,
  getSwaDictList,
} from '@/api/swaDict'
import WarningBar from '@/components/warningBar/warningBar.vue'
import {ref} from 'vue'
import {useRouter} from 'vue-router'
import {ElMessage} from 'element-plus'
import {formatBoolean, formatDate} from '@/utils/format'

const router = useRouter()

const formData = ref({
  name: null,
  type: null,
  status: true,
  detail: null,
})
const rules = ref({
  name: [
    {
      required: true,
      message: '请输入字典名（中）',
      trigger: 'blur',
    },
  ],
  type: [
    {
      required: true,
      message: '请输入字典名（英）',
      trigger: 'blur',
    },
  ],
  detail: [
    {
      required: true,
      message: '请输入描述',
      trigger: 'blur',
    },
  ],
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.status === '') {
    searchInfo.value.status = null
  }
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const getTableData = async () => {
  const table = await getSwaDictList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const toDetail = (row) => {
  router.push({
    name: 'dictionaryDetail',
    params: {
      id: row.ID,
    },
  })
}

const dialogFormVisible = ref(false)
const type = ref('')
const updateSwaDictFunc = async (row) => {
  console.log('字典变更：', row)
  const res = await findSwaDict({ID: row.ID, status: row.status})
  console.log('字典变更结果：', res)
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.resysDictionary
    dialogFormVisible.value = true
  }
}
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: null,
    type: null,
    status: true,
    detail: null,
  }
}
const deleteSwaDictFunc = async (row) => {
  row.visible = false
  const res = await deleteSwaDict({ID: row.ID})
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功',
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const dialogForm = ref(null)
const enterDialog = async () => {
  dialogForm.value.validate(async (valid) => {
    if (!valid) return
    let res
    console.log('字典详情：', formData.value)
    switch (type.value) {
      case 'create':
        res = await createSwaDict(formData.value)
        break
      case 'update':
        res = await updateSwaDict(formData.value)
        break
      default:
        res = await createSwaDict(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage.success('操作成功')
      closeDialog()
      getTableData()
    }
  })
}
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}
</script>

<style></style>
