# ValidateFX Terraform Provider – OS-Specific Installation & Troubleshooting

This document provides platform-specific installation and troubleshooting guidance for installing the **ValidateFX Terraform provider**.  
It focuses on Terraform plugin discovery behavior for **Terraform v1.5+** across **Windows**, **macOS**, and **Linux** systems.

---

## 🧩 Overview

Terraform automatically discovers provider binaries located in specific filesystem paths under the user’s home directory.  
When installing the ValidateFX provider manually (without `terraform init`), it’s important to place the binary correctly based on your OS and architecture.

The general directory layout follows this structure:

