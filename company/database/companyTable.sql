DROP TABLE IF EXISTS company;
CREATE TABLE company(
    id INT AUTO_INCREMENT NOT NULL,
    name  VARCHAR(500) NOT NULL,
    creator  VARCHAR(500) NOT NULL,
     PRIMARY KEY (`id`)
);

INSERT INTO company
(name,creator)
VALUES 
("apple","steve jobs"),
("zip2","elon musk");