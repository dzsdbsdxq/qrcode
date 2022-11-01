package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// App struct
type App struct {
	ctx context.Context
}
type QrLogoData struct {
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) OpenLogoAlertDialog() QrLogoData {
	filters := []runtime.FileFilter{
		{
			DisplayName: "Images Files (*.jpg;*.png)",
			Pattern:     "*.jpg;*.png",
		},
	}
	dialogOptions := runtime.OpenDialogOptions{
		DefaultDirectory:           "",
		DefaultFilename:            "",
		Title:                      "选择图片",
		Filters:                    filters,
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: true,
	}

	filePath, err := runtime.OpenFileDialog(a.ctx, dialogOptions)

	if err != nil {
		return QrLogoData{FileName: "", FilePath: ""}
	}
	if filePath == "" {
		return QrLogoData{FileName: "", FilePath: ""}
	}
	return QrLogoData{
		FilePath: doImageToBase(filePath),
		FileName: path.Base(filePath),
	}
}

func (a *App) OfficeQrGenerateSave(data string) error {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}
	dialogOptions := runtime.SaveDialogOptions{
		DefaultDirectory: path,
		DefaultFilename:  "output.png",
		Title:            "选择保存文件",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "图片文件",
				Pattern:     "*.png",
			},
		},
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: true,
	}
	filePath, _ := runtime.SaveFileDialog(a.ctx, dialogOptions)

	//成图片文件并把文件写入到buffer
	image, _ := base64.StdEncoding.DecodeString(data)

	return ioutil.WriteFile(filePath, image, 0666)
}

// DoImageToBase 图片转base64
func doImageToBase(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	//获取图片类型
	pytpe := ext(filePath)
	BaseHead := fmt.Sprintf("data:image/%s;base64", pytpe)
	str := fmt.Sprintf("%s,%s", BaseHead, base64.StdEncoding.EncodeToString(b))
	return str
}

// Ext 自定义获取后缀名
func ext(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return ""
}
