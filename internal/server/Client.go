package server

import (
	"SixProject/internal/model"
	"errors"
	"fmt"
	"fyne.io/fyne/v2/widget"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"time"
)

//计时器
var timer = time.AfterFunc(time.Duration(time.Minute*1), TimeOut)

//客户端 Client
type Client struct {

	//用户id
	Username string
	//连接的socket
	Socket *websocket.Conn
	//发送的消息
	Send chan []byte
}

var UserClient Client

//定义客户端结构体的read方法
func (c *Client) Read() {

	defer func() {

		fmt.Println("读方法中 ，我要关闭了")
		if e := recover(); e != nil {
			model.SendButton.Disable()
			model.StatusLabel2.Text = "connect fail"
			model.StatusLabel2.Refresh()
			model.ServerButton.Text = "connect"
			model.ServerButton.Refresh()
		}
		c.Socket.Close()

	}()

	for {

		//读取消息
		_, message, err := c.Socket.ReadMessage()
		fmt.Println("读到消息")

		//如果有错误信息，就注销这个连接然后关闭
		if err != nil {
			log.Println(err)
			fmt.Println("有错误，关闭")
			log.Println("ERROR:用户" + c.Username + "在读送消息时，连接已关闭，发生错误，退出关闭")
			panic(errors.New("连接出错，关闭连接"))
			break
		}
		//如果没有错误信息就把信息放入broadcast
		//jsonMessage, _ := proto.Marshal(&model.GeneralReward{Msg: string(message),Type: 1})
		//var generalMsg = &model.GeneralReward{}
		var generalMsg = &model.GeneralReward{}
		e := proto.Unmarshal(message, generalMsg)
		if e != nil {
			log.Println("ERROR:用户" + c.Username + "在读送消息时，protobuf反序列化消息出错，关闭连接返回")
			return
		}
		log.Println("INFO:用户" + c.Username + "收到一条消息，Msg:" + generalMsg.Msg + ",Type:" + strconv.Itoa(int(generalMsg.Type)))

		if generalMsg.Type == 1 || generalMsg.Type == 4 || generalMsg.Type == 3 {

			model.TalkList.Add(widget.NewLabel(generalMsg.Username + ":" + generalMsg.Msg))
			model.TalkList.Refresh()
			continue
		}

		if generalMsg.Type == 2 {
			timer.Reset(time.Duration(time.Minute * 1))
			//fmt.Println("收到pong")
			continue
		}

		if generalMsg.Type == 5 {
			//fmt.Println("接收到list消息")

			//fmt.Println("收到list："+generalMsg.Msg)
			model.UserList.Remove(model.ListMsg)
			model.ListMsg = widget.NewLabel(generalMsg.Msg)
			model.UserList.Add(model.ListMsg)
			model.UserList.Refresh()
			continue
		}

	}
}

func (c *Client) Write() {
	//fmt.Println("我要写东西")
	defer func() {
		fmt.Println("写方法中 ，我要关闭了")
		log.Println("INFO:用户" + model.Username + "关闭websocket连接")
		UserClient.Socket.Close()
	}()

	for {

		/*select {
		case message,_:= <- c.Send:
			var generalMsg = &model.Msg{}
			json.Unmarshal(message,generalMsg)
			if generalMsg.Type == 3{
				c.Socket.WriteMessage(websocket.TextMessage, message)
				return
			}
			err :=c.Socket.WriteMessage(websocket.TextMessage, message)
			if (err != nil){
				log.Println(err.Error())
				return
			}
		}*/

		select {

		//从send里读消息
		case message, ok := <-UserClient.Send:
			//如果没有消息
			if !ok {

				/*c.Socket.WriteMessage(websocket.CloseMessage, []byte{
				})*/
				break
			}

			var generalMsg = &model.GeneralReward{}
			err := proto.Unmarshal(message, generalMsg)
			if err != nil {
				log.Println("ERROR:用户" + c.Username + "在发送消息时，protobuf序列化消息出错，关闭连接返回")
				return
			}
			log.Println("INFO:用户" + c.Username + "向服务端发送了一条消息，Msg:" + generalMsg.Msg + ",Type:" + strconv.Itoa(int(generalMsg.Type)))

			if generalMsg.Type == 3 {
				UserClient.Socket.WriteMessage(websocket.TextMessage, message)
				UserClient.Socket.Close()
				model.SendButton.Disable()
				return
			}
			//有消息就写入，发送给web端
			UserClient.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func TimeOut() {
	UserClient.Socket.Close()
	model.StatusLabel2.Text = "disconnect"
	model.StatusLabel2.Refresh()
}

/*func (c *Client) WriteM(){
	defer Wg.Done()
	for {
		fmt.Print("请输入:")
		reader := bufio.NewReader(os.Stdin)
		data, _ := reader.ReadString('\n')
		var generalMsg = &model.Msg{}
		json.Unmarshal([]byte(data),generalMsg)
		fmt.Println("json解析：")
		fmt.Println(generalMsg)
		c.Send	<- []byte(data)
		c.Socket.WriteMessage(1, []byte(data))
	}
}*/
