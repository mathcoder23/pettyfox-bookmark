export const StrUtil = {
    format: function (origin = '') {
        if (arguments.length <= 1) return origin
        let s = arguments[0]

        for (let i = 1; i < arguments.length; i++) {
            s = s.replace(new RegExp('\\{' + (i - 1) + '\\}', 'g'), arguments[i])
        }
        return s
    },
    /**
     * 处理前端空格提交数据的问题,以及删除空字符串的属性
     * @param obj
     * @returns {string|*}
     */
    trimSubmit: (obj) => {
        if (typeof obj === 'string') {
            return obj.trim()
        } else if (typeof obj === 'object') {
            const fields = Object.keys(obj)
            for (const field of fields) {
                if (field.indexOf("_") === 0 || field.indexOf("$") === 0) {
                    delete obj[field]
                }
            }
            for (const index in obj) {
                if (typeof obj[index] === 'undefined') {
                    delete obj[index]
                } else if (typeof obj[index] === 'string') {
                    obj[index] = obj[index].trim()
                    if (obj[index].length === 0 || obj[index] === 'undefined') {
                        delete obj[index]
                    }
                }
            }
        }
        return obj
    },
    /**
     * 对象数组，自定义转换规则提取字符串
     * @param stream
     * @param eachCb
     * @returns {string}
     */
    stream2mapStr(stream, eachCb) {
        let str = ''
        if (stream instanceof Array) {
            for (const each of stream) {
                str += eachCb(each)
            }
            str = str.substr(0, str.length - 1)
        }
        return str
    },
    getArrayObjectItem(array, key, value) {
        if (array) {
            for (const index in array) {
                if (array[index][key] === value) {
                    return array[index]
                }
            }
        }
        return {}
    },
    isNotBlank(str) {
        return typeof str === 'string' && str.trim().length > 0
    }

}
