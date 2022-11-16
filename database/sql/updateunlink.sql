UPDATE players 
SET discord_id = NULL,
link_code = NULL,
count_unlink = count_unlink + 1
WHERE steam_id = ?;