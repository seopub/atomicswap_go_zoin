package memdb

import (
	"testing"

	"seopub/btcsuite_goleveldb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
