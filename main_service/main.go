package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/KouKouChan/YuriCore/main_service/start"
	"github.com/KouKouChan/YuriCore/utils"
)

var (
	// SERVERVERSION 版本号
	SERVERVERSION = "v1.3"
)

func main() {
	fmt.Println("YuriCore Server", SERVERVERSION)
	fmt.Println("Initializing process ...")

	ExePath, err := utils.GetExePath()
	if err != nil {
		panic(err)
	}

	start.Init(ExePath)

	go initTCP()
	go initUDP()

	ch := make(chan os.Signal)
	defer close(ch)
	signal.Notify(ch, syscall.SIGINT)
	_ = <-ch

	fmt.Println("Press CTRL+C again to close server")

	signal.Notify(ch, syscall.SIGINT)
	_ = <-ch
}
