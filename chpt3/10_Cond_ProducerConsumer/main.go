package main

import (
	"fmt"
	"sync"
)

// 并发场景：夏天喝Cola需要冰块，从制冰机取冰块，制冰机是Producer，取冰的人是Consumer
// 制冰机在机箱满的时候要停止生产，机箱空的时候，取冰的人需要停止取冰。

// 最多冰块数，制冰机的容量
const maxCnt = 3

// 最少冰块数，代表杯子中没有冰块了
const minCnt = 0

type iceCube int

type cup struct {
	iceCubes []iceCube
}

// 多协程通过sync Cond来进行生产者消费者的模拟
// 类比Java 加锁 + 循环&&等待 + 唤醒，在Go中也是，经典范式
// Q: 为什么加锁
// A：
//  1. 加锁 获得程序执行权
//  2. 不加锁情况下，如果生产冰块同时还能从杯子中拿出冰块，万一生产速率 < 取出速率，杯子就空了；反之，杯子益出冰块，都是非预期情况
//
// Q：为什么需要cond.Wait(), cond.Signal()?
// A：因为要通知阻塞的协程重新获取执行权(获取锁)
func ProducerAndConsumerSimulationWithSyncCond() {
	stopCh := make(chan struct{})

	lc := new(sync.Mutex)
	cond := sync.NewCond(lc)

	cup := cup{
		iceCubes: make([]iceCube, 3, 3),
	}

	// Concumser
	go func() {
		for {
			cond.L.Lock()
			for len(cup.iceCubes) == minCnt {
				cond.Wait()
			}
			// 删除头部的冰块
			cup.iceCubes = cup.iceCubes[1:]
			fmt.Println("consume 1 iceCube, left iceCubes -> ", len(cup.iceCubes))
			cond.Signal() // 唤醒producer
			cond.L.Unlock()
		}
	}()

	// Producer
	go func() {
		for {
			cond.L.Lock()
			for len(cup.iceCubes) == maxCnt {
				// 当已经打到最多冰块时，停止制冰
				cond.Wait()
			}
			// 杯子中新添加进一个冰块
			cup.iceCubes = append(cup.iceCubes, 1)
			fmt.Println("producer 1 iceCube, left iceCubes ", len(cup.iceCubes))
			cond.Signal() // 唤醒等待取冰的协程
			cond.L.Unlock()
		}
	}()

	for {
		select {
		case <-stopCh:
			return
		default:
		}
	}
}

func main() {

}
