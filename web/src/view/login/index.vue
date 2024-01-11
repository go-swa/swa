<template>
  <div id="userLayout">
    <div class="login_panel">
      <div class="login_panel_form">
        <div class="login_panel_form_title">
          <p class="login_panel_form_title_p">{{ $XIDW_CONFIG.appName }}</p>
        </div>
        <el-form
            ref="loginForm"
            :model="loginFormData"
            :rules="rules"
            :validate-on-rule-change="false"
            label-width="auto"
            @keyup.enter="submitForm"
        >
          <el-form-item class="itemColor" label="用户名：" prop="username">
            <el-input v-model="loginFormData.username" placeholder="请输入用户名" size="large" suffix-icon="user"/>
          </el-form-item>
          <el-form-item label="密 码：" prop="password">
            <el-input v-model="loginFormData.password" placeholder="请输入密码" show-password size="large" type="password"/>
          </el-form-item>
          <el-form-item v-if="loginFormData.openCaptcha" label="验证码：" prop="captcha">
            <div class="vPicBox">
              <el-input v-model="loginFormData.captcha" placeholder="请输入验证码" size="large" style="flex:1;padding-right: 20px;"/>
              <div class="vPic">
                <img v-if="picPath" :src="picPath" alt="请输入验证码" @click="loginVerify()">
              </div>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button size="large" style="width: 100%; margin-left: 16%" type="primary" @click="submitForm">登 录</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
}
</script>

<script setup>
import {captcha} from '@/api/user'
import {checkDB} from '@/api/initdb'
import BottomInfo from '@/components/bottomInfo/bottomInfo.vue'
import {reactive, ref} from 'vue'
import {ElMessage} from 'element-plus'
import {useRouter} from 'vue-router'
import {useUserStore} from '@/pinia/modules/user'
import XIDW_CONFIG from '@/core/config'
import $XIDW_CONFIG from '@/core/config'

const router = useRouter()
const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error('请输入正确的用户名'))
  } else {
    callback()
  }
}
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('请输入正确的密码'))
  } else {
    callback()
  }
}

const loginVerify = () => {
  captcha({}).then(async (ele) => {
    rules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: 'blur',
    })
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
    loginFormData.openCaptcha = ele.data.openCaptcha
  })
}
loginVerify()

const loginForm = ref(null)
const picPath = ref('')
const loginFormData = reactive({
  username: 'admin',
  password: '123456',
  captcha: '',
  captchaId: '',
  openCaptcha: false,
})

const rules = reactive({
  username: [{validator: checkUsername, trigger: 'blur'}],
  password: [{validator: checkPassword, trigger: 'blur'}],
  captcha: [
    {
      message: '验证码格式不正确',
      trigger: 'blur',
    },
  ],
})
const userStore = useUserStore()
const login = async () => {
  return await userStore.LoginIn(loginFormData)
}

const submitForm = () => {
  loginForm.value.validate(async (v) => {
    if (v) {
      const flag = await login()
      if (!flag) {
        loginVerify()
      }
    } else {
      ElMessage({
        type: 'error',
        message: '请正确填写登录信息',
        showClose: true,
      })
      loginVerify()
      return false
    }
  })
}

const checkInit = async () => {
  const res = await checkDB()
  if (res.code === 0) {
    if (res.data?.needInit) {
      userStore.NeedInit()
      router.push({name: 'Init'})
    } else {
      ElMessage({
        type: 'info',
        message: '已配置数据库信息，无法初始化',
      })
    }
  }
}

</script>

<style lang="scss">
@import "@/styles/newLogin.scss";
</style>
