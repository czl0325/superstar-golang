# superstar-golang
golang后台实现的球星库


本文主要是对golang的iris库，xorm库的应用。
首先你要安装iris库，本库安装比较麻烦，具体参考我的另外一篇文章  [泪流满面的golang1.12.7-iris库安装配置全过程](https://www.jianshu.com/p/1d6cde7d355c)
PS: go mod tidy -v.  自动下载依赖包，同时删除多余的包

1。 创建conf文件夹，创建db.go文件，配置一下自己的数据库。

2.  使用navicat创建superstar的数据库，建立一张表
![](https://upload-images.jianshu.io/upload_images/2587882-d8ae1f66cd6e668b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1000)

3.  通过xorm工具来生成models

4.  
