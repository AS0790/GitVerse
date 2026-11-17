# Network Latency & Concurrency Benchmarks

This document records the performance optimizations and benchmark results achieved across GitVerse versions.

---

## ⚡ Comparative Latency Metrics

| Metric | Sequential Fetching (v1.2) | Parallel Concurrency (v1.3) | Improvement |
| :--- | :---: | :---: | :---: |
| **Cold Profile Query (`TheVicky1`)** | ~3,400 ms | **~1,700 ms** | **50% Faster** |
| **Warm In-Memory Cache Query** | ~1,200 ms | **17 ms** | **98.5% Faster** |
| **Parallel Badge Fan-Out** | 4 roundtrips (serial) | **1 roundtrip (parallel)** | **4x Concurrency** |
| **TLS Handshake Latency** | 250ms per request | **0ms (reused keep-alive)** | **Zero Handshake Overhead** |

---

## 🔬 Benchmark Methodology

Benchmarks were measured using `node` testing against active GitHub profiles, testing both cold DNS/TLS resolution and warm keep-alive agent connection pooling across 32 sockets.

