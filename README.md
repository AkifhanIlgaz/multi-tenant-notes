# Multi-Tenant Notes Demo

This project demonstrates three common multi-tenant architecture approaches used in SaaS (Software as a Service) applications. Each folder implements a different strategy for tenant data isolation:

- **separate-db-per-tenant**: Each tenant has a completely separate database. This provides strong isolation but can be harder to manage at scale.
- **shared-db-separate-schema**: All tenants share a single database, but each tenant has its own schema. This balances isolation and manageability.
- **shared-schema-and-db**: All tenants share the same database and schema. Tenant data is separated by a tenant identifier in each table. This is the simplest to manage but offers the least isolation.

You can explore each folder to see how these approaches are implemented in practice.
