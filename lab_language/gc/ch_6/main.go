// https://mp.weixin.qq.com/s/jGGCccMOx4s5asG2IXWNMQ
package main

import (
	"fmt"
	"runtime"
	"time"
)

func MapSlice() {
	a := make(map[byte][]byte, 1e8)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}

	runtime.KeepAlive(a)
}

// Map 值类型的Key/Value不超过128字节 GC将会忽略
func MapArray() {
	a := make(map[byte][128]byte, 1e8)
	//a := make(map[byte][16]int, 1e8)
	//a := make(map[byte][64]int16, 1e8)
	//a := make(map[byte][32]int32, 1e8)
	//a := make(map[byte][16]int64, 1e8)
	//a := make(map[byte][32]float32, 1e8)
	//a := make(map[byte][16]float64, 1e8)
	//a := make(map[byte][128]bool, 1e8)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}

	runtime.KeepAlive(a)
}

// Map 值类型的Value超过128字节会变成pointer类型
func MapArrayExceed128() {
	a := make(map[byte][129]byte, 1e8)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}

	runtime.KeepAlive(a)
}

func main() {
	//MapSlice()
	MapArray()
	//MapArrayExceed128()
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
