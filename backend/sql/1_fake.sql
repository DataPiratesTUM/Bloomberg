INSERT INTO organisations (id, name) VALUES ('5599931a-6f8e-442d-b52e-8c297af7cb8e', 'Bob E.V.');

INSERT INTO users (id, name, organisation) VALUES ('4e805cc9-fe3b-4649-96fc-f39634a557cd', 'Bob', '5599931a-6f8e-442d-b52e-8c297af7cb8e');
INSERT INTO users (id, name) VALUES ('c85e9879-f808-450d-8ab3-c6f5ab0e9d0c', 'Spongebob');

INSERT INTO securities (id, name, creator, description, ttl_1, ttl_2, funding_goal) VALUES (
    '3e8b7701-9d3e-407a-b78a-d8fa4d07bff5', 
    'LMU',
    '4e805cc9-fe3b-4649-96fc-f39634a557cd',
    'Unexzellent',
    100,
    100000000,
    1000000000
);
