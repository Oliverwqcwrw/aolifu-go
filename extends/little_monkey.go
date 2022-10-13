package main

import "fmt"

type Monkey struct {
	Name string
}

func (monkey *Monkey) climbing() {
	fmt.Println("monkey climbing")
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

type LittleMonkey struct {
	Monkey
}

func (littleMonkey *LittleMonkey) Flying() {
	fmt.Println("littleMonkey can fly")
}

func (littleMonkey *LittleMonkey) Swimming() {
	fmt.Println("littleMonkey can swim")
}
