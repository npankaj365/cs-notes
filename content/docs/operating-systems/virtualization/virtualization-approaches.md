---
title: Virtualization Approaches
weight: 2
type: docs
---
# Virtualization Approaches

| Full Virtualization     | Paravirtualization      |
| ------------- | ------------- |
| Complete Hardware Emulation | Does not fully emulate hardware |
| Requires no changes to the guest OS | Requires the guest OS to be modified |
|Performance without Hardware-assisted Virtualization (like Intel VT-x or AMD-V) suffers due to binary translation for handling privileged instructions. | Can offer better performance without hardware assistance as there is lesser overhead of translating instructions.|
|Compatible with a wider range of OS, as it requires no OS modifications | Compatibility is limited.|
|Hypervisor intercepts and translates all instructions from the guest OS to the physical hardware. This includes privileged instructions that would normally interact directly with hardware. |  Hypervisor provies an interface for the guest oS to use for privileged operations, memory management, and I/O operations. |



Full Virtualization wasn't fully available on x86 before 2005, but is more commonly used lately, since most modern hardware has virtualization hardware support built-in. Paravirtualization, on the other hand, is not widely used as it used to before. Instead, **Hybrid virtualization** combines full virtualization techniques with paravirtualized drivers to overcome limitations with hardware-assisted full virtualization. Specifically, it overcomes the high CPU overheads coming from the many VM traps that get called while using an unmodified guest OS in a hardware-assisted full virtualization approach. 
