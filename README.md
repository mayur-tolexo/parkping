# parkping

ParkPing is a platform that helps people contact vehicle owners safely using a Fastag scan or vehicle number, without revealing personal phone numbers.

If a vehicle is wrongly parked or in an emergency:

- Scan the Fastag / enter vehicle number
- Check if the vehicle is registered
- Make a masked call or send a message to the owner

## Features

- REST APIs built with Go
- FASTag / QR Token–based vehicle identification
- Contact **Message** and **Call** handlers
- Swagger (OpenAPI) documentation
- Docker & Docker Compose support
- Separate **development** and **production** configurations
- Environment-based configuration (`APP_ENV`)

---

## Tech Stack

- **Language:** Go
- **Web Framework:** net/http (or gin, if applicable)
- **API Docs:** swaggo/swag (Swagger)
- **Containerization:** Docker
- **Orchestration:** Docker Compose

---
## Sample .env
```sh
APP_ENV=dev
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=appdb
DB_SSLMODE=disable
```
