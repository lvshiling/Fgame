package dao

const (
	lingTongDevActivitySql = `SELECT 
						C.serverId,
						A.playerId,
						C.name,
						A.advancedId,
						B.power 
					FROM 
						t_player_lingtong_develop A 
					INNER JOIN 
						t_player_lingtong_power B ON A.playerId = B.playerId AND B.classType=?
					INNER JOIN 
					    t_player C ON A.playerId = C.id  AND C.serverId=?
					INNER JOIN 
					    t_player_lingtong D  ON A.playerId=D.playerId
					WHERE 
						B.power >0 
					AND
						A.advancedId >= ?
					AND 
						A.classType = B.classType
					ORDER  BY 
					      A.advancedId DESC, 
						  B.power DESC,
						  D.level DESC
					LIMIT ?`
)
