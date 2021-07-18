let app = new Vue({
    el: '#app',
    data: {
        baseUrl: 'http://localhost:3000',
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
        tab: null
    },
    async mounted() {
        this.tab = await this.getTab()
        this.$set(this.form, "url", this.tab.url)
        this.$set(this.form, "name", this.tab.title)
        console.log('tab', this.tab)
    },
    methods: {
        getTab() {
            return new Promise(((resolve, reject) => {
                chrome.tabs.getSelected(null, function (tab) {
                    resolve(tab);
                });
            }))

        },
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
                });
            } else {
                layer.msg("请输入url地址")
            }
        },
        clickCreate() {
            this.form = {}

        }
    },
    watch: {
        search() {
            this.changeSearch()
        }
    }
})