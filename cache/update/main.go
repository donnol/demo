// 缓存更新的几种策略
package main

func main() {

}

// CacheAsidePattern 缓存在旁边模式
// 从缓存取，没取到时，去数据库取，取到后放入缓存
// 更新数据库时，让缓存失效，但是并不马上更新缓存，等取的时候再更新
func CacheAsidePattern() {
	// 从缓存取
	_ = func() string {
		var key string
		value, ok := cache(key)
		if !ok { // 缓存中没有，则从数据库拿
			value = db(key)
			setCache(key, value) // 从数据库拿到后，放到缓存中
		}
		return value
	}

	// 更新
	_ = func() {
		var key, value string
		setDB(key, value)
		deleteCache(key) // 更新了数据库后，让缓存失效
	}
}

// ReadThroughPattern 读通过
// 在查询操作中更新缓存，也就是说，当缓存失效的时候（过期或LRU换出），Cache Aside是由调用方负责把数据加载入缓存，而Read Through则用缓存服务自己来加载
func ReadThroughPattern() {
	// 从缓存取
	_ = func() string {
		var key string
		value, ok := cache(key)
		if !ok { // 缓存中没有，缓存服务自己从数据库获取数据
			setCache(key, value)
		}
		return value
	}
}

// WriteThroughPattern 写通过
// 当有数据更新的时候，如果没有命中缓存，直接更新数据库，然后返回。如果命中了缓存，则更新缓存，然后再由Cache自己更新数据库（这是一个同步操作）
func WriteThroughPattern() {
	// 更新
	_ = func() {
		var key, value string
		value, ok := cache(key)
		if ok {
			setCache(key, value) // 缓存服务更新数据库
		} else { // 没命中，直接更新数据库
			setDB(key, value)
		}
	}
}

// WriteBehindCachingPattern 写在缓存后面，又称Write Back
// 类似Linux文件系统的Page Cache的算法
// 在更新数据的时候，只更新缓存，不更新数据库，而我们的缓存会异步地批量更新数据库
func WriteBehindCachingPattern() {
	// 更新
	_ = func() {
		var key, value string
		setCache(key, value)
		go setDB(key, value) // 异步更新数据库
	}
}

func cache(key string) (value string, ok bool) {
	return
}

func setCache(key, value string) {

}

func deleteCache(key string) {

}

func db(key string) (value string) {
	return
}

func setDB(key, value string) {

}
