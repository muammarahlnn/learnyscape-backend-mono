INSERT INTO roles (name)
VALUES
    ('Admin'),
    ('Lecturer'),
    ('Student');

INSERT INTO users (role_id, username, email, hash_password, full_name)
VALUES (1, 'admin', 'admin@unhas.ac.id', 'Admin123', 'Admin 1');