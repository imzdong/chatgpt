#### 项目目的

构建本地的chatgpt

#### 项目架构

前后端分离，[前端](https://github.com/imzdong/chatgpt-web)使用vue，后端使用go

#### 项目目录

```text
├── cmd
├── internal
├── templates
├── pkg
├── api
├── config
├── web
├── templates/
│   └── index.html
├── static/
│   └── style.css
└── go.mod
├── main.go
```

```text
cmd 目录：该目录用于存储主要的可执行文件。在一个大型的 Web 项目中，可能会存在多个可执行文件，例如 Web 服务器、命令行工具等。
internal 目录：该目录用于存储项目内部的私有代码。通常情况下，内部代码只能被项目本身访问，不能被其他项目或库访问。
pkg 目录：该目录用于存储项目的库代码。库代码通常会被其他项目或库引用和复用，因此需要保持相对独立和可复用。
api 目录：该目录用于存储 API 接口的定义和实现代码。在一个 Web 项目中，可能会存在多个 API 接口，因此需要将它们独立到一个目录中进行管理。
config 目录：该目录用于存储项目的配置文件。配置文件通常包括数据库连接、密钥等敏感信息，因此需要妥善保管和管理。
web 目录：该目录用于存储 Web 相关的代码，例如路由、控制器、模板等。在一个大型的 Web 项目中，可能会存在多个子目录，例如 web/api 目录用于存储 API 相关的代码，web/admin 目录用于存储后台管理界面相关的代码等。
static 目录：该目录用于存储静态文件，例如图像、样式表、JavaScript 文件等。通常情况下，静态文件可以通过 HTTP 服务器直接访问。
templates 目录：该目录用于存储模板文件。模板文件通常包括 HTML、CSS、JavaScript 等代码，用于生成动态的 Web 页面。
```

#### 项目构建
* 构建以上目录结构
* 构建go模块
    ```shell
    go mod init chatgpt
    ```
* 运行go项目
  ```shell
    go run main.go
  ```
  ```text
    vue项目通过build打包的dist。
  其中的index.html文件移动到templates下面，其他的css和js移动到static下面。
  并且修改index.html引入的css和js的路径，添加/static
```
  