--
-- PostgreSQL database dump
--

-- Dumped from database version 12.4
-- Dumped by pg_dump version 12.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: incidents; Type: TABLE; Schema: public; Owner: hikeshi
--

CREATE TABLE public.incidents (
    id uuid NOT NULL,
    date date NOT NULL,
    date_closed date NOT NULL,
    severity character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    summary text NOT NULL,
    scope character varying(255) NOT NULL,
    responsible_party character varying(255) NOT NULL,
    result character varying(255) NOT NULL,
    mitigation character varying(255) NOT NULL,
    affected_customers character varying(255) NOT NULL,
    root_cause character varying(255) NOT NULL,
    slack_channel character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.incidents OWNER TO hikeshi;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: hikeshi
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO hikeshi;

--
-- Name: users; Type: TABLE; Schema: public; Owner: hikeshi
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO hikeshi;

--
-- Name: incidents incidents_pkey; Type: CONSTRAINT; Schema: public; Owner: hikeshi
--

ALTER TABLE ONLY public.incidents
    ADD CONSTRAINT incidents_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: hikeshi
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: hikeshi
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

