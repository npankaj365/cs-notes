---
weight: 1
bookCollapseSection: true
title: "Frameworks"
---

# Threat Modeling Frameworks

## STRIDE

Microsoft's mnemonic for threat categories. Apply per-element on DFD.

| Threat | Definition | Violated Property | Typical Target |
|--------|------------|-------------------|----------------|
| **S**poofing | Impersonating something/someone | Authentication | External entities, processes |
| **T**ampering | Modifying data or code | Integrity | Data flows, data stores |
| **R**epudiation | Denying actions | Non-repudiation | Processes |
| **I**nformation Disclosure | Exposing data | Confidentiality | Data flows, data stores |
| **D**enial of Service | Degrading availability | Availability | All elements |
| **E**levation of Privilege | Gaining unauthorized access | Authorization | Processes |

### STRIDE-per-Element

| DFD Element | Applicable Threats |
|-------------|-------------------|
| External Entity | S, R |
| Process | S, T, R, I, D, E |
| Data Store | T, R, I, D |
| Data Flow | T, I, D |

## DREAD (Risk Scoring)

Score 1-10 for each factor, average for overall risk.

| Factor | Question |
|--------|----------|
| **D**amage | How bad if exploited? |
| **R**eproducibility | How easy to reproduce? |
| **E**xploitability | How easy to exploit? |
| **A**ffected users | How many impacted? |
| **D**iscoverability | How easy to find? |

{{< hint warning >}}
DREAD is deprecated at Microsoft due to subjectivity. Consider CVSS or custom risk matrices.
{{< /hint >}}

## Attack Trees

Hierarchical decomposition of attack goals.

```
[Root: Compromise User Account]
├── [OR] Steal Credentials
│   ├── [AND] Phishing
│   │   ├── Craft email
│   │   └── User clicks link
│   ├── [OR] Credential stuffing
│   └── [OR] Keylogger
├── [OR] Session Hijacking
│   ├── XSS to steal cookie
│   └── Network sniffing (no TLS)
└── [OR] Account Recovery Abuse
    └── Social engineer support
```

### Annotations

- **Cost**: Resources needed (time, money, skill)
- **Probability**: Likelihood of success
- **Detection**: Chance of being caught
- **Boolean**: AND (all children needed) vs OR (any child sufficient)

## Cyber Kill Chain (Lockheed Martin)

Sequential phases of an intrusion:

| Phase | Description | Defender Action |
|-------|-------------|-----------------|
| 1. Reconnaissance | Research target | OSINT monitoring |
| 2. Weaponization | Create exploit payload | Threat intel |
| 3. Delivery | Transmit payload (email, USB, web) | Email filtering, WAF |
| 4. Exploitation | Trigger vulnerability | Patching, sandboxing |
| 5. Installation | Persist on system | EDR, integrity monitoring |
| 6. C2 | Establish command channel | Network monitoring, DNS filtering |
| 7. Actions on Objectives | Exfil, destroy, ransom | DLP, backups |

## MITRE ATT&CK

Knowledge base of adversary tactics and techniques based on real-world observations.

| Tactic | Example Techniques |
|--------|-------------------|
| Initial Access | Phishing, supply chain compromise |
| Execution | PowerShell, scheduled tasks |
| Persistence | Registry run keys, implant |
| Privilege Escalation | Token manipulation, exploits |
| Defense Evasion | Obfuscation, rootkits |
| Credential Access | Dumping, keylogging |
| Lateral Movement | Pass-the-hash, RDP |
| Exfiltration | C2 channel, cloud storage |

## When to Use What

| Framework | Best For |
|-----------|----------|
| STRIDE | Systematic design review, DFD-based analysis |
| Attack Trees | Specific attack goal analysis, cost modeling |
| Kill Chain | Incident response, detection gap analysis |
| MITRE ATT&CK | Detection engineering, red team planning |

---
*Written by Claude Opus 4.5*
