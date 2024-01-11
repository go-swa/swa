export default function Position() {
    return {
        name: 'position',
        apply: 'serve',
        transform(code, id) {
            const index = id.lastIndexOf('.')
            const ext = id.substr(index + 1)
            if (ext.toLowerCase() === 'vue') {
                return codeLineTrack(code, id)
            }
        },
    }
}

const codeLineTrack = (code, id) => {
    const lineList = code.split('\n')
    const newList = []
    lineList.forEach((item, index) => {
        newList.push(addLineAttr(item, index + 1, id))
    })
    return newList.join('\n')
}

const addLineAttr = (lineStr, line, id) => {
    if (!/^\s+</.test(lineStr)) {
        return lineStr
    }

    const reg = /((((^(\s)+\<))|(^\<))[\w-]+)|(<\/template)/g
    let leftTagList = lineStr.match(reg)
    if (leftTagList) {
        leftTagList = Array.from(new Set(leftTagList))
        leftTagList.forEach((item) => {
            const skip = [
                'KeepAlive',
                'template',
                'keep-alive',
                'transition',
                'el-',
                'El',
                'router-view',
            ]
            if (item && !skip.some((i) => item.indexOf(i) > -1)) {
                const reg = new RegExp(`${item}`)
                const location = `${item} code-location="${id}:${line}"`
                lineStr = lineStr.replace(reg, location)
            }
        })
    }
    return lineStr
}
