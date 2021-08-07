package view

import (
	"SixProject/internal/handler"
	"SixProject/internal/model"
	"SixProject/internal/server"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func InitView() {
	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())
	w := app.NewWindow("聊天室")
	var w1 = w
	fr1 := fyne.NewSize(100, 40)
	fr2 := fyne.NewSize(200, 40)

	//用户名
	model.UsernameLable = widget.NewLabel("username")
	//输入框
	model.UsernameEntry = widget.NewEntry()
	model.UsernameEntry.SetPlaceHolder("input username")

	usernameBox1 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fr1), model.UsernameLable)
	usernameBox2 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fr2), model.UsernameEntry)
	/*usernameEntry.OnChanged= func(s string) {
		fmt.Println(s)
	}*/
	usernameBox := container.NewHBox(usernameBox1, usernameBox2)

	//server
	model.ServerLabel = widget.NewLabel("server")
	model.ServerEntry = widget.NewEntry()
	model.ServerEntry.SetPlaceHolder("input server")
	model.PortEntry = widget.NewEntry()
	model.PortEntry.SetPlaceHolder("input port")
	serverBox1 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fr1), model.ServerLabel)
	serverBox2 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fr2), model.ServerEntry)
	serverBox3 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fr2), model.PortEntry)
	//连接状态
	model.StatusLabel1 = widget.NewLabel("connect status:")
	model.StatusLabel2 = widget.NewLabel("")

	//列表
	model.UserListLabel = widget.NewLabel("userList:")
	model.UserList = container.NewVBox()

	//连接按钮
	model.ServerButton = &widget.Button{}
	model.ServerButton.Text = "connect"
	model.ServerButton.OnTapped = clickButtun(w1)
	serverBox4 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fr1), model.ServerButton)
	serverBox5 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fr2), model.StatusLabel1)
	serverBox6 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fr1), model.StatusLabel2)
	serverBox := container.NewHBox(serverBox1, serverBox2, serverBox3, serverBox4, serverBox5, serverBox6)

	//聊天主体
	model.TalkList = container.NewVBox()

	model.Src = container.NewScroll(model.UserList)

	model.TalkSrc = container.NewScroll(model.TalkList)

	//fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(400,660)),src)
	talkBox11 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(400, 60)), model.UserListLabel)
	talkBox12 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(400, 600)), model.Src)
	userMaxList := container.NewVBox(talkBox11, talkBox12)
	talkBox2 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(600, 660)), model.TalkSrc)
	talkBox1 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(400, 660)), userMaxList)

	bodyBox := container.NewHBox(talkBox1, talkBox2)
	//bodyBox := widget.NewHBox(contain,multiEntry)

	//尾部
	model.MessageEntry = widget.NewMultiLineEntry()
	model.MessageEntry.SetPlaceHolder("input Message")
	lastBox1 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(600, 100)), model.MessageEntry)

	model.SendButton = &widget.Button{}
	model.SendButton.Text = "Send"
	model.SendButton.OnTapped = SendMassage()
	lastBox2 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(400, 100)), model.SendButton)
	lastBox := container.NewHBox(lastBox1, lastBox2)

	content1 := container.NewVBox(usernameBox, serverBox, bodyBox, lastBox)
	w.SetContent(content1)
	//sendButton.OnTapped=clickButtun(serverButton,statusLabel2,userList)
	/*content := container.NewVBox(
		widget.NewLabel("The top row of the VBox"),
		container.NewHBox(
			widget.NewLabel("Label 1"),
			widget.NewLabel("Label 2"),
		),
	)

	content.Add(widget.NewButton("Add more items", func() {
		content.Add(widget.NewLabel("Added"))
	}))
	*/
	//w.SetContent(content)
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}

func clickButtun(w fyne.Window) func() {
	return func() {
		if model.UsernameEntry.Text == "" {
			//dialog.ShowInformation("Info", "Please input username", w)
			dialog.ShowError(errors.New("Please input username"), w)
			return
		}
		model.Username = model.UsernameEntry.Text
		if model.ServerEntry.Text == "" {
			//dialog.ShowInformation("Info", "Please input server", w)
			dialog.ShowError(errors.New("Please input server"), w)
			return
		}
		model.Server = model.ServerEntry.Text
		if model.PortEntry.Text == "" {
			//dialog.ShowInformation("Info", "Please input port", w)
			dialog.ShowError(errors.New("Please input port"), w)
			return
		}
		model.Port = model.PortEntry.Text

		if model.ServerButton.Text == "connect" {

			//开始websocket连接
			handler.Login()

		} else {
			handler.Exit()

		}
	}
}

func SendMassage() func() {
	return func() {

		if server.UserClient.Socket != nil {

			handler.SendMsg()
			model.MessageEntry.Text = ""
			model.MessageEntry.Refresh()
		}
		return
	}

}

func makeListTab() *widget.List {
	data := []string{"wuaynqing", "www", "list"}

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(nil), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			for i, v := range data {
				if i == id {
					item.(*fyne.Container).Objects[1].(*widget.Label).SetText(v)
				}

			}

		},
	)

	return list
}
