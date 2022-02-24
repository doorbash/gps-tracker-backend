**Docker Compose:**
```
  gps-tracker:
    restart: always
    logging:
      driver: "json-file"
      options:
        max-size: "200m"
    volumes:
      - ./docker/gps-tracker/db.sqlite:/db.sqlite
    image: ghcr.io/doorbash/gps-tracker-backend:latest
```