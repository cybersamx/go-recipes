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


func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// --- SQL ---

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

// --- XORM ---

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
