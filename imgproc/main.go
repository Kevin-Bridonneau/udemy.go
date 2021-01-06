package main

import (
	"flag"
	"fmt"
	"time"

	"udemy.go/imgproc/filter"
	"udemy.go/imgproc/task"
)

func main() {
	var srcDir = flag.String("src", "", "Input directory")
	var dstDir = flag.String("dst", "", "Output directory")
	var filterType = flag.String("filter", "grayscale", "grayscale/blur")
	var taskType = flag.String("type", "waitgrp", "waitgrp / chan")
	var poolSize = flag.Int("pool", 4, "PoolSize integer for chanel tasker")
	flag.Parse()
	var f filter.Filter

	switch *filterType {
	case "grayscale":
		f = filter.Grayscale{}
	case "blur":
		f = filter.Blur{}
	}

	var t task.Tasker
	switch *taskType {
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcDir, *dstDir, f)
	case "chan":
		t = task.NewChanTask(*srcDir, *dstDir, f, *poolSize)
	}

	start := time.Now()
	t.Process()
	elapsed := time.Since(start)
	fmt.Printf("Image processing took %v\n", elapsed)

}
