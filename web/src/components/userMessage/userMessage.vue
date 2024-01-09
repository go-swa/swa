<script>
export default {
  name: 'UserMessage'
}
</script>
<template>
  <div style="width: 40px">
    <el-icon><Bell /></el-icon>
    <span class="span_msg" v-if="msgCount<999"> {{ msgCount }}</span>
    <span class="span_msg" v-else>.....</span>
  </div>
</template>

<script setup>
import { onUnmounted, ref } from 'vue'
const msgCount = ref(1)
const timer = ref(null)

timer.value = setInterval(() => {
  msgCount.value++
}, 1000 * 1)

onUnmounted(() => {
  clearInterval(timer.value)
  timer.value = null
})
</script>

<style lang="scss" scoped>

.span_msg{
  top:-25px;
  left:22px;
  display:block;
  position:relative;
  background: red;
  border-radius: 10px;
  width:25px;
  font-size: 12px;
}
</style>
