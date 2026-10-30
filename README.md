# Zyvor Pharma QMS

Pharmaceutical manufacturing and quality management system.

**Production batch → laboratory testing → quality approval → distribution → recall tracking.**

Focused on small and midsize pharmaceutical manufacturers, contract manufacturers, and quality laboratories. Distinct from [Zyvora Pharmacy](https://zyvorai.github.io/zyvor-pharma-site/) (medical stores).

## Status

Early development. Core domain models and UI skeleton in progress.  
**Not a validated system.** Validation documentation is built alongside the software. Deployment readiness does not establish regulatory compliance (e.g. 21 CFR Part 11, EU GMP Annex 11, WHO).

## Stack

- **Frontend**: React 19 + TypeScript + Vite + Tailwind CSS — Apple-inspired Liquid Glass aesthetic (top navigation, command palette, glassmorphism, dark-first).
- **Backend**: Go (chi, pgx, sqlc) with append-only audit trails and electronic signature hooks.
- **Database**: PostgreSQL (strong consistency, audit-friendly).
- **Infrastructure**: Kubernetes + Envoy (API gateway). Local development via Docker Compose.
- **AI**: Source-linked SOP answers and quality-report drafts that require human review before any action.

## Modules (initial scope)

| Module | Scope |
|--------|-------|
| Manufacturing | Batch records, recipes, material consumption, yield |
| Quality | Deviations, CAPA, change control, SOP approvals |
| Laboratory | Samples, test results, specifications, certificates of analysis |
| Traceability | Supplier lots, batch genealogy, quarantine, recalls |
| Equipment | Calibration, cleaning, maintenance records |
| Governance | Roles, audit trails, electronic sign-offs, training records |
| AI assistant | Source-linked SOP answers and quality-report drafts requiring review |

## Quick start (local)

```bash
# Backend
cd backend && go run ./cmd/server

# Frontend
cd frontend && npm install && npm run dev
```

See `docker-compose.yml` for full stack (Postgres + backend + frontend).

## License

Apache-2.0. See [LICENSE](LICENSE).

Part of the Zyvor suite. Own repository, domain model, and release cycle.
