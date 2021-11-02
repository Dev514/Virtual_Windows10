package main
import (
    "fmt"
    "os"
    "time"
    "image/color"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
    "github.com/faiface/beep"
    "github.com/faiface/beep/mp3"
    "github.com/faiface/beep/speaker"
    "fyne.io/fyne/v2/layout"
)
var f *os.File
var format beep.Format
var streamer beep.StreamSeekCloser
var pause bool = false
func showGroove() {
    go func(msg string) {
        fmt.Println(msg)
        if streamer == nil {
        } else {
            fmt.Println(fmt.Sprint(streamer.Len()))
        }
    }("going")
    time.Sleep(time.Second)
    
    w := myApp.NewWindow("audio player...")
    w.Resize(fyne.NewSize(400, 400))
    r,_:=LoadResourceFromPath("logos\\music_logo.png")
    w.SetIcon(r)
    logo := canvas.NewImageFromFile("logos\\music_logo.png")
    logo.FillMode = canvas.ImageFillOriginal

    toolbar := widget.NewToolbar(
        widget.NewToolbarSpacer(),
        widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
            speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
            speaker.Play(streamer)
        }),
        widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
            if !pause {
                pause = true
                speaker.Lock()
            } else if pause {
                pause = false
                speaker.Unlock()
            }
        }),
        widget.NewToolbarAction(theme.MediaStopIcon(), func() {
             speaker.Close()
        }),
        widget.NewToolbarSpacer(),
    )
    
    label := widget.NewLabel("Audio MP3..")
    label.Alignment = fyne.TextAlignCenter
    label2 := widget.NewLabel("Play MP3..")
    label2.Alignment = fyne.TextAlignCenter
    browse_files := widget.NewButton("Browse...", func() {
        fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) {
            streamer, format, _ = mp3.Decode(uc)
            label2.Text = uc.URI().Name()
            label2.Refresh()
        }, w)
        fd.Show()
        fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
    })
    color1:=color.NRGBA{R: 255, G:153, B:102, A: 255}
    color2:=color.NRGBA{R:153,G:0,B:255,A:255}
    cstbtn:=CustomBtn(browse_files,color1,color2)
    c := container.NewVBox(label,cstbtn, label2, toolbar)
    w.SetContent(
        container.NewBorder(logo, nil, nil, nil, c),
    )
    w.Show()
}


func CustomBtn(btn *widget.Button,color1, color2 color.Color) *fyne.Container{
    
    gradient:= canvas.NewLinearGradient(color2,color1,45)
    container:=container.New(
        layout.NewMaxLayout(),
        gradient,
        btn,
    )
    return container

}