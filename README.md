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