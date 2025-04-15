CREATE TABLE IF NOT EXISTS organizations
(
    id          INTEGER PRIMARY KEY,

    name        VARCHAR(256) NOT NULL
        CONSTRAINT uq_organization_name UNIQUE,

    parent_id   INTEGER,
    external_id INTEGER,
    group_id    INTEGER,
    type        INTEGER,
    created_at  TIMESTAMP    NOT NULL
);

CREATE INDEX ix_organization_external_id
    ON organizations (external_id);

CREATE INDEX ix_organization_group_id
    ON organizations (group_id);
