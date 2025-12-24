/*
 Navicat Premium Data Transfer

 Source Server         : postgre_server
 Source Server Type    : PostgreSQL
 Source Server Version : 110022
 Source Host           : localhost:5432
 Source Catalog        : indico_db
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 110022
 File Encoding         : 65001

 Date: 24/12/2025 18:56:44
*/


-- ----------------------------
-- Sequence structure for vouchers_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."vouchers_id_seq";
CREATE SEQUENCE "public"."vouchers_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for vouchers
-- ----------------------------
DROP TABLE IF EXISTS "public"."vouchers";
CREATE TABLE "public"."vouchers" (
  "id" "pg_catalog"."int8" NOT NULL DEFAULT nextval('vouchers_id_seq'::regclass),
  "code" "pg_catalog"."varchar" COLLATE "pg_catalog"."default" NOT NULL,
  "name" "pg_catalog"."varchar" COLLATE "pg_catalog"."default" NOT NULL,
  "description" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "discount" "pg_catalog"."numeric" NOT NULL,
  "max_usage" "pg_catalog"."int8" NOT NULL DEFAULT 1,
  "used_count" "pg_catalog"."int8" DEFAULT 0,
  "valid_from" "pg_catalog"."timestamptz" NOT NULL,
  "valid_until" "pg_catalog"."timestamptz" NOT NULL,
  "is_active" "pg_catalog"."bool" DEFAULT true,
  "created_at" "pg_catalog"."timestamptz",
  "updated_at" "pg_catalog"."timestamptz",
  "deleted_at" "pg_catalog"."timestamptz"
)
;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."vouchers_id_seq"
OWNED BY "public"."vouchers"."id";
SELECT setval('"public"."vouchers_id_seq"', 24, true);

-- ----------------------------
-- Indexes structure for table vouchers
-- ----------------------------
CREATE UNIQUE INDEX "idx_vouchers_code" ON "public"."vouchers" USING btree (
  "code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_vouchers_deleted_at" ON "public"."vouchers" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table vouchers
-- ----------------------------
ALTER TABLE "public"."vouchers" ADD CONSTRAINT "vouchers_pkey" PRIMARY KEY ("id");
