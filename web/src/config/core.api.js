//提供core层的，api具体实现


import {BookmarkApi} from "../api/BookmarkApi";
import {CoreApi} from "../core/api/core/CoreApi";

CoreApi.init("http://localhost:8080")
export const MyApi = {
    //实现的方法写在其中
    BookmarkApi: BookmarkApi
}
