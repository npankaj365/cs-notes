---
weight: 1
bookFlatSection: false
title: "Franking"
---

# Franking

Message franking enables abuse reporting in end-to-end encrypted messaging while preserving sender deniability for non-reported messages.

## Core Idea

Sender commits to message content; platform can verify reported messages without seeing all messages.

## Protocol (Facebook/Meta)

1. **Sending**: Sender creates commitment `c = HMAC(k_f, m)` where `k_f` is franking key
2. **Delivery**: Platform attaches its own tag `t = HMAC(k_p, c)` without seeing plaintext
3. **Reporting**: Recipient reveals `(m, k_f)` to platform
4. **Verification**: Platform recomputes commitment, verifies its tag

## Security Properties

| Property | Guarantee |
|----------|-----------|
| Sender binding | Sender cannot deny reported message |
| Platform binding | Platform cannot forge reports |
| Confidentiality | Unreported messages remain private |
| Deniability | Non-reported messages have plausible deniability |

## Cryptographic Construction

```
Franking tag: f = (c, k_f) where c = Commit(k_f, m)
Platform tag: t = MAC(k_platform, c || sender || recipient || timestamp)
Report: (m, k_f, c, t)
Verify: Open(c, k_f) = m AND VerifyMAC(t)
```

## Threat Model

- Malicious sender trying to deny abuse
- Malicious recipient trying to frame sender
- Curious platform (should not learn unreported content)

## Limitations

- Requires recipient cooperation for reporting
- Screenshot-based reporting still possible (weaker guarantees)
- Metadata still visible to platform
- Group messaging complicates attribution

## See Also

- Asymmetric Message Franking (Grubbs et al.)
- Compactly Committing AE (Dodis et al.)


---
*Written by Claude Opus 4.5*
