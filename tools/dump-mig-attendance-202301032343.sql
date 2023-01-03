--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6 (Homebrew)
-- Dumped by pg_dump version 15.1

-- Started on 2023-01-03 23:43:33 WIB

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

DROP DATABASE "mig-attendance";
--
-- TOC entry 3604 (class 1262 OID 16475)
-- Name: mig-attendance; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "mig-attendance" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';


ALTER DATABASE "mig-attendance" OWNER TO postgres;

\connect -reuse-previous=on "dbname='mig-attendance'"

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

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: sholeh
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO sholeh;

--
-- TOC entry 3605 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: sholeh
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 212 (class 1259 OID 16492)
-- Name: mig_attendance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mig_attendance (
    id integer NOT NULL,
    user_id integer NOT NULL,
    activity text,
    check_in timestamp without time zone,
    check_out timestamp without time zone
);


ALTER TABLE public.mig_attendance OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 16491)
-- Name: mig_attendance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mig_attendance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mig_attendance_id_seq OWNER TO postgres;

--
-- TOC entry 3607 (class 0 OID 0)
-- Dependencies: 211
-- Name: mig_attendance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mig_attendance_id_seq OWNED BY public.mig_attendance.id;


--
-- TOC entry 210 (class 1259 OID 16483)
-- Name: mig_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mig_users (
    id integer NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL
);


ALTER TABLE public.mig_users OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16482)
-- Name: mig_users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mig_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mig_users_id_seq OWNER TO postgres;

--
-- TOC entry 3608 (class 0 OID 0)
-- Dependencies: 209
-- Name: mig_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mig_users_id_seq OWNED BY public.mig_users.id;


--
-- TOC entry 3450 (class 2604 OID 16495)
-- Name: mig_attendance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mig_attendance ALTER COLUMN id SET DEFAULT nextval('public.mig_attendance_id_seq'::regclass);


--
-- TOC entry 3449 (class 2604 OID 16486)
-- Name: mig_users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mig_users ALTER COLUMN id SET DEFAULT nextval('public.mig_users_id_seq'::regclass);


--
-- TOC entry 3598 (class 0 OID 16492)
-- Dependencies: 212
-- Data for Name: mig_attendance; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.mig_attendance VALUES (2, 1, 'activity 2', '2022-01-02 00:00:00', '2022-01-02 23:59:59');
INSERT INTO public.mig_attendance VALUES (3, 1, 'activity 3', '2022-01-03 00:00:00', '2022-01-03 23:59:59');
INSERT INTO public.mig_attendance VALUES (4, 1, 'activity 4', '2022-01-04 00:00:00', '2022-01-04 23:59:59');
INSERT INTO public.mig_attendance VALUES (5, 1, 'activity 5', '2022-01-05 00:00:00', '2022-01-05 23:59:59');
INSERT INTO public.mig_attendance VALUES (6, 1, 'activity 6', '2022-01-06 00:00:00', '2022-01-06 23:59:59');
INSERT INTO public.mig_attendance VALUES (7, 1, 'activity 7', '2022-01-07 00:00:00', '2022-01-07 23:59:59');
INSERT INTO public.mig_attendance VALUES (8, 1, 'activity 8', '2022-01-08 00:00:00', '2022-01-08 23:59:59');
INSERT INTO public.mig_attendance VALUES (9, 1, 'activity 9', '2022-01-09 00:00:00', '2022-01-09 23:59:59');
INSERT INTO public.mig_attendance VALUES (10, 1, 'activity 10', '2022-01-10 00:00:00', '2022-01-10 23:59:59');
INSERT INTO public.mig_attendance VALUES (11, 2, 'activity 11', '2022-01-11 00:00:00', '2022-01-11 23:59:59');
INSERT INTO public.mig_attendance VALUES (12, 2, 'activity 12', '2022-01-12 00:00:00', '2022-01-12 23:59:59');
INSERT INTO public.mig_attendance VALUES (13, 2, 'activity 13', '2022-01-13 00:00:00', '2022-01-13 23:59:59');
INSERT INTO public.mig_attendance VALUES (14, 2, 'activity 14', '2022-01-14 00:00:00', '2022-01-14 23:59:59');
INSERT INTO public.mig_attendance VALUES (15, 2, 'activity 15', '2022-01-15 00:00:00', '2022-01-15 23:59:59');
INSERT INTO public.mig_attendance VALUES (16, 2, 'activity 16', '2022-01-16 00:00:00', '2022-01-16 23:59:59');
INSERT INTO public.mig_attendance VALUES (17, 2, 'activity 17', '2022-01-17 00:00:00', '2022-01-17 23:59:59');
INSERT INTO public.mig_attendance VALUES (18, 2, 'activity 18', '2022-01-18 00:00:00', '2022-01-18 23:59:59');
INSERT INTO public.mig_attendance VALUES (19, 2, 'activity 19', '2022-01-19 00:00:00', '2022-01-19 23:59:59');
INSERT INTO public.mig_attendance VALUES (1, 1, 'my activity today is debugging my code', '2022-01-01 00:00:00', '2022-01-01 23:59:59');


--
-- TOC entry 3596 (class 0 OID 16483)
-- Dependencies: 210
-- Data for Name: mig_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.mig_users VALUES (1, 'Karyawan 1', 'user@gmail.com', '$2a$12$vJsjS3x9rvz/.S2DyyeKx.i5x7uRaiw11AUn575eTtBnMeGNyRDjK');
INSERT INTO public.mig_users VALUES (2, 'Karyawan 2', 'user2@email.com', '$2a$12$GuPu/4GpJ3qv6rg5NqMX7e9iUaITy8Bvu95CVrkzqhp2IrghBm5X6');


--
-- TOC entry 3609 (class 0 OID 0)
-- Dependencies: 211
-- Name: mig_attendance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mig_attendance_id_seq', 19, true);


--
-- TOC entry 3610 (class 0 OID 0)
-- Dependencies: 209
-- Name: mig_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mig_users_id_seq', 2, true);


--
-- TOC entry 3455 (class 2606 OID 16499)
-- Name: mig_attendance mig_attendance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mig_attendance
    ADD CONSTRAINT mig_attendance_pkey PRIMARY KEY (id);


--
-- TOC entry 3452 (class 2606 OID 16490)
-- Name: mig_users mig_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mig_users
    ADD CONSTRAINT mig_users_pkey PRIMARY KEY (id);


--
-- TOC entry 3453 (class 1259 OID 16501)
-- Name: idx_user_id_check_in; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_user_id_check_in ON public.mig_attendance USING btree (user_id, date(check_in));


--
-- TOC entry 3606 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: sholeh
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2023-01-03 23:43:33 WIB

--
-- PostgreSQL database dump complete
--

