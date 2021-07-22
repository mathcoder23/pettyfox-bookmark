# 小狐书签
> 轻量级的私服化个人书签

# 特性
- 私有化书签管理
- 轻量级后端
- 支持Chrome浏览器插件

# 技术架构
- 后端开发-Go(irls)
- 后端数据存储-Redis
- 全文搜索-RedisSearch
- 前端开发-Vue
> 此架构在后端初始消耗内存大概在20M
# 部署
## 插件安装
> 使用Chrome的拓展插件页，加载chrome-plug文件夹
## 后端部署
> 配备好yml文件中的redis地址
```shell script
# 后台
#docker-compose -f docker-compose.yml up -d
# 调试
#docker-compose -f docker-compose.yml up --build
# 临时
docker-compose -f docker-compose.yml up
```

## 配置
部署完成后在浏览器添加标签页会出现小狐书签管理，点击右上角的配置服务，输入后端服务器地址
