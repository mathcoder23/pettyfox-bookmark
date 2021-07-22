<template>
    <div @click="menuVisible=false">
        <bookmark-edit-dialog ref="dialog" @update="updateList"/>
        <bookmark-index-dialog ref="indexDialog"/>
        <div style="display: flex;justify-content: flex-end">
            <el-button @click="clickResetIndex">重建索引</el-button>
        </div>
        <div class="search search--open">
            <form class="search__form" action="">
                <!--                <input v-model="search" class="search__input" name="search" type="search"-->
                <!--                       placeholder="Search" autocomplete="off" autofocus="autofocus" autocapitalize="off"-->
                <!--                       spellcheck="false">-->
                <el-autocomplete
                        class="search__input"
                        autofocus="autofocus"
                        v-model="search"
                        :fetch-suggestions="querySearchAsync"
                        placeholder="Search"
                        @select="handleSelect"
                ></el-autocomplete>
            </form>
        </div>
        <div style="margin-left:60px">
            <div class="list">
                <div class="item" @click="clickCreate" v-if="this.search.length === 0">
                    <div class="item-plus">
                        +
                    </div>
                </div>
                <div class="item" v-for="item in list"
                     @click="clickItem(item)"
                     @contextmenu.prevent="openMenu($event,item)">
                    <el-tooltip v-if="item.name.length >0" class="item-name" effect="dark" :content="item.name"
                                placement="right">
                        <div class="item-name"><span>{{formatText(item.name)}}</span></div>
                    </el-tooltip>

                    <el-tooltip v-if="item.desc.length >0" class="item-desc" effect="dark" :content="item.desc"
                                placement="right">
                        <div class="item-desc">{{formatText(item.desc,20)}}</div>
                    </el-tooltip>

                </div>

            </div>
        </div>

        <div id="contextmenu"
             v-if="menuVisible"
             :style="menuStyle"
             class="menu">
            <div class="contextmenu__item"
                 @click="clickEdit(currentItem)">编辑
            </div>
            <div class="contextmenu__item"
                 @click="clickRemove(currentItem)">删除
            </div>
            <div class="contextmenu__item"
                 @click="clickGetIndex(currentItem)">查看索引
            </div>
        </div>
    </div>
</template>

<script>
    import BookmarkEditDialog from "./bookmark/BookmarkEditDialog";
    import BookmarkIndexDialog from "./bookmark/BookmarkIndexDialog";

    export default {
        name: "Home",
        components: {BookmarkIndexDialog, BookmarkEditDialog},
        data() {
            return {
                menuStyle: {},
                currentItem: null,
                menuVisible: false,
                // baseUrl: 'http://api.pettyfox.top:10004',
                baseUrl: 'http://localhost:8080',
                message: 'Hello Vue!',
                search: '',
                currentSelect: {},
                form: {},
                menuList: [{
                    name: '首页'
                }, {
                    name: 'Home'
                }, {
                    name: '个人'
                }],
                list: []
            }
        },
        mounted() {
            this.updateList()
        },
        methods: {
            formatText(text, len = 15) {
                if (text) {
                    if (text.length > len) {
                        return text.substring(0, len) + "..."
                    } else {
                        return text.substring(0, text.length)
                    }
                } else {
                    return ''
                }
            },
            clickGetIndex() {
                this.$refs.indexDialog.open(this.currentItem.id)
            },
            async updateList() {
                if (this.search.length === 0) {
                    let rep = await this.$api.BookmarkApi.list()
                    this.list = rep.data
                } else {

                    let rep = await this.$api.BookmarkApi.search({
                        keyword: this.search
                    })
                    this.list = rep.data
                }

            },
            async querySearchAsync(key, cb) {
                if (this.search.length === 0) {
                    cb([])
                    return
                }
                let rep = await this.$api.BookmarkApi.searchSuggest({keyword: this.search})
                let suggest = []
                if (rep.data && rep.data.length > 0) {
                    for (let item of rep.data) {
                        suggest.push({
                            value: item
                        })
                    }
                }
                console.log('sss', suggest)
                cb(suggest)
            },
            handleSelect(item) {
                console.log(item);
            },
            clickItem(item) {
                window.open(item.url)
            },
            clickEdit(item) {
                this.$refs.dialog.open(item)
            },
            clickRemove(item) {
                this.$confirm('确认要删除此书签', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.$api.BookmarkApi.remove({ids: [item.id]}).then(rep => {
                        if (rep.code === 200) {
                            this.$message.success("删除成功")
                            this.updateList()
                        } else {
                            this.$message.error(rep.msg)
                        }
                    })
                }).catch(() => {

                });
            },
            clickResetIndex(item) {
                this.$confirm('确定要重建索引吗', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.$api.BookmarkApi.resetIndex({ids: [item.id]}).then(rep => {
                        if (rep.code === 200) {
                            this.$message.success("索引重建成功")
                            this.updateList()
                        } else {
                            this.$message.error(rep.msg)
                        }
                    })
                }).catch(() => {

                });
            },
            clickMenu(item) {
                this.currentSelect = item
            },
            openMenu(e, item) {
                console.log(e, item)
                this.menuStyle = {
                    "left": e.x + "px",
                    "top": e.y + "px"
                }
                this.currentItem = item
                this.menuVisible = true
            },
            clickCreate() {
                this.$refs.dialog.open({})
            }
        },
        watch: {
            search() {
                this.updateList()
            }
        }
    }
</script>

