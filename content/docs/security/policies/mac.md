---
title: Mandatory Access Controls (MAC)
weight: 1
---
# Mandatory Access Controls (MAC)

System-enforced access policy where security labels (not users) determine access rights. Contrasts with DAC where resource owners control permissions.

## MAC vs DAC

| Aspect | MAC | DAC |
|--------|-----|-----|
| Policy authority | System/admin | Resource owner |
| Label assignment | Mandatory | Discretionary |
| Trojan resistance | Strong | Weak |
| Flexibility | Low | High |

## Core Models

### Bell-LaPadula (Confidentiality)

- **Simple Security**: No read up (subject cannot read higher classification)
- **Star Property**: No write down (subject cannot write to lower classification)
- Prevents information leakage to lower clearance levels

### Biba (Integrity)

- Dual of Bell-LaPadula
- **Simple Integrity**: No read down
- **Star Integrity**: No write up
- Prevents corruption from untrusted sources

### Clark-Wilson (Commercial Integrity)

- Well-formed transactions via Transformation Procedures (TPs)
- Constrained Data Items (CDIs) only modified by certified TPs
- Separation of duties enforced

## Linux Implementations

| Framework | Approach | Use Case |
|-----------|----------|----------|
| SELinux | Type Enforcement + MLS | High-security servers |
| AppArmor | Path-based profiles | Desktop/container confinement |
| Smack | Simple labels | Embedded/IoT |
| TOMOYO | Path-based, learning mode | Gradual policy development |

### SELinux Type Enforcement

```
# Process httpd_t can read files labeled httpd_sys_content_t
allow httpd_t httpd_sys_content_t:file { read open getattr };
```

### AppArmor Profile

```
/usr/bin/nginx {
  /var/www/** r,
  /var/log/nginx/** w,
  deny /etc/shadow r,
}
```

## When to Use MAC

- Multi-tenant systems with strict isolation requirements
- Systems processing data at multiple classification levels
- Defense against privilege escalation and confused deputy
- Container/sandbox hardening

---
*Written by Claude Opus 4.5*
