#Client for minio
  This document describes how to use the cli to operate objects with minio.
  Base sdk is from official resources:  
  [golang-client-quickstart-guide](https://docs.min.io/docs/golang-client-quickstart-guide.html)     
  [Github](https://github.com/minio/minio-go): there are many example codes in /examples/s3/*
##tree of mino
    .
    ├── api
    │   ├── bucket.go # 桶操作相关函数
    │   ├── client.go # 认证相关函数
    │   └── object.go # 对象操作相关函数
    ├── cmd
    │   ├── minio.go # cobra & cli 
    │   └── root.go # copy from mogdb operator --- cmd.root
    ├── config
    │   └── config.go # 定义 Credentials
    ├── go.mod
    ├── go.sum
    ├── main.go # cmd.Execute()
    ├── main_test.go # test /api/*
    ├── Readme.md


##Uses of cli
Base command is mgo, first level subcommand is minio, second level subcommand is bucket and object.
### mgo minio bucket 
bucket等同于文件夹

    mgo minio bucket ls # 列出桶
    mgo minio bucket make bucket-name # 新建一个桶
    mgo minio bucket rm bucket-name # 删除一个桶

### mgo minio object
object等同于文件

    mgo minio object ls bucket-name --prefix prefix-name # 列出bucket-name下的带有前缀prefix-name的对象文件
    mgo minio object getf bucket-name object-name -p filepath # 把bucket-name:object-name文件保存到 filepath， -p filepath可省略，则直接在本程序的当前目录下保存
    mgo minio object putf bucket-name object-name -p filepath # 把filepath文件上传到 bucket-name:object-name -p filepath不可省略
    mgo minio object stat bucket-name object-name # 把bucket-name:object-name文件的状态打印出来
    mgo minio object rm bucket-name object-name # 删除minio服务器的对象bucket-name:object-name
    mgo minio object rmall bucket-name # 删除minio桶 bucket-name 内所有对象






