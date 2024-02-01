package gorm

import (
	"unicode/utf8"
	"unsafe"
)

type Builder struct {
	data []byte
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (this_ *Builder) String() string {
	return unsafe.String(unsafe.SliceData(this_.data), len(this_.data))
}

func (this_ *Builder) Len() int {
	return len(this_.data)
}

func (this_ *Builder) Cap() int {
	return cap(this_.data)
}

func (this_ *Builder) Reset() {
	this_.data = this_.data[:0]
}

func (this_ *Builder) grow() {
	if cap(this_.data) == 0 {
		this_.data = make([]byte, 0, 40960)
	} else {
		tmp := make([]byte, len(this_.data), cap(this_.data)*2)
		copy(tmp, this_.data)
		this_.data = tmp
	}
}

func (this_ *Builder) Grow(n int) {
	if n < 0 {
		panic("strings.Builder.Grow: negative count")
	}
	if cap(this_.data)-len(this_.data) < n {
		this_.grow()
	}
}

func (this_ *Builder) Write(p []byte) (int, error) {
	this_.Grow(len(p))
	this_.data = append(this_.data, p...)
	return len(p), nil
}

func (this_ *Builder) WriteByte(c byte) error {
	this_.Grow(1)
	this_.data = append(this_.data, c)
	return nil
}

func (this_ *Builder) WriteRune(r rune) (int, error) {
	this_.Grow(4)
	n := len(this_.data)
	this_.data = utf8.AppendRune(this_.data, r)
	return len(this_.data) - n, nil
}

func (this_ *Builder) WriteString(s string) (int, error) {
	this_.Grow(len(s))
	this_.data = append(this_.data, s...)
	return len(s), nil
}
