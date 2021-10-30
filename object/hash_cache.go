package object

type HashCache struct {
	Value map[Hashable]HashKey
}

func NewHashCache() *HashCache {
	ch := &HashCache{Value: make(map[Hashable]HashKey)}

	return ch
}

func (hc *HashCache) Get(key Hashable, hashFn func() HashKey) HashKey {
	if cached, ok := hc.Value[key]; ok {
		return cached
	}

	h := hashFn()
	hc.Value[key] = h

	return h
}
