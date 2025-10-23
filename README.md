# Product Service
Product Service is a microservice responsible for managing product and category data. It is designed using an Event-Driven Architecture and integrates with a Search Engine for efficient product discovery.

---

## Futures
- Product Management
- Category Managemenent
- Search Engine

---

## Tech Stack
- **Programming Language:** Go
- **Database:** PostgreSQL
- **Logging:** Logrus
- **Validation:** Go Validator
- **API Documentation:** Swagger (via swaggo)
- **Containerization:** Docker
- **Deployment:** Kubernetes
- **Caching:** Memchached
- **Message Broker:** RabbitMQ
- **Search Engine:** Elasticsearch

---

## Setup
```bash
# clone repository
https://github.com/nurmanhadi/microservices-product-service.git

# go to directory
cd product-service

# install dependency
go mod tidy
```

---

## Environment Variable
Create file `.env`
```bash
DB_HOST=localhost
DB_PORT=5433
DB_DATABASE=product_service
DB_USERNAME=product_service
DB_PASSWORD=product_service
DB_MAX_POOL_CONNS=10
DB_MAX_IDLE_CONNS=5
DB_CONN_MAX_LIFETIME=300

BROKER_HOST=localhost
BROKER_PORT=5672
BROKER_USERNAME=admin
BROKER_PASSWORD=admin
BROCKER_VIRTUAL_HOSTS=ecommerce
BROCKER_QUEUE_PRODUCT_CONSUMER=product_consumer
```

---

## Usage

### Run Localy
```bash
go run main.go
```

---

## API Documentation
Access url Endpoint `/swagger/index.html`

---

## License
This project is licensed under the MIT License.

---

## Author
**Nurman Hadi**  
Backend Developer (Golang, Microservices)  
GitHub: [nurmanhadi](https://github.com/nurmanhadi)