<template>
  <div>
    <div class="clearfix sticky-button" style="margin: 18px">
      <el-button class="fl-right" type="primary" @click="authDataEnter">确 定</el-button>
      <el-button class="fl-left" type="primary" @click="all">全选</el-button>
      <el-button class="fl-left" type="primary" @click="self">本角色</el-button>
      <el-button class="fl-left" type="primary" @click="selfAndChildren">本角色及子角色</el-button>
    </div>
    <div class="tree-content">
      <el-checkbox-group v-model="subRoles" @change="selectRole">
        <el-checkbox v-for="(item,key) in roles" :key="key" :label="item">{{ item.roleName }}</el-checkbox>
      </el-checkbox-group>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Datas'
}
</script>

<script setup>
import { getSubRoles, setSubRoles } from '@/api/swaRole'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
const props = defineProps({
  row: {
    default: function() {
      return {}
    },
    type: Object
  },
  role: {
    default: function() {
      return []
    },
    type: Array
  }
})

const roles = ref([])
const needConfirm = ref(false)
const roundRole = (roleData) => {
  roleData && roleData.forEach(item => {
    const obj = {}
    obj.roleId = item.roleId
    obj.roleName = item.roleName
    roles.value.push(obj)
    if (item.children && item.children.length) {
      roundRole(item.children)
    }
  })
}

const subRoles = ref([])
const init = async() => {
  roundRole(props.role)
  console.log('配置资源', props.role)
  console.log('配置资源row', props.row.roleId)
  const res = await getSubRoles(props.row.roleId)
  console.log('获取子角色：', res)
  if (res.code === 0) {
    subRoles.value = res.data.subRoles
    console.log('子角色：', subRoles.value)
  }
}

init()

const enterAndNext = () => {
  authDataEnter()
}

const emit = defineEmits(['changeRow'])
const all = () => {
  subRoles.value = [...roles.value]
  emit('changeRow', 'subRoles', subRoles.value)
  needConfirm.value = true
}
const self = () => {
  subRoles.value = roles.value.filter(item => item.roleId === props.row.roleId)
  emit('changeRow', 'subRoles', subRoles.value)
  needConfirm.value = true
}
const selfAndChildren = () => {
  const arrBox = []
  getChildrenId(props.row, arrBox)
  subRoles.value = roles.value.filter(item => arrBox.indexOf(item.roleId) > -1)
  emit('changeRow', 'subRoles', subRoles.value)
  needConfirm.value = true
}
const getChildrenId = (row, arrBox) => {
  arrBox.push(row.roleId)
  row.children && row.children.forEach(item => {
    getChildrenId(item, arrBox)
  })
}
const authDataEnter = async() => {
  const res = await setSubRoles(props.row)
  console.log('设置用户资源权限：', props.row)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '资源设置成功' })
  }
}

const selectRole = () => {
  emit('changeRow', 'subRoles', subRoles.value)
  needConfirm.value = true
}

defineExpose({
  enterAndNext,
  needConfirm
})

</script>
