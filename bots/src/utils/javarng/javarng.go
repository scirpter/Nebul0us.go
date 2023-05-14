package javarng

type JavaRandom struct {
	seed int64
}

func NewJavaRandom(seed int64) *JavaRandom {
	return WithSeed(seed)
}

func WithSeed(seed int64) *JavaRandom {
	seed = (seed ^ 0x5DEECE66D) & ((1 << 48) - 1)
	return &JavaRandom{seed: seed}
}

func (r *JavaRandom) SetSeed(seed int64) {
	seed = (seed ^ 0x5DEECE66D) & ((1 << 48) - 1)
	r.seed = seed
}

func (r *JavaRandom) Seed() int64 {
	return r.seed
}

func (r *JavaRandom) Next(bits int32) int32 {
	r.seed = (r.seed*0x5DEECE66D + 0xB) & ((1 << 48) - 1)
	return int32(uint64(r.seed) >> (48 - bits))
}

func (r *JavaRandom) NextBytes(bytes []byte) {
	i := 0
	for i < len(bytes) {
		rnd := r.NextInt32(0)
		n := min(len(bytes)-i, 4)
		for n > 0 {
			bytes[i] = byte(rnd)
			i += 1
			n -= 1
			rnd >>= 8
		}
	}
}

func (r *JavaRandom) NextInt32(bound int32) int32 {
	if bound <= 0 {
		panic("bound must be positive")
	}
	if (bound & -bound) == bound {
		return int32(int64(bound) * int64(r.Next(31)) >> 31)
	}
	var bits, val int32
	for {
		bits = r.Next(31)
		val = bits % bound
		if bits-val+bound-1 >= 0 {
			break
		}
	}
	return val
}

func (r *JavaRandom) NextInt64() int64 {
	return (int64(r.Next(32)) << 32) + int64(r.Next(32))
}

func (r *JavaRandom) NextBoolean() bool {
	return r.Next(1) != 0
}

func (r *JavaRandom) NextFloat() float32 {
	return float32(r.Next(24)) / (1 << 24)
}

func (r *JavaRandom) NextDouble() float64 {
	return float64((int64(r.Next(26))<<27)+int64(r.Next(27))) / (1 << 53)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
