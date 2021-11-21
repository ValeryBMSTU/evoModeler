--
-- PostgreSQL database dump
--

-- Dumped from database version 10.19 (Ubuntu 10.19-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.19 (Ubuntu 10.19-0ubuntu0.18.04.1)

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

DROP DATABASE IF EXISTS postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'ru_RU.UTF-8' LC_CTYPE = 'ru_RU.UTF-8';


ALTER DATABASE postgres OWNER TO postgres;

\connect postgres

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
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: Agent; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Agent" (
    id integer NOT NULL,
    target_function real NOT NULL,
    id_parent integer,
    id_generation integer NOT NULL,
    id_template integer
);


ALTER TABLE public."Agent" OWNER TO postgres;

--
-- Name: Agent_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Agent_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Agent_id_seq" OWNER TO postgres;

--
-- Name: Agent_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Agent_id_seq" OWNED BY public."Agent".id;


--
-- Name: Gen; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Gen" (
    id integer NOT NULL,
    "position" integer NOT NULL,
    value integer NOT NULL,
    id_agent integer NOT NULL
);


ALTER TABLE public."Gen" OWNER TO postgres;

--
-- Name: Geb_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Geb_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Geb_id_seq" OWNER TO postgres;

--
-- Name: Geb_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Geb_id_seq" OWNED BY public."Gen".id;


--
-- Name: GenTemplate; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."GenTemplate" (
    id integer NOT NULL,
    id_gen integer NOT NULL,
    id_template integer NOT NULL
);


ALTER TABLE public."GenTemplate" OWNER TO postgres;

--
-- Name: GenTemplate_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."GenTemplate_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."GenTemplate_id_seq" OWNER TO postgres;

--
-- Name: GenTemplate_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."GenTemplate_id_seq" OWNED BY public."GenTemplate".id;


--
-- Name: Generation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Generation" (
    id integer NOT NULL,
    order_number integer NOT NULL,
    id_session integer NOT NULL,
    extra_params text
);


ALTER TABLE public."Generation" OWNER TO postgres;

--
-- Name: Generation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Generation_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Generation_id_seq" OWNER TO postgres;

--
-- Name: Generation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Generation_id_seq" OWNED BY public."Generation".id;


--
-- Name: GeneticAlgorithm; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."GeneticAlgorithm" (
    id integer NOT NULL,
    name text NOT NULL,
    create_date date NOT NULL,
    description text,
    config text
);


ALTER TABLE public."GeneticAlgorithm" OWNER TO postgres;

--
-- Name: Issue; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Issue" (
    id integer NOT NULL,
    name character varying NOT NULL,
    description text
);


ALTER TABLE public."Issue" OWNER TO postgres;

--
-- Name: Issue_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Issue_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Issue_id_seq" OWNER TO postgres;

--
-- Name: Issue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Issue_id_seq" OWNED BY public."Issue".id;


--
-- Name: Session; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Session" (
    id integer NOT NULL,
    id_user integer NOT NULL,
    is_deleted boolean DEFAULT false NOT NULL
);


ALTER TABLE public."Session" OWNER TO postgres;

--
-- Name: Task; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Task" (
    id integer NOT NULL,
    name text NOT NULL,
    create_date date NOT NULL,
    description text,
    "id_GA" integer,
    id_user integer NOT NULL,
    id_solver integer NOT NULL
);


ALTER TABLE public."Task" OWNER TO postgres;

--
-- Name: Session_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Session_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Session_id_seq" OWNER TO postgres;

--
-- Name: Session_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Session_id_seq" OWNED BY public."Task".id;


--
-- Name: Session_id_seq1; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Session_id_seq1"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Session_id_seq1" OWNER TO postgres;

--
-- Name: Session_id_seq1; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Session_id_seq1" OWNED BY public."Session".id;


--
-- Name: Solver; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Solver" (
    id integer NOT NULL,
    name character varying NOT NULL,
    description text,
    model text,
    id_issue integer NOT NULL
);


ALTER TABLE public."Solver" OWNER TO postgres;

--
-- Name: Template; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Template" (
    id integer NOT NULL,
    length integer NOT NULL
);


ALTER TABLE public."Template" OWNER TO postgres;

--
-- Name: Template_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Template_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Template_id_seq" OWNER TO postgres;

--
-- Name: Template_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Template_id_seq" OWNED BY public."Template".id;


--
-- Name: User; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."User" (
    id integer NOT NULL,
    login character varying NOT NULL,
    pass character varying NOT NULL
);


ALTER TABLE public."User" OWNER TO postgres;

--
-- Name: User_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."User_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."User_id_seq" OWNER TO postgres;

--
-- Name: User_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."User_id_seq" OWNED BY public."User".id;


--
-- Name: ga_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ga_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ga_id_seq OWNER TO postgres;

--
-- Name: ga_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ga_id_seq OWNED BY public."GeneticAlgorithm".id;


--
-- Name: solver_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.solver_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.solver_id_seq OWNER TO postgres;

