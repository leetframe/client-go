# LeetFrame Go Client

### Quick Start
```
go mod tidy
go run scripts/testPublish.go
```

### Schema
```
client schema:
- events
  - id (auto increment/guid)
  - name (unique string)
  - location (string)
    - filename (string)
    - function name (string)
    - line number (int)
  - expectation (string)
  - metadata (json; structs:[], tags:[], stacktrace:"")
- event_occurences
  - event_id
  - timestamp
  - metadata (overriding attributes)
- projects
  - id (auto increment/guid)
  - name (unique string)
  - created_at
  - updated_at
  - tags (json, e.g. team:hermes)
```

# TODOs
- [ ] implement generic interface usable with pubsublite, pubsub, and kafka
- [ ] produce messages in local with client-go
