# Flags

演示如何使用命令行的flag参数

## 运行示例

```shell
INT_FLAG=109 go run main.go --string_flag="串串来了" --bool_flag=true
```

可以看到程序打印出获取到的flag参数信息

```text
字符串flag值： 串串来了
字符串缺省值： 我是缺省值
整形flag值：109
布尔值flag值：true
```
