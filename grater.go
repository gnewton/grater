package grater
// Copyright 2018 The Glen Newton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"log"
	"time"
	"errors"
)

type Grater struct {
     ticker *time.Ticker
     name string
}

// TODO
func NewFileGrater(d time.Duration, counter *uint64, filename string) *Grater{
     return nil
}    



func NewGrater(d time.Duration, counter *uint64) (*Grater, error){
     if counter == nil {
     		return nil, errors.New("Error: Counter is nil")
     }

     if int64(d) <= 0 {
     	return nil, errors.New("Error: Duration must be >0")
     }    
	var previousCounter uint64 = 0
	rater := new(Grater)
	
	rater.ticker = time.NewTicker(d)

	go func() {
		log.Println("Grater: START")	
		for _ = range rater.ticker.C {
			log.Println("Grater:", *counter-previousCounter)
			previousCounter = *counter
		}
		log.Println("Grater: END")		
	}()
	return rater, nil
}

func (r *Grater)Stop(){
     r.ticker.Stop()
     log.Println("Grater: STOP")	
}


