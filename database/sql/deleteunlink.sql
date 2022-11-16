-- Delete lines in `gloves`, `weapons` and `weapons_timestamps`
DELETE FROM `gloves`
WHERE steamid = ?;

DELETE FROM `weapons`
WHERE steamid = ?;

DELETE FROM `weapons_timestamps`
WHERE steamid = ?;