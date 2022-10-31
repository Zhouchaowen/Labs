package main

import (
	"fmt"
	"runtime"
	"time"
)

func MapKeyString() {
	a := make(map[string]int, 1e8)

	for i := 0; i < 2; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("MapKeyString GC took %s \n", time.Since(start))
	}

	runtime.KeepAlive(a)
}

// Map 值类型的Key/Value不超过128字节 GC将会忽略
func MapKeyArray() {
	a := make(map[[32]byte]int, 1e8)

	for i := 0; i < 2; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("MapKeyArray GC took %s \n", time.Since(start))
	}

	runtime.KeepAlive(a)
}

func main() {
	MapKeyString()
	MapKeyArray()
}

/*
func MapBucketType(t *types.Type) *types.Type {
	if t.MapType().Bucket != nil {
		return t.MapType().Bucket
	}

	keytype := t.Key()
	elemtype := t.Elem()
	types.CalcSize(keytype)
	types.CalcSize(elemtype)
	if keytype.Width > MAXKEYSIZE { 			// MAXKEYSIZE = 128
		keytype = types.NewPtr(keytype)			// 变为指针
	}
	if elemtype.Width > MAXELEMSIZE {
		elemtype = types.NewPtr(elemtype)
	}
	.......
}
*/
