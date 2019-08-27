package dao

const (
	forceRankActivitySql = `SELECT 
						A.serverId,
						B.playerId,
						A.name,
						C.allianceName AS gangName,
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
				   where 
				        B.power >0 
				   ORDER BY 
				        B.power DESC,
				  		B.level DESC, 
				   		B.playerId ASC 
				   LIMIT ?`
)

const (
	mountRankActivitySql = `SELECT 
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
							A.advancedId >= ? 
						ORDER  BY 
						      A.advancedId DESC, 
							  A.power DESC,
							  C.power DESC 
						LIMIT ?`
)

const (
	wingRankActivitySql = `SELECT 
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
						A.advancedId >= ?
					ORDER BY 
					    A.advancedId DESC, 
	  					A.power DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	bodyShieldRankActivitySql = `SELECT 
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
							A.advancedId >= ?
						ORDER  BY 
							A.advancedId DESC, 
  							A.power DESC,
  							C.power DESC 
						LIMIT ?`
)

const (
	weaponRankActivitySql = `SELECT 
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
	gangRankActivitySql = ` SELECT 
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
	shenFaRankActivitySql = `SELECT 
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
							A.advancedId >= ?
						ORDER BY 
							A.advancedId DESC, 
  							A.power DESC,
  							C.power DESC 
						LIMIT ?`
)

const (
	lingYuRankActivitySql = `SELECT 
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
						AND A.advancedId >= ?
						ORDER BY 
							A.advancedId DESC, 
  							A.power DESC,
  							C.power DESC 
						LIMIT ?`
)

const (
	featherRankActivitySql = `SELECT 
						B.serverId,
						A.playerId,
						B.name,
						A.featherId AS advancedId,
						A.fpower As power
					FROM 
						t_player_wing A
					INNER JOIN
						t_player B ON A.playerId = B.id AND B.serverId=?
					INNER JOIN
						t_player_property C ON A.playerId=C.playerId 
					WHERE 
						A.fpower >0 
					AND 
						A.featherId >=?
					ORDER BY 
						A.featherId DESC, 
	  					A.fpower DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	shieldRankActivitySql = `SELECT 
							B.serverId,
							A.playerId,
							B.name,
							A.shieldId AS advancedId,
							A.spower As power
						FROM 
							t_player_body_shield A
						INNER JOIN
							t_player B ON A.playerId = B.id AND B.serverId=?
						INNER JOIN
							t_player_property C ON A.playerId=C.playerId 
						WHERE
							 A.spower >0 
						AND 
							A.shieldId >= ?
						ORDER  BY 
							A.shieldId DESC, 
  							A.spower DESC,
  							C.power DESC 
						LIMIT ?`
)
const (
	chargeRankActivitySql = `SELECT
								B.serverId,
								A.playerId,
								A.goldNum AS num,
								B.name,
								C.power 
							FROM
								t_player_open_activity_charge A
								INNER JOIN t_player B ON A.playerId = B.id 
								AND B.serverId = ?
								INNER JOIN t_player_property C ON A.playerId = C.playerId 
							WHERE
								A.deleteTime = 0 
								AND A.groupId = ? 
								AND A.goldNum >= ? 
								AND A.endTime = ?
								AND C.power > 0 
							ORDER BY
								A.goldNum DESC,
								C.power DESC,
								A.playerId ASC 
								LIMIT ?`
)

const (
	lingTongForceRankActivitySql = `SELECT
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
									AND A.basePower >= ? 
									AND C.power > 0 
								ORDER BY
									A.basePower DESC,
									C.power DESC,
									A.playerId ASC 
									LIMIT ?`
)

const (
	goldEquipForceRankActivitySql = `SELECT
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
									AND A.power >= ? 
									AND C.power > 0 
								ORDER BY
									A.power DESC,
									C.power DESC,
									A.playerId ASC 
									LIMIT ?`
)

const (
	dianXingForceRankActivitySql = `SELECT
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
									AND A.power >= ? 
									AND C.power > 0 
								ORDER BY
									A.power DESC,
									C.power DESC,
									A.playerId ASC 
									LIMIT ?`
)

const (
	shenQiForceRankActivitySql = `SELECT
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
									AND A.power >= ? 
									AND C.power > 0 
								ORDER BY
									A.power DESC,
									C.power DESC,
									A.playerId ASC 
									LIMIT ?`
)

const (
	mingGeForceRankActivitySql = `SELECT
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
									AND A.power >= ? 
									AND C.power > 0 
								ORDER BY
									A.power DESC,
									C.power DESC,
									A.playerId ASC 
									LIMIT ?`
)

const (
	shengHenForceRankActivitySql = `SELECT
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
									AND A.power >= ? 
									AND C.power > 0 
								ORDER BY
									A.power DESC,
									C.power DESC,
									A.playerId ASC 
									LIMIT ?`
)

const (
	zhenFaForceRankActivitySql = `SELECT
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
									AND A.power >= ? 
									AND C.power > 0 
								ORDER BY
									A.power DESC,
									C.power DESC,
									A.playerId ASC 
									LIMIT ?`
)

const (
	tuLongEquipForceRankActivitySql = `SELECT
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
									AND A.power >= ? 
									AND C.power > 0 
								ORDER BY
									A.power DESC,
									C.power DESC,
									A.playerId ASC 
									LIMIT ?`
)

const (
	babyForceRankActivitySql = `SELECT
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
									AND A.power >= ? 
									AND C.power > 0 
								ORDER BY
									A.power DESC,
									C.power DESC,
									A.playerId ASC 
									LIMIT ?`
)

const (
	costRankActivitySql = `SELECT
							B.serverId,
							A.playerId,
							A.goldNum AS num,
							B.name,
							C.power 
						FROM
							t_player_open_activity_cost A
							INNER JOIN t_player B ON A.playerId = B.id 
							AND B.serverId =?
							INNER JOIN t_player_property C ON A.playerId = C.playerId 
						WHERE
							A.deleteTime = 0 
							AND A.groupId = ? 
							AND A.goldNum >= ? 
							AND A.endTime = ?
							AND C.power > 0 
						ORDER BY
							A.goldNum DESC,
							C.power DESC,
							A.playerId ASC
							LIMIT ?`
)

const (
	anQiRankActivitySql = `SELECT 
							B.serverId,
							A.playerId,
							B.name,
							A.advancedId,
							A.power 
						FROM 
							t_player_anqi A
						INNER JOIN
							t_player B ON A.playerId = B.id AND B.serverId= ?
						INNER JOIN
							t_player_property C ON A.playerId = C.playerId 
						WHERE
							 A.power >0 
						AND 
							A.advancedId >= ?
						ORDER  BY 
							A.advancedId DESC, 
  							A.power DESC,
  							C.power DESC 
						LIMIT ?`
)

const (
	charmRankActivitySql = `SELECT
							B.serverId,
							A.playerId,
							A.addNum AS num,
							B.name,
							C.power 
						FROM
							t_player_activity_add_num A
							INNER JOIN t_player B ON A.playerId = B.id 
							AND B.serverId = ?
							INNER JOIN t_player_property C ON A.playerId = C.playerId 
							INNER JOIN t_player_marry D ON A.playerId = D.playerId
						WHERE
							A.deleteTime = 0 
							AND A.groupId = ? 
							AND D.status = ?
							AND A.addNum >= ? 
							AND A.endTime = ?
							AND C.power > 0 
							AND D.spouseId != 0
						ORDER BY
							A.addNum DESC,
							C.power DESC,
							A.playerId ASC 
							LIMIT ?`
)

const (
	countRankActivitySql = `SELECT
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
								AND A.endTime = ?
							ORDER BY 
								A.times DESC,
								C.power DESC,
								A.playerId ASC
								LIMIT ?`
)

const (
	faBaoRankActivitySql = `SELECT 
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
						A.advancedId >= ?
					ORDER BY 
					    A.advancedId DESC, 
	  					A.power DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	xianTiRankActivitySql = `SELECT 
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
						A.advancedId >= ?
					ORDER BY 
					    A.advancedId DESC, 
	  					A.power DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	levelRankActivitySql = `SELECT
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
	shiHunFanRankActivitySql = `SELECT 
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
						A.advancedId >= ?
					ORDER BY 
					    A.advancedId DESC, 
	  					A.power DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	tianMoTiRankActivitySql = `SELECT 
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
						A.advancedId >= ?
					ORDER BY 
					    A.advancedId DESC, 
	  					A.power DESC,
	  					C.power DESC 
					LIMIT ?`
)

const (
	feiShengRankActivitySql = `SELECT
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
	marryDevelopRankActivitySql = `SELECT
										B.serverId,
										A.playerId,
										A.addNum AS num,
										B.name,
										C.power 
									FROM
										t_player_activity_add_num A
										INNER JOIN t_player B ON A.playerId = B.id 
										AND B.serverId = ?
										INNER JOIN t_player_property C ON A.playerId = C.playerId 
									WHERE
										A.deleteTime = 0 
										AND A.groupId = ? 
										AND A.addNum >= ? 
										AND A.endTime = ?
										AND C.power > 0 
									ORDER BY
										A.addNum DESC,
										C.power DESC,
										A.playerId ASC 
										LIMIT ?`
)

const (
	zhuanshengRankActivitySql = `SELECT
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
