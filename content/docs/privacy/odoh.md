---
weight: 2
title: "Oblivious DNS over HTTPS (ODoH)"
---

# Oblivious DNS over HTTPS (ODoH)

ODoH decouples DNS query content from client identity by introducing a proxy between client and resolver.

## Problem with DoH/DoT

DNS over HTTPS/TLS encrypts queries from eavesdroppers but the resolver still sees:
- Client IP address
- Full query content
- Query timing

Resolver becomes a privacy bottleneck—must be fully trusted.

## ODoH Architecture

```
Client → Proxy (Oblivious Relay) → Target (Resolver)
```

| Entity | Knows Client IP | Knows Query |
|--------|-----------------|-------------|
| Proxy | Yes | No (encrypted) |
| Target | No (sees proxy IP) | Yes (decrypts) |
| Colluding both | Yes | Yes |

## Protocol (RFC 9230)

1. Client fetches Target's HPKE public key (via DNS or HTTPS)
2. Client encrypts query: `ct = HPKE.Seal(pk_target, query)`
3. Client sends `ct` to Proxy
4. Proxy forwards to Target (cannot decrypt)
5. Target decrypts, resolves, encrypts response
6. Response returns via Proxy

### Key Encapsulation

Uses HPKE (Hybrid Public Key Encryption):
- KEM: X25519 or P-256
- KDF: HKDF-SHA256
- AEAD: AES-128-GCM or ChaCha20-Poly1305

## Security Properties

| Property | Guarantee |
|----------|-----------|
| Query confidentiality | Proxy cannot read queries |
| Client anonymity | Target cannot identify client |
| Response integrity | AEAD authentication |
| Forward secrecy | Per-query ephemeral keys |

## Threat Model

**Assumes non-colluding Proxy and Target.** If they collude:
- Full deanonymization possible
- Reduces to standard DoH

Mitigations:
- Use Proxy and Target from different jurisdictions/organizations
- Tor-like multi-hop extensions (not in spec)

## Comparison

| Protocol | Encryption | Hides from ISP | Hides from Resolver |
|----------|------------|----------------|---------------------|
| Plain DNS | None | No | No |
| DoT/DoH | Yes | Yes | No |
| ODoH | Yes | Yes | Yes* |
| DNSCrypt Anonymized | Yes | Yes | Yes* |

*Assuming non-collusion

## Deployments

- Cloudflare: `odoh.cloudflare-dns.com`
- Apple Private Relay (uses ODoH-like design)
- Fastly as relay provider

## Limitations

- Additional latency (extra hop)
- Requires trust in non-collusion assumption
- Key distribution bootstrapping problem
- Limited resolver support

## See Also

- [RFC 9230 - Oblivious DNS over HTTPS](https://datatracker.ietf.org/doc/html/rfc9230)
- HPKE (RFC 9180)
- Oblivious HTTP (OHTTP) - generalization of ODoH

---
*Written by Claude Opus 4.5*
