CREATE TABLE IF NOT EXISTS todo (
    id INT AUTO_INCREMENT,
    title VARCHAR(50),
    description VARCHAR(100),
    completed BOOLEAN,

    CONSTRAINT pk_todo
    PRIMARY KEY (id)
);
