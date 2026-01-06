---
weight: 4
title: "VOPRF"
---

# Verifiable Oblivious Pseudorandom Function (VOPRF)

Two-party protocol where server evaluates PRF on client's input without learning the input, and client verifies correct evaluation without learning the key.

## OPRF vs VOPRF vs POPRF

| Variant | Server Learns Input | Client Verifies | Public Metadata |
|---------|---------------------|-----------------|-----------------|
| OPRF | No | No | No |
| VOPRF | No | Yes | No |
| POPRF | No | Yes | Yes |

## Protocol (RFC 9497)

### Setup
- Server generates keypair: `(sk, pk)` where `pk = sk·G`
- Client knows `pk`

### Blind
Client blinds input `x`:
```
r ←$ Zq          // random scalar
P = H2C(x)       // hash-to-curve
B = r·P          // blinded point
```

### Evaluate
Server computes on blinded input:
```
Z = sk·B         // evaluate
π = DLEQ(sk, G, pk, B, Z)  // proof
return (Z, π)
```

### Unblind & Verify
Client:
```
verify DLEQ(π, G, pk, B, Z)
N = r⁻¹·Z = r⁻¹·sk·r·P = sk·P
output = H(x, N)
```

## DLEQ Proof

Proves `log_G(pk) = log_B(Z)` (same secret key used) without revealing `sk`.

Schnorr-style:
```
Prover:                         Verifier:
  t ←$ Zq
  A = t·G, D = t·B
  c = H(G, pk, B, Z, A, D)
  s = t - c·sk
  send (c, s)
                                A' = s·G + c·pk
                                D' = s·B + c·Z
                                check c = H(G, pk, B, Z, A', D')
```

## Ciphersuites (RFC 9497)

| Suite | Curve | Hash | Security |
|-------|-------|------|----------|
| P256-SHA256 | NIST P-256 | SHA-256 | 128-bit |
| P384-SHA384 | NIST P-384 | SHA-384 | 192-bit |
| P521-SHA512 | NIST P-521 | SHA-512 | 256-bit |
| ristretto255-SHA512 | Ristretto255 | SHA-512 | 128-bit |
| decaf448-SHAKE256 | Decaf448 | SHAKE256 | 224-bit |

## Security Properties

| Property | Guarantee |
|----------|-----------|
| Obliviousness | Server learns nothing about input |
| Verifiability | Client detects malicious evaluation |
| Pseudorandomness | Output indistinguishable from random without key |
| Unlinkability | Cannot link blind request to unblinded output |

## Applications

| Application | How VOPRF is Used |
|-------------|-------------------|
| Privacy Pass | Blind token signing |
| OPAQUE | Password hardening without server learning password |
| Private Set Intersection | Compare elements without revealing sets |
| Rate limiting | Anonymous but bounded access |

## OPAQUE Integration

VOPRF hardens password authentication:
```
Client                              Server
  |                                   |
  |-- blind(pwd) -------------------->|
  |                                   |
  |<-- OPRF(sk, blind(pwd)) ---------|
  |                                   |
  rwd = unblind(response)             |
  key = KDF(rwd)                      |
```

Server never sees password; compromise doesn't yield offline dictionary attack.

## When to Use

- Need server-assisted computation without revealing input
- Client must verify server behaved honestly (use VOPRF over OPRF)
- Token issuance with unlinkability (Privacy Pass)
- Password hardening (OPAQUE)
- Rate limiting with anonymity

## When Not to Use

- No need for verifiability (plain OPRF is simpler)
- Server is fully trusted (just use HMAC)
- Need public verifiability by third parties (use blind signatures)
- Input domain is small/enumerable (server can brute-force)
- Quantum adversaries (current constructions not post-quantum)

## See Also

- [RFC 9497 - Oblivious Pseudorandom Functions](https://datatracker.ietf.org/doc/html/rfc9497)
- [RFC 9380 - Hashing to Elliptic Curves](https://datatracker.ietf.org/doc/html/rfc9380)
- Privacy Pass (RFC 9576-9578)
- OPAQUE (RFC 9497)

---
*Written by Claude Opus 4.5*
