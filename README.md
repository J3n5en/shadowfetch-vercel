# ShadowFetch for Vercel

## 功能特点

- 支持所有HTTP方法（GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS）
- 密码保护，防止未授权访问
- 可选的Chrome浏览器模拟功能
- 自动转发HTTP头和请求体
- 部署在Vercel平台，利用全球边缘网络

## 环境变量配置

项目使用以下环境变量进行配置：

| 环境变量 | 说明 | 默认值 |
|---------|------|-------|
| `TARGET_URL` | 目标服务器的URL，所有请求将被转发到这个URL | `https://httpbin.org/` |
| `PASSWORD` | 访问代理服务的密码，用于URL路径中 | `J3` |
| `DEBUG_MODE` | 是否启用调试模式，设置为"1"开启 | `0` |
| `IMPERSONATE_CHROME` | 是否模拟Chrome浏览器，设置为"1"开启 | `1` |

## 使用方法

1. 部署此项目到Vercel
2. 设置必要的环境变量
3. 通过以下格式访问代理服务：

```
https://your-vercel-domain.vercel.app/{PASSWORD}/{path}
```

例如，如果您的密码是`J3`，并且想要访问`https://httpbin.org/get`，则URL应为：

```
https://your-vercel-domain.vercel.app/J3/get
```
