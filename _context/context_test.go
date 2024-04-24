package _context

import (
	"fmt"
	"git.ziji.fun/junyang/go-common/_assert"
	"git.ziji.fun/junyang/go-common/_cmd"
	"net/http"
	"testing"
	"time"
)

type server struct {
	ctx     *Context
	handler *http.Server
}

func (this *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	this.ctx = New(w, r, false)
}
func (this *server) Run() {
	this.handler = &http.Server{
		Addr:    ":50001",
		Handler: http.HandlerFunc(this.ServeHTTP),
	}
	if err := this.handler.ListenAndServe(); nil != err {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestContext_GetGet(t *testing.T) {
	// no need tot est
	t.SkipNow()
}
func TestContext_GetGetAll(t *testing.T) {
	s := &server{}
	go func(s *server) {
		s.Run()
	}(s)
	time.Sleep(time.Second * 5)
	_cmd.ExecuteAsString("curl", "-X", "GET", "http://127.0.0.1:50001/api/test?get1=get1&get2=get2")
	{
		var expect string = "get1"
		get := s.ctx.GET["get1"]
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "get2"
		get := s.ctx.GET["get2"]
		_assert.Equal(t, expect, get)
	}
	getAll := s.ctx.GetGetAll()
	s.ctx.Get()
	{
		var expect string = "get1"
		get := getAll["get1"]
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "get2"
		get := getAll["get2"]
		_assert.Equal(t, expect, get)
	}
	s.handler.Close()
}
func TestContext_GetPost(t *testing.T) {
	//// no need to test
	//t.SkipNow()

	s := &server{}
	s.Run()

	//go func(s *server) {
	//	s.Run()
	//}(s)
	//time.Sleep(time.Second * 5)
	//_cmd.ExecuteAsString("curl", "-X", "POST", "http://127.0.0.1:50001/api/test?get1=get1&get2=get2", "-H", "Content-type: application/x-www-form-urlencoded", "-d", "post1=post1&post2=post2")
	//fmt.Println(s.ctx.RAW)
	//s.handler.Close()

}
func TestContext_GetPostAll(t *testing.T) {
	// application/x-www-form-urlencoded
	{
		s := &server{}
		go func(s *server) {
			s.Run()
		}(s)
		time.Sleep(time.Second * 5)
		_cmd.ExecuteAsString("curl", "-X", "POST", "http://127.0.0.1:50001/api/test?get1=get1&get2=get2", "-H", "Content-type: application/x-www-form-urlencoded", "-d", "post1=post1&post2=post2")
		{
			var expect string = "post1"
			get := s.ctx.POST["post1"]
			_assert.Equal(t, expect, get)
		}
		{
			var expect string = "post2"
			get := s.ctx.POST["post2"]
			_assert.Equal(t, expect, get)
		}
		postAll := s.ctx.GetPostAll()
		{
			var expect string = "post1"
			get := postAll["post1"]
			_assert.Equal(t, expect, get)
		}
		{
			var expect string = "post2"
			get := postAll["post2"]
			_assert.Equal(t, expect, get)
		}
		s.handler.Close()
	}
	// application/json
	{
		s := &server{}
		go func(s *server) {
			s.Run()
		}(s)
		time.Sleep(time.Second * 5)
		_cmd.ExecuteAsString("curl", "-X", "POST", "http://127.0.0.1:50001/api/test?get1=get1&get2=get2", "-H", "Content-type: application/json", "-d", `{"json1":"json1","json2":"json2"}`)
		{
			var expect string = "json1"
			get := s.ctx.POST["json1"]
			_assert.Equal(t, expect, get)
		}
		{
			var expect string = "json2"
			get := s.ctx.POST["json2"]
			_assert.Equal(t, expect, get)
		}
		postAll := s.ctx.GetPostAll()
		{
			var expect string = "json1"
			get := postAll["json1"]
			_assert.Equal(t, expect, get)
		}
		{
			var expect string = "json2"
			get := postAll["json2"]
			_assert.Equal(t, expect, get)
		}
		s.handler.Close()
	}
}
func TestContext_GetRequest(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestContext_GetRequestAll(t *testing.T) {
	s := &server{}
	go func(s *server) {
		s.Run()
	}(s)
	time.Sleep(time.Second * 5)
	_cmd.ExecuteAsString("curl", "-X", "POST", "http://127.0.0.1:50001/api/test?get1=get1", "-H", "Content-type: application/x-www-form-urlencoded", "-d", "post1=post1")
	{
		var expect string = "get1"
		get := s.ctx.REQUEST["get1"]
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "post1"
		get := s.ctx.REQUEST["post1"]
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "get1"
		get := s.ctx.GetRequest("get1").String().Value()
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "post1"
		get := s.ctx.GetRequest("post1").String().Value()
		_assert.Equal(t, expect, get)
	}
	requestAll := s.ctx.GetRequestAll()
	{
		var expect string = "get1"
		get := requestAll["get1"]
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "post1"
		get := requestAll["post1"]
		_assert.Equal(t, expect, get)
	}
	s.handler.Close()
}
