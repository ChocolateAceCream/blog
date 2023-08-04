/* eslint-disable */

// replace str from index poStart to poEnd with sub
// if poStart == poEnd, then it's an insert function
export const insertAt = (str, sub, poStart, poEnd) => `${str.slice(0, poStart)}${sub}${str.slice(poEnd)}`;

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

// camel to underline
export const toSQLLine = (str) => {
    if (str === 'ID') return 'ID'
    return str.replace(/([A-Z])/g, "_$1").toLowerCase();
}

// underline to camel
export const toHump = (name) => {
    return name.replace(/\_(\w)/g, function(all, letter) {
        return letter.toUpperCase();
    });
}