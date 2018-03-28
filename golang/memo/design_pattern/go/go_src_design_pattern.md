# go 源码里面的 设计模式

1 工厂模式(factory)

```go
// https://golang.org/src/image/format.go?s=2205:2252#L68

...
// Sniff determines the format of r's data.
func sniff(r reader) format {
    for _, f := range formats {
        b, err := r.Peek(len(f.magic))
        if err == nil && match(f.magic, b) {
            return f
        }
    }
    return format{}
}

// Decode decodes an image that has been encoded in a registered format.
// The string returned is the format name used during format registration.
// Format registration is typically done by an init function in the codec-
// specific package.
func Decode(r io.Reader) (Image, string, error) {
    rr := asReader(r)
    f := sniff(rr)
    if f.decode == nil {
        return nil, "", ErrFormat
    }
    m, err := f.decode(rr)
    return m, f.name, err
}
...
```

2 外观模式(facade)

```go
// https://golang.org/src/io/ioutil/ioutil.go?s=3442:3483#L105

...
type nopCloser struct {
    io.Reader
}

func (nopCloser) Close() error { return nil }

// NopCloser returns a ReadCloser with a no-op Close method wrapping
// the provided Reader r.
func NopCloser(r io.Reader) io.ReadCloser {
    return nopCloser{r}
}
...
```

3 适配器模式(Adapter)

```go
// 同上面的外观模式
```

4 单例模式(singleton)

```go
// https://golang.org/src/sync/once.go?s=1137:1164#L25

...
func (o *Once) Do(f func()) {
    if atomic.LoadUint32(&o.done) == 1 {
        return
    }
    // Slow-path.
    o.m.Lock()
    defer o.m.Unlock()
    if o.done == 0 {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}
...
```

5 工厂方法模式(factorymethod)

6 抽象工厂模式(abstractfactory)

7 (builder)

8 原型模式(prototype)

9 中介者模式(mediator)

10 代理模式(proxy)

11 观察者模式(observer)

12 命令模式(command)

13 迭代器模式(iterator)

14 组合模式(composite)

15 模板方法模式(templatemethod)

16 策略模式(strategy)

17 状态模式(state)

18 备忘录模式(memento)

19 享元模式(flyweight)

20 解释器模式(interpreter)

21 装饰模式(decorator)

22 职责链模式(chain)

23 桥接模式(bridge)

24 访问者模式(visitor)
