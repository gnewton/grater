package grater

// Copyright 2018 The Glen Newton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"bufio"
	"errors"
	"log"
	"strconv"
	"time"
)

type Grater struct {
	ticker *time.Ticker
	name   string
	writer *bufio.Writer
	outFunc      outFunc
}


func NewFileGrater(d time.Duration, counter *uint64, w *bufio.Writer) (*Grater, error) {
	return graterInit(d, counter, writerOut, w)
}

func NewGrater(d time.Duration, counter *uint64) (*Grater, error) {
	return graterInit(d, counter, logOut, nil)
}

func (r *Grater) Stop() {
	r.ticker.Stop()
}

func graterInit(d time.Duration, counter *uint64, o outFunc, w *bufio.Writer) (*Grater, error) {
	if counter == nil {
		return nil, errors.New("Error: Counter is nil")
	}

	if int64(d) <= 0 {
		return nil, errors.New("Error: Duration must be >0")
	}

	if o == nil{
		return nil, errors.New("outFunc is nil")
	}
	
	grater := new(Grater)
	grater.outFunc = o
	grater.writer = w
	grater.ticker = time.NewTicker(d)

	go counterTicker(grater, counter)

	return grater, nil
}


func counterTicker(r *Grater, counter *uint64) {
	if r == nil || counter == nil {
		log.Fatal("Grater or counter are nil")
	}

	var previousCounter uint64 = 0

	for _ = range r.ticker.C {
		err := r.outFunc(r, strconv.FormatUint(*counter-previousCounter, 10))
		if err != nil{
			log.Fatal("Grater or counter are nil")
		}
		previousCounter = *counter
	}

}

