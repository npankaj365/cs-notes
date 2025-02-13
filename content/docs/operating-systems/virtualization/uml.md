---
title: User-mode Linux
weight: 3
type: docs
---
# User-mode Linux (UML)

## What is it?
A virtualization system for Linux based on architectural port of the Linux kernel to its own system call interface, which enables multiple virtual Linux kernel-based OSes (guests) to run as an application within a normal Linux system (host). It is already part of the main kernel. 

> Linux has also been ported to a number of architectures without a PMMU, although functionality is then obviously somewhat limited. Linux has also been ported to itself. You can now run the kernel as a userspace application - this is called UserMode Linux (UML). [1]

## Why is it used?
- Enables you to boot a full Linux kernel inside of a user space. Think of it as VMWare but only for Linux. 

## Where is it used?
- Widely used for development and much easier to debug -- mostly kernel development
- Hosting virtual servers
- For setting up educational systems
- As a sandbox



## References

[1] [On what hardware does it run - Linux](https://www.kernel.org/doc/html/v6.13/admin-guide/README.html#on-what-hardware-does-it-run)