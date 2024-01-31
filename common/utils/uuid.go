package utils

import (
	"encoding/binary"
	"sync"
	"time"
)

var (
	mux          sync.Mutex
	prvTimestamp int64
	instanceId   int
	usedNumbers  = make(map[string]bool)
)

func GenerateUniqueUUid() int64 {
	var ret int64 = 0

	binsID := make([]byte, 8)
	baseB := make([]byte, 8)
	instanceB := make([]byte, 4)

	var instanceMod = instanceId % 256 // max 256 instance

	mux.Lock()
	defer mux.Unlock()

	t := time.Now().UnixMilli()
	if t <= prvTimestamp {
		ret = prvTimestamp + 1
	} else {
		ret = t
	}
	prvTimestamp = ret

	binary.BigEndian.PutUint64(baseB, uint64(ret))
	binary.BigEndian.PutUint32(instanceB, uint32(instanceMod))

	// set first 6byte
	binsID[1] = baseB[2]
	binsID[2] = baseB[3]
	binsID[3] = baseB[4]
	binsID[4] = baseB[5]
	binsID[5] = baseB[6]
	binsID[6] = baseB[7]

	// next 1 byte for instance id
	binsID[7] = instanceB[3]

	ret = int64(binary.BigEndian.Uint64(binsID))

	return ret
}

func GetCurrentTimestamp() int64 {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	return timestamp
}

func CodeOPT() int {
	var ret int = 0

	binsID := make([]byte, 8)
	baseB := make([]byte, 8)
	instanceB := make([]byte, 4)

	var instanceMod = instanceId % 256 // max 256 instance

	mux.Lock()
	defer mux.Unlock()

	t := time.Now().UnixMilli()
	if t <= prvTimestamp {
		ret = int(prvTimestamp + 1)
	} else {
		ret = int(t)
	}
	prvTimestamp = int64(ret)

	binary.BigEndian.PutUint64(baseB, uint64(ret))
	binary.BigEndian.PutUint32(instanceB, uint32(instanceMod))

	// set first 6byte
	binsID[1] = baseB[2]
	binsID[2] = baseB[3]
	binsID[3] = baseB[4]
	binsID[4] = baseB[5]
	binsID[5] = baseB[6]
	binsID[6] = baseB[7]

	// next 1 byte for instance id
	binsID[7] = instanceB[3]

	ret = int(binary.BigEndian.Uint64(binsID))

	// giữ 6 chữ số cuối của ret
	ret = ret % 1000000

	return ret
}
