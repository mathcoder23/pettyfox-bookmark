<template>
    <el-dialog :visible="visible"  @close="visible=false">
        <bookmark-edit @close="visible=false" ref="edit" @update="$emit('update')"/>
    </el-dialog>
</template>

<script>
    import BookmarkEdit from "./BookmarkEdit";

    export default {
        name: "BookmarkEditDialog",
        components: {BookmarkEdit},
        data() {
            return {
                visible: false,
                form: {}
            }
        },
        methods: {
            open(form) {
                this.form = form
                this.visible = true;
                this.$nextTick(() => {
                    this.$refs.edit.loadForm(this.form)
                })

            }
        }
    }
</script>

<style scoped>

</style>
