/* eslint-disable */
export const toUpperCase = (str) => {
    if (str[0]) {
        return str.replace(str[0], str[0].toUpperCase())
    } else {
        return ''
    }
}

export const toLowerCase = (str) => {
    if (str[0]) {
        return str.replace(str[0], str[0].toLowerCase())
    } else {
        return ''
    }
}

export const toSQLLine = (str) => {
    if (str === 'ID') return 'ID'
    return str.replace(/([A-Z])/g, "_$1").toLowerCase();
}

export const toHump = (name) => {
    return name.replace(/\_(\w)/g, function (all, letter) {
        return letter.toUpperCase();
    });
}
