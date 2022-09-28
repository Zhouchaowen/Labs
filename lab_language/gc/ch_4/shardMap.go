package ch_4

import "sync"

type MapShards []MapShard

const ShardCount = 1024

type MapShard struct {
	items map[string]string
	mu    sync.RWMutex
}

func NewMapShards() MapShards {
	m := make(MapShards, ShardCount)
	for i := 0; i < ShardCount; i++ {
		m[i] = MapShard{
			items: map[string]string{},
		}
	}
	return m
}

func (m MapShards) GetShard(key string) *MapShard {
	return &m[uint(Sum64(key))%uint(ShardCount)]
}

func (m MapShards) Get(key string) (string, bool) {
	shard := m.GetShard(key)
	shard.mu.RLock()
	defer shard.mu.RUnlock()
	value, ok := shard.items[key]
	return value, ok
}

func (m MapShards) Set(key string, value string) {
	shard := m.GetShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	shard.items[key] = value
}

const (
	// offset64 FNVa offset basis. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	offset64 = 14695981039346656037
	// prime64 FNVa prime value. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	prime64 = 1099511628211
)

// Sum64 gets the string and returns its uint64 hash value.
func Sum64(key string) uint64 {
	var hash uint64 = offset64
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= prime64
	}

	return hash
}
