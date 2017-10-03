package main

import (
//"fmt"
)

func geo_fence(x float64, y float64, points [][]float64) bool {
	count := 0
	x1, y1 := points[0][0], points[0][1]
	x1_part := (y1 > y) || ((x1-x > 0) && (y1 == y))
	var a = []float64{x1, y1}
	p := append(points, a)
	for i := range p {
		if i == 0 {
			continue
		}
		point := p[i]
		x2, y2 := point[0], point[1]
		x2_part := (y2 > y) || ((x2 > x) && (y2 == y))
		if x2_part == x1_part {
			x1, y1 = x2, y2
			continue
		}
		mul := (x1-x)*(y2-y) - (x2-x)*(y1-y)
		if mul > 0 {
			count += 1
		} else {
			if mul < 0 {
				count -= 1
			}
		}
		x1, y1 = x2, y2
		x1_part = x2_part
	}

	if count == 2 || count == -2 {
		return true
	} else {
		return false
	}
}

func main() {
	points := [][]float64{{116.239117, 40.006047}, {116.250103, 39.835425}, {116.483563, 39.856513}, {116.494549, 39.995527}}
	x := 116.402538
	y := 39.922896
	for i := 1; i < 10000; i++ {
		_ = geo_fence(x+float64(i)*0.01, y+float64(i)*0.01, points)
	}
}
