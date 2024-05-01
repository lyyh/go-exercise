package main

import (
	"fmt"
	"runtime"
	"strings"
)

type ExecTimeResp struct {
	ImageDownloadTime float64 `json:"imageDownloadTime,omitempty"` // 图片下载耗时
	ImageDecodeTime   float64 `json:"imageDecodeTime,omitempty"`   // 图片解码耗时
	ModelInferTime    float64 `json:"modelInferTime,omitempty"`    // 算法推理时间
	TotalTime         float64 `json:"totalTime,omitempty"`         // 算法内部整体执行时间
	Location          *Location
}

type Location struct {
	X          int64             `json:"x"`
	Security   map[string]string `json:"security"`
	Intimidate []Intimidate      `json:"intimidate"`
}

type Intimidate struct {
	Type  string  `json:"type"`
	Score float64 `json:"score"`
	State string  `json:"state"`
	XMin  int     `json:"xmin"`
	YMin  int     `json:"ymin"`
	XMax  int     `json:"xmax"`
	YMax  int     `json:"ymax"`
	Angle int     `json:"angle"`
}

// func (e ExecTimeResp) String() string {
// 	fmt.Println("adfasdf")
// 	return fmt.Sprintf("{ImageDownloadTime: %v}", e.ImageDecodeTime)
// }

func (l *Location) String() string {
	return fmt.Sprintf("{X:%+v,Security:%+v,Intimidate:%+v}", l.X, l.Security, l.Intimidate)
}

// stack represents a stack of program counters.
type stack []uintptr

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

type V struct {
}

func (v *V) Format(s fmt.State, verb rune) {
	fmt.Println("format")
}

func main() {
	// var a = callers()
	// fmt.Printf("%+v", *a)

	stack := make([]byte, 1024*1024)
	n := runtime.Stack(stack, true)
	fmt.Println(string(stack[:n]))

	v := V{}
	a := fmt.Sprintf("%+v", v)
	fmt.Println(a)

	var s = "http://localhost:8080/3"
	fmt.Println(strings.Trim(s, "/"))
	xs := strings.Split(strings.Trim(s, "/"), "/")
	fmt.Println(xs[0])

	// resp := ExecTimeResp{
	// 	Location: &Location{
	// 		X: 1,
	// 		Security: map[string]string{
	// 			"asdf": "123123",
	// 		},
	// 		Intimidate: []Intimidate{Intimidate{}},
	// 	},
	// }
	// fmt.Printf("%+v", resp)
	//
	// // fmt.Sprintf("%%.%df", 3)
	// var a float64
	// var b float64
	// a = 100
	// b = 5.1
	// res := b / a
	// value, _ := strconv.ParseFloat(fmt.Sprintf(fmt.Sprintf("%%.%df", 3), res), 64)
	// fmt.Println(value)
	// resp := ExecTimeResp{}
	// fmt.Printf("resp:%+v", resp)
	// marshal, err := json.Marshal(resp)
	// if err != nil {
	// 	fmt.Printf("%s", err.Error())
	// }
	// fmt.Printf("aa %s", string(marshal))
	// resp := ExecTimeResp{}
	// err := json.Unmarshal([]byte(`{"imageDownloadTime":0.04425782896578312,"imageDecodeTime":0.07776033389382064,"modelInferTime":0.0523483669385314,"totalTime":0.17530754185281694}`), &resp)
	// if err != nil {
	// 	fmt.Printf("%s", err.Error())
	// }
	// fmt.Printf("%+v", resp)

	// namesNumMap := make(map[string]int)
	// namesNumMap["a"] += 1
	// fmt.Printf("%+v", namesNumMap)
	// a := make(map[string]string)
	// fmt.Printf("%s", a["a"])
	// s := make(map[string]map[string]map[string]bool)
	// s["a"] = make(map[string]map[string]bool)
	// // s["a"]["b"] = make(map[string]bool
	//
	// v := "123"
	// fmt.Println(fmt.Sprintf("s:", v))
	// a := s["a"]
	// a["b"] = make(map[string]bool)

}
