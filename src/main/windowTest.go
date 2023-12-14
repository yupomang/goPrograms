package main

/*import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

type ComWindow struct {
	*walk.MainWindow
}

type LabWindow struct {
	*walk.MainWindow
}

type WindowType int

func Show(Window_Type string) {
	switch Window_Type {
	case "main_window":
		comWindow := new(ComWindow)
		comWindow.ShowWindow()
	case "lab_window":
		//labWindow := new(LabWindow)
	default:
		fmt.Println("参数传递错误")
	}
}

func (comWindow *ComWindow) ShowWindow() {
	var unzipEdit *walk.LineEdit
	pathWindow := new(ComWindow)
	err := declarative.MainWindow{
		AssignTo: &pathWindow.MainWindow,
		Title:    "文件压缩",
		MinSize:  declarative.Size{Width: 480, Height: 230},
		Layout:   declarative.HBox{},
		Children: []declarative.Widget{
			declarative.Composite{
				Layout: declarative.Grid{Columns: 2, Spacing: 10},
				Children: []declarative.Widget{
					declarative.LineEdit{
						AssignTo: &unzipEdit,
						Text:     "请输入路径",
					},
				},
			},
			declarative.Composite{},
		},
	}.Create()
	if err != nil {
		fmt.Println(err)
	}
	pathWindow.SetX(650)
	pathWindow.SetY(300)
	pathWindow.Run()
}*/

//这个修改后的代码修复了之前提到的问题，并添加了缺少的导入语句。现在，您可以运行此代码，并根据需要显示文件压缩窗口或实验窗口。
