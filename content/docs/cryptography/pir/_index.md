---
weight: 3
title: "Private Information Retrieval"
bookCollapseSection: true
---

# Private Information Retrieval (PIR)

PIR allows a client to retrieve an item from a database without revealing which item was accessed. Unlike ORAM, the server doesn't store client state—each query is independent.

## Taxonomy

| Type | Servers | Assumption | Communication |
|------|---------|------------|---------------|
| Information-theoretic PIR | ≥2 (non-colluding) | None | O(N^(1/k)) for k servers |
| Computational PIR (cPIR) | 1 | Crypto hardness | O(polylog N) possible |
| Symmetric PIR (SPIR) | ≥1 | Varies | Client learns only queried item |

## Information-Theoretic PIR

- **Chor et al. (1995)**: Original 2-server scheme, O(N^(1/3)) communication
- **Key insight**: XOR of responses from non-colluding servers reveals only queried bit
- Lower bound: single server IT-PIR requires Ω(N) communication (trivial download)

## Computational PIR

Single-server schemes rely on:
- **Homomorphic encryption**: Query = encrypted index, server computes on ciphertext
- **Lattice-based**: SealPIR, OnionPIR (RLWE-based)
- **DPF-based**: Distributed Point Functions for 2-server with sublinear computation

### Notable Constructions

| Scheme | Basis | Communication | Server Computation |
|--------|-------|---------------|-------------------|
| Kushilevitz-Ostrovsky | QR | O(N^ε) | O(N) |
| SealPIR | RLWE | O(√N) | O(N) |
| OnionPIR | RLWE | O(1) | O(N) |
| Checklist | DPF | O(√N) | O(√N) amortized |

## PIR vs ORAM

| | PIR | ORAM |
|--|-----|------|
| Server state | Stateless | Stateful |
| Client state | None | O(log N) typical |
| Write support | Read-only (usually) | Read/write |
| Use case | Public DB, many clients | Private storage, single client |

## Applications

- Private DNS queries (Oblivious DoH)
- Private contact discovery (Signal)
- Anonymous credential lookup
- Private media streaming
- Certificate transparency checks

## When to Use

- Public database with many clients (stateless queries)
- Read-only access patterns
- Query privacy matters more than server efficiency
- Can tolerate O(N) server computation per query
- Non-colluding server assumption is realistic (for IT-PIR)

## When Not to Use

- Need write access (use ORAM)
- Single client with private data (ORAM more efficient)
- Cannot afford O(N) server work per query
- Low-latency requirements
- Database changes frequently (invalidates preprocessing)

## Practical Considerations

- Server computation still O(N) for most schemes (must touch all data)
- Batch queries and preprocessing can amortize costs
- Hybrid approaches: PIR + local cache
- Privacy vs performance tradeoff remains significant

---
*Written by Claude Opus 4.5*
