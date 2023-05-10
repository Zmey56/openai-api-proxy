package proxy

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"github.com/Zmey56/openai-api-proxy/log"
	"hash/fnv"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var (
	hostAndProcessSalt     uint32
	hostAndProcessSaltOnce sync.Once
)

func HostAndProcessSalt() uint32 {
	hostAndProcessSaltOnce.Do(func() {
		// Generate process prefix (12 bits)
		hash := fnv.New32()

		hostname, err := os.Hostname()
		if err == nil {
			_, err = hash.Write([]byte(hostname))
			if err != nil {
				panic(err)
			}
		} else {
			log.Warning.Print("Failed to get hostname.")
			log.Warning.Printf("os.Hostname failed. %v", hostname)
		}

		pid := os.Getpid()
		_, err = hash.Write([]byte{
			byte(pid >> 24),
			byte(pid >> 16),
			byte(pid >> 8),
			byte(pid),
		})
		if err != nil {
			log.Debug.Printf("hash.Write failed. %v", err)
		}

		hostAndProcessSalt = hash.Sum32()
	})
	return hostAndProcessSalt
}

func cryptoRandomBytes(l int) []byte {
	bytes := make([]byte, l)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Warning.Printf("Failed to read 4 random bytes from crypto/rand.")
		log.Debug.Printf("rand.Read failed. %v", err)
	}
	return bytes
}

func cryptoRandomInt64(l int) int64 {
	if l <= 0 || l > 8 {
		panic("l should be between 1 and 8")
	}
	bytes := cryptoRandomBytes(l)
	result := int64(bytes[0])
	for i := 1; i < l; i++ {
		shift := uint(i * 8)
		result = result | (int64(bytes[i]) << shift)
	}
	return result
}

var (
	base32Encoding     = base32.HexEncoding.WithPadding(base32.NoPadding)
	base32Encoding8Len = base32Encoding.EncodedLen(8)
)

func newID64Generator() *id64Generator {
	// Generate process prefix (12 bits)
	hostAndProcessSalt := int64(HostAndProcessSalt() << 24)

	// Generate sequence number (will use only 20 bits)
	sequenceID := cryptoRandomInt64(3)

	return &id64Generator{
		hostAndProcessSalt: hostAndProcessSalt,
		sequenceID:         sequenceID,
	}
}

type id64Generator struct {
	hostAndProcessSalt int64
	sequenceID         int64
	epochTime          int64
}

type ID64 int64

func (id ID64) HostAndProcessSalt() int8 {
	return int8((id & 0x0000000FFF000000) >> 24)
}

func (id ID64) Time() time.Time {
	return time.Unix(int64(id)>>32, 0)
}

func (id ID64) SequenceNum() int {
	return int(id & 0x0000000000FFFFFF)
}

func (id ID64) String() string {
	bytes := []byte{
		byte(id >> 56),
		byte(id >> 48),
		byte(id >> 40),
		byte(id >> 32),
		byte(id >> 24),
		byte(id >> 16),
		byte(id >> 8),
		byte(id),
	}
	return base32Encoding.EncodeToString(bytes)
}

func (generator *id64Generator) New() ID64 {
	previousEpochTime := atomic.LoadInt64(&generator.epochTime)
	epochTime := time.Now().Unix()

	sequenceID := int64(0)
	if previousEpochTime != epochTime {
		// Just to be sure that it will be impossible to guess how many
		// New ids is getting generated in a second
		sequenceID = cryptoRandomInt64(3)
		atomic.StoreInt64(&generator.epochTime, epochTime)
		atomic.StoreInt64(&generator.sequenceID, sequenceID)
	} else {
		sequenceID = atomic.AddInt64(&generator.sequenceID, 1)
	}

	return ID64(
		epochTime<<32 |
			generator.hostAndProcessSalt |
			(0x0000000000FFFFFF & sequenceID))
}

func Int64ToID64(val int64) ID64 {
	return ID64(val)
}

func StringToID64(val string) (ID64, error) {
	if len(val) > base32Encoding8Len {
		return ID64(0), errors.New("invalid ID64")
	}
	// We trim =, so we need to add it back
	bytes, err := base32Encoding.DecodeString(val)
	if err != nil || len(bytes) != 8 {
		if err != nil {
			log.Debug.Printf("base32.HexEncoding.DecodeString failed. %v", err)
		}
		return ID64(0), errors.New("invalid ID64")
	}
	return ID64(
		int64(bytes[0])<<56 |
			int64(bytes[1])<<48 |
			int64(bytes[2])<<40 |
			int64(bytes[3])<<32 |
			int64(bytes[4])<<24 |
			int64(bytes[5])<<16 |
			int64(bytes[6])<<8 |
			int64(bytes[7]),
	), nil
}
