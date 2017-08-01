package main

import (
	"github.com/badforlabor/cellnet"
	"github.com/badforlabor/cellnet/proto/gamedef"
	"github.com/badforlabor/cellnet/socket"
	"os"
	"os/signal"
	log "github.com/badforlabor/glog"
	"flag"
	"strings"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
	"github.com/badforlabor/cellnet/proto/newprotocol"
)
/*
	第三方包：
	go get github.com/badforlabor/cellnet
	go get github.com/badforlabor/glog
	go get github.com/davyxu/golog
	go get github.com/golang/protobuf/proto
	go get github.com/davyxu/pbmeta

	protobuf依赖：
		下载一个protoc （https://github.com/google/protobuf/releases）
		go install github.com/golang/protobuf/protoc-gen-go
		go get github.com/golang/protobuf

*/

//var log *golog.Logger = golog.New("test")

func main() {

	flag.Parse()

	// 创建一个log目录，与exe文件同级别
	dir := getCurPath()
	dir = strings.Join([]string{dir, "/log"}, "")
	os.Mkdir(dir, os.ModeDir)

	// 设置日志大小，日志路径
	log.MaxSize = 1024 * 1024
	log.Cheat(dir)
	// 设置关闭程序的时候，flush
	defer log.Flush()

	queue := cellnet.NewEventQueue()

	evd := socket.NewAcceptor(queue).Start("127.0.0.1:7201")

	socket.NewRegisterMessage(evd, newprotocol.PID_TestEchoACK, func(content interface{}, ses cellnet.Session) {
		msg := content.(*newprotocol.TestEchoACK)

		log.Infoln("server recv:", msg.Content)

		ses.Send(&newprotocol.TestEchoACK{
			Content: msg.Content,
		})

	})

	time.Sleep(3 * time.Second)
	queue.StartLoop()

	fmt.Println("1111111111111111")
	client()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Infoln("exit, sig:", sig)
}

func getCurPath() string {
	fmt.Println(os.Args[0])
	file, _ := exec.LookPath(os.Args[0])

	if len(file) == 0 {
		file = os.Args[0]
	}

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)

	//将全路径用\\分割，得到4段，①E: ②golang ③test ④a.exe
	splitstring := strings.Split(path, "\\")

	//size为4
	size := len(splitstring)

	//将全路径用最后一段(④a.exe)进行分割，得到2段，①E:\\golang\\test\\ ②a.exe
	splitstring = strings.Split(path, splitstring[size-1])

	//将①(E:\\golang\\test\\)中的\\替换为/，最终得到结果E:/golang/test/
	rst := strings.Replace(splitstring[0], "\\", "/", size-1)
	return rst
}

func client() {

	queue := cellnet.NewEventQueue()

	evd := socket.NewConnector(queue).Start("127.0.0.1:7201")

	socket.NewRegisterMessage(evd, newprotocol.PID_TestEchoACK, func(content interface{}, ses cellnet.Session) {
		msg := content.(*newprotocol.TestEchoACK)

		log.Infoln("client recv:", msg.Content)
	})


	socket.NewRegisterMessage(evd, socket.Event_SessionConnected, func(content interface{}, ses cellnet.Session) {

		ses.Send(&newprotocol.TestEchoACK{
			Content: "hello",
		})

	})

	socket.NewRegisterMessage(evd,  socket.Event_SessionConnectFailed, func(content interface{}, ses cellnet.Session) {

		msg := content.(*gamedef.SessionConnectFailed)

		log.Infoln(msg.Reason)

	})

	queue.StartLoop()

	//signal.WaitAndExpect(1, "not recv data")

}