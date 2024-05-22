CREATE TABLE consumer_limits (
    id INT PRIMARY KEY,
    consumer_id int not null,
    tenor int not null ,
    limit_amount decimal(10,2) not null default 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY tenor(tenor),
    KEY limit_amount(limit_amount),
    FOREIGN KEY (consumer_id) REFERENCES consumers(id)
);