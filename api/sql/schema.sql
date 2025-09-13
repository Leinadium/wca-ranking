CREATE DATABASE app;
CREATE DATABASE datalake;
CREATE DATABASE dump;

CREATE TABLE app.states (
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_name          VARCHAR(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (state_id)
);

CREATE TABLE app.registered_users (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    register_date       DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (wca_id),
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE
);

CREATE TABLE datalake.export_date (
    last_update         DATETIME,
    PRIMARY KEY (last_update)
);

CREATE TABLE datalake.competitors (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    wca_name            VARCHAR(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (wca_id)
);

CREATE TABLE datalake.competitions (
    competition_id      VARCHAR(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (competition_id),
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE
);

CREATE TABLE datalake.competitions_by_person_and_state (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    n_competitions      INT NOT NULL,
    PRIMARY KEY (wca_id, state_id),
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE
);

CREATE TABLE datalake.estimated_state_for_user (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (wca_id),
    FOREIGN KEY (wca_id) REFERENCES datalake.competitors(wca_id) ON DELETE CASCADE,
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE
);

CREATE TABLE datalake.all_persons_with_states (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    state_id            CHAR(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (wca_id),
    FOREIGN KEY (state_id) REFERENCES app.states(state_id) ON DELETE CASCADE
);

CREATE TABLE datalake.ranking_single (
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

CREATE TABLE datalake.ranking_average (
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

CREATE TABLE datalake.sum_of_ranks (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    sum_single          INT,
    sum_average         INT,
    sum_sum             INT,
    PRIMARY KEY (wca_id),
    FOREIGN KEY (wca_id) REFERENCES datalake.competitors(wca_id) ON DELETE CASCADE
);

CREATE TABLE datalake.ranking_sum (
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

CREATE TABLE dump.competitions_by_person_and_country (
    wca_id              VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    country_name        VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    n_competitions      INT,
    PRIMARY KEY (wca_id, country_name),
    FOREIGN KEY (wca_id) REFERENCES datalake.competitors(wca_id) ON DELETE CASCADE,
    INDEX (country_name)
);

CREATE TABLE dump.results_by_state (
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

CREATE TABLE dump.Results (
    personId            VARCHAR(10) NOT NULL,
    eventId             VARCHAR(6) NOT NULL,
    competitionId       VARCHAR(20) NOT NULL,
    value1              INT(11),
    value2              INT(11),
    value3              INT(11),
    value4              INT(11),
    value5              INT(11),
    best                INT(11)
);

CREATE TABLE dump.Competitions (
    id      VARCHAR(10) NOT NULL,
    name    VARCHAR(80) NOT NULL
);