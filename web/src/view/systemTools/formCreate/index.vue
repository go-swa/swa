<template>
  <div style="height:80vh">
    <iframe :src="`${basePath}:${basePort}/form-generator/#/`" frameborder="0" height="100%" width="100%"/>
  </div>
</template>

<script>
export default {
  name: 'FormGenerator'
}
</script>

<script setup>
import {ref} from 'vue'

const basePath = ref(import.meta.env.VITE_BASE_PATH)
const basePort = ref(import.meta.env.VITE_SERVER_PORT)
</script>