--
-- Name: solver_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.solver_id_seq OWNED BY public."Solver".id;


--
-- Name: Agent id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Agent" ALTER COLUMN id SET DEFAULT nextval('public."Agent_id_seq"'::regclass);


--
-- Name: Gen id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Gen" ALTER COLUMN id SET DEFAULT nextval('public."Geb_id_seq"'::regclass);


--
-- Name: GenTemplate id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."GenTemplate" ALTER COLUMN id SET DEFAULT nextval('public."GenTemplate_id_seq"'::regclass);


--
-- Name: Generation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Generation" ALTER COLUMN id SET DEFAULT nextval('public."Generation_id_seq"'::regclass);


--
-- Name: GeneticAlgorithm id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."GeneticAlgorithm" ALTER COLUMN id SET DEFAULT nextval('public.ga_id_seq'::regclass);


--
-- Name: Issue id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Issue" ALTER COLUMN id SET DEFAULT nextval('public."Issue_id_seq"'::regclass);


--
-- Name: Session id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session" ALTER COLUMN id SET DEFAULT nextval('public."Session_id_seq1"'::regclass);


--
-- Name: Solver id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Solver" ALTER COLUMN id SET DEFAULT nextval('public.solver_id_seq'::regclass);


--
-- Name: Task id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task" ALTER COLUMN id SET DEFAULT nextval('public."Session_id_seq"'::regclass);


--
-- Name: Template id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Template" ALTER COLUMN id SET DEFAULT nextval('public."Template_id_seq"'::regclass);


--
-- Name: User id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User" ALTER COLUMN id SET DEFAULT nextval('public."User_id_seq"'::regclass);


--
-- Data for Name: Agent; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Agent" (id, target_function, id_parent, id_generation, id_template) FROM stdin;
\.


--
-- Data for Name: Gen; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Gen" (id, "position", value, id_agent) FROM stdin;
\.


--
-- Data for Name: GenTemplate; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."GenTemplate" (id, id_gen, id_template) FROM stdin;
\.


--
-- Data for Name: Generation; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Generation" (id, order_number, id_session, extra_params) FROM stdin;
\.


--
-- Data for Name: GeneticAlgorithm; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."GeneticAlgorithm" (id, name, create_date, description, config) FROM stdin;
\.


--
-- Data for Name: Issue; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Issue" (id, name, description) FROM stdin;
1	TSP	The traveling salesman problem (also called the travelling salesperson problem or TSP) asks the following question: "Given a list of cities and the distances between each pair of cities, what is the shortest possible route that visits each city exactly once and returns to the origin city?" It is an NP-hard problem in combinatorial optimization, important in theoretical computer science and operations research.
\.


--
-- Data for Name: Session; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Session" (id, id_user, is_deleted) FROM stdin;
3	6	t
4	7	f
5	8	t
6	10	f
7	11	f
8	12	f
9	12	f
10	12	f
11	12	f
12	12	f
13	12	f
14	12	f
15	12	f
16	12	f
17	12	f
18	12	f
19	12	f
20	12	f
21	12	f
22	12	f
23	12	f
24	12	f
25	12	f
26	12	f
27	12	f
28	12	f
29	12	f
30	12	f
\.


--
-- Data for Name: Solver; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Solver" (id, name, description, model, id_issue) FROM stdin;
1	Ant	In computer science and operations research, the ant colony optimization algorithm (ACO) is a probabilistic technique for solving computational problems which can be reduced to finding good paths through graphs.	\N	1
\.


--
-- Data for Name: Task; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Task" (id, name, create_date, description, "id_GA", id_user, id_solver) FROM stdin;
\.


--
-- Data for Name: Template; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Template" (id, length) FROM stdin;
\.


--
-- Data for Name: User; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."User" (id, login, pass) FROM stdin;
1	Bill	sasasas
2	Test	123456
3	Test2	123456
4	HelloBoy	123456
5	HelloBoy2	123456
6	HelloBoy3	123456
7	Arrr	123456
8	bobaAndBiba	123456
10	adgsthrgreg	arggsegeagfraef
11	argargaegeargae	rgaegaegergegae
12	lolkek	topkek
\.


--
-- Name: Agent_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Agent_id_seq"', 1, false);


--
-- Name: Geb_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Geb_id_seq"', 1, false);


--
-- Name: GenTemplate_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."GenTemplate_id_seq"', 1, false);


--
-- Name: Generation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Generation_id_seq"', 1, false);


--
-- Name: Issue_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Issue_id_seq"', 1, true);


--
-- Name: Session_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Session_id_seq"', 1, false);


--
-- Name: Session_id_seq1; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Session_id_seq1"', 30, true);


--
-- Name: Template_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Template_id_seq"', 1, false);


--
-- Name: User_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."User_id_seq"', 12, true);


--
-- Name: ga_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ga_id_seq', 1, false);


--
-- Name: solver_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.solver_id_seq', 1, true);


--
-- Name: Agent agent_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Agent"
    ADD CONSTRAINT agent_pk PRIMARY KEY (id);


--
-- Name: GeneticAlgorithm ga_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."GeneticAlgorithm"
    ADD CONSTRAINT ga_pk PRIMARY KEY (id);


--
-- Name: Gen geb_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Gen"
    ADD CONSTRAINT geb_pk PRIMARY KEY (id);


--
-- Name: Generation generation_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Generation"
    ADD CONSTRAINT generation_pk PRIMARY KEY (id);


--
-- Name: GenTemplate gentemplate_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."GenTemplate"
    ADD CONSTRAINT gentemplate_pk PRIMARY KEY (id);


--
-- Name: Issue issue_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Issue"
    ADD CONSTRAINT issue_pk PRIMARY KEY (id);


--
-- Name: Session session_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT session_pk PRIMARY KEY (id);


--
-- Name: Solver solver_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Solver"
    ADD CONSTRAINT solver_pk PRIMARY KEY (id);


--
-- Name: Task task_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task"
    ADD CONSTRAINT task_pk PRIMARY KEY (id);


--
-- Name: Template template_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Template"
    ADD CONSTRAINT template_pk PRIMARY KEY (id);


--
-- Name: User user_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT user_pk PRIMARY KEY (id);


--
-- Name: ga_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX ga_id_uindex ON public."GeneticAlgorithm" USING btree (id);


--
-- Name: ga_name_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX ga_name_uindex ON public."GeneticAlgorithm" USING btree (name);


--
-- Name: geb_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX geb_id_uindex ON public."Gen" USING btree (id);


--
-- Name: generation_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX generation_id_uindex ON public."Generation" USING btree (id);


--
-- Name: gentemplate_id_gen_id_template_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX gentemplate_id_gen_id_template_uindex ON public."GenTemplate" USING btree (id_gen, id_template);


--
-- Name: gentemplate_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX gentemplate_id_uindex ON public."GenTemplate" USING btree (id);


--
-- Name: issue_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX issue_id_uindex ON public."Issue" USING btree (id);


--
-- Name: issue_name_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX issue_name_uindex ON public."Issue" USING btree (name);


--
-- Name: session_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX session_id_uindex ON public."Session" USING btree (id);


--
-- Name: solver_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX solver_id_uindex ON public."Solver" USING btree (id);


--
-- Name: solver_name_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX solver_name_uindex ON public."Solver" USING btree (name);


--
-- Name: task_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX task_id_uindex ON public."Task" USING btree (id);


--
-- Name: template_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX template_id_uindex ON public."Template" USING btree (id);


--
-- Name: user_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_id_uindex ON public."User" USING btree (id);


--
-- Name: user_login_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_login_uindex ON public."User" USING btree (login);


--
-- Name: Agent agent_agent_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Agent"
    ADD CONSTRAINT agent_agent_id_fk FOREIGN KEY (id_parent) REFERENCES public."Agent"(id);


--
-- Name: Agent agent_generation_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Agent"
    ADD CONSTRAINT agent_generation_id_fk FOREIGN KEY (id_generation) REFERENCES public."Generation"(id);


--
-- Name: Agent agent_template_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Agent"
    ADD CONSTRAINT agent_template_id_fk FOREIGN KEY (id_template) REFERENCES public."Template"(id);


--
-- Name: Gen gen_agent_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Gen"
    ADD CONSTRAINT gen_agent_id_fk FOREIGN KEY (id_agent) REFERENCES public."Agent"(id);


--
-- Name: Generation generation_session_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Generation"
    ADD CONSTRAINT generation_session_id_fk FOREIGN KEY (id_session) REFERENCES public."Task"(id);


--
-- Name: GenTemplate gentemplate_gen_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."GenTemplate"
    ADD CONSTRAINT gentemplate_gen_id_fk FOREIGN KEY (id_gen) REFERENCES public."Gen"(id);


--
-- Name: GenTemplate gentemplate_template_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."GenTemplate"
    ADD CONSTRAINT gentemplate_template_id_fk FOREIGN KEY (id_template) REFERENCES public."Template"(id);


--
-- Name: Task session_geneticalgorithm_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task"
    ADD CONSTRAINT session_geneticalgorithm_id_fk FOREIGN KEY ("id_GA") REFERENCES public."GeneticAlgorithm"(id);


--
-- Name: Session session_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT session_user_id_fk FOREIGN KEY (id_user) REFERENCES public."User"(id);


--
-- Name: Solver solver_issue_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Solver"
    ADD CONSTRAINT solver_issue_id_fk FOREIGN KEY (id_issue) REFERENCES public."Issue"(id);


--
-- Name: Task task_solver_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task"
    ADD CONSTRAINT task_solver_id_fk FOREIGN KEY (id_solver) REFERENCES public."Solver"(id);


--
-- Name: Task task_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task"
    ADD CONSTRAINT task_user_id_fk FOREIGN KEY (id_user) REFERENCES public."User"(id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

