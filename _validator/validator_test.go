package _validator

import (
	"fmt"
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestNewInt64(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestInt64_Default(t *testing.T) {
	{
		var expect int64 = 1000
		get := NewInt64("int64", nil).Default(1000).Value()
		_assert.Equal(t, expect, get)
	}
}
func TestInt64_CodeMessage(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestInt64_EnsureMin(t *testing.T) {
	(func() {
		NewInt64("int64", 9).EnsureMin(9)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewInt64("int64", 9).EnsureMin(10)
	})()
}
func TestInt64_EnsureMax(t *testing.T) {
	(func() {
		NewInt64("int64", 9).EnsureMax(9)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewInt64("int64", 9).EnsureMax(8)
	})()
}
func TestInt64_EnsureBetween(t *testing.T) {
	(func() {
		NewInt64("int64", 9).EnsureBetween(8, 10)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewInt64("int64", 9).EnsureBetween(1, 8)
	})()
}
func TestInt64_EnsureLength(t *testing.T) {
	(func() {
		NewInt64("int64", 9).EnsureLength(1)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewInt64("int64", 9).EnsureLength(2)
	})()
}
func TestInt64_EnsureLengthMin(t *testing.T) {
	(func() {
		NewInt64("int64", 9).EnsureLengthMin(1)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewInt64("int64", 9).EnsureLengthMin(2)
	})()
}
func TestInt64_EnsureLengthMax(t *testing.T) {
	(func() {
		NewInt64("int64", 9).EnsureLengthMax(1)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewInt64("int64", 99).EnsureLengthMax(1)
	})()
}
func TestInt64_EnsureLengthBetween(t *testing.T) {
	(func() {
		NewInt64("int64", 9).EnsureLengthBetween(1, 2)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewInt64("int64", 999).EnsureLengthBetween(1, 2)
	})()
}
func TestInt64_EnsureIn(t *testing.T) {
	(func() {
		NewInt64("int64", 2).EnsureIn(1, 2, 3)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewInt64("int64", 4).EnsureIn(1, 2, 3)
	})()
}
func TestInt64_String(t *testing.T) {
	var expect string = "1"
	get := NewInt64("int64", 1).String().Value()
	_assert.Equal(t, expect, get)
}
func TestInt64_Bool(t *testing.T) {
	{
		var expect bool = true
		get := NewInt64("int64", 1).Bool().Value()
		_assert.Equal(t, expect, get)
	}
	{
		var expect bool = false
		get := NewInt64("int64", 0).Bool().Value()
		_assert.Equal(t, expect, get)
	}
}
func TestInt64_Float64(t *testing.T) {
	var expect float64 = 1
	get := NewInt64("int", 1).Float64().Value()
	_assert.EqualByFloat(t, expect, get)
}
func TestInt64_Value(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestNewString(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestString_Default(t *testing.T) {
	{
		var expect string = "hello world!"
		get := NewString("string", nil).Default("hello world!").Value()
		_assert.Equal(t, expect, get)
	}
}
func TestString_CodeMessage(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestString_EnsureEmpty(t *testing.T) {
	(func() {
		NewString("string", nil).EnsureEmpty()
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewString("string", "hello world!").EnsureEmpty()
	})()
}
func TestString_EnsureNotEmpty(t *testing.T) {
	(func() {
		NewString("string", "hello world!").EnsureNotEmpty()
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewString("string", nil).EnsureNotEmpty()
	})()
}
func TestString_EnsureLength(t *testing.T) {
	(func() {
		NewString("string", nil).EnsureLength(0)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewString("string", "hello world!").EnsureLength(0)
	})()
}
func TestString_EnsureLengthMin(t *testing.T) {
	(func() {
		NewString("string", "123").EnsureLengthMin(3)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewString("string", "12").EnsureLengthMin(3)
	})()
}
func TestString_EnsureLengthMax(t *testing.T) {
	(func() {
		NewString("string", "123").EnsureLengthMax(3)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewString("string", "1234").EnsureLengthMax(3)
	})()
}
func TestString_EnsureLengthBetween(t *testing.T) {
	(func() {
		NewString("string", "123").EnsureLengthBetween(1, 3)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewString("string", "1234").EnsureLengthBetween(1, 3)
	})()
}
func TestString_EnsureIn(t *testing.T) {
	(func() {
		NewString("string", "123").EnsureIn("1", "12", "123")
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewString("string", "1234").EnsureIn("1", "12", "123")
	})()
}
func TestString_EnsureRegexp(t *testing.T) {
	(func() {
		NewString("string", "2024-02-04 02:04:00").EnsureRegexp(`^\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}$`)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewString("string", "2024-02-04 02:04:Aa").EnsureRegexp(`^\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}$`)
	})()
}
func TestString_Bool(t *testing.T) {
	{
		var expect bool = true
		get := NewString("string", "0").Bool().Value()
		_assert.Equal(t, expect, get)
	}
	{
		var expect bool = false
		get := NewString("string", "").Bool().Value()
		_assert.Equal(t, expect, get)
	}
}
func TestString_Float64(t *testing.T) {
	var expect float64 = 3.141592
	get := NewString("string", "3.141592").Float64().Value()
	_assert.EqualByFloat(t, expect, get)
}
func TestString_Int64(t *testing.T) {
	var expect int64 = 99
	get := NewString("string", "99").Int64().Value()
	_assert.Equal(t, expect, get)
}
func TestNewBool(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestBool_Default(t *testing.T) {
	{
		var expect bool = true
		get := NewBool("bool", nil).Default(true).Value()
		_assert.Equal(t, expect, get)
	}
	{
		var expect bool = false
		get := NewBool("bool", nil).Default(false).Value()
		_assert.Equal(t, expect, get)
	}
}
func TestBool_CodeMessage(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestBool_EnsureFalse(t *testing.T) {
	(func() {
		NewBool("bool", false).EnsureFalse()
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewBool("bool", true).EnsureFalse()
	})()
}
func TestBool_EnsureTrue(t *testing.T) {
	(func() {
		NewBool("bool", true).EnsureTrue()
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewBool("bool", false).EnsureTrue()
	})()
}
func TestBool_EnsureIn(t *testing.T) {
	(func() {
		NewBool("bool", false).EnsureIn(false)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewBool("bool", false).EnsureIn(true)
	})()
}
func TestBool_Float64(t *testing.T) {
	{
		var expect float64 = 1
		get := NewBool("bool", true).Float64().Value()
		_assert.EqualByFloat(t, expect, get)
	}
	{
		var expect float64 = 0
		get := NewBool("bool", false).Float64().Value()
		_assert.EqualByFloat(t, expect, get)
	}
}
func TestBool_Int64(t *testing.T) {
	{
		var expect int64 = 1
		get := NewBool("bool", true).Int64().Value()
		_assert.Equal(t, expect, get)
	}
	{
		var expect int64 = 0
		get := NewBool("bool", false).Int64().Value()
		_assert.EqualByFloat(t, expect, get)
	}
}
func TestBool_String(t *testing.T) {
	{
		var expect string = "1"
		get := NewBool("bool", true).String().Value()
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = ""
		get := NewBool("bool", false).String().Value()
		_assert.EqualByFloat(t, expect, get)
	}
}
func TestBool_Value(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestNewFloat64(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestFloat64_Default(t *testing.T) {
	{
		var expect float64 = 3.141592
		get := NewFloat64("float64", nil).Default(3.141592).Value()
		_assert.EqualByFloat(t, expect, get)
	}
}
func TestFloat64_CodeMessage(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestFloat64_EnsureMin(t *testing.T) {
	(func() {
		NewFloat64("float64", 3.141592).EnsureMin(3.14159)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewFloat64("float64", 3.141592).EnsureMin(3.1415926)
	})()
}
func TestFloat64_EnsureMax(t *testing.T) {
	(func() {
		NewFloat64("float64", 3.141592).EnsureMax(3.1415926)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewFloat64("float64", 3.141592).EnsureMax(3.14159)
	})()
}
func TestFloat64_EnsureBetween(t *testing.T) {
	(func() {
		NewFloat64("float64", 3.141592).EnsureBetween(3.14159, 3.1415926)
	})()
	(func() {
		defer func() {
			err := recover()
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}()
		NewFloat64("float64", 3.141592).EnsureBetween(3.1415926, 3.14159261)
	})()
}
func TestFloat64_Int64(t *testing.T) {
	var expect int64 = 3
	get := NewFloat64("float64", 3.141592).Int64().Value()
	_assert.Equal(t, expect, get)
}
func TestFloat64_Bool(t *testing.T) {
	{
		var expect bool = true
		get := NewFloat64("float64", 3.141592).Bool().Value()
		_assert.Equal(t, expect, get)
	}
	{
		var expect bool = false
		get := NewFloat64("float64", 0).Bool().Value()
		_assert.Equal(t, expect, get)
	}
}
func TestFloat64_Value(t *testing.T) {
	// no need to test
	t.SkipNow()
}
