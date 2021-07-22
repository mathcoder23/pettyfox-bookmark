//提供core层的，api具体实现


import {BookmarkApi} from "../api/BookmarkApi";
import {CoreApi} from "../core/api/core/CoreApi";

CoreApi.init("http://api.pettyfox.top:10004")
export const MyApi = {
    //实现的方法写在其中
    BookmarkApi: BookmarkApi
}
