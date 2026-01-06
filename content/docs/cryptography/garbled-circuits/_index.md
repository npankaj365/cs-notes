---
weight: 1
bookFlatSection: false
title: "Garbled Circuits"
---

# Garbled Circuits

## Policy Circuits
Before we dive into garbled circuits, let us understand policy circuits in general. 

Two types:
- Boolean Circuits
- General Circuits

The following is an example of Boolean Circuits. 

{{< code language="go" source="/extras/boolean-circuits-toy.go" >}}

## Yao's Garbled Circuits

Garbled circuits enable two-party secure computation where parties learn only the output, not each other's inputs.

### Protocol Overview

1. **Garbler** (Alice):
   - Assigns two random labels (k⁰, k¹) to each wire (representing 0 and 1)
   - For each gate, encrypts output labels under input label combinations
   - Sends garbled circuit + her input labels to Bob

2. **Evaluator** (Bob):
   - Receives garbled circuit and Alice's input labels
   - Obtains his own input labels via Oblivious Transfer
   - Evaluates gate-by-gate: decrypts exactly one output label per gate
   - Decodes final output using output mapping

### Gate Garbling

For AND gate with input wires A, B and output wire C:

| A | B | C=A∧B | Encrypted Entry |
|---|---|-------|-----------------|
| 0 | 0 | 0 | Enc(kᴬ⁰, kᴮ⁰, kᶜ⁰) |
| 0 | 1 | 0 | Enc(kᴬ⁰, kᴮ¹, kᶜ⁰) |
| 1 | 0 | 0 | Enc(kᴬ¹, kᴮ⁰, kᶜ⁰) |
| 1 | 1 | 1 | Enc(kᴬ¹, kᴮ¹, kᶜ¹) |

Rows are randomly permuted so evaluator can't infer which row decrypts.

### Optimizations

| Technique | Improvement |
|-----------|-------------|
| Point-and-permute | O(1) decryption attempts instead of O(4) |
| Free XOR | XOR gates require no garbling/communication |
| Half-gates | AND gates need only 2 ciphertexts (was 4) |
| Row reduction | Derive one entry, send 3 instead of 4 |

### Security Properties

- **Privacy**: Evaluator sees only one label per wire (learns nothing about other input)
- **One-time**: Same garbled circuit cannot be reused (labels would leak)
- **Malicious security**: Requires cut-and-choose or authenticated garbling

### Applications

- Private set intersection
- Secure auctions
- Privacy-preserving ML inference
- Two-party key generation

## When to Use

- Two-party computation with semi-honest or malicious adversaries
- Function is known at compile time (circuit can be pre-generated)
- Low round complexity matters (constant rounds)
- Boolean or arithmetic circuits with moderate depth
- One-shot computation (fresh garbling each time is acceptable)

## When Not to Use

- More than two parties (use secret sharing-based MPC)
- Reactive/stateful computation (garbled circuits are one-shot)
- Very deep circuits (communication scales with circuit size)
- Function depends on runtime data (circuit structure must be fixed)
- Symmetric computation where both parties contribute equally (garbler does more work)


---
*Written by Claude Opus 4.5*
