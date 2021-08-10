package ctrl

import (
	"SixProject/internal/model"
	"SixProject/internal/service"
	"SixProject/internal/ws"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func ClickButtun(w fyne.Window) func() {
	return func() {
		if model.UsernameEntry.Text == "" {
			//dialog.ShowInformation("Info", "Please input username", w)
			dialog.ShowError(errors.New("Please input username"), w)
			return
		}
		model.Username = model.UsernameEntry.Text
		if model.ServerEntry.Text == "" {
			//dialog.ShowInformation("Info", "Please input service", w)
			dialog.ShowError(errors.New("Please input service"), w)
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
			service.Login()

		} else {
			service.Exit()

		}
	}
}

func SendMassage(w fyne.Window) func() {
	return func() {

		if ws.UserClient.Socket != nil {
			if model.MessageEntry.Text == "" || model.MessageEntry.Text == " " {
				dialog.ShowError(errors.New("Massage not null"), w)
				return
			}
			service.SendMsg()
			model.MessageEntry.Text = ""
			model.MessageEntry.Refresh()
		} else {

			dialog.ShowError(errors.New("Websocket not connected"), w)
			model.MessageEntry.Text = ""
			model.MessageEntry.Refresh()
		}

	}

}
