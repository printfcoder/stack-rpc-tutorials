# 原型文件

## 生成指令

```bash
protoc --proto_path=. --go_out=paths=source_relative:./ micro_benchmark.proto --micro_out=paths=source_relative:./
```
