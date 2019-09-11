package main

import "fmt"

type Logger struct{}

func (logger *Logger) Init(packagename string) error {
	return nil
}
func (logger *Logger) Log(log string) {
	fmt.Println(log)
}
func (logger *Logger) Debug(log string) {
	fmt.Println(log)
}
func (logger *Logger) Fatal(err error) {
	fmt.Println(err)
}
