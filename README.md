# 🚗 Parking Lot System

A multi-floor, concurrent parking lot system built with Go and Gin.

---

## ✅ Features

- 🧠 Configurable per-floor layout using `config.yaml`
- 📊 Multi-vehicle type support: Bicycle, Motorcycle, Automobile
- 🔄 Thread-safe `Park`, `Unpark`, `Search`, and `Available` APIs
- ⚙️ Clean architecture: config, service, model, and handler separation
- 🧾 Easily extendable with YAML-based setup

---

## 🏁 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/YOUR_USERNAME/parking-lot-system.git
cd parking-lot-system
```

### 2. Create `config.yaml`

```yaml
floors: 2
layout:
  - # Floor 0
    - [B-1, M-1, A-1]
    - [X-0, M-1, A-1]
  - # Floor 1
    - [A-1, A-1, M-1]
    - [B-1, X-0, M-1]
```

> `B-1` = Bicycle  
> `M-1` = Motorcycle  
> `A-1` = Automobile  
> `X-0` = Inactive

### 3. Run the application

```bash
go run main.go
```

> The app runs at `http://localhost:8080`

---

## 📡 API Endpoints

| Method | Endpoint                            | Description             |
|--------|-------------------------------------|-------------------------|
| POST   | `/api/v1/parkings`                  | Park a vehicle          |
| POST   | `/api/v1/parkings/unpark`           | Unpark a vehicle        |
| GET    | `/api/v1/parkings/available?type=A` | List available spots    |
| GET    | `/api/v1/parkings/search/:plate`    | Find last known spot    |

Example body for `POST /api/v1/parkings`:
```json
{
  "type": "A",
  "vehicleNumber": "AB1234CD"
}
```

---

## 🛠 Project Structure

```
parking-lot/
├── main.go               # Server entry point
├── config/               # YAML config loader
├── model/                # Spot & vehicle models
├── service/              # Parking logic & concurrency
├── handler/              # HTTP handlers
└── config.yaml           # Floor layout config
```

---

## 📋 Requirements

- Go 1.18+
- Gin Web Framework
- gopkg.in/yaml.v3

---

## ⚖️ License

MIT © tamamsalimi  
