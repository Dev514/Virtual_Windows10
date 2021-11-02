package main

import(
	"time"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"image/color"
	//"fyne.io/fyne/v2/theme"
)

type fn func()
var myApp fyne.App=app.New()
var w fyne.Window=myApp.NewWindow("Windows 10")

func showTime(clock *fyne.Container){
	
	fomatted:=time.Now().Format("03:04")
	date:=time.Now().Format("01-02-2006")
	clock.Objects=clock.Objects[: len(clock.Objects)-2]
	clockT:=canvas.NewText(fomatted,color.White)
	clockT.TextSize=12
	dateT:=canvas.NewText(date,color.White)
	dateT.TextSize=11
	clock.Add(clockT)
	clock.Add(dateT)

}

func main(){
	r,_:=LoadResourceFromPath("logos\\windows10.png")
    w.SetIcon(r)
	
	w.Resize(fyne.NewSize(1400,800))
	backimg:=canvas.NewImageFromFile("img4.jpg")

	logo,_:=fyne.LoadResourceFromPath("logos\\windows_logo.png")
	window_icon:=newTappableIcon(logo,"Windows",showNotepad)
	clock:=canvas.NewText("",color.White)
	clock.TextSize=12
	date:=canvas.NewText("",color.White)
	date.TextSize=11
	clockCon:=container.NewVBox(
		clock,
		date,
	)
	showTime(clockCon)
	calc_logo,_:=fyne.LoadResourceFromPath("logos\\calc_logo.png")
	calc_icon:=newTappableIcon(calc_logo,"Calculator",showCalc)

	notepad_logo,_:=fyne.LoadResourceFromPath("logos\\notepad_logo.png")
	notepad_icon:=newTappableIcon(notepad_logo,"Notepad",showNotepad)

	gallery_logo,_:=fyne.LoadResourceFromPath("logos\\gallery_logo.png")
	gallery_icon:=newTappableIcon(gallery_logo,"Gallery",showGallery)

	weather_logo,_:=fyne.LoadResourceFromPath("logos\\weather_icon.png")
	weather_icon:=newTappableIcon(weather_logo,"Weather App",showWeather)

	groove_logo,_:=fyne.LoadResourceFromPath("logos\\music_logo.png")
	groove_icon:=newTappableIcon(groove_logo,"Groove Music",showGroove)

	snipping_logo,_:=fyne.LoadResourceFromPath("logos\\snipping_logo.png")
	snipping_icon:=newTappableIcon(snipping_logo,"Snipping Tool",showSnipping)

	calc_iconc:=newTappableIcon(calc_logo,"Snipping Tool",showCalc)
	notepad_iconc:=newTappableIcon(notepad_logo,"Snipping Tool",showNotepad)
	weather_iconc:=newTappableIcon(weather_logo,"Snipping Tool",showWeather)
	groove_iconc:=newTappableIcon(groove_logo,"Snipping Tool",showGroove)
	snipping_iconc:=newTappableIcon(snipping_logo,"Snipping Tool",showSnipping)
	gallery_iconc:=newTappableIcon(gallery_logo,"Snipping Tool",showGallery)

	snText:=canvas.NewText("Snipping Tool",color.White)
	snText.Alignment=fyne.TextAlignCenter
	snc:=container.NewBorder(nil, snText,nil,nil,snipping_iconc,)

	wText:=canvas.NewText("Weather App",color.White)
	wText.Alignment=fyne.TextAlignCenter
	wac:=container.NewBorder(nil, wText,nil,nil,weather_iconc,)

	nText:=canvas.NewText("Notepad",color.White)
	nText.Alignment=fyne.TextAlignCenter
	nc:=container.NewBorder(nil, nText,nil,nil,notepad_iconc,)

	gmText:=canvas.NewText("Groove Music",color.White)
	gmText.Alignment=fyne.TextAlignCenter
	gmc:=container.NewBorder(nil, gmText,nil,nil,groove_iconc,)

	galText:=canvas.NewText("Gallery",color.White)
	galText.Alignment=fyne.TextAlignCenter
	gc:=container.NewBorder(nil, galText,nil,nil,gallery_iconc,)

	calText:=canvas.NewText("Calculator",color.White)
	calText.Alignment=fyne.TextAlignCenter
	cc:=container.NewBorder(nil,calText,nil,nil,calc_iconc,)


	content:=container.New(
		layout.NewMaxLayout(),
		backimg,
		    container.NewHBox(
			
			container.NewGridWithRows(9,
			snc,
			 gmc,
			 wac,
			 gc,
			 nc,
			 cc,
			 layout.NewSpacer(),
			 layout.NewSpacer(),
			 layout.NewSpacer(),		
		),
		),
		)
	//btn:=widget.NewButton("ksdjsid",func(){})
	
	input:=widget.NewEntry()
	inputCon:=container.New(
		layout.NewMaxLayout(),
		canvas.NewLinearGradient(color.White,color.White,45),
		input,
	)
	input.SetPlaceHolder("Type here to search")
	
	bottom:=container.NewGridWithColumns(2,
		container.NewGridWithColumns(2,
			container.New(
				layout.NewFormLayout(),
				window_icon,
				inputCon,
			),
			container.NewGridWithColumns(6,
				
				calc_icon,
				groove_icon,
				notepad_icon,
				gallery_icon,
				weather_icon,
				snipping_icon,
				
			),
		),
		container.NewGridWithColumns(8,
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
			clockCon,
			
		),
		
		
	)
	gradient:=canvas.NewLinearGradient(color.NRGBA{R:36,G:36,B:38,A:255},color.NRGBA{R:36,G:36,B:39,A:255},45)
	btnc:=container.New(
		layout.NewMaxLayout(),
		gradient,
		bottom,

	)
	go func(){
		t:=time.NewTicker(time.Second)
		for range t.C {
			showTime(clockCon)
		}
	}()
	
	w.SetContent(
		container.NewBorder(nil,btnc,nil,nil,content),
		
	)
	
	w.ShowAndRun()

}

type tappableIcon struct {
	widget.Icon
	 query string
	fun fn
}

func newTappableIcon(res fyne.Resource,que string,fun fn) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)
	icon.query=que
	icon.fun=fun
	return icon
}
func (t *tappableIcon) Tapped(_ *fyne.PointEvent) {
	t.fun()
	fmt.Println("I have been tapped"+t.query)
}

func (t *tappableIcon) TappedSecondary(_ *fyne.PointEvent) {
	
	fmt.Println("thiaoa")
}


type Resource interface {
    Name() string
    Content() []byte
}
func LoadResourceFromPath(path string) (Resource, error) {
    bytes, err := ioutil.ReadFile(filepath.Clean(path))
    if err != nil {
        return nil, err
    }
    name := filepath.Base(path)
    return NewStaticResource(name, bytes), nil
}

func NewStaticResource(name string, content []byte) *StaticResource {
    return &StaticResource{
        StaticName:    name,
        StaticContent: content,
    }
}
func (r *StaticResource) Name() string {
    return r.StaticName
}
func (r *StaticResource) Content() []byte {
    return r.StaticContent
}
type StaticResource struct {
    StaticName    string
    StaticContent []byte
}