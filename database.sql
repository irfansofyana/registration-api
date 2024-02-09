/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/**
  Installing extensions
 */
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/** This is the table for the users. */
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    full_name VARCHAR(60) NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    password VARCHAR(64) NOT NULL,
    successful_login_count BIGINT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

/**
  A lot of queries using phone number will be executed, so we need to index the phone number.
 */
CREATE INDEX idx_users_phone_number ON users (phone_number);