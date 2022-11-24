1. GET
2. POST
3. REST API
4. Database persistance bukan in memmory
5. Autentikasi

===

1. API Framework => Echo
2. ORM => Gorm
3. Unit Test => BDD (Ginkgo)
4. Database => PostgreSQL
5. Use Case => CRUD Product

Product:
- ID [AI - INT - PK]
- Name [Str]
- Qty [INT]
- AuditTable [Extend]

Audit Table:
- CreatedAt
- CreatedBy
- UpdatedAt
- UpdatedBy
- DeletedAt
- DeletedBy

API

Usecase/Application

Domain

Infrastructre