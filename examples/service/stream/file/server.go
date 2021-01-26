package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/stack-labs/stack-rpc"
	file2 "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/stream/file"
	"github.com/stack-labs/stack-rpc/util/errors"
)

type File struct{}

// 文件接收方法
func (g *File) File(ctx context.Context, file file2.File_FileStream) error {
	//将接受到的内容储存到临时文件中
	temp, err := ioutil.TempFile("", "stack")
	if err != nil {
		return errors.InternalServerError("file.service", err.Error())
	}
	for {
		b, err := file.Recv()
		if err != nil {
			return errors.InternalServerError("file.service", err.Error())
		}
		if b.Len == -1 {
			//完成标志
			break
		}
		if _, err := temp.Write(b.Byte); err != nil {
			return errors.InternalServerError("file.service", err.Error())
		}
	}
	println(temp.Name())
	// 发送文件信息
	return file.SendMsg(&file2.FileMsg{
		FileName: temp.Name(),
	})
}

// 文件处理方法
func (g *File) DealFile(ctx context.Context, req *file2.DealFileRequest, rsp *file2.DealFileRespond) error {
	// 通过文件名获取到文件内容
	// 计算文件md5
	hash := md5.New()
	file, err := os.OpenFile(req.FileName, os.O_RDONLY, 0755)
	if err != nil {
		return errors.InternalServerError("file.service", err.Error())
	}
	_, _ = io.Copy(hash, file)
	MD5Str := hex.EncodeToString(hash.Sum(nil))
	// 加上param
	hash2 := md5.New()
	hash2.Write([]byte(MD5Str + req.Param))
	rsp.Md5 = hex.EncodeToString(hash2.Sum(nil))
	println(req.FileName + "|" + rsp.Md5)
	return nil
}

func main() {
	// 创建服务，除了服务名，其它选项可加可不加，比如Version版本号、Metadata元数据等
	service := stack.NewService(
		stack.Name("file.service"),
		stack.Version("latest"),
	)
	service.Init()

	// 注册服务
	_ = file2.RegisterFileHandler(service.Server(), new(File))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
