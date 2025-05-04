INSERT INTO "user" VALUES (1, NOW(), NOW(), 'Local User', 'local@user.com', '')
ON CONFLICT (id) DO NOTHING;
