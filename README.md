# superstar-golang
golang后台实现的球星库


根据一凡老师教程自己一步一步敲的，中间遇到很多坑。列举下：


1.   shared/layout.html错误
```
[ERRO] 2019/03/19 15:45 html/template: "shared/layout.html" is undefined
[INFO] 2019/03/19 15:45 500 29.252036ms ::1 GET /
[ERRO] 2019/03/19 15:45 html/template: "shared/layout.html" is undefined
```

这个是路径的问题，修改下路径
在bootstrapper.go  中
```
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./web/views")
	b.SetupSessions(24*time.Hour,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	b.SetupErrorHandlers()

	// static files
	b.Favicon(StaticAssets + Favicon)
	b.HandleDir(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	// middleware, after static files
	b.Use(recover.New())
	b.Use(logger.New())

	return b
}
```
关键在于b.SetupViews("./web/views")，views是在web目录下，多写一个web目录


2.  StaticAssets 路径报错

同理，需要web目录
```
const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./web/public/"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
)
```

3. bootstrap的css样式加载不出来

同理，加入web目录
```
<link href="/web/public/css/bootstrap.min.css" rel="stylesheet">
```

```
<script src="/web/public/js/jquery.min.js"></script>
<script>window.jQuery || document.write('<script src="/web/public/js/jquery.min.js"><\/script>')</script>
<script src="/web/public/js/bootstrap.min.js"></script>
```

4. edit.html报错
错误在于本句话
```
{{if gt .info.Id 0}}编辑{{else}}添加{{end}}球员信息
```
语法错误，但是不知道怎么改？暂时替换成   "添加球员信息"  成功！

5.  最后一个问题，为什么老是的xorm生成的model有form。我的没有？？？何解？？？
