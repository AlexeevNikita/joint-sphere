CREATE TABLE "user"
(
    id              BIGSERIAL PRIMARY KEY,
    email           VARCHAR(320) UNIQUE NOT NULL,
    username        VARCHAR(50)         NOT NULL,
    surname         VARCHAR(50)         NOT NULL,
    name            VARCHAR(50)         NOT NULL,
    patronymic      VARCHAR(50),
    country         VARCHAR(50),
    birthdate       DATE,
    additional_info TEXT
);

CREATE TABLE "subscription"
(
    id          BIGSERIAL PRIMARY KEY,
    producer_id BIGINT NOT NULL,
    consumer_id BIGINT NOT NULL,
    FOREIGN KEY (producer_id) REFERENCES public.user(id),
    FOREIGN KEY (consumer_id) REFERENCES public.user(id)
);

CREATE INDEX idx_producer_id ON Subscription (producer_id);
CREATE INDEX idx_consumer_id ON Subscription (consumer_id);