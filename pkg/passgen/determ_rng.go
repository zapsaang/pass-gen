package passgen

import (
	"crypto/sha256"
	"encoding/binary"
)

type determRNG struct {
	seed    []byte
	counter uint64
	buffer  []byte  
	ptr     int   
}

func newDetermRNG(seedStr string) *determRNG {
	return &determRNG{
		seed:    []byte(seedStr),
		counter: 0,
		buffer:  nil, 
		ptr:     0,
	}
}

func (r *determRNG) nextByte() byte {
	if r.buffer == nil || r.ptr >= len(r.buffer) {
		r.refill()
	}
	b := r.buffer[r.ptr]
	r.ptr++
	return b
}

func (r *determRNG) refill() {
	h := sha256.New()
	h.Write(r.seed)

	ctrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(ctrBytes, r.counter)
	h.Write(ctrBytes)

	r.buffer = h.Sum(nil)
	r.ptr = 0
	r.counter++
}

func (r *determRNG) Intn(max int) int {
	if max <= 0 {
		return 0
	}

	if max <= 256 {
		limit := 256 - (256 % max)
		for {
			b := int(r.nextByte())
			if b < limit {
				return b % max
			}
		}
	}

	limit := 65536 - (65536 % max)
	for {
		b1 := r.nextByte()
		b2 := r.nextByte()
		val := int(b1)<<8 | int(b2)

		if val < limit {
			return val % max
		}
	}
}
