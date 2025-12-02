package passgen

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func BenchmarkGenerate_Length8(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 8, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Length16(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 16, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Length32(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 32, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Length64(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 64, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Length128(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 128, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Length256(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 256, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Length512(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 512, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Length1024(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 1024, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Length2048(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 2048, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Length4096(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 4096, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_LevelLow(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 32, Level: LevelLow}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_LevelMedium(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 32, Level: LevelMedium}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_LevelStrong(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 32, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_ShortInput(b *testing.B) {
	cfg := Config{Input: "ab", Salt: "salt", Length: 32, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_MediumInput(b *testing.B) {
	cfg := Config{Input: strings.Repeat("a", 100), Salt: "salt", Length: 32, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_LongInput(b *testing.B) {
	cfg := Config{Input: strings.Repeat("a", 500), Salt: "salt", Length: 32, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_MaxInput(b *testing.B) {
	cfg := Config{Input: strings.Repeat("a", 1000), Salt: "salt", Length: 32, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_NoSalt(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "", Length: 32, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_ShortSalt(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "s", Length: 32, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_LongSalt(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: strings.Repeat("salt", 100), Length: 32, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkDetermRNG_NextByte_Sequential(b *testing.B) {
	rng := newDetermRNG("benchseed")
	b.ResetTimer()
	for b.Loop() {
		rng.nextByte()
	}
}

func BenchmarkDetermRNG_NextByte_NewInstance(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		rng := newDetermRNG("benchseed")
		rng.nextByte()
	}
}

func BenchmarkDetermRNG_Intn_Max10(b *testing.B) {
	rng := newDetermRNG("benchseed")
	b.ResetTimer()
	for b.Loop() {
		rng.Intn(10)
	}
}

func BenchmarkDetermRNG_Intn_Max26(b *testing.B) {
	rng := newDetermRNG("benchseed")
	b.ResetTimer()
	for b.Loop() {
		rng.Intn(26)
	}
}

func BenchmarkDetermRNG_Intn_Max62(b *testing.B) {
	rng := newDetermRNG("benchseed")
	b.ResetTimer()
	for b.Loop() {
		rng.Intn(62)
	}
}

func BenchmarkDetermRNG_Intn_Max83(b *testing.B) {
	rng := newDetermRNG("benchseed")
	b.ResetTimer()
	for b.Loop() {
		rng.Intn(83)
	}
}

func BenchmarkDetermRNG_Intn_Max100(b *testing.B) {
	rng := newDetermRNG("benchseed")
	b.ResetTimer()
	for b.Loop() {
		rng.Intn(100)
	}
}

func BenchmarkDetermRNG_Intn_Max200(b *testing.B) {
	rng := newDetermRNG("benchseed")
	b.ResetTimer()
	for b.Loop() {
		rng.Intn(200)
	}
}

func BenchmarkDetermRNG_Intn_Max255(b *testing.B) {
	rng := newDetermRNG("benchseed")
	b.ResetTimer()
	for b.Loop() {
		rng.Intn(255)
	}
}

func BenchmarkDetermRNG_Intn_Max256(b *testing.B) {
	rng := newDetermRNG("benchseed")
	b.ResetTimer()
	for b.Loop() {
		rng.Intn(256)
	}
}

func BenchmarkDetermRNG_Refill_Once(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		rng := newDetermRNG("benchseed")
		rng.refill()
	}
}

func BenchmarkDetermRNG_Refill_Multiple(b *testing.B) {
	rng := newDetermRNG("benchseed")
	b.ResetTimer()
	for b.Loop() {
		rng.refill()
	}
}

func BenchmarkDetermRNG_ShortSeed(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		rng := newDetermRNG("s")
		rng.Intn(100)
	}
}

func BenchmarkDetermRNG_MediumSeed(b *testing.B) {
	seed := strings.Repeat("a", 64)
	b.ResetTimer()
	for b.Loop() {
		rng := newDetermRNG(seed)
		rng.Intn(100)
	}
}

func BenchmarkDetermRNG_LongSeed(b *testing.B) {
	seed := strings.Repeat("a", 1024)
	b.ResetTimer()
	for b.Loop() {
		rng := newDetermRNG(seed)
		rng.Intn(100)
	}
}

func BenchmarkDetermRNG_VeryLongSeed(b *testing.B) {
	seed := strings.Repeat("a", 4096)
	b.ResetTimer()
	for b.Loop() {
		rng := newDetermRNG(seed)
		rng.Intn(100)
	}
}

func BenchmarkGenerateRandomString_Length8(b *testing.B) {
	for b.Loop() {
		GenerateRandomString(8)
	}
}

func BenchmarkGenerateRandomString_Length16(b *testing.B) {
	for b.Loop() {
		GenerateRandomString(16)
	}
}

func BenchmarkGenerateRandomString_Length32(b *testing.B) {
	for b.Loop() {
		GenerateRandomString(32)
	}
}

func BenchmarkGenerateRandomString_Length64(b *testing.B) {
	for b.Loop() {
		GenerateRandomString(64)
	}
}

func BenchmarkGenerateRandomString_Length128(b *testing.B) {
	for b.Loop() {
		GenerateRandomString(128)
	}
}

func BenchmarkGenerateRandomString_Length256(b *testing.B) {
	for b.Loop() {
		GenerateRandomString(256)
	}
}

func BenchmarkGenerateRandomString_Length512(b *testing.B) {
	for b.Loop() {
		GenerateRandomString(512)
	}
}

func BenchmarkGenerateRandomString_Length1024(b *testing.B) {
	for b.Loop() {
		GenerateRandomString(1024)
	}
}

func BenchmarkGenerateRandomString_Length4096(b *testing.B) {
	for b.Loop() {
		GenerateRandomString(4096)
	}
}

func BenchmarkGenerate_Parallel(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 32, Level: LevelStrong}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Generate(cfg)
		}
	})
}

func BenchmarkGenerate_Parallel_Length128(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 128, Level: LevelStrong}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Generate(cfg)
		}
	})
}

func BenchmarkGenerate_Parallel_Length1024(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 1024, Level: LevelStrong}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Generate(cfg)
		}
	})
}

func BenchmarkGenerateRandomString_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GenerateRandomString(32)
		}
	})
}

func BenchmarkGenerateRandomString_Parallel_Length128(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GenerateRandomString(128)
		}
	})
}

