CREATE TABLE IF NOT EXISTS users (
  id VARCHAR(100) PRIMARY KEY,
  username VARCHAR(100),
  password VARCHAR(100),
  role VARCHAR(30),
  reset_token VARCHAR(100),
  is_active VARCHAR(30)
);