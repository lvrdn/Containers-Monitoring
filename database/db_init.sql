DROP TABLE IF EXISTS "containers";
CREATE TABLE containers (
    "address" varchar(100) PRIMARY KEY,
    "last_ping" timestamp,
    "last_success_ping"  timestamp
);