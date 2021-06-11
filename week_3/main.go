package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)
/**
1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
**/
func serveApp() error {
	return errors.New("firrst")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request){
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		return err
	}
	return nil
}

func serveDebug() error{
	return nil
	if err := http.ListenAndServe("127.0.0.1:8081", http.DefaultServeMux); err !=nil{
		log.Fatal(err)
	}
	return nil
}

func main()  {
	group, ctx := errgroup.WithContext(context.Background())
	// 初始化连接
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request){
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	server := http.Server{
		Handler: mux,
		Addr: ":8080",
	}

	serverOut := make(chan int)

	// group1 启动server
	group.Go(func() error {
		return server.ListenAndServe()
	})

	// group2 关闭server
	group.Go(func() error {
		select {
		case <- serverOut:
			fmt.Println("server exit")
			// func退出会触发g.cancel, ctx.done会收到信号
		case <- ctx.Done():
			fmt.Println("errgoup exit")
		}
		return server.Shutdown(ctx)
	})

	// group3 signal的注册
	group.Go(func() error {
		quit := make(chan os.Signal)
		// sigint 用户ctrl+c, sigterm程序退出
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <- ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os exit: %v", sig)


		}
	})

	err := group.Wait()
	fmt.Println(err)
	fmt.Println(ctx.Err())
}
