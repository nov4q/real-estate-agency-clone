CREATE TABLE IF NOT EXISTS Tenants
(
    tenantId INTEGER NOT NULL AUTO_INCREMENT,
    firstName VARCHAR(30) NOT NULL,
    lastName VARCHAR(30) NOT NULL,
    userName VARCHAR(30) NOT NULL,
    password CHAR(128) NOT NULL, -- 128 characters long password stands for password encrypted with SHA-512
    PRIMARY KEY(tenantId)
);

CREATE TABLE IF NOT EXISTS Apartments
(
    apartmentId INTEGER AUTO_INCREMENT NOT NULL,
    city VARCHAR(100) NOT NULL,
    address VARCHAR(100) NOT NULL,
    area REAL(10,3) NOT NULL,
    tenantId INTEGER,
    rentPrice REAL(10, 2) NOT NULL,
    PRIMARY KEY(apartmentId),
    FOREIGN KEY(tenantId) REFERENCES Tenants(tenantId)
);

CREATE TABLE IF NOT EXISTS Transactions
(
    transactionId INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    price REAL(10, 2) NOT NULL,
    tenantId INTEGER NOT NULL,
    apartmentId INTEGER NOT NULL,
    date DATE NOT NULL,
    status INTEGER NOT NULL,
    FOREIGN KEY (tenantId) REFERENCES Tenants(tenantId),
    FOREIGN KEY (apartmentId) REFERENCES Apartments(apartmentId)
);

CREATE TABLE IF NOT EXISTS Payments -- apartment costs
(
    paymentId INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    price REAL(10, 2) NOT NULL,
    date DATE NOT NULL,
    expiry DATE NOT NULL,
    status VARCHAR(20) NOT NULL,
    apartmentId INTEGER NOT NULL,
    FOREIGN KEY(apartmentId) REFERENCES Apartments(apartmentId)
);

CREATE TABLE IF NOT EXISTS Histories
(
    historyId INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    apartmentId INTEGER NOT NULL,
    tenantId INTEGER NOT NULL,
    rentBegin DATE NOT NULL,
    rentEnd DATE NOT NULL,
    FOREIGN KEY (apartmentId) REFERENCES Apartments(apartmentId),
    FOREIGN KEY (tenantId) REFERENCES Tenants(tenantId)
);

CREATE TABLE IF NOT EXISTS Payments
(
    paymentId INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    price INTEGER NOT NULL,
    date DATE NOT NULL,
    expiry DATE NOT NULL,
    status INTEGER NOT NULL,
    apartmentId INTEGER NOT NULL,
    FOREIGN KEY (apartmentId) REFERENCES Apartments(apartmentId)
);
