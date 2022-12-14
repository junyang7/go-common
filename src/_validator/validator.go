package _validator

import (
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/junyang7/go-common/src/_is"
	"github.com/junyang7/go-common/src/_slice"
	"strings"
)

type Int struct {
	name        string
	value       int
	codeMessage *_codeMessage.CodeMessage
}

func NewInt(name string, value interface{}) *Int {
	return &Int{
		name:        name,
		value:       _as.Int(value),
		codeMessage: _codeMessage.ErrParameter,
	}
}
func (this *Int) Default(value int) *Int {
	if _is.Empty(this.value) {
		this.value = value
	}
	return this
}
func (this *Int) CodeMessage(codeMessage *_codeMessage.CodeMessage) *Int {
	this.codeMessage = codeMessage
	return this
}
func (this *Int) EnsureMin(min int) *Int {
	_interceptor.Insure(this.value >= min).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Int) EnsureMax(max int) *Int {
	_interceptor.Insure(this.value <= max).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Int) EnsureBetween(min int, max int) *Int {
	_interceptor.Insure(this.value >= min && this.value <= max).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Int) EnsureLength(length int) *Int {
	l := len(_as.String(this.value))
	_interceptor.Insure(l == length).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Int) EnsureLengthMin(min int) *Int {
	l := len(_as.String(this.value))
	_interceptor.Insure(l >= min).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Int) EnsureLengthMax(max int) *Int {
	l := len(_as.String(this.value))
	_interceptor.Insure(l <= max).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Int) EnsureLengthBetween(min int, max int) *Int {
	l := len(_as.String(this.value))
	_interceptor.Insure(l >= min && l <= max).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Int) EnsureIn(value ...int) *Int {
	_interceptor.Insure(_slice.In(this.value, value)).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Int) Value() int {
	return this.value
}

type String struct {
	name        string
	value       string
	codeMessage *_codeMessage.CodeMessage
}

func NewString(name string, value interface{}) *String {
	return &String{
		name:        name,
		value:       strings.TrimSpace(_as.String(value)),
		codeMessage: _codeMessage.ErrParameter,
	}
}
func (this *String) Default(value string) *String {
	if _is.Empty(this.value) {
		this.value = value
	}
	return this
}
func (this *String) CodeMessage(codeMessage *_codeMessage.CodeMessage) *String {
	this.codeMessage = codeMessage
	return this
}
func (this *String) EnsureEmpty() *String {
	_interceptor.Insure(_is.Empty(this.value)).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *String) EnsureNotEmpty() *String {
	_interceptor.Insure(!_is.Empty(this.value)).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *String) EnsureLength(length int) *String {
	l := len(this.value)
	_interceptor.Insure(l == length).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *String) EnsureLengthMin(min int) *String {
	l := len(this.value)
	_interceptor.Insure(l >= min).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *String) EnsureLengthMax(max int) *String {
	l := len(this.value)
	_interceptor.Insure(l <= max).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *String) EnsureLengthBetween(min int, max int) *String {
	l := len(this.value)
	_interceptor.Insure(l >= min && l <= max).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *String) EnsureIn(value ...string) *String {
	_interceptor.Insure(_slice.In(this.value, value)).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *String) Value() string {
	return this.value
}

type Bool struct {
	name        string
	value       bool
	codeMessage *_codeMessage.CodeMessage
}

func NewBool(name string, value interface{}) *Bool {
	return &Bool{
		name:        name,
		value:       _as.Bool(value),
		codeMessage: _codeMessage.ErrParameter,
	}
}
func (this *Bool) Default(value bool) *Bool {
	if _is.Empty(this.value) {
		this.value = value
	}
	return this
}
func (this *Bool) CodeMessage(codeMessage *_codeMessage.CodeMessage) *Bool {
	this.codeMessage = codeMessage
	return this
}
func (this *Bool) EnsureIn(value ...bool) *Bool {
	_interceptor.Insure(_slice.In(this.value, value)).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Bool) Value() bool {
	return this.value
}

type Float64 struct {
	name        string
	value       float64
	codeMessage *_codeMessage.CodeMessage
}

func NewFloat64(name string, value interface{}) *Float64 {
	return &Float64{
		name:        name,
		value:       _as.Float64(value),
		codeMessage: _codeMessage.ErrParameter,
	}
}
func (this *Float64) Default(value float64) *Float64 {
	if _is.Empty(this.value) {
		this.value = value
	}
	return this
}
func (this *Float64) CodeMessage(codeMessage *_codeMessage.CodeMessage) *Float64 {
	this.codeMessage = codeMessage
	return this
}
func (this *Float64) EnsureMin(min float64) *Float64 {
	_interceptor.Insure(this.value >= min).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Float64) EnsureMax(max float64) *Float64 {
	_interceptor.Insure(this.value <= max).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Float64) EnsureBetween(min float64, max float64) *Float64 {
	_interceptor.Insure(this.value >= min && this.value <= max).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Float64) EnsureLength(length int) *Float64 {
	l := len(_as.String(this.value))
	_interceptor.Insure(l == length).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Float64) EnsureLengthMin(min int) *Float64 {
	l := len(_as.String(this.value))
	_interceptor.Insure(l >= min).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Float64) EnsureLengthMax(max int) *Float64 {
	l := len(_as.String(this.value))
	_interceptor.Insure(l <= max).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Float64) EnsureLengthBetween(min int, max int) *Float64 {
	l := len(_as.String(this.value))
	_interceptor.Insure(l >= min && l <= max).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Float64) EnsureIn(value ...float64) *Float64 {
	_interceptor.Insure(_slice.In(this.value, value)).
		CodeMessage(this.codeMessage).
		Data(map[string]interface{}{this.name: this.value}).
		Do()
	return this
}
func (this *Float64) Value() float64 {
	return this.value
}
