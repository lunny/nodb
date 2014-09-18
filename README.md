# NoDB

[中文](https://github.com/lunny/nodb/blob/master/README_CN.md)

Nodb is a fork of [ledisdb](https://github.com/siddontang/ledisdb) and shrink version. It's get rid of all C or other language codes and only keep Go's. It aims to provide a nosql database library rather than a redis like server. So if you want a redis like server, ledisdb is the best choose.

Nodb is a pure Go and high performance NoSQL database library. It supports some data structure like kv, list, hash, zset, bitmap, set.

Nodb now use goleveldb as backend to store data.

## Features

+ Rich data structure: KV, List, Hash, ZSet, Bitmap, Set.
+ Stores lots of data, over the memory limit. 
+ Supports lua scripting.
+ Supports expiration and ttl.
+ Easy to embed in your own Go application.

## Install

    go get github.com/lunny/nodb

## Package Example
    
    import "github.com/lunny/nodb"

    l, _ := nodb.Open(cfg)
    db, _ := l.Select(0)

    db.Set(key, value)

    db.Get(key)

## Links

+ [Ledisdb Official Website](http://ledisdb.com)
+ [GoDoc](https://godoc.org/github.com/lunny/nodb)
+ [GoWalker](https://gowalker.org/github.com/lunny/nodb)


## Thanks

Gmail: siddontang@gmail.com
