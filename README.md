# 🚀 CommerceOps Platform

## Overview

CommerceOps Platform is an end-to-end DevOps project built using Go, Docker, Kubernetes, Helm, Terraform, GitHub Actions, Argo CD, Prometheus, and Grafana.

The project demonstrates modern DevOps practices including:

- CI/CD Automation
- GitOps Deployment
- Infrastructure as Code
- Containerization
- Kubernetes Orchestration
- Monitoring & Observability

---

## Architecture

```text
Developer
    │
    ▼
GitHub
    │
    ▼
GitHub Actions
    │
    ▼
Docker Hub
    │
    ▼
Argo CD
    │
    ▼
Kubernetes
    │
    ▼
Prometheus
    │
    ▼
Grafana
```

Terraform provisions Kubernetes resources through Infrastructure as Code.

---

## Tech Stack

| Category | Technology |
|-----------|------------|
| Language | Go |
| Containerization | Docker |
| Orchestration | Kubernetes |
| Package Management | Helm |
| Infrastructure as Code | Terraform |
| CI/CD | GitHub Actions |
| GitOps | Argo CD |
| Monitoring | Prometheus |
| Visualization | Grafana |
| Registry | Docker Hub |

---

## Features

- Go Web Application
- Docker Multi-Stage Build
- Kubernetes Deployment
- Helm Chart Packaging
- GitHub Actions CI/CD
- Argo CD GitOps Workflow
- Terraform Infrastructure Provisioning
- Prometheus Monitoring
- Grafana Dashboards
- Custom HTTP Metrics

---

## Project Structure

```text
ecommerce-devops-platform/
│
├── .github/workflows
├── app
├── helm
├── kubernetes
├── monitoring
├── terraform
└── README.md
```

---

## Monitoring

Application metrics are exposed through:

```text
/metrics
```

Prometheus scrapes metrics and Grafana visualizes:

- HTTP Requests
- CPU Usage
- Memory Usage
- Application Health

---

## Deployment Workflow

```text
Code Commit
     │
     ▼
GitHub
     │
     ▼
GitHub Actions
     │
     ▼
Docker Hub
     │
     ▼
Argo CD
     │
     ▼
Kubernetes
     │
     ▼
Prometheus
     │
     ▼
Grafana
```

---

## Author

### Tharun Kumaran

Aspiring DevOps Engineer

GitHub: (https://github.com/tharunkumaran05-ship-it)
