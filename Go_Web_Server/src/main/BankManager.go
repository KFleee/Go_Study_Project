package main

import (
	"time"
	"sync"
)

type Lock struct {
	lock *sync.RWMutex
	timeAccessed time.Time "last access time"
}

type Bank struct {
	lock sync.RWMutex
	userLock map[string]Lock
	MaxLifeTime int64 "用户锁最长存活时间s"
	GcLifeTime int64 "垃圾回收周期s"
}

func (bank *Bank) LockRead(userId string) (Lock, error) {
	bank.lock.RLock()
	if userLock, ok := bank.userLock[userId]; ok {
		go bank.LockUpdate(userId)
		bank.lock.RUnlock()
		return userLock, nil
	}
	bank.lock.RUnlock()
	bank.lock.Lock()
	if userLock, ok := bank.userLock[userId]; ok {
		go bank.LockUpdate(userId)
		bank.lock.Unlock()
		return userLock, nil
	}
	var lock sync.RWMutex
	var timeAccessed time.Time = time.Now()
	var userLock Lock = Lock{
		lock: &lock,
		timeAccessed: timeAccessed,
	}
	bank.userLock[userId] = userLock
	bank.lock.Unlock()
	return userLock, nil
}

func (bank *Bank) LockUpdate(userId string) error {
	bank.lock.Lock()
	defer bank.lock.Unlock()
	if userLock, ok := bank.userLock[userId]; ok {
		userLock.timeAccessed = time.Now()
		bank.userLock[userId] = userLock
		return nil
	}
	return nil
}

func (bank *Bank) LockGc() {
	bank.lock.RLock()
	for k, v := range bank.userLock {
		if (v.timeAccessed.Unix() + bank.MaxLifeTime) < time.Now().Unix() {
			bank.lock.RUnlock()
			bank.lock.Lock()
			delete(bank.userLock, k)
			bank.lock.Unlock()
			bank.lock.RLock()
		}
	}
	bank.lock.RUnlock()
}

func (bank *Bank) Gc() {
	bank.LockGc()
	time.AfterFunc(time.Duration(bank.GcLifeTime) * time.Second, func() {bank.Gc()})
}