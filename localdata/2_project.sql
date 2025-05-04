INSERT INTO project VALUES ('test-project', NOW(), NOW(), 'Test Project', 'Test Project Description')
ON CONFLICT (id) DO NOTHING;