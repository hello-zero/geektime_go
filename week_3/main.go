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
	if err := http.ListenAndServe("127.0.0.1:8081", http.DefaultServeMux); err !=nil{
		log.Fatal(err)
	}
	return nil
}

func main()  {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request){
			fmt.Fprintln(resp, "Hello, QCon!")
		})
		if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
			return err
		}
		return nil
	})

	group.Go(func() error {
		if err := http.ListenAndServe("127.0.0.1:8081", http.DefaultServeMux); err !=nil{
			log.Fatal(err)
		}
		return nil
	})

	group.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <- ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os exit: %v", sig)


		}
	})

	err := group.Wait()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ctx.Err())
}
