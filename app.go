package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/sync/errgroup"
	"net/http"
	"time"
	"log"
)

var (
	group errgroup.Group
)

func main() {
	log.Println("当前运行环境:", gin.Mode())

	serverInner := &http.Server{
		Addr:         fmt.Sprintf(":%d", ChatConfig.InnerPort),
		Handler:      initRouterInner(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	serverOutter := &http.Server{
		Addr:         fmt.Sprintf(":%d", ChatConfig.OutterPort),
		Handler:      initRouterOutter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	group.Go(func() error {
		return serverInner.ListenAndServe()
	})

	group.Go(func() error {
		return serverOutter.ListenAndServe()
	})

	if err := group.Wait(); err != nil {
		log.Printf("run error: %+v\n", err)
	}

	//var a int = 1
	//var b *int = &a
	//var c **int = &b
	//var x int = *b
	//fmt.Println("a = ", a)                       // a =  1
	//fmt.Println("&a = ", &a)                     // &a =  0xc420080008
	//fmt.Println("*&a = ", *&a)                   // *&a =  1
	//fmt.Println("b = ", b)                       // b =  0xc420080008
	//fmt.Println("&b = ", &b)                     // &b =  0xc42008a018
	//fmt.Println("*&b = ", *&b)                   // *&b =  0xc420080008
	//fmt.Println("*b = ", *b)                     // *b =  1
	//fmt.Println("c = ", c)                       // c =  0xc42008a018
	//fmt.Println("*c = ", *c)                     // *c =  0xc420080008
	//fmt.Println("&c = ", &c)                     // &c =  0xc42008a020
	//fmt.Println("*&c = ", *&c)                   // *&c =  0xc42008a018
	//fmt.Println("**c = ", **c)                   // **c =  1
	//fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c) // ***&*&*&*&c =  1
	//fmt.Println("x = ", x)                       // x =  1
}
