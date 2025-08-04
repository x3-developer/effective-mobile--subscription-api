CREATE TABLE subscriptions
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    price      BIGINT       NOT NULL,
    user_id    UUID         NOT NULL,
    start_date TIMESTAMP    NOT NULL,
    end_date   TIMESTAMP
);