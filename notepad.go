package main

import(
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	 "fyne.io/fyne/v2/widget"
	 //"path/filepath"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
	"image/color"
	"io/ioutil"
	"fyne.io/fyne/v2/layout"
	"strconv"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)
var count int=0
func showNotepad(){
	
	//a:=app.New()
	w:=myApp.NewWindow("Notepad")
	r,_:=LoadResourceFromPath("logos\\notepad_logo.png")
    w.SetIcon(r)
	w.Resize(fyne.NewSize(500,600))
	input:=widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text here .....")

	newBtn:=widget.NewButton("New",func(){
		input.SetText("")
		count++

	})
   
	editBtn:=widget.NewButton("Edit",func(){
		openFile:=dialog.NewFileOpen(
			func(read fyne.URIReadCloser,_ error){
				ReadData,_:=ioutil.ReadAll(read)
				output:= fyne.NewStaticResource("Notepad",ReadData)
				input.SetText(string(output.StaticContent))
				
			},w)
			openFile.SetFilter(
				storage.NewExtensionFileFilter([] string{".txt"}),
			)
			openFile.Show()
	})

	saveBtn:=widget.NewButton("Save",func(){
		saveFile:=dialog.NewFileSave(
			func(uc fyne.URIWriteCloser,_ error){
				textData:= []byte(input.Text)
				uc.Write(textData)
			},w )
			saveFile.SetFileName("Untitled"+strconv.Itoa(count)+".txt")
			count++
			saveFile.Show()
	})


	openBtn:=widget.NewButton("Open",func(){
		openFile:=dialog.NewFileOpen(
			func(read fyne.URIReadCloser,_ error){
				ReadData,_:=ioutil.ReadAll(read)
				output:= fyne.NewStaticResource("Notepad",ReadData)

				viewData:=widget.NewMultiLineEntry()
				viewData.SetText(string(output.StaticContent))

				
				w:=fyne.CurrentApp().NewWindow(
					string(output.StaticName))
					btnc:=container.New(layout.NewMaxLayout(),
					canvas.NewLinearGradient(color.NRGBA{R:24,G:153,B:204,A:255},color.NRGBA{R:24,G:153,B:204,A:255},45),
					widget.NewButton("Save",func(){
						saveFile:=dialog.NewFileSave(
							func(uc fyne.URIWriteCloser,_ error){
								textData:= []byte(input.Text)
								uc.Write(textData)
							},w )
							saveFile.SetFileName("Untitled"+strconv.Itoa(count)+".txt")
							count++
							saveFile.Show()
					}))
					w.SetContent(container.NewBorder(nil,btnc,nil,nil,container.NewScroll(viewData)))
					w.Resize(fyne.NewSize(500,600))
					w.Show()
			},w)
			openFile.SetFilter(
				storage.NewExtensionFileFilter([] string{".txt"}),
			)
			openFile.Show()
	}) 


	
	btnContainer:=container.NewHBox(
		newBtn,
		openBtn,
		editBtn,
		saveBtn,
		)
    
	Container:=container.New(
		layout.NewMaxLayout(),
		input,
	)
	w.SetContent(container.NewBorder(btnContainer,nil,nil,nil,Container))
	w.Show()
}

