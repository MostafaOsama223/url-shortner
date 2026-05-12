## URL Shortner
### Introduction
---
This repository implements the URL Shortner discussed on Chapter 8 in Alex Xu's book "System Design Interview". The purpose of this project is to put the system under heavy tests/loads (concurrent requests, heavy traffic, etc...), given limited server resources.

### Hardware
---
Raspberry pi 4 model B will be used as the host.

**Specs:**
- RAM: 4 GB LPDDR4-3200
- CPU: Quad core Cortex-A72 (ARM v8) 1.8GHz
- Storage: 16 GB

*It is intentional to use a limited device to be a motive for implement the optimal solution. A VM could also be used, but I've a Raspberry :)*