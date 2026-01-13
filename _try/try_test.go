package _try

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestTryCatch(t *testing.T) {
	{
		var caughtError any
		logic := func() {
			panic("Test panic")
		}
		New(logic).
			Catch(func(err any) {
				caughtError = err
			}).
			Do()
		_assert.Equal(t, "Test panic", caughtError)
	}
}
func TestNoPanic(t *testing.T) {
	{
		var caughtError any
		logic := func() {}
		New(logic).
			Catch(func(err any) {
				caughtError = err
			}).
			Do()
		_assert.Nil(t, caughtError)
	}
}
func TestMultipleCatchHandlers(t *testing.T) {
	{
		var caughtError any
		logic := func() {
			panic("Test panic")
		}
		New(logic).
			Catch(func(err any) {
				caughtError = "Handler 1: " + err.(string)
			}).
			Catch(func(err any) {
				caughtError = "Handler 2: " + err.(string)
			}).
			Do()
		_assert.Equal(t, "Handler 2: Test panic", caughtError)
	}
}
func TestCatchHandlesDifferentErrors(t *testing.T) {
	{
		var caughtError any
		logic := func() {
			panic("Custom Error")
		}
		New(logic).
			Catch(func(err any) {
				if err.(string) == "Custom Error" {
					caughtError = "Caught Custom Error"
				}
			}).
			Do()
		_assert.Equal(t, "Caught Custom Error", caughtError)
	}
}
func TestNilLogicFunction(t *testing.T) {
	{
		var caughtError any
		logic := func() {}
		New(logic).
			Catch(func(err any) {
				caughtError = err
			}).
			Do()
		_assert.Nil(t, caughtError)
	}
}
func TestEmptyCatch(t *testing.T) {
	{
		var caughtError any
		logic := func() {
			panic("Test panic")
		}
		New(logic).
			Catch(func(err any) {
			}).
			Do()
		_assert.Nil(t, caughtError)
	}
}
func TestRecoverAfterPanic(t *testing.T) {
	{
		var caughtError any
		logic := func() {
			defer func() {
				if err := recover(); err != nil {
					caughtError = "Recovered: " + err.(string)
				}
			}()
			panic("Panic occurred")
		}
		New(logic).
			Catch(func(err any) {}).
			Do()
		_assert.Equal(t, "Recovered: Panic occurred", caughtError)
	}
}
