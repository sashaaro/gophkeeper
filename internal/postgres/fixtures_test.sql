INSERT INTO "user" (id, login, pass)
VALUES
    ('01ef6697-3190-6984-9572-74563c32efde', 'test', '123');

INSERT INTO secret (id, user_id, name, kind)
VALUES
    ('11111111-3190-6984-9572-74563c32efde', '01ef6697-3190-6984-9572-74563c32efde', 'testCred', 'credentials');

INSERT INTO credentials (id, login, password)
VALUES
    ('11111111-3190-6984-9572-74563c32efde', 'bob' , 'z$3a!');
