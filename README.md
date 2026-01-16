# Go Gin – Production-Grade SRE Observability Project

This repository demonstrates a complete, production-ready SRE setup built around a Go Gin backend.  
It includes secure containerization, runtime hardening, metrics instrumentation, alerting, and a full observability stack.

---

## 📌 Features

### 1. **Backend (Go + Gin)**
- `/health` endpoint for liveness checks
- `/metrics` Prometheus endpoint
- Lightweight Gin router
- Graceful shutdown
- Zero dynamic dependencies (fully static binary)

---

### 2. **Containerization (Docker)**
- Multi-stage Dockerfile (builder → distroless runner)
- Non-root user
- Read-only file system
- Seccomp profile applied
- AppArmor (optional)
- CPU / Memory / PID limits
- Minimal attack surface

---

### 3. **Observability Stack**
Includes:

- **Prometheus** (metrics scraping)
- **Grafana** (dashboards)
- **cAdvisor** (container metrics)
- **Custom Prometheus alerts**

Dashboards include:
- Goroutines
- Backend CPU usage
- Backend memory usage
- Container CPU usage
- Container memory usage
- Container network activity

---

### 4. **Prometheus Alerting**
Alerts configured:

- **High CPU usage**  
- **High API latency (P90)**  
- Ready to integrate with Slack / Email via Alertmanager

---

## 📁 Project Structure


# Go Gin – Production-Grade SRE Observability Project

This repository demonstrates a complete, production-ready SRE setup built around a Go Gin backend.  
It includes secure containerization, runtime hardening, metrics instrumentation, alerting, and a full observability stack.

---

## 📌 Features

### 1. **Backend (Go + Gin)**
- `/health` endpoint for liveness checks
- `/metrics` Prometheus endpoint
- Lightweight Gin router
- Graceful shutdown
- Zero dynamic dependencies (fully static binary)

---

### 2. **Containerization (Docker)**
- Multi-stage Dockerfile (builder → distroless runner)
- Non-root user
- Read-only file system
- Seccomp profile applied
- AppArmor (optional)
- CPU / Memory / PID limits
- Minimal attack surface

---

### 3. **Observability Stack**
Includes:

- **Prometheus** (metrics scraping)
- **Grafana** (dashboards)
- **cAdvisor** (container metrics)
- **Custom Prometheus alerts**

Dashboards include:
- Goroutines
- Backend CPU usage
- Backend memory usage
- Container CPU usage
- Container memory usage
- Container network activity

---

### 4. **Prometheus Alerting**
Alerts configured:

- **High CPU usage**  
- **High API latency (P90)**  
- Ready to integrate with Slack / Email via Alertmanager

---

## 📁 Project Structure

# Go Gin – Production-Grade SRE Observability Project

This repository demonstrates a complete, production-ready SRE setup built around a Go Gin backend.  
It includes secure containerization, runtime hardening, metrics instrumentation, alerting, and a full observability stack.

---

## 📌 Features

### 1. **Backend (Go + Gin)**
- `/health` endpoint for liveness checks
- `/metrics` Prometheus endpoint
- Lightweight Gin router
- Graceful shutdown
- Zero dynamic dependencies (fully static binary)

---

### 2. **Containerization (Docker)**
- Multi-stage Dockerfile (builder → distroless runner)
- Non-root user
- Read-only file system
- Seccomp profile applied
- AppArmor (optional)
- CPU / Memory / PID limits
- Minimal attack surface

---

### 3. **Observability Stack**
Includes:

- **Prometheus** (metrics scraping)
- **Grafana** (dashboards)
- **cAdvisor** (container metrics)
- **Custom Prometheus alerts**

Dashboards include:
- Goroutines
- Backend CPU usage
- Backend memory usage
- Container CPU usage
- Container memory usage
- Container network activity

---

### 4. **Prometheus Alerting**
Alerts configured:

- **High CPU usage**  
- **High API latency (P90)**  
- Ready to integrate with Slack / Email via Alertmanager

---

## 📁 Project Structure

- # Go Gin – Production-Grade SRE Observability Project

