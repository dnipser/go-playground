package main

import (
	"fmt"
	"math"
	"reflect"
)

type Point interface {
	getX() int
	getY() int
	distance() float64
	toString() string
}

type DecartPoint struct {
	x, y int
}

func (dp *DecartPoint) getX() int {
	return dp.x
}

func (dp *DecartPoint) getY() int {
	return dp.y
}

func (dp *DecartPoint) distance() float64 {
	res := math.Sqrt(float64(dp.y*dp.y + dp.x*dp.x))
	return res
}

func (dp *DecartPoint) toString() string {
	return fmt.Sprint(dp.getX(), dp.getY(), dp.distance())
}

func createPoint(x, y int) Point {
	return &DecartPoint{x, y}
}

func printPoint(p Point) {
	fmt.Println(p.toString())
}

func printObject(obj interface{}) {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)
	fmt.Println(objType, objValue)
}

func runInterfaceExample() {
	point := createPoint(4, 6)
	printPoint(point)
	printObject(point)
}
