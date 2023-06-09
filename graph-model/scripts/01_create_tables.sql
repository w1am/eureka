DROP TYPE IF EXISTS entity_type CASCADE;

CREATE TYPE entity_type AS ENUM (
  'Person',
  'Organization',
  'Location',
  'Event'
);

DROP TABLE IF EXISTS entity;

CREATE TABLE entity (
    node_id       BIGSERIAL,
    entity_type   entity_type NOT NULL,
    entity_id     TEXT NOT NULL,
    PRIMARY KEY (node_id)
);

CREATE UNIQUE INDEX entity_entity_type_entity_id_unique_indx ON entity (entity_type, entity_id);

DROP TABLE IF EXISTS relation;

CREATE table relation (
    parent        BIGINT NOT NULL,
    child         BIGINT NOT NULL,
    PRIMARY KEY (parent, child)
);

CREATE INDEX relation_child_parent_indx ON relation(child, parent);
