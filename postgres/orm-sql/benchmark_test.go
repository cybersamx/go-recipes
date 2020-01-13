package main

import (
	_ "github.com/lib/pq"
	"os"
	"testing"
)

const (
	numSeeds = 1000
)

// --- SQL ---

func benchmarkInsertSQL(b *testing.B, n int) {
	clearTables()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		insertDataSQL(n)
	}
}

func benchmarkUpdateSQL(b *testing.B, n int) {
	clearTables()
	insertDataSQL(numSeeds)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		updateDataSQL(n)
	}
}

func benchmarkSelectSQL(b *testing.B, n int) {
	clearTables()
	insertDataSQL(numSeeds)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		selectDataSQL(n)
	}
}

// --- XORM ---

func benchmarkInsertXORM(b *testing.B, n int) {
	clearTables()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		insertDataXORM(n)
	}
}

func benchmarkUpdateXORM(b *testing.B, n int) {
	clearTables()
	insertDataSQL(numSeeds)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		updateDataXORM(n)
	}
}

func benchmarkSelectXORM(b *testing.B, n int) {
	clearTables()
	insertDataSQL(numSeeds)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		selectDataXORM(n)
	}
}

// --- GORM ---

func benchmarkInsertGORM(b *testing.B, n int) {
	clearTables()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		insertDataGORM(n)
	}
}

func benchmarkUpdateGORM(b *testing.B, n int) {
	clearTables()
	insertDataSQL(numSeeds)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		updateDataGORM(n)
	}
}

func benchmarkSelectGORM(b *testing.B, n int) {
	clearTables()
	insertDataSQL(numSeeds)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		selectDataGORM(n)
	}
}


// --- Main Test Func ---

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// --- SQL ---

// Insert

func BenchmarkInsertSQL_1(b *testing.B) {
	benchmarkInsertSQL(b, 1)
}

func BenchmarkInsertSQL_10(b *testing.B) {
	benchmarkInsertSQL(b, 10)
}

func BenchmarkInsertSQL_100(b *testing.B) {
	benchmarkInsertSQL(b, 100)
}

func BenchmarkInsertSQL_1000(b *testing.B) {
	benchmarkInsertSQL(b, 1000)
}

// Update

func BenchmarkUpdateSQL_1(b *testing.B) {
	benchmarkUpdateSQL(b, 1)
}

func BenchmarkUpdateSQL_10(b *testing.B) {
	benchmarkUpdateSQL(b, 10)
}

func BenchmarkUpdateSQL_100(b *testing.B) {
	benchmarkUpdateSQL(b, 100)
}

func BenchmarkUpdateSQL_1000(b *testing.B) {
	benchmarkUpdateSQL(b, 1000)
}

// Select

func BenchmarkSelectSQL_1(b *testing.B) {
	benchmarkSelectSQL(b, 1)
}

func BenchmarkSelectSQL_10(b *testing.B) {
	benchmarkSelectSQL(b, 10)
}

func BenchmarkSelectSQL_100(b *testing.B) {
	benchmarkSelectSQL(b, 100)
}

func BenchmarkSelectSQL_1000(b *testing.B) {
	benchmarkSelectSQL(b, 1000)
}

// --- XORM ---

// Insert

func BenchmarkInsertXORM_1(b *testing.B) {
	benchmarkInsertXORM(b, 1)
}

func BenchmarkInsertXORM_10(b *testing.B) {
	benchmarkInsertXORM(b, 10)
}

func BenchmarkInsertXORM_100(b *testing.B) {
	benchmarkInsertXORM(b, 100)
}

func BenchmarkInsertXORM_1000(b *testing.B) {
	benchmarkInsertXORM(b, 1000)
}

// Update

func BenchmarkUpdateXORM_1(b *testing.B) {
	benchmarkUpdateXORM(b, 1)
}

func BenchmarkUpdateXORM_10(b *testing.B) {
	benchmarkUpdateXORM(b, 10)
}

func BenchmarkUpdateXORM_100(b *testing.B) {
	benchmarkUpdateXORM(b, 100)
}

func BenchmarkUpdateXORM_1000(b *testing.B) {
	benchmarkUpdateXORM(b, 1000)
}

// Select

func BenchmarkSelectXORM_1(b *testing.B) {
	benchmarkSelectXORM(b, 1)
}

func BenchmarkSelectXORM_10(b *testing.B) {
	benchmarkSelectXORM(b, 10)
}

func BenchmarkSelectXORM_100(b *testing.B) {
	benchmarkSelectXORM(b, 100)
}

func BenchmarkSelectXORM_1000(b *testing.B) {
	benchmarkSelectXORM(b, 1000)
}

// --- GORM ---

// Insert

func BenchmarkInsertGORM_1(b *testing.B) {
	benchmarkInsertGORM(b, 1)
}

func BenchmarkInsertGORM_10(b *testing.B) {
	benchmarkInsertGORM(b, 10)
}

func BenchmarkInsertGORM_100(b *testing.B) {
	benchmarkInsertGORM(b, 100)
}

func BenchmarkInsertGORM_1000(b *testing.B) {
	benchmarkInsertGORM(b, 1000)
}

// Update

func BenchmarkUpdateGORM_1(b *testing.B) {
	benchmarkUpdateGORM(b, 1)
}

func BenchmarkUpdateGORM_10(b *testing.B) {
	benchmarkUpdateGORM(b, 10)
}

func BenchmarkUpdateGORM_100(b *testing.B) {
	benchmarkUpdateGORM(b, 100)
}

func BenchmarkUpdateGORM_1000(b *testing.B) {
	benchmarkUpdateGORM(b, 1000)
}

// Select

func BenchmarkSelectGORM_1(b *testing.B) {
	benchmarkSelectGORM(b, 1)
}

func BenchmarkSelectGORM_10(b *testing.B) {
	benchmarkSelectGORM(b, 10)
}

func BenchmarkSelectGORM_100(b *testing.B) {
	benchmarkSelectGORM(b, 100)
}

func BenchmarkSelectGORM_1000(b *testing.B) {
	benchmarkSelectGORM(b, 1000)
}
