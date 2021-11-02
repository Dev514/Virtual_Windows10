package main

import(
	"fmt"
	"strconv"
	"encoding/json"
	"math"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	 "io/ioutil"
	 "image/color"
	// "path/filepath"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"strings"
	"net/http"
	//"fyne.io/fyne/v2/theme"
)
var temperature float64 = 0
   var location string = "Mumbai"
   var woeid int64 = 12586539
   var weather string= "clear"
   var searchUrl string= "https://www.metaweather.com/api/location/search/?query="
   var locationUrl string= "https://www.metaweather.com/api/location/"
	
func showWeather(){
	var containe *fyne.Container
	
	w:=myApp.NewWindow("Weather")
	r,_:=LoadResourceFromPath("logos\\weather_icon.png")
    w.SetIcon(r)
	w.Resize(fyne.NewSize(400,650))
	img_path:="weather_img\\"
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter City...")
	image:=canvas.NewImageFromFile(img_path+weather+".png")
	temp_s := fmt.Sprintf("%.0f"+"°C", temperature)
	label1:=canvas.NewText(temp_s,color.White)
	label1.Alignment=fyne.TextAlignCenter
	label1.TextStyle = fyne.TextStyle{Bold: true}
	label1.TextSize=40
	label2:=canvas.NewText(location,color.White)
	label2.Alignment=fyne.TextAlignCenter
	label2.TextStyle = fyne.TextStyle{Bold: true}
	label2.TextSize=50
	
	
	
		
	con:=container.New(
		layout.NewMaxLayout(),
		image,
		label1,
		container.NewVBox(
			layout.NewSpacer(),
			label2,
		),
		)
		btn:=widget.NewButton("", func() {
			if input.Text!=""{
			fetch(input.Text)
			image=canvas.NewImageFromFile(img_path+weather+".png")
			
			
			temp_s = fmt.Sprintf("%.0f"+"°C", temperature)
			label1=canvas.NewText(temp_s,color.White)
			label1.Alignment=fyne.TextAlignCenter
			 label1.TextStyle = fyne.TextStyle{Bold: true}
			 label1.TextSize=40

			 label2=canvas.NewText(location,color.White)
			label2.Alignment=fyne.TextAlignCenter
			 label2.TextStyle = fyne.TextStyle{Bold: true}
			 label2.TextSize=50
			 con.Objects=con.Objects[:len(con.Objects)-3]
			 con.Add(image)
			 con.Add(label1)
			 con.Add(container.NewVBox(
				layout.NewSpacer(),
				label2,
			),)
			 
			 
			fmt.Println(weather)
			}
		})
        gradient:= canvas.NewLinearGradient(color.White,color.White,45)
		input_cont:=container.New(
			layout.NewMaxLayout(),
			gradient,
			input,
		)
	content := container.NewVBox(input_cont,CustBtn(btn) ,)
	
    containe=container.New(
        layout.NewMaxLayout(),
        con,
		content,
		
    )
	
    w.SetContent(containe)
	fmt.Println(woeid)
	fmt.Println(temperature)
	fmt.Println(weather)
	w.Show()
    
}

func fetch(input string){
	res,err:=http.Get(searchUrl+input)
	if err==nil{
	defer res.Body.Close()
	body,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println(err)
	}else{

	weather,err:=UnmarshalLocation(body)
	if(err!=nil){
		fmt.Println(err)
	}else{
	woeid=weather[0].Woeid
	location=weather[0].Title
	fetch_location() 
	}
	 }
	}
}

 func fetch_location(){
	wans:=strconv.Itoa(int(woeid))
	 loc_res,err:=http.Get(locationUrl+wans)
	 if err!=nil{
		 fmt.Println(err)
	 }else{
	 defer loc_res.Body.Close()
	 body,err:=ioutil.ReadAll(loc_res.Body)
	 if err==nil{
	 weather_det,err:=UnmarshalWeather(body)
	 if err==nil{
	 consol_weather:=weather_det.ConsolidatedWeather
	 data:=consol_weather[0]
	 temperature=math.Round(data.TheTemp)
	 weather=strings.ToLower(strings.ReplaceAll(data.WeatherStateName," ",""))}
	  }
	   }

 }

type Location []LocationElement

func UnmarshalLocation(data []byte) (Location, error) {
	var r Location
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Location) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type LocationElement struct {
	Title        string `json:"title"`        
	LocationType string `json:"location_type"`
	Woeid        int64  `json:"woeid"`        
	LattLong     string `json:"latt_long"`    
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()


func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	ConsolidatedWeather []ConsolidatedWeather `json:"consolidated_weather"`
	Time                string                `json:"time"`                
	SunRise             string                `json:"sun_rise"`            
	SunSet              string                `json:"sun_set"`             
	TimezoneName        string                `json:"timezone_name"`       
	Parent              Parent                `json:"parent"`              
	Sources             []Source              `json:"sources"`             
	Title               string                `json:"title"`               
	LocationType        string                `json:"location_type"`       
	Woeid               int64                 `json:"woeid"`               
	LattLong            string                `json:"latt_long"`           
	Timezone            string                `json:"timezone"`            
}

type ConsolidatedWeather struct {
	ID                   int64   `json:"id"`                    
	WeatherStateName     string  `json:"weather_state_name"`    
	WeatherStateAbbr     string  `json:"weather_state_abbr"`    
	WindDirectionCompass string  `json:"wind_direction_compass"`
	Created              string  `json:"created"`               
	ApplicableDate       string  `json:"applicable_date"`       
	MinTemp              float64 `json:"min_temp"`              
	MaxTemp              float64 `json:"max_temp"`              
	TheTemp              float64 `json:"the_temp"`              
	WindSpeed            float64 `json:"wind_speed"`            
	WindDirection        float64 `json:"wind_direction"`        
	AirPressure          float64 `json:"air_pressure"`          
	Humidity             int64   `json:"humidity"`              
	Visibility           float64 `json:"visibility"`            
	Predictability       int64   `json:"predictability"`        
}

type Parent struct {
	Title        string `json:"title"`        
	LocationType string `json:"location_type"`
	Woeid        int64  `json:"woeid"`        
	LattLong     string `json:"latt_long"`    
}

type Source struct {
	Title     string `json:"title"`     
	Slug      string `json:"slug"`      
	URL       string `json:"url"`       
	CrawlRate int64  `json:"crawl_rate"`
}

func CustBtn(btn *widget.Button) *fyne.Container{
    
    Text:= canvas.NewText("Search",color.White)
	Text.Alignment=fyne.TextAlignCenter
	Text.TextStyle=fyne.TextStyle{Bold:true}
    container:=container.New(
        layout.NewMaxLayout(),
        Text,
        btn,
    )
    return container

}

