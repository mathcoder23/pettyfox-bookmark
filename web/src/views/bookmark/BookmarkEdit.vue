<template>
    <el-form ref="form" :model="form" label-width="80px">
        <el-form-item label="标题">
            <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="书签分组">
            <el-select v-model="form.region" disabled placeholder="请选择标签分组">
                <el-option label="分组1" value="shanghai"></el-option>
            </el-select>
        </el-form-item>
        <el-form-item label="书签地址">
            <el-input v-model="form.url"></el-input>
        </el-form-item>
        <el-form-item label="描述">
            <el-input type="textarea" v-model="form.desc"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button @click="handleSubmit">保存</el-button>
            <el-button @click="$emit('close')">取消</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
    import {StrUtil} from "../../utils/StrUtil";

    export default {
        name: "BookmarkEdit",
        data() {
            return {
                form: {}
            }
        },
        methods: {
            loadForm(form) {
                this.form = Object.assign({}, this.form, form)
            },
            handleSubmit() {

                if (!StrUtil.isNotBlank(this.form.url)) {
                    this.$message("请输入书签地址")
                    return
                }
                this.$api.BookmarkApi.save(this.form).then(rep => {
                    if (rep.code === 200) {
                        this.$emit('close')
                        this.$emit('update')
                    } else {
                        this.$message.warning(rep.msg)
                    }
                })
            }
        }
    }
</script>

<style scoped>

</style>
