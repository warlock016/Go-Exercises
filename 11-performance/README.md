# Module 11: Performance

**Focus Area:** Benchmarking, Profiling, and Optimization
**Prerequisites:** Module 09 (Concurrency), Modules 01-04 (Fundamentals)
**Estimated Time:** 10-12 hours
**Exercises:** 12

---

## Overview

Performance optimization in Go requires measurement, not guesswork. This module teaches you to use Go's powerful benchmarking and profiling tools to find bottlenecks and apply targeted optimizations.

**What You'll Learn:**
- Writing and interpreting benchmarks
- Reducing memory allocations
- CPU and memory profiling with pprof
- Understanding escape analysis and inlining
- Cache-friendly data structures
- Benchmarking concurrent code correctly

---

## Module Structure

### Tier 1: Introduction (Exercises 01-03)

| Exercise | Concept | Time |
|----------|---------|------|
| 01_benchmarking_basics | b.N, b.ResetTimer, sub-benchmarks | 30-35 min |
| 02_memory_allocation | b.ReportAllocs, reducing allocations | 35-40 min |
| 03_string_building | strings.Builder vs concatenation | 30-35 min |

### Tier 2: Application (Exercises 04-07)

| Exercise | Concept | Time |
|----------|---------|------|
| 04_slice_preallocation | make with capacity, append performance | 35-40 min |
| 05_map_optimization | Map sizing, struct keys vs string keys | 40-45 min |
| 06_sync_pool | Object pooling, GC pressure reduction | 45-50 min |
| 07_profiling_cpu | pprof CPU profiling, flame graphs | 50-60 min |

### Tier 3: Integration (Exercises 08-10)

| Exercise | Concept | Time |
|----------|---------|------|
| 08_profiling_memory | Heap profiling, finding allocation sites | 50-60 min |
| 09_escape_analysis | Stack vs heap, -gcflags="-m" | 45-50 min |
| 10_inlining | Function inlining, when it helps/hurts | 45-50 min |

### Tier 4: Mastery (Exercises 11-12)

| Exercise | Concept | Time |
|----------|---------|------|
| 11_cache_optimization | Cache-friendly data structures | 55-65 min |
| 12_concurrent_performance | Benchmarking concurrent code | 60-75 min |

---

## Key Commands

### Run Benchmarks
```bash
go test -bench=. -benchmem
```

### CPU Profile
```bash
go test -cpuprofile=cpu.prof -bench=BenchmarkX
go tool pprof -http=:8080 cpu.prof
```

### Memory Profile
```bash
go test -memprofile=mem.prof -bench=BenchmarkX
go tool pprof -http=:8080 mem.prof
```

### Escape Analysis
```bash
go build -gcflags="-m" ./...
```

### Compare Benchmarks
```bash
go install golang.org/x/perf/cmd/benchstat@latest
go test -bench=. -count=10 > old.txt
# make changes
go test -bench=. -count=10 > new.txt
benchstat old.txt new.txt
```

---

## The Optimization Loop

1. **Measure** - Benchmark to establish baseline
2. **Profile** - Find the actual bottleneck
3. **Optimize** - Apply targeted fix
4. **Verify** - Benchmark to confirm improvement
5. **Repeat** - Until goals met

---

## Success Criteria

By completing this module, you should be able to:
- Write meaningful benchmarks
- Use pprof to identify bottlenecks
- Reduce allocations through pooling and preallocation
- Read escape analysis output
- Design cache-friendly data structures
- Make optimization decisions based on data
