package archive

// EvictionPolicy represents the eviction policy for removing KVPs from the map
type EvictionPolicy interface {
	// Evict and return the next key to be removed according to the policy (_, false if no keys are evicted)
	Evict() (string, bool)
	// Admit a key to the policy manager
	Admit(string)
	// Disregard (ignore) a key from the policy manager
	Disregard(string)
}
