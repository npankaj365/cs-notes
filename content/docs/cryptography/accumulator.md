---
weight: 5
title: "Accumulator"
---

# Cryptographic Accumulator

Compact representation of a set that allows succinct membership (or non-membership) proofs without revealing other elements.

## Core Idea

- **Accumulator value**: Single short digest representing entire set
- **Witness**: Proof that element is (or isn't) in the set
- **Verification**: Check witness against accumulator in O(1) or O(log n)

## Types

| Type | Membership | Non-membership | Dynamic | Assumption |
|------|------------|----------------|---------|------------|
| RSA Accumulator | Yes | Yes (with trapdoor) | Yes | Strong RSA |
| Bilinear Accumulator | Yes | Yes | Yes | q-SDH |
| Merkle Tree | Yes | Yes (with sorted tree) | Yes | Hash collision resistance |
| Bloom Filter | Yes (probabilistic) | No | Append-only | None (probabilistic) |

## RSA Accumulator (Benaloh-de Mare, Barić-Pfitzmann)

### Setup
- `n = pq` (RSA modulus, trapdoor = factorization)
- `g` generator of QR_n
- Hash function `H: {0,1}* → primes`

### Accumulate
```
A = g^(∏ H(xᵢ)) mod n
```

### Witness Generation
For element x with prime representative e = H(x):
```
w = g^(∏ H(xⱼ) for j≠i) mod n
```

### Verification
```
w^e ≡ A (mod n)
```

### Non-membership (requires trapdoor or Bezout coefficients)
Prove `gcd(e, ∏eᵢ) = 1` using Bezout: `ae + b(∏eᵢ) = 1`

## Bilinear Accumulator (Nguyen)

Uses pairing `e: G₁ × G₂ → G_T`

### Setup
- Secret `s`, public params `(g, g^s, g^s², ..., g^sᵠ)`

### Accumulate
```
A = g^∏(s + xᵢ)
```

### Witness
```
w = g^(∏(s + xⱼ) for j≠i)
```

### Verification
```
e(w, g^s · g^xᵢ) = e(A, g)
```

## Merkle Tree as Accumulator

- Root hash = accumulator value
- Witness = authentication path (O(log n) hashes)
- Non-membership: sorted tree + prove neighbors

| Operation | Complexity |
|-----------|------------|
| Accumulate | O(n) |
| Witness size | O(log n) |
| Verify | O(log n) |
| Update | O(log n) |

## Dynamic Accumulators

Support efficient updates without recomputing all witnesses:

| Operation | RSA | Bilinear | Merkle |
|-----------|-----|----------|--------|
| Add element | O(1) acc, O(n) witnesses | O(1) acc, O(n) witnesses | O(log n) |
| Remove element | O(1) acc, O(n) witnesses | O(1) acc, O(n) witnesses | O(log n) |
| Batch update | Possible | Possible | O(k log n) |

## Security Properties

| Property | Definition |
|----------|------------|
| Collision-freeness | Cannot create valid witness for non-member |
| Undeniability | Cannot deny membership of actual member |
| Indistinguishability | Accumulator hides set contents |

## Applications

| Application | How Accumulator is Used |
|-------------|------------------------|
| Revocation lists (CRL) | Non-membership proves certificate not revoked |
| Anonymous credentials | Prove credential in valid set without revealing which |
| Blockchain (UTXO) | Compact representation of unspent outputs |
| Group signatures | Prove membership in signing group |
| Authenticated data structures | Succinct proofs for outsourced databases |

## When to Use

- Need succinct membership proofs (constant or log-size)
- Set is large but proofs must be small
- Verifier shouldn't learn other set elements
- Anonymous credential systems
- Stateless blockchain clients (UTXO commitments)

## When Not to Use

- Set is small (just send the set)
- Need efficient non-membership without trapdoor (use Merkle with sorted leaves)
- Witness update frequency is high (O(n) witness updates per change)
- Post-quantum requirement (RSA/bilinear not PQ-secure; use Merkle)
- Probabilistic membership is acceptable (Bloom filter is simpler)

## See Also

- [Batching Techniques for Accumulators](https://eprint.iacr.org/2018/1188) (Boneh et al.)
- Vector Commitments (generalization)
- Verkle Trees (vector commitment + Merkle structure)

---
*Written by Claude Opus 4.5*
