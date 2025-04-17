-- create "organizations" table
CREATE TABLE "organizations" (
  "id" bigserial NOT NULL,
  "name" character varying(128) NOT NULL,
  "type" smallint NOT NULL,
  "parent_id" bigint NULL,
  "external_id" bigint NULL,
  "group_id" bigint NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_organizations_external_id" to table: "organizations"
CREATE INDEX "idx_organizations_external_id" ON "organizations" ("external_id");
-- create index "idx_organizations_group_id" to table: "organizations"
CREATE INDEX "idx_organizations_group_id" ON "organizations" ("group_id");
-- create index "idx_organizations_name" to table: "organizations"
CREATE UNIQUE INDEX "idx_organizations_name" ON "organizations" ("name");
