DELETE gloves, weapons, weapons_timestamps
FROM gloves
JOIN weapons ON (weapons.steamid = gloves.steamid)
JOIN weapons_timestamps ON (weapons_timestamps.steamid = weapons.steamid)
WHERE gloves.steamid = ?;