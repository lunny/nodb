# NoDB

[English](https://github.com/lunny/nodb/blob/master/README.md)

Nodb 是 [ledisdb](https://github.com/siddontang/ledisdb) 的克隆和缩减版本。该版本去掉了所有C和其它语言的依赖，只保留Go语言的。目标是提供一个Nosql数据库的开发库而不是提供一个像Redis那样的服务器。因此如果你想要的是一个独立服务器，你可以直接选择ledisdb。

Nodb 是一个纯Go的高性能 NoSQL 数据库。他支持 kv, list, hash, zset, bitmap, set 等数据结构。

Nodb 当前底层使用 goleveldb 来存储数据。

## 特性

+ 丰富的数据结构支持： KV, List, Hash, ZSet, Bitmap, Set。
+ 永久存储并且不受内存的限制。
+ 高性能那个。
+ 可以方便的嵌入到你的应用程序中。

## 安装

    go get github.com/lunny/nodb

## 例子
    
    import "github.com/lunny/nodb"

    l, _ := nodb.Open(cfg)
    db, _ := l.Select(0)

    db.Set(key, value)

    db.Get(key)

## 链接

+ [Ledisdb Official Website](http://ledisdb.com)
+ [GoDoc](https://godoc.org/github.com/lunny/nodb)
+ [GoWalker](https://gowalker.org/github.com/lunny/nodb)


## 感谢

Gmail: siddontang@gmail.com
