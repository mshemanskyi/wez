package main

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TimezoneID     string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type Current struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMb       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipMm         float64   `json:"precip_mm"`
	PrecipIn         float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelslikeC       float64   `json:"feelslike_c"`
	FeelslikeF       float64   `json:"feelslike_f"`
	VisKm            float64   `json:"vis_km"`
	VisMiles         float64   `json:"vis_miles"`
	UV               float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

type Day struct {
	MaxTempC          float64   `json:"maxtemp_c"`
	MaxTempF          float64   `json:"maxtemp_f"`
	MinTempC          float64   `json:"mintemp_c"`
	MinTempF          float64   `json:"mintemp_f"`
	AvgTempC          float64   `json:"avgtemp_c"`
	AvgTempF          float64   `json:"avgtemp_f"`
	MaxWindMph        float64   `json:"maxwind_mph"`
	MaxWindKph        float64   `json:"maxwind_kph"`
	TotalPrecipMm     float64   `json:"totalprecip_mm"`
	TotalPrecipIn     float64   `json:"totalprecip_in"`
	TotalSnowCm       float64   `json:"totalsnow_cm"`
	AvgVisKm          float64   `json:"avgvis_km"`
	AvgVisMiles       float64   `json:"avgvis_miles"`
	AvgHumidity       float64   `json:"avghumidity"`
	DailyWillItRain   int       `json:"daily_will_it_rain"`
	DailyChanceOfRain int       `json:"daily_chance_of_rain"`
	DailyWillItSnow   int       `json:"daily_will_it_snow"`
	DailyChanceOfSnow int       `json:"daily_chance_of_snow"`
	Condition         Condition `json:"condition"`
	UV                float64   `json:"uv"`
}

type Hour struct {
	TimeEpoch    int64     `json:"time_epoch"`
	Time         string    `json:"time"`
	TempC        float64   `json:"temp_c"`
	TempF        float64   `json:"temp_f"`
	IsDay        int       `json:"is_day"`
	Condition    Condition `json:"condition"`
	WindMph      float64   `json:"wind_mph"`
	WindKph      float64   `json:"wind_kph"`
	WindDegree   int       `json:"wind_degree"`
	WindDir      string    `json:"wind_dir"`
	PressureMb   float64   `json:"pressure_mb"`
	PressureIn   float64   `json:"pressure_in"`
	PrecipMm     float64   `json:"precip_mm"`
	PrecipIn     float64   `json:"precip_in"`
	Humidity     int       `json:"humidity"`
	Cloud        int       `json:"cloud"`
	FeelslikeC   float64   `json:"feelslike_c"`
	FeelslikeF   float64   `json:"feelslike_f"`
	WindChillC   float64   `json:"windchill_c"`
	WindChillF   float64   `json:"windchill_f"`
	HeatIndexC   float64   `json:"heatindex_c"`
	HeatIndexF   float64   `json:"heatindex_f"`
	DewpointC    float64   `json:"dewpoint_c"`
	DewpointF    float64   `json:"dewpoint_f"`
	WillItRain   int       `json:"will_it_rain"`
	ChanceOfRain int       `json:"chance_of_rain"`
	WillItSnow   int       `json:"will_it_snow"`
	ChanceOfSnow int       `json:"chance_of_snow"`
	VisKm        float64   `json:"vis_km"`
	VisMiles     float64   `json:"vis_miles"`
	GustMph      float64   `json:"gust_mph"`
	GustKph      float64   `json:"gust_kph"`
	UV           float64   `json:"uv"`
}

type ForecastDay struct {
	Date      string `json:"date"`
	DateEpoch int64  `json:"date_epoch"`
	Day       Day    `json:"day"`
	Hour      []Hour `json:"hour"`
}

type Forecast struct {
	ForecastDay []ForecastDay `json:"forecastday"`
}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type opts struct {
	ApiKey   string `long:"apikey" description:"Weather api key" required:"true"`
	Location string `short:"l" long:"location" description:"Location: city" default:"Chernivtsi" required:"false"`
	Forecast string `short:"f" long:"forecast" description:"Forecast days" default:"1" required:"false"`
	Weather  string `short:"w" long:"weather" description:"Get weather" required:"false"`
}
