## Usage
**Docker Compose:**
```
  gps-tracker:
    restart: always
    logging:
      driver: "json-file"
      options:
        max-size: "200m"
    image: ghcr.io/doorbash/gps-tracker-backend:latest
```