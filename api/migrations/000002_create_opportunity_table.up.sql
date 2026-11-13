CREATE TABLE opportunity (
    id UUID NOT NULL DEFAULT uuid_generate_v1mc(),
    title TEXT NOT NULL,
    company TEXT NOT NULL,
    category TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT opportunity_pk PRIMARY KEY (id),
    CONSTRAINT opportunity_title_un UNIQUE (title)
)