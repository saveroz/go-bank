BEGIN;
CREATE TABLE IF NOT EXISTS account(
  id varchar(255) NOT NULL PRIMARY KEY,
  name varchar(100) NOT NULL,
  balance int NOT NULL,
  account_no int NOT NULL,
  created timestamp NOT NULL,
  updated timestamp NOT NULL,
  CONSTRAINT account_no_unique_idx UNIQUE(account_no)
);
COMMIT;