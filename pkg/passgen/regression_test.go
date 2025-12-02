package passgen

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const (
	regressionDir  = "testdata"
	regressionFile = "regression_1m.jsonl.gz"
	dataCount      = 1_000_000
)

var levels = []Level{LevelLow, LevelMedium, LevelStrong}

type RegressionCase struct {
	Cfg      Config `json:"c"`
	Expected string `json:"e"`
}

func TestConsistency_Massive(t *testing.T) {
	filePath := filepath.Join(regressionDir, regressionFile)
	if err := ensureRegressionData(t, filePath); err != nil {
		t.Fatalf("Failed to generate test data: %v", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open regression file: %v", err)
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		t.Fatalf("Failed to create gzip reader: %v", err)
	}
	defer gzr.Close()

	t.Logf("Starting verification of %d records from %s...", dataCount, filePath)

	scanner := bufio.NewScanner(gzr)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	count := 0
	failed := 0
	start := time.Now()

	for scanner.Scan() {
		count++
		line := scanner.Bytes()

		var tc RegressionCase
		if err := json.Unmarshal(line, &tc); err != nil {
			t.Errorf("Line %d: JSON decode error: %v", count, err)
			continue
		}

		got, err := Generate(tc.Cfg)
		if err != nil {
			t.Errorf("Line %d: Generate returned error: %v", count, err)
			failed++
			continue
		}

		if got != tc.Expected {
			t.Errorf("\nMISMATCH at Line %d:\nConfig: %+v\nExpected: %s\nGot:      %s",
				count, tc.Cfg, tc.Expected, got)
			failed++

			if failed >= 10 {
				t.Fatal("Too many failures, aborting test.")
			}
		}

		if count%100_000 == 0 {
			t.Logf("Verified %d/%d records...", count, dataCount)
		}
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("Scanner error: %v", err)
	}

	duration := time.Since(start)
	t.Logf("PASS: Verified %d records in %s. Throughput: %.0f/s",
		count, duration, float64(count)/duration.Seconds())
}

func ensureRegressionData(t *testing.T, filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		t.Logf("Regression data file found: %s", filePath)
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	t.Logf("Generating %d regression records to %s (compressed)...", dataCount, filePath)

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	gw, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	if err != nil {
		return err
	}
	defer gw.Close()

	enc := json.NewEncoder(gw)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	start := time.Now()
	for i := range dataCount {
		cfg := Config{
			Input:  randString(rnd, rnd.Intn(50)+5),
			Salt:   randString(rnd, rnd.Intn(20)+5),
			Length: rnd.Intn(64) + 8,
			Level:  randLevel(rnd),
		}

		if i%1000 == 0 {
			cfg.Length = rnd.Intn(1000) + 100
		}

		pwd, err := Generate(cfg)
		if err != nil {
			return fmt.Errorf("generate failed during data creation: %v", err)
		}

		rc := RegressionCase{
			Cfg:      cfg,
			Expected: pwd,
		}

		if err := enc.Encode(rc); err != nil {
			return err
		}

		if (i+1)%100_000 == 0 {
			t.Logf("Generated %d/%d...", i+1, dataCount)
		}
	}

	t.Logf("Generation complete. Time: %s", time.Since(start))
	return nil
}

var testRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*")

func randString(rnd *rand.Rand, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = testRunes[rnd.Intn(len(testRunes))]
	}
	return string(b)
}

func randLevel(rnd *rand.Rand) Level {
	return levels[rnd.Intn(len(levels))]
}
