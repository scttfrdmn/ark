-- Rollback initial schema

-- Drop triggers
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
DROP TRIGGER IF EXISTS update_training_modules_updated_at ON training_modules;
DROP TRIGGER IF EXISTS update_user_training_progress_updated_at ON user_training_progress;
DROP TRIGGER IF EXISTS update_policies_updated_at ON policies;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop tables (in reverse order of dependencies)
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS policies;
DROP TABLE IF EXISTS audit_logs;
DROP TABLE IF EXISTS user_training_progress;
DROP TABLE IF EXISTS training_modules;
DROP TABLE IF EXISTS users;

-- Drop extensions
DROP EXTENSION IF EXISTS "uuid-ossp";
