package main

import "time"

type Point struct {
	X, Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 150}
}
