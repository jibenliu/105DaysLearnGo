package validate

import (
	"context"
	"errors"
	"goProjects/daylearning/others/grpcDemo/order"
	"golang.org/x/sync/errgroup"
	"time"
)

var (
	ErrPreAuthorizationTimeout = errors.New("pre-authorization request timeout")
	ErrInventoryRequestTimeout = errors.New("check inventory request timeout")
	ErrItemOutOfStock          = errors.New("sorry one or more items in your order is out of stock")
)

// preAuthorizePayment 对支付方式进行预授权并返回错误。
// 如果预先授权成功，则返回 nil
func preAuthorizePayment(ctx context.Context, payment *order.PaymentMethod, orderAmount float32) error {
	// 在这里执行昂贵的授权逻辑——在这个例子中我们使用 sleep
	// 并返回 nil 来表示成功的授权
	timer := time.NewTimer(3 * time.Second)
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ErrPreAuthorizationTimeout
	}
}

// checkInventory 返回一个布尔值和一个错误，表示是否所有商品是否都有库存
//(true, nil) 表示所有商品都有库存并且没有遇到错误
func checkInventory(ctx context.Context, items []*order.Item) (bool, error) {
	// 在这里执行昂贵的库存检查逻辑 - 在这个例子中我们使用 sleep
	timer := time.NewTimer(2 * time.Second)
	select {
	case <-timer.C:
		return true, nil
	case <-ctx.Done():
		return false, ErrInventoryRequestTimeout
	}
}

// getOrderTotal 计算订单总数
func getOrderTotal(items []*order.Item) float32 {
	var total float32
	for _, item := range items {
		total += item.Price
	}
	return total
}
func validateOrder(ctx context.Context, items []*order.Item, payment *order.PaymentMethod) error {
	g, errCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return preAuthorizePayment(errCtx, payment, getOrderTotal(items))
	})
	g.Go(func() error {
		itemsInStock, err := checkInventory(errCtx, items)
		if err != nil {
			return err
		}
		if !itemsInStock {
			return ErrItemOutOfStock
		}
		return nil
	})
	return g.Wait()
}
