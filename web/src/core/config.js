/**
 * 网站配置文件
 */

const config = {
    appName: 'SWA微服务',
    appLogo: './assets/swa-48.png',
    showViteLogo: true
}

export const viteLogo = (env) => {
    if (config.showViteLogo) {
        const chalk = require('chalk')
        console.log(
            chalk.green(
                `> 欢迎使用熙艾迪威(XIDW)公司产品：微服务管理平台(SWA:Service Weaver Admin)，开源地址：https://gitee.com/go-swa/swa`
            )
        )
        console.log(
            chalk.green(
                `> 当前版本:v2.1.8`
            )
        )
        console.log(
            chalk.green(
                `> 加群方式:微信号：xidwauthor QQ群：547882356 邮箱：xidwauthor@sina.com`
            )
        )
        console.log(
            chalk.green(
                `> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`
            )
        )
        console.log('\n')
    }
}

export default config
