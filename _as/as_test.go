package _as

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestString(t *testing.T) {
	// []byte
	{
		var expect string = "hello world!"
		var give []byte = []byte("hello world!")
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// string
	{
		var expect string = "3.141592"
		var give string = "3.141592"
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// int8
	{
		var expect string = "8"
		var give int8 = 8
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// int16
	{
		var expect string = "16"
		var give int16 = 16
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// int32
	{
		var expect string = "32"
		var give int32 = 32
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// int64
	{
		var expect string = "64"
		var give int64 = 64
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// int
	{
		var expect string = "0"
		var give int = 0
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// uint8
	{
		var expect string = "8"
		var give uint8 = 8
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// uint16
	{
		var expect string = "16"
		var give uint16 = 16
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// uint32
	{
		var expect string = "32"
		var give uint32 = 32
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// uint64
	{
		var expect string = "64"
		var give uint64 = 64
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// uint
	{
		var expect string = "0"
		var give uint = 0
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// float32
	{
		var expect string = "3.141592"
		var give float32 = 3.141592
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// float64
	{
		var expect string = "3.141592"
		var give float64 = 3.141592
		get := String(give)
		_assert.Equal(t, expect, get)
	}
	// bool
	{
		// false
		{
			var expect string = ""
			var give bool = false
			get := String(give)
			_assert.Equal(t, expect, get)
		}
		// true
		{
			var expect string = "1"
			var give bool = true
			get := String(give)
			_assert.Equal(t, expect, get)
		}
	}
	// other
	{
		var expect string = ""
		var give chan string = nil
		get := String(give)
		_assert.Equal(t, expect, get)
	}
}
func TestBool(t *testing.T) {
	// []byte
	{
		// true
		{
			var expect bool = true
			var give []byte = []byte("hello world!")
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give []byte = []byte("")
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// string
	{
		// true
		{
			var expect bool = true
			var give string = "3.141592"
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give string = ""
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int8
	{
		// true
		{
			var expect bool = true
			var give int8 = 8
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give int8 = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int16
	{
		// true
		{
			var expect bool = true
			var give int16 = 16
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give int16 = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int32
	{
		// true
		{
			var expect bool = true
			var give int32 = 32
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give int32 = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int64
	{
		// true
		{
			var expect bool = true
			var give int64 = 64
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give int64 = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int
	{
		// true
		{
			var expect bool = true
			var give int = 1
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give int = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint8
	{
		// true
		{
			var expect bool = true
			var give uint8 = 8
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give uint8 = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint16
	{
		// true
		{
			var expect bool = true
			var give uint16 = 16
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give uint16 = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint32
	{
		// true
		{
			var expect bool = true
			var give uint32 = 32
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give uint32 = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint64
	{
		// true
		{
			var expect bool = true
			var give uint64 = 64
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give uint64 = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint
	{
		// true
		{
			var expect bool = true
			var give uint = 1
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give uint = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// float32
	{
		// true
		{
			var expect bool = true
			var give float32 = 3.141592
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give float32 = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// float64
	{
		// true
		{
			var expect bool = true
			var give float64 = 3.141592
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give float64 = 0
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// bool
	{
		// true
		{
			var expect bool = true
			var give bool = true
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
		// false
		{
			var expect bool = false
			var give bool = false
			get := Bool(give)
			_assert.Equal(t, expect, get)
		}
	}
	// other
	{
		var expect bool = false
		var give chan string = nil
		get := Bool(give)
		_assert.Equal(t, expect, get)
	}
}
func TestByteList(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestFloat64(t *testing.T) {
	// []byte
	{
		{
			var expect float64 = 0
			var give []byte = []byte("hello world!")
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 0
			var give []byte = []byte("3.141592")
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// string
	{
		{
			var expect float64 = 0
			var give string = "hello world!"
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 0
			var give string = ""
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 0
			var give string = "0"
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give string = "1"
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 3.141592
			var give string = "3.141592"
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// int8
	{
		{
			var expect float64 = 0
			var give int8 = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = -1
			var give int8 = -1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give int8 = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// int16
	{
		{
			var expect float64 = 0
			var give int16 = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = -1
			var give int16 = -1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give int16 = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// int32
	{
		{
			var expect float64 = 0
			var give int32 = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = -1
			var give int32 = -1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give int32 = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// int64
	{
		{
			var expect float64 = 0
			var give int64 = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = -1
			var give int64 = -1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give int64 = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// int
	{
		{
			var expect float64 = 0
			var give int = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = -1
			var give int = -1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give int = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// uint8
	{
		{
			var expect float64 = 0
			var give uint8 = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give uint8 = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// uint16
	{
		{
			var expect float64 = 0
			var give uint16 = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give uint16 = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// uint32
	{
		{
			var expect float64 = 0
			var give uint32 = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give uint32 = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// uint64
	{
		{
			var expect float64 = 0
			var give uint64 = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give uint64 = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// uint
	{
		{
			var expect float64 = 0
			var give uint = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give uint = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// float32
	{
		{
			var expect float64 = 0
			var give float32 = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give float32 = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = -1
			var give float32 = -1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 3.141592
			var give float32 = 3.141592
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = -3.141592
			var give float32 = -3.141592
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// float64
	{
		{
			var expect float64 = 0
			var give float64 = 0
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give float64 = 1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = -1
			var give float64 = -1
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 3.141592
			var give float64 = 3.141592
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = -3.141592
			var give float64 = -3.141592
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// bool
	{
		{
			var expect float64 = 0
			var give bool = false
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
		{
			var expect float64 = 1
			var give bool = true
			get := Float64(give)
			_assert.EqualByFloat(t, expect, get)
		}
	}
	// other
	{
		var expect float64 = 0
		var give chan string = nil
		get := Float64(give)
		_assert.EqualByFloat(t, expect, get)
	}
}
func TestInt64(t *testing.T) {
	// []byte
	{
		{
			var expect int64 = 0
			var give []byte = []byte("hello world!")
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give []byte = []byte("1")
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -1
			var give []byte = []byte("-1")
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 0
			var give []byte = []byte("3.1415926")
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// string
	{
		{
			var expect int64 = 0
			var give string = "hello world!"
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give string = "1"
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -1
			var give string = "-1"
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 0
			var give string = "3.1415926"
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int8
	{
		{
			var expect int64 = 0
			var give int8 = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give int8 = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -1
			var give int8 = -1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int16
	{
		{
			var expect int64 = 0
			var give int16 = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give int16 = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -1
			var give int16 = -1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int32
	{
		{
			var expect int64 = 0
			var give int32 = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give int32 = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -1
			var give int32 = -1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int64
	{
		{
			var expect int64 = 0
			var give int64 = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give int64 = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -1
			var give int64 = -1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int
	{
		{
			var expect int64 = 0
			var give int = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give int = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -1
			var give int = -1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint8
	{
		{
			var expect int64 = 0
			var give uint8 = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give uint8 = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint16
	{
		{
			var expect int64 = 0
			var give uint16 = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give uint16 = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint32
	{
		{
			var expect int64 = 0
			var give uint32 = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give uint32 = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint64
	{
		{
			var expect int64 = 0
			var give uint64 = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give uint64 = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint
	{
		{
			var expect int64 = 0
			var give uint = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give uint = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// float32
	{
		{
			var expect int64 = 0
			var give float32 = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give float32 = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -1
			var give float32 = -1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 3
			var give float32 = 3.141592
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -3
			var give float32 = -3.141592
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 3
			var give float32 = 3.92153
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -3
			var give float32 = -3.92153
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// float64
	{
		{
			var expect int64 = 0
			var give float64 = 0
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give float64 = 1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -1
			var give float64 = -1
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 3
			var give float64 = 3.141592
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -3
			var give float64 = -3.141592
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 3
			var give float64 = 3.92153
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = -3
			var give float64 = -3.92153
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// bool
	{
		{
			var expect int64 = 0
			var give bool = false
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int64 = 1
			var give bool = true
			get := Int64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// other
	{
		var expect int64 = 0
		var give chan string = nil
		get := Int64(give)
		_assert.Equal(t, expect, get)
	}
}
func TestInt(t *testing.T) {
	// []byte
	{
		{
			var expect int = 0
			var give []byte = []byte("hello world!")
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give []byte = []byte("1")
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -1
			var give []byte = []byte("-1")
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 0
			var give []byte = []byte("3.1415926")
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// string
	{
		{
			var expect int = 0
			var give string = "hello world!"
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give string = "1"
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -1
			var give string = "-1"
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 0
			var give string = "3.1415926"
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int8
	{
		{
			var expect int = 0
			var give int8 = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give int8 = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -1
			var give int8 = -1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int16
	{
		{
			var expect int = 0
			var give int16 = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give int16 = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -1
			var give int16 = -1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int32
	{
		{
			var expect int = 0
			var give int32 = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give int32 = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -1
			var give int32 = -1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int64
	{
		{
			var expect int = 0
			var give int64 = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give int64 = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -1
			var give int64 = -1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int
	{
		{
			var expect int = 0
			var give int = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give int = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -1
			var give int = -1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint8
	{
		{
			var expect int = 0
			var give uint8 = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give uint8 = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint16
	{
		{
			var expect int = 0
			var give uint16 = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give uint16 = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint32
	{
		{
			var expect int = 0
			var give uint32 = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give uint32 = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint64
	{
		{
			var expect int = 0
			var give uint64 = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give uint64 = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint
	{
		{
			var expect int = 0
			var give uint = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give uint = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// float32
	{
		{
			var expect int = 0
			var give float32 = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give float32 = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -1
			var give float32 = -1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 3
			var give float32 = 3.141592
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -3
			var give float32 = -3.141592
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 3
			var give float32 = 3.92153
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -3
			var give float32 = -3.92153
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// float64
	{
		{
			var expect int = 0
			var give float64 = 0
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give float64 = 1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -1
			var give float64 = -1
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 3
			var give float64 = 3.141592
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -3
			var give float64 = -3.141592
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 3
			var give float64 = 3.92153
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = -3
			var give float64 = -3.92153
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// bool
	{
		{
			var expect int = 0
			var give bool = false
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect int = 1
			var give bool = true
			get := Int(give)
			_assert.Equal(t, expect, get)
		}
	}
	// other
	{
		var expect int = 0
		var give chan string = nil
		get := Int(give)
		_assert.Equal(t, expect, get)
	}
}
func TestUint64(t *testing.T) {
	// []byte
	{
		{
			var expect uint64 = 0
			var give []byte = []byte("hello world!")
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give []byte = []byte("1")
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 0
			var give []byte = []byte("3.1415926")
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// string
	{
		{
			var expect uint64 = 0
			var give string = "hello world!"
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give string = "1"
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 0
			var give string = "3.1415926"
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int8
	{
		{
			var expect uint64 = 0
			var give int8 = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give int8 = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int16
	{
		{
			var expect uint64 = 0
			var give int16 = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give int16 = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int32
	{
		{
			var expect uint64 = 0
			var give int32 = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give int32 = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int64
	{
		{
			var expect uint64 = 0
			var give int64 = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give int64 = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int
	{
		{
			var expect uint64 = 0
			var give int = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give int = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint8
	{
		{
			var expect uint64 = 0
			var give uint8 = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give uint8 = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint16
	{
		{
			var expect uint64 = 0
			var give uint16 = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give uint16 = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint32
	{
		{
			var expect uint64 = 0
			var give uint32 = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give uint32 = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint64
	{
		{
			var expect uint64 = 0
			var give uint64 = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give uint64 = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint
	{
		{
			var expect uint64 = 0
			var give uint = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give uint = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// float32
	{
		{
			var expect uint64 = 0
			var give float32 = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give float32 = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 3
			var give float32 = 3.141592
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 3
			var give float32 = 3.92153
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// float64
	{
		{
			var expect uint64 = 0
			var give float64 = 0
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give float64 = 1
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 3
			var give float64 = 3.141592
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 3
			var give float64 = 3.92153
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// bool
	{
		{
			var expect uint64 = 0
			var give bool = false
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint64 = 1
			var give bool = true
			get := Uint64(give)
			_assert.Equal(t, expect, get)
		}
	}
	// other
	{
		var expect uint64 = 0
		var give chan string = nil
		get := Uint64(give)
		_assert.Equal(t, expect, get)
	}
}
func TestUint(t *testing.T) {
	// []byte
	{
		{
			var expect uint = 0
			var give []byte = []byte("hello world!")
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give []byte = []byte("1")
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 0
			var give []byte = []byte("3.1415926")
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// string
	{
		{
			var expect uint = 0
			var give string = "hello world!"
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give string = "1"
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 0
			var give string = "3.1415926"
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int8
	{
		{
			var expect uint = 0
			var give int8 = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give int8 = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int16
	{
		{
			var expect uint = 0
			var give int16 = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give int16 = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int32
	{
		{
			var expect uint = 0
			var give int32 = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give int32 = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int64
	{
		{
			var expect uint = 0
			var give int64 = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give int64 = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// int
	{
		{
			var expect uint = 0
			var give int = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give int = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint8
	{
		{
			var expect uint = 0
			var give uint8 = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give uint8 = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint16
	{
		{
			var expect uint = 0
			var give uint16 = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give uint16 = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint32
	{
		{
			var expect uint = 0
			var give uint32 = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give uint32 = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint64
	{
		{
			var expect uint = 0
			var give uint64 = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give uint64 = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// uint
	{
		{
			var expect uint = 0
			var give uint = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give uint = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// float32
	{
		{
			var expect uint = 0
			var give float32 = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give float32 = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 3
			var give float32 = 3.141592
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 3
			var give float32 = 3.92153
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// float64
	{
		{
			var expect uint = 0
			var give float64 = 0
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give float64 = 1
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 3
			var give float64 = 3.141592
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 3
			var give float64 = 3.92153
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// bool
	{
		{
			var expect uint = 0
			var give bool = false
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
		{
			var expect uint = 1
			var give bool = true
			get := Uint(give)
			_assert.Equal(t, expect, get)
		}
	}
	// other
	{
		var expect uint = 0
		var give chan string = nil
		get := Uint(give)
		_assert.Equal(t, expect, get)
	}
}
