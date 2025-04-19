# ShadowFetch Vercel

一个基于Vercel的反向代理服务，可以用来代理任何HTTP请求，特别适合代理PostHog等服务。

## 功能特点

- 支持所有HTTP方法（GET, POST, PUT, DELETE等）
- 支持静态资源缓存
- 易于部署和配置
- 使用Vercel Edge Functions，全球分布式部署

## 使用方法

### 部署到Vercel

1. Fork或克隆此仓库
2. 修改配置（可选）
   - 在`api/edge.js`中修改`API_HOST`和`ASSET_HOST`为你想要代理的目标
   - 或者使用`api/posthog.js`专门代理PostHog服务
   - 在`vercel.json`中根据需要调整配置
3. 使用Vercel CLI部署：
   ```bash
   vercel login
   vercel
   ```

### 配置

#### 通用代理
你可以在`api/edge.js`文件中修改以下配置：

```javascript
const API_HOST = "httpbin.org" // 修改为你想要代理的目标
const ASSET_HOST = "httpbin.org" // 修改为静态资源的目标
```

#### PostHog代理
你可以在`api/posthog.js`文件中修改以下配置：

```javascript
const API_HOST = "us.i.posthog.com" // 美国区域，可改为 "eu.i.posthog.com" 使用欧盟区域
const ASSET_HOST = "us-assets.i.posthog.com" // 美国区域，可改为 "eu-assets.i.posthog.com" 使用欧盟区域
```

## 本地开发

1. 安装依赖：
   ```bash
   npm install
   ```

2. 启动本地开发服务器：

   使用Vercel CLI（推荐）：
   ```bash
   npm run dev
   ```

   或者使用简单的Node.js服务器：
   ```bash
   npm run dev:local
   ```

## 注意事项

- 此代理服务仅用于学习和测试目的
- 请遵守Vercel的服务条款和目标网站的使用条款
- 不要用于非法用途

## 许可证

MIT
