CREATE TABLE Customers (
    ID int NOT NULL AUTO_INCREMENT,
    Name varchar(255) NOT NULL,
    ZipCode varchar(255),
    City varchar(255),
    Birthdate varchar(255),
    Status varchar(255),
    PRIMARY KEY (ID)
);

INSERT INTO Customers
    (Name,ZipCode,City,Birthdate,Status)
VALUES
    ('Ehsan','234234555','Tehran','18,Nov,1992','1'),
    ('Reza','5678568','Shiraz','7,Jan,1992','1'),
    ('Ali','7686430123','Rasht','12,Jun,1992','1'),
    ('Mohammad','2342556065','Mazandaran','24,Mar,1992','1'),
    ('Ghasem','7896846352','Mashhad','4,Apr,1992','1'),
    ('Bahar','4578678679','Bushehr','16,May,1992','1'),
    ('Farid','346535243','Kish','10,Jul,1992','1');