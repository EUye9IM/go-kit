package debounce_test

import (
	"testing"
	"time"

	"github.com/EUye9IM/go-kit/debounce"
)

const d = 10 * time.Millisecond

type Input []struct {
	data     []int
	duration time.Duration
}
type Output [][]int

func testChan(t *testing.T, src chan<- int, dst <-chan []int,
	input Input, output Output) {
	go func() {
		for _, i := range input {
			for _, v := range i.data {
				src <- v
				t.Log("send", v)
			}
			<-time.After(i.duration)
		}
		close(src)
	}()
	i := 0
	for d := range dst {
		if i >= len(output) {
			t.Fatal("unexpect data")
			break
		}
		t.Log("recv", d, "expect", output[i])
		if len(d) != len(output[i]) {
			t.Fatal("unexpect data len")
		}
		for i, v := range output[i] {
			if v != d[i] {
				t.Fatal("unexpect data value")
			}
		}
		i++
	}
	if i != len(output) {
		t.Fatal("unexpect close")
	}
}
func TestNewChan(t *testing.T) {
	srcCh := make(chan int)
	dstCh := debounce.NewChan(srcCh, 10*d)

	testInput := Input{
		{[]int{1, 2}, 8 * d},
		{[]int{3}, 8 * d},
		{[]int{4, 5}, 12 * d}, //send once
		{[]int{6}, 12 * d},    //send once
		{[]int{7, 8}, 0},      //send once , close instantly
	}
	testOutput := [][]int{
		{1, 2, 3, 4, 5},
		{6},
		{7, 8},
	}
	testChan(t, srcCh, dstCh, testInput, testOutput)
}
func TestNewChanClose(t *testing.T) {
	srcCh := make(chan int)
	dstCh := debounce.NewChan(srcCh, 10*d)

	testInput := Input{
		{[]int{1, 2}, 12 * d}, //send once
	}
	testOutput := [][]int{
		{1, 2},
	}
	testChan(t, srcCh, dstCh, testInput, testOutput)
}
