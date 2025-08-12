-------------------------------
-- name: GetPersonSingle :many
SELECT
    dlk.wca_id      AS wca_id,
    ct.wca_name     AS name,
    al.state_id     AS state_id,
    CASE WHEN ru.wca_id is not null THEN true ELSE false END AS registered,
    dlk.event_id    AS event_id,
    dlk.ranking     AS ranking,
    dlk.single      AS best,
    comp.id         AS competition_id,
    comp.name       AS competition_name,
    comp2.state_id  AS competition_state,
    -- dmp.roundTypeId AS round,
    dmp.value1      AS time_1,
    dmp.value2      AS time_2,
    dmp.value3      AS time_3,
    dmp.value4      AS time_4,
    dmp.value5      AS time_5,
    STR_TO_DATE(CONCAT(comp.year, ',', comp.endMonth, ',', comp.endDay), '%Y,%m,%d') AS ts
FROM
    datalake.ranking_single dlk
        LEFT JOIN datalake.competitors ct on (dlk.wca_id = ct.wca_id)
        LEFT JOIN datalake.all_persons_with_states al on (dlk.wca_id = al.wca_id)
        LEFT JOIN app.registered_users ru on (dlk.wca_id = ru.wca_id)
        LEFT JOIN dump.Results dmp on (dlk.wca_id = dmp.personId and dlk.event_id = dmp.eventId)
        LEFT JOIN dump.Competitions comp on (dmp.competitionId = comp.id)
        LEFT JOIN datalake.competitions comp2 on (dmp.competitionId = comp2.competition_id)
WHERE
    dlk.single = dmp.best
    AND dlk.wca_id = sqlc.arg(wcaId)

-------------------------------
-- name: GetPersonAverage :many
SELECT
    dlk.wca_id      AS wca_id,
    ct.wca_name     AS name,
    al.state_id     AS state_id,
    CASE WHEN ru.wca_id is not null THEN true ELSE false END AS registered,
    dlk.event_id    AS event_id,
    dlk.ranking     AS ranking,
    dlk.average     AS best,
    comp.id         AS competition_id,
    comp.name       AS competition_name,
    comp2.state_id  AS competition_state,
    -- dmp.roundTypeId AS round,
    dmp.value1      AS time_1,
    dmp.value2      AS time_2,
    dmp.value3      AS time_3,
    dmp.value4      AS time_4,
    dmp.value5      AS time_5,
    STR_TO_DATE(CONCAT(comp.year, ',', comp.endMonth, ',', comp.endDay), '%Y,%m,%d') AS ts
FROM
    datalake.ranking_average dlk
        LEFT JOIN datalake.competitors ct on (dlk.wca_id = ct.wca_id)
        LEFT JOIN datalake.all_persons_with_states al on (dlk.wca_id = al.wca_id)
        LEFT JOIN app.registered_users ru on (dlk.wca_id = ru.wca_id)
        LEFT JOIN dump.Results dmp on (dlk.wca_id = dmp.personId and dlk.event_id = dmp.eventId)
        LEFT JOIN dump.Competitions comp on (dmp.competitionId = comp.id)
        LEFT JOIN datalake.competitions comp2 on (dmp.competitionId = comp2.competition_id)
WHERE
    dlk.average = dmp.average
    AND dlk.wca_id = sqlc.arg(wcaId)

-------------------------------
-- name: GetPersonInfo :one
SELECT
    ct.wca_name AS name,
    al.state_id AS state_id,
    CASE WHEN ru.wca_id IS NOT NULL THEN true ELSE false END AS registered,
    cb.n_competitions AS state_competitions,
    cb2.total AS total_competitions
FROM
    datalake.competitors ct
        LEFT JOIN app.registered_users ru ON ru.wca_id = ct.wca_id
        LEFT JOIN datalake.all_persons_with_states al on ct.wca_id = al.wca_id
        LEFT JOIN datalake.competitions_by_person_and_state cb ON cb.wca_id = ct.wca_id
        LEFT JOIN (
            SELECT personId, COUNT(DISTINCT competitionId) AS total
            FROM dump.Results
            WHERE personId = @wcaId
        ) as cb2 ON cb2.personId = ct.wca_id
WHERE
    ct.wca_id = @wcaId

-------------------------------
-- name: GetRankingSingle :many
SELECT
	rs.wca_id 			AS wca_id,
	cpr.wca_name 		AS name,
	rs.state_id 		AS state_id,
	CASE WHEN ru.wca_id is not null THEN true ELSE false END AS registered,
	rs.event_id 		AS event_id,
	rs.ranking 			AS ranking,
	rs.single 			AS best,
	comp.id         	AS competition_id,
    comp.name       	AS competition_name,
    cpn.state_id  		AS competition_state,
    dmp.value1      	AS time_1,
    dmp.value2      	AS time_2,
    dmp.value3      	AS time_3,
    dmp.value4      	AS time_4,
    dmp.value5      	AS time_5,
    STR_TO_DATE(CONCAT(comp.year, ',', comp.endMonth, ',', comp.endDay), '%Y,%m,%d') AS ts
FROM
	datalake.ranking_single rs
		LEFT JOIN datalake.competitors cpr on rs.wca_id = cpr.wca_id
		LEFT JOIN app.registered_users ru on rs.wca_id = ru.wca_id
		LEFT JOIN dump.Results dmp on (rs.wca_id = dmp.personId and rs.event_id = dmp.eventId)
		LEFT JOIN dump.Competitions comp on dmp.competitionId = comp.id
		LEFT JOIN datalake.competitions cpn on dmp.competitionId = cpn.competition_id 
WHERE
	rs.single = dmp.best
	AND rs.state_id = sqlc.arg(stateId)
	AND rs.event_id = sqlc.arg(eventId)
ORDER BY rs.ranking ASC
