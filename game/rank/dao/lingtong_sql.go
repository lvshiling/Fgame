package dao

const (
	lingTongDevSql = `SELECT 
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
						A.advancedId >= 0
					AND 
						A.classType = B.classType
					ORDER  BY 
						  A.advancedId DESC, 
						  B.power DESC,
						  D.level DESC
					LIMIT ?`
)

const (
	lingTongLevelRankSql = `SELECT
						B.serverId,
						A.playerId,
						B.name,
						A.level AS num,
						A.power 
					FROM
						t_player_lingtong A
						INNER JOIN t_player B ON A.playerId = B.id 
					WHERE
						B.serverId =?
						AND A.power > 0 
						AND A.level >= 0
					ORDER BY
						A.level DESC,
						A.power DESC 
						LIMIT ?`
)
