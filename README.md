# GopherUp Tbilisi – March 2024

[![codecov](https://codecov.io/gh/itimky/gopher-up-march-2024/graph/badge.svg?token=GAFHEGYAUQ)](https://codecov.io/gh/itimky/gopher-up-march-2024)
[![Go Report Card](https://goreportcard.com/badge/github.com/itimky/gopher-up-march-2024)](https://goreportcard.com/report/github.com/itimky/gopher-up-march-2024)

<!-- https://mermaid.js.org/syntax/classDiagram.html -->

```mermaid
---
title: Orders
---
classDiagram
    direction LR
    namespace http-contract {
        class CreateOrderV1 {
        }
    }
    namespace http-handler {
        class Handler {
            HandleCreateOrderV1(..)
        }
        
        class HTTPDomainInterface["domain"] {
            <<interface>>
        }
    }
    namespace orders {
        class Orders {
            CreateOrder(..)
        }
        
        class OrdersModels["models"] {
        }
        
        class OrdersDBInterface["db"] {
            <<interface>>
        }
    }
    namespace orders-db {
        class InMemDB {
            AddOrder(..)
        }
    }
    namespace db-contract {
        class Order {
            
        }
    }

    CreateOrderV1 <-- Handler
    HTTPDomainInterface --> OrdersModels
    OrdersModels <-- InMemDB
    OrdersDBInterface <|.. InMemDB
    InMemDB --> Order
```