<style>
    body {
        /*兼容浏览器版本*/
        -webkit-background-size: cover;
        -o-background-size: cover;
        background-size: cover;
        margin: 0;
        scrollbar-width: none; /* firefox */
        -ms-overflow-style: none; /* IE 10+ */
        overflow-x: hidden;
        overflow-y: auto;
    }

    body::-webkit-scrollbar {
        display: none; /* Chrome Safari */
    }


    /* Layout for search container */
    .search {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        text-align: center;
    }


    .search__form {
        margin: 5em 0;
        width: 75%;
    }

    .search__input input {
        outline: none;
        background-color: transparent;
        border: 0;
        height: 70px;
        font-family: inherit;
        font-size: 3vw !important;
        line-height: 1;
        display: inline-block;
        box-sizing: border-box;
        width: 100%;
        max-width: 1200px;
        color: white;
        border-bottom: 4px solid #13E8E9;
    }

    .search__input input:focus {
        border: 0;
        border-bottom: 4px solid #13E8E9;
    }

    .search__input input::-webkit-input-placeholder {
        opacity: 0.1;
        /* WebKit, Blink, Edge */
        color: #fff;
    }

    .search__input input::-moz-placeholder {
        opacity: 0.1;
        /* Mozilla Firefox 19+ */
        color: #fff;
    }

    .search__input input:-ms-input-placeholder {
        opacity: 0.1;
        /* Internet Explorer 10-11 */
        color: #fff;
    }

    .search__input input::-webkit-search-cancel-button,
    .search__input input::-webkit-search-decoration {
        -webkit-appearance: none;
    }

    .search__input input::-ms-clear {
        display: none;
    }

    .search__info {
        font-size: 90%;
        font-weight: bold;
        display: block;
        color: white;
        letter-spacing: 2px;
        width: 75%;
        margin: 0 auto;
        padding: 0.85em 0;
        text-align: right;
    }

    /************************/
    /* Transitions 			*/
    /************************/


    .search--open {
        pointer-events: auto;
        opacity: 1;
    }

    /* Close button */
    .btn--search-close {
        opacity: 0;
        transform: scale3d(0.8, 0.8, 1);
        transition: opacity 0.6s, transform 0.6s;
        transition-timing-function: cubic-bezier(0.2, 1, 0.3, 1);
    }

    .search--open .btn--search-close {
        opacity: 1;
        transform: scale3d(1, 1, 1);
    }

    /* Search form with input and description */
    .search__form {
        opacity: 0;
        transform: scale3d(0.7, 0.7, 1);
        transition: opacity 0.6s, transform 0.6s;
        transition-timing-function: cubic-bezier(0.2, 1, 0.3, 1);
    }

    .search--open .search__form {
        opacity: 1;
        transform: scale3d(1, 1, 1);
    }

    @media screen and (max-width: 40em) {
        .btn--search-close {
            font-size: 1.25em;
        }
    }

    .list {
        display: flex;
        flex-flow: wrap;
        width: 80%;
        justify-content: center;
        align-content: center;
        margin: 0 auto;
    }

    .list .item {
        border-radius: 5px;
        box-shadow: 0 8px 12px rgba(0, 0, 0, .2);
        width: 250px;
        height: 100px;
        background: rgb(38 36 45);
        margin: 20px;
        color: rgb(212 212 213);
        cursor: pointer;
        user-select: none;
        display: flex;
        flex-direction: column;
    }

    .list .item .item-plus {
        font-size: 4vw;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
    }

    .list .item:hover {
        box-shadow: 0 2px 5px #13E8E9;
    }

    .list .item .item-name {
        padding: 10px;
        border-bottom: 1px solid #13E8E9;
    }

    .list .item .item-desc {
        color: grey;
        padding: 10px;
    }

    .edit {
        width: 420px;
        padding: 10px;
        background: rgba(32, 33, 36, 0.9);
        display: flex;
        justify-content: space-around;
        flex-wrap: wrap;
    }

    .edit .form {
        display: flex;
        flex-direction: column;
        width: 200px;
    }

    .edit .w2 {
        display: flex;
        flex-direction: column;
        width: 400px;
    }

    .edit .form input {
        display: block;
        box-sizing: border-box;
        width: 100%;
        padding: 8px 12px;
        border-style: solid;
        border-width: 1px;
        font-size: 14px;
        line-height: 1.3;
        border-radius: 4px;
        transition: border-color .2s ease-in-out;
        outline: none;
        margin: 0;
        font-family: inherit;
        background-color: hsla(0, 0%, 100%, .04);
        color: white;
    }

    .edit .form .title {
        color: hsla(0, 0%, 100%, .4);
        margin: 10px;
    }

    .edit-btn {
        margin-top: 5px;
        text-align: center;
        padding: 8px 16px;
        font-weight: 500;
        font-size: 14px;
        cursor: pointer;
        border: 0;
        line-height: 1.3;
        color: white;
        background-color: #1686f0;
        user-select: none;
    }


    .contextmenu__item {
        display: block;
        line-height: 34px;
        text-align: center;
    }

    .contextmenu__item:not(:last-child) {
        border-bottom: 1px solid rgba(0, 0, 0, 0.1);
    }

    .menu {
        position: fixed;
        background-color: black;
        width: 100px;
        /*height: 106px;*/
        font-size: 12px;
        color: #fff;
        border-radius: 4px;
        -webkit-box-sizing: border-box;
        box-sizing: border-box;
        border-radius: 3px;
        border: 1px solid rgba(0, 0, 0, 0.15);
        box-shadow: 0 6px 12px rgba(0, 0, 0, 0.175);
        white-space: nowrap;
        z-index: 1000;
    }

    .contextmenu__item:hover {
        cursor: pointer;
        background: hsla(0, 0%, 100%, .4);
        border-color: hsla(0, 0%, 100%, .4);
        color: #fff;
    }
</style>
