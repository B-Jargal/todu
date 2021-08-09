
CREATE SEQUENCE public.owners_id_seq;

CREATE TABLE public.owners (
                id BIGINT NOT NULL DEFAULT nextval('public.owners_id_seq'),
                todu_id BIGINT NOT NULL,
                role VARCHAR(32) NOT NULL,
                name VARCHAR(100) NOT NULL,
                primary_alias VARCHAR(100) NOT NULL,
                CONSTRAINT owners_pk PRIMARY KEY (id)
);
COMMENT ON COLUMN public.owners.role IS 'Super admin or product owner';


ALTER SEQUENCE public.owners_id_seq OWNED BY public.owners.id;