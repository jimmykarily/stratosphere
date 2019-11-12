package main

import (
	"encoding/json"
	"image"
	"time"

	maps "github.com/flopp/go-staticmaps"
)

type Activity struct {
	Id            int
	Type          int
	Parent        int
	StartTime     int
	EndTime       int
	Calories      float64
	CurrentStatus int
	ContentJSON   string
	PointsData    string
	Content       *Content
	Image         image.Image
}

type Content struct {
	TrackId                    int     `json:"track_id"`
	Calories                   int     `json:"calorie"`
	EndTime                    int     `json:"end_time"`
	StartTime                  int     `json:"start_time"`
	Distance                   int     `json:"distance"`
	SportType                  int     `json:"sport_type"`
	Duration                   int     `json:"duration"`
	PausedTime                 int     `json:"paused _ime"`
	UpdateCount                int     `json:"update_count"`
	ChildList                  []int   `json:"child_list"` // Array of int?
	PauseInfo                  []int   `json:"pause_info"` // Array of int?
	ParentTrackId              int     `json:"parent_trackid"`
	Speed                      float64 `json:"speed"`
	Pace                       float64 `json:"pace"`
	Latitude                   float64 `json:"latitude"`
	Longitude                  float64 `json:"longitude"`
	StepFrequency              float64 `json:"step_freq"`
	AverageHeartRate           int     `json:"avg_heart_rate"`
	MaxHeartRate               int     `json:"max_heart_rate"`
	MinHeartRate               int     `json:"min_heart_rate"`
	AveragePace                float64 `json:"total_pace"`
	MaxPace                    float64 `json:"best_pace"`
	MinPace                    float64 `json:"min_pace"`
	AverageSpeed               float64 `json:"total_speed"`
	MaxSpeed                   float64 `json:"max_speed"`
	MaxStepFrequency           float64 `json:"max_step_freq"`
	ClimbUp                    int     `json:"climb_up"`
	ClimbDown                  int     `json:"climb_down"`
	HighAltitude               float64 `json:"high_altitude"`
	LowAltitude                float64 `json:"low_altitude"`
	KmPerIndividual            int     `json:"km_per_individual"`
	StepCount                  int     `json:"step_count"`
	AverageStride              float64 `json:"avg_stride"`
	TotalClimbDisacend         int     `json:"total_climb_disacend"`
	MaxCadence                 int     `json:"max_cadence"` // Bicycle cadence
	AverageCadence             int     `json:"ave_cadence"`
	DeviceType                 int     `json:"device_type"`
	DeviceSource               int     `json:"device_source"`
	SummaryVersion             int     `json:"summary_version"`
	DescendDistance            int     `json:"descend_distance"`
	AscendTime                 int     `json:"ascend_time"`
	DescendTime                int     `json:"descend_time"`
	Swolf                      int     `json:"swolf"`
	TotalStrokes               int     `json:"total_strokes"` // Swimming?
	TotalTrips                 int     `json:"total_trips"`
	AverageStrokeSpeed         int     `json:"avg_stoke_speed"` // Typo?
	MaxStrokeSpeed             int     `json:"max_stoke_speed"` // Typo?
	AverageDistancePerStroke   int     `json:"avg_dis_per_stroke"`
	SwimPoolLength             int     `json:"swim_pool_length"`
	TeValue                    int     `json:"te_value"`
	SwimStyle                  int     `json:"swim_style"`
	Unit                       int     `json:"unit"`
	IntervalType               int     `json:"interval_type"`
	DownhillNum                int     `json:"downhill_num"`
	DownhillMaxAltitudeDescent int     `json:"downhill_max_altitude_descend"`
	Strokes                    int     `json:"strokes"`
	ForeHand                   int     `json "fore_hand"` // Tennis
	BackHand                   int     `json:"back_hand"`
	Serve                      int     `json:"serve"`
	SecHaldStartTime           int     `json:"sec_half_start_time"`
	RopeSkippingTotalCounts    int     `json:"rope_skipping_total_counts"`
	RopeSkippingMaxFrequency   int     `json:"rope_skipping_max_frequency"`
	RopeSkippingAvgFrequency   int     `json:"rope_skiping_avg_frequency"`
	RopeSkippingRestTime       int     `json:"rope_skipping_rest_time"`
	WeatherInfo                string  `json:"weather_info"`
	PersonalBest               string  `json:"pb"`
	Marathon                   string  `json:"marathon"`
	Model                      string  `json:"model"`
}

func (a *Activity) ParseContent() {
	json.Unmarshal([]byte(a.ContentJSON), &(a.Content))
}

func (a *Activity) GenerateMap() {
	ctx := maps.NewContext()
	ctx.SetSize(1000, 1000)
	paths, err := maps.ParsePathString(a.PointsData)
	if err != nil {
		panic("Couldn't parse path for activity: " + err.Error())
	}
	path := paths[0]
	path.Weight = 1.0

	ctx.AddPath(path)
	img, err := ctx.Render()
	if err != nil {
		panic(err)
	}
	a.Image = img
}

func (c *Content) StartTimeStr() string {
	return time.Unix(0, int64(c.StartTime)*int64(time.Millisecond)).Format("January 2, 2006")
}
