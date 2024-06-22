/*
debounce 用于消抖

例如fsnotify会短时间内产生多个event，通过 NewChan 产生一个新的 Chan，
该 chan 能够过滤短时间内相同的消息并适时返回一个消息切片
*/
package debounce

import "time"

// NewChan 新建消抖后Channel。
// srcChan <-chan T 消息来源chan
// duration time.Duration 消抖时长
func NewChan[T any](srcChan <-chan T, duration time.Duration) <-chan []T {
	dstChan := make(chan []T, cap(srcChan))
	timer := time.NewTimer(duration)
	if !timer.Stop() {
		<-timer.C
	}
	go func() {
		defer close(dstChan)
		for data := range srcChan {
			dataLst := []T{data}
			timer.Reset(duration)

			shouldSend := false
			for !shouldSend {
				select {
				case <-timer.C:
					timer.Stop()
					shouldSend = true
				case data, ok := <-srcChan:
					if ok {
						dataLst = append(dataLst, data)
						timer.Reset(duration)
					} else {
						timer.Stop()
						dstChan <- dataLst
						return
					}
				}
			}
			dstChan <- dataLst
		}
	}()
	return dstChan
}
