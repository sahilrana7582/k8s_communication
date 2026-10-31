# K8s Communication — SRE API

A Kubernetes management API built with **Golang, Gin, and Kubernetes Client-Go**.

The project provides REST APIs to communicate with a Kubernetes cluster and manage core Kubernetes resources such as Pods, Deployments, and Services.

The goal is to build a lightweight SRE backend that exposes Kubernetes operations through HTTP endpoints, making cluster inspection and resource management accessible to external applications and automation tools.

## Features

- **Kubernetes Client Integration** — Connect to Kubernetes clusters using Client-Go.
- **Pod Management** — Retrieve and inspect Pods.
- **Deployment Management** — Retrieve and inspect Deployments.
- **Service Management** — Retrieve and inspect Kubernetes Services.
- **Health Check** — Expose an API health endpoint.
- **REST API** — HTTP endpoints powered by Gin.
- **Layered Architecture** — Separate handlers, business logic, Kubernetes communication, and routing.

## Tech Stack

| Technology | Purpose |
|---|---|
| Go | Backend development |
| Gin | HTTP framework |
| Kubernetes Client-Go | Kubernetes API communication |
| Kubernetes | Container orchestration |

## Project Structure

```text
sre-api/
├── cmd/
│   └── main.go
│
├── internal/
│   ├── config/
│   │   └── config.go
│   │
│   ├── handler/
│   │   ├── health_handler.go
│   │   ├── pod_handler.go
│   │   ├── deployment_handler.go
│   │   └── service_handler.go
│   │
│   ├── k8s/
│   │   ├── client.go
│   │   ├── pod.go
│   │   ├── deployment.go
│   │   └── service.go
│   │
│   ├── service/
│   │   ├── pod_service.go
│   │   ├── deployment_service.go
│   │   └── service_service.go
│   │
│   └── routes/
│       ├── routes.go
│       ├── health_routes.go
│       ├── pod_routes.go
│       ├── deployment_routes.go
│       └── service_routes.go
│
├── go.mod
└── go.sum
```

## Architecture

The application follows a layered architecture.

```text
                Client
                  |
                  v
             Gin Router
                  |
                  v
              Handlers
                  |
                  v
               Services
                  |
                  v
          Kubernetes Client
                  |
                  v
          Kubernetes API Server
                  |
                  v
             K8s Cluster
```

### Layer Responsibilities

**1. Handler Layer (`internal/handler`)**

Responsible for:
- Receiving HTTP requests.
- Extracting URL parameters and query parameters.
- Validating request input.
- Calling the appropriate service.
- Returning HTTP responses.

**2. Service Layer (`internal/service`)**

Responsible for:
- Application-level business logic.
- Coordinating Kubernetes operations.
- Handling service-level errors.
- Preparing data for handlers.

**3. Kubernetes Layer (`internal/k8s`)**

Responsible for:
- Initializing Kubernetes clients.
- Communicating