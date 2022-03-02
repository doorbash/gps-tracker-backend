**Database:**
```
CREATE TABLE "Device" (
	"id"	TEXT NOT NULL,
	PRIMARY KEY("id")
);
CREATE TABLE "LatLng" (
	"device_id"	TEXT NOT NULL,
	"datetime"	INTEGER NOT NULL,
	"lat"	REAL NOT NULL,
	"lng"	REAL NOT NULL,
	"alt"	REAL NOT NULL,
	"hdop"	REAL NOT NULL,
	"pdop"	REAL NOT NULL,
	"vdop"	REAL NOT NULL,
	PRIMARY KEY("device_id","datetime")
);
```

```
INSERT INTO Device(id) VALUES("00000000000000000000");
```

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
