# CI CI Pipeline

| Job | Runs On | Needs |
|-----|---------|-------|
| build | ubuntu-latest |  |
| test | ubuntu-latest | build |

## Diagram
```mermaid
flowchart TD
    build[build]
    test[test]
    build --> test

```
