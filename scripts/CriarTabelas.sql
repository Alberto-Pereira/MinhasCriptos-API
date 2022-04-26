--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

-- Started on 2022-04-20 17:44:48

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
-- TOC entry 4 (class 2615 OID 24618)
-- Name: minhascriptosprincipal; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA minhascriptosprincipal;


ALTER SCHEMA minhascriptosprincipal OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 214 (class 1259 OID 24637)
-- Name: criptomoeda; Type: TABLE; Schema: minhascriptosprincipal; Owner: postgres
--

CREATE TABLE minhascriptosprincipal.criptomoeda (
    id integer NOT NULL,
    "tipoMoeda" text NOT NULL,
    "dataDeCompra" date NOT NULL,
    "quantidadeComprada" double precision NOT NULL,
    "precoDeCompra" double precision NOT NULL,
    "valorDaUnidadeNoDiaDeCompra" double precision NOT NULL,
    usuario_id integer NOT NULL
);


ALTER TABLE minhascriptosprincipal.criptomoeda OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 24635)
-- Name: criptomoeda_id_seq; Type: SEQUENCE; Schema: minhascriptosprincipal; Owner: postgres
--

CREATE SEQUENCE minhascriptosprincipal.criptomoeda_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE minhascriptosprincipal.criptomoeda_id_seq OWNER TO postgres;

--
-- TOC entry 3323 (class 0 OID 0)
-- Dependencies: 212
-- Name: criptomoeda_id_seq; Type: SEQUENCE OWNED BY; Schema: minhascriptosprincipal; Owner: postgres
--

ALTER SEQUENCE minhascriptosprincipal.criptomoeda_id_seq OWNED BY minhascriptosprincipal.criptomoeda.id;


--
-- TOC entry 213 (class 1259 OID 24636)
-- Name: criptomoeda_usuario_id_seq; Type: SEQUENCE; Schema: minhascriptosprincipal; Owner: postgres
--

CREATE SEQUENCE minhascriptosprincipal.criptomoeda_usuario_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE minhascriptosprincipal.criptomoeda_usuario_id_seq OWNER TO postgres;

--
-- TOC entry 3324 (class 0 OID 0)
-- Dependencies: 213
-- Name: criptomoeda_usuario_id_seq; Type: SEQUENCE OWNED BY; Schema: minhascriptosprincipal; Owner: postgres
--

ALTER SEQUENCE minhascriptosprincipal.criptomoeda_usuario_id_seq OWNED BY minhascriptosprincipal.criptomoeda.usuario_id;


--
-- TOC entry 211 (class 1259 OID 24627)
-- Name: usuario; Type: TABLE; Schema: minhascriptosprincipal; Owner: postgres
--

CREATE TABLE minhascriptosprincipal.usuario (
    id integer NOT NULL,
    nome text NOT NULL,
    email text NOT NULL,
    senha text NOT NULL
);


ALTER TABLE minhascriptosprincipal.usuario OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 24626)
-- Name: usuario_id_seq; Type: SEQUENCE; Schema: minhascriptosprincipal; Owner: postgres
--

CREATE SEQUENCE minhascriptosprincipal.usuario_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE minhascriptosprincipal.usuario_id_seq OWNER TO postgres;

--
-- TOC entry 3325 (class 0 OID 0)
-- Dependencies: 210
-- Name: usuario_id_seq; Type: SEQUENCE OWNED BY; Schema: minhascriptosprincipal; Owner: postgres
--

ALTER SEQUENCE minhascriptosprincipal.usuario_id_seq OWNED BY minhascriptosprincipal.usuario.id;


--
-- TOC entry 3172 (class 2604 OID 24640)
-- Name: criptomoeda id; Type: DEFAULT; Schema: minhascriptosprincipal; Owner: postgres
--

ALTER TABLE ONLY minhascriptosprincipal.criptomoeda ALTER COLUMN id SET DEFAULT nextval('minhascriptosprincipal.criptomoeda_id_seq'::regclass);


--
-- TOC entry 3173 (class 2604 OID 24641)
-- Name: criptomoeda usuario_id; Type: DEFAULT; Schema: minhascriptosprincipal; Owner: postgres
--

ALTER TABLE ONLY minhascriptosprincipal.criptomoeda ALTER COLUMN usuario_id SET DEFAULT nextval('minhascriptosprincipal.criptomoeda_usuario_id_seq'::regclass);


--
-- TOC entry 3171 (class 2604 OID 24630)
-- Name: usuario id; Type: DEFAULT; Schema: minhascriptosprincipal; Owner: postgres
--

ALTER TABLE ONLY minhascriptosprincipal.usuario ALTER COLUMN id SET DEFAULT nextval('minhascriptosprincipal.usuario_id_seq'::regclass);


--
-- TOC entry 3175 (class 2606 OID 24658)
-- Name: usuario email; Type: CONSTRAINT; Schema: minhascriptosprincipal; Owner: postgres
--

ALTER TABLE ONLY minhascriptosprincipal.usuario
    ADD CONSTRAINT email UNIQUE (email);


--
-- TOC entry 3177 (class 2606 OID 24634)
-- Name: usuario usuario_pkey; Type: CONSTRAINT; Schema: minhascriptosprincipal; Owner: postgres
--

ALTER TABLE ONLY minhascriptosprincipal.usuario
    ADD CONSTRAINT usuario_pkey PRIMARY KEY (id);


--
-- TOC entry 3178 (class 2606 OID 24644)
-- Name: criptomoeda usuario_id; Type: FK CONSTRAINT; Schema: minhascriptosprincipal; Owner: postgres
--

ALTER TABLE ONLY minhascriptosprincipal.criptomoeda
    ADD CONSTRAINT usuario_id FOREIGN KEY (usuario_id) REFERENCES minhascriptosprincipal.usuario(id);


-- Completed on 2022-04-20 17:44:49

--
-- PostgreSQL database dump complete
--

