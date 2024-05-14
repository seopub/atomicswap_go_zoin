package table

import (
	"testing"

	"seopub/btcsuite_goleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
