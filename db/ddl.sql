---------------------------------------------
---------------------------------------------
-------------------- app --------------------
CREATE database app; 

-- app.states:
-- a table to associate states id with names
-- e.g.: RJ - Rio de Janeiro
CREATE TABLE IF NOT EXISTS app.states (
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_name          VARCHAR(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (state_id)
);

-- app.registered_users
-- a table to store users who registered in our plataform
-- stores the wca id used, the registered state and when he did registered
CREATE TABLE IF NOT EXISTS app.registered_users (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    register_date       DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (wca_id),
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE
);

---------------------------------------------
---------------------------------------------
------------------ datalake -----------------
CREATE DATABASE datalake;

-- datalake.export_date
-- a table to store when was the last update of the datalake
CREATE TABLE IF NOT EXISTS datalake.export_date (
    last_update         DATETIME,
    PRIMARY KEY (last_update)
);

-- datalake.competitors
-- a table to store the competitiors' name and id
CREATE TABLE IF NOT EXISTS datalake.competitors (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    wca_name            VARCHAR(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (wca_id)
);

-- datalake.competitions
-- a table to store which competition happened in which state
CREATE TABLE IF NOT EXISTS datalake.competitions (
    competition_id      VARCHAR(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (competition_id),
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE
);

-- datalake.competitions_by_person_and_state
-- a table to store how much competitions the competitor competed in each state
CREATE TABLE IF NOT EXISTS datalake.competitions_by_person_and_state (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    n_competitions      INT NOT NULL,
    PRIMARY KEY (wca_id, state_id),
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE
);

-- datalake.estimated_state_for_user
-- a table to store the estimated/guessed state for the competitor
CREATE TABLE IF NOT EXISTS datalake.estimated_state_for_user (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (wca_id),
    FOREIGN KEY (wca_id) REFERENCES datalake.competitors(wca_id) ON DELETE CASCADE,
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE
);

-- datalake.all_persons_with_states
-- a table to join the estimated/guessed state with the registered state to each competitor
CREATE TABLE IF NOT EXISTS datalake.all_persons_with_states (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (wca_id),
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE
);

-- datalake.ranking_single
-- a table to store the current ranking for single results
-- the ranking is for (user, state, event)
CREATE TABLE IF NOT EXISTS datalake.ranking_single (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    event_id            VARCHAR(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    single              INT(11),
    ranking             INT NOT NULL,
    PRIMARY KEY (wca_id, event_id),
    FOREIGN KEY (wca_id) REFERENCES datalake.competitors(wca_id) ON DELETE CASCADE,
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE,
    INDEX (ranking)
);

-- datalake.ranking_average
-- a table to store the current ranking for average results
-- the ranking is for (user, state, event)
CREATE TABLE IF NOT EXISTS datalake.ranking_average (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    event_id            VARCHAR(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    average             INT(11),
    ranking             INT NOT NULL,
    PRIMARY KEY (wca_id, event_id),
    FOREIGN KEY (wca_id) REFERENCES datalake.competitors(wca_id) ON DELETE CASCADE,
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE,
    INDEX (ranking)
);

-- datalake.sum_of_ranks
-- a table to store the state of all other ranks summed
-- the ranking is for (user, state, event)
CREATE TABLE IF NOT EXISTS datalake.sum_of_ranks (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    sum_single          INT,
    sum_average         INT,
    sum_sum             INT,
    PRIMARY KEY (wca_id),
    FOREIGN KEY (wca_id) REFERENCES datalake.competitors(wca_id) ON DELETE CASCADE
);

-- datalake.ranking_sum
-- a table to store the current ranking for sum of ranks
CREATE TABLE IF NOT EXISTS datalake.ranking_sum (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    ranking_single      INT NOT NULL,
    ranking_average     INT NOT NULL,
    PRIMARY KEY (wca_id),
    FOREIGN KEY (wca_id) REFERENCES datalake.competitors(wca_id) ON DELETE CASCADE,
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE,
    INDEX (ranking_single),
    INDEX (ranking_average)
);

---------------------------------------------
---------------------------------------------
------------------ datalake -----------------
CREATE DATABASE dump;

-- dump.competitions_by_person_and_country
-- a temporary table to list how many competitions by country a competitor have
CREATE TABLE IF NOT EXISTS dump.competitions_by_person_and_country (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    country_name        VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    n_competitions      INT,
    PRIMARY KEY (wca_id, country_name),
    FOREIGN KEY (wca_id) REFERENCES datalake.competitors(wca_id) ON DELETE CASCADE,
    INDEX (country_name)
);

-- dump.results_by_state
-- a temporary table to group results of a competitor with its state
CREATE TABLE IF NOT EXISTS dump.results_by_state (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    event_id            VARCHAR(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    average             INT(11),
    single              INT(11),
    PRIMARY KEY (wca_id, state_id, event_id),
    FOREIGN KEY (wca_id) REFERENCES datalake.competitors(wca_id) ON DELETE CASCADE,
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE,
    INDEX (single),
    INDEX (average)
);

---------------------------------------------
---------------------------------------------
----------------- procedure -----------------
DELIMITER //

CREATE PROCEDURE IF NOT EXISTS app.update_after_insert ()
LANGUAGE SQL
NOT DETERMINISTIC
MODIFIES SQL DATA
BEGIN

-- assuming dump has been filled with WCA sql dump.

-- set string encode
ALTER DATABASE dump CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- applying encoding to tables
ALTER TABLE dump.Competitions CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
ALTER TABLE dump.Persons CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
ALTER TABLE dump.Results CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
ALTER TABLE dump.Events CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
ALTER TABLE dump.RanksSingle CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
ALTER TABLE dump.RanksAverage CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- creating indexes on dump tables
-- TODO: CHECK USAGE OF ALL INDEXES
CREATE INDEX IF NOT EXISTS idx_comp_id ON dump.Competitions (id);
CREATE INDEX IF NOT EXISTS idx_comp_countryId ON dump.Competitions (countryId);

CREATE INDEX IF NOT EXISTS idx_res_eventId ON dump.Results (eventId);
CREATE INDEX IF NOT EXISTS idx_res_best ON dump.Results (best);
CREATE INDEX IF NOT EXISTS idx_res_personId ON dump.Results (personId);
CREATE INDEX IF NOT EXISTS idx_res_competitionId ON dump.Results (competitionId);

CREATE INDEX IF NOT EXISTS idx_ra_personId ON dump.RanksAverage (personId);
CREATE INDEX IF NOT EXISTS idx_ra_eventId ON dump.RanksAverage (eventId);
CREATE INDEX IF NOT EXISTS idx_rs_personId ON dump.RanksSingle (personId);
CREATE INDEX IF NOT EXISTS idx_rs_eventId ON dump.RanksSingle (eventId);

CREATE INDEX IF NOT EXISTS idx_per_id ON dump.Persons (id);
CREATE INDEX IF NOT EXISTS idx_per_country ON dump.Persons (countryId);

-- filling competitions
REPLACE INTO datalake.competitions(
    competition_id,
    state_id
)
    SELECT
        c.id            AS competition_id,
        s.state_id      AS state_id
    FROM
        (
            SELECT
                c.id AS id,
                SUBSTRING_INDEX(c.cityName, ', ', -1) AS name
            FROM dump.Competitions c
            WHERE c.countryId = 'Brazil'
        ) AS c 
        LEFT JOIN 
            app.states s ON c.name = s.state_name
        WHERE s.state_id IS NOT NULL
;

REPLACE INTO datalake.competitors(
    wca_id,
    wca_name
)
    SELECT
        p.id                        AS wca_id,
        p.name                      AS wca_name
    FROM
        dump.Persons p
    WHERE
        p.countryId = 'Brazil'
;

-- loading number of competitions each person in each state
REPLACE INTO datalake.competitions_by_person_and_state (
    wca_id,
    state_id,
    n_competitions
)
    SELECT
        p.wca_id                    AS wca_id, 
        lake_c.state_id             AS state_id,
        COUNT(DISTINCT dump_c.id)   AS n_competitions
    FROM
        datalake.competitors p,
        dump.Results r,
        datalake.competitions lake_c,
        dump.Competitions dump_c
    WHERE
        p.wca_id = r.personId AND
        dump_c.id = r.competitionId AND
        lake_c.competition_id = dump_c.id
    GROUP BY p.wca_id, lake_c.state_id
    -- ON DUPLICATE KEY UPDATE (n_competitions)
;

REPLACE INTO dump.competitions_by_person_and_country (
    wca_id,
    country_name,
    n_competitions
)
    SELECT
        p.wca_id                    AS wca_id,
        c.countryId                 AS country_name,
        COUNT(DISTINCT c.id)        AS n_competitions
    FROM
        datalake.competitors p,
        dump.Competitions c,
        dump.Results r
    WHERE
        p.wca_id = r.personId AND
        c.id = r.competitionId
    GROUP BY p.wca_id, c.countryId
    -- ON DUPLICATE KEY UPDATE (n_competitions)
;

REPLACE INTO datalake.estimated_state_for_user (
    wca_id,
    state_id
)
    SELECT
        p.wca_id                    AS wca_id,
        ces_max_state.state_id      AS state_id
    FROM
        datalake.competitors p
            LEFT JOIN (
                SELECT wca_id, state_id
                FROM datalake.competitions_by_person_and_state c1
                WHERE n_competitions = (
                    SELECT MAX(n_competitions)
                    FROM datalake.competitions_by_person_and_state c2
                    WHERE c1.wca_id = c2.wca_id
                )
            ) AS ces_max_state ON p.wca_id = ces_max_state.wca_id
            LEFT JOIN  (
                SELECT wca_id, country_name
                FROM dump.competitions_by_person_and_country c1
                WHERE n_competitions = (
                    SELECT MAX(n_competitions)
                    FROM dump.competitions_by_person_and_country c2
                    WHERE c1.wca_id = c2.wca_id
                )
            ) AS cec_max_country on p.wca_id = cec_max_country.wca_id
    WHERE
        cec_max_country.country_name = 'Brazil' AND
        ces_max_state.state_id IS NOT NULL
    -- ON DUPLICATE KEY UPDATE (state_id)
;


-- REPLACE INTO app.registered_users (
--     wca_id,
--     state_id
-- )
--     SELECT
--         wca_id,
--         state_id
--     FROM
--         app.next_update_users
-- ;

-- TRUNCATE TABLE app.next_update_users;


TRUNCATE TABLE dump.competitions_by_person_and_country;


END //
--------------------------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------------------------------------------------

CREATE PROCEDURE IF NOT EXISTS app.refresh ()
LANGUAGE SQL
NOT DETERMINISTIC
MODIFIES SQL DATA
BEGIN

REPLACE INTO datalake.all_persons_with_states (
    wca_id,
    state_id
)
    SELECT
        es.wca_id                           AS wca_id,
        COALESCE(re.state_id, es.state_id)  AS state_id
    FROM
        datalake.estimated_state_for_user es
        LEFT JOIN app.registered_users re
            ON es.wca_id = re.wca_id
;


REPLACE INTO dump.results_by_state (
    wca_id,
    state_id,
    event_id,
    single,
    average
)
    SELECT
        al.wca_id           AS wca_id,
        al.state_id         AS state_id,
        rs.eventId          AS event_id,
        rs.best             AS single,
        ra.best             AS average
    FROM
        datalake.all_persons_with_states al
            LEFT JOIN dump.RanksSingle rs
                ON al.wca_id = rs.personId
            LEFT JOIN dump.RanksAverage ra
                ON rs.personId = ra.personId
                AND rs.eventId = ra.eventId
    WHERE al.state_id IS NOT NULL
    AND rs.eventId IS NOT NULL
;

-- remove old values if the state has changed
-- as the replace doesnt remove these old values
DELETE
FROM dump.results_by_state rs
WHERE EXISTS (
    SELECT wca_id, state_id
    FROM datalake.all_persons_with_states al
    WHERE
        rs.wca_id = al.wca_id
        AND rs.state_id != al.state_id
);

REPLACE INTO datalake.ranking_single (
    wca_id,
    event_id,
    state_id,
    single,
    ranking
)
    SELECT
        rs1.wca_id              AS wca_id,
        rs1.event_id            AS event_id,
        rs1.state_id            AS state_id,
        rs1.single              AS single,
        dense_rank() OVER (
            PARTITION BY rs2.event_id, rs2.state_id ORDER BY rs2.single ASC
        ) AS ranking
    FROM
        dump.results_by_state rs1
            LEFT JOIN (
                SELECT
                    wca_id,
                    event_id,
                    state_id,
                    -- COALESCE(NULLIF(NULLIF(NULLIF(single,0),-1),-2), 9999999999) AS single
                    CASE WHEN COALESCE(single, 0) <= 0 THEN 9999999999 ELSE single END AS single
                FROM
                    dump.results_by_state
            ) AS rs2
                ON rs1.wca_id = rs2.wca_id 
                    AND rs1.event_id = rs2.event_id 
                    AND rs1.state_id = rs2.state_id
;

REPLACE INTO datalake.ranking_average(
    wca_id,
    state_id,
    event_id,
    average,
    ranking
)
    SELECT
        rs1.wca_id          AS wca_id,
        rs1.state_id        AS state_id,
        rs1.event_id        AS event_id,
        rs1.average         AS average,
        dense_rank() OVER (
            PARTITION BY rs2.event_id, rs2.state_id ORDER BY rs2.average ASC
        ) AS ranking
    FROM
        dump.results_by_state rs1
            LEFT JOIN (
                SELECT
                    wca_id,
                    event_id,
                    state_id,
                    -- COALESCE(NULLIF(NULLIF(NULLIF(average,0),-1),-2), 9999999999) AS average
                    CASE WHEN COALESCE(average, 0) <= 0 THEN 9999999999 ELSE average END AS average
                FROM
                    dump.results_by_state
            ) AS rs2
                ON rs1.wca_id = rs2.wca_id
                    AND rs1.event_id = rs2.event_id
                    AND rs1.state_id = rs2.state_id
;

REPLACE INTO datalake.sum_of_ranks(
    wca_id,
    sum_single,
    sum_average,
    sum_sum
)
    SELECT
        cc.wca_id,
        ss.sum_single,
        sa.sum_average,
        ss.sum_single + sa.sum_average sum_sum
    FROM
        datalake.competitors cc
            LEFT JOIN (
                SELECT
                    c.wca_id,
                    SUM(rs.ranking) sum_single
                FROM
                    datalake.competitors c
                        RIGHT JOIN datalake.ranking_single rs ON c.wca_id = rs.wca_id
                GROUP BY
                    wca_id 
            ) AS ss on cc.wca_id = ss.wca_id

            LEFT JOIN (
                SELECT
                    c.wca_id,
                    SUM(r.ranking) sum_average
                FROM
                    datalake.competitors c
                        RIGHT JOIN datalake.ranking_average r ON c.wca_id = r.wca_id
                GROUP BY
                    wca_id
            ) AS sa on cc.wca_id = sa.wca_id
;

-- SELECT
--         rs1.wca_id          AS wca_id,
--         rs1.state_id        AS state_id,
--         rs1.event_id        AS event_id,
--         NULLIF(NULLIF(NULLIF(rs1.average, 0),-1),-2)    AS average,
--         dense_rank() OVER (
--             PARTITION BY rs2.event_id, rs2.state_id ORDER BY rs2.average ASC
--         ) AS ranking

-- REPLACE INTO datalake.ranking_sum(
--     wca_id,
--     state_id,
--     ranking_single,
--     ranking_average
-- )
--     SELECT
--     FROM
--         datalake.all_persons_with_states al
--             LEFT JOIN (
--                 SELECT
--                     sr.wca_id,
--                     dense_rank() OVER (
--                         PARTITION BY sr.wca_id ORDER BY sr.sum_single ASC
--                     ) AS ranking
--                 FROM
--                     dalake.sum_of_ranks sr
--             )

-- ;

END //
