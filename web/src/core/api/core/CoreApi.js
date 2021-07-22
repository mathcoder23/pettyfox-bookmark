import ax from 'axios'
import {StrUtil} from "../../utils/StrUtil";

let request = null
let axios = null
let baseUrl = ''
let qsString = (params, init = 0) => {
    let qs = ""
    if ("object" === typeof params) {
        qs = init === 0 ? "?" : ''
        let count = 0
        for (let p in params) {
            if (typeof params[p] === "number" ||
                typeof params[p] === "boolean" ||
                typeof params[p] === "string") {
                if (count === 0) {
                    count++
                    qs += p + "=" + params[p]
                } else {
                    qs += "&" + p + "=" + params[p]
                }
            }
        }
    }
    return qs
}

function delCookie() {
    var keys = document.cookie.match(/[^ =;]+(?==)/g)
    if (keys) {
        for (var i = keys.length; i--;) {
            document.cookie = keys[i] + '=0;path=/;expires=' + new Date(0).toUTCString() // 清除当前域名下的,例如：m.ratingdog.cn
            document.cookie = keys[i] + '=0;path=/;domain=' + document.domain + ';expires=' + new Date(0).toUTCString() // 清除当前域名下的，例如 .m.ratingdog.cn
            document.cookie = keys[i] + '=0;path=/;domain=ratingdog.cn;expires=' + new Date(0).toUTCString() // 清除一级域名下的或指定的，例如 .ratingdog.cn
        }
    }
}

let init = (url) => {
    baseUrl = url
    axios = ax
    if (axios == null) {
        console.error('请配置axios')
    }
    // axios.defaults.retry = 4;
    // axios.defaults.retryDelay = 1000;
    request = axios.create({
        baseURL: baseUrl,
        timeout: 30000,
        // responseType: 'json',
        // validateStatus:()=>{return true}
    });


    //HTTPrequest拦截
    request.interceptors.request.use(config => {
        console.log('request', config)

        return config
    }, error => {
        return Promise.reject(error)
    });

    //响应拦截
    request.interceptors.response.use(res => {
        console.log('response', res)

        return res.data;
    }, error => {

        return Promise.reject(new Error(error));
    })


}

const apiPost = (config, data, params) => {
    return new Promise((resolve, reject) => {
        let conf = {method: "post"}
        if (typeof config === "string") {
            conf.url = config
        } else {
            conf.url = config.uri
        }
        // if(typeof params !== "undefined" && params !== null){
        //     conf.url = conf.url+qsString(params)
        //     conf.params = params
        // }
        conf.params = params
        conf.data = data
        conf.meta = Object.assign({flag: 'apiPost'}, config.meta)

        request(conf).then(rep => resolve(rep)).catch(err => {
            reject(err)
        })
    })

}
const apiRequest = (config) => {
    if (config.data) {
        config.data = StrUtil.trimSubmit(config.data)
    }
    if (config.params) {
        config.params = StrUtil.trimSubmit(config.params)
    }
    return request(config)
}
const apiFormPost = (config, data, params) => {
    return new Promise((resolve, reject) => {
        let conf = {method: "post"}
        if (typeof config === "string") {
            conf.url = config
        } else {
            conf.url = config.uri
        }
        // if(typeof params !== "undefined" && params !== null){
        //     conf.url = conf.url+qsString(params)
        //     conf.params = params
        // }
        conf.params = params
        conf.data = qsString(data, 1)
        conf.meta = Object.assign({flag: 'apiPost'}, config.meta)
        conf.headers = {'Content-Type': 'application/x-www-form-urlencoded'}
        request(conf).then(rep => resolve(rep)).catch(err => {
            reject(err)
        })
    })

}
const apiPut = (config, params, data) => {
    return new Promise((resolve, reject) => {
        let conf = {method: "put"}
        if (typeof config === "string") {
            conf.url = config
        } else {
            conf.url = config.uri
        }
        conf.url = conf.url + qsString(params)
        conf.data = data
        conf.params = params
        request(conf).then(rep => resolve(rep)).catch(err => {
            reject(err)
        })
    })

}
const apiDel = (config, params, data) => {
    return new Promise((resolve, reject) => {
        let conf = {method: "delete"}
        if (typeof config === "string") {
            conf.url = config
        } else {
            conf.url = config.uri
        }
        conf.url = conf.url + qsString(params)
        conf.data = data
        conf.params = params
        request(conf).then(rep => resolve(rep)).catch(err => {
            reject(err)
        })
    })
}
const apiGet = (config, params) => {
    return new Promise((resolve, reject) => {
        let conf = {method: "get"}
        if (typeof config === "string") {
            conf.url = config
        } else {
            conf.url = config.uri
        }
        // conf.url = conf.url+qsString(params)
        conf.params = params
        request(conf).then(rep => resolve(rep)).catch(err => {
            reject(err)
        })
    })

}
const apiList = (config, params) => {
    return new Promise((resolve, reject) => {
        let conf = {method: "get"}
        if (typeof config === "string") {
            conf.url = config
        } else {
            conf.url = config.uri
        }
        // conf.url = conf.url+qsString(params)
        conf.params = params
        request(conf).then(rep => resolve(rep)).catch(err => {
            reject(err)
        })
    })

}

/**
 * 目前暂定api接口返回的都是数据，逻辑分离的。即resolve返回数据，当状态码异常在reject中
 * @type {{init: init, request: null, driver: null, intercept: null, notify: null, url: null, apiList: (function(*=, *=): Promise<any>), apiGet: (function(*=, *=): Promise<any>), apiPost: (function(*=, *=, *): Promise<any>), apiPut: (function(*=, *=, *): Promise<any>), apiDel: (function(*=, *=, *): Promise<any>)}}
 */
export const CoreApi = {
    init:init,
    apiList: apiList,
    apiGet: apiGet,
    apiFormPost: apiFormPost,
    apiPost: apiPost,
    apiRequest: apiRequest,
    apiPut: apiPut,
    apiDel: apiDel
}
