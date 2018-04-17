# unexpected directory layout 意外的目录布局

    分析：在机器上拥有多个 GOPATH 的情况下，又在项目里使用了 vendor，恰好 vendor 里的库跟其它项目的相同时，出现了该错误。

    解决：构建时指定 GOPATH，如 GOPATH=xxx go build # xxx 只有一个目录
