CREATE TABLE validators
(
    consensus_address TEXT NOT NULL PRIMARY KEY, /* Validator consensus address */
    consensus_pubkey  TEXT NOT NULL UNIQUE /* Validator consensus public key */
);

CREATE TABLE blocks
(
    height           BIGINT UNIQUE PRIMARY KEY,
    hash             TEXT                        NOT NULL UNIQUE,
    num_txs          INTEGER DEFAULT 0,
    total_gas        BIGINT  DEFAULT 0,
    proposer_address TEXT,
    timestamp        TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE pre_commits
(
    validator_address TEXT                        NOT NULL,
    height            BIGINT                      NOT NULL,
    timestamp         TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    voting_power      BIGINT                      NOT NULL,
    proposer_priority BIGINT                      NOT NULL,
    UNIQUE (validator_address, timestamp)
);

CREATE TABLE transactions
(
    hash         TEXT    NOT NULL,
    height       BIGINT  NOT NULL REFERENCES blocks (height),
    success      BOOLEAN NOT NULL,

    -- memo and signatures are part of the transaction body
    memo         TEXT,
    signatures   TEXT[]  NOT NULL,

    -- signer_infos and fee are part of the transaction auth info
    signer_infos JSONB   NOT NULL DEFAULT '[]'::JSONB,
    fee          JSONB   NOT NULL DEFAULT '{}'::JSONB,

    -- Additional fields from the transaction response
    gas_wanted   BIGINT           DEFAULT 0,
    gas_used     BIGINT           DEFAULT 0,
    raw_log      TEXT,
    logs         JSONB,

    -- partition_id is used to partition the table and make queries faster
    partition_id BIGINT  NOT NULL DEFAULT 0,

    CONSTRAINT unique_tx UNIQUE (hash, partition_id)
) PARTITION BY LIST (partition_id);

CREATE TABLE messages
(
    index            BIGINT NOT NULL,
    type             TEXT   NOT NULL,
    value            JSONB  NOT NULL,

    -- reference to the transaction table
    transaction_hash TEXT   NOT NULL,
    partition_id     BIGINT NOT NULL DEFAULT 0,
    FOREIGN KEY (transaction_hash, partition_id) REFERENCES transactions (hash, partition_id),

    CONSTRAINT unique_message_per_tx UNIQUE (transaction_hash, index, partition_id)
) PARTITION BY LIST (partition_id);


CREATE TABLE message_involved_accounts
(
    user_address     TEXT   NOT NULL,

    -- reference to the messages table
    message_index    BIGINT NOT NULL,
    transaction_hash TEXT   NOT NULL,
    partition_id     BIGINT NOT NULL DEFAULT 0,
    FOREIGN KEY (transaction_hash, message_index, partition_id) REFERENCES messages (transaction_hash, index, partition_id),

    CONSTRAINT unique_message_account UNIQUE (message_index, transaction_hash, user_address, partition_id)
) PARTITION BY LIST (partition_id);

CREATE TABLE pruning
(
    last_pruned_height BIGINT NOT NULL
);