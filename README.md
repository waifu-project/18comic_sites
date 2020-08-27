# 18comic get sites

获取镜像站, 发布地址: http://jmcomic.xyz

通过脚本格式化, 部署之后访问接口 `/api/mirror`, 返回结果演示

```json
[
    {
        "title": "海外分流",
        "api": "18comic.org"
    },
    {
        "title": "JM主站",
        "api": "18comic.vip"
    },
    {
        "title": "JM中國主站",
        "api": "18comic1.biz"
    },
    {
        "title": "分流1 ",
        "api": "18comic2.biz"
    },
    {
        "title": "分流2",
        "api": "18comic3.biz"
    }
]
```

typescript interface

```typescript
export interface item {
  title: string
  api: string
}
```