CREATE TABLE IF NOT EXISTS "states" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "products" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL INDEX,
  "code" varchar(50) INDEX,
  "cost" float ,
  "sale_price" float NOT NULL,
  "units" int NOT NULL,
  "description" varchar(255),
  "id_state" int NOT NULL,
  "created_at" date NOT NULL
);

ALTER TABLE "products" ADD FOREIGN KEY ("id_state") REFERENCES "states" ("id");

CREATE TABLE IF NOT EXISTS "purchase" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "total" float NOT NULL,
  "created_at" date NOT NULL
);

CREATE TABLE IF NOT EXISTS "purchase_details" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "id_purchase" int NOT NULL
  "unit_cost" float NOT NULL,
  "units" int NOT NULL,
  "id_product" int NOT NULL
);

ALTER TABLE "purchase_details" ADD FOREIGN KEY ("id_purchase") REFERENCES "purchase" ("id");
ALTER TABLE "purchase_details" ADD FOREIGN KEY ("id_product") REFERENCES "products" ("id");

CREATE TABLE IF NOT EXISTS "sales" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "total" float NOT NULL,
  "created_at" date NOT NULL
);

CREATE TABLE IF NOT EXISTS "sales_details" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "id_sale" int NOT NULL,
  "id_product" int NOT NULL,
  "units" int NOT NULL,
  "unit_price" float NOT NULL,
  "total" float NOT NULL
  
);

ALTER TABLE "sales_details" ADD FOREIGN KEY ("id_product") REFERENCES "products" ("id");

ALTER TABLE "sales_details" ADD FOREIGN KEY ("id_sale") REFERENCES "sales" ("id");

CREATE TABLE IF NOT EXISTS "losses" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "total" float NOT NULL,
  "created_at" date NOT NULL
  
);

CREATE TABLE IF NOT EXISTS "loss_details" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "id_loss" int NOT NULL,
  "id_product" int NOT NULL,
  "unit_cost" float NOT NULL,
  "units" int NOT NULL,
  "total" float NOT NULL
);

ALTER TABLE "loss_details" ADD FOREIGN KEY ("id_loss") REFERENCES "losses" ("id");

ALTER TABLE "loss_details" ADD FOREIGN KEY ("id_product") REFERENCES "products" ("id");


CREATE TABLE IF NOT EXISTS "role" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "users" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(50) NOT NULL,
  "username" varchar(50) NOT NULL UNIQUE,
  "email" varchar(50) NOT NULL UNIQUE,
  "password" varchar(200) NOT NULL,
  "id_role" int NOT NULL
);

ALTER TABLE "users" ADD FOREIGN KEY ("id_role") REFERENCES "role" ("id");


CREATE TABLE IF NOT EXISTS "modules" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(50) NOT NULL UNIQUE
);

insert into modules (name) values ('users'),('inventory');

CREATE TABLE IF NOT EXISTS "operations" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(50) NOT null,
  "id_module" int NOT NULL
);

ALTER TABLE "operations" ADD FOREIGN KEY ("id_module") REFERENCES "modules" ("id");

CREATE TABLE IF NOT EXISTS "role_operations" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "id_role" int NOT NULL,
  "id_operation" int  NOT NULL
);

ALTER TABLE "role_operations" ADD FOREIGN KEY ("id_role") REFERENCES "role" ("id");
ALTER TABLE "role_operations" ADD FOREIGN KEY ("id_operation") REFERENCES "operations" ("id");


INSERT INTO operations (name,id_module) VALUES
	 ('create',1),
	 ('update',1),
	 ('read',1),
	 ('delete',1),
	 ('create',2),
	 ('update',2),
	 ('delele',2),
	 ('read',2)

INSERT INTO role (name) VALUES ('administrator');

insert into role_operations (id_role,id_operation) values (1,1),(1,2),(1,3),(1,4),(1,5),(1,6),(1,7),(1,8);
