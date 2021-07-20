
let app = new Vue({
    el: '#app',
    data: {
        baseUrl: 'http://api.pettyfox.top:10004',
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
        list: [
            {name: 'aaa', desc: 'baaa'},
            {name: 'aaa', desc: 'baaa'},
            {name: 'aaa', desc: 'baaa'},
            {name: 'aaa', desc: 'baaa'},
            {name: 'aaa', desc: 'baaa'},
            {name: 'aaa', desc: 'baaa'},
            {name: 'aaa', desc: 'baaa'},
            {name: 'aaa', desc: 'baaa'},
            {name: 'aaa', desc: 'baaa'},
            {name: 'bb', desc: 'baaa'}
        ]
    },
    mounted() {
        this.updateList()
    },
    methods: {
        updateList() {
            $.get(this.baseUrl + '/bookmark/list', (result) => {
                this.list = result.data
                console.log('aa', this.list)
            });
        },
        clickItem(item) {
            window.open(item.url)
        },
        clickEdit(item) {
            this.form = Object.assign({}, this.form, JSON.parse(JSON.stringify(item)))
            layer.open({
                type: 1,
                title: false,
                closeBtn: 0,
                shadeClose: true,
                skin: 'yourclass',
                content: $('.edit')
            });
        },
        changeSearch() {
            if (this.search.length === 0) {
                this.updateList()
            } else {
                $.get(this.baseUrl + '/bookmark/search?keyword=' + this.search, (result) => {
                    this.list = result.data
                    console.log('aa', this.list)
                });
            }

        },
        clickMenu(item) {
            this.currentSelect = item
        },
        clickSave() {
            console.log('aa', this.form)
            if (this.form.url && this.form.url.length > 0) {
                $.post(this.baseUrl + '/bookmark/add', this.form, (result) => {
                    console.log('aaa', result)
                    layer.closeAll()
                    this.updateList()
                });
            } else {
                layer.msg("请输入url地址")
            }
        },
        // batchCreate() {
        //
        //     for (let key in temp.data.iconDict) {
        //         let item = temp.data.iconDict[key]
        //         let obj = {
        //             name: item.label,
        //             url: item.url,
        //             desc: item.desc
        //         }
        //         $.post(this.baseUrl + '/bookmark/add', obj, (result) => {
        //
        //         });
        //         console.log('obj', obj)
        //     }
        // },
        clickCreate() {
            this.form = {}
            layer.open({
                type: 1,
                title: false,
                closeBtn: 0,
                shadeClose: true,
                skin: 'yourclass',
                content: $('.edit')
            });
        }
    },
    watch: {
        search() {
            this.changeSearch()
        }
    }
})
