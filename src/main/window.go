package main

import (
	"archive/zip"
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"io"
	"os"
)

type Window interface {
	ShowWindow() //展示窗体界面
}

// 创建压缩解压缩的界面类
type ComWindow struct {
	Window
	*walk.MainWindow
}

// 展示压缩、解压缩成功失败提示信息的界面类
type LabWindow struct {
	Window
}

// 创建界面类对象
func Show(Window_Type string) {
	var Win Window
	switch Window_Type {
	case "main_window":
		Win = &ComWindow{}
	case "lab_window":
		Win = &LabWindow{}
	default:
		fmt.Println("参数传递错误")
	}
	Win.ShowWindow()
}

var labText *walk.Label //用来展示提示信息的Label
var Text string         //保存提示信息
// 首先实现ShowWindow方法,展示出空白的窗体
func (comWindow *ComWindow) ShowWindow() {
	var unzipEdit *walk.LineEdit        //选择“解压文件”文本框
	var saveUnZipEdit *walk.LineEdit    //解压后文件存放路径文本框
	var unzipButn *walk.PushButton      //选择解压文件按钮
	var saveUnZipButn *walk.PushButton  //创建用来选择解压后文件保存路径的按钮
	var zipEdit *walk.LineEdit          //创建一个用来展示要压缩的文件路径的文本框
	var zipButn *walk.PushButton        //选择要压缩文件的按钮
	var saveZipEdit *walk.LineEdit      //用来展示压缩后的文件的存放路径的文本框
	var saveZipButn *walk.PushButton    //选择压缩后文件存放路径按钮
	var startUnZipButn *walk.PushButton //开始解压按钮
	var startZipButn *walk.PushButton   //开始压缩按钮

	pathWindow := new(ComWindow)
	//var pathWindow ComWindow
	//pathWindow.Window = &WindowImpl{}
	//pathWindow.MainWindow = new(walk.MainWindow)

	err := declarative.MainWindow{
		AssignTo: &pathWindow.MainWindow, //关联主窗体
		Title:    "文件压缩",                 //窗口标题名称
		//MinSize:  declarative.Size{Width: 480, Height: 230}, //指定窗口的宽度，高度
		Bounds: declarative.Rectangle{X: 250, Y: 360, Width: 480, Height: 230},
		//布局
		Layout: declarative.HBox{}, //水平布局
		Children: []declarative.Widget{
			//左边区域
			declarative.Composite{
				Layout: declarative.Grid{Columns: 2, Spacing: 10}, //左边区域分为两列布局
				Children: []declarative.Widget{
					declarative.LineEdit{ //表示的文本框
						AssignTo: &unzipEdit, //将创建好的文本框与变量关联，后面可以更具该变量获取文本框中的值
						Text:     "请输入路径",
					},
					//添加选择解压文件的按钮
					declarative.PushButton{
						AssignTo: &unzipButn,
						Text:     "选择解压文件",
						OnClicked: func() {
							//匿名函数
							//fmt.Println(unzipButn.Text())
							//弹出选择文件对话框
							filePath := pathWindow.OpenFileManager()
							unzipEdit.SetText(filePath) //将返回的文件路径赋值给文本框
							//fmt.Println(filePath)
						},
					},
					// 创建展示解压后文件存放路径的文本框
					declarative.LineEdit{
						AssignTo:    &saveUnZipEdit,
						Text:        "解压后文件的路径",
						ToolTipText: "请输入解压后文件的存放路径",
					},
					//创建用来选择解压后文件存放路径的按钮
					declarative.PushButton{
						AssignTo: &saveUnZipButn,
						Text:     "选择保存路径",
						OnClicked: func() {
							filePath := pathWindow.OpenDirManager()
							saveUnZipEdit.SetText(filePath + "\\")
						},
					},
					//创建一个用来展示要压缩的文件路径的文本框
					declarative.LineEdit{
						AssignTo: &zipEdit,
						Text:     "选择要压缩的文件路径",
					},
					// 添加压缩文件按钮
					declarative.PushButton{
						AssignTo: &zipButn,
						Text:     "选择压缩文件",
						OnClicked: func() {
							filePath := pathWindow.OpenFileManager()
							zipEdit.SetText(filePath)
						},
					},
					//添加压缩好的文件存放路径展示的文本框
					declarative.LineEdit{
						AssignTo: &saveZipEdit,
						Text:     "请输入要压缩的文件存放路径",
					},
					//创建一个选择压缩后文件路径的按钮
					declarative.PushButton{
						AssignTo: &saveZipButn,
						Text:     "选择保存",
						OnClicked: func() {
							filePath := pathWindow.OpenDirManager()
							saveZipEdit.SetText(filePath + "\\")
						},
					},
					//用来展示压缩与解压缩后相应的提示信息
					declarative.Label{
						AssignTo: &labText,
						Text:     "",
					},
				},
			},
			//右边区域
			declarative.Composite{
				Layout: declarative.Grid{Rows: 2, Spacing: 40},
				Children: []declarative.Widget{
					declarative.PushButton{
						AssignTo: &startUnZipButn,
						Text:     "开始解压",
						OnClicked: func() {
							//解压文件，传递了压缩文件的路径和解压后文件存放的路径
							pathWindow.StartToUnZip(unzipEdit.Text(), saveUnZipEdit.Text())
							Text = "文件解压成功"
							Show("lab_window")
						},
					},
					declarative.PushButton{
						AssignTo: &startZipButn,
						Text:     "开始压缩",
						OnClicked: func() {
							pathWindow.StartToZip(zipEdit.Text(), saveZipEdit.Text())
						},
					},
				},
			},
		},
	}.Create() //创建窗口
	if err != nil {
		fmt.Println(err)
	}
	//窗口的展示，需要通过坐标来指定
	pathWindow.SetX(650) //X坐标
	pathWindow.SetY(300) //Y坐标
	pathWindow.Run()     //启动窗口
}

