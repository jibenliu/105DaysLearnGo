package dispatch

import (
	"fmt"
	"goProjects/daylearning/others/grpcDemo/order"
	"sync"
	"time"
)

// OrderDispatcher 是一个守护进程，它使用 sync 创建一个工作池。waitGroup 并发地
// 处理和分发订单
type OrderDispatcher struct {
	ordersCh   chan *order.Order
	orderLimit int // 并发处理的最大订单数
}

// NewOrderDispatcher 创建一个新的 OrderDispatcher
func NewOrderDispatcher(orderLimit int, bufferSize int) OrderDispatcher {
	return OrderDispatcher{
		ordersCh:   make(chan *order.Order, bufferSize), // initiliaze as a buffered channel
		orderLimit: orderLimit,
	}
}

// SubmitOrder 提交订单进行处理
func (d OrderDispatcher) SubmitOrder(order *order.Order) {
	go func() {
		d.ordersCh <- order
	}()
}

// Start 在后台启动调度程序
func (d OrderDispatcher) Start() {
	go d.processOrders()
}

// Shutdown 通过关闭订单来关闭 OrderDispatcher
// 注意：这个函数应该只在最后一个订单到达订单通道之后才执行。
// 向一个封闭的通道提交命令会引起 panic。
func (d OrderDispatcher) Shutdown() {
	close(d.ordersCh)
}

// processOrders 使用“for range”和一个 sync.waitGroup 在后台处理所有传入的订单
func (d OrderDispatcher) processOrders() {
	limiter := make(chan struct{}, d.orderLimit)
	var wg sync.WaitGroup
	// 连续地处理从订单通道接收到的订单
	// 当通道关闭时，此循环将终止
	for ord := range d.ordersCh {
		limiter <- struct{}{}
		wg.Add(1)
		go func(order *order.Order) {
			// TODO: 触发执行流程，将订单组装成一个包裹并发货，
			// 这里我们 sleep 并打印
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("Order (%v) has shipped \n", order)
			<-limiter
			wg.Done()
		}(ord)
	}
	wg.Wait()
}
func main() {
	dispatcher := NewOrderDispatcher(3, 100)
	dispatcher.Start()
	defer dispatcher.Shutdown()
	dispatcher.SubmitOrder(&order.Order{Items: []*order.Item{{Description: "iPhone Screen Protector", Price: 9.99}}})
	dispatcher.SubmitOrder(&order.Order{Items: []*order.Item{{Description: "iPhone Case", Price: 19.99}}})
	dispatcher.SubmitOrder(&order.Order{Items: []*order.Item{{Description: "Pixel Case", Price: 14.99}}})
	dispatcher.SubmitOrder(&order.Order{Items: []*order.Item{{Description: "Bluetooth Speaker", Price: 29.99}}})
	dispatcher.SubmitOrder(&order.Order{Items: []*order.Item{{Description: "4K Monitor", Price: 159.99}}})
	dispatcher.SubmitOrder(&order.Order{Items: []*order.Item{{Description: "Inkjet Printer", Price: 79.99}}})

	time.Sleep(5 * time.Second) // 仅为了测试
}
