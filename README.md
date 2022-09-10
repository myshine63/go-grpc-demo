### 微服务
#### 这个是练习微服务时写的代码
#### 环境搭建
1. 安装protobuf
* 下载地址：https://github.com/protocolbuffers/protobuf/releases/tag/v21.5
* 下载解压后，在window系统的path环境变量中添加路径，测试是否可用 `protobuf -h`
2. 为go语言安装插件
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```
```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
3. 编译proto文件命令,注意要进入pb的文件夹
```shell
protoc --go_out=. ./*proto
```
```shell
protoc --go-grpc_out=. ./*proto
```
4. 安装consul
* 下载地址：https://www.consul.io/downloads，选择amd64表示是64位系统。
* 将consul可执行程序的路径添加到系统path环境变量，运行`consul -version`查看是否可用
5. consul agent 常用命令
* -bind 指定consul服务的ip地址,0.0.0.0表示选择可用的地址
* -http-port 指定consul访问端口
* -client 表面那些机器可以访问端口。默认本机，设置成0.0.0.0表示所有机器都可以访问
* -config-dir consul服务配置文件的地址
* -data-dir consul运行时的数据保存信息
* -dev 直接以默认配置启动consul
* -node 服务发现的名字，即consul的名字
* -rejoin consul启动的时候，可以加入到的consul集群
* -server 以服务的方式启动consul,同时允许其他的consul访问当前的consul
* -ui 可以用浏览器访问consul
6. consul启动命令demo
```shell
consul agent -server -bind=0.0.0.0 -config-dir=d:/consul/consul.d -data-dir=d:/consul/data -node=consul-demo -ui
```