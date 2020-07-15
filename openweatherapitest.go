package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"apitest-go/config"
)

//json変換に必要なstruct
type WeatherAPI struct {
	Weather []struct {
		Description string `json: "description"`
	} `json: "weather"`
	Main struct {
		Temp     float64 `json: "temp`
		Humidity int     `json: "humidity"`
		Temp_Max float64 `json: "temp_max"`
	} `json: "Main"`
	Name string `json: "name"`
}

//ケルビンを摂氏に変換
func changeKelvin(v float64) int {
	k := v - 273.15
	//c := strconv.FormatFloat(k, 'f', 0, 64)
	c := int(k)
	return c
}

func main() {
	//TownIDとAPIKEYをconfig.iniから取得
	townid := config.Config.Town_Id
	apikey := config.Config.Api_Key
	//Current weather dataのベースURL
	baseurl := "https://api.openweathermap.org/data/2.5/weather?id="
	//API call作成
	url := baseurl + townid + "&appid=" + apikey + "&lang=ja"

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Not access this url: %v", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var data WeatherAPI

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
	}
	//天気の詳細
	tenki := data.Weather[0].Description
	//直近の気温
	t := data.Main.Temp
	//湿度
	h := data.Main.Humidity
	//直近の最高気温
	tmax := data.Main.Temp_Max
	//都市の名前
	cname := data.Name
	//ケルビンを摂氏に変換
	tc := changeKelvin(t)
	tmaxc := changeKelvin(tmax)

	fmt.Printf("今日の%vの天気は%v\n", cname, tenki)
	fmt.Printf("気温%v度、湿度は%v%%\n", tc, h)
	//25度以下のとき
	if tmaxc <= 25 {
		fmt.Printf("最高気温は%v度で肌寒いでしょう\n", tmaxc)
		//25度より高いとき
	} else if tmaxc > 25 {
		fmt.Printf("最高気温は%v度で暑いでしょう\n", tmaxc)
	}
}
