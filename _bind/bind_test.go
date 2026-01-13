package _bind

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

type InnerStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
type InnerDetails struct {
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func TestDo_BasicTypes(t *testing.T) {
	{
		data := map[string]interface{}{
			"id":   1,
			"name": "Test",
			"flag": true,
		}
		v := struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Flag bool   `json:"flag"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Name, "Test")
		_assert.Equal(t, v.Flag, true)
	}
	{
		data := map[string]interface{}{
			"id": 1,
		}
		v := struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Name, "")
	}
	{
		data := map[string]interface{}{
			"id":   "string_instead_of_int",
			"flag": true,
		}
		v := struct {
			ID   int  `json:"id"`
			Flag bool `json:"flag"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 0)
		_assert.Equal(t, v.Flag, true)
	}
}
func TestDo_StructBinding(t *testing.T) {
	{
		data := map[string]interface{}{
			"id":   1,
			"info": map[string]interface{}{"name": "Test", "value": 42},
		}
		v := struct {
			ID   int         `json:"id"`
			Info InnerStruct `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Info.Name, "Test")
		_assert.Equal(t, v.Info.Value, 42)
	}
	{
		data := map[string]interface{}{
			"info": map[string]interface{}{"name": "Test"},
		}
		v := struct {
			ID   int         `json:"id"`
			Info InnerStruct `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 0)
		_assert.Equal(t, v.Info.Name, "Test")
		_assert.Equal(t, v.Info.Value, 0)
	}
}
func TestDo_RecursiveStructBinding(t *testing.T) {
	{
		data := map[string]interface{}{
			"id":   1,
			"info": map[string]interface{}{"name": "Test", "value": 42},
			"details": map[string]interface{}{
				"address": "123 Main St",
				"phone":   "123-456-7890",
			},
		}
		v := struct {
			ID      int          `json:"id"`
			Info    InnerStruct  `json:"info"`
			Details InnerDetails `json:"details"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Info.Name, "Test")
		_assert.Equal(t, v.Info.Value, 42)
		_assert.Equal(t, v.Details.Address, "123 Main St")
		_assert.Equal(t, v.Details.Phone, "123-456-7890")
	}
}
func TestDo_NilPointer(t *testing.T) {
	{
		var v *struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		data := map[string]interface{}{
			"id":   1,
			"name": "Test",
		}
		err := Do(v, data)
		_assert.NoError(t, err)
	}
}
func TestDo_EmptyData(t *testing.T) {
	{
		data := map[string]interface{}{}
		v := struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 0)
		_assert.Equal(t, v.Name, "")
	}
}
func TestDo_InvalidDataType(t *testing.T) {
	{
		data := map[string]interface{}{
			"id":   []int{1, 2},
			"flag": true,
		}
		v := struct {
			ID   int  `json:"id"`
			Flag bool `json:"flag"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 0)
		_assert.Equal(t, v.Flag, true)
	}
}
func TestDo_NestedMapBinding(t *testing.T) {
	{
		data := map[string]interface{}{
			"info": map[string]interface{}{"name": "Test", "value": 42},
		}
		v := struct {
			Info map[string]interface{} `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Info["name"], "Test")
		_assert.Equal(t, v.Info["value"], 42)
	}
	{
		data := map[string]interface{}{
			"info": map[string]interface{}{"name": "Test", "value": "42"},
		}
		v := struct {
			Info map[string]string `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Info["name"], "Test")
		_assert.Equal(t, v.Info["value"], "42")
	}
	{
		data := map[string]interface{}{
			"info": map[string]interface{}{"a": 1, "b": 2},
		}
		v := struct {
			Info map[string]int `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Info["a"], 1)
		_assert.Equal(t, v.Info["b"], 2)
	}
	{
		data := map[string]interface{}{
			"info": map[string]interface{}{"x": 1.5, "y": 2.7},
		}
		v := struct {
			Info map[string]float64 `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Info["x"], 1.5)
		_assert.Equal(t, v.Info["y"], 2.7)
	}
}
func TestDo_NonStructInput(t *testing.T) {
	{
		var v int
		data := map[string]interface{}{
			"id": 1,
		}
		err := Do(v, data)
		_assert.NoError(t, err)
	}
	{
		var v string
		data := map[string]interface{}{
			"name": "Test",
		}
		err := Do(v, data)
		_assert.NoError(t, err)
	}
	{
		var v []int
		data := map[string]interface{}{
			"items": []int{1, 2, 3},
		}
		err := Do(v, data)
		_assert.NoError(t, err)
	}
}
func TestDo_NilData(t *testing.T) {
	{
		v := struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{}
		err := Do(&v, nil)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 0)
		_assert.Equal(t, v.Name, "")
	}
}
func TestDo_TypeConversions_Comprehensive(t *testing.T) {
	{
		data := map[string]interface{}{
			"int_val":    int(100),
			"int8_val":   int8(8),
			"int16_val":  int16(16),
			"int32_val":  int32(32),
			"int64_val":  int64(64),
			"uint_val":   uint(200),
			"uint8_val":  uint8(8),
			"uint16_val": uint16(16),
			"uint32_val": uint32(32),
			"uint64_val": uint64(64),
		}
		v := struct {
			IntVal    int    `json:"int_val"`
			Int8Val   int    `json:"int8_val"`
			Int16Val  int    `json:"int16_val"`
			Int32Val  int    `json:"int32_val"`
			Int64Val  int64  `json:"int64_val"`
			UintVal   uint   `json:"uint_val"`
			Uint8Val  uint   `json:"uint8_val"`
			Uint16Val uint   `json:"uint16_val"`
			Uint32Val uint   `json:"uint32_val"`
			Uint64Val uint64 `json:"uint64_val"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.IntVal, 100)
		_assert.Equal(t, v.Int8Val, 8)
		_assert.Equal(t, v.Int16Val, 16)
		_assert.Equal(t, v.Int32Val, 32)
		_assert.Equal(t, v.Int64Val, int64(64))
		_assert.Equal(t, v.UintVal, uint(200))
		_assert.Equal(t, v.Uint8Val, uint(8))
		_assert.Equal(t, v.Uint16Val, uint(16))
		_assert.Equal(t, v.Uint32Val, uint(32))
		_assert.Equal(t, v.Uint64Val, uint64(64))
	}
	{
		data := map[string]interface{}{
			"float32_val":   float32(3.14),
			"float64_val":   float64(2.718),
			"int_to_float":  int(100),
			"uint_to_float": uint(200),
		}
		v := struct {
			Float32Val  float64 `json:"float32_val"`
			Float64Val  float64 `json:"float64_val"`
			IntToFloat  float64 `json:"int_to_float"`
			UintToFloat float64 `json:"uint_to_float"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.EqualByFloat(t, v.Float32Val, 3.140000104904175)
		_assert.Equal(t, v.Float64Val, 2.718)
		_assert.Equal(t, v.IntToFloat, 100.0)
		_assert.Equal(t, v.UintToFloat, 200.0)
	}
	{
		data := map[string]interface{}{
			"str_to_int":    "123",
			"str_to_int64":  "456",
			"str_to_uint":   "789",
			"str_to_uint64": "999",
			"str_to_float":  "3.14",
			"str_to_bool":   "true",
		}
		v := struct {
			StrToInt    int     `json:"str_to_int"`
			StrToInt64  int64   `json:"str_to_int64"`
			StrToUint   uint    `json:"str_to_uint"`
			StrToUint64 uint64  `json:"str_to_uint64"`
			StrToFloat  float64 `json:"str_to_float"`
			StrToBool   bool    `json:"str_to_bool"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.StrToInt, 123)
		_assert.Equal(t, v.StrToInt64, int64(456))
		_assert.Equal(t, v.StrToUint, uint(789))
		_assert.Equal(t, v.StrToUint64, uint64(999))
		_assert.Equal(t, v.StrToFloat, 3.14)
		_assert.Equal(t, v.StrToBool, true)
	}
	{
		data := map[string]interface{}{
			"int_to_str":   int(123),
			"int64_to_str": int64(456),
			"uint_to_str":  uint(789),
			"float_to_str": float64(3.14),
			"bool_to_str":  true,
		}
		v := struct {
			IntToStr   string `json:"int_to_str"`
			Int64ToStr string `json:"int64_to_str"`
			UintToStr  string `json:"uint_to_str"`
			FloatToStr string `json:"float_to_str"`
			BoolToStr  string `json:"bool_to_str"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.IntToStr, "123")
		_assert.Equal(t, v.Int64ToStr, "456")
		_assert.Equal(t, v.UintToStr, "789")
		_assert.Equal(t, v.FloatToStr, "3.14")
		_assert.Equal(t, v.BoolToStr, "true")
	}
	{
		data := map[string]interface{}{
			"bool_true":         true,
			"bool_false":        false,
			"int_to_bool_true":  int(1),
			"int_to_bool_false": int(0),
			"str_to_bool_true":  "non-empty",
			"str_to_bool_false": "",
		}
		v := struct {
			BoolTrue       bool `json:"bool_true"`
			BoolFalse      bool `json:"bool_false"`
			IntToBoolTrue  bool `json:"int_to_bool_true"`
			IntToBoolFalse bool `json:"int_to_bool_false"`
			StrToBoolTrue  bool `json:"str_to_bool_true"`
			StrToBoolFalse bool `json:"str_to_bool_false"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.BoolTrue, true)
		_assert.Equal(t, v.BoolFalse, false)
		_assert.Equal(t, v.IntToBoolTrue, true)
		_assert.Equal(t, v.IntToBoolFalse, false)
		_assert.Equal(t, v.StrToBoolTrue, true)
		_assert.Equal(t, v.StrToBoolFalse, false)
	}
	{
		data := map[string]interface{}{
			"bytes":     []byte("hello"),
			"str_bytes": "world",
		}
		v := struct {
			Bytes    []byte `json:"bytes"`
			StrBytes []byte `json:"str_bytes"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, string(v.Bytes), "hello")
		_assert.Equal(t, string(v.StrBytes), "world")
	}
	{
		data := map[string]interface{}{
			"zero_int":       0,
			"zero_float":     0.0,
			"zero_str":       "0",
			"negative_int":   -100,
			"negative_float": -3.14,
			"negative_str":   "-50",
			"large_int":      int64(9223372036854775807),
		}
		v := struct {
			ZeroInt       int     `json:"zero_int"`
			ZeroFloat     float64 `json:"zero_float"`
			ZeroStr       int     `json:"zero_str"`
			NegativeInt   int     `json:"negative_int"`
			NegativeFloat float64 `json:"negative_float"`
			NegativeStr   int     `json:"negative_str"`
			LargeInt      int64   `json:"large_int"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ZeroInt, 0)
		_assert.Equal(t, v.ZeroFloat, 0.0)
		_assert.Equal(t, v.ZeroStr, 0)
		_assert.Equal(t, v.NegativeInt, -100)
		_assert.Equal(t, v.NegativeFloat, -3.14)
		_assert.Equal(t, v.NegativeStr, -50)
		_assert.Equal(t, v.LargeInt, int64(9223372036854775807))
	}
}
func TestDo_TypeConversions_NestedStruct(t *testing.T) {
	type NestedStruct struct {
		ID      int     `json:"id"`
		Price   float64 `json:"price"`
		Enabled bool    `json:"enabled"`
		Name    string  `json:"name"`
	}
	{
		data := map[string]interface{}{
			"product": map[string]interface{}{
				"id":      "123",
				"price":   "99.99",
				"enabled": "true",
				"name":    "Product Name",
			},
		}
		v := struct {
			Product NestedStruct `json:"product"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Product.ID, 123)
		_assert.Equal(t, v.Product.Price, 99.99)
		_assert.Equal(t, v.Product.Enabled, true)
		_assert.Equal(t, v.Product.Name, "Product Name")
	}
	{
		data := map[string]interface{}{
			"product": map[string]interface{}{
				"id":      int8(100),
				"price":   float32(3.14),
				"enabled": int(1),
				"name":    int(12345),
			},
		}
		v := struct {
			Product NestedStruct `json:"product"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Product.ID, 100)
		_assert.EqualByFloat(t, v.Product.Price, 3.140000104904175)
		_assert.Equal(t, v.Product.Enabled, true)
		_assert.Equal(t, v.Product.Name, "12345")
	}
	{
		type Level2 struct {
			Value int `json:"value"`
		}
		type Level1 struct {
			Level2 Level2 `json:"level2"`
		}
		data := map[string]interface{}{
			"level1": map[string]interface{}{
				"level2": map[string]interface{}{
					"value": "999",
				},
			},
		}
		v := struct {
			Level1 Level1 `json:"level1"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Level1.Level2.Value, 999)
	}
	{
		type MixedStruct struct {
			ID       int     `json:"id"`
			Price    float64 `json:"price"`
			Quantity string  `json:"quantity"`
			Total    int     `json:"total"`
		}
		data := map[string]interface{}{
			"order": map[string]interface{}{
				"id":       "100",
				"price":    "199.99",
				"quantity": 5,
				"total":    "999",
			},
		}
		v := struct {
			Order MixedStruct `json:"order"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Order.ID, 100)
		_assert.Equal(t, v.Order.Price, 199.99)
		_assert.Equal(t, v.Order.Quantity, "5")
		_assert.Equal(t, v.Order.Total, 999)
	}
}
func TestDo_MoreTypeConversions(t *testing.T) {
	{
		data := map[string]interface{}{
			"int8":   int8(8),
			"int16":  int16(16),
			"int32":  int32(32),
			"uint":   uint(100),
			"uint8":  uint8(8),
			"uint16": uint16(16),
			"uint32": uint32(32),
		}
		v := struct {
			Int8   int  `json:"int8"`
			Int16  int  `json:"int16"`
			Int32  int  `json:"int32"`
			Uint   uint `json:"uint"`
			Uint8  uint `json:"uint8"`
			Uint16 uint `json:"uint16"`
			Uint32 uint `json:"uint32"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Int8, 8)
		_assert.Equal(t, v.Int16, 16)
		_assert.Equal(t, v.Int32, 32)
		_assert.Equal(t, v.Uint, uint(100))
		_assert.Equal(t, v.Uint8, uint(8))
		_assert.Equal(t, v.Uint16, uint(16))
		_assert.Equal(t, v.Uint32, uint(32))
	}
	{
		data := map[string]interface{}{
			"float32": float32(3.14),
			"float64": float64(2.718),
		}
		v := struct {
			Float32 float64 `json:"float32"`
			Float64 float64 `json:"float64"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.EqualByFloat(t, v.Float32, 3.140000104904175)
		_assert.Equal(t, v.Float64, 2.718)
	}
	{
		data := map[string]interface{}{
			"bytes": []byte("hello"),
			"str":   "world",
		}
		v := struct {
			Bytes []byte `json:"bytes"`
			Str   string `json:"str"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, string(v.Bytes), "hello")
		_assert.Equal(t, v.Str, "world")
	}
	{
		data := map[string]interface{}{
			"id":    "123",
			"price": "99.99",
			"flag":  "true",
		}
		v := struct {
			ID    int     `json:"id"`
			Price float64 `json:"price"`
			Flag  bool    `json:"flag"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 123)
		_assert.Equal(t, v.Price, 99.99)
		_assert.Equal(t, v.Flag, true)
	}
}
func TestDo_DeepNesting(t *testing.T) {
	type Level3 struct {
		Value string `json:"value"`
	}
	type Level2 struct {
		Level3 Level3 `json:"level3"`
	}
	type Level1 struct {
		Level2 Level2 `json:"level2"`
	}
	{
		data := map[string]interface{}{
			"level2": map[string]interface{}{
				"level3": map[string]interface{}{
					"value": "deep",
				},
			},
		}
		v := struct {
			Level1 Level1 `json:"level1"`
		}{}
		data2 := map[string]interface{}{
			"level1": data,
		}
		err := Do(&v, data2)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Level1.Level2.Level3.Value, "deep")
	}
}
func TestDo_StructFieldWithNilValue(t *testing.T) {
	{
		data := map[string]interface{}{
			"id":   1,
			"info": nil,
		}
		v := struct {
			ID   int         `json:"id"`
			Info InnerStruct `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Info.Name, "")
		_assert.Equal(t, v.Info.Value, 0)
	}
}
func TestDo_NonPointerStruct(t *testing.T) {
	{
		data := map[string]interface{}{
			"id":   1,
			"name": "Test",
		}
		v := struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{}
		err := Do(v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 0)
		_assert.Equal(t, v.Name, "")
	}
}
func TestDo_NoJsonTag(t *testing.T) {
	{
		data := map[string]interface{}{
			"ID":   1,
			"Name": "Test",
		}
		v := struct {
			ID   int
			Name string
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Name, "Test")
	}
}
func TestDo_StructFieldNotMap(t *testing.T) {
	{
		data := map[string]interface{}{
			"id":   1,
			"info": "not_a_map",
		}
		v := struct {
			ID   int         `json:"id"`
			Info InnerStruct `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Info.Name, "")
		_assert.Equal(t, v.Info.Value, 0)
	}
}
func TestDo_MapFieldNotMap(t *testing.T) {
	{
		data := map[string]interface{}{
			"info": "not_a_map",
		}
		v := struct {
			Info map[string]interface{} `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Info, map[string]interface{}(nil))
	}
}
func TestDo_MapFieldNonStringKey(t *testing.T) {
	{
		data := map[string]interface{}{
			"info": map[string]interface{}{"a": 1},
		}
		v := struct {
			Info map[int]string `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Info, map[int]string(nil))
	}
}
func TestDo_MapFieldUnsupportedElemType(t *testing.T) {
	{
		data := map[string]interface{}{
			"info": map[string]interface{}{"a": true, "b": false},
		}
		v := struct {
			Info map[string]bool `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, len(v.Info), 0)
	}
}
func TestDo_MapFieldNilInitialization(t *testing.T) {
	{
		data := map[string]interface{}{
			"info": map[string]interface{}{"a": 1, "b": 2},
		}
		v := struct {
			Info map[string]int `json:"info"`
		}{}
		_assert.Equal(t, v.Info, map[string]int(nil))
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Info["a"], 1)
		_assert.Equal(t, v.Info["b"], 2)
	}
}
func TestDo_NonByteSlice(t *testing.T) {
	{
		data := map[string]interface{}{
			"items": []int{1, 2, 3},
		}
		v := struct {
			Items []int `json:"items"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Items, []int(nil))
	}
}
func TestDo_UnsupportedFieldType(t *testing.T) {
	{
		data := map[string]interface{}{
			"ID": 1,
			"Ch": make(chan int),
		}
		v := struct {
			ID int
			Ch chan int
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Ch, (chan int)(nil))
	}
}
func TestDo_FieldCannotSet(t *testing.T) {
	{
		data := map[string]interface{}{
			"id":   1,
			"name": "Test",
		}
		v := struct {
			ID   int    `json:"id"`
			name string `json:"name"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.name, "")
	}
}
func TestDo_MapFieldWithUnsupportedElemInLoop(t *testing.T) {
	{
		data := map[string]interface{}{
			"info": map[string]interface{}{
				"a": 1,
				"b": []int{1, 2},
				"c": 2,
			},
		}
		v := struct {
			Info map[string]int `json:"info"`
		}{}
		err := Do(&v, data)
		_assert.NoError(t, err)
		_assert.Equal(t, v.Info["a"], 1)
		_assert.Equal(t, v.Info["b"], 0)
		_assert.Equal(t, v.Info["c"], 2)
	}
}
func TestDo_MultipleBindings(t *testing.T) {
	{
		v := struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{}
		data1 := map[string]interface{}{
			"id":   1,
			"name": "Alice",
		}
		err := Do(&v, data1)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Name, "Alice")
		_assert.Equal(t, v.Age, 0)
		data2 := map[string]interface{}{
			"age": 25,
		}
		err = Do(&v, data2)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Name, "Alice")
		_assert.Equal(t, v.Age, 25)
		data3 := map[string]interface{}{
			"id":   2,
			"name": "Bob",
			"age":  30,
		}
		err = Do(&v, data3)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 2)
		_assert.Equal(t, v.Name, "Bob")
		_assert.Equal(t, v.Age, 30)
	}
	{
		type Info struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		}
		v := struct {
			ID   int  `json:"id"`
			Info Info `json:"info"`
		}{}
		data1 := map[string]interface{}{
			"id": 1,
			"info": map[string]interface{}{
				"name":  "Test1",
				"value": 10,
			},
		}
		err := Do(&v, data1)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Info.Name, "Test1")
		_assert.Equal(t, v.Info.Value, 10)
		data2 := map[string]interface{}{
			"info": map[string]interface{}{
				"name": "Test2",
			},
		}
		err = Do(&v, data2)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Info.Name, "Test2")
		_assert.Equal(t, v.Info.Value, 10)
		data3 := map[string]interface{}{
			"id": 2,
			"info": map[string]interface{}{
				"value": 20,
			},
		}
		err = Do(&v, data3)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 2)
		_assert.Equal(t, v.Info.Name, "Test2")
		_assert.Equal(t, v.Info.Value, 20)
	}
	{
		v := struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{}
		data1 := map[string]interface{}{
			"id":   1,
			"name": "First",
		}
		err := Do(&v, data1)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 1)
		_assert.Equal(t, v.Name, "First")
		data2 := map[string]interface{}{
			"id":   2,
			"name": "Second",
		}
		err = Do(&v, data2)
		_assert.NoError(t, err)
		_assert.Equal(t, v.ID, 2)
		_assert.Equal(t, v.Name, "Second")
	}
}
