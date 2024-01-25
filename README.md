# GO Clean Architecture Code Base

In this project are already included an integration with database, Currently it's available for (Mysql, PostgreSQL).
It's might be usefull to know that libraries that I've used on this Project:
- Gorm -> ORM
- GoMock -> Test
- DotEnv -> Environment Variables
- zerolog -> Logging
- Goose -> Migration

## Prerequisites

- Go installed on your machine

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/elhardian/go-clean-architecture.git
   ```

2. Go to Project Directory:
    ```bash
    cd go-clean-architecture
    ```

3. Setup The Project 
    Setup project is includes `install dependencies` and `run migrations`
    ```bash
    make setup
    ```

## Usage

### Run the Application
To run the application, use the following command:
```bash
make run
```