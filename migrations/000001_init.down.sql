DO $$
DECLARE
    tbl text;
BEGIN
    -- Loop through all tables in the public schema
    FOR tbl IN
        SELECT tablename
        FROM pg_tables
        WHERE schemaname = 'public'
    LOOP
        EXECUTE format('DROP TABLE IF EXISTS public.%I CASCADE;', tbl);
    END LOOP;
END$$;
