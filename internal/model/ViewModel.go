package model

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	Server   string
	Port     string
	Username string
)

var (
	UsernameLable *widget.Label
	UsernameEntry *widget.Entry
	ServerLabel   *widget.Label
	ServerEntry   *widget.Entry
	PortEntry     *widget.Entry
	StatusLabel1  *widget.Label
	StatusLabel2  *widget.Label
	ServerButton  *widget.Button
	UserList      *fyne.Container
	TalkList      *fyne.Container
	UserListLabel *widget.Label
	Src           *container.Scroll
	ListMsg       *widget.Label
	TalkSrc       *container.Scroll
	MessageEntry  *widget.Entry
	SendButton    *widget.Button
)
