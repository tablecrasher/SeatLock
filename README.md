# SeatLock

A cinema seat-booking backend that guarantees no two people ever get sold the same seat, even under heavy concurrent traffic.

## Problem statement

When many users try to book the same seat at once, a naive "check if free, then book" flow lets two requests both see the seat as free before either writes — resulting in double bookings. SeatLock exists to make that impossible.

## Approaches

**1. Synchronous (FIFO)** — Requests handled one at a time against a plain in-memory map. Safe only because there's no real concurrency; doesn't scale.

**2. Pessimistic locking** — A mutex locks the store for every check-then-book. Correct, but the lock is local to one process, so it breaks down as soon as the app runs on more than one instance.

**3. Optimistic (no locking)** — Check-and-book collapses into a single atomic write that just fails if someone else got there first. No lock held, no blocking.

**Final solution** — Built on the optimistic approach using Redis's atomic `SET NX` with a TTL: booking a seat is one atomic write, holds auto-expire if a checkout is abandoned, and correctness holds across any number of app instances. Verified with a test that fires 100,000 concurrent requests at one seat and confirms exactly one wins ([service_test.go](internal/booking/service_test.go)).

## Demo

<!-- Add a screenshot or GIF of the seat map here -->

Pick a movie, click a seat to hold it for 2 minutes with a live countdown, then confirm or release. Other users see holds and bookings update in near real time.

## Install & run

Requires Go 1.24+ and Docker.

```bash
git clone https://github.com/tablecrasher/cinema-booking-system.git
cd cinema-booking-system

docker-compose up -d   # starts Redis
go run ./cmd            # starts the server

open http://localhost:8080
```
