<template>
    <el-dialog :visible="visible" @close="visible=false">
        <div class="tag-list" v-if="data.desc">
            <sapn class="title">书签描述:</sapn>
            <div class="item" v-for="item in data.desc">
                <el-tag>{{item}}</el-tag>
            </div>
        </div>
        <div class="tag-list" v-if="data.name">
            <sapn class="title">书签标题:</sapn>
            <div class="item" v-for="item in data.name">
                <el-tag>{{item}}</el-tag>
            </div>
        </div>
        <div class="tag-list" v-if="data.url">
            <sapn class="title">书签地址:</sapn>
            <div class="item" v-for="item in data.url">
                <el-tag>{{item}}</el-tag>
            </div>
        </div>

    </el-dialog>
</template>

<script>
    export default {
        name: "BookmarkIndexDialog",
        components: {},
        data() {
            return {
                visible: false,
                form: {},
                data: {}
            }
        },
        methods: {
            batchCreate() {

                for (let i in temp.data.iconDict) {
                    let item = temp.data.iconDict[i]
                    let obj = {
                        name: item.label,
                        url: item.url,
                        desc: item.desc
                    }
                    this.$api.BookmarkApi.save(obj)
                }
            },
            async open(id) {
                this.visible = true
                let rep = await this.$api.BookmarkApi.getIndex({id: id})
                console.log('rep', rep)
                if (rep.code === 200) {
                    this.data = rep.data
                } else {
                    this.$message.warning(rep.msg)
                }
            }
        }
    }
</script>

<style scoped lang="scss">
    .tag-list {
        display: flex;
        flex-wrap: wrap;

        .title {
            margin: 10px;
        }

        .item {
            margin-left: 5px;
        }
    }
</style>
