package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"fyneTest/component"
	"image/color"
)

type MainUi struct {
	Window        fyne.Window
	Toolbar       *fyne.Container
	SidebarTop    *fyne.Container
	SidebarLeft   *fyne.Container
	SidebarRight  *fyne.Container
	SidebarBottom *fyne.Container
	Statusbar     *fyne.Container
	WindowContent *fyne.Container
}

func CreateMainUi(w fyne.Window) *MainUi {
	icoHaitunwan, _ := fyne.LoadResourceFromPath("resources/haitunwan.png")
	
	// top toolbar
	toolbar := container.NewBorder(nil, widget.NewSeparator(), nil, nil,
		widget.NewToolbar(
			widget.NewToolbarAction(icoHaitunwan, func() {
				confirmDialog := dialog.NewConfirm("toolbaritem",
					"toolbaritem tapped",
					func(b bool) {
					
					}, w)
				confirmDialog.Show()
			}),
		),
	)
	
	// sidebarTop as bread
	sidebarTop := container.NewBorder(nil, widget.NewSeparator(), nil, nil,
		container.NewHBox(
			widget.NewButton("top", func() {
				confirmDialog := dialog.NewConfirm("top",
					"top will place toolbar  items",
					func(b bool) {
					
					}, w)
				confirmDialog.Show()
			})),
	)
	top := container.NewVBox(toolbar, sidebarTop)
	
	// left dock
	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)
	
	list := widget.NewListWithData(
		data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
	
	panelLeft := container.NewHBox(
		container.NewBorder(
			container.NewVBox(
				widget.NewLabel("项目"),
				widget.NewSeparator(),
			),
			nil,
			nil,
			nil,
			list,
		),
		widget.NewSeparator(),
	)
	tagBtn := component.NewTagButton(icoHaitunwan, "Go", panelLeft)
	tagBtn2 := component.NewTagButton(icoHaitunwan, "java", panelLeft)
	sidebarLeft := container.NewBorder(nil, nil, nil, widget.NewSeparator(),
		container.NewVBox(tagBtn, tagBtn2))
	
	// right dock
	panelRight := container.NewHBox(
		widget.NewSeparator(),
		container.NewBorder(
			container.NewVBox(
				widget.NewLabel("Profile"),
				widget.NewSeparator(),
			),
			nil,
			nil,
			nil,
			container.NewHBox(widget.NewLabel("panel right")),
		),
	)
	tagBtn3 := component.NewTagButton(icoHaitunwan, "Maven", panelRight)
	tagBtn4 := component.NewTagButton(icoHaitunwan, "Database", panelRight)
	sidebarRight := container.NewBorder(nil, nil, widget.NewSeparator(), nil,
		container.NewVBox(tagBtn3, tagBtn4))
	
	// center
	mainEditor := container.NewMax(
		canvas.NewRectangle(color.RGBA{R: 131, G: 175, B: 155, A: 1}),
		container.NewDocTabs(
			container.NewTabItem("main.go",
				container.NewVBox(
					canvas.NewRectangle(color.RGBA{R: 131, G: 175, B: 155, A: 1.0}),
					widget.NewRichTextWithText("test color and background"),
				),
			),
			container.NewTabItem("TagButton.go", widget.NewLabel("TagButton")),
		),
	)
	panelCenter := container.NewMax(
		container.NewBorder(
			nil,
			widget.NewButton("editor status info ", func() {
			
			}),
			nil,
			nil,
			mainEditor,
		))
	
	// panelBottom
	panelBottom := container.NewVBox(widget.NewLabel("panel bottom"))
	tagBtn5 := component.NewTagButton(icoHaitunwan, "Terminal", panelBottom)
	tagBtn6 := component.NewTagButton(icoHaitunwan, "Terminal", panelBottom)
	sidebarBottom := container.NewHBox(tagBtn5, tagBtn6)
	// bottom statusbar
	statusbar := container.NewHBox(
		widget.NewLabelWithStyle("status info", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	)
	
	sidebarBottomWithSep := container.NewVBox(
		widget.NewSeparator(),
		panelBottom,
		sidebarBottom,
	)
	statusbarWithSep := container.NewVBox(widget.NewSeparator(), statusbar)
	bottom := container.NewVBox(
		sidebarBottomWithSep,
		statusbarWithSep,
	)
	
	center := container.NewBorder(nil, nil, panelLeft, panelRight, panelCenter)
	windowContent := container.NewMax(
		container.NewBorder(top, bottom, sidebarLeft, sidebarRight, center),
	)
	
	mainUi := &MainUi{
		Window:        w,
		Toolbar:       toolbar,
		SidebarTop:    sidebarTop,
		SidebarLeft:   sidebarLeft,
		SidebarRight:  sidebarRight,
		SidebarBottom: sidebarBottomWithSep,
		Statusbar:     statusbarWithSep,
		WindowContent: windowContent,
	}
	
	// 被隐藏的组件所在的容器要刷新
	tagBtn.SetToggleCallback(func() {
		center.Refresh()
	})
	tagBtn2.SetToggleCallback(func() {
		center.Refresh()
	})
	tagBtn3.SetToggleCallback(func() {
		center.Refresh()
	})
	tagBtn4.SetToggleCallback(func() {
		center.Refresh()
	})
	
	// menu
	menu := mainUi.createMenu()
	mainUi.Window.SetMainMenu(menu)
	
	// 默认隐匿
	//panelLeft.Hide()
	panelRight.Hide()
	ToggleContainer(mainUi.Toolbar)
	ToggleContainer(mainUi.SidebarTop)
	ToggleContainer(mainUi.SidebarLeft)
	ToggleContainer(mainUi.SidebarRight)
	ToggleContainer(mainUi.SidebarBottom)
	
	return mainUi
}

func (m *MainUi) createMenu() *fyne.MainMenu {
	
	return fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Open", func() {
				dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
					uri := closer.URI()
					fmt.Println(uri)
				}, m.Window)
			}),
		),
		fyne.NewMenu("View",
			fyne.NewMenuItem("Toggle Toolbar", func() {
				ToggleContainer(m.Toolbar)
			}),
			fyne.NewMenuItem("Toggle Sidebar Top", func() {
				ToggleContainer(m.SidebarTop)
			}),
			fyne.NewMenuItem("Toggle Sidebar Left", func() {
				ToggleContainer(m.SidebarLeft)
				m.WindowContent.Refresh()
			}),
			fyne.NewMenuItem("Toggle Sidebar Right", func() {
				ToggleContainer(m.SidebarRight)
				m.WindowContent.Refresh()
			}),
			fyne.NewMenuItem("Toggle Sidebar Bottom", func() {
				ToggleContainer(m.SidebarBottom)
			}),
			fyne.NewMenuItem("Toggle Statusbar", func() {
				ToggleContainer(m.Statusbar)
			}),
		),
	)
}

func ToggleContainer(c Visible) bool {
	if c.Visible() {
		c.Hide()
		return false
	} else {
		c.Show()
		return true
	}
}

type Visible interface {
	Visible() bool
	Hide()
	Show()
}
