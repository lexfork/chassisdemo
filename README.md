# chassisdemo
go-chassis demo

# 微服务列表
* API： 服务对外接口
* RestServer： Rest服务
* RpcServer： RPC服务

# 运行
```$xslt
heavyrains-MacBook-Pro:chassisdemo heavyrainlee$ ./chassisdemo 
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  chassisdemo [command]

Available Commands:
  api         A brief description of your command
  help        Help about any command
  restserver  A brief description of your command
  rpcserver   A brief description of your command

Flags:
  -c, --config string   config file path (default is ./${command}/conf/)
  -h, --help            help for chassisdemo
  -t, --toggle          Help message for toggle

Use "chassisdemo [command] --help" for more information about a command.

```

### 运行API服务
```$xslt
heavyrains-MacBook-Pro:chassisdemo heavyrainlee$ ./chassisdemo api
```
> 测试时使用单实例

### 运行Rest服务
```$xslt
heavyrains-MacBook-Pro:chassisdemo heavyrainlee$ ./chassisdemo restserver
```
> 可以运行多个，运行多个时，可执行文件和config文件需要放到不同的目录，然后修改端口号

### 运行RPC服务
```$xslt
heavyrains-MacBook-Pro:chassisdemo heavyrainlee$ ./chassisdemo rpcserver
```
> 可以运行多个，运行多个时，可执行文件和config文件需要放到不同的目录，然后修改端口号

# 测试
当所有服务都成功启动后，执行下面两个命令检查结果
```$xslt
curl localhost:5000/sayresthello/heavyrain
curl localhost:5000/sayrpchello/heavyrain
```