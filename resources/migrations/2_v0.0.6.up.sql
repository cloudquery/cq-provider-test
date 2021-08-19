ALTER TABLE "migrate_resource"
    ADD COLUMN IF NOT EXISTS upgrade_column_2 integer;
