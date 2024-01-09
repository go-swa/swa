<template>
  <div>
    <div style="border: 1px solid #538ea4">
      <span class="div_span">服务器状态</span>
      <el-row :gutter="5" class="system_state">
        <el-col :span="6">
          <el-card v-if="state.os" class="card_item">
            <template #header>
              <div>操作系统</div>
            </template>
            <div style="margin-top: -20px">
              <el-row v-for="(v, k) in state.os" :key="k" :gutter="5" style="height: 30px">
                <el-col v-if="stateDict[k]" :span="12">{{stateDict[k]}}</el-col>
                <el-col v-else :span="12">{{k}}</el-col>
                <el-col :span="12" v-text="v" />
              </el-row>
            </div>
          </el-card>
        </el-col>

        <el-col :span="6">
          <el-card
            v-if="state.cpu"
            class="card_item"
            :body-style="{ height: '180px', 'overflow-y': 'scroll' }"
          >
            <template #header>
              <div>处理器</div>
            </template>
            <div style="margin-top: -20px">
              <el-row :gutter="10">
                <el-col :span="12">物理内核数</el-col>
                <el-col :span="12" v-text="state.cpu.cores" />
              </el-row>
              <el-row v-for="(item, index) in state.cpu.cpus" :key="index" :gutter="10" style="height: 30px">
                <el-col :span="12">内核{{ index }}:</el-col>
                <el-col
                  :span="12"
                >
                  <el-progress
                    type="line"
                    :percentage="+item.toFixed(0)"
                    :color="colors"
                  />
                </el-col>
              </el-row>
            </div>
          </el-card>
        </el-col>

        <el-col :span="6">
          <el-card v-if="state.disk" class="card_item">
            <template #header>
              <div>硬盘</div>
            </template>
            <div style="margin-top: -30px">
              <el-row :gutter="10">
                <el-col :span="14">
                  <el-row v-for="(v, k) in state.disk" :key="k" :gutter="70" style="height: 30px">
                    <el-col v-if="stateDict[k]" :span="14">{{ stateDict[k]}}</el-col>
                    <el-col v-else :span="14">{{ k }}</el-col>
                    <el-col :span="110" v-text="v" />
                  </el-row>
                </el-col>
                <el-col :span="10">
                  <el-progress
                    type="dashboard"
                    :percentage="state.disk.usedPercent"
                    :color="colors"
                  />
                </el-col>
              </el-row>
            </div>
          </el-card>
        </el-col>

        <el-col :span="6">
          <el-card v-if="state.ram" class="card_item">
            <template #header>
              <div>内存</div>
            </template>
            <div style="margin-top: -30px">
              <el-row :gutter="10">
                <el-col :span="16">
                  <el-row v-for="(v, k) in state.ram" :key="k" :gutter="10" style="height: 30px">
                    <el-col v-if="stateDict[k]" :span="12">{{ stateDict[k]}}</el-col>
                    <el-col v-else :span="12">{{ k }}</el-col>
                    <el-col :span="12" v-text="v" />
                  </el-row>
                </el-col>
                <el-col :span="8">
                  <el-progress
                    type="dashboard"
                    :percentage="state.ram.usedPercent"
                    :color="colors"
                  />
                </el-col>
              </el-row>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script setup>
import { getSystemState } from '@/api/system'
import { onUnmounted, ref } from 'vue'

const timer = ref(null)
const state = ref({})
const stateDict = ref({})
const colors = ref([
  { color: '#5cb87a', percentage: 20 },
  { color: '#e6a23c', percentage: 40 },
  { color: '#f56c6c', percentage: 80 }
])

const reload = async() => {
  const res = await getSystemState()
  console.log('getSystemState:', res)
  if (res.code === 0) {
    state.value = res.data.server
    stateDict.value = res.data.stateDict
    console.log('serve state:', res.data)
  }
}


reload()

timer.value = setInterval(() => {
  reload()
}, 1000 * 10)

onUnmounted(() => {
  clearInterval(timer.value)
  timer.value = null
})

</script>

<script>
export default {
  name: 'State',
}
</script>

<style>
.system_state {
  padding: 10px;
}

.card_item {
  margin-top: -7px;
  height: 200px;
}
</style>
