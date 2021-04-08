package pool

import (
	"testing"
)
import "sync"

type AFreeCoder struct {
	officialAccount string
	article         string
	content         []string
	placeHolder     string
}

// 为了真实模拟，这里禁止编译器使用内联优化
////go:noinline
func NewAFreeCoder() *AFreeCoder {
	return &AFreeCoder{
		officialAccount: "码农的自由之路",
		content:         make([]string, 10000, 10000),
		placeHolder:     "如果觉得有用，欢迎关注哦~",
	}
}

func (a *AFreeCoder) Write() {
	a.article = "Go 并发之性能提升杀器 Pool"
}

func f(concurrentNum int) {
	var w sync.WaitGroup
	w.Add(concurrentNum)
	for i := 0; i < concurrentNum; i++ {
		go func() {
			defer w.Done()
			a := NewAFreeCoder()
			a.Write()
		}()
	}
	w.Wait()
}

func fUsePool(concurrentNum int) {
	var w sync.WaitGroup
	p := sync.Pool{
		New: func() interface{} {
			return NewAFreeCoder()
		},
	}

	w.Add(concurrentNum)
	for i := 0; i < concurrentNum; i++ {
		go func() {
			defer w.Done()
			a := p.Get().(*AFreeCoder)
			defer p.Put(a)
			a.Write()
		}()
	}
	w.Wait()
}

//注意命名规范 Benchmark+首字母大写的方法名 参数固定
func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(i)
	}
}
//注意命名规范 Benchmark+首字母大写的方法名 参数固定
func BenchmarkFUserPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fUsePool(i)
	}
}
