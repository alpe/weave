package iavl

import (
	"github.com/tendermint/iavl"
	dbm "github.com/tendermint/tmlibs/db"

	"github.com/confio/weave/store"
)

// CommitStore manages a iavl committed state
type CommitStore struct {
	tree *iavl.VersionedTree
}

var _ store.CommitKVStore = CommitStore{}

// NewCommitStore creates a new store with disk backing
func NewCommitStore(path, name string) CommitStore {
	// Create the underlying leveldb datastore which will
	// persist the Merkle tree inner & leaf nodes.
	db, err := dbm.NewGoLevelDB(name, path)
	if err != nil {
		panic(err)
	}

	// numHistory := int64(100)
	// mainLoader := store.NewIAVLStoreLoader(db, cacheSize, numHistory)

	cacheSize := 10000
	tree := iavl.NewVersionedTree(cacheSize, db)
	tree.Load()
	return CommitStore{tree}
}

// MockCommitStore creates a new in-memory store for testing
func MockCommitStore() CommitStore {
	var db dbm.DB = dbm.NewMemDB()
	cacheSize := 10000
	tree := iavl.NewVersionedTree(cacheSize, db)
	return CommitStore{tree}
}

// Get returns the value at last committed state
// returns nil iff key doesn't exist. Panics on nil key.
func (s CommitStore) Get(key []byte) []byte {
	version := s.tree.LatestVersion()
	_, val := s.tree.GetVersioned(key, version)
	return val
}

// Commit the next version to disk, and returns info
func (s CommitStore) Commit() store.CommitID {
	version := s.tree.LatestVersion() + 1
	hash, err := s.tree.SaveVersion(version)
	if err != nil {
		panic(err)
	}
	return store.CommitID{
		Version: int64(version),
		Hash:    hash,
	}
}

// LoadLatestVersion loads the latest persisted version.
// If there was a crash during the last commit, it is guaranteed
// to return a stable state, even if older.
func (s CommitStore) LoadLatestVersion() error {
	return s.tree.Load()
}

// LatestVersion returns info on the latest version saved to disk
func (s CommitStore) LatestVersion() store.CommitID {
	return store.CommitID{
		Version: int64(s.tree.LatestVersion()),
		Hash:    s.tree.Hash(),
	}
}

// Adapter returns a wrapped version of the tree.
//
// Data writen here is stored in the tip of the version tree,
// and will be writen to disk on Commit. There is no way
// to rollback writes here, without throwing away the CommitStore
// and re-loading from disk.
func (s CommitStore) Adapter() store.CacheableKVStore {
	var kv store.KVStore = adapter{s.tree.Tree()}
	return store.BTreeCacheable{kv}
}

// CacheWrap wraps the Adapter with a cache, so it may be writen
// or discarded as needed.
func (s CommitStore) CacheWrap() store.KVCacheWrap {
	return s.Adapter().CacheWrap()
}

// func (b *Bonsai) GetVersionedWithProof(key []byte, version int64) ([]byte, iavl.KeyProof, error) {
//   return b.Tree.GetVersionedWithProof(key, uint64(version))
// }

// TODO: create batch and reader and wrap the rest in btree...

// adapter converts the working iavl.Tree to match these interfaces
type adapter struct {
	tree *iavl.Tree
}

var _ store.KVStore = adapter{}

// Get returns nil iff key doesn't exist. Panics on nil key.
func (a adapter) Get(key []byte) []byte {
	_, val := a.tree.Get(key)
	return val
}

// Has checks if a key exists. Panics on nil key.
func (a adapter) Has(key []byte) bool {
	return a.tree.Has(key)
}

// Set adds a new value
func (a adapter) Set(key, value []byte) {
	a.tree.Set(key, value)
}

// Delete removes from the tree
func (a adapter) Delete(key []byte) {
	a.tree.Remove(key)
}

// NewBatch returns a batch that can write multiple ops atomically
func (a adapter) NewBatch() store.Batch {
	return store.NewNonAtomicBatch(a)
}

// Iterator over a domain of keys in ascending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
func (a adapter) Iterator(start, end []byte) store.Iterator {
	var res []store.Model
	add := func(key []byte, value []byte) bool {
		m := store.Model{Key: key, Value: value}
		res = append(res, m)
		return false
	}
	a.tree.IterateRange(start, end, true, add)
	return store.NewSliceIterator(res)
}

// ReverseIterator over a domain of keys in descending order. End is exclusive.
// Start must be greater than end, or the Iterator is invalid.
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
func (a adapter) ReverseIterator(start, end []byte) store.Iterator {
	var res []store.Model
	add := func(key []byte, value []byte) bool {
		m := store.Model{Key: key, Value: value}
		res = append(res, m)
		return false
	}
	a.tree.IterateRange(start, end, false, add)
	return store.NewSliceIterator(res)
}
