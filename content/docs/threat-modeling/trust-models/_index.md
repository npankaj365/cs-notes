---
weight: 3
bookCollapseSection: true
title: "Trust Models"
---

# Trust Models & Adversary Capabilities

Defining what the adversary can and cannot do. Critical for meaningful security/privacy guarantees.

## Standard Adversary Models

| Model | Behavior | Capabilities | Use Case |
|-------|----------|--------------|----------|
| Honest | Follows protocol | None adversarial | Baseline |
| Semi-honest (HbC) | Follows protocol, tries to learn | Observes all internal state | MPC, PIR |
| Covert | Cheats if undetectable | Deviates if detection prob < ε | Practical MPC |
| Malicious | Arbitrary deviation | Full protocol deviation | Strongest guarantees |

## Semi-Honest (Honest-but-Curious)

- Executes protocol correctly
- Attempts to infer additional information from transcript
- Models curious cloud providers, passive eavesdroppers
- Sufficient for many practical scenarios

```
Example: Cloud storage provider
- Honestly stores encrypted files
- Tries to learn content from access patterns
- Solution: ORAM
```

## Malicious Adversary

- Arbitrary deviation from protocol
- May send malformed messages, abort, collude
- Requires stronger techniques:
  - Zero-knowledge proofs
  - Authenticated garbling
  - Verifiable computation

## Collusion Models

| Model | Description |
|-------|-------------|
| No collusion | Each party independent |
| Threshold (t-of-n) | Up to t parties collude |
| Dishonest majority | > n/2 parties malicious |
| Honest majority | < n/2 parties malicious |

### Threshold Assumptions

| Threshold | Enables |
|-----------|---------|
| t < n | Basic MPC possible |
| t < n/2 | Information-theoretic MPC |
| t < n/3 | Byzantine agreement without setup |

## Network Adversary Models

| Model | Capabilities |
|-------|--------------|
| Passive/Eavesdropper | Read all traffic |
| Active/MitM | Read, modify, inject, drop |
| Global | Observe all network links |
| Local | Observe subset of links |
| Adaptive | Compromise nodes based on observations |
| Static | Corruption set fixed before execution |

## TEE Threat Models

| Adversary | Capabilities | Trust |
|-----------|--------------|-------|
| Software attacker | Control OS, hypervisor | Trust CPU + enclave code |
| Physical attacker | Probe memory bus | Trust CPU package |
| Side-channel attacker | Measure timing, power, cache | Trust varies by mitigation |
| Microarchitectural | Exploit speculative execution | Trust specific CPU + patches |

### What TEEs Typically DON'T Protect Against

- Side-channel attacks (timing, cache, power)
- Denial of service (OS controls scheduling)
- Rollback attacks (without monotonic counters)
- Physical attacks on CPU package
- Bugs in enclave code

## Cryptographic Assumptions

| Assumption | Broken By |
|------------|-----------|
| DLog (discrete log) | Quantum (Shor's) |
| RSA | Quantum (Shor's) |
| DDH | Quantum |
| LWE/RLWE | Unknown (believed post-quantum) |
| Hash collision resistance | Quantum (Grover's, √ speedup) |
| Random Oracle Model | Instantiation may fail |

## Modeling Checklist

When specifying a threat model, define:

1. **Adversary goal**: Confidentiality, integrity, availability, privacy?
2. **Adversary position**: Insider, outsider, network, physical?
3. **Adversary capabilities**: Computational (PPT?), storage, network access?
4. **Corruption model**: Static/adaptive, threshold, collusion?
5. **Trust assumptions**: What is NOT compromised?
6. **Cryptographic assumptions**: Hardness assumptions relied upon
7. **Out of scope**: Explicit non-goals

## Common Pitfalls

- Assuming "encrypted = secure" (ignores metadata, access patterns)
- Trusting the client in client-server model
- Ignoring side channels in threat model
- Unstated collusion assumptions
- Assuming honest majority without justification
- Confusing authentication with authorization

---
*Written by Claude Opus 4.5*
