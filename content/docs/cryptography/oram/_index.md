---
weight: 2
title: "Oblivious RAM"
bookCollapseSection: true
---

# Oblivious RAM (ORAM)

ORAM hides memory access patterns from an adversary observing memory operations. Even if data is encrypted, access patterns can leak information (e.g., which records are accessed, frequency, ordering).

## Threat Model

- Adversary sees all memory addresses accessed (reads/writes)
- Cannot distinguish between reads and writes
- Goal: make access sequences computationally indistinguishable regardless of input

## Key Constructions

| Construction | Client Storage | Bandwidth Overhead | Notes |
|-------------|----------------|-------------------|-------|
| Square Root ORAM | O(√N) | O(√N) | Original Goldreich-Ostrovsky |
| Hierarchical ORAM | O(1) | O(log³N) | Goldreich-Ostrovsky |
| Path ORAM | O(log N) | O(log N) | Practical, widely used |
| Ring ORAM | O(log N) | O(log N) | Improved constants over Path |
| Circuit ORAM | O(1) | O(log N) | Optimized for MPC |

## Path ORAM (Shi et al., 2011)

Most practical construction. Key ideas:
- Binary tree of buckets, each holding O(log N) blocks
- Each block mapped to random leaf; accessed via root-to-leaf path
- **Invariant**: block always on its assigned path
- Access: read path, find block, write back to root, remap to new random leaf
- Stash: client-side overflow buffer (O(log N) expected size)

## Applications

- Secure cloud storage (hide access patterns from cloud provider)
- Oblivious data structures (maps, heaps, etc.)
- Private information retrieval
- Secure multi-party computation
- TEE memory protection (hide patterns from side-channels)

## When to Use

- Single client with private data on untrusted storage
- Access pattern leakage is a real threat (side-channel adversary)
- Read/write workloads (not read-only)
- Data fits in structured storage (arrays, trees, maps)
- TEE applications where memory bus is observable

## When Not to Use

- Read-only public database (use PIR instead)
- Multiple clients accessing shared data (ORAM is single-client)
- Latency-critical applications (O(log N) overhead per access)
- Small datasets where full download is acceptable
- Adversary cannot observe access patterns anyway

## Limitations

- Minimum O(log N) bandwidth overhead (lower bound proven)
- Write-back traffic doubles bandwidth
- Stash overflow probability must be negligible
- Expensive for small random accesses

---
*Written by Claude Opus 4.5*
