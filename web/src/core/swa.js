/*
 * gin-vue-admin web框架组
 *
 * */
import { register } from './global'

export default {
  install: (app) => {
    register(app)
    console.log(`
      > 欢迎使用熙艾迪威(XIDW)公司 go-swa 系列软件
      > 软件名：SWA微服务管理平台(Service Weaver Admin)
      > 软件架构：基本google service weaver微服务架构
      > 软件优势：架构天生微服务，单体编码,不需编写和调试微服务之间接口,自动微服务拆分部署
      > 更多优势,网络搜索 'service weaver'
      > 软件特色：第一家开源基本google service weaver的软件后台
      > 开源地址：https://gitee.com/go-swa/swa 
      > 当前版本:v2.1.8                                                       
      > 技术交流,软件服务,联系方式:
        微信号：xidwauthor 
        QQ群：547882356 
        邮箱：xidwauthor@sina.com 
      > 默认前端文件运行地址:http://127.0.0.1:8083 
      > 基于swa的数据中台即将开源,敬请期待
    `)
  }
}
