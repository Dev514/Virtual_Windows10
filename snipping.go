package main
import "C"
import (
  "fmt"
  "math"
  "github.com/go-vgo/robotgo"
  "github.com/kbinani/screenshot"
  "image/png"
  "strconv"
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/canvas"
  "fyne.io/fyne/v2/dialog"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/widget"
   "image"
)

func mouseClick() *image.RGBA{
	var x1,y1,x2,y2 int
	lmk := robotgo.AddMouse("left" )
	if(lmk){
		fmt.Println("left Click")
		x1, y1 = robotgo.GetMousePos()
		fmt.Println("left pos: ", x1, y1)

	}

	lmk2 := robotgo.AddMouse("left" )
	if(lmk2){
		fmt.Println("left Click")
		x2, y2 = robotgo.GetMousePos()
		fmt.Println("left pos: ", x2, y2)

	}
    

	cropBoxWidth := int(math.Abs(float64(x1) - float64(x2)))
	 cropBoxHeight := int(math.Abs(float64(y1) - float64(y2)))
	 cropBoxX := int(math.Min(float64(x1), float64(x2)))
	 cropBoxY := int(math.Min(float64(y1), float64(y2)))
	 img, err := screenshot.Capture(cropBoxX,cropBoxY,cropBoxWidth,cropBoxHeight)
	 if err != nil {
		panic(err)
	}

	return img
}


func showSnipping() {
	var count int=0
	var imgs image.Image=nil
	dup:=false
	
	w:=myApp.NewWindow("Snipping Tool")
    w.Resize(fyne.NewSize(500,500))
	r,_:=LoadResourceFromPath("logos\\snipping_logo.png")
    w.SetIcon(r)
	
	content:=container.NewVBox()
	newbutton := widget.NewButton("New", func() {
		if(dup){
			content.Objects=content.Objects[:len(content.Objects)-1]
			w.Resize(fyne.NewSize(500,500))
		}
	  imgs=mouseClick()
	
	imgtab:=canvas.NewImageFromImage(imgs)
	imgtab.FillMode=canvas.ImageFillOriginal
	content.Add(imgtab)
	dup=true

	})

	savebutton:=widget.NewButton("Save",func(){
		
		saveFile:=dialog.NewFileSave(
			func(uc fyne.URIWriteCloser,_ error){
				 png.Encode(uc,imgs)
				
			},w )
			saveFile.SetFileName("screenshot"+strconv.Itoa(count)+".png")
			count++
			saveFile.Show()
	})

	buttons:=container.NewGridWithColumns(2,
	newbutton,
	savebutton,
	)
	content.Add(buttons)
	w.SetContent(content)
	w.Show()
}

