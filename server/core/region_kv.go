package core

import (
	"context"
	"sync"
	"time"

	"github.com/pingcap/kvproto/pkg/metapb"
)

var dirtyFlushTick = time.Second

// RegionKV is used to save regions.
type RegionKV struct {
	*leveldbKV
	mu           sync.RWMutex
	batchRegions map[string]*metapb.Region
	batchSize    int
	cacheSize    int
	flushRate    time.Duration
	flushTime    time.Time
	ctx          context.Context
	cancel       context.CancelFunc
}