This repository demonstrates a complete, production-ready SRE setup built around a Go Gin backend.  
It includes secure containerization, runtime hardening, metrics instrumentation, alerting, and a full observability stack.

---

## 📌 Features

### 1. **Backend (Go + Gin)**
- `/health` endpoint for liveness checks
- `/metrics` Prometheus endpoint
- Lightweight Gin router
- Graceful shutdown
- Zero dynamic dependencies (fully static binary)

---

### 2. **Containerization (Docker)**
- Multi-stage Dockerfile (builder → distroless runner)
- Non-root user
- Read-only file system
- Seccomp profile applied
- AppArmor (optional)
- CPU / Memory / PID limits
- Minimal attack surface

---

### 3. **Observability Stack**
Includes:

- **Prometheus** (metrics scraping)
- **Grafana** (dashboards)
- **cAdvisor** (container metrics)
- **Custom Prometheus alerts**

Dashboards include:
- Goroutines
- Backend CPU usage
- Backend memory usage
- Container CPU usage
- Container memory usage
- Container network activity

---

### 4. **Prometheus Alerting**
Alerts configured:

- **High CPU usage**  
- **High API latency (P90)**  
- Ready to integrate with Slack / Email via Alertmanager

---

## 📁 Project Structure

```go-gin-sre-observability/
│
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── Dockerfile
│   ├── seccomp.json
│   ├── .dockerignore
│   └── README.md
│
├── observability/
│   ├── prometheus/
│   │   ├── prometheus.yml
│   │   ├── alert-rules.yml
│   │   └── README.md
│   │
│   ├── grafana/
│   │   ├── dashboards/
│   │   │   └── gin-dashboard.json
│   │   └── provisioning/
│   │
│   └── cadvisor/
│
├── docker-compose.yml
├── .gitignore
└── README.md

```
---

## 🚀 How to Run

## 🚀 How to Run the Entire Stack

This project uses Docker Compose to start all services:

- Go Gin backend
- Prometheus
- Grafana
- cAdvisor

### 1. Clone the repository

```bash
git clone https://github.com/<your-username>/go-gin-sre-observability.git
cd go-gin-sre-observability

### 2. start all services

```bash
docker-compose up -d
```
#### ->this will automatically start :

- backend (http://localhost:8080)

- Prometheus (http://localhost:9090)

- Grafana (http://localhost:3000)

- cAdvisor (http://localhost:8085)
 
#### -->validate everything is running 

######backend:

```bash
http://localhost:8080/health
```
###### -->Metrics endpoints:

```bash
http://localhost:8080/metrics
```
###### -->Prometheous:

```bash
http://localhost:9090
```

###### -->Grafana:

```bash
http://localhost:3000
```
###### -->login:

```bash
username: admin
password: admin
```
#### 4.Import the Grafana Dashboard 

###### 1. Go to Grafana -->Dashboard --> Import
###### 2. Upload the file:

```bash
observability/grafana/dashboards/gin-dashboard.json
```

#### 5. STop the stack

```bash
docker-composse down
```
```yaml

---

# ⭐ OPTIONAL PROFESSIONAL ADD-ONS

If you want to make your README even stronger, you can add:

### “Run Only the Backend”

---
---

# Graphs

#### In this i have used 

- prometheus
- grafana 
- alertmanager
 
---> it monitor cpu spike , memory usage , threads etc..

###dashboards

```

![ Dashboard-1](/home/sanjay-vasudevan/Pictures/Screenshots
/Screenshot from 2026-01-16 20-45-58.png)

```
```
![ Dashboard-1](/home/sanjay-vasudevan/Pictures/Screenshots
/'Screenshot from 2026-01-16 20-46-05.png')

```
```
![ Dashboard-1](/home/sanjay-vasudevan/Pictures/Screenshots
/'Screenshot from 2026-01-16 20-46-12.png')
```
```
![ Dashboard-1](/home/sanjay-vasudevan/Pictures/Screenshots
/'Screenshot from 2026-01-16 20-46-33.png')

```
```
![ Dashboard-1](/home/sanjay-vasudevan/Pictures/Screenshots
/'Screenshot from 2026-01-16 20-46-40.png')
```

#### Thankyou for visiting

---

