package main

import (
	"io"
	"net/http"
)

// 统一请求入口函数
func index(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"hello")
}

func main(){
        // 启动8083 端口
	http.ListenAndServe(":8083",http.HandlerFunc(index))   // 无论路由如何去写，都会进入到 index 函数中
}