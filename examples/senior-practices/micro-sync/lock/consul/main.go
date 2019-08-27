package main

import (
	"github.com/micro/go-micro/sync/lock"
	"github.com/micro/go-micro/sync/lock/consul"
	"github.com/micro/go-micro/util/log"
	"time"
)

func main() {
	// 地址
	nodes := lock.Nodes("127.0.0.1:8500")

	resourceId := "id"

	go func() {
		lc := consul.NewLock(nodes)
		log.Logf("协程一获取锁...")
		// 获取锁
		err := lc.Acquire(resourceId)
		if err != nil {
			log.Logf("[ERR] 协程一未得到锁")
			return
		}

		log.Logf("协程一得到锁，等一秒")
		time.Sleep(1 * time.Second)

		// 释放锁
		log.Logf("协程一释放锁")
		err = lc.Release(resourceId)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		lc := consul.NewLock(nodes)
		log.Logf("协程二获取锁...")
		// 获取锁
		err := lc.Acquire(resourceId)
		if err != nil {
			log.Logf("[ERR] 协程二未得到锁")
			return
		}

		log.Logf("协程二得到锁，等一秒")
		time.Sleep(1 * time.Second)

		// 释放锁
		log.Logf("协程二释放锁")
		err = lc.Release(resourceId)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// 等协程跑完
	time.Sleep(5 * time.Second)
}
