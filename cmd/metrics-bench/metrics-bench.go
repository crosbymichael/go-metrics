package main

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"time"
)

func main() {
	r := metrics.NewRegistry()
	for i := 0; i < 10000; i++ {
//		r.RegisterCounter(fmt.Sprintf("%d", i), metrics.NewCounter())
//		r.RegisterGauge(fmt.Sprintf("%d", i), metrics.NewGauge())
//		r.RegisterHistogram(fmt.Sprintf("%d", i), metrics.NewHistogram(metrics.NewUniformSample(1028)))
//		r.RegisterHistogram(fmt.Sprintf("%d", i), metrics.NewHistogram(metrics.NewExpDecaySample(1028, 0.015)))
		r.RegisterMeter(fmt.Sprintf("%d", i), metrics.NewMeter())
	}
	time.Sleep(600e9)
}
