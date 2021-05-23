# Golang Basic - Setup Golang Environment

[toc]

## Golang 安装

- MAC下面安装Golang:

```shell
brew update
brew install golang
brew cleanup
go version
```
- MAC下面升级Golang

```shell
brew upgrade go
go version
brew cleanup
```

- 环境变量
    * GOPATH: 默认为$HOME/go
    * GOROOT: Go语言编译、工具、标准库等的安装路径,默认为/usr/local/go,homebrew 安装之后目录为: `$(brew –prefix golang)/libexec`
      
- 自行修改以上两个环境变量
修改配置文件比如~/.zshrc, ~/.bashrc等
  
```shell
export GOPATH="${HOME}/.go" # 这个路径可以自己修改
export GOROOT="$(brew --prefix golang)/libexec"  # 默认路径，也可以自己修改到golang的不同安装路径
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin" #加入到bin目录
```

## Golang IDE

- goland，付费软件，如何使用可以自行查找
- [VS code](https://code.visualstudio.com/)

## 配置VS-Code

- 安装VS-Code的[Golang 插件](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
- 一个golang的项目，VS-code默认开始安装需要的工具
```shell
Tools environment: GOPATH=/Users/Patrick/go
Installing 8 tools at /Users/Patrick/go/bin in module mode.
  gopkgs
  go-outline
  gotests
  gomodifytags
  impl
  goplay
  dlv
  gopls

Installing github.com/uudashr/gopkgs/v2/cmd/gopkgs (/Users/Patrick/go/bin/gopkgs) SUCCEEDED
Installing github.com/ramya-rao-a/go-outline (/Users/Patrick/go/bin/go-outline) SUCCEEDED
Installing github.com/cweill/gotests/gotests (/Users/Patrick/go/bin/gotests) SUCCEEDED
Installing github.com/fatih/gomodifytags (/Users/Patrick/go/bin/gomodifytags) SUCCEEDED
Installing github.com/josharian/impl (/Users/Patrick/go/bin/impl) SUCCEEDED
Installing github.com/haya14busa/goplay/cmd/goplay (/Users/Patrick/go/bin/goplay) SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv (/Users/Patrick/go/bin/dlv) SUCCEEDED
Installing golang.org/x/tools/gopls (/Users/Patrick/go/bin/gopls) SUCCEEDED

```

常用的这些安装工具的说明:
  - gocode。代码补全，注意，网上有很多文章用的是https://github.com/nsf/gocode，这是错误的，因为nsf/gocode现在已经不维护了。
  - godef。跳转到代码对应定义的，这个可以说是各种编程语言IDE必备的功能。
  - guru。guru原来叫做oracle，能获得全部代码对应引用
  - gorename。重命名Go源码文件
  - goimports。自动帮你格式化import语句的，它来帮你决定import顺序
  - gopkgs。列出Go包。
  - golangci-lint。旧的用法是gometalinter，现在应该选择快5倍的golangci-lint，这个工具用来做静态检查工具，它会反馈代码风格并给出建议，这样就可以在源码保存是运行它能告诉你现有代码中有什么问题
  - go-outline。当前文件做符号(Symbol)搜索，相当于文件大纲
  - go-symbols。工作区符号搜索
  - gogetdoc。鼠标悬停时可以显示文档(方法和类的签名等帮助信息)
  - gomodifytags。提供tags管理，可以对struct的tag增删改
  - gotests。用例测试
  - impl。自动生成接口实现代码(stubs)
  - gopkgs。导入包时自动补全(列出可导入的包列表，主要用于包导入时的提示功能)
  - fillstruct。根据结构体定义在使用时自动填充默认值

- 配置Visual Studio Code

安装以上内容之后，基本上不要额外配置，就可以使用visual studio code开发了



