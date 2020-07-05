package main

import (
	"fmt"
	"net/http"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort), // 监听TCP地址
		Handler:        router,                               //http句柄，用于处理程序响应http请求
		ReadTimeout:    setting.ReadTimeout,                  //允许读取的最大时间
		WriteTimeout:   setting.WriteTimeout,                 //允许写入最大的时间
		MaxHeaderBytes: 1 << 20,                              //请求头的最大字节数
	}

	s.ListenAndServe()
}
