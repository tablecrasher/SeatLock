# SeatLock
## SeatLock is a cinema booking system written in Go that allows multiple buyers to book tickets concurrently. 
### Demo:
<img width="3018" height="1532" alt="Adobe Express - Screen Recording 2026-07-15 at 12 22 48 AM" src="https://github.com/user-attachments/assets/a430ec00-7b38-4706-a23b-0524468bb047" />

## Install & run

Requires Go 1.24+ and Docker.

```bash
git clone https://github.com/tablecrasher/cinema-booking-system.git
cd cinema-booking-system

docker-compose up -d   # starts Redis
go run ./cmd            # starts the server

open http://localhost:8080
```
