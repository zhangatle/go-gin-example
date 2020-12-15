package main

import (
	"fmt"
	"github.com/zhangatle/go-gin-example/pkg/setting"
	"github.com/zhangatle/go-gin-example/routers"
	"net/http"
)

func main()  {
	r := routers.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPort),
		Handler: r,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
