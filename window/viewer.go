package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	_ "github.com/lengzhao/font/autoload" //这个可以让你识别中文
	"net/url"
	"time"
)

func NewWindow() {
	myApp := app.New()
	// 创建一个窗口对象
	myWindow := myApp.NewWindow("zaneway`s Cert Viewer")
	body := newBody()
	myWindow.SetContent(body)
	myWindow.Resize(fyne.Size{800, 600})
	myWindow.ShowAndRun()

}

func newBody() *fyne.Container {
	// 表头
	url, _ := url.Parse("https://github.com/zaneway/CertViewer")
	link := widget.NewHyperlink("^-^  欢迎访问全球最大的同性交友网站  ^-^", url)
	//超链接显示在中间
	centerLink := container.NewCenter(link)
	//时间显示在最右侧
	rightTime := container.NewHBox(layout.NewSpacer(), refreshTimeSeconds())
	//build tab
	tabs := container.NewAppTabs(container.NewTabItemWithIcon("certificate", theme.InfoIcon(), Structure()), container.NewTabItemWithIcon("asn1", theme.ZoomInIcon(), Asn1Structure()))

	//填充布局
	body := container.NewVBox(
		centerLink,
		rightTime,
		tabs,
	)
	return body

}

const DateTime = "2006-01-02 15:04:05"

func refreshTimeSeconds() *widget.Label {
	//填充当前时间
	nowTime := widget.NewLabel(time.Now().Format(DateTime))
	//异步线程更新时间
	go func() {
		for range time.Tick(time.Second) {
			nowTime.SetText(time.Now().Format(DateTime))
		}
	}()
	return nowTime
}
