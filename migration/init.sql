CREATE TABLE casbin_rule (
    id SERIAL PRIMARY KEY,
    ptype VARCHAR(100),
    v0 VARCHAR(100),
    v1 VARCHAR(100),
    v2 VARCHAR(100),
    v3 VARCHAR(100),
    v4 VARCHAR(100),
    v5 VARCHAR(100)
);

INSERT INTO casbin_rule(ptype, v0, v1, v2) VALUES('p', 'admin', '/project', '*');
INSERT INTO casbin_rule(ptype, v0, v1, v2) VALUES('p', 'admin', '/channel', '*');
INSERT INTO casbin_rule(ptype, v0, v1, v2) VALUES('p', 'user', '/project', 'GET');
INSERT INTO casbin_rule(ptype, v0, v1, v2) VALUES('p', 'user', '/channel', 'GET');

INSERT INTO casbin_rule(ptype, v0, v1) VALUES('g', 'alice', 'admin');
INSERT INTO casbin_rule(ptype, v0, v1) VALUES('g', 'bob', 'user');