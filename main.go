package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"rpctest01/client"
	"rpctest01/server"
	"strconv"
	"sync"
	"time"
)

//main函数
func main() {
	go server.RegisterAndServeOnTcp() //先启动服务端
	time.Sleep(1e9)
	wg := new(sync.WaitGroup) //waitGroup用于阻塞主线程防止提前退出
	callTimes := 10
	wg.Add(callTimes)
	for i := 0; i < callTimes; i++ {
		go func() {
			//使用hello world加一个随机数作为参数
			argString := "hello world " + strconv.Itoa(rand.Int(), nil)
			resultString, err := client.Echo(argString)
			if err != nil {
				log.Fatal("error calling:", err)
			}
			if resultString != argString {
				fmt.Println("error")
			} else {
				fmt.Printf("echo:%s\n", resultString)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// 作者：熊纪元
// 链接：https://juejin.im/post/5c4d7005f265da61223ab198
// 来源：掘金
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
