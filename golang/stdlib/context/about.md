# context

上下文

在一个请求中，使用了多个 goroutine 去处理它时，如果这个请求突然被取消了，那么关联的 goroutine 都要被取消，否则会浪费很多资源
