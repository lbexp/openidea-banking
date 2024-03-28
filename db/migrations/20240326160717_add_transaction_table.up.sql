CREATE TABLE IF NOT EXISTS transactions (
  transaction_id UUID PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  currency VARCHAR(10) NOT NULL,
  balance DECIMAL(12, 2) NOT NULL,
  proof_image_url VARCHAR(255),
  bank_account_number VARCHAR(255) NOT NULL,
  bank_name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
)