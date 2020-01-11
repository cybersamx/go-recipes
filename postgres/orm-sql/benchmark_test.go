package main

import (
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func benchmarkInsertSQL(b *testing.B, n int) {
	clearTables()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		insertDataSQL(n)
	}
}

func benchmarkInsertXORM(b *testing.B, n int) {
	clearTables()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		insertDataXORM(n)
	}
}

func benchmarkUpdateSQL(b *testing.B, n int) {
	insertDataSQL(n)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		updateDataSQL(n)
	}
}

func benchmarkUpdateXORM(b *testing.B, n int) {
	insertDataXORM(n)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		updateDataXORM(n)
	}
}

func benchmarkSelectSQL(b *testing.B, n int) {
	insertDataSQL(n)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		selectDataSQL(n)
	}
}

func benchmarkSelectXORM(b *testing.B, n int) {
	insertDataXORM(n)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		selectDataXORM(n)
	}
}


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

func BenchmarkInsertSQL_10000(b *testing.B) {
	benchmarkInsertSQL(b, 10000)
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

func BenchmarkUpdateSQL_10000(b *testing.B) {
	benchmarkUpdateSQL(b, 10000)
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

func BenchmarkSelectSQL_10000(b *testing.B) {
	benchmarkSelectSQL(b, 10000)
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

func BenchmarkInsertXORM_10000(b *testing.B) {
	benchmarkInsertXORM(b, 10000)
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

func BenchmarkUpdateXORM_10000(b *testing.B) {
	benchmarkUpdateXORM(b, 10000)
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

func BenchmarkSelectXORM_10000(b *testing.B) {
	benchmarkSelectXORM(b, 10000)
}
