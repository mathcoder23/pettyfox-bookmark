//提供core层的，api具体实现


import {BookmarkApi} from "../api/BookmarkApi";
import {CoreApi} from "../core/api/core/CoreApi";

let url = localStorage.getItem("baseUrl")
if (!url) {
    url = "http://localhost:10004"
}
CoreApi.init(url)
export const MyApi = {
    //实现的方法写在其中
    BookmarkApi: BookmarkApi
}
