CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL, -- can be positive or negative
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL, -- must be positive
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  
  -- ENFORCE INTEGRITY: Prevent negative or zero value transfers
  CONSTRAINT "transfer_amount_positive" CHECK ("amount" > 0),
  -- ENFORCE INTEGRITY: Prevent transferring money to the same account
  CONSTRAINT "transfer_to_different_account" CHECK ("from_account_id" <> "to_account_id")
);

-- Indexing for fast account ownership lookups
CREATE INDEX ON "accounts" ("owner");

-- Indexing for fast entry history retrieval per account
CREATE INDEX ON "entries" ("account_id", "created_at" DESC);

-- Clean, ultra-efficient pagination indexes for transaction history
CREATE INDEX idx_transfers_from_account_pagination ON transfers (from_account_id, created_at DESC, id DESC);
CREATE INDEX idx_transfers_to_account_pagination ON transfers (to_account_id, created_at DESC, id DESC);

-- Foreign Key Constraints
ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE RESTRICT;
ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id") ON DELETE RESTRICT;
ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id") ON DELETE RESTRICT;

