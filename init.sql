-- Drop tables if they exist (for idempotency)
DROP TABLE IF EXISTS coupons;
DROP TABLE IF EXISTS campaigns;

-- Enable extension for UUID generation (PostgreSQL only)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create campaigns table
CREATE TABLE campaigns (
                           id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                           name TEXT NOT NULL,
                           total_count INTEGER NOT NULL,
                           start_at TIMESTAMPTZ NOT NULL,
                           end_at TIMESTAMPTZ,
                           created_at TIMESTAMPTZ DEFAULT NOW(),
                           updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Create coupons table
CREATE TABLE coupons (
                         code VARCHAR(10) PRIMARY KEY,
                         campaign_id UUID NOT NULL REFERENCES campaigns(id) ON UPDATE CASCADE ON DELETE CASCADE,
                         user_id TEXT NOT NULL,
                         issued_at TIMESTAMPTZ DEFAULT NOW(),
                         CONSTRAINT idx_campaign_user UNIQUE (campaign_id, user_id)
);

-- Insert a sample campaign
INSERT INTO campaigns (id, name, total_count, start_at, end_at)
VALUES (
           '00000000-0000-0000-0000-000000000001',
           'Sample Campaign',
           1000000,
           NOW() - INTERVAL '10 minutes',
           NOW() + INTERVAL '1 day'
       );
