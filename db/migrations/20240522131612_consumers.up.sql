CREATE TABLE consumers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    full_name varchar(255) NOT NULL,
    nik varchar(16) NOT NULL,
    legal_name varchar(50) NOT NULL,
    pob varchar(50) NOT NULL,
    dob date NOT NULL ,
    salary decimal(10,2) NOT NULL,
    image_identity text not null,
    image_selfie text not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY full_name(full_name),
    KEY nik(nik),
    KEY legal_name(legal_name),
    KEY pob(pob),
    KEY dob(dob),
    KEY salary(salary)
);