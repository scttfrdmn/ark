-- Ark Database Initialization
-- This script runs automatically when PostgreSQL container starts

-- Create extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Future: Add initial schema migrations here
-- For now, migrations will be handled by migration tool

-- Create a healthcheck table
CREATE TABLE IF NOT EXISTS schema_version (
    version INTEGER PRIMARY KEY,
    applied_at TIMESTAMP DEFAULT NOW()
);

-- Insert initial version
INSERT INTO schema_version (version) VALUES (0)
ON CONFLICT (version) DO NOTHING;
