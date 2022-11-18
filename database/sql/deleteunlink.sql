DELETE FROM gloves
WHERE steamid = ?;

DELETE FROM weapons
WHERE steamid = ?;

DELETE FROM weapons_timestamps
WHERE steamid = ?;