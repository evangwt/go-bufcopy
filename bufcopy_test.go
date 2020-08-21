package bufcopy

import (
	"bytes"
	"io"
	"testing"
)

func TestBufCopy(t *testing.T) {
	buf := New()
	str := bytes.Repeat([]byte("ABC"), 100)
	src := bytes.NewReader(str)
	var dst bytes.Buffer
	_, err := buf.Copy(&dst, src)
	if err != nil {
		t.Fatal(err)
	}
	if dst.String() != string(str) {
		t.Fatal("incorrect content")
	}
}

func BenchmarkIoCopySmall(b *testing.B) {
	str := bytes.Repeat([]byte{1}, 512)
	src := bytes.NewReader(str)
	var dst bytes.Buffer

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		io.Copy(&dst, src)
		src.Reset(str)
	}
	b.ReportAllocs()
}

func BenchmarkBufCopySmall(b *testing.B) {
	buf := New()
	str := bytes.Repeat([]byte{1}, 512)
	src := bytes.NewReader(str)
	var dst bytes.Buffer

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Copy(&dst, src)
		src.Reset(str)
	}
	b.ReportAllocs()
}

func BenchmarkIoCopyLarge(b *testing.B) {
	str := bytes.Repeat([]byte{1}, 32*1024)
	src := bytes.NewReader(str)
	var dst bytes.Buffer

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		io.Copy(&dst, src)
		src.Reset(str)
	}
	b.ReportAllocs()
}

func BenchmarkBufCopyLarge(b *testing.B) {
	buf := New()
	str := bytes.Repeat([]byte{1}, 32*1024)
	src := bytes.NewReader(str)
	var dst bytes.Buffer

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Copy(&dst, src)
		src.Reset(str)
	}
	b.ReportAllocs()
}
