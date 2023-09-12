package main

import "sync"

func main() {

}

type singleton struct {
}

var (
	once     sync.Once
	instance *singleton
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
