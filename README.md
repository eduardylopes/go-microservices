# Microservices Project

This repository contains a set of microservices designed to perform various tasks within a system.

## Microservices Overview

![Microservices Interaction](https://github.com/eduardylopes/go-microservices/assets/60992933/f3c1ef9e-113e-4027-84e1-9e41f5483976)

- **authentication-service:** Responsible for authenticating users.
- **broker-service:** Intercepts all requests and redirects them to the provided service.
- **listener-service:** Responsible for listening to a queue.
- **logger-service:** Responsible for logging events.
- **mailer-service:** Responsible for sending emails.

## Getting Started

To run the project, follow the steps below:

### Prerequisites

Make sure you have the following installed:

- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://linuxhint.com/install-make-ubuntu/)

### Install dependencies

To ensure the proper functioning of each microservice, it is essential to install their dependencies. Fortunately, there is a convenient script to install all dependencies at once.

```bash
make install_dependencies
```

### Build and Run

To build and start the project, use the following command:

```bash
make up_build
```

### Database Migration

To synchronize migrations and create a table named 'users' with a dummy user, use the following command:

```bash
make migrateup
```

### Accessing Interfaces

- To test the microservices, go to http://localhost:3000.
- To view the email interface, access http://localhost:8025.
- To access the MongoDB interface, visit http://localhost:8082 using the login credentials: admin / password.
