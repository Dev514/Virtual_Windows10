package main
import (
	"image/color"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	//"path/filepath"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
	"strings"
	"fmt"
	"io/ioutil"
	"fyne.io/fyne/v2/layout"
)
func showGallery(){
	
	w:=myApp.NewWindow("Photos")
	r,_:=LoadResourceFromPath("logos\\gallery_logo.png")
    w.SetIcon(r)
	w.Resize(fyne.NewSize(1200,600))
	root_src:="D:\\Picture"
	files,err:=ioutil.ReadDir(root_src)
	if err!=nil{
		panic(err)
	}
	tabs:= container.NewAppTabs()
 	for _,file:=range files{
		fmt.Println(file.Name(),file.IsDir())
		if file.IsDir()==false{
			extension:=strings.Split(file.Name(),".")[1]
			if extension=="png" || extension=="jpg"{
				image:=canvas.NewImageFromFile(root_src+"\\"+file.Name())
				
				tabs.Append(container.NewTabItem(file.Name(),image))
			}
		}
	}
	
	tabs.SetTabLocation(container.TabLocationLeading )
	color1:=color.NRGBA{R: 255, G:204, B:102, A: 255}
    color2:=color.NRGBA{R:255,G:153,B:51,A:255}
	gradient:= canvas.NewLinearGradient(color2,color1,135)
    container:=container.New(
        layout.NewMaxLayout(),
        gradient,
        tabs,
    )
	w.SetContent(container)
	w.Show()
}

