ALTER TABLE opportunity DROP CONSTRAINT opportunity_title_company_un;
ALTER TABLE opportunity ADD CONSTRAINT opportunity_title_un UNIQUE (title);