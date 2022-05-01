BEGIN;
CREATE TABLE IF NOT EXISTS "transaction"(
  id varchar(255) NOT NULL PRIMARY KEY,
  account_no int NOT NULL,
  amount int NOT NULL,
  type varchar(64) NOT NULL,
  transaction_type varchar(64) NOT NULL,
  description varchar(255) NOT NULL,
  created timestamptz NOT NULL,
  updated timestamptz NOT NULL
);
COMMIT;