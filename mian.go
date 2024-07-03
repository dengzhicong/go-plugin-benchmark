package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/pkujhd/goloader"
)

func main() {
	linker, err := goloader.ReadObjs([]string{"goloader.o"}, []string{"main"})
	if err != nil {
		panic(err)
	}

	run := "main.RandInt"

	symPtr := make(map[string]uintptr)
	err = goloader.RegSymbol(symPtr)

	if err != nil {
		panic(err)
	}

	codeModule, err := goloader.Load(linker, symPtr)
	if err != nil {
		panic(err)
	}

	a := linker.SymMap["go-plugin-benchmark/fib.tailRecursiveFibonacci"]
	if a != nil {
		fmt.Println(a)
	}

	runFuncPtr, ok := codeModule.Syms[run]
	if !ok || runFuncPtr == 0 {
		fmt.Println("Load error! not find function:", run)
		return
	}
	funcPtrContainer := (uintptr)(unsafe.Pointer(&runFuncPtr))
	runFunc := *(*func() int)(unsafe.Pointer(&funcPtrContainer))
	result := runFunc()
	fmt.Println("result:{}", result)

	os.Stdout.Sync()
	codeModule.Unload()
}
