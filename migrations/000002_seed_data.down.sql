-- Rollback seed data

-- Delete sample policies
DELETE FROM policies WHERE name IN ('s3-training-gate', 's3-bucket-limit', 'ec2-instance-limit');

-- Delete sample training modules
DELETE FROM training_modules WHERE name IN ('s3-basics', 's3-security', 'iam-basics', 'ec2-basics');
