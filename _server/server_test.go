package _server

import (
	"fmt"
	"git.ziji.fun/junyang/go-common/_assert"
	"git.ziji.fun/junyang/go-common/_cmd"
	"git.ziji.fun/junyang/go-common/_context"
	"git.ziji.fun/junyang/go-common/_directory"
	"git.ziji.fun/junyang/go-common/_hash"
	"git.ziji.fun/junyang/go-common/_router"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestWeb(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestWebEngine_Addr(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestWebEngine_Root(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestWebEngine_Run(t *testing.T) {
	{
		var response string = ""
		wg := sync.WaitGroup{}
		go func() {
			path := _directory.Current() + "/test_web_run"
			web := Web()
			web.Addr(":50001").Root(path).Run()
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			time.Sleep(time.Second * 9)
			response = _cmd.ExecuteAsString("curl", "http://127.0.0.1:50001/")
			wg.Done()
		}()
		wg.Wait()
		var expect string = "58041c180a3c476c01c77cb32d3c12f3"
		get := _hash.Md5(response)
		_assert.Equal(t, expect, get)
	}
}
func TestApi(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestApiEngine_Addr(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestApiEngine_Origin(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestApiEngine_Router(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestApiEngine_Run(t *testing.T) {
	{
		var response string = ``
		s := Api().Addr(`:50001`).Router(&_router.Router{
			Rule:       `^/(\w+)/test`,
			MethodList: []string{"ANY"},
			Call: func(ctx *_context.Context) {
				ctx.JSON(nil)
			},
			IsRegexp:      true,
			ParameterList: []string{"a"},
		})
		go func() {
			defer func() {
				err := recover()
				if err != http.ErrServerClosed {
					panic(err)
				}
			}()
			s.Run()
		}()
		time.Sleep(time.Second * 9)
		response = _cmd.ExecuteAsString("curl", "http://127.0.0.1:50001/api/test")
		s.handler.Close()
		fmt.Println(response)

	}
}
