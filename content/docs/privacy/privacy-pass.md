---
weight: 3
title: "Privacy Pass"
---

# Privacy Pass

Protocol for anonymous authorization tokens. User proves prior verification (e.g., CAPTCHA) without linking current request to that verification event.

## Problem

CAPTCHAs and rate-limiting create UX friction. Without Privacy Pass:
- User solves CAPTCHA → server remembers IP/cookie
- Tracking across requests, privacy erosion

## Core Idea

1. User solves challenge once, receives batch of **blinded tokens**
2. Server signs tokens without seeing their values
3. User unblinds tokens, stores them
4. Later: user redeems token to bypass challenge
5. Server verifies signature but cannot link to issuance

## Protocol Flow

```
Issuance:
  Client                           Issuer
    |--- blind(token) --------------->|
    |<-- sign(blind(token)) ----------|
    |                                 |
  unblind(signed) = σ

Redemption:
  Client                           Origin
    |--- (token, σ) ----------------->|
    |         verify(token, σ)        |
    |<-- access granted --------------|
```

## Cryptographic Primitives

### VOPRF (Verifiable Oblivious PRF)

- Issuer has secret key `k`
- Client blinds input: `B = blind(t, r)`
- Issuer evaluates: `Z = kB` (without learning `t`)
- Client unblinds: `σ = unblind(Z, r) = H(t)^k`
- Verification via DLEQ proof

### Token Types (RFC 9576, 9577, 9578)

| Type | Primitive | Public Metadata | Private Metadata |
|------|-----------|-----------------|------------------|
| Publicly Verifiable | RSA Blind Sig | Yes | No |
| Privately Verifiable | VOPRF | No | No |
| Rate-Limited | VOPRF + | Yes | Yes |

## Security Properties

| Property | Guarantee |
|----------|-----------|
| Unlinkability | Redemption unlinkable to issuance |
| Unforgeability | Cannot create valid tokens without issuer |
| One-show | Token can only be redeemed once |
| Privacy | Issuer learns nothing during blind signing |

## Architecture (RFC 9576)

```
┌────────┐    ┌────────┐    ┌────────┐
│ Client │────│ Origin │────│ Issuer │
└────────┘    └────────┘    └────────┘
     │             │              │
     │  challenge  │              │
     │<────────────│              │
     │         token request      │
     │────────────────────────────>
     │         blind signature    │
     │<────────────────────────────
     │  redemption │              │
     │────────────>│   verify     │
     │             │──────────────>
```

## Deployments

- **Cloudflare**: CAPTCHA bypass via browser extension
- **Apple Private Access Tokens**: Device attestation without tracking
- **Chrome Trust Tokens**: Anti-fraud without third-party cookies

## Apple Private Access Tokens

Extends Privacy Pass for device attestation:
- Attester (Apple) vouches for device legitimacy
- Origin gets proof of real device without device fingerprint
- Used to bypass CAPTCHAs on iOS/macOS

## Limitations

- Token hoarding: malicious users stockpile tokens
- Issuer centralization: trust concentration
- Metadata leakage: token type/issuer visible at redemption
- Sybil attacks: one verification → many tokens

## See Also

- [RFC 9576 - Privacy Pass Architecture](https://datatracker.ietf.org/doc/html/rfc9576)
- [RFC 9578 - Privacy Pass Issuance Protocol](https://datatracker.ietf.org/doc/html/rfc9578)
- Blind signatures (Chaum)
- Anonymous credentials

---
*Written by Claude Opus 4.5*
