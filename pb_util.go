package wal

import (
	"fmt"
	types "github.com/ashwaniYDV/my-wal/types"
	"google.golang.org/protobuf/proto"
)

// MustMarshal marshals the wal entry to bytes
func MustMarshal(entry *types.WAL_Entry) []byte {
	marshaledEntry, err := proto.Marshal(entry)
	// this err means something is wrong in proto, so we should panic
	if err != nil {
		panic(fmt.Sprintf("Marshal should never fail (%v)", err))
	}

	return marshaledEntry
}

// MustUnmarshal unmarshals the bytes to wal entry
func MustUnmarshal(data []byte, entry *types.WAL_Entry) {
	// this err means something is wrong in proto, so we should panic
	if err := proto.Unmarshal(data, entry); err != nil {
		panic(fmt.Sprintf("Unmarshal should never fail (%v)", err))
	}
}
