package handler

import (
	"SixProject/internal/model"
	"SixProject/internal/server"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"time"
)

//创建websocket长链接
func Login() {

	defer func() {
		if e := recover(); e != nil {
			log.Println("ERROR:websocket连接失败")

		}
	}()
	host := model.Server + ":" + model.Port
	u := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   "ws",
	}

	re := http.Header{}
	re.Add("username", model.Username)

	fmt.Println(u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), re)
	if err != nil {
		fmt.Println("websocket创建失败")
		fmt.Println(err.Error())
		log.Println("ERROR:websocket创建失败")
		model.StatusLabel2.Text = "connect fail"
		model.StatusLabel2.Refresh()
		panic("websocket创建失败")

	}
	log.Println("INFO:用户" + model.Username + "websocket连接成功")
	server.UserClient.Socket = c
	server.UserClient.Username = model.Username
	server.UserClient.Send = make(chan []byte)
	model.StatusLabel2.Text = "connected"
	model.StatusLabel2.Refresh()
	model.ServerButton.Text = "disconnect"
	model.ServerButton.Refresh()
	model.SendButton.Enable()

	//jsonStr ,r:=json.Marshal(model.Msg{Msg: model.Username+"登陆",Username: model.Username,Type: 1})
	/*if r!=nil{
		fmt.Println(r.Error())
	}*/
	//server.UserClient.Send<-jsonStr

	go server.UserClient.Read()
	go server.UserClient.Write()

	Ping()
	UserList()
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for range ticker.C {
			//fmt.Println("3秒打印一次")
			Ping()

		}
		ticker.Stop()
	}()

	go func() {
		ticker := time.NewTicker(4 * time.Second)
		for range ticker.C {
			//fmt.Println("4秒打印一次")

			UserList()
		}
		ticker.Stop()
	}()
}

//发送ping
func Ping() {
	if server.UserClient.Socket != nil {
		jsonStr, _ := proto.Marshal(&model.GeneralReward{Msg: "Ping", Username: model.Username, Type: 2})

		server.UserClient.Send <- jsonStr
	}

}

//发送list消息
func UserList() {
	if server.UserClient.Socket != nil {
		jsonStr, _ := proto.Marshal(&model.GeneralReward{Msg: "是我，你爹！快打钱！", Username: model.Username, Type: 5})
		server.UserClient.Send <- jsonStr
	}

}

//发送消息
func SendMsg() {
	if server.UserClient.Socket != nil {
		jsonStr, _ := proto.Marshal(&model.GeneralReward{Msg: model.MessageEntry.Text, Username: model.Username, Type: 1})

		server.UserClient.Send <- jsonStr
	}

}

//退出
func Exit() {
	if server.UserClient.Socket != nil {
		jsonStr, _ := proto.Marshal(&model.GeneralReward{Msg: model.Username + ":" + "login out", Username: model.Username, Type: 3})
		server.UserClient.Send <- jsonStr
		model.StatusLabel2.Text = "disconnect"
		model.StatusLabel2.Refresh()
		model.ServerButton.Text = "connect"
		model.ServerButton.Refresh()
	}

}
