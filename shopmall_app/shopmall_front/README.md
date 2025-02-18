# HelloWorld 前端项目

## 项目概述
本项目是 HelloWorld 应用的前端部分，基于 React + TypeScript + Vite 构建的现代化 Web 应用。

## 功能特性
- 用户认证与授权
- 响应式界面设计
- 实时数据交互
- 主题定制支持

## 技术架构
### 核心技术栈
- **React 18**: 用于构建用户界面的 JavaScript 库
- **TypeScript**: 添加静态类型支持的 JavaScript 超集
- **Vite**: 现代前端构建工具，提供极速的开发体验

### 项目结构
```
src/
├── api/        # API 接口封装
├── assets/     # 静态资源文件
├── components/ # 可复用组件
├── pages/      # 页面组件
├── App.tsx     # 应用入口组件
└── main.tsx    # 应用启动文件
```

## 使用组件
- **React Router**: 路由管理
- **Axios**: HTTP 客户端
- **ESLint**: 代码质量检查
- **Prettier**: 代码格式化工具

## 本地开发

### 环境要求
- Node.js 16+
- npm 7+ 或 yarn

### 安装依赖
```bash
npm install
# 或
yarn
```

### 依赖重装指南
如果在安装依赖时遇到问题，可以尝试以下步骤重新安装：

```bash
# 清除 npm 缓存
npm cache clean --force

# 删除现有的 node_modules 目录和 package-lock.json
rm -rf node_modules package-lock.json

# 重新安装依赖
npm install
```

对于使用 yarn 的用户：
```bash
# 清除 yarn 缓存
yarn cache clean

# 删除现有的 node_modules 目录和 yarn.lock
rm -rf node_modules yarn.lock

# 重新安装依赖
yarn
```

### 启动开发服务器
```bash
npm run dev
# 或
yarn dev
```

### 构建生产版本
```bash
npm run build
# 或
yarn build
```

## ESLint 配置
项目使用 ESLint 进行代码质量控制，支持类型检查。主要配置如下：

```js
export default tseslint.config({
  languageOptions: {
    parserOptions: {
      project: ['./tsconfig.node.json', './tsconfig.app.json'],
      tsconfigRootDir: import.meta.dirname,
    },
  },
})
```

## 贡献指南
1. Fork 本仓库
2. 创建特性分支
3. 提交变更
4. 发起 Pull Request

## 许可证
MIT License
