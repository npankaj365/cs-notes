---
title: Replay Attack
weight: 1
---
# Replay Attack

## One-Line Essence:
Capturing a valid transmitted data and maliciously repeating it later to trick a system

## Real-World Analogy
Alice capturing Bob's "Hey Google" message to unlock Bob's phone in his absence.

## Key Points

- Attacker records legitimate protocol messages and retransmits them later
- Works because message is valid—authentication succeeds despite replay
- Particularly dangerous in authentication, financial transactions, command-and-control

## Attack Scenarios

| Context | Attack | Impact |
|---------|--------|--------|
| Auth tokens | Replay captured JWT/session token | Account takeover |
| Payment systems | Re-submit signed transaction | Double spending |
| Keyless entry | Replay RF signal | Vehicle theft |
| API requests | Replay signed request | Duplicate operations |

## Defense Patterns

| Defense | Mechanism | Trade-off |
|---------|-----------|-----------|
| Nonces | Server provides unique challenge per request | State management |
| Timestamps | Reject messages outside time window | Clock synchronization |
| Sequence numbers | Reject out-of-order/duplicate seq | State tracking |
| Session binding | Tie to TLS session/channel | Protocol complexity |

## Code Pattern

```python
# Defense: timestamp + nonce validation
def verify_request(msg, sig, timestamp, nonce):
    # Reject if too old (5 minute window)
    if abs(time.time() - timestamp) > 300:
        raise ReplayError("Timestamp outside window")

    # Reject if nonce already seen
    if nonce in seen_nonces:
        raise ReplayError("Duplicate nonce")

    seen_nonces.add(nonce)  # Store with TTL
    return verify_signature(msg, sig)
```

## Common Pitfalls

- Using timestamps alone without nonce (attacker replays within window)
- Not binding authentication to channel (token works across sessions)
- Accepting nonces without expiration (memory exhaustion)
- Ignoring replay in "secure" channels (TLS doesn't prevent application-layer replay)

## See Also

- Reflection attacks
- MITM attacks
- Challenge-response protocols

## References

- [RFC 4120 - Kerberos V5](https://datatracker.ietf.org/doc/html/rfc4120) (replay cache)
- [FIDO2/WebAuthn](https://www.w3.org/TR/webauthn/) (challenge-response)

---
*Written by Claude Opus 4.5*
