import {CoreApi} from "../core/api/core/CoreApi";

export const BookmarkApi = {
    list: () => {
        return CoreApi.apiList('/bookmark/list')
    },
    search: (query) => {
        return CoreApi.apiList('/bookmark/search',query)
    },
    resetIndex: (query) => {
        return CoreApi.apiPost('/bookmark/resetIndex',query)
    },
    getIndex: (query) => {
        return CoreApi.apiGet('/bookmark/getIndex',query)
    },
    searchSuggest: (query) => {
        return CoreApi.apiList('/bookmark/searchSuggest',query)
    },
    save: (data) => {
        return CoreApi.apiPost('/bookmark/save', data)
    },
    remove: (data) => {
        return CoreApi.apiPost('/bookmark/remove', data)
    }
}
