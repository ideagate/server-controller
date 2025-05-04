INSERT INTO application VALUES ('test-application', 'test-project', NOW(), NOW(), 'Test Application', 'Test Application Description')
ON CONFLICT (id, project_id) DO NOTHING;