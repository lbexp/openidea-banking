CREATE TABLE IF NOT EXISTS balances (
  balance_id UUID PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  currency VARCHAR(10) NOT NULL,
  balance DECIMAL(12, 2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  UNIQUE(user_id, currency),
  FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
)