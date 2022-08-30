package main

import (
	"net/http"
)

type Controller struct {
	Data interface{}
	Req  *http.Request
	Res  http.ResponseWriter
}

func (p *Controller) Out() {
	if p.Data == nil {
		panic("继承者没有输出任何数据")
	}
	p.Res.Write([]byte(p.Data.(string)))
}

type Login struct {
	Controller
}

// 对应 http POST 方式
func (p *Controller) Post() {
	// WEB开发常见的控制器方法, 用于处理客户 http Post 方式的请求
	// 继承者应该覆盖这个方法, 否则认为不允许这样访问, 那就返回 403 拒绝访问
	p.Res.WriteHeader(403)
}

func (p *Login) Post() {
	if p.Req.Form.Get("login_name") == "" {
		p.Data = "无效的登录名"
		return
	}
	p.Data = "登录成功"
}

// 官方 net/http 包要求的接口方法
func (p *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Req = r
	p.Res = w
	if r.Method == "POST" {
		p.Post()
		p.Out()
	}
}

type HttpPostController interface {
	ServerHTTP(w http.ResponseWriter, r *http.Request)
	Post()
	Out()
}

func main() {
	http.ListenAndServe(":8000", nil)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		p := &Login{}
		p.ServeHTTP(w, r)
	})
}
