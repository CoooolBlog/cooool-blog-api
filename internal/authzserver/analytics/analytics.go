/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package analytics

import (
	"sync"
	"sync/atomic"
	"time"

	log "github.com/kristianhuang/go-cmp/rollinglog"
	"github.com/kristianhuang/go-cmp/storage"
	"gopkg.in/vmihailenco/msgpack.v2"
)

const (
	analyticsKeyName                 = "cooool-blog-system-analytics"
	recordsBufferForcedFlushInterval = 1 * time.Second
)

type AnalyticsRecord struct {
	TimeStamp  int64     `json:"timestamp"`
	Username   string    `json:"username"`
	Effect     string    `json:"effect"`
	Conclusion string    `json:"conclusion"`
	Request    string    `json:"request"`
	Policies   string    `json:"policies"`
	Deciders   string    `json:"deciders"`
	ExpireAt   time.Time `json:"expire-at" bson:"expire-at"`
}

var analytics *Analytics

// SetExpiry set expiration time to a key.
func (a *AnalyticsRecord) SetExpiry(expiresInSeconds int64) {
	expiry := time.Duration(expiresInSeconds) * time.Second
	if expiresInSeconds == 0 {
		expiry = 24 * 365 * 100 * time.Hour
	}

	t := time.Now()
	t2 := t.Add(expiry)
	a.ExpireAt = t2
}

// Analytics will record analytics data to a redis back end as defined in the Config object.
type Analytics struct {
	store                      storage.AnalyticsHandler
	poolSize                   int
	recordsChan                chan *AnalyticsRecord
	workerBufferSize           uint64
	recordsBufferFlushInterval uint64
	shouldStop                 uint32
	poolWg                     sync.WaitGroup
}

func NewAnalytics(options *AnalyticsOptions, store storage.AnalyticsHandler) *Analytics {
	ps := options.PoolSize
	recordsBufferSize := options.RecordsBufferSize
	workerBufferSize := recordsBufferSize / uint64(ps)
	log.Debug("Analytics pool worker buffer size", log.Uint64("workerBufferSize", workerBufferSize))

	recordsChan := make(chan *AnalyticsRecord, recordsBufferSize)

	analytics = &Analytics{
		store:                      store,
		poolSize:                   ps,
		recordsChan:                recordsChan,
		workerBufferSize:           workerBufferSize,
		recordsBufferFlushInterval: options.FlushInterval,
	}

	return analytics
}

func GetAnalytics() *Analytics {
	return analytics
}

func (a *Analytics) Start() {
	a.store.Connect()
	atomic.SwapUint32(&a.shouldStop, 0)
	for i := 0; i < a.poolSize; i++ {
		a.poolWg.Add(1)
		go a.recordWorker()
	}
}

// Stop stop the analytics service.
func (a *Analytics) Stop() {
	// flag to stop sending records into channel
	atomic.SwapUint32(&a.shouldStop, 1)
	// close channel to stop workers
	close(a.recordsChan)
	// wait for all workers to be done
	a.poolWg.Wait()
}

func (a *Analytics) RecordHit(record *AnalyticsRecord) error {
	// check if we should stop sending records 1st
	if atomic.LoadUint32(&a.shouldStop) > 0 {
		return nil
	}

	// just send record to channel consumed by pool of workers
	// leave all data crunching and Redis I/O work for pool workers.
	a.recordsChan <- record

	return nil
}

func (a *Analytics) recordWorker() {
	defer a.poolWg.Done()

	// this is buffer to send one pipelined command to redis
	// use r.recordsBufferSize as cap to reduce slice re-allocations
	recordsBuffer := make([][]byte, 0, a.workerBufferSize)

	// read records from channel and process
	lastSentTS := time.Now()
	for {
		var readyToSend bool
		select {
		case record, ok := <-a.recordsChan:
			// check if channel was closed and it is time to exit from worker
			if !ok {
				// send what is left in buffer
				a.store.AppendToSetPipelined(analyticsKeyName, recordsBuffer)

				return
			}

			// we have new record - prepare it and add to buffer

			if encoded, err := msgpack.Marshal(record); err != nil {
				log.Errorf("Error encoding analytics data: %s", err.Error())
			} else {
				recordsBuffer = append(recordsBuffer, encoded)
			}

			// identify that buffer is ready to be sent
			readyToSend = uint64(len(recordsBuffer)) == a.workerBufferSize

		case <-time.After(time.Duration(a.recordsBufferFlushInterval) * time.Millisecond):
			// nothing was received for that period of time
			// anyways send whatever we have, don't hold data too long in buffer
			readyToSend = true
		}

		// send data to Redis and reset buffer
		if len(recordsBuffer) > 0 && (readyToSend || time.Since(lastSentTS) >= recordsBufferForcedFlushInterval) {
			a.store.AppendToSetPipelined(analyticsKeyName, recordsBuffer)
			recordsBuffer = recordsBuffer[:0]
			lastSentTS = time.Now()
		}
	}
}

// DurationToMillisecond convert time duration type to float64.
func DurationToMillisecond(d time.Duration) float64 {
	return float64(d) / 1e6
}
