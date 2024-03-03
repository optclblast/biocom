CREATE TYPE nomenclature_type AS ENUM (
        'product', 'assembly_unit', 'component', 'service'
);

CREATE TABLE nomenclature (
  id varchar(36)  PRIMARY KEY unique,
  company_id varchar(36)  not null,
  name text,
  created_at timestamp default CURRENT_TIMESTAMP,
  updated_at timestamp default CURRENT_TIMESTAMP,
  deleted_at timestamp default null,
  type nomenclature_type
);

create index index_nomenclature_id_company_id 
        on nomenclature (id, company_id);

CREATE TABLE nomenclature_composition (
  nomenclature_id varchar(36)  PRIMARY KEY,
  part_id varchar(36)  NOT NULL,
  amount float DEFAULT 0
);

create index index_nomenclature_composition_id_part_id 
        on nomenclature_composition (nomenclature_id, part_id);

CREATE TABLE storage_nomenclature (
  storage_id varchar(36) not null,
  nomenclature_id varchar(36) not null,
  amount float DEFAULT 0,
  created_at timestamp,
  updated_at timestamp
);

create index index_nomenclature_id_storage_id_nomenclature_id 
        on storage_nomenclature (storage_id, nomenclature_id);

CREATE TABLE storages (
  id varchar(36) PRIMARY KEY unique,
  name text,
  address text,
  created_at timestamp,
  updated_at timestamp
);

ALTER TABLE storage_nomenclature ADD FOREIGN KEY (storage_id) REFERENCES storages (id);

ALTER TABLE storage_nomenclature ADD FOREIGN KEY (nomenclature_id) REFERENCES nomenclature (id);
