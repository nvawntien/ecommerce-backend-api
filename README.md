#  Go E-commerce Backend API

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Framework](https://img.shields.io/badge/Gin-Framework-000000?style=flat&logo=gin)
![Database](https://img.shields.io/badge/PostgreSQL-14+-336791?style=flat&logo=postgresql)
![Redis](https://img.shields.io/badge/Redis-Cache%20%26%20OTP-DC382D?style=flat&logo=redis&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)

A robust, scalable, and secure RESTful API for an E-commerce platform built with **Golang**. This project adheres to **Clean Architecture** principles and utilizes modern tools like **Redis** for caching and **Viper** for configuration management to ensure production readiness.

##  Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Architecture](#-architecture)
- [Database & Caching](#-database--caching)
- [Getting Started](#-getting-started)
- [API Documentation](#-api-documentation)
- [Project Structure](#-project-structure)

##  Features

- **Authentication & Security**
    - Secure user registration and login.
    - **OTP (One-Time Password)** verification stored in **Redis** with TTL (Time-To-Live).
    - JWT implementation for stateless authentication.
    - Password hashing using `bcrypt`.

- **Product Management**
    - Complete CRUD operations for products and variants.
    - Optimized partial updates using pointers.
    - Inventory management with concurrency control.

- **Order Processing (Transactional)**
    - **ACID Compliance:** Uses database transactions for data integrity.
    - Server-side price calculation and stock deduction.
    - Order history and status tracking.

- **System & Ops**
    - **Centralized Configuration:** Managed via **Viper** (supports `.env`, `.yaml`, flags).
    - **Log Rotation:** Implemented with **Lumberjack** to prevent disk overflow in production environments.

##  Tech Stack

- **Core:** [Go (Golang)](https://go.dev/)
- **Web Framework:** [Gin Gonic](https://gin-gonic.com/)
- **Database:** PostgreSQL (Primary), Redis (Caching/Transient).
- **Configuration:** [Viper](https://github.com/spf13/viper) - For 12-factor app configuration.
- **Logging:** [Lumberjack](https://github.com/natefinch/lumberjack) - For rolling log files.
- **Data Access:** `sqlx` & `lib/pq`.
- **Security:** `golang-jwt/jwt/v5`, `crypto/bcrypt`.
- **Utilities:** `google/uuid`.

##  Architecture

The project follows **Clean Architecture** to decouple business logic from infrastructure.



- **Controller Layer:** Handles HTTP requests and validation.
- **Service Layer:** Core business logic (e.g., calculating totals, verifying OTPs).
- **Repository Layer:** Abstracted data access for PostgreSQL and Redis.
- **Models:** DTOs and Database Entities.

##  Database & Caching

### PostgreSQL (Primary)
- `users`, `products`, `product_variants`, `orders`, `reviews`.
- Uses `pgcrypto` for UUID generation.

### Redis (Transient Store)
- **OTP Storage:** Stores authentication codes with expiration (e.g., 5 minutes).
- **Session/Cache:** (Optional) Caching hot product data for low latency.

##  Getting Started

### Prerequisites
- Go 1.21+
- PostgreSQL
- **Redis** (Required for OTP features)
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone [https://github.com/nvawntien/go-ecommerce-backend.git](https://github.com/nvawntien/go-ecommerce-backend.git)
   cd go-ecommerce-backend