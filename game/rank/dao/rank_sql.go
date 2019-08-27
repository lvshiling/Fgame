package dao

const (
	forceRankSql = `SELECT 
						A.serverId,
						B.playerId,
						A.name,
						D.name AS gangName,
						B.power,
						B.level,
						A.role,
						A.sex 
				   FROM 
				        t_player A  
				   INNER JOIN 
				        t_player_property B ON A.id=B.playerId AND A.serverId=?
				   INNER JOIN
						t_player_alliance C ON A.id=C.playerId
					LEFT JOIN 
						t_alliance D ON C.allianceId=D.id
				   where 
				        B.power >0 
				   ORDER BY 
				        B.power DESC,
				  		B.level DESC, 
				   		B.playerId ASC 
				   LIMIT ?`
)

const (
	mountRankSql = `SELECT 
						B.serverId,
						A.playerId,
						B.name,
						A.advancedId,
						A.power 
					FROM 
					    t_player_mount A 
					INNER JOIN 
					    t_player B ON A.playerId = B.id  AND B.serverId=?
					INNER JOIN 
					    t_player_property C  ON A.playerId=C.playerId
					WHERE 
						A.power >0 
					AND
						A.advancedId >= 0
					ORDER  BY 
					      A.advancedId DESC, 
						  A.power DESC,
						  C.power DESC 
					LIMIT ?`
)

