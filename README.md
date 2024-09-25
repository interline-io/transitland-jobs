# Transitland Jobs

This package contains interfaces for a Job Queue, and implementations for a local in-memory queue (`local`), a Redis-backed queue (`redis`), and a PostgreSQL backed queue (`river`). Each implementation may be wrapped using `jobs/JobLogger` to add detailed logging.