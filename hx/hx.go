package hx

import (
	"fmt"

	"github.com/a-h/templ"
)

type HX map[string]string

func (h HX) Set(attr, val string) {
	h[fmt.Sprintf("hx-%s", attr)] = val
}

func New(opts ...func(HX)) HX {
	hx := HX{}
	for _, opt := range opts {
		opt(hx)
	}
	return hx
}

// REST

func GET(path string) func(HX) {
	return func(h HX) { h.Set("get", string(templ.URL(path))) }
}

func POST(path string) func(HX) {
	return func(h HX) { h.Set("post", string(templ.URL(path))) }
}

func PUT(path string) func(HX) {
	return func(h HX) { h.Set("put", string(templ.URL(path))) }
}

func DELETE(path string) func(HX) {
	return func(h HX) { h.Set("delete", string(templ.URL(path))) }
}

// Core

func PushUrl(value string) func(HX) {
	return func(h HX) { h.Set("push-url", value) }
}

func Select(value string) func(HX) {
	return func(h HX) { h.Set("select", value) }
}

func SelectOob(value string) func(HX) {
	return func(h HX) { h.Set("select", value) }
}

func Swap(value string) func(HX) {
	return func(h HX) { h.Set("swap", value) }
}

func SwapOob(value string) func(HX) {
	return func(h HX) { h.Set("swap-oob", value) }
}

func Target(value string) func(HX) {
	return func(h HX) { h.Set("target", value) }
}

func Trigger(value string) func(HX) {
	return func(h HX) { h.Set("trigger", value) }
}

func Vals(value string) func(HX) {
	return func(h HX) { h.Set("vals", value) }
}
