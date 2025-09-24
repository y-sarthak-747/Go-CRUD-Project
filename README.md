# 📘 Student-Teacher Notification Service (Go + Hexagonal Architecture)

Welcome! 🎉  
This repository is a **learning project** for engineers new to Go (especially those from Java/Spring Boot).  

It builds up step by step:

- ✅ Basic CRUD with Postgres  
- 📊 Observability with Prometheus  
- 👩‍🏫 Adding Teacher domain  
- ⚡ Caching with Redis  
- 📩 Event-driven messaging with Kafka  

Each step is captured in its own **branch**, so you can follow the journey incrementally.  

---

## 📚 Before Starting This Project — Recommended Resources

If you’re new to Go, it’s best to get comfortable with the basics first.  
I highly recommend going through the following resources and completing these specific chapters **before diving into this project**:

- 🎥 **YouTube (Go Crash Course)**: [Go Crash Course](https://www.youtube.com/watch?v=un6ZyFkqFKo&t=18089s)  
  - Great for an overview and hands-on examples.  

- 📘 **Boot.dev (Learn Golang)**: [Learn Golang](https://boot.dev/learn/learn-golang)  
  - Make sure to complete at least these chapters:  
    1. Variables  
    2. Conditionals  
    3. Functions  
    4. Structs  
    5. Interfaces  
    6. Errors  
    7. Pointers (Chapter 10)  
    8. Packages & Modules (Chapter 11)  

👉 Once you’ve gone through these topics, you’ll be ready to understand and work with this project effectively.  


---


## 🪜 Learning Path (Branches)

- **`main`**  
  - Basic CRUD for **Student**  
  - Postgres persistence  
  - REST APIs (`GET`, `POST`, `PUT`, `DELETE`)  
  - Hexagonal layering (`models`, `ports`, `repositories`, `services`, `controllers`, `routes`)  

- **`prom`**  
  - Added **Prometheus endpoint** (`/metrics`)  
  - Middleware for request instrumentation  
  - Custom metrics  

- **`teacher_domain`**  
  - Added **Teacher** model, repository, service, and controller  
  - Showcases how multiple domains co-exist in hexagonal structure  

- **`redis`**  
  - Integrated **Redis caching** for students  
  - Added **Cache Hit / Cache Miss** metrics  
  - Hybrid repository pattern (DB + Redis)  

- **`kafka`**  
  - Teachers can send messages to students  
  - **Kafka Producer** publishes notification events  
  - **Kafka Consumer** listens, processes, and saves to Postgres (`student_teacher_notifications`)  


---

# 🔷 Hexagonal Architecture (Beginner-Friendly Explanation)

Hexagonal Architecture (also called **Ports and Adapters**) is a way of organizing code so that your **business logic** is at the center, and all the “outside world stuff” (like databases, HTTP APIs, caches, Kafka) just **plug into it**.  

Think of it like a hexagon socket wrench — the inside shape is always the same, but you can attach different heads (adapters) depending on what you need.  

---

## Why Hexagonal? 🤔

- **Keep business logic clean** → No `gorm`, `redis`, `kafka` code in your core domain.  
- **Easier to test** → Replace real database with a mock repo in tests.  
- **Swapable infrastructure** → Move from Postgres → MySQL or Redis → Memcached without changing business rules.  
- **Clear separation of concerns** → Each layer has its own job.  

---

## Layers in This Project

- **Domain Layer**  
  - Contains entities like `Student`, `Teacher`, `Notification`.  
  - Knows nothing about databases, HTTP, or Kafka.  
  - Example: `Student` struct has `ID`, `Name`, `Number`.  

- **Ports (Interfaces)**  
  - Define contracts the domain needs (like `StudentRepository` or `NotificationRepository`).  
  - Example:  
    ```go
    type NotificationRepository interface {
        Save(notification *Notification) error
    }
    ```  

- **Application Layer (Services)**  
  - Implements business use-cases by depending on **ports**.  
  - Example: `NotificationService` takes an event and tells the repository to save it.  
  - It doesn’t care if the repo is Postgres, Redis, or something else.  

- **Infrastructure Layer (Adapters)**  
  - Actual implementations that talk to the outside world.  
  - Examples:  
    - Postgres repo → saves data in DB.  
    - Redis repo → caches students.  
    - Kafka producer/consumer → sends and receives events.  
    - HTTP controllers → expose REST APIs.  

- **Config & Bootstrap**  
  - Wires everything together when the app starts.  
  - Example: `main.go` calls `config.ConnectDatabase()`, `InitKafkaProducer()`, then starts routes and consumers.  

---

## Example Flow: Teacher Sends a Notification

1. Teacher calls the API → `POST /teachers/send`.  
2. Controller receives request, turns it into a `NotificationEvent`.  
3. Kafka Producer sends it to the topic `student_notifications`.  
4. Kafka Consumer listens, deserializes the event, and hands it to `NotificationService`.  
5. `NotificationService` maps event to `Notification` entity.  
6. Calls `NotificationRepository.Save(notification)` (port).  
7. Postgres repository (adapter) inserts into `student_teacher_notifications` table.  

---


## 💡 Beginner Tips

- Start at **`domain/models`** → see what entities exist.  
- Look at **`ports`** → understand the contracts.  
- Check **`services`** → see how business rules are applied.  
- Then explore **controllers** and **kafka** → to see how the outside world interacts.  
- If you’re stuck or curious → **just ask ChatGPT**!  
  - I personally used ChatGPT extensively while building this project.  
  - It helped me understand tricky Go concepts, hexagonal architecture details, and best practices.  
  - Don’t hesitate — treat it like your coding buddy. 🚀  


---


## ⚙️ Project Setup

To run this project, you need Go installed (1.20+ recommended).  
Use the following commands during development:

- Run the application: `go run main.go`  
- Build the application: `go build main.go`  
- Download dependencies: `go mod tidy`  
- Run a single file (example): `go run hello.go`  

Always run `go mod tidy` after adding new dependencies so your `go.mod` and `go.sum` stay clean.

---

## 📦 Dependencies

These are the main dependencies required for this project:

- Gin (HTTP framework): `go get github.com/gin-gonic/gin`  
- GORM (ORM for Postgres): `go get gorm.io/gorm` and `go get gorm.io/driver/postgres`  
- Go-Redis (Redis client): `go get github.com/redis/go-redis/v9`  
- Prometheus client for Go: `go get github.com/prometheus/client_golang/prometheus` and `go get github.com/prometheus/client_golang/prometheus/promhttp`  
- Sarama (Kafka client for Go): `go get github.com/IBM/sarama`  

---

## 📡 APIs (curl examples)

1. Create Student  
   curl -X POST http://localhost:8080/students \
   -H "Content-Type: application/json" \
   -d '{"name":"Alice","number":"1234"}'

2. Get Student by ID  
   curl -X GET http://localhost:8080/students/1

3. Create Teacher  
   curl -X POST http://localhost:8080/teachers \
   -H "Content-Type: application/json" \
   -d '{"name":"Mr. Smith","subject":"Math"}'

4. Get Teacher by ID  
   curl -X GET http://localhost:8080/teachers/1

5. Teacher Sends Notification to Student  
   curl -X POST http://localhost:8080/teachers/send \
   -H "Content-Type: application/json" \
   -d '{"student_id": 1, "teacher_id": 2, "message": "Don’t forget your math homework!"}'


---


## 🗄️ Database Schemas

The following tables are created in Postgres:

1. **students**  
   - id (serial primary key)  
   - name (text)  
   - number (text)  

2. **teachers**  
   - id (serial primary key)  
   - name (text)  
   - subject (text)  

3. **student_teacher_notifications**  
   - id (serial primary key)  
   - student_id (int, foreign key to students.id)  
   - teacher_id (int, foreign key to teachers.id)  
   - message (text)  
   - created_at (timestamp)  

### SQL Queries to Create Tables

CREATE TABLE students (  
    id SERIAL PRIMARY KEY,  
    name TEXT NOT NULL,  
    number TEXT NOT NULL  
);

CREATE TABLE teachers (  
    id SERIAL PRIMARY KEY,  
    name TEXT NOT NULL,  
    subject TEXT NOT NULL  
);

CREATE TABLE student_teacher_notifications (  
    id SERIAL PRIMARY KEY,  
    student_id INT NOT NULL REFERENCES students(id),  
    teacher_id INT NOT NULL REFERENCES teachers(id),  
    message TEXT NOT NULL,  
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  
);

## 📩 Kafka

- Kafka Topic used: student_notifications  
- Teachers publish messages to this topic.  
- The Kafka consumer listens to this topic, processes events, and saves them into the student_teacher_notifications table.


---


## 🐳 Running with Docker Compose (Optional)

This project also includes a docker-compose.yml file that sets up all the required infrastructure services:

- Postgres (for Student/Teacher/Notification tables)  
- Redis (for caching)  
- Kafka + Zookeeper (for event-driven messaging)  
- Prometheus (for metrics scraping)  

If you don’t want to install these tools manually on your system, you can simply run:

docker compose up -d

This will start all the containers in the background. Your Go service can then connect to Postgres, Redis, and Kafka running inside Docker. Prometheus will also be available to scrape metrics.

👉 This step is optional. If you already have Postgres, Redis, or Kafka installed locally, you can use your own setup. If you don’t, docker compose up is the quickest way to get started.
