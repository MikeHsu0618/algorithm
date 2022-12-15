package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice; expect 5 but get %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var test = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for i := 0; i < len(test); i++ {
		orderFinished = []string{}
		eatTime = test[i].delay
		sleepTime = test[i].delay
		thinkTime = test[i].delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice; expect 5 but get %d", len(orderFinished))
		}
	}
}
