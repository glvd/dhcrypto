1.生成.so文件

```
go build -buildmode=c-shared -o crypto.so
```

2. 生成.a文件
```
go build -buildmode=c-archive -o crypto.a
```
