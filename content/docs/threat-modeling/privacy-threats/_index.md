---
weight: 2
bookCollapseSection: true
title: "Privacy Threats"
---

# Privacy Threat Modeling

Privacy threats are distinct from security threats. A system can be secure (no unauthorized access) but not private (authorized parties learn too much).

## Pfitzmann-Hansen Terminology

Foundational vocabulary for privacy properties.

### Anonymity Set

Group of subjects who might have performed an action. Larger set = more anonymity.

### Core Properties

| Property | Definition | Negation |
|----------|------------|----------|
| **Anonymity** | Subject not identifiable within anonymity set | Identifiability |
| **Unlinkability** | Cannot link two items (actions, identities, messages) | Linkability |
| **Undetectability** | Cannot distinguish if item exists | Detectability |
| **Unobservability** | Undetectability + anonymity of involved subjects | Observability |
| **Pseudonymity** | Use of pseudonym (linkable within scope, unlinkable across) | - |

### Relationship Hierarchy

```
Unobservability
      ↓ implies
Anonymity + Undetectability
      ↓ implies
Unlinkability
```

## LINDDUN

Privacy-focused threat modeling framework (counterpart to STRIDE).

| Threat | Definition | Example |
|--------|------------|---------|
| **L**inkability | Link two items of interest | Correlating browsing sessions via fingerprint |
| **I**dentifiability | Identify subject from data | Re-identification from "anonymous" dataset |
| **N**on-repudiation | Cannot deny action (privacy violation when unwanted) | Signed audit logs expose whistleblower |
| **D**etectability | Determine existence of item | Traffic analysis reveals VPN usage |
| **D**isclosure | Excessive data exposure | App collects more than needed |
| **U**nawareness | Data subject not informed | Hidden tracking pixels |
| **N**on-compliance | Violating regulations/policies | GDPR breach |

### LINDDUN Process

1. Define DFD with privacy-relevant data flows
2. Map LINDDUN threats to DFD elements
3. Use threat trees for systematic enumeration
4. Identify Privacy Enhancing Technologies (PETs)
5. Document residual risks

## k-Anonymity and Extensions

| Model | Guarantee | Limitation |
|-------|-----------|------------|
| k-anonymity | Each record indistinguishable from k-1 others on quasi-identifiers | Homogeneity attack, background knowledge |
| l-diversity | Each equivalence class has l distinct sensitive values | Skewness attack |
| t-closeness | Distribution of sensitive attribute in class close to overall distribution | Computational cost |

## Differential Privacy Threat Model

- **Adversary**: Has access to query results, arbitrary auxiliary information
- **Guarantee**: Cannot determine if specific individual is in dataset
- **Mechanism**: Add calibrated noise (Laplace, Gaussian)
- **Parameter ε**: Privacy budget (lower = more private)

## Privacy Attack Categories

| Attack | Description | Mitigations |
|--------|-------------|-------------|
| Re-identification | Link "anonymous" data to identity | k-anonymity, DP |
| Inference | Deduce sensitive info from aggregates | DP, query auditing |
| Membership inference | Determine if record was in training set | DP training, regularization |
| Attribute inference | Infer unknown attribute from known ones | Minimize data collection |
| Linkage | Join datasets on quasi-identifiers | Data minimization, synthetic data |
| Traffic analysis | Learn from metadata/timing | Padding, mixing, constant-rate |

## When to Use

- Systems processing PII or sensitive data
- GDPR/CCPA compliance requirements
- Anonymous communication systems
- ML pipelines with privacy requirements
- Any system where "honest-but-curious" is the threat model

## See Also

- [LINDDUN](https://linddun.org/)
- Pfitzmann & Hansen, "Anonymity, Unlinkability, Unobservability, Pseudonymity, and Identity Management"

---
*Written by Claude Opus 4.5*
