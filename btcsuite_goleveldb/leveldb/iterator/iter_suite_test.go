package iterator_test

import (
	"testing"

	"seopub/btcsuite_goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
