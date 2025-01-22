package _is

import (
	"github.com/junyang7/go-common/src/_assert"
	"testing"
)

func TestEmpty(t *testing.T) {
	{
		{
			var give bool
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give bool = false
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give bool = true
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give int8
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int8 = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int8 = 1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int8 = -1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give int16
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int16 = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int16 = 1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int16 = -1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give int32
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int32 = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int32 = 1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int32 = -1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give int64
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int64 = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int64 = 1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int64 = -1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give int
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int = 1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give int = -1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give uint8
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give uint8 = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give uint8 = 1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give uint16
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give uint16 = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give uint16 = 1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give uint32
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give uint32 = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give uint32 = 1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give uint64
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give uint64 = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give uint64 = 1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give uint
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give uint = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give uint = 1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give float32
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give float32 = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give float32 = 0.1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give float32 = -0.1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give float64
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give float64 = 0
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give float64 = 0.1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give float64 = -0.1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give string
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give string = ""
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give string = " "
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give string = "Hello world!"
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give any
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give interface{}
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give []int
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give []int = []int{}
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give []int = []int{1}
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give [3]int
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give [3]int = [3]int{}
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give [3]int = [3]int{1}
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give map[string]string
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give map[string]string = map[string]string{}
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give map[string]string = map[string]string{"a": "A"}
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
	{
		{
			var give chan int
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give chan int = make(chan int)
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give chan int = make(chan int, 2)
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give chan int = make(chan int, 2)
			give <- 1
			var expect bool = false
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
		{
			var give chan int = make(chan int, 2)
			give <- 1
			<-give
			var expect bool = true
			get := Empty(give)
			_assert.Equal(t, expect, get)
		}
	}
}
