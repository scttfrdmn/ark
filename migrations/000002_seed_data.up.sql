-- Seed data for development and testing

-- Insert sample training modules

-- S3 Basics
INSERT INTO training_modules (name, title, description, category, difficulty, estimated_minutes, content) VALUES
('s3-basics', 'S3 Basics', 'Learn the fundamentals of Amazon S3 object storage', 's3', 'beginner', 15,
'{"sections": [
  {"title": "What is S3?", "content": "Amazon S3 is object storage built to store and retrieve any amount of data from anywhere."},
  {"title": "Buckets and Objects", "content": "S3 stores data as objects within buckets. Each object consists of data, metadata, and a unique identifier."},
  {"title": "Security", "content": "Use bucket policies and IAM policies to control access. Enable encryption at rest and in transit."}
]}'::jsonb);

-- S3 Security
INSERT INTO training_modules (name, title, description, category, difficulty, estimated_minutes, content, prerequisites) VALUES
('s3-security', 'S3 Security Best Practices', 'Learn how to secure your S3 buckets and data', 's3', 'intermediate', 20,
'{"sections": [
  {"title": "Bucket Policies", "content": "Control access to your buckets using bucket policies."},
  {"title": "Encryption", "content": "Encrypt data at rest using SSE-S3, SSE-KMS, or SSE-C."},
  {"title": "Access Logging", "content": "Enable server access logging to track requests."},
  {"title": "Versioning", "content": "Enable versioning to protect against accidental deletion."}
]}'::jsonb,
ARRAY['s3-basics']);

-- IAM Basics
INSERT INTO training_modules (name, title, description, category, difficulty, estimated_minutes, content) VALUES
('iam-basics', 'IAM Fundamentals', 'Understand AWS Identity and Access Management', 'iam', 'beginner', 20,
'{"sections": [
  {"title": "What is IAM?", "content": "IAM enables you to manage access to AWS services and resources securely."},
  {"title": "Users and Groups", "content": "Create IAM users and organize them into groups."},
  {"title": "Roles", "content": "Use IAM roles to grant temporary access to AWS resources."},
  {"title": "Policies", "content": "Define permissions using JSON policy documents."}
]}'::jsonb);

-- EC2 Basics
INSERT INTO training_modules (name, title, description, category, difficulty, estimated_minutes, content) VALUES
('ec2-basics', 'EC2 Fundamentals', 'Learn about Amazon EC2 virtual servers', 'ec2', 'beginner', 25,
'{"sections": [
  {"title": "What is EC2?", "content": "Amazon EC2 provides scalable computing capacity in the cloud."},
  {"title": "Instance Types", "content": "Choose the right instance type for your workload."},
  {"title": "Security Groups", "content": "Control inbound and outbound traffic with security groups."},
  {"title": "Key Pairs", "content": "Use key pairs for secure SSH access to instances."}
]}'::jsonb);

-- Insert sample policies

-- Training gate for S3
INSERT INTO policies (name, description, policy_type, rules, applies_to) VALUES
('s3-training-gate', 'Require S3 basics training before creating buckets', 'training_gate',
'{"required_modules": ["s3-basics"], "actions": ["s3:CreateBucket"]}'::jsonb,
ARRAY['researcher']);

-- Resource limit
INSERT INTO policies (name, description, policy_type, rules, applies_to) VALUES
('s3-bucket-limit', 'Limit number of S3 buckets per user', 'resource_limit',
'{"resource_type": "s3:bucket", "max_count": 10}'::jsonb,
ARRAY['researcher']);

-- EC2 instance limit
INSERT INTO policies (name, description, policy_type, rules, applies_to) VALUES
('ec2-instance-limit', 'Limit number of EC2 instances per user', 'resource_limit',
'{"resource_type": "ec2:instance", "max_count": 5, "max_total_vcpus": 16}'::jsonb,
ARRAY['researcher']);
