package gorm

import (
	"runtime"
	"unicode/utf8"
	"unsafe"
)

type d struct {
	data []byte
}

type Builder struct {
	d *d
}

func (this_ *Builder) String() string {
	if this_.d == nil {
		return ""
	}
	return unsafe.String(unsafe.SliceData(this_.d.data), len(this_.d.data))
}

func (this_ *Builder) Len() int {
	if this_.d == nil {
		return 0
	}
	return len(this_.d.data)
}

func (this_ *Builder) Cap() int {
	if this_.d == nil {
		return 0
	}
	return cap(this_.d.data)
}

func (this_ *Builder) Reset() {
	if this_.d == nil {
		return
	}
	this_.d.data = this_.d.data[:0]
}

func (this_ *Builder) grow() {
	if cap(this_.d.data) == 0 {
		this_.d.data = GetSample(40960)[:0]
		runtime.SetFinalizer(this_.d, func(d *d) {
			PutSample(d.data)
		})
	} else {
		tmp := GetSample(cap(this_.d.data) * 2)
		copy(tmp, this_.d.data)
		PutSample(this_.d.data)
		this_.d.data = tmp
	}
}

func (this_ *Builder) Grow(n int) {
	if n < 0 {
		panic("strings.Builder.Grow: negative count")
	}
	if this_.d == nil {
		this_.d = &d{}
	}
	if cap(this_.d.data)-len(this_.d.data) < n {
		this_.grow()
	}
}

func (this_ *Builder) Write(p []byte) (int, error) {
	this_.Grow(len(p))
	this_.d.data = append(this_.d.data, p...)
	return len(p), nil
}

func (this_ *Builder) WriteByte(c byte) error {
	this_.Grow(1)
	this_.d.data = append(this_.d.data, c)
	return nil
}

func (this_ *Builder) WriteRune(r rune) (int, error) {
	this_.Grow(4)
	n := len(this_.d.data)
	this_.d.data = utf8.AppendRune(this_.d.data, r)
	return len(this_.d.data) - n, nil
}

func (this_ *Builder) WriteString(s string) (int, error) {
	this_.Grow(len(s))
	this_.d.data = append(this_.d.data, s...)
	return len(s), nil
}
