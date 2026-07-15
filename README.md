# SeatLock
## SeatLock is a cinema booking system written in Go that allows multiple buyers to book tickets concurrently. 
## Architecture

```mermaid
flowchart LR
    Browser["🎟️ Browser UI<br/>(static/index.html)"]

    subgraph Server["Go HTTP Server :8080"]
        Mux["net/http mux<br/>(cmd/main.go)"]
        Handler["booking.Handler<br/>(internal/booking/handler.go)"]
        Service["booking.Service"]
        Store["BookingStore interface"]
        RedisStore["RedisStore"]
        MemStore["In-memory stores<br/>(tests / local)"]
    end

    Redis[("Redis<br/>(docker-compose)")]

    Browser -->|"GET /movies<br/>GET /movies/{id}/seats<br/>POST /movies/{id}/seats/{seatID}/hold<br/>PUT /sessions/{id}/confirm<br/>DELETE /sessions/{id}"| Mux
    Mux --> Handler --> Service --> Store
    Store --> RedisStore
    Store -.-> MemStore
    RedisStore -->|"seat:{movieID}:{seatID} → session JSON<br/>session:{sessionID} → seat key"| Redis
```

## Seat lifecycle

```mermaid
sequenceDiagram
    participant U as Buyer
    participant S as Server
    participant R as Redis

    U->>S: POST /movies/{movieID}/seats/{seatID}/hold
    S->>R: SET seat:{movieID}:{seatID} (TTL 2 min)
    R-->>S: OK (seat locked)
    S-->>U: 201 session_id + expires_at

    alt Buyer confirms in time
        U->>S: PUT /sessions/{sessionID}/confirm
        S->>R: PERSIST key (TTL removed)
        S-->>U: Seat booked permanently
    else Buyer releases
        U->>S: DELETE /sessions/{sessionID}
        S->>R: DEL seat + session keys
    else TTL expires
        R->>R: Key auto-expires
        Note over R: Seat is free again
    end
```
## Install & run
## Problem


Requires Go 1.24+ and Docker.

```bash
git clone https://github.com/tablecrasher/cinema-booking-system.git
cd cinema-booking-system

docker-compose up -d   # starts Redis
go run ./cmd            # starts the server

open http://localhost:8080
```
