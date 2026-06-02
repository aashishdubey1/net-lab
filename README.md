# net-lab: Networking in Go — From TCP to WebSockets

This repo documents my hands-on journey through backend networking fundamentals, built entirely in Go. Each folder represents a day's dedicated project

| Day | Project                                         | Description                                                                                                                          |
| --- | ----------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| 1   | [tcp-echo-server](./day-01-tcp-echo)            | A raw TCP server that echoes back any message sent by a client. Multiple clients handled concurrently via goroutines.                |
| 2   | [tcp-chat-server](./day-02-tcp-chat)            | Multi-client broadcast chat over TCP. Messages sent by one client are relayed to all connected clients via a central hub.            |
| 3   | [udp-ping-server](./day-03-udp-ping)            | Stateless UDP ping/pong server and client. Measures round-trip time and demonstrates tradeoffs between UDP and TCP.                  |
| 4   | [http-from-scratch](./day-04-http-from-scratch) | HTTP/1.1 server built on top of raw TCP — no net/http. Manually parses request lines, headers, and writes compliant responses.       |
| 5   | [rest-api-server](./day-05-rest-api)            | CRUD REST API using net/http with custom routing, middleware chain for logging and timing, and JSON request/response handling.       |
| 6   | [reverse-proxy](./day-06-reverse-proxy)         | HTTP reverse proxy with round-robin load balancing across multiple upstream Go servers and automatic health check failover.          |
| 7   | [websocket-server](./day-07-websocket)          | Realtime WebSocket chat server with a browser UI. Upgrades HTTP connections and handles graceful shutdown with context cancellation. |
