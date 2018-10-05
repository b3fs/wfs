package server

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/embed"
	"github.com/pingcap/kvproto/pkg/pdpb"
)

const (
	etcdTimeout           = time.Second * 3
	serverMetricsInterval = time.Minute
	// pdRootPath for all pd servers.
	pdRootPath      = "/pd"
	pdAPIPrefix     = "/pd/"
	pdClusterIDPath = "/pd/cluster_id"
)

// EnableZap enable the zap logger in embed etcd.
var EnableZap = false

// Server is the pd server.
type Server struct {
	// Server state.
	isServing int64
	leader    atomic.Value

	// Configs and initial fields.
	cfg         *Config
	etcdCfg     *embed.Config
	scheduleOpt *scheduleOption
	handler     *Handler

	serverLoopCtx    context.Context
	serverLoopCancel func()
	serverLoopWg     sync.WaitGroup

	// Etcd and cluster informations.
	etcd      *embed.Etcd
	client    *clientv3.Client
	id        uint64 // etcd server id.
	clusterID uint64 // pd cluster id.
	rootPath  string
	member    *pdpb.Member // current PD's info.
	// memberValue is the serialized string of `member`. It will be save in
	// etcd leader key when the PD node is successfully elected as the leader
	// of the cluster. Every write will use it to check leadership.
	memberValue string

	// Server services.
	// for id allocator, we can use one allocator for
	// store, region and peer, because we just need
	// a unique ID.
	idAlloc *idAllocator
	// for kv operation.
	kv *core.KV
	// for namespace.
	classifier namespace.Classifier
	// for raft cluster
	cluster *RaftCluster
	// For tso, set after pd becomes leader.
	ts            atomic.Value
	lastSavedTime time.Time
	// For async region heartbeat.
	hbStreams *heartbeatStreams
}
