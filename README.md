# Archive of Being

[![Golang](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-v1.7.4-000000?logo=gin)](https://github.com/gin-gonic/gin)
[![Postgres](https://img.shields.io/badge/PostgreSQL-13-336791?logo=postgresql)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-6.2-CB3837?logo=redis)](https://redis.io/)
[![JWT](https://img.shields.io/badge/JWT-json_web_token-FE6C9A?logo=jwt)](https://jwt.io/)

Archive of Being is a secure digital journal where every personal moment is captured and preserved for eternity using cutting-edge technologies.

## Table of Contents
- [Introduction](#introduction)
- [Technologies](#technologies)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction
**Archive of Being** is a digital diaries designed to safely store your personal journey. Developed with performance and scalability in mind, this project leverages Golang, the Gin framework, PostgreSQL, Redis, and JWT to provide a secure and high-performance platform for preserving your life's data.

## Technologies
- **Golang:** The core language powering the application for robust performance and efficiency.
- **Gin:** A high-performance web framework that simplifies routing and middleware management.
- **Postgres:** A reliable relational database ensuring consistent and secure data storage.
- **Redis:** An in-memory data store for caching, accelerating data retrieval.
- **JWT:** JSON Web Tokens provide secure, stateless authentication for user sessions.

## Features
- **Digital Journal:** Log your daily experiences and important milestones.
- **Secure Data Storage:** Keep your data safe with PostgreSQL and secure JWT-based authentication.
- **Fast Performance:** Benefit from Redis caching for instant data access.
- **Scalable Architecture:** Easily extend the application as your repository of moments grows.
- **Philosophical Reflection:** A digital space that not only archives data but also inspires reflection on the journey of life.

## Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/yourusername/archive-of-being.git
    cd archive-of-being
    ```

2. **Set up environment variables:**

   Create a `.env` file in the root directory and configure your settings:
    ```env
    DATABASE_URL=postgres://username:password@localhost:5432/dbname
    REDIS_URL=redis://localhost:6379
    JWT_SECRET=your_secret_key
    ```

3. **Install dependencies:**
    ```bash
    go mod download
    ```

4. **Run database migrations:**
    ```bash
    go run cmd/migrate/main.go
    ```

5. **Start the application:**
    ```bash
    go run main.go
    ```

## Usage
After starting the application, access it at [http://localhost:8080](http://localhost:8080). Use the API endpoints to register, log in, and start adding your personal entries. For detailed API documentation, refer to the [docs](docs/API.md).

## Contributing
Contributions are welcome! Please check out our [contributing guidelines](CONTRIBUTING.md) before submitting a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements
Thank you for choosing Archive of Being. We hope this digital journal becomes a timeless repository of your life's journey.