const (
	wingRankSql = `SELECT 
						B.serverId,
						A.playerId,
						B.name,
						A.advancedId,
						A.power 
					FROM 
						t_player_wing A
					INNER JOIN
						t_player B ON A.playerId = B.id AND B.serverId=?
					INNER JOIN
					    t_player_property C ON A.playerId=C.playerId 
					WHERE 
						A.power >0 
					AND 
						A.advancedId >0 
					ORDER BY 
					    A.advancedId DESC, 
	  					A.power DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	bodyShieldRankSql = `SELECT 
							B.serverId,
							A.playerId,
							B.name,
							A.advancedId,
							A.power 
						FROM 
							t_player_body_shield A 
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN  
							t_player_property C ON A.playerId=C.playerId 
						WHERE 
							A.power >0 
						AND 
							A.advancedId >0 
						ORDER  BY 
							A.advancedId DESC, 
  							A.power DESC,
  							C.power DESC 
						LIMIT ?`
)

const (
	weaponRankSql = `SELECT 
						B.serverId,
						A.playerId,
						B.name,
						A.star,
						A.weaponWear,
						A.power,
						B.role,
						B.sex 
					FROM 
						t_player_weapon_info A
					INNER JOIN
						t_player B ON A.playerId = B.id AND B.serverId=?
					INNER JOIN 
						t_player_property C  ON A.playerId=C.playerId
					WHERE 
					    A.power >0  
					ORDER BY 
					    A.star DESC, 
						A.power DESC,
						C.power DESC 
					LIMIT ?`
)

const (
	gangRankSql = ` SELECT 
						  B.serverId,
						  A.id,
						  A.name AS gangName ,
						  B.name As leadName,
						  A.mengzhuId ,
						  A.totalForce,
						  B.role,
						  B.sex
					FROM 
						t_alliance A
					INNER JOIN 
						t_player B ON A.mengzhuId = B.id  AND B.serverId=?
					WHERE 
						A.deleteTime=0 
					AND 
						A.totalForce >0 
					ORDER BY 
						A.totalForce DESC, 
						A.id ASC 
					LIMIT ?`
)

const (
	shenFaRankSql = `SELECT 
							B.serverId,
							A.playerId,
							B.name,
							A.advancedId,
							A.power 
						FROM 
							t_player_shenfa A
						INNER JOIN 
							t_player B ON  A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON  A.playerId=C.playerId 
						WHERE
							A.power >0 
						AND 
							A.advancedId > 0
						ORDER BY 
							A.advancedId DESC, 
  							A.power DESC,
  							C.power DESC 
						LIMIT ?`
)

const (
	lingYuRankSql = `SELECT 
							B.serverId,
							A.playerId,
							B.name,
							A.advancedId,
							A.power 
						FROM 
							t_player_lingyu A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN	
							t_player_property C ON A.playerId=C.playerId 
						WHERE 
							A.power >0 
						AND A.advancedId >0 
						ORDER BY 
							A.advancedId DESC, 
  							A.power DESC,
  							C.power DESC 
						LIMIT ?`
)

const (
	featherRankSql = `SELECT 
						B.serverId,
						A.playerId,
						B.name,
						A.featherId AS advancedId,
						A.fpower AS power
					FROM 
						t_player_wing A
					INNER JOIN
						t_player B ON A.playerId = B.id AND B.serverId=?
					INNER JOIN
						t_player_property C ON A.playerId=C.playerId 
					WHERE 
						A.fpower >0 
					AND 
						A.featherId >0 
					ORDER BY 
						A.featherId DESC, 
	  					A.fpower DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	shieldRankSql = `SELECT 
							B.serverId,
							A.playerId,
							B.name,
							A.shieldId AS advancedId,
							A.spower AS power
						FROM 
							t_player_body_shield A
						INNER JOIN
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN
							t_player_property C ON A.playerId=C.playerId 
						WHERE
							 A.spower >0 
						AND 
							A.shieldId >0 
						ORDER  BY 
							A.shieldId DESC, 
  							A.spower DESC,
  							C.power DESC 
						LIMIT ?`
)
const (
	chargeRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.goldNum AS num,
							B.name,
							C.power 
						FROM
							t_player_open_activity_charge A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
						AND 
							C.power > 0 
						ORDER BY
							A.goldNum DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	lingTongForceRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.basePower AS num,
							B.name,
							C.power 
						FROM
							t_player_lingtong A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
						AND 
							C.power > 0 
						ORDER BY
							A.basePower DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	goldEquipForceRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.power AS num,
							B.name,
							C.power 
						FROM
							t_player_goldequip A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
						AND 
							C.power > 0 
						ORDER BY
							A.power DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	dianXingForceRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.power AS num,
							B.name,
							C.power 
						FROM
							t_player_dianxing A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
						AND 
							C.power > 0 
						ORDER BY
							A.power DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	shenQiForceRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.power AS num,
							B.name,
							C.power 
						FROM
							t_player_shenqi A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
						AND 
							C.power > 0 
						ORDER BY
							A.power DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	mingGeForceRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.power AS num,
							B.name,
							C.power 
						FROM
							t_player_mingge A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
						AND 
							C.power > 0 
						ORDER BY
							A.power DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	shengHenForceRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.power AS num,
							B.name,
							C.power 
						FROM
							t_player_shenghen A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
						AND 
							C.power > 0 
						ORDER BY
							A.power DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	zhenFaForceRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.power AS num,
							B.name,
							C.power 
						FROM
							t_player_zhenfa_power A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
						AND 
							C.power > 0 
						ORDER BY
							A.power DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	tuLongEquipForceRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.power AS num,
							B.name,
							C.power 
						FROM
							t_player_tulong_equip A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
						AND 
							C.power > 0 
						ORDER BY
							A.power DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	babyForceRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.power AS num,
							B.name,
							C.power 
						FROM
							t_player_baby_power A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
						AND 
							C.power > 0 
						ORDER BY
							A.power DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	costRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.goldNum AS num,
							B.name,
							C.power 
						FROM
							t_player_open_activity_cost A
						INNER JOIN 
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN 
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
						AND 
							C.power > 0 
						AND 
							A.goldNum > 0
						ORDER BY
							A.goldNum DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	anQiRankSql = `SELECT 
							B.serverId,
							A.playerId,
							B.name,
							A.advancedId,
							A.power 
						FROM 
							t_player_anqi A
						INNER JOIN
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN
							t_player_property C ON A.playerId=C.playerId 
						WHERE
							 A.power >0 
						AND 
							A.advancedId >0 
						ORDER  BY 
							A.advancedId DESC, 
  							A.power DESC,
  							C.power DESC 
						LIMIT ?`
)

const (
	charmRankSql = `SELECT
						B.serverId,
						A.playerId,
						B.name,
						A.charm AS num,
						A.power 
					FROM
						t_player_property A
						INNER JOIN t_player B ON A.playerId = B.id 
					WHERE
						B.serverId =?
						AND A.power > 0 AND A.charm > 0 
					ORDER BY
						A.charm DESC,
						A.power DESC 
						LIMIT ?`
)

