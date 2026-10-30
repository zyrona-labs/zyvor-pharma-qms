# Validation Approach

Zyvor Pharma QMS is built with validation in mind. This document tracks the intended validation strategy.

**Current status: not validated.**

## Planned activities

- **IQ (Installation Qualification)**: documented installation of the system in a controlled environment.
- **OQ (Operational Qualification)**: testing of core workflows (batch creation → sample → test result → QA approval → release).
- **PQ (Performance Qualification)**: end-to-end performance with representative data volumes and concurrent users.
- **Audit trail**: append-only, immutable event log covering create/update/approve/reject actions.
- **Electronic signatures**: role-based sign-off with meaning of signature recorded.
- **Access control**: least-privilege roles (Manufacturing, Laboratory, Quality, Release).

Validation protocols and evidence will be developed alongside the software. No claim of regulatory compliance is made until formal validation is completed and documented.
