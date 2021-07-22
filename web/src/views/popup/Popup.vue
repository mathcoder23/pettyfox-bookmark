<template>
    <div>
        <bookmark-edit @close="onClose" ref="edit" @update="$emit('update')"/>
    </div>
</template>

<script>
    import BookmarkEdit from "../bookmark/BookmarkEdit";

    export default {
        name: "Popup",
        components: {BookmarkEdit},
        data() {
            return {
                tab: null
            }
        },
        async mounted() {
            this.tab = await this.getChromeTab()
            this.$nextTick(() => {

                this.$refs.edit.loadForm({
                    name: this.tab.title,
                    url: this.tab.url
                })
            })
        }, methods: {
            onClose() {
                window.close()
            },
            getChromeTab() {
                return new Promise(((resolve, reject) => {
                    chrome.tabs.getSelected(null, function (tab) {
                        resolve(tab);
                    });
                }))

            },
        }
    }
</script>

<style>
    body {
        width: 700px;
    }
</style>