// 打开文件选择对话框
func (mv *ComWindow) OpenFileManager() (filePath string) {
	//1: 创建文件对话框的对象
	dlg := new(walk.FileDialog)
	dlg.Title = "选择文件"
	dlg.Filter = "所有文档(*.*)|*.*|文本文档(*.txt)|*.txt"
	//2: 打开文件对话框
	b, err := dlg.ShowOpen(mv) //如果单击对话框中的“打开”按钮，返回true，否则返回false
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	//3: 获取选中的文件
	filePath = dlg.FilePath //获取选中的文件的路径
	return filePath
}

// 打开浏览文件夹的窗口
func (mv *ComWindow) OpenDirManager() (filePath string) {
	//1: 创建对话框的对像
	dlg := new(walk.FileDialog)
	//2: 打开窗口
	_, err := dlg.ShowBrowseFolder(mv) // 展示浏览文件夹的窗口
	if err != nil {
		fmt.Println(err)
	}
	//3: 获取选中的路径，并且返回
	filePath = dlg.FilePath
	return filePath
}

// 实现文件解压操作
func (mv *ComWindow) StartToUnZip(file string, saveFile string) {
	//1: 获取第一个文本框中，要解压的文件的路径，并且读取压缩文件中的内容
	reader, err := zip.OpenReader(file)
	if err != nil {
		fmt.Println(err)
	}
	defer reader.Close()
	//2：循环遍历压缩包中的文件
	for _, file := range reader.File {
		rc, err := file.Open() //打开从压缩文件中获取到的文件或者是文件夹
		if err != nil {
			fmt.Println(err)
		}
		defer rc.Close()
		//构建完整的文件夹或者是文件的存放路径（文件夹或者是文件存放路径+文件夹或者文件名称）
		newName := saveFile + file.Name
		newName, err = UTF8ToGBK(newName)
		if err != nil {
			fmt.Println(err)
		}
		//判断是否为文件夹。IsDir():如果是文件夹，该方法返回的是true
		if file.FileInfo().IsDir() {
			//创建文件夹
			err := os.MkdirAll(newName, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}
		//判断是否为文件
		if !file.FileInfo().IsDir() {
			f, err := os.Create(newName)
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()
			//读取压缩包中文件的内容，然后写入到新创建的文件中
			//read write
			_, err1 := io.Copy(f, rc)
			if err1 != nil {
				fmt.Println(err1)
			}
		}
	}
}

// 将提示信息打印在label上
func (lab *LabWindow) ShowWindow() {
	labText.SetText(Text)
}

// 实现文件压缩操作
func (mv *ComWindow) StartToZip(filePath string, savePath string) {
	//1: 获取第四个文本框中的值，然后创建压缩文件
	d, err := os.Create(savePath)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()
	//2: 获取第三个文本框中的值，打开该文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//1: 获取第四个文本框中的值，然后创建压缩文件
	/*fileNameIndex := strings.Split(file.Name(), "\\")
	fileName := fileNameIndex[len(fileNameIndex)-1]
	currentTime := time.Now()
	timeString := currentTime.Format("20060102150405")
	savePath1 := savePath + fileName*/
	//获取.txt后缀之前的文件路径
	/*lastDotIndex := strings.LastIndex(savePath1, ".")
	fileNameWithoutExt := savePath1[:lastDotIndex]*/
	//获取文件后缀
	/*fileNameParts := strings.Split(savePath1, ".")
	fileExt := fileNameParts[len(fileNameParts)-1]*/
	/*savePath2 := fmt.Sprintf("%s_%s.%s", fileNameWithoutExt, timeString, "zip")
	d, err := os.Create(savePath2)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()*/
	//3: 将压缩的文件写入到压缩包中
	//3.1 要获取要压缩的文件的信息
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		fmt.Println(err)
	}
	//3.2 要将压缩的文件写入到压缩包中
	w := zip.NewWriter(d) //更具创建的压缩包，创建了一个Writer指针，通过该指针可以对压缩包进行操作
	defer w.Close()
	writer, err := w.CreateHeader(header)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(writer, file)

}