func BenchmarkGenerateRandomString_Parallel_Length1024(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GenerateRandomString(1024)
		}
	})
}

func BenchmarkGenerate_Allocs_Short(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 16, Level: LevelStrong}
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Allocs_Medium(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 128, Level: LevelStrong}
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Allocs_Long(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 1024, Level: LevelStrong}
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerateRandomString_Allocs_Short(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		GenerateRandomString(16)
	}
}

func BenchmarkGenerateRandomString_Allocs_Medium(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		GenerateRandomString(128)
	}
}

func BenchmarkGenerateRandomString_Allocs_Long(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		GenerateRandomString(1024)
	}
}

func BenchmarkDetermRNG_Allocs(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		rng := newDetermRNG("benchseed")
		for range 100 {
			rng.Intn(100)
		}
	}
}

func BenchmarkGenerate_Throughput_16B(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 16, Level: LevelStrong}
	b.SetBytes(16)
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Throughput_64B(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 64, Level: LevelStrong}
	b.SetBytes(64)
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Throughput_256B(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 256, Level: LevelStrong}
	b.SetBytes(256)
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Throughput_1KB(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 1024, Level: LevelStrong}
	b.SetBytes(1024)
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Throughput_4KB(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 4096, Level: LevelStrong}
	b.SetBytes(4096)
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_Batch100(b *testing.B) {
	cfgs := make([]Config, 100)
	for i := range 100 {
		cfgs[i] = Config{
			Input:  fmt.Sprintf("user%d@example.com", i),
			Salt:   "global_salt",
			Length: 16,
			Level:  LevelStrong,
		}
	}
	b.ResetTimer()
	for b.Loop() {
		for _, cfg := range cfgs {
			Generate(cfg)
		}
	}
}

func BenchmarkGenerate_DifferentUsers(b *testing.B) {
	b.ResetTimer()
	i := 0
	for b.Loop() {
		i += 1
		cfg := Config{
			Input:  fmt.Sprintf("user%d@example.com", i%1000),
			Salt:   "salt",
			Length: 16,
			Level:  LevelStrong,
		}
		Generate(cfg)
	}
}

func BenchmarkGenerate_DifferentSites(b *testing.B) {
	sites := []string{"google.com", "github.com", "twitter.com", "facebook.com", "amazon.com"}
	b.ResetTimer()
	i := 0
	for b.Loop() {
		i += 1
		cfg := Config{
			Input:  "user@example.com+" + sites[i%len(sites)],
			Salt:   "master_salt",
			Length: 20,
			Level:  LevelStrong,
		}
		Generate(cfg)
	}
}

func BenchmarkGenerate_MinLength(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 1, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_MaxLength(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 4096, Level: LevelStrong}
	b.ResetTimer()
	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkGenerate_ScalabilityCPU(b *testing.B) {
	cpuCounts := []int{1, 2, 4, runtime.NumCPU()}
	for _, numCPU := range cpuCounts {
		if numCPU > runtime.NumCPU() {
			continue
		}
		b.Run(fmt.Sprintf("CPU%d", numCPU), func(b *testing.B) {
			runtime.GOMAXPROCS(numCPU)
			cfg := Config{Input: "benchmark", Salt: "salt", Length: 32, Level: LevelStrong}
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					Generate(cfg)
				}
			})
		})
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func BenchmarkGenerate_Warmup(b *testing.B) {
	cfg := Config{Input: "benchmark", Salt: "salt", Length: 32, Level: LevelStrong}
	for range 1000 {
		Generate(cfg)
	}

	for b.Loop() {
		Generate(cfg)
	}
}

func BenchmarkDetermRNG_Warmup(b *testing.B) {
	rng := newDetermRNG("warmup")
	for range 10000 {
		rng.Intn(100)
	}
	b.ResetTimer()
	for b.Loop() {
		rng.Intn(100)
	}
}
