Student-Teacher Notification Service (Go + Hexagonal Architecture)

Welcome! 🎉
This repository is designed as a learning project for engineers new to Go (especially those coming from Java/Spring Boot).

We’ll build up a service step by step:

CRUD with Postgres

Metrics with Prometheus

Adding new domains (Teacher)

Caching with Redis

Event-driven messaging with Kafka

Each step is captured in a separate branch, so you can explore the journey incrementally.

🚀 Branches & Learning Path

main → Basic CRUD for Student

Postgres persistence

REST APIs (GET, POST, PUT, DELETE)

Hexagonal layering (Models, Ports, Repositories, Services, Controllers, Routes)

prometheus-metrics → Observability

Added Prometheus endpoint (/metrics)

Custom metrics for API requests

Middleware for instrumentation

teacher-domain → Extending the domain

Added a Teacher model, repository, service, and controller

Demonstrated how multiple domains co-exist in the same hexagonal structure

redis-cache → Performance

Integrated Redis for student caching

Added Cache Hit / Cache Miss metrics

Hybrid repository pattern (DB + Redis)

kafka-events → Event-driven architecture

Teachers can send messages to students

Kafka producer publishes notification events

Kafka consumer listens, processes, and persists to Postgres (student_teacher_notifications table)
