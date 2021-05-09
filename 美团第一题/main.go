package main

import "fmt"

type z struct {
	x, y, u, v, w int
}

func main() {
	var n int
	var min = 9999
	fmt.Scan(&n)
	var zList = make([]*z, 0)
	for i := 0;i < n; i++ {
		zTmp := &z{}
		fmt.Scan(&zTmp.x, &zTmp.y, &zTmp.u, &zTmp.v, &zTmp.w)
		zList = append(zList, zTmp)
		for i := 0; i < len(zList);i ++ {
			if zList[i].x == zList[i].y && zList[i].y ==1 {
				if zList.
			}
		}
	}
}
