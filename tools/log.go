package tools

import (
	"fmt"
	"github.com/golog"
	"net/http"
	"net/rpc"
)


var slog *golog.Logger
var clog *golog.Logger

func init(){
	slog = golog.New_ex("[server]",golog.Create_MyWrite_File("./serverlog.txt"))
	clog = golog.New_ex("[client]",golog.Create_MyWrite_File("./clientlog.txt"))
}

type Logtxt struct{
}

type Arg struct{
	Logdata string
}

func (self* Logtxt) ServerWriteLog(arg *Arg,reply *string) error{
	slog.Debugf(arg.Logdata)
	return nil
}

func (self* Logtxt) ClientWriteLog(arg *Arg,reply *string) error{
	clog.Debugf(arg.Logdata)
	return nil
}

func Run_log_rpc(){
	ac := new(Logtxt)
	rpc.Register(ac)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234",nil)
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println("RPC服务已经停止")
}
