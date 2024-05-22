CREATE TABLE transactions(
    id INT PRIMARY KEY,
    consumer_id INT not null ,
    contract_number varchar(40) not null,
    otr decimal(10,2) not null,
    fee_admin decimal(10, 2) not null ,
    installment_amount int not null,
    total_interest decimal(10, 2) not null,
    asset_name varchar(100) not null,
    transaction_date date not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY contract_number(contract_number),
    KEY asset_name(asset_name),
    KEY transaction_date(transaction_date),
    FOREIGN KEY (consumer_id) REFERENCES consumers(id)
);