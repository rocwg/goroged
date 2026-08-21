# ged

`ged` is a small, type-safe Go library for building Edge applications.

The project provides a small set of primitives for the boundary between HTTP consumers and backend services.

## Positioning

`ged` is an Edge Application library.

It is not:

- an API Gateway
- a reverse proxy
- a service mesh
- a workflow engine
- a business framework

`ged` intentionally stays small.

## Core Capabilities

### Request Context

Propagate request-level information through the application context.

```text
HTTP
 ↓
Request Context
 ↓
gRPC Metadata
```

Currently supported:

- Request ID
- Trace ID

### Identity Context

Carry authenticated identity information through the Edge request context.

Currently supported:

- User ID
- Tenant ID

### gRPC Metadata Propagation

Automatically propagate Edge context into outgoing gRPC metadata.

```
HTTP Request
    │
    ▼
ged/request
    │
    ▼
ged/identity
    │
    ▼
ged/grpc
    │
    ▼
gRPC Metadata
    │
    ▼
Provider
```

### Authentication Middleware

`ged` defines the authentication boundary through the `Authenticator` interface.

Authentication implementation remains outside the library.

### HTTP JSON Handler

`ged/http` provides a small generic helper for the common:

```
HTTP JSON
    ↓
Request
    ↓
Application/RPC
    ↓
Response
    ↓
HTTP JSON
```

flow.

## Design Principles

### 1. Small

`ged` should provide only capabilities that are repeatedly required by Go Edge applications.

### 2. Type-safe

Prefer Go types and interfaces over dynamic configuration and reflection.

### 3. Application-owned

`ged` provides primitives.

The application owns:

- routes
- business logic
- provider clients
- authentication implementations
- logging policy
- observability policy
- configuration
- error policy

### 4. Transport-aware, Business-unaware

`ged` understands Edge transport concerns such as:

- HTTP
- gRPC
- request context
- identity propagation

It does not understand business domains.

### 5. No Gateway Ambition

`ged` is not intended to become a replacement for KrakenD, Envoy, Kong, or a service mesh.

The goal is a small Go-native foundation for applications that need an Edge layer.

## Example

```go
router := gin.New()

router.Use(request.Middleware())
router.Use(
    authentication.Middleware(
        authenticator,
    ),
)
```

```go
mux := http.NewServeMux()

mux.HandleFunc("GET /api/v1/...", handler)
```

Outgoing gRPC calls automatically receive the request and identity metadata through the gRPC client interceptor.

## Module

```
github.com/rocwg/ged
```

## Status

```
ged v0.1
```

The v0.1 API is intentionally small and considered frozen.

这份 README 有一个很重要的作用：

> **防止未来的我们把 `ged` 越做越胖。** 

---
