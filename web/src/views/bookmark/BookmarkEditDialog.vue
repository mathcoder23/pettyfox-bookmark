<template>
    <el-dialog :visible="visible">
        <bookmark-edit @close="visible=false" ref="edit" @update="$emit('update')"/>
    </el-dialog>
</template>

<script>
    import BookmarkEdit from "./BookmarkEdit";

    let temp = {
        "data": {
            "groupDict": {
                "294d521e-f722-4874-b4db-ad887f438305": {
                    "data": [
                        "1a83be58-c1a6-4e5c-ab78-9d9f4d0c3918"
                    ],
                    "icoSrc": "https://static.monknow.com/newtab/icon-group/61df5dde-0afd-4104-9b87-418bd62f95df.svg",
                    "label": "生活",
                    "uuid": "294d521e-f722-4874-b4db-ad887f438305"
                },
                "3a67f870-a7b1-4be2-a9c3-f7c20d259bc1": {
                    "data": [
                        "6b433195-b51c-4281-9673-3b48bfd00cd4",
                        "9bcfc82c-1c7d-4c91-a7c5-04dbf4458780"
                    ],
                    "icoSrc": "https://static.monknow.com/newtab/icon-group/15cb3663-ba77-44f6-91ab-2ded15fdec30.svg",
                    "label": "开发",
                    "uuid": "3a67f870-a7b1-4be2-a9c3-f7c20d259bc1"
                },
                "55df8225-f86b-4741-885e-a09ad501c325": {
                    "data": [
                        "a1d19640-f094-4eb7-b199-13ab518a7e10",
                        "3fd9fdc0-a711-4800-904e-f4b76b842472",
                        "af67c476-adae-470a-9183-73f0de29603d",
                        "9f6fba14-a054-43f9-8f15-fa35269e0d01",
                        "c57db4ac-7dd5-411d-a8e7-d2a2764c363a",
                        "ad584ea0-8886-40d3-96c7-e7777a538cbf",
                        "26ae3273-71d1-4b8c-831b-81f1250eb550"
                    ],
                    "icoSrc": "https://static.monknow.com/newtab/icon-group/61df5dde-0afd-4104-9b87-418bd62f95df.svg",
                    "label": "Home",
                    "uuid": "55df8225-f86b-4741-885e-a09ad501c325"
                },
                "8cf6b7e0-ff1b-49b5-8eb7-84346e308627": {
                    "data": [
                        "82b47318-5c66-4381-8124-41102636c76f"
                    ],
                    "icoSrc": "https://static.monknow.com/newtab/icon-group/61df5dde-0afd-4104-9b87-418bd62f95df.svg",
                    "label": "技术",
                    "uuid": "8cf6b7e0-ff1b-49b5-8eb7-84346e308627"
                },
                "a86403c1-660c-436b-88f2-f46d23513711": {
                    "data": [
                        "7f4620cd-179f-4a6a-aa58-42ff0ab3b7bc",
                        "6adced17-78ea-47b0-8b21-09a8c8b71db3"
                    ],
                    "icoSrc": "https://static.monknow.com/newtab/icon-group/9ac2df78-617e-4090-85ea-e68caf566a47.svg",
                    "label": "电影",
                    "uuid": "a86403c1-660c-436b-88f2-f46d23513711"
                },
                "e071f0b3-6e2e-49b6-92df-742c40854370": {
                    "data": [
                        "1a341d3f-b135-4fdb-95d2-a164821a4d73",
                        "29026e99-dd36-4be6-9ad2-0f45d3b61257",
                        "3b3a4f0b-df87-4c14-b1ce-9c4318a534df",
                        "755b16f1-bcb4-47a4-88fe-c1cfe6660a20",
                        "c1916a7d-dee2-40ad-ae5f-b970d13a6672",
                        "a1f4dd0a-adfe-4626-829d-bad8bceb7130",
                        "51b401be-3b72-4575-8163-259bbf6dcfd7",
                        "ef15afed-eab8-4729-a0dd-422b799a44eb",
                        "175db664-3049-4b81-a29e-c4d941f5a303",
                        "57f92ef1-15b9-4650-ac88-e7b5c80b7749",
                        "3540416b-df35-484d-a3df-e7749e2627b1",
                        "f0b16b35-5f9f-4d0f-996f-c81343f640de",
                        "174487ba-5de5-4ea6-aba7-e53d34b97f40",
                        "d6bbe31e-dbde-4b0c-8a72-0c269ad1a6e6",
                        "4cc236f3-e451-4240-b04f-879b052c2bcf",
                        "96774e24-7c99-4ca4-9d4a-f7d67b6bab8c",
                        "160a86aa-0767-4838-8fc0-cf940a167490",
                        "176066af-1610-4c73-8eef-c898a117502f",
                        "ce5d9f93-deaa-45f0-ad42-b77111eb5a88",
                        "b76c8db3-eb00-4692-9d41-1aa5ac609392",
                        "c79b3f4a-806b-4fe7-8788-6cda2e43e132"
                    ],
                    "icoSrc": "https://static.monknow.com/newtab/icon-group/61df5dde-0afd-4104-9b87-418bd62f95df.svg",
                    "label": "知识",
                    "uuid": "e071f0b3-6e2e-49b6-92df-742c40854370"
                },
                "ee571baf-8db4-4cb9-b544-97a997aecd63": {
                    "data": [
                        "ef20b9c3-b801-4010-a08f-b0c18379352e",
                        "c421c45a-cabc-440c-9786-7014d0e12f81",
                        "74600bc7-2d83-4e5f-8393-d803592a2710",
                        "c1818ed1-e08b-4b54-9636-75e20d6e42c6",
                        "9a463dc2-4ccc-44be-bbc6-d9aa121b5041",
                        "6efff3d9-7476-4986-ac4c-643a96dd9193"
                    ],
                    "icoSrc": "https://static.monknow.com/newtab/icon-group/61df5dde-0afd-4104-9b87-418bd62f95df.svg",
                    "label": "公司",
                    "uuid": "ee571baf-8db4-4cb9-b544-97a997aecd63"
                },
                "fb339dbd-2af4-4434-a0e8-5b234795d964": {
                    "data": [
                        "570a3519-5a58-459f-87de-41f3a73dcd8f",
                        "8b23b137-a9c0-44d1-ae86-afed1ee074c7",
                        "df68b8d9-3516-4ae3-b8e2-8d703693b3ec",
                        "77279812-abf7-4ed3-b2aa-0636a5dcbf7f",
                        "68809a49-85ef-4d07-a649-5b2e2cc18ae0",
                        "a7d9599b-0764-4e5f-af5e-218ad03bca22",
                        "df00e31d-d41b-4bd1-a1d8-df09cec78b23",
                        "fdcd26fb-4798-4eb2-a0ef-06d944e9b2c0"
                    ],
                    "icoSrc": "https://static.monknow.com/newtab/icon-group/5664820e-0e6e-4040-9fa8-7a7491a5f2b8.svg",
                    "label": "Tool",
                    "uuid": "fb339dbd-2af4-4434-a0e8-5b234795d964"
                },
                "fd7fd429-2a78-4eba-98af-04d59619cfb4": {
                    "data": [
                        "70158d50-85ba-411a-b465-d3a8af22d781",
                        "da723fcc-af81-40f2-9362-54402ba551c9",
                        "da44211c-390d-4c1c-88c0-92c6cfd9e518",
                        "344c8022-678f-4650-bcf5-838452b8edd5",
                        "9caff45f-2a17-41f2-8985-56a11ecb24b5"
                    ],
                    "icoSrc": "https://static.monknow.com/newtab/icon-group/5664820e-0e6e-4040-9fa8-7a7491a5f2b8.svg",
                    "label": "常用工具",
                    "uuid": "fd7fd429-2a78-4eba-98af-04d59619cfb4"
                }
            },
            "groups": [
                "fd7fd429-2a78-4eba-98af-04d59619cfb4",
                "ee571baf-8db4-4cb9-b544-97a997aecd63",
                "3a67f870-a7b1-4be2-a9c3-f7c20d259bc1",
                "e071f0b3-6e2e-49b6-92df-742c40854370",
                "55df8225-f86b-4741-885e-a09ad501c325",
                "a86403c1-660c-436b-88f2-f46d23513711",
                "fb339dbd-2af4-4434-a0e8-5b234795d964",
                "294d521e-f722-4874-b4db-ad887f438305",
                "8cf6b7e0-ff1b-49b5-8eb7-84346e308627"
            ],
            "iconDict": {
                "160a86aa-0767-4838-8fc0-cf940a167490": {
                    "backgroundColor": {
                        "data": "#9326e9",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "管理",
                    "id": null,
                    "label": "记忆管理平台",
                    "type": "text",
                    "url": "https://www.pettyfox.top/admin/index.html#/wel/index",
                    "uuid": "160a86aa-0767-4838-8fc0-cf940a167490"
                },
                "174487ba-5de5-4ea6-aba7-e53d34b97f40": {
                    "backgroundColor": {
                        "data": "#97ffd8",
                        "type": "pure"
                    },
                    "desc": "随心写作，自由表达",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/37b581a34bd7b8315c2b2c9cb125feeb.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "2021年中高级最全Java面试题，肝了！ - 知乎",
                    "type": "image",
                    "url": "https://zhuanlan.zhihu.com/p/346994550",
                    "uuid": "174487ba-5de5-4ea6-aba7-e53d34b97f40"
                },
                "175db664-3049-4b81-a29e-c4d941f5a303": {
                    "backgroundColor": {
                        "data": "#9326e9",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "概述",
                    "id": null,
                    "label": "概述 · Casbin",
                    "type": "text",
                    "url": "https://casbin.org/docs/zh-CN/overview",
                    "uuid": "175db664-3049-4b81-a29e-c4d941f5a303"
                },
                "176066af-1610-4c73-8eef-c898a117502f": {
                    "backgroundColor": {
                        "data": "#ee3b3b",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "中国",
                    "id": null,
                    "label": "中国亲戚关系计算器",
                    "type": "text",
                    "url": "http://passer-by.com/relationship/",
                    "uuid": "176066af-1610-4c73-8eef-c898a117502f"
                },
                "1a341d3f-b135-4fdb-95d2-a164821a4d73": {
                    "backgroundColor": {
                        "data": "#373737",
                        "type": "pure"
                    },
                    "desc": "全球极客挚爱的技术成长平台",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/de4034e7d1785ffcd072fdcdcae3a3df.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "140. 单词拆分 II - 力扣（LeetCode）",
                    "type": "image",
                    "url": "https://leetcode-cn.com/problems/word-break-ii/",
                    "uuid": "1a341d3f-b135-4fdb-95d2-a164821a4d73"
                },
                "1a83be58-c1a6-4e5c-ab78-9d9f4d0c3918": {
                    "backgroundColor": {
                        "data": "#97ffd8",
                        "type": "pure"
                    },
                    "desc": "随心写作，自由表达",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/37b581a34bd7b8315c2b2c9cb125feeb.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "2021 年夏天有哪些值得购买的硬防晒（防晒衣、防晒帽、冰袖、遮阳伞等）用品 - 知乎",
                    "type": "image",
                    "url": "https://zhuanlan.zhihu.com/p/158869059",
                    "uuid": "1a83be58-c1a6-4e5c-ab78-9d9f4d0c3918"
                },
                "26ae3273-71d1-4b8c-831b-81f1250eb550": {
                    "backgroundColor": {
                        "data": "#cc0000",
                        "type": "pure"
                    },
                    "desc": "原创IT技术文章分享及交流平台",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/578b373fa7a0b44814f4d43d29ea51cf.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "IntelliJ IDEA 2021年最新激活码免费获取（IDEA永久激活注册码）_wx5f2a5b33cc3f2的技术博客_51CTO博客",
                    "type": "image",
                    "url": "https://blog.51cto.com/u_14896618/2819585",
                    "uuid": "26ae3273-71d1-4b8c-831b-81f1250eb550"
                },
                "29026e99-dd36-4be6-9ad2-0f45d3b61257": {
                    "backgroundColor": {
                        "data": "#f5363f",
                        "type": "pure"
                    },
                    "desc": "消费门户网站，实时推送优质的网购优惠信息，真实的原创购物攻略",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/93b4c585a7f6c393af0b897388185ec8.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "什么值得买 | 分享每一种值",
                    "type": "image",
                    "url": "https://www.smzdm.com/",
                    "uuid": "29026e99-dd36-4be6-9ad2-0f45d3b61257"
                },
                "344c8022-678f-4650-bcf5-838452b8edd5": {
                    "backgroundColor": {
                        "data": "#333333",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "V2",
                    "id": null,
                    "label": "V2ray Linux客户端v2rayA下载安装及使用教程 支持VMess/VLESS/SS/SSR/Trojan - V2ray XTLS黑科技",
                    "type": "text",
                    "url": "https://v2xtls.org/v2ray-linux%e5%ae%a2%e6%88%b7%e7%ab%afv2raya%e4%b8%8b%e8%bd%bd%e5%ae%89%e8%a3%85%e5%8f%8a%e4%bd%bf%e7%94%a8%e6%95%99%e7%a8%8b-%e6%94%af%e6%8c%81vmess-vless-ss-ssr-trojan-pingtunnel/",
                    "uuid": "344c8022-678f-4650-bcf5-838452b8edd5"
                },
                "3540416b-df35-484d-a3df-e7749e2627b1": {
                    "backgroundColor": {
                        "data": "#007ce7",
                        "type": "pure"
                    },
                    "desc": "有问题，上知乎。可信赖的问答社区。",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/d485e51391e9a96b80caa34516bd00be.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "Java 中new String(字面量) 中 字面量 是何时进入字符串常量池的? - 知乎",
                    "type": "image",
                    "url": "https://www.zhihu.com/question/55994121",
                    "uuid": "3540416b-df35-484d-a3df-e7749e2627b1"
                },
                "3b3a4f0b-df87-4c14-b1ce-9c4318a534df": {
                    "backgroundColor": {
                        "data": "#16d9c4",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "By",
                    "id": null,
                    "label": "ByteDance IconPark",
                    "type": "text",
                    "url": "https://iconpark.oceanengine.com/official",
                    "uuid": "3b3a4f0b-df87-4c14-b1ce-9c4318a534df"
                },
                "3fd9fdc0-a711-4800-904e-f4b76b842472": {
                    "backgroundColor": {
                        "data": "#ee3b3b",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "CR",
                    "id": null,
                    "label": "CRUD/FORM通用文档 - Avue.js",
                    "type": "text",
                    "url": "https://avuejs.com/doc/form-crud",
                    "uuid": "3fd9fdc0-a711-4800-904e-f4b76b842472"
                },
                "4cc236f3-e451-4240-b04f-879b052c2bcf": {
                    "backgroundColor": {
                        "data": "#97ffd8",
                        "type": "pure"
                    },
                    "desc": "随心写作，自由表达",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/37b581a34bd7b8315c2b2c9cb125feeb.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "图解：什么是红黑树？ - 知乎",
                    "type": "image",
                    "url": "https://zhuanlan.zhihu.com/p/273829162",
                    "uuid": "4cc236f3-e451-4240-b04f-879b052c2bcf"
                },
                "51b401be-3b72-4575-8163-259bbf6dcfd7": {
                    "backgroundColor": {
                        "data": "#97ffd8",
                        "type": "pure"
                    },
                    "desc": "随心写作，自由表达",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/37b581a34bd7b8315c2b2c9cb125feeb.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "Redis Stream——作为消息队列的典型应用场景 - 知乎",
                    "type": "image",
                    "url": "https://zhuanlan.zhihu.com/p/47470912",
                    "uuid": "51b401be-3b72-4575-8163-259bbf6dcfd7"
                },
                "570a3519-5a58-459f-87de-41f3a73dcd8f": {
                    "backgroundColor": {
                        "data": "#188fff",
                        "type": "pure"
                    },
                    "desc": "MONKNOW official website",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/f7365f50ea753205a8dbeb24bb8dcd2e.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": 1820,
                    "label": "MONKNOW",
                    "type": "image",
                    "url": "https://www.monknow.com",
                    "uuid": "570a3519-5a58-459f-87de-41f3a73dcd8f"
                },
                "57f92ef1-15b9-4650-ac88-e7b5c80b7749": {
                    "backgroundColor": {
                        "data": "#29b5ed",
                        "type": "pure"
                    },
                    "desc": "程序员分享平台",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/178b77b9665ae4f324a42226641ef4c8.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "[译] Envoy 为什么能战胜 Ngnix：线程模型分析篇 - 开发者头条",
                    "type": "image",
                    "url": "https://toutiao.io/posts/l9csrw/preview",
                    "uuid": "57f92ef1-15b9-4650-ac88-e7b5c80b7749"
                },
                "68809a49-85ef-4d07-a649-5b2e2cc18ae0": {
                    "backgroundColor": {
                        "data": "#1a73e8",
                        "type": "pure"
                    },
                    "desc": "Browser extensions",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/8413ddc9362f03a6902200356f3500df.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": 1708,
                    "label": "Extensions",
                    "type": "image",
                    "url": "internal://extensions",
                    "uuid": "68809a49-85ef-4d07-a649-5b2e2cc18ae0"
                },
                "6b433195-b51c-4281-9673-3b48bfd00cd4": {
                    "backgroundColor": {
                        "data": "#ee3b3b",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "卡片",
                    "id": null,
                    "label": "卡片管理——管理平台",
                    "type": "text",
                    "url": "https://www.pettyfox.top/admin/index.html#/card-group/card",
                    "uuid": "6b433195-b51c-4281-9673-3b48bfd00cd4"
                },
                "6efff3d9-7476-4986-ac4c-643a96dd9193": {
                    "backgroundColor": {
                        "data": "#478cc6",
                        "type": "pure"
                    },
                    "desc": "Beautiful free stock photos",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/88e2bc74ca7cf36c9118bac2e3c7a408.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "Hippopx - 精美的免版权图库",
                    "type": "image",
                    "url": "https://www.hippopx.com/zh/",
                    "uuid": "6efff3d9-7476-4986-ac4c-643a96dd9193"
                },
                "70158d50-85ba-411a-b465-d3a8af22d781": {
                    "backgroundColor": {
                        "data": "#16d9c4",
                        "type": "pure"
                    },
                    "desc": "基于私有云的流程图管理平台",
                    "icoText": "Dr",
                    "id": null,
                    "label": "我的流程图",
                    "type": "text",
                    "url": "http://www.pettyfox.top/drawio/admin/drawio.html",
                    "uuid": "70158d50-85ba-411a-b465-d3a8af22d781"
                },
                "74600bc7-2d83-4e5f-8393-d803592a2710": {
                    "backgroundColor": {
                        "data": "#fcb138",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "知识",
                    "id": null,
                    "label": "正式服-老师端-知识图谱",
                    "type": "text",
                    "url": "http://kg-tea.eface99.com:12000/#/user/login?redirect=%2F",
                    "uuid": "74600bc7-2d83-4e5f-8393-d803592a2710"
                },
                "755b16f1-bcb4-47a4-88fe-c1cfe6660a20": {
                    "backgroundColor": {
                        "data": "#9326e9",
                        "type": "pure"
                    },
                    "desc": "开源免费商用图标库iconfont",
                    "icoText": "Ic",
                    "id": null,
                    "label": "IcoMoon App - Icon Font, SVG, PDF & PNG Generator",
                    "type": "text",
                    "url": "https://icomoon.io/app/#/select",
                    "uuid": "755b16f1-bcb4-47a4-88fe-c1cfe6660a20"
                },
                "77279812-abf7-4ed3-b2aa-0636a5dcbf7f": {
                    "backgroundColor": {
                        "data": "#ee3b3b",
                        "type": "pure"
                    },
                    "desc": "Browser bookmarks",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/629a51cf5304c1855cc1da0e8287903f.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": 1709,
                    "label": "Bookmarks",
                    "type": "image",
                    "url": "internal://bookmarks",
                    "uuid": "77279812-abf7-4ed3-b2aa-0636a5dcbf7f"
                },
                "7f4620cd-179f-4a6a-aa58-42ff0ab3b7bc": {
                    "backgroundColor": {
                        "data": "#4b9ff5",
                        "type": "pure"
                    },
                    "desc": "4K资源下载、影音行业报道及影音器材评测!",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/f8065ea672c259be1161d33627943db3.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "Hao4K专注家庭影音-4K资源下载、影音行业报道及影音器材评测",
                    "type": "image",
                    "url": "https://www.hao4k.cn/",
                    "uuid": "7f4620cd-179f-4a6a-aa58-42ff0ab3b7bc"
                },
                "82b47318-5c66-4381-8124-41102636c76f": {
                    "backgroundColor": {
                        "data": "#333333",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "6.",
                    "id": null,
                    "label": "6.JVM —— 程序计数器 | Mr.Ai",
                    "type": "text",
                    "url": "https://leslieaibin.github.io/2021/05/12/Java/JVM/6.%E7%A8%8B%E5%BA%8F%E8%AE%A1%E6%95%B0%E5%99%A8/",
                    "uuid": "82b47318-5c66-4381-8124-41102636c76f"
                },
                "8b23b137-a9c0-44d1-ae86-afed1ee074c7": {
                    "backgroundColor": {
                        "data": "#71ce47",
                        "type": "pure"
                    },
                    "desc": "MONKNOW official wallpaper library",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/70381b0dbcca06e2eb30679040c6dcfa.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": 554,
                    "label": "Wallpaper Lib",
                    "type": "image",
                    "url": "internal://wallpaper-lib",
                    "uuid": "8b23b137-a9c0-44d1-ae86-afed1ee074c7"
                },
                "96774e24-7c99-4ca4-9d4a-f7d67b6bab8c": {
                    "backgroundColor": {
                        "data": "#276ce6",
                        "type": "pure"
                    },
                    "desc": "流程图共享协作工具",
                    "icoText": "ht",
                    "id": null,
                    "label": "https://miro.com",
                    "type": "text",
                    "url": "https://miro.com/",
                    "uuid": "96774e24-7c99-4ca4-9d4a-f7d67b6bab8c"
                },
                "9a463dc2-4ccc-44be-bbc6-d9aa121b5041": {
                    "backgroundColor": {
                        "data": "#409dff",
                        "type": "pure"
                    },
                    "desc": "基于 Vue 2.0 的桌面端组件库",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/bb6888645ce6d8e175b2cdc0e056dab7.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "组件 | Element",
                    "type": "image",
                    "url": "https://element.eleme.cn/#/zh-CN/component/installation",
                    "uuid": "9a463dc2-4ccc-44be-bbc6-d9aa121b5041"
                },
                "9bcfc82c-1c7d-4c91-a7c5-04dbf4458780": {
                    "backgroundColor": {
                        "data": "#ee3b3b",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "小狐后台",
                    "id": null,
                    "label": "本地-小狐记忆",
                    "type": "text",
                    "url": "http://localhost:7777/#",
                    "uuid": "9bcfc82c-1c7d-4c91-a7c5-04dbf4458780"
                },
                "9caff45f-2a17-41f2-8985-56a11ecb24b5": {
                    "backgroundColor": {
                        "data": "#fcb138",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "四川",
                    "id": null,
                    "label": "四川自考网—高等教育自学考试门户",
                    "type": "text",
                    "url": "https://www.tfzikao.com/",
                    "uuid": "9caff45f-2a17-41f2-8985-56a11ecb24b5"
                },
                "9f6fba14-a054-43f9-8f15-fa35269e0d01": {
                    "backgroundColor": {
                        "data": "#479f05",
                        "type": "pure"
                    },
                    "desc": "大型视频网站，专业的网络视频播放平台。",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/49d02ac31b2efd6c6da10ad9ffa3d909.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "Pandakill第一季 2016-09-23：pandakill正式篇第一期",
                    "type": "image",
                    "url": "https://www.iqiyi.com/v_19rr9gqva8.html#curid=542114700_9b1b8f41f4d4cc761c46fc0b88e70707",
                    "uuid": "9f6fba14-a054-43f9-8f15-fa35269e0d01"
                },
                "a1d19640-f094-4eb7-b199-13ab518a7e10": {
                    "backgroundColor": {
                        "data": "#288dff",
                        "type": "pure"
                    },
                    "desc": "记录，成为更好的自己",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/fab0e3d155dc35bfda1167f02b58b127.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "有道云笔记",
                    "type": "image",
                    "url": "https://note.youdao.com/web/#/file/recent/markdown/WEB54e031dc2ea37182d66781a1531701fc/",
                    "uuid": "a1d19640-f094-4eb7-b199-13ab518a7e10"
                },
                "a1f4dd0a-adfe-4626-829d-bad8bceb7130": {
                    "backgroundColor": {
                        "data": "#ffffff",
                        "type": "pure"
                    },
                    "desc": "促进软件开发及相关领域知识与创新的传播",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/a21a0a44157b1249c277ad5f922287a0.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "消息队列（五）如何保证消息的顺序性？ - InfoQ 写作平台",
                    "type": "image",
                    "url": "https://xie.infoq.cn/article/c84491a814f99c7b9965732b1",
                    "uuid": "a1f4dd0a-adfe-4626-829d-bad8bceb7130"
                },
                "a7d9599b-0764-4e5f-af5e-218ad03bca22": {
                    "backgroundColor": {
                        "data": "#fcb138",
                        "type": "pure"
                    },
                    "desc": "Browser history",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/2305d028ff04720a3b114e80c90caad3.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": 1707,
                    "label": "History",
                    "type": "image",
                    "url": "internal://history",
                    "uuid": "a7d9599b-0764-4e5f-af5e-218ad03bca22"
                },
                "ad584ea0-8886-40d3-96c7-e7777a538cbf": {
                    "backgroundColor": {
                        "data": "#16d9c4",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "在线",
                    "id": null,
                    "label": "在线Unicode编码解码 - 码工具",
                    "type": "text",
                    "url": "https://www.matools.com/code-convert-unicode",
                    "uuid": "ad584ea0-8886-40d3-96c7-e7777a538cbf"
                },
                "af67c476-adae-470a-9183-73f0de29603d": {
                    "backgroundColor": {
                        "data": "#209e85",
                        "type": "pure"
                    },
                    "desc": "提升英语词汇及听说读写能力",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/9b86e6c258c5ef71b4e974d4aad5fde5.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "记单词",
                    "type": "image",
                    "url": "https://web.shanbay.com/wordsweb/#/study/entry",
                    "uuid": "af67c476-adae-470a-9183-73f0de29603d"
                },
                "b76c8db3-eb00-4692-9d41-1aa5ac609392": {
                    "backgroundColor": {
                        "data": "#009a61",
                        "type": "pure"
                    },
                    "desc": "中国领先的新一代开发者社区和专业的技术媒体",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/a18fec96cb7b41d92b6ff569e8f25022.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "将项目发布到 maven 中央仓库踩过的坑 - SegmentFault 思否",
                    "type": "image",
                    "url": "https://segmentfault.com/a/1190000009450347",
                    "uuid": "b76c8db3-eb00-4692-9d41-1aa5ac609392"
                },
                "c1818ed1-e08b-4b54-9636-75e20d6e42c6": {
                    "backgroundColor": {
                        "data": "#ffffff",
                        "type": "pure"
                    },
                    "desc": "Ant Design 的 Vue 实现",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/5d0713940785faea9d106a12434dd59e.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "表格 Table - Ant Design Vue",
                    "type": "image",
                    "url": "https://www.antdv.com/components/table-cn/",
                    "uuid": "c1818ed1-e08b-4b54-9636-75e20d6e42c6"
                },
                "c1916a7d-dee2-40ad-ae5f-b970d13a6672": {
                    "backgroundColor": {
                        "data": "#fcb138",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "Gi",
                    "id": null,
                    "label": "GitHub 文件加速",
                    "type": "text",
                    "url": "https://shrill-pond-3e81.hunsh.workers.dev/",
                    "uuid": "c1916a7d-dee2-40ad-ae5f-b970d13a6672"
                },
                "c421c45a-cabc-440c-9786-7014d0e12f81": {
                    "backgroundColor": {
                        "data": "#ee3b3b",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "知识",
                    "id": null,
                    "label": "正式服-知识图谱管理平台",
                    "type": "text",
                    "url": "http://kg-admin.eface99.com:12000/#/user/login?redirect=%2F",
                    "uuid": "c421c45a-cabc-440c-9786-7014d0e12f81"
                },
                "c57db4ac-7dd5-411d-a8e7-d2a2764c363a": {
                    "backgroundColor": {
                        "data": "#333333",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "AS",
                    "id": null,
                    "label": "ASCII编码对照表_911查询",
                    "type": "text",
                    "url": "http://ascii.911cha.com/",
                    "uuid": "c57db4ac-7dd5-411d-a8e7-d2a2764c363a"
                },
                "c79b3f4a-806b-4fe7-8788-6cda2e43e132": {
                    "backgroundColor": {
                        "data": "#97ffd8",
                        "type": "pure"
                    },
                    "desc": "随心写作，自由表达",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/37b581a34bd7b8315c2b2c9cb125feeb.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "sonic——可替代Elasticsearch的简单搜索引擎 - 知乎",
                    "type": "image",
                    "url": "https://zhuanlan.zhihu.com/p/65790063",
                    "uuid": "c79b3f4a-806b-4fe7-8788-6cda2e43e132"
                },
                "ce5d9f93-deaa-45f0-ad42-b77111eb5a88": {
                    "backgroundColor": {
                        "data": "#16d9c4",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "Wo",
                    "id": null,
                    "label": "WordReference Word of the Day",
                    "type": "text",
                    "url": "https://daily.wordreference.com/",
                    "uuid": "ce5d9f93-deaa-45f0-ad42-b77111eb5a88"
                },
                "d6bbe31e-dbde-4b0c-8a72-0c269ad1a6e6": {
                    "backgroundColor": {
                        "data": "#97ffd8",
                        "type": "pure"
                    },
                    "desc": "随心写作，自由表达",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/37b581a34bd7b8315c2b2c9cb125feeb.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "一致性哈希算法（consistent hashing） - 知乎",
                    "type": "image",
                    "url": "https://zhuanlan.zhihu.com/p/129049724",
                    "uuid": "d6bbe31e-dbde-4b0c-8a72-0c269ad1a6e6"
                },
                "da44211c-390d-4c1c-88c0-92c6cfd9e518": {
                    "backgroundColor": {
                        "data": "#333333",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "V2",
                    "id": null,
                    "label": "V2Ray客户端 - V2ray XTLS黑科技",
                    "type": "text",
                    "url": "https://v2xtls.org/v2ray%e5%ae%a2%e6%88%b7%e7%ab%af/",
                    "uuid": "da44211c-390d-4c1c-88c0-92c6cfd9e518"
                },
                "da723fcc-af81-40f2-9362-54402ba551c9": {
                    "backgroundColor": {
                        "data": "#288dff",
                        "type": "pure"
                    },
                    "desc": "记录，成为更好的自己有道",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/fab0e3d155dc35bfda1167f02b58b127.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "有道云笔记",
                    "type": "image",
                    "url": "https://note.youdao.com/web/#/file/recent/note/WEBf275e6b8bd3fbd18822b763e43ad56c4/",
                    "uuid": "da723fcc-af81-40f2-9362-54402ba551c9"
                },
                "df00e31d-d41b-4bd1-a1d8-df09cec78b23": {
                    "backgroundColor": {
                        "data": "#9226e9",
                        "type": "pure"
                    },
                    "desc": "Browser downloads",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/d0359c6069b4b274fe4bb5b9ab326fb2.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": 1993,
                    "label": "Downloads",
                    "type": "image",
                    "url": "internal://downloads",
                    "uuid": "df00e31d-d41b-4bd1-a1d8-df09cec78b23"
                },
                "df68b8d9-3516-4ae3-b8e2-8d703693b3ec": {
                    "backgroundColor": {
                        "data": "#16d9c2",
                        "type": "pure"
                    },
                    "desc": "To Do gives you focus",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/030a4276f1d10e30cc0c50f017e96417.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": 2954,
                    "label": "ToDo",
                    "type": "image",
                    "url": "internal://todos",
                    "uuid": "df68b8d9-3516-4ae3-b8e2-8d703693b3ec"
                },
                "ef15afed-eab8-4729-a0dd-422b799a44eb": {
                    "backgroundColor": {
                        "data": "#97ffd8",
                        "type": "pure"
                    },
                    "desc": "随心写作，自由表达",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/37b581a34bd7b8315c2b2c9cb125feeb.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "RBAC用户、角色、权限、组设计方案 - 知乎",
                    "type": "image",
                    "url": "https://zhuanlan.zhihu.com/p/63769951",
                    "uuid": "ef15afed-eab8-4729-a0dd-422b799a44eb"
                },
                "ef20b9c3-b801-4010-a08f-b0c18379352e": {
                    "backgroundColor": {
                        "data": "#85d724",
                        "type": "pure"
                    },
                    "desc": "",
                    "icoText": "Sw",
                    "id": null,
                    "label": "知识后台-Swagger UI",
                    "type": "text",
                    "url": "http://localhost:20008/swagger-ui/index.html#/",
                    "uuid": "ef20b9c3-b801-4010-a08f-b0c18379352e"
                },
                "f0b16b35-5f9f-4d0f-996f-c81343f640de": {
                    "backgroundColor": {
                        "data": "#ffffff",
                        "type": "pure"
                    },
                    "desc": "专业IT技术发表平台",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/02ff6285310d63c7021a6398cd7e4e18.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": null,
                    "label": "绝了!华为技术专家居然把JVM内存模型讲解这么细致!_JavaEdge全是干货的技术号-CSDN博客_jvm内存模型",
                    "type": "image",
                    "url": "https://blog.csdn.net/qq_33589510/article/details/116423830",
                    "uuid": "f0b16b35-5f9f-4d0f-996f-c81343f640de"
                },
                "fdcd26fb-4798-4eb2-a0ef-06d944e9b2c0": {
                    "backgroundColor": {
                        "data": "#ffffff",
                        "type": "pure"
                    },
                    "desc": "A hidden Chrome dinosaur game",
                    "icoScalePercentage": 100,
                    "icoSrc": {
                        "data": "https://static.monknow.com/newtab/icon/9884c390fef75337f4606f2dc48c0f11.png",
                        "isOfficial": true,
                        "mimeType": "image/png",
                        "uploaded": true
                    },
                    "id": 2006,
                    "label": "Chrome Dino",
                    "type": "image",
                    "url": "internal://dino",
                    "uuid": "fdcd26fb-4798-4eb2-a0ef-06d944e9b2c0"
                },
                "6adced17-78ea-47b0-8b21-09a8c8b71db3": {
                    "id": null,
                    "uuid": "6adced17-78ea-47b0-8b21-09a8c8b71db3",
                    "url": "https://cn.bing.com/search?q=go+%E9%83%A8%E7%BD%B2docker&PC=U316&FORM=CHROMN",
                    "label": "go 部署docker - 国内版 Bing",
                    "type": "image",
                    "desc": "国际领先的搜索引擎，提供全球信息搜索服务。ab",
                    "backgroundColor": {
                        "type": "pure",
                        "data": "#0c8484"
                    },
                    "icoSrc": {
                        "isOfficial": true,
                        "uploaded": true,
                        "mimeType": "image/png",
                        "data": "https://static.monknow.com/newtab/icon/e48c9d25fca6f96d50010c66d4f4814b.png"
                    },
                    "icoScalePercentage": 100
                }
            }
        },
        "setting": {
            "borderRadiusPercentage": 20,
            "columnGapPercentage": 50,
            "displayAddBtn": true,
            "displayNotificationBadge": true,
            "displayShadow": true,
            "iconLayout": "particular",
            "opacityPercentage": 100,
            "openTarget": "currentTab",
            "rememberLastVisitedGroup": true,
            "rowGapPercentage": 40,
            "scrollToSwitchGroup": true,
            "sizePercentage": 100
        },
        "updaterVersion": 3,
        "version": 6
    }
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
            open(form) {
                // this.batchCreate()
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