const (
	countRankSql = `SELECT
						B.serverId,
						A.playerId,
						A.times AS num,
						B.name,
						C.power 
					FROM
						t_player_activity_num_record A
						INNER JOIN t_player B ON A.playerId = B.id 
						AND B.serverId =?
						INNER JOIN t_player_property C ON A.playerId = C.playerId 
					WHERE
						A.deleteTime = 0 
						AND A.groupId = ? 
						AND A.times >= ? 
						AND C.power > 0 
					ORDER BY
						A.times DESC,
						C.power DESC,
						A.playerId ASC
						LIMIT ?`
)

const (
	faBaoRankSql = `SELECT 
						B.serverId,
						A.playerId,
						B.name,
						A.advancedId,
						A.power 
					FROM 
						t_player_fabao A
					INNER JOIN
						t_player B ON A.playerId = B.id AND B.serverId=?
					INNER JOIN
					    t_player_property C ON A.playerId=C.playerId 
					WHERE 
						A.power >0 
					AND 
						A.advancedId >0 
					ORDER BY 
					    A.advancedId DESC, 
	  					A.power DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	xianTiRankSql = `SELECT 
						B.serverId,
						A.playerId,
						B.name,
						A.advancedId,
						A.power 
					FROM 
						t_player_xianti A
					INNER JOIN
						t_player B ON A.playerId = B.id AND B.serverId=?
					INNER JOIN
					    t_player_property C ON A.playerId=C.playerId 
					WHERE 
						A.power >0 
					AND 
						A.advancedId >0 
					ORDER BY 
					    A.advancedId DESC, 
	  					A.power DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	levelRankSql = `SELECT
						B.serverId,
						A.playerId,
						B.name,
						A.level AS num,
						A.power 
					FROM
						t_player_property A
						INNER JOIN t_player B ON A.playerId = B.id 
					WHERE
						B.serverId =?
						AND A.power > 0 
						AND A.level >= ?
					ORDER BY
						A.level DESC,
						A.power DESC 
						LIMIT ?`
)

const (
	shiHunFanRankSql = `SELECT 
						B.serverId,
						A.playerId,
						B.name,
						A.advancedId,
						A.power 
					FROM 
						t_player_shihunfan A
					INNER JOIN
						t_player B ON A.playerId = B.id AND B.serverId=?
					INNER JOIN
					    t_player_property C ON A.playerId=C.playerId 
					WHERE 
						A.power >0 
					AND 
						A.advancedId >0 
					ORDER BY 
					    A.advancedId DESC, 
	  					A.power DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	tianmoTiRankSql = `SELECT 
						B.serverId,
						A.playerId,
						B.name,
						A.advancedId,
						A.power 
					FROM 
						t_player_tianmo A
					INNER JOIN
						t_player B ON A.playerId = B.id AND B.serverId=?
					INNER JOIN
					    t_player_property C ON A.playerId=C.playerId 
					WHERE 
						A.power >0 
					AND 
						A.advancedId >0 
					ORDER BY 
					    A.advancedId DESC, 
	  					A.power DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	feiShengRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.feiLevel AS num,
							B.name,
							C.power 
						FROM
							t_player_fei_sheng A
							INNER JOIN t_player B ON A.playerId = B.id  AND B.serverId =?
							INNER JOIN t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
							AND A.feiLevel >= ?
							AND C.power > 0 
						ORDER BY
							A.feiLevel DESC,
							C.power DESC,
							A.playerId ASC
							LIMIT ?`
)

const (
	marryDevelopRankSql = `SELECT 
								B.serverId,
								A.playerId,
								A.developLevel AS num,
								B.name,
								C.power 
							FROM
								t_player_marry A
								INNER JOIN t_player B ON A.playerId = B.id  AND B.serverId =?
								INNER JOIN t_player_property C ON A.playerId = C.playerId 
							WHERE
								A.deleteTime = 0 
								AND A.developLevel >= ?
								AND C.power > 0 
							ORDER BY
								A.developLevel DESC,
								C.power DESC,
								A.playerId ASC
								LIMIT ?`
)

const (
	zhuanshengRankSql = `SELECT
							B.serverId,
							A.playerId,
							A.zhuanSheng AS num,
							B.name,
							A.power 
						FROM
							t_player_property A 
							INNER JOIN t_player B ON B.id = A.playerId 
							AND B.serverId =?
						WHERE 
							A.zhuansheng >= ?
							AND A.power > 0  
						ORDER BY 
							A.zhuanSheng DESC,
							A.level DESC 
							LIMIT ?`
)
