/*
 * @Author: 张侦翔 302562238@qq.com
 * @Date: 2023-02-15 11:43:04
 * @LastEditors: 张侦翔 302562238@qq.com
 * @LastEditTime: 2023-02-15 11:53:58
 * @FilePath: \web\src\plugins\cookies\index.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import Cookies from 'js-cookie'

const cookies = {}

/**
 * @description 存储 cookie 值
 * @param {String} name cookie name
 * @param {String} value cookie value
 * @param {Object} setting cookie setting
 */
cookies.set = function (name = 'default', value = '', cookieSetting = {}) {
    const currentCookieSetting = {
        expires: 1,
        sameSite: 'lax'
    }
    Object.assign(currentCookieSetting, cookieSetting)
    Cookies.set(`${name}`, value, currentCookieSetting)
}

/**
 * @description 拿到 cookie 值
 * @param {String} name cookie name
 */
cookies.get = function (name = 'default') {
    return Cookies.get(`${name}`)
}

/**
 * @description 拿到 cookie 全部的值
 */
cookies.getAll = function () {
    return Cookies.get()
}

/**
 * @description 删除 cookie
 * @param {String} name cookie name
 */
cookies.remove = function (name = 'default') {
    return Cookies.remove(`${name}`)
}

export default cookies
