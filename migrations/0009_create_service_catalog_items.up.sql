CREATE TABLE service_catalog_items (
    id text PRIMARY KEY,
    name text NOT NULL,
    normalized_name text NOT NULL,
    default_price_rupiah bigint NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT service_catalog_items_default_price_positive_check
        CHECK (default_price_rupiah > 0)
);

CREATE UNIQUE INDEX service_catalog_items_normalized_name_unique
    ON service_catalog_items (normalized_name);

CREATE INDEX service_catalog_items_active_name_idx
    ON service_catalog_items (is_active, normalized_name);
