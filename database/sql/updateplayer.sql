UPDATE players
SET discord_id = ?,
link_code = NULL
WHERE
steam_id = ?
AND
link_code = ?;