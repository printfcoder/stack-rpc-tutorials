package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)

// exampleCall 方法负责处理/example/call路由
func exampleCall(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// 获取传入参数值
	name := r.Form.Get("name")

	if len(name) == 0 {
		http.Error(
			w,
			errors.BadRequest("go.micro.api.example", "no content").Error(),
			400,
		)
		return
	}

	// 序列化响应参数
	b, _ := json.Marshal(map[string]interface{}{
		"message": "exampleCall 收到了你的消息，" + name,
	})

	w.Write(b)
}

// exampleFooBar 方法负责处理 /example/foo/bar 路由
func exampleFooBar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(
			w,
			errors.BadRequest("go.micro.api.example", "require post").Error(),
			400,
		)
		return
	}

	if len(r.Header.Get("Content-Type")) == 0 {
		http.Error(
			w,
			errors.BadRequest("go.micro.api.example", "need content-type").Error(),
			400,
		)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(
			w,
			errors.BadRequest("go.micro.api.example", "expect application/json").Error(),
			400,
		)
		return
	}

	// 获取传入参数值
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	data := struct {
		Name string `json:"name"`
	}{}
	json.Unmarshal(bodyBytes, &data)

	// 序列化响应参数
	b, _ := json.Marshal(map[string]interface{}{
		"message": "exampleFooBar 收到了你的消息，" + data.Name,
	})

	w.Write(b)
}

// uploadFile 方法负责处理/example/foo/upload上传文件的路由
func uploadFile(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		t, _ := template.New("foo").Parse(`<html>
<head>
       <title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="http://127.0.0.1:8080/example/foo/upload" method="post">
    <input type="file" name="uploadfile" />
   <br />
   保存目录： <input type="text" name="path" /> 如 /Users/me/Downloads/test/
     <br />
    <input type="submit" name='上传' value="upload" />
</form>
</body>
</html>`)
		t.Execute(w, nil)

		return
	}

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)

	path := r.PostForm.Get("path")
	f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func main() {
	// 我们为了方便就直接使用go-web包，因为API需要从注册中心获取服务信息，而go-web包已经有注册服务的能力
	service := web.NewService(
		web.Name("go.micro.api.example"),
	)

	service.HandleFunc("/example/call", exampleCall)
	service.HandleFunc("/example/foo/bar", exampleFooBar)
	service.HandleFunc("/example/foo/upload", uploadFile)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
