# Golang Iris Web Server Starter Kit

一个基于go语言的iris框架搭建的开发环境，可以直接套用并进行开发，省去各种麻烦事。

`A golang web server starter kit with iris.`

##Go环境

默认使用Go语言版本 `1.11`，并启用`Go Modules`

## 目录结构

- src
    - structs  `存放系统所需的基础结构体`
        - config.go `系统配置的结构体，存放配置文件config.yml的解析内容`
        - restful.go `定义Web接口的RESTful结构体`
        - query_signature.go `定义Query参数校验相关的结构体，一般是用于不同web服务之间的私密通讯`
    - utils    `基本的、常用的工具函数等`
        - common.go `内含JSON解析相关方法`
        - global.go `提供载入和解析系统配置文件的函数`
        - query_signature.go   `提供Query参数校验相关的功能`
    
- config.yml    `项目配置文件`