package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating Single Instance.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single Instance allready created-1")
		}
	} else {
		fmt.Println("Single Instance allready created-2")
	}
	return singleInstance
}
