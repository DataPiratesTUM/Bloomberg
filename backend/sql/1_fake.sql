INSERT INTO users (id, name) VALUES ('4e805cc9-fe3b-4649-96fc-f39634a557cd', 'Bob');

INSERT INTO organisations (id, name) VALUES ('5599931a-6f8e-442d-b52e-8c297af7cb8e', 'Bob E.V.');

INSERT INTO securities (id, name, creator, description, ttl_1, ttl_2) VALUES (
    '3e8b7701-9d3e-407a-b78a-d8fa4d07bff5', 
    'LMU',
    '5599931a-6f8e-442d-b52e-8c297af7cb8e',
    'Unexzellent',
    now() + INTERVAL '1 DAY',
    now() + INTERVAL '2 DAY'
);