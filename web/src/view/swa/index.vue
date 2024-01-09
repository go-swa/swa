<template>
  <div id="layoutXidw">
    <div class="common-layout">
      <el-container class="layout-base">
        <el-header class="layout-top">
          <LayoutHeader />
        </el-header>

        <el-container class="layout-main">
          <el-aside class="layout-aside">
            <LayoutSider />
          </el-aside>

          <el-container class="right-main">
            <el-header class="layout-history">
              <LayoutHistory />
            </el-header>

            <Transition name="swa">
              <el-main class="box-main">
                <div class="x-page">
                  <router-view class="box-routerView" />
                </div>
              </el-main>
            </Transition>
          </el-container>
        </el-container>
      </el-container>
    </div>
  </div>
</template>

<script setup>
import LayoutHeader from '@/view/swa/xidwHeader/LayoutHeader.vue'
import LayoutSider from '@/view/swa/xidwSider/LayoutSider.vue'
import LayoutHistory from '@/view/swa/xidwHistory/LayoutHistory.vue'
</script>

<style lang="scss">
@import '@/view/swa/xidwLayout.scss';
</style>
