<template>
    <div class="hello">
        <a-list item-layout="horizontal" :data-source="data">
            <a-list-item slot="renderItem" slot-scope="item, index">
                <a-list-item-meta
                        description="Ant Design, a design language for background applications, is refined by Ant UED Team"
                >
                    <a slot="title" href="https://www.antdv.com/">{{ item.title }}</a>
                    <a-avatar
                            slot="avatar"
                            src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                    />
                </a-list-item-meta>
            </a-list-item>
        </a-list>
    </div>
</template>

<script>
    const data = [
        {
            title: 'Ant Design Title 1',
        },
        {
            title: 'Ant Design Title 2',
        },
        {
            title: 'Ant Design Title 3',
        },
        {
            title: 'Ant Design Title 4',
        },
    ];
    export default {
        name: 'HelloWorld',
        data() {
            return {data}
        },
        mounted() {
            this.$api.GoodsApi.getList()
        },
        props: {
            msg: String
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
