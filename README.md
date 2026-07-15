# SeatLock

A cinema booking system written in Go that lets multiple users book tickets concurrently without double-booking a seat.

## Demo
<img width="3018" height="1532" alt="Adobe Express - Screen Recording 2026-07-15 at 9 17 53 AM" src="https://github.com/user-attachments/assets/cb210deb-dd3d-43c1-86c8-1e86709dfacd" />

## Problem Statement
How do we build a fast booking system that prevents double-bookings?

**Strategy 1: Synchronous Approach**
Users buy tickets in first-come-first-served order at a register. This prevents double-booking but is slow, since every purchase blocks the next.

**Strategy 2: Optimistic Concurrency**
Sell seats online with no lock while a purchase is in progress. This is faster than the synchronous approach, but two users can end up racing for the same seat — one of them enters their card details only to find the seat was taken out from under them.

**Strategy 3: Pessimistic Locking**
Sell seats online, but when a user starts checkout they're given a short-lived lock on the seat. This is the best of both worlds: it's fast, and no one wastes time entering payment details for a seat someone else just took.

SeatLock implements Strategy 3, using Redis as the lock store.

## How it works
- Picking a seat calls `HOLD`, which places a lock in Redis with a TTL (`SET NX` + expiry) so only one user can hold a given seat at a time.
- If the user confirms before the hold expires, the lock is persisted (TTL removed) and the booking becomes permanent.
- If the user cancels, or the TTL expires first, the seat is released automatically and becomes available again.

## API

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/movies` | List available movies |
| `GET` | `/movies/{movieID}/seats` | List seat status for a movie |
| `POST` | `/movies/{movieID}/seats/{seatID}/hold` | Place a timed hold on a seat |
| `PUT` | `/sessions/{sessionID}/confirm` | Confirm a held seat, making the booking permanent |
| `DELETE` | `/sessions/{sessionID}` | Release a held seat |

## Getting Started

**Requirements:** Go 1.24+, Docker

```bash
# start Redis (+ Redis Commander UI on :8081)
docker compose up -d

# run the server
go run ./cmd
```

The app is served at [http://localhost:8080](http://localhost:8080).

## Tech Stack
- Go (`net/http`, no framework)
- Redis for distributed seat locks
- Vanilla HTML/JS frontend (`static/`)
