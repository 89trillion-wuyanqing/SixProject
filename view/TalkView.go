package view

import (
	"SixProject/internal/ctrl"
	"SixProject/internal/model"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

	//service
	model.ServerLabel = widget.NewLabel("service")
	model.ServerEntry = widget.NewEntry()
	model.ServerEntry.SetPlaceHolder("input service")
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
	model.ServerButton.OnTapped = ctrl.ClickButtun(w1)
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
	model.SendButton.OnTapped = ctrl.SendMassage(w1)
	lastBox2 := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(400, 100)), model.SendButton)
	lastBox := container.NewHBox(lastBox1, lastBox2)

	content1 := container.NewVBox(usernameBox, serverBox, bodyBox, lastBox)
	w.SetContent(content1)

	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}
