CREATE TABLE public.users (
    id serial PRIMARY KEY,
    email varchar,
    first_name varchar,
    last_name varchar,
    password varchar,
    user_active integer DEFAULT 0,
    created_at timestamp without time zone DEFAULT current_timestamp,
    updated_at timestamp without time zone DEFAULT current_timestamp
);

INSERT INTO public.users("email","first_name","last_name","password","user_active","created_at","updated_at")
VALUES ('admin@example.com', 'Admin', 'User', '$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe', 1, '2022-03-14 00:00:00', '2022-03-14 00:00:00');