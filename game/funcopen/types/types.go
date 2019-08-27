package types

type FuncOpenType int32

const (
	FuncOpenTypeMain       FuncOpenType = 1000 //主界面
	FuncOpenTypeFunc                    = 1001 //功能面板
	FuncOpenTypeQuest                   = 1002 //任务对话面板
	FuncOpenTypeInventory               = 1003 //背包
	FuncOpenTypeAlliance                = 1004 //仙盟
	FuncOpenTypeRole                    = 1005 //角色
	FuncOpenTypeSheZhi                  = 1006 //设置
	FuncOpenTypeShangCheng              = 1007 //商城
	FuncOpenTypesFriend                 = 1008 //好友
	FuncOpenTypeTeam                    = 1009 //组队
	FuncOpenTypeBloodPool               = 1010 //血池
	FuncOpenTypeZhuanSheng              = 1011 //转生
	//FuncOpenTypeFashionStar                           = 10102 //时装升星
	FuncOpenTypeSkill              = 10010 //技能（职业技能标签）
	FuncOpenTypeJueXue             = 10011 //绝学
	FuncOpenTypeXinFa              = 10012 //心法
	FuncOpenTypeTelent             = 10013 //天赋
	FuncOpenTypeMount              = 10020 //坐骑（主标签与升阶标签）
	FuncOpenTypeMountUpgrade       = 10021 //坐骑升级（吃幻化丹）
	FuncOpenTypeMountCul           = 10022 //坐骑培养（吃培养丹）
	FuncOpenTypeMountEquipment     = 10023 //坐骑装备
	FuncOpenTypeMountSkill         = 10024 //坐骑技能
	FuncOpenTypeHuaLingMount       = 10025 //坐骑化灵
	FuncOpenTypeMountUnreal        = 10026 //坐骑幻化
	FuncOpenTypeMountJinJie        = 10027 //坐骑进阶标签页
	FuncOpenTypeMountHuanHua       = 10028 //坐骑幻化标签页
	FuncOpenTypeMountEquipQiangHua = 10029 //坐骑装备强化
	FuncOpenTypeEquipmentStrength  = 10030 //锻造（装备强化标签）
	FuncOpenTypeEquimentStar       = 10031 //装备升星

	FuncOpenTypeEquimentXiLian              = 10032 //装备洗炼
	FuncOpenTypeEquimentKaiGuang            = 10033 //装备开光
	FuncOpenTypeBaGuaXiangQian              = 10034 //八卦镶嵌
	FuncOpenTypeFuShi                       = 10035 //八卦符石
	FuncOpenTypeSoul                        = 10040 //帝魂（主标签，包含激活与佩戴）
	FuncOpenTypeSoulUpgrade                 = 10041 //帝魂升级
	FuncOpenTypeSoulAwaken                  = 10042 //帝魂顿悟
	FuncOpenTypeSoulStrength                = 10043 //帝魂强化
	FuncOpenTypeBodyShield                  = 10050 //护体盾（主标签与升阶标签）
	FuncOpenTypeBodyShieldCul               = 10051 //护盾培养（标签页）
	FuncOpenTypeShield                      = 10052 //神盾尖刺
	FuncOpenTypeDanEat                      = 10060 //仙丹服用
	FuncOpenTypeDanAlchemy                  = 10061 //仙丹炼成
	FuncOpenTypeWeapon                      = 10070 //兵魂（激活与佩戴标签）
	FuncOpenTypeWeaponStar                  = 10071 //兵魂升星
	FuncOpenTypeWeaponAwaken                = 10072 //兵魂觉醒
	FuncOpenTypeGemDigged                   = 10080 //宝石（挖矿）
	FuncOpenTypeGemEmbedded                 = 10081 //宝石（镶嵌）
	FuncOpenTypeGemTaoZhuang                = 10082 //宝石（套装）
	FuncOpenTypeWing                        = 10090 //战翼（主标签与进阶）
	FuncOpenTypeFeather                     = 10092 //护体仙羽
	FuncOpenTypeWingUpgrade                 = 10091 //战翼（吃幻化丹）
	FuncOpenTypeWingRune                    = 10093 //战翼符文
	FuncOpenTypeWingSkill                   = 10094 //战翼技能
	FuncOpenTypeHuaLingWing                 = 10095 //战翼化灵
	FuncOpenTypeWingShengJi                 = 10096 //战翼升级
	FuncOpenTypeWingRuneUpgrade             = 10097 //战翼符文强化
	FuncOpenTypeWingAdvencedLabel           = 10098 //战翼进阶标签页
	FuncOpenTypeWingUnrealLabel             = 10099 //战翼幻化标签页
	FuncOpenTypeTitle                       = 10100 //外观（称号标签）
	FuncOpenTypeFashion                     = 10101 //外观（时装标签）
	FuncOpenTypeDingZhiTitle                = 10102 //定制称号
	FuncOpenTypeTitleUpstar                 = 10103 //称号升星
	FuncOpenTypeJingJie                     = 10110 //境界（主标签）
	FuncOpenTypeSynthesis                   = 10120 //合成(宝石)
	FuncOpenTypeShenfa                      = 10130 //身法
	FuncOpenTypeShenfaAdvanced              = 10131 //身法进阶
	FuncOpenTypeShenfaShengJi               = 10132 //身法升级
	FuncOpenTypeHuaLingShenFa               = 10133 //身法化灵
	FuncOpenTypeShenFaUnreal                = 10134 //身法幻化
	FuncOpenTypeShenFaEquip                 = 10135 //身法装备
	FuncOpenTypeShenFaSkill                 = 10136 //身法技能
	FuncOpenTypeShenFaEquipUpgrade          = 10137 //身法装备强化
	FuncOpenTypeShenFaUnrealLabel           = 10138 //身法幻化标签页
	FuncOpenTypeShenFaExpert                = 10139 //身法达人
	FuncOpenTypeLingYu                      = 10140 //领域
	FuncOpenTypeLingYuAdvanced              = 10141 //领域进阶
	FuncOpenTypeLingYuShengJi               = 10142 //领域升级
	FuncOpenTypeHuaLingLingYu               = 10143 //领域化灵
	FuncOpenTypeLingYuUnreal                = 10144 //领域幻化
	FuncOpenTypeLingYuEquip                 = 10145 //领域装备
	FuncOpenTypeLingYuSkill                 = 10146 //领域技能
	FuncOpenTypeLingYuEquipUpgrade          = 10147 //领域装备强化
	FuncOpenTypeLingYuUnrealLabel           = 10148 //领域幻化标签页
	FuncOpenTypeLingYuExpert                = 10149 //领域达人
	FuncOpenTypeYuanGodGold                 = 10150 //元神金装
	FuncOpenTypeYuanGodStrengthen           = 10151 //元神金装·升星
	FuncOpenTypeYuanGodRecasting            = 10152 //元神金装·重铸
	FuncOpenTypeYuanGodEat                  = 10153 //元神金装·吞噬
	FuncOpenTypeYuanGodUpgrade              = 10154 //元神金装.强化
	FuncOpenTypeYuanGodKaiGuang             = 10155 //元神金装.开光
	FuncOpenTypeYuanGodQiangHuaTaoZhuang    = 10156 //强化（套装）
	FuncOpenTypeYuanGodOpenLightTaoZhuang   = 10157 //开光（套装）
	FuncOpenTypeYuanGodUpstarTaoZhuang      = 10158 //升星（套装）
	FuncOpenTypeGoldEquipExtend             = 10159 //装备·继承
	FuncOpenTypeAnQi                        = 10160 //暗器
	FuncOpenTypeAnQiAdvanced                = 10161 //暗器升阶
	FuncOpenTypeAnQiJiGuan                  = 10162 //暗器机关
	FuncOpenTypeAnQiSkill                   = 10163 //暗器技能
	FuncOpenTypeHuaLingAnQi                 = 10164 //暗器化灵
	FuncOpenTypeAnQiPeiYang                 = 10165 //暗器培养
	FuncOpenTypeAnQiShengJi                 = 10166 //暗器升级
	FuncOpenTypeAnQiEquipUpgrade            = 10167 //暗器机关强化
	FuncOpenTypeAnQiExpert                  = 10168 //暗器达人
	FuncOpenTypeAnQiAdvencedRew             = 10169 //暗器升阶奖励
	FuncOpenTypeTianShu                     = 10170 //天书
	FuncOpenTypeFaBao                       = 10180 //法宝
	FuncOpenTypeFaBaoAdvanced               = 10181 //法宝进阶
	FuncOpenTypeFaBaoUnreal                 = 10182 //法宝幻化
	FuncOpenTypeFaBaoTongLing               = 10183 //法宝通灵
	FuncOpenTypeFaBaoSuit                   = 10184 //法宝配饰
	FuncOpenTypeFaBaoSkill                  = 10185 //法宝技能
	FuncOpenTypeFaBaoShengJi                = 10186 //法宝升级
	FuncOpenTypeHuaLingFaBao                = 10187 //法宝化灵
	FuncOpenTypeFaBaoEquipUpgrade           = 10188 //法宝配饰强化
	FuncOpenTypeFaBaoUnrealLabel            = 10189 //法宝幻化标签页
	FuncOpenTypeXianTi                      = 10190 //仙体
	FuncOpenTypeXianTiAdvanced              = 10191 //仙体进阶
	FuncOpenTypeXianTiUnreal                = 10192 //仙体幻化
	FuncOpenTypeXianTiSkill                 = 10193 //仙体技能
	FuncOpenTypeXianTiLingYu                = 10194 //仙体灵玉
	FuncOpenTypeXianTiShengJi               = 10195 //仙体升级
	FuncOpenTypeHuaLingXianTi               = 10196 //仙体化灵
	FuncOpenTypeXianTiEquipUpgrade          = 10197 //仙体灵玉强化
	FuncOpenTypeXianTiUnrealLabel           = 10198 //仙体幻化丹
	FuncOpenTypeXianTiExpert                = 10199 //仙体达人
	FuncOpenTypeLiveness                    = 10200 //活跃度系统
	FuncOpenTypeXueDu                       = 10210 //血盾
	FuncOpenTypeFuBen                       = 10220 //副本玩法
	FuncOpenTypeAdvanced                    = 10230 //升阶
	FuncOpenTypeMingRiKaiQi                 = 10240 //明日开启
	FuncOpenTypeExpBox                      = 10250 //经验魔瓶
	FuncOpenTypeDianXing                    = 10260 //点星
	FuncOpenTypeDianXingJieFeng             = 10261 //点星解封
	FuncOpenTypeWardrobe                    = 10270 //衣橱
	FuncOpenTypeShiHunFan                   = 10280 //噬魂幡
	FuncOpenTypeShiHunFanAdvanced           = 10281 //噬魂幡进阶
	FuncOpenTypeShiHunFanUnreal             = 10282 //噬魂幡幻化
	FuncOpenTypeShiHunFanShengJi            = 10283 //噬魂幡升级
	FuncOpenTypeShiHunFanEquip              = 10284 //噬魂幡装备
	FuncOpenTypeShiHunFanSkill              = 10285 //噬魂幡技能
	FuncOpenTypeHuaLingShiHunFan            = 10286 //噬魂幡化灵
	FuncOpenTypeShiHunFanEquipUpgrade       = 10287 //噬魂幡装备强化
	FuncOpenTypeShiHunFanExpert             = 10288 //噬魂幡达人
	FuncOpenTypeShiHunFanAdvencedRew        = 10289 //噬魂幡升阶奖励
	FuncOpenTypeTianMo                      = 10290 //天魔体
	FuncOpenTypeTianMoAdvanced              = 10291 //天魔体进阶
	FuncOpenTypeTianMoUnreal                = 10292 //天魔体幻化
	FuncOpenTypeTianMoUplevel               = 10293 //天魔体升级
	FuncOpenTypeTianMoEquipment             = 10294 //天魔体装备
	FuncOpenTypeTianMoSkill                 = 10295 //天魔体技能
	FuncOpenTypeHuaLingTianMoTi             = 10296 //天魔体化灵
	FuncOpenTypeTianMoEquipUpgrade          = 10297 //天魔体装备强化
	FuncOpenTypeTianMoExpert                = 10298 //天魔体达人
	FuncOpenTypeTianMoAdvencedRew           = 10299 //天魔体升阶奖励
	FuncOpenTypeLingTong                    = 10300 //灵童
	FuncOpenTypeLingTongUpgrade             = 10301 //灵童升级
	FuncOpenTypeLingTongPeiYang             = 10302 //灵童培养
	FuncOpenTypeLingTongHuanHua             = 10303 //灵童幻化
	FuncOpenTypeLingTongUpstar              = 10304 //灵童升星
	FuncOpenTypeLingTongEquip               = 10305 //灵童装备
	FuncOpenTypeLingTongEquipStrengthen     = 10306 //灵童装备强化
	FuncOpenTypeLingTongEquipAdvanced       = 10307 //灵童装备升阶
	FuncOpenTypeShenZhuLingTongEquip        = 10308 //灵童装备灵锻
	FuncOpenTypeLingTongEquipWuXingLingZhu  = 10309 //灵童装备五行灵珠
	FuncOpenTypeLingTongWeapon              = 10310 //灵兵
	FuncOpenTypeLingTongWeaponAdvanced      = 10311 //灵兵进阶
	FuncOpenTypeLingTongWeaponUnreal        = 10312 //灵兵幻化
	FuncOpenTypeLingTongWeaponUpgrade       = 10313 //灵兵升级
	FuncOpenTypeLingTongWeaponEquip         = 10314 //灵兵装备
	FuncOpenTypeLingTongWeaponSkill         = 10315 //灵兵技能
	FuncOpenTypeHuaLingLingBing             = 10316 //灵兵化灵
	FuncOpenTypeLingTongWeaponEquipUpgrade  = 10317 //灵兵装备强化
	FuncOpenTypeLingTongWeaponUnrealLabel   = 10318 //灵兵幻化
	FuncOpenTypeLingTongWeaponExpert        = 10319 //灵兵达人
	FuncOpenTypeLingTongMount               = 10320 //灵骑
	FuncOpenTypeLingTongMountAdvanced       = 10321 //灵骑进阶
	FuncOpenTypeLingTongMountUnreal         = 10322 //灵骑幻化
	FuncOpenTypeLingTongMountUpgrade        = 10323 //灵骑升级
	FuncOpenTypeLingTongMountEquip          = 10324 //灵骑装备
	FuncOpenTypeLingTongMountSkill          = 10325 //灵骑技能
	FuncOpenTypeLingTongMountPeiYang        = 10326 //灵骑技能
	FuncOpenTypeHuaLingLingQi               = 10327 //灵骑化灵
	FuncOpenTypeLingTongMountEquipUpgrade   = 10328 //灵骑装备强化
	FuncOpenTypeLingTongMountExpert         = 10329 //灵骑达人
	FuncOpenTypeLingTongWing                = 10330 //灵翼
	FuncOpenTypeLingTongWingAdvanced        = 10331 //灵翼进阶
	FuncOpenTypeLingTongWingUnreal          = 10332 //灵翼幻化
	FuncOpenTypeLingTongWingUpgrade         = 10333 //灵翼升级
	FuncOpenTypeLingTongWingEquip           = 10334 //灵翼装备
	FuncOpenTypeLingTongWingSkill           = 10335 //灵翼技能
	FuncOpenTypeHuaLingLingWing             = 10336 //灵翼化灵
	FuncOpenTypeLingTongWingEquipUpgrade    = 10337 //灵翼装备强化
	FuncOpenTypeLingTongWingUnrealLabel     = 10338 //灵翼幻化
	FuncOpenTypeLingTongWingExpert          = 10339 //灵翼达人
	FuncOpenTypeLingTongShenFa              = 10340 //灵身
	FuncOpenTypeLingTongShenFaAdvanced      = 10341 //灵身进阶
	FuncOpenTypeLingTongShenFaUnreal        = 10342 //灵身幻化
	FuncOpenTypeLingTongShenFaUpgrade       = 10343 //灵身升级
	FuncOpenTypeLingTongShenFaEquip         = 10344 //灵身装备
	FuncOpenTypeLingTongShenFaSkill         = 10345 //灵身技能
	FuncOpenTypeHuaLingLingShen             = 10346 //灵身化灵
	FuncOpenTypeLingTongShenFaEquipUpgrade  = 10347 //灵身装备强化
	FuncOpenTypeLingTongShenFaUnrealLabel   = 10348 //灵身幻化
	FuncOpenTypeLingTongShenFaExpert        = 10349 //灵身达人
	FuncOpenTypeLingTongLingYu              = 10350 //灵域
	FuncOpenTypeLingTongLingYuAdvanced      = 10351 //灵身进阶
	FuncOpenTypeLingTongLingYuUnreal        = 10352 //灵身幻化
	FuncOpenTypeLingTongLingYuUpgrade       = 10353 //灵身升级
	FuncOpenTypeLingTongLingYuEquip         = 10354 //灵身装备
	FuncOpenTypeLingTongLingYuSkill         = 10355 //灵身技能
	FuncOpenTypeHuaLingLingArea             = 10356 //灵域化灵
	FuncOpenTypeLingTongLingYuEquipUpgrade  = 10357 //灵域装备强化
	FuncOpenTypeLingTongLingYuUnrealLabel   = 10358 //灵域幻化
	FuncOpenTypeLingTongLingYuExpert        = 10359 //灵域达人
	FuncOpenTypeLingTongFaBao               = 10360 //灵宝
	FuncOpenTypeLingTongFaBaoAdvanced       = 10361 //灵宝进阶
	FuncOpenTypeLingTongFaBaoUnreal         = 10362 //灵宝幻化
	FuncOpenTypeLingTongFaBaoUpgrade        = 10363 //灵宝升级
	FuncOpenTypeLingTongFaBaoEquip          = 10364 //灵宝装备
	FuncOpenTypeLingTongFaBaoSkill          = 10365 //灵宝技能
	FuncOpenTypeLingTngFaBaoTongLing        = 10366 //灵宝通灵
	FuncOpenTypeHuaLingLingBao              = 10367 //灵宝化灵
	FuncOpenTypeLingTongFaBaoEquipUpgrade   = 10368 //灵宝装备强化
	FuncOpenTypeLingTongFaBaoTongLingLabel  = 10369 //灵宝通灵
	FuncOpenTypeLingTongXianTi              = 10370 //灵体
	FuncOpenTypeLingTongXianTiAdvanced      = 10371 //灵宝进阶
	FuncOpenTypeLingTongXianTiUnreal        = 10372 //灵宝幻化
	FuncOpenTypeLingTongXianTiUpgrade       = 10373 //灵宝升级
	FuncOpenTypeLingTongXianTiEquip         = 10374 //灵宝装备
	FuncOpenTypeLingTongXianTiSkill         = 10375 //灵宝技能
	FuncOpenTypeHuaLingLingTi               = 10376 //灵体化灵
	FuncOpenTypeLingTongXianTiEquipUpgrade  = 10377 //灵体装备强化
	FuncOpenTypeLingTongXianTiUnrealLabel   = 10378 //灵体幻化
	FuncOpenTypeLingTongXianTiExpert        = 10379 //灵体达人
	FuncOpenTypeFeiSheng                    = 10400 //飞升
	FuncOpenTypeKaiFuMuBiao                 = 10410 // 开发目标
	FuncOpenTypeHongBao                     = 10420 //红包
	FuncOpenTypeMingGe                      = 10500 // 命格
	FuncOpenTypeMingGong                    = 10501 // 命宫
	FuncOpenTypeMingPan                     = 10502 // 命盘
	FuncOpenTypeMingGeHeCheng               = 10503 // 命格合成
	FuncOpenTypeMingGeXunBao                = 10504 // 命格寻宝
	FuncOpenTypeShenQi                      = 10510 // 神器
	FuncOpenTypeShenQiSmelt                 = 10511 // 神器淬炼
	FuncOpenTypeShenQiQiLing                = 10512 // 神器器灵
	FuncOpenTypeShenQiZhuLing               = 10513 // 神器注灵
	FuncOpenTypeShenQiTaoZhuang             = 10514 // 神器套装
	FuncOpenTypeShenQiResolve               = 10515 // 神器分解
	FuncOpenTypeShenQiXunBao                = 10516 // 神器寻宝
	FuncOpenTypeShengHen                    = 10520 //圣痕
	FuncOpenTypeShengHenXunBao              = 10521 //圣痕寻宝
	FuncOpenTypeTuLongEquip                 = 10530 //屠龙装备
	FuncOpenTypeTuLongQiangHua              = 10531 //屠龙装备-强化
	FuncOpenTypeTuLongRongHe                = 10532 //屠龙装备-融合
	FuncOpenTypeTuLongZhuanHua              = 10533 //屠龙装备-转化
	FuncOpenTypeTuLongSkill                 = 10534 //屠龙装备-技能
	FuncOpenTypeTuLongXunBao                = 10535 //屠龙装备-寻宝
	FuncOpenTypeYingLingPu                  = 10540 // 英灵谱
	FuncOpenTypeZhiZhunYingLingPu           = 10541 // 至尊英灵谱
	FuncOpenTypeZhenFa                      = 10550 //阵法
	FuncOpenTypeZhenQi                      = 10551 //阵旗
	FuncOpenTypeZhenHuo                     = 10552 //阵法仙火
	FuncOpenTypeZhenFaSuit                  = 10553 //阵法套装
	FuncOpenTypeDongFang                    = 10560 //洞房
	FuncOpenTypeBaoBao                      = 10561 //宝宝
	FuncOpenTypeBaoBaoLearn                 = 10562 //四书五经
	FuncOpenTypeBaoBaoToy                   = 10563 //玩具
	FuncOpenTypeBaoBaoZhuanShi              = 10564 //投胎转世
	FuncOpenTypeFuQiFuBen                   = 10565 //夫妻副本
	FuncOpenTypeTrad                        = 10570 //交易市场
	FuncOpenTypeQiXueQiang                  = 10580 //泣血枪
	FuncOpenTypeChuangShiPreview            = 10590 //创世之战预告
	FuncOpenTypeHuanLeJingCai               = 10601 // 欢乐竞猜
	FuncOpenTypeBiWuDaHui                   = 10600 // 比武大会
	FuncOpenTypeJieYi                       = 10610 // 结义
	FuncOpenTypeXiongDiWeiMing              = 10611 // 兄弟威名
	FuncOpenTypeWeiMingUpLev                = 10612 // 威名升级
	FuncOpenTypeXiongDiToken                = 10613 // 兄弟信物
	FuncOpenTypeTokenUpLev                  = 10614 // 信物强化
	FuncOpenTypeWushuangWeapon              = 10620 //无双神器
	FuncOpenTypeGodCasting                  = 10630 //神铸
	FuncOpenTypeForgeSoul                   = 10631 //锻魂
	FuncOpenTypeCastingSpirit               = 10632 //铸灵
	FuncOpenTypeGodCastingInherit           = 10633 //神铸继承
	FuncOpenTypeZhenXiBoss                  = 10634 //神铸BOSS
	FuncOpenTypeDingShiBoss                 = 10635 //定时BOSS
	FuncOpenTypeShengShouBoss               = 10636 //圣兽秘境
	FuncOpenTypeChuangShiZhiZhan            = 10640 //创世之战
	FuncOpenTypeMountDaRen                  = 11020 //坐骑达人
	FuncOpenTypeMountJinJieRew              = 11021 //坐骑升阶奖励
	FuncOpenTypeShenZhuMount                = 11022 //坐骑装备神铸
	FuncOpenTypeTongLingMount               = 11023 //坐骑通灵
	FuncOpenTypeJueXingMount                = 11024 //坐骑觉醒
	FuncOpenTypeWingExpert                  = 11090 //战翼达人
	FuncOpenTypeWingAdvencedRew             = 11091 //战翼升阶奖励
	FuncOpenTypeShenZhuWing                 = 11092 //战翼装备神铸
	FuncOpenTypeTongLingWing                = 11093 //战翼通灵
	FuncOpenTypeJueXingWing                 = 11094 //战翼觉醒
	FuncOpenTypeShenFaAdvencedRew           = 11130 //身法升阶奖励
	FuncOpenTypeShenZhuShenFa               = 11131 //身法装备神铸
	FuncOpenTypeTongLingShenFa              = 11132 //身法通灵
	FuncOpenTypeJueXingShenFa               = 11133 //身法觉醒
	FuncOpenTypeLingYuAdvencedRew           = 11140 //领域升阶奖励
	FuncOpenTypeShenZhuLingYu               = 11141 //领域装备神铸
	FuncOpenTypeTongLingLingYu              = 11142 //领域通灵
	FuncOpenTypeJueXingLingYu               = 11143 //领域觉醒
	FuncOpenTypeShenZhuAnQi                 = 11160 //暗器装备神铸
	FuncOpenTypeTongLingAnQi                = 11161 //暗器通灵
	FuncOpenTypeJueXingAnQi                 = 11162 //暗器觉醒
	FuncOpenTypeFaBaoExpert                 = 11180 //法宝达人
	FuncOpenTypeFaBaoAdvencedRew            = 11181 //法宝升阶奖励
	FuncOpenTypeShenZhuFaBao                = 11182 //法宝装备神铸
	FuncOpenTypeTongLingFaBao               = 11183 //法宝通灵
	FuncOpenTypeJueXingFaBao                = 11184 //法宝觉醒
	FuncOpenTypeXianTiAdvencedRew           = 11190 //仙体升阶奖励
	FuncOpenTypeShenZhuXianTi               = 11191 //仙体装备神铸
	FuncOpenTypeTongLingXianTi              = 11192 //仙体通灵
	FuncOpenTypeJueXingXianTi               = 11193 //仙体觉醒
	FuncOpenTypeShenZhuShiHunFan            = 11280 //噬魂幡装备神铸
	FuncOpenTypeTongLingShiHunFan           = 11281 //噬魂幡通灵
	FuncOpenTypeJueXingShiHunFan            = 11282 //噬魂幡觉醒
	FuncOpenTypeShenZhuTianMoTi             = 11290 //天魔体装备神铸
	FuncOpenTypeTongLingTianMoTi            = 11291 //天魔体通灵
	FuncOpenTypeJueXingTianMoTi             = 11292 //天魔体觉醒
	FuncOpenTypeLingTongWeaponAdvencedRew   = 11310 //灵兵升阶奖励
	FuncOpenTypeShenZhuLingTongWeapon       = 11311 //灵兵装备神铸
	FuncOpenTypeTongLingLingTongWeapon      = 11312 //灵兵通灵
	FuncOpenTypeJueXingLingTongWeapon       = 11313 //灵兵觉醒
	FuncOpenTypeLingTongMountAdvencedRew    = 11320 //灵骑升阶奖励
	FuncOpenTypeShenZhuLingTongMount        = 11321 //灵骑装备神铸
	FuncOpenTypeTongLingLingTongMount       = 11322 //灵骑通灵
	FuncOpenTypeJueXingLingTongMount        = 11323 //灵骑觉醒
	FuncOpenTypeLingTongWingAdvencedRew     = 11330 //灵翼升阶奖励
	FuncOpenTypeShenZhuLingTongWing         = 11331 //灵翼装备神铸
	FuncOpenTypeTongLingLingTongWing        = 11332 //灵翼通灵
	FuncOpenTypeJueXingLingTongWing         = 11333 //灵翼觉醒
	FuncOpenTypeLingTongShenFaAdvencedRew   = 11340 //灵身升阶奖励
	FuncOpenTypeShenZhuLingTongShenFa       = 11341 //灵身装备神铸
	FuncOpenTypeTongLingLingTongShenFa      = 11342 //灵身通灵
	FuncOpenTypeJueXingLingTongShenFa       = 11343 //灵身觉醒
	FuncOpenTypeLingTongLingYuAdvencedRew   = 11350 //灵域升阶奖励
	FuncOpenTypeShenZhuLingTongLingYu       = 11351 //灵域装备神铸
	FuncOpenTypeTongLingLingTongLingYu      = 11352 //灵域通灵
	FuncOpenTypeJueXingLingTongLingYu       = 11353 //灵域觉醒
	FuncOpenTypeLingTongFaBaoExpert         = 11360 //灵宝达人
	FuncOpenTypeLingTongFaBaoAdvencedRew    = 11361 //灵宝升阶奖励
	FuncOpenTypeShenZhuLingTongFaBao        = 11362 //灵宝装备神铸
	FuncOpenTypeTongLingLingTongFaBao       = 11363 //灵宝通灵
	FuncOpenTypeJueXingLingTongFaBao        = 11364 //灵宝觉醒
	FuncOpenTypeLingTongXianTiAdvencedRew   = 11370 //灵体升阶奖励
	FuncOpenTypeShenZhuLingTongXianTi       = 11371 //灵体装备神铸
	FuncOpenTypeTongLingLingTongXianTi      = 11372 //灵体通灵
	FuncOpenTypeJueXingLingTongXianTi       = 11373 //灵体觉醒
	FuncOpenTypePrivateCustomer             = 11380 //专属客服
	FuncOpenTypeMyHouse                     = 11390 //我的房子
	FuncOpenTypeZhouKa                      = 11400 //周卡
	FuncOpenTypeRing                        = 11410 // 特戒
	FuncOpenTypeRingStrengthen              = 11411 // 特戒强化
	FuncOpenTypeRingAdvance                 = 11412 // 特戒进阶
	FuncOpenTypeRingFuse                    = 11413 // 特戒融合
	FuncOpenTypeRingJingLing                = 11414 // 特戒净灵
	FuncOpenTypeRingXunBao                  = 11415 // 特戒寻宝
	FuncOpenTypeShangGuZhiLing              = 11420 //上古之灵
	FuncOpenTypeShangGuZhiLingUpLevel       = 11421 //上古之灵升级
	FuncOpenTypeShangGuZhiLingLingWen       = 11422 //上古之灵灵纹
	FuncOpenTypeShangGuZhiLingUpRank        = 11423 //上古之灵升阶
	FuncOpenTypeShangGuZhiLingLingLian      = 11424 //上古之灵灵炼
	FuncOpenTypeShangGuZhiLingMiLing        = 11425 //上古之灵觅灵
	FuncOpenTypeActivity                    = 20100 //活动
	FuncOpenTypeYueXiaQingYuan              = 20101 //月下情缘
	FuncOpenTypeJiuXiaoChengZhan            = 20201 //九霄城战
	FuncOpenTypeQuiz                        = 20301 //仙尊答题
	FuncOpenTypeShaLuZhiDu                  = 20401 //杀戮之都
	FuncOpenTypeQiangLongYi                 = 30101 //抢龙椅
	FuncOpenTypeTianJieTa                   = 30201 //天劫塔
	FuncOpenTypeCaiLiaoFuBen                = 30240 //材料副本
	FuncOpenTypeCaiLiaoMount                = 30241 //材料副本.坐骑
	FuncOpenTypeCaiLiaoWing                 = 30242 //材料副本.战翼
	FuncOpenTypeCaiLiaoShenFa               = 30243 //材料副本.身法
	FuncOpenTypeCaiLiaoFaBao                = 30244 //材料副本.法宝
	FuncOpenTypeCaiLiaoXianTi               = 30245 //材料副本.仙体
	FuncOpenTypeCaiLiaoTianMo               = 30246 //材料副本.天魔
	FuncOpenTypeCaiLiaoLingTong             = 30247 //材料副本.灵童
	FuncOpenTypeCaiLiaoLingBing             = 30248 //材料副本.灵兵
	FuncOpenTypeCaiLiaoLingYu               = 30249 //材料副本.灵域
	FuncOpenTypeBossHuanJing                = 30250 //boss幻境
	FuncOpenTypeOutlandBoss                 = 30260 //外域boss
	FuncOpenTypeDailyQuest                  = 30270 //日环任务
	FuncOpenTypeBaGuaMiJing                 = 30280 //八卦秘境
	FuncOpenTypeTuMo                        = 30301 //屠魔任务
	FuncOpenTypeXinaFu                      = 30401 //秘境仙府
	FuncOpenTypeExpFuBen                    = 30402 //经验副本
	FuncOpenTypeSilverFuBen                 = 30403 //银两副本
	FuncOpenTypeCoupleFuBen                 = 30404 //双修副本
	FuncOpenTypeTeamCopy                    = 30405 //组队副本
	FuncOpenTypeTeamCopySilver              = 30406 //组队副本.银两
	FuncOpenTypeTeamCopyXingChen            = 30407 //组队副本.星尘
	FuncOpenTypeTeamCopyXueMo               = 30408 //组队副本.血魔
	FuncOpenTypeTeamCopyZhuangShengEquip    = 30409 //组队副本.转生装备
	FuncOpenTypeTeamCopyWeapon              = 30410 //组队副本.兵魂
	FuncOpenTypeTeamCopyLingTong            = 30411 //组队副本.灵童
	FuncOpenTypeTeamCopyStrength            = 30412 //组队副本.强化
	FuncOpenTypeTeamCopyUpstar              = 30413 //组队副本.升星
	FuncOpenTypeDiHunYiJi                   = 30501 //帝魂遗迹
	FuncOpenTypePaiHangBang                 = 30601 //排行榜
	FuncOpenTypeResFind                     = 30701 //资源找回
	FuncOpenTypeTinaJiPai                   = 30801 //天机牌
	FuncOpenTypeZhuoYaoJi                   = 30901 //捉妖记
	FuncOpenTypeFourGod                     = 31001 //四神遗迹
	FuncOpenTypeArena                       = 31101 //3v3竞技场
	FuncOpenTypePersonalTransportation      = 31201 //个人镖车
	FuncOpenTypeAllianceTransportation      = 31202 //仙盟镖车
	FuncOpenTypeDragon                      = 31301 //神龙现世
	FuncOpenTypeWorldBoss                   = 31401 //世界BOSS
	FuncOpenTypeCrossWorldBoss              = 31402 //魔界巢穴(跨服boss)
	FuncOpenTypeMarry                       = 31501 //结婚
	FuncOpenTypeMarryPre                    = 31502 //结婚预告
	FuncOpenTypeDingQing                    = 31503 //定期信物
	FuncOpenTypeMarryJiNian                 = 31504 //结婚纪念
	FuncOpenTypeMarryHunJie                 = 31505 //结婚婚戒
	FuncOpenTypeMarryAiQingShu              = 31506 //结婚爱情树
	FuncOpenTypeMarryBiaoBai                = 31507 //表白
	FuncOpenTypeLingChiFighting             = 31601 //灵池争夺
	FuncOpenTypeCangLongQiJu                = 31701 //苍龙棋局
	FuncOpenTypeCrossTuLong                 = 31801 //跨服屠龙
	FuncOpenTypeCrossLianYu                 = 31901 //无间炼狱
	FuncOpenTypeDaBaoTower                  = 32001 //打宝塔
	FuncOpenTypeMyBoss                      = 32101 //个人BOSS
	FuncOpenTypeBoss                        = 32201 //BOSS
	FuncOpenTypeGodSiege                    = 32301 //神兽攻城
	FuncOpenTypeDenseWat                    = 32401 //金银密窟
	FuncOpenTypeShenMoWar                   = 32501 //神魔战场
	FuncOpenTypeCangJingGe                  = 32601 //藏经阁
	FuncOpenTypeEquipBaoKu                  = 32701 //装备宝库
	FuncOpenTypeMaterialBaoKu               = 32702 //材料宝库
	FuncOpenTypeAllianceDaily               = 32801 //仙盟日常
	FuncOpenTypeAllianceBoss                = 32802 //仙盟Boss
	FuncOpenTypeAllianceAltar               = 32901 //仙盟圣坛
	FuncOpenTypeXianTaoDaHui                = 33001 //仙桃大会
	FuncOpenTypeShenYuZhiZhan               = 33101 //神域之战
	FuncOpenTypeLongGongTanBao              = 33201 //龙宫探宝
	FuncOpenTypeYuXiZhiZhan                 = 33301 //玉玺之战
	FuncOpenTypeFirstCharge                 = 40000 //首冲奖励
	FuncOpenTypeFirstChargeCycleDay         = 40001 //每日首充
	FuncOpenTypeVIPPluss                    = 40002 //至尊会员
	FuncOpenTypeXianZunCard                 = 40003 // 仙尊特权卡
	FuncOpenTypeWelfareHall                 = 40010 //福利大厅
	FuncOpenTypeWelfareLogin                = 40011 //登录奖励
	FuncOpenTypeWelfareUplevel              = 40012 //升级礼包
	FuncOpenTypeWelfareOnline               = 40013 //在线礼包
	FuncOpenTypeWelfareSign                 = 40014 //签到(暂略
	FuncOpenTypeWelfareGiftBag              = 40015 //激活码兑换
	FuncOpenTypeOpenActivity                = 40020 //开服活动
	FuncOpenTypeInvestLevel                 = 40021 //投资计划
	FuncOpenTypeInvestSevenDay              = 40022 //7日投资
	FuncOpenTypeFeedbackCharge              = 40023 //充值返利
	FuncOpenTypeFeedbackCost                = 40024 //消费返利
	FuncOpenTypeRankCharge                  = 40025 //充值排行
	FuncOpenTypeRankCost                    = 40026 //消费排行
	FuncOpenTypeContinueCharge              = 40027 //连续充值
	FuncOpenTypeXunHuanRankCost             = 40028 //消费排行
	FuncOpenTypeRankSevenDay                = 40030 //7日冲刺
	FuncOpenTypeRankMount                   = 40031 //坐骑排行
	FuncOpenTypeRankWing                    = 40032 //战翼排行
	FuncOpenTypeRankBodyShield              = 40033 //护盾排行
	FuncOpenTypeRankLingyu                  = 40034 //领域排行
	FuncOpenTypeRankShenfa                  = 40035 //身法排行
	FuncOpenTypeRankFeather                 = 40036 //护体仙羽排行
	FuncOpenTypeRankShield                  = 40037 //盾刺排行
	FuncOpenTypeRankAnQi                    = 40200 //暗器排行
	FuncOpenTypeRankFaBao                   = 40201 //法宝排行
	FuncOpenTypeRankXianTi                  = 40202 //仙体排行
	FuncOpenTypeMassacre                    = 40038 //戮仙刃
	FuncOpenTypeGoldEquiBag                 = 40040 //转生大礼包
	FuncOpenTypeLongFengChengXiangOpen      = 40050 //龙凤呈祥开服
	FuncOpenTypeLongFengChengXiangMerge     = 40051 //龙凤呈祥合服
	FuncOpenTypeLuckyTurnplate              = 40060 //幸运转盘
	FuncOpenTypeHuHuShengFeng               = 40070 //虎虎生风
	FuncOpenTypeShenQiRenZhu                = 40071 // 神器认主
	FuncOpneTypeShenQiRenZhuJieRi           = 40072 //神器认主节日
	FuncOpenTypeHeiHuoShop                  = 40080 //黑货商店
	FuncOpenTypeFeiHuangTengDa              = 40090 //飞黄腾达
	FuncOpenTypeSingleCharge                = 40100 //开服单笔充值
	FuncOpenTypeSingleChargeMaxRew          = 40101 //单笔充值
	FuncOpenTypeYiZheTeHui                  = 40110 //一折特惠
	FuncOpenTypeSystemJinJie                = 40120 //系统进阶活动
	FuncOpenTypeSystemJinJieMount           = 40121 //系统进阶活动(坐骑)
	FuncOpenTypeSystemJinJieWing            = 40122 //系统进阶活动(战翼)
	FuncOpenTypeSystemJinJieAnQi            = 40123 //系统进阶活动(暗器)
	FuncOpenTypeSystemJinJieBodyShiled      = 40124 //系统进阶活动(护体盾)
	FuncOpenTypeSystemJinJieLingYu          = 40125 //系统进阶活动(领域)
	FuncOpenTypeSystemJinJieShenFa          = 40126 //系统进阶活动(身法)
	FuncOpenTypeSystemJinJieFaBao           = 40128 //系统进阶活动(法宝)
	FuncOpenTypeQiTianLeiChong              = 40130 //七天累充
	FuncOpenTypeXinaShiXuYuan               = 40140 //限时许愿
	FuncOpenTypeYuanBaoSongBuTing           = 40150 //元宝送不停
	FuncOpenTypeNiChongWoSong               = 40160 //你充我送
	FuncOpenTypeNiChongWoSongTwo            = 40161 //你充我送
	FuncOpenTypeShengJieSongJiang           = 40170 //升阶送奖励
	FuncOpenTypeYaoHu                       = 40180 //妖狐来袭
	FuncOpenTypeGoldLaBa                    = 40190 //元宝拉霸
	FuncOpenTypeChristmasParty              = 40260 //圣诞嘉年华
	FuncOpenTypeChristmasXueRen             = 40261 //圣诞雪人
	FuncOpenTypeChristmasTree               = 40262 //圣诞树
	FuncOpenTypeChristmasGift               = 40263 //圣诞豪礼
	FuncOpenTypeChristmasShop               = 40264 //圣诞商店
	FuncOpenTypeChristmasGoldCar            = 40265 //驯鹿元宝车
	FuncOpenTypeNewYearParty                = 40280 //元旦嘉年华
	FuncOpenTypeNewYearJinZhuSongLi         = 40281 //金猪送礼
	FuncOpenTypeNewYearJinZhuNaoCheng       = 40282 //金猪闹城
	FuncOpenTypeNewYearDaPaoDaQiu           = 40283 //大炮打气球
	FuncOpenTypeNewYearKuangHuan            = 40284 //元旦狂欢
	FuncOpenTypeNewYearZhaoCaiJinBao        = 40285 //招财进宝
	FuncOpenTypeMadeResource                = 40290 //经验炼制
	FuncOpenTypeCycleSingleCharge           = 40300 //每日单笔
	FuncOpenTypeCycleSingleChargeTwo        = 40301 //每日单笔2
	FuncOpenTypeCycleSingleChargeThree      = 40302 //每日单笔3
	FuncOpenTypeBossFirstKill               = 40310 //BOSS首杀
	FuncOpenTypeKanJiaGift                  = 40320 //砍价大礼包
	FuncOpenTypeChaoZhiTaoCan               = 40330 //超值套餐
	FuncOpenTypeYangJinJi                   = 40340 //养鸡生金蛋
	FuncOpenTypeLingQiChongCi               = 40351 //灵骑冲刺
	FuncOpenTypeShiHunChongCi               = 40352 //噬魂冲刺
	FuncOpenTypeLingYi                      = 40353 //灵翼冲刺
	FuncOpenTypeChongCi                     = 40354 //天魔冲刺
	FuncOpenTypeLingBingChongCi             = 40355 //灵兵冲刺
	FuncOpenTypeLingBaoChongCi              = 40356 //灵宝冲刺
	FuncOpenTypeLingTiChongCi               = 40357 //灵体冲刺
	FuncOpenTypeLingYuChongCi               = 40358 //灵域冲刺
	FuncOpenTypeLingShenChongCi             = 40359 //灵身冲刺
	FuncOpenTypeLingQiJinJie                = 40361 //灵骑进阶
	FuncOpenTypeShiHunJinJie                = 40362 //噬魂进阶
	FuncOpenTypeLingYiJinJie                = 40363 //灵翼进阶
	FuncOpenTypeLingMoJinJie                = 40364 //天魔进阶
	FuncOpenTypeLingBingJinJie              = 40365 //灵兵进阶
	FuncOpenTypeLingBaoJinJie               = 40366 //灵宝进阶
	FuncOpenTypeLingTiJinJie                = 40367 //灵体进阶
	FuncOpenTypeLingYuJinJie                = 40368 //灵域进阶
	FuncOpenTypeLingShenJinJie              = 40369 //灵身进阶
	FuncOpenTypeCrazyBox                    = 40370 //疯狂宝箱
	FuncOpenTypeRankLevel                   = 40380 //等级排行
	FuncOpenTypeRankCharm                   = 40381 //魅力排行
	FuncOpenTypeRankBiaoBai                 = 40382 //表白排行
	FuncOpenTypeChongZhiSongFangZi          = 40390 //充值送房子
	FuncOpenTypeZaJinDan                    = 40400 //砸金蛋
	FuncOpenTypeMingRenPu                   = 40410 //名人谱
	FuncOpenTypeXunHuanMingRenPu            = 40411 //循环名人谱
	FuncOpenTypeShouChongFanGold            = 40420 //首充返元宝
	FuncOpenTypeShouChongFanBindGold        = 40430 //首充返绑元
	FuncOpenTypeYuanBaoFanLi                = 40421 //元宝返利
	FuncOpenTypeBangYuanFanLi               = 40431 //绑元返利
	FuncOpenTypeShouChongFanBei             = 40440 //首冲翻倍
	FuncOpenTypeNewFirstCharge              = 40441 //首充翻倍
	FuncOpenTypeJingCaiHuoDong              = 40660 //精彩活动
	FuncOpenTypeXunHuanCycleSingle          = 40661 //每日单笔
	FuncOpenTypeXunHuanShouChongFanYuanBao  = 40662 //首充返元宝
	FuncOpenTypeXunHuanShouChongFanBangYuan = 40663 //首充返绑元
	FuncOpenTypeXunHuanChongZhiPaiHang      = 40664 //充值排行
	FuncOpenTypeXunHuanJinJieFanHuan        = 40665 //升阶返还
	FuncOpenTypeXunHuanHuanLeChouJiang      = 40666 //欢乐抽签
	FuncOpenTypeXunHuanMoJin                = 40667 //摸金
	FuncOpenTypeXunHuanGuaGuaLe             = 40668 //刮刮乐
	FuncOpenTypeXunHuanZhiZhunLunPan        = 40669 //至尊轮盘
	FuncOpenTypeXunHuanQuanMinQiangGou      = 40670 //全民抢购
	FuncOpenTypeXunHuanJuHuaSuan            = 40671 //聚划算
	FuncOpenTypeXunHuanXianShiLiBao         = 40672 //限时礼包
	FuncOpenTypeXunHuanMaiYiSongYi          = 40673 //买一送一
	FuncOpenTypeXunHuanZhaoCaiJinBao        = 40674 //招财进宝
	FuncOpenTypeXunHuanHuangJinKuangGong    = 40675 //黄金矿工
	FuncOpenTypeXunHuanShenMiBaoXiang       = 40676 //神秘宝箱
	FuncOpenTypeXunHuanYiXiaoBoDa           = 40677 //以小博大
	FuncOpenTypeXunHuanXunHuanChongZhi      = 40700 //循环充值
	FuncOpenTypeXunHuanJiFenShangCheng      = 40701 //积分商城
	FuncOpenTypeXunHuanShengJieHaoLi        = 40702 //升阶豪礼
	FuncOpenTypeXunHuanXingYuXinYuan        = 40703 //星语心愿
	FuncOpenTypeXunHuanBaiBeiFanLi          = 40800 //百倍返利
	FuncOpenTypeXunHuanChaoZhiHuanGou       = 40801 //超值换购
	FuncOpenTypeXunHuanZaJinZhu             = 40900 //砸金猪
	FuncOpenTypeXunHuanMeiLiPaiHang         = 40902 //魅力排行
	FuncOpenTypeXunHuanBiaoBaiPaiHang       = 40903 //表白排行
	FuncOpenTypeXunHuanShenMiBox            = 40904 //循环-神秘宝箱
	FuncOpenTypeHouseExtended               = 40905 //房产活动
	FuncOpenTypeXunHuanMount                = 40906 // 坐骑进阶
	FuncOpenTypeXunHuanWing                 = 40907 // 战翼进阶
	FuncOpenTypeXunHuanAnQi                 = 40908 // 暗器进阶
	FuncOpenTypeXunHuanXianTi               = 40909 // 仙体进阶
	FuncOpenTypeXunHuanLingYu               = 40910 // 领域进阶
	FuncOpenTypeXunHuanShenFa               = 40911 // 身法进阶
	FuncOpenTypeXunHuanFaBao                = 40912 // 法宝进阶
	FuncOpenTypeXunHuanLingQi               = 40913 // 灵骑进阶
	FuncOpenTypeXunHuanShiHunFan            = 40914 // 噬魂进阶
	FuncOpenTypeXunHuanLingYi               = 40915 // 灵翼进阶
	FuncOpenTypeXunHuanTianMoTi             = 40916 // 天魔进阶
	FuncOpenTypeXunHuanLingBing             = 40917 // 灵兵进阶
	FuncOpenTypeXunHuanLingBao              = 40918 // 灵宝进阶
	FuncOpenTypeXunHuanLingTi               = 40919 // 灵体进阶
	FuncOpenTypeXunHuanLingTongYu           = 40920 // 灵域进阶
	FuncOpenTypeXunHuanLingShen             = 40921 // 灵身进阶
	FuncOpenTypeChildrenDay                 = 40930 // 六一狂欢
	FuncOpenTypeChildrenDayBinFenTangGuo    = 40931 // 缤纷糖果
	FuncOpenTypeChildrenDayHuanLeQiQiu      = 40932 // 欢乐气球
	FuncOpenTypeChildrenDayLiuYiLiBao       = 40933 // 六一礼包
	FuncOpenTypeChildrenDayChongZhiFangSong = 40934 // 充值放送
	FuncOpenTypeDuanWuZongQingDuanWu        = 40940 // 粽情端午
	FuncOpenTypeDuanWuDuanWuMingRen         = 40941 // 端午名人
	FuncOpenTypeDuanWuSaiLongZhou           = 40942 // 赛龙舟
	FuncOpenTypeDuanWuWuWeiZongHe           = 40943 // 五味粽盒
	FuncOpenTypeDuanWuChongZhiFangSong      = 40944 // 充值放送
	FuncOpenTypeDuanWuQiYuDao               = 40945 // 奇遇岛
	FuncOpenType618KuangHuan                = 40950 // 618狂欢
	FuncOpenTypeJueZhan618                  = 40951 //决战618
	FuncOpenTypeQiangHongBao                = 40952 // 抢红包
	FuncOpenTypeNianZhongChongZhiFangSong   = 40953 // 充值放送
	FuncOpenTypeNianZhongDaZu               = 40954 // 年中大促
	FuncOpenType618MuBiao                   = 40955 // 618目标
	FuncOpenTypeXiaLingYing                 = 40960 // 夏令营
	FuncOpenTypeQingLiangYiXia              = 40961 // 清凉一夏
	FuncOpenTypeHaiLuoZhiSheng              = 40962 // 海螺之声
	FuncOpenTypeHaiDaoTanXian               = 40963 // 海岛探险
	FuncOpenTypeShaTanWaBao                 = 40964 // 沙滩挖宝
	FuncOpenTypeChongZhiFangSong            = 40965 // 充值放送
	FuncOpenTypeWuLianDianFeng              = 40970 // 武炼巅峰
	FuncOpenTypeZhanLeiGu                   = 40971 // 战鼓擂
	FuncOpenTypeYiQuanZhiBa                 = 40972 // 一拳制霸
	FuncOpenTypeTianFengMuBiao              = 40973 // 巅峰目标
	FuncOpenTypeWuLianChongZhiFangSong      = 40974 // 武炼巅峰之充值放送
	FuncOpenTypeZhanLiLiBao                 = 40975 // 战力礼包
	FuncOpenTypeLeiTaiZhuLi                 = 40976 // 擂台助力
	FuncOpenTypeMerge                       = 41000 //合服庆典
	FuncOpenTypeArenaPvpActivity            = 41001 //比武助力
	FuncOpenTypeMergeActivity               = 41010 //合服乐翻天
	FuncOpenTypeBindGoldBowl                = 41020 //绑元聚宝盆
	FuncOpenTypeGoldBowl                    = 41021 //元宝聚宝盆
	FuncOpenTypeLuckyDrew                   = 41030 //幸运大转盘
	FuncOpenTypeZhaKuang                    = 41040 //炸矿
	FuncOpenTypeAdvancedBless               = 41050 //升阶狂欢日
	FuncOpenTypeAdvancedRew                 = 41051 //升阶狂欢
	FuncOpenTypeAdvancedReturn              = 41052 //升阶返还
	FuncOpenTypeAdvancedCrit                = 41053 //暴击日
	FuncOpenTypeAdvancedBlessDay            = 41054 //祝福日
	FuncOpenTypeAdvancedBlessDayMax         = 41055 //祝福日（最近档次）
	FuncOpenTypeMergeXunHuanBaoJi           = 41056 //暴击日
	FuncOpenTypeMergeXunHuanAdvancedReturn  = 41057 //升阶返还
	FuncOpenTypeMergeXunHuanBlessDay        = 41058 //祝福日
	FuncOpenTypeFashionShop                 = 41060 //时装商店
	FuncOpenTypeFashionShopXunHuan          = 41061 //时装商店(循环活动)
	FuncOpenTypeZhuWeiChengZhan             = 41070 //助威城战
	FuncOpenTypeHeFuQiangGou                = 41080 //合服抢购
	FuncOpenTypeMergeHeFuQiangGou           = 41081 //合服抢购(合服循环)
	FuncOpenTypeSanJieMiBao                 = 41101 //三界秘宝
	FuncOpenTypeXianRenZhiLu                = 41102 //仙人指路
	FuncOpenTypeLianJinLu                   = 41103 //炼金炉
	FuncOpenTypeXiuXianDianJi               = 41104 //修仙典籍
	FuncOpenTypeTongTianTa                  = 41105 //通天塔
	FuncOpenTypeYunYinXianRen               = 41106 //云隐仙人
	FuncOpenTypeSanJieJinBang               = 41107 //三界金榜
	FuncOpenTypeMergeInvest                 = 41201 //合服投资
	FuncOpenTypeZhuanshengChongci           = 41202 //转生冲刺
	FuncOpenTypeShengJieHaoLi               = 42001 //升阶豪礼
	FuncOpenTypeXinYuXinYuan                = 42002 //星语心愿
	FuncOpenTypeXuanHuanChongZhi            = 42003 //循环充值
	FuncOpenTypeChaoJiHuanGou               = 42004 //超值换购
	FuncOpenTypeHuanLeDaLie                 = 43001 //欢乐打猎
	FuncOpenTypeZhuanZhuanLe                = 43002 //转转乐
	FuncOpenTypeXiaoLiFeiBiao               = 43003 //小李飞镖
	FuncOpenTypeHuanLeLieQiang              = 43004 //欢乐猎枪
	FuncOpenTypeGuaGuaLe                    = 43005 //刮刮乐
	FuncOpenTypeKaiFuZhaKuang               = 43006 //开服炸矿
	FuncOpenTypeXiangYunZhuanPan            = 43007 //幸运转盘
	FuncOpenTypeBaoKuCritDay                = 43008 //宝库暴击
	FuncOpenTypeSinglePlay                  = 50001 //单人玩法
	FuncOpenTypeContendPK                   = 50101 //竞争pk
	FuncOpenTypeOpenActivityCombine         = 50102 //开服乐翻天
	FuncOpenTypeXianFu                      = 50103 //仙府
	FuncOpenTypeChallengeQuest              = 50104 //挑战任务
	FuncOpenTypeLingTongYiJiMenu            = 50105 //灵童一级菜单
	FuncOpenTypeMarryMenu                   = 50106 //结婚
	FuncOpenTypeForgeTotalEntra             = 50107 //锻造总入口
	FuncOpenTypeBossEntra                   = 50108 //BOSS总入口
)

func (ft FuncOpenType) Valid() bool {
	switch ft {
	case FuncOpenTypeMain,
		FuncOpenTypeFunc,
		FuncOpenTypeQuest,
		FuncOpenTypeInventory,
		FuncOpenTypeAlliance,
		FuncOpenTypeRole,
		FuncOpenTypeSheZhi,
		FuncOpenTypeShangCheng,
		FuncOpenTypesFriend,
		FuncOpenTypeTeam,
		FuncOpenTypeSkill,
		FuncOpenTypeJueXue,
		FuncOpenTypeXinFa,
		FuncOpenTypeMount,
		FuncOpenTypeMountUpgrade,
		FuncOpenTypeMountCul,
		FuncOpenTypeMountEquipment,
		FuncOpenTypeEquipmentStrength,
		FuncOpenTypeEquimentStar,
		FuncOpenTypeEquimentXiLian,
		FuncOpenTypeEquimentKaiGuang,
		FuncOpenTypeBaGuaXiangQian,
		FuncOpenTypeSoul,
		FuncOpenTypeSoulUpgrade,
		FuncOpenTypeSoulAwaken,
		FuncOpenTypeSoulStrength,
		FuncOpenTypeBodyShield,
		FuncOpenTypeBodyShieldCul,
		FuncOpenTypeDanEat,
		FuncOpenTypeDanAlchemy,
		FuncOpenTypeWeapon,
		FuncOpenTypeWeaponStar,
		FuncOpenTypeWeaponAwaken,
		FuncOpenTypeGemDigged,
		FuncOpenTypeGemEmbedded,
		FuncOpenTypeWing,
		FuncOpenTypeWingUpgrade,
		FuncOpenTypeTitle,
		FuncOpenTypeFashion,
		FuncOpenTypeJingJie,
		FuncOpenTypeSynthesis,
		FuncOpenTypeActivity,
		FuncOpenTypeYueXiaQingYuan,
		FuncOpenTypeJiuXiaoChengZhan,
		FuncOpenTypeQiangLongYi,
		FuncOpenTypeTianJieTa,
		FuncOpenTypeTuMo,
		FuncOpenTypeXinaFu,
		FuncOpenTypeDiHunYiJi,
		FuncOpenTypePaiHangBang,
		FuncOpenTypeResFind,
		FuncOpenTypeTinaJiPai,
		FuncOpenTypeZhuoYaoJi,
		FuncOpenTypeFourGod,
		FuncOpenTypeArena,
		FuncOpenTypePersonalTransportation,
		FuncOpenTypeAllianceTransportation,
		FuncOpenTypeDragon,
		FuncOpenTypeWorldBoss,
		FuncOpenTypeLingYu,
		FuncOpenTypeShenfa,
		FuncOpenTypeMarry,
		FuncOpenTypeMarryPre,
		FuncOpenTypeShield,
		FuncOpenTypeFeather,
		FuncOpenTypeYuanGodGold,
		FuncOpenTypeYuanGodStrengthen,
		FuncOpenTypeYuanGodRecasting,
		FuncOpenTypeYuanGodEat,
		FuncOpenTypeLingChiFighting,
		FuncOpenTypeShenfaAdvanced,
		FuncOpenTypeLingYuAdvanced,
		FuncOpenTypeExpFuBen,
		FuncOpenTypeSilverFuBen,
		FuncOpenTypeCangLongQiJu,
		FuncOpenTypeFirstCharge,
		FuncOpenTypeWelfareHall,
		FuncOpenTypeWelfareLogin,
		FuncOpenTypeWelfareUplevel,
		FuncOpenTypeWelfareOnline,
		FuncOpenTypeWelfareSign,
		FuncOpenTypeWelfareGiftBag,
		FuncOpenTypeOpenActivity,
		FuncOpenTypeInvestLevel,
		FuncOpenTypeInvestSevenDay,
		FuncOpenTypeFeedbackCharge,
		FuncOpenTypeFeedbackCost,
		FuncOpenTypeRankCharge,
		FuncOpenTypeRankCost,
		FuncOpenTypeRankSevenDay,
		FuncOpenTypeRankMount,
		FuncOpenTypeRankWing,
		FuncOpenTypeRankBodyShield,
		FuncOpenTypeRankLingyu,
		FuncOpenTypeRankShenfa,
		FuncOpenTypeRankFeather,
		FuncOpenTypeRankShield,
		FuncOpenTypeAnQi,
		FuncOpenTypeAnQiAdvanced,
		FuncOpenTypeTianShu,
		FuncOpenTypeCoupleFuBen,
		FuncOpenTypeCrossWorldBoss,
		FuncOpenTypeBloodPool,
		FuncOpenTypeZhuanSheng,
		FuncOpenTypeFirstChargeCycleDay,
		FuncOpenTypeVIPPluss,
		//FuncOpenTypeFashionStar,
		FuncOpenTypeCrossTuLong,
		FuncOpenTypeSinglePlay,
		FuncOpenTypeContendPK,
		FuncOpenTypeMerge,
		FuncOpenTypeMergeActivity,
		FuncOpenTypeBindGoldBowl,
		FuncOpenTypeGoldBowl,
		FuncOpenTypeLuckyDrew,
		FuncOpenTypeZhaKuang,
		FuncOpenTypeAdvancedBless,
		FuncOpenTypeCrossLianYu,
		FuncOpenTypeMassacre,
		FuncOpenTypeGoldEquiBag,
		FuncOpenTypeDaBaoTower,
		FuncOpenTypeMyBoss,
		FuncOpenTypeBoss,
		FuncOpenTypeOpenActivityCombine,
		FuncOpenTypeGodSiege,
		FuncOpenTypeAnQiJiGuan,
		FuncOpenTypeWingRune,
		FuncOpenTypeMountSkill,
		FuncOpenTypeAnQiSkill,
		FuncOpenTypeWingSkill,
		FuncOpenTypeYuanGodUpgrade,
		FuncOpenTypeYuanGodKaiGuang,
		FuncOpenTypeLongFengChengXiangOpen,
		FuncOpenTypeLongFengChengXiangMerge,
		FuncOpenTypeLuckyTurnplate,
		FuncOpenTypeHuHuShengFeng,
		FuncOpenTypeHeiHuoShop,
		FuncOpenTypeFeiHuangTengDa,
		FuncOpenTypeSingleCharge,
		FuncOpenTypeFaBao,
		FuncOpenTypeFaBaoAdvanced,
		FuncOpenTypeFaBaoUnreal,
		FuncOpenTypeFaBaoTongLing,
		FuncOpenTypeFaBaoSuit,
		FuncOpenTypeFaBaoSkill,
		FuncOpenTypeXianTi,
		FuncOpenTypeXianTiAdvanced,
		FuncOpenTypeXianTiUnreal,
		FuncOpenTypeXianTiSkill,
		FuncOpenTypeXianTiLingYu,
		FuncOpenTypeLiveness,
		FuncOpenTypeXueDu,
		FuncOpenTypeFuBen,
		FuncOpenTypeAdvanced,
		FuncOpenTypeMingRiKaiQi,
		FuncOpenTypeExpBox,
		FuncOpenTypeCaiLiaoFuBen,
		FuncOpenTypeCaiLiaoMount,
		FuncOpenTypeCaiLiaoWing,
		FuncOpenTypeCaiLiaoShenFa,
		FuncOpenTypeCaiLiaoFaBao,
		FuncOpenTypeCaiLiaoXianTi,
		FuncOpenTypeBossHuanJing,
		FuncOpenTypeOutlandBoss,
		FuncOpenTypeYiZheTeHui,
		FuncOpenTypeSystemJinJie,
		FuncOpenTypeQiTianLeiChong,
		FuncOpenTypeXinaShiXuYuan,
		FuncOpenTypeYuanBaoSongBuTing,
		FuncOpenTypeNiChongWoSong,
		FuncOpenTypeShengJieSongJiang,
		FuncOpenTypeYaoHu,
		FuncOpenTypeGoldLaBa,
		FuncOpenTypeChristmasParty,
		FuncOpenTypeChristmasXueRen,
		FuncOpenTypeChristmasTree,
		FuncOpenTypeChristmasGift,
		FuncOpenTypeChristmasShop,
		FuncOpenTypeChristmasGoldCar,
		FuncOpenTypeDailyQuest,
		FuncOpenTypeBaGuaMiJing,
		FuncOpenTypeSystemJinJieMount,
		FuncOpenTypeSystemJinJieWing,
		FuncOpenTypeSystemJinJieAnQi,
		FuncOpenTypeSystemJinJieBodyShiled,
		FuncOpenTypeSystemJinJieLingYu,
		FuncOpenTypeSystemJinJieShenFa,
		FuncOpenTypeSystemJinJieFaBao,
		FuncOpenTypeRankAnQi,
		FuncOpenTypeRankFaBao,
		FuncOpenTypeRankXianTi,
		FuncOpenTypeXianFu,
		FuncOpenTypeQuiz,
		FuncOpenTypeTeamCopy,
		FuncOpenTypeNewYearParty,
		FuncOpenTypeNewYearJinZhuSongLi,
		FuncOpenTypeNewYearJinZhuNaoCheng,
		FuncOpenTypeNewYearDaPaoDaQiu,
		FuncOpenTypeNewYearKuangHuan,
		FuncOpenTypeNewYearZhaoCaiJinBao,
		FuncOpenTypeChallengeQuest,
		FuncOpenTypeDenseWat,
		FuncOpenTypeDianXing,
		FuncOpenTypeDianXingJieFeng,
		FuncOpenTypeTeamCopySilver,
		FuncOpenTypeTeamCopyXingChen,
		FuncOpenTypeTeamCopyXueMo,
		FuncOpenTypeTeamCopyZhuangShengEquip,
		FuncOpenTypeTeamCopyWeapon,
		FuncOpenTypeCaiLiaoTianMo,
		FuncOpenTypeWardrobe,
		FuncOpenTypeFaBaoShengJi,
		FuncOpenTypeXianTiShengJi,
		FuncOpenTypeLingYuShengJi,
		FuncOpenTypeShenfaShengJi,
		FuncOpenTypeTianMo,
		FuncOpenTypeTianMoAdvanced,
		FuncOpenTypeTianMoUnreal,
		FuncOpenTypeTianMoUplevel,
		FuncOpenTypeTianMoEquipment,
		FuncOpenTypeTianMoSkill,
		FuncOpenTypeShiHunFan,
		FuncOpenTypeShiHunFanAdvanced,
		FuncOpenTypeShiHunFanUnreal,
		FuncOpenTypeShiHunFanShengJi,
		FuncOpenTypeShiHunFanEquip,
		FuncOpenTypeShiHunFanSkill,
		FuncOpenTypeContinueCharge,
		FuncOpenTypeMadeResource,
		FuncOpenTypeCycleSingleCharge,
		FuncOpenTypeHuanLeDaLie,
		FuncOpenTypeZhuanZhuanLe,
		FuncOpenTypeXiaoLiFeiBiao,
		FuncOpenTypeHuanLeLieQiang,
		FuncOpenTypeGuaGuaLe,
		FuncOpenTypeKaiFuZhaKuang,
		FuncOpenTypeXiangYunZhuanPan,
		FuncOpenTypeBossFirstKill,
		FuncOpenTypeKanJiaGift,
		FuncOpenTypeChaoZhiTaoCan,
		FuncOpenTypeYangJinJi,
		FuncOpenTypeLingQiChongCi,
		FuncOpenTypeShiHunChongCi,
		FuncOpenTypeLingYi,
		FuncOpenTypeChongCi,
		FuncOpenTypeLingBingChongCi,
		FuncOpenTypeLingBaoChongCi,
		FuncOpenTypeLingTiChongCi,
		FuncOpenTypeLingYuChongCi,
		FuncOpenTypeLingShenChongCi,
		FuncOpenTypeLingQiJinJie,
		FuncOpenTypeShiHunJinJie,
		FuncOpenTypeLingYiJinJie,
		FuncOpenTypeLingMoJinJie,
		FuncOpenTypeLingBingJinJie,
		FuncOpenTypeLingBaoJinJie,
		FuncOpenTypeLingTiJinJie,
		FuncOpenTypeLingYuJinJie,
		FuncOpenTypeLingShenJinJie,
		FuncOpenTypeFashionShop,
		FuncOpenTypeZhuWeiChengZhan,
		FuncOpenTypeHeFuQiangGou,
		FuncOpenTypeCrazyBox,
		FuncOpenTypeRankLevel,
		FuncOpenTypeChongZhiSongFangZi,
		FuncOpenTypeZaJinDan,
		FuncOpenTypeMingRenPu,
		FuncOpenTypeShouChongFanGold,
		FuncOpenTypeShouChongFanBindGold,
		FuncOpenTypeLingTongYiJiMenu,
		FuncOpenTypeLingTong,
		FuncOpenTypeLingTongUpgrade,
		FuncOpenTypeLingTongPeiYang,
		FuncOpenTypeLingTongHuanHua,
		FuncOpenTypeLingTongWeapon,
		FuncOpenTypeLingTongWeaponAdvanced,
		FuncOpenTypeLingTongWeaponUnreal,
		FuncOpenTypeLingTongWeaponUpgrade,
		FuncOpenTypeLingTongWeaponEquip,
		FuncOpenTypeLingTongWeaponSkill,
		FuncOpenTypeLingTongMount,
		FuncOpenTypeLingTongMountAdvanced,
		FuncOpenTypeLingTongMountUnreal,
		FuncOpenTypeLingTongMountUpgrade,
		FuncOpenTypeLingTongMountEquip,
		FuncOpenTypeLingTongMountSkill,
		FuncOpenTypeLingTongMountPeiYang,
		FuncOpenTypeLingTongWing,
		FuncOpenTypeLingTongWingAdvanced,
		FuncOpenTypeLingTongWingUnreal,
		FuncOpenTypeLingTongWingUpgrade,
		FuncOpenTypeLingTongWingEquip,
		FuncOpenTypeLingTongWingSkill,
		FuncOpenTypeLingTongShenFa,
		FuncOpenTypeLingTongShenFaAdvanced,
		FuncOpenTypeLingTongShenFaUnreal,
		FuncOpenTypeLingTongShenFaUpgrade,
		FuncOpenTypeLingTongShenFaEquip,
		FuncOpenTypeLingTongShenFaSkill,
		FuncOpenTypeLingTongLingYu,
		FuncOpenTypeLingTongLingYuAdvanced,
		FuncOpenTypeLingTongLingYuUnreal,
		FuncOpenTypeLingTongLingYuUpgrade,
		FuncOpenTypeLingTongLingYuEquip,
		FuncOpenTypeLingTongLingYuSkill,
		FuncOpenTypeLingTongFaBao,
		FuncOpenTypeLingTongFaBaoAdvanced,
		FuncOpenTypeLingTongFaBaoUnreal,
		FuncOpenTypeLingTongFaBaoUpgrade,
		FuncOpenTypeLingTongFaBaoEquip,
		FuncOpenTypeLingTongFaBaoSkill,
		FuncOpenTypeLingTngFaBaoTongLing,
		FuncOpenTypeLingTongXianTi,
		FuncOpenTypeLingTongXianTiAdvanced,
		FuncOpenTypeLingTongXianTiUnreal,
		FuncOpenTypeLingTongXianTiUpgrade,
		FuncOpenTypeLingTongXianTiEquip,
		FuncOpenTypeLingTongXianTiSkill,
		FuncOpenTypeShouChongFanBei,
		FuncOpenTypeAdvancedRew,
		FuncOpenTypeAdvancedReturn,
		FuncOpenTypeAdvancedCrit,
		FuncOpenTypeAdvancedBlessDay,
		FuncOpenTypeCaiLiaoLingTong,
		FuncOpenTypeCaiLiaoLingBing,
		FuncOpenTypeCaiLiaoLingYu,
		FuncOpenTypeTeamCopyLingTong,
		FuncOpenTypeTeamCopyStrength,
		FuncOpenTypeTeamCopyUpstar,
		FuncOpenTypeShenMoWar,
		FuncOpenTypeFeiSheng,
		FuncOpenTypeKaiFuMuBiao,
		FuncOpenTypeHongBao,
		FuncOpenTypeTelent,
		FuncOpenTypeHuaLingMount,
		FuncOpenTypeHuaLingWing,
		FuncOpenTypeHuaLingShenFa,
		FuncOpenTypeHuaLingAnQi,
		FuncOpenTypeHuaLingFaBao,
		FuncOpenTypeHuaLingXianTi,
		FuncOpenTypeHuaLingLingYu,
		FuncOpenTypeHuaLingShiHunFan,
		FuncOpenTypeHuaLingTianMoTi,
		FuncOpenTypeHuaLingLingBing,
		FuncOpenTypeHuaLingLingQi,
		FuncOpenTypeHuaLingLingWing,
		FuncOpenTypeHuaLingLingShen,
		FuncOpenTypeHuaLingLingArea,
		FuncOpenTypeHuaLingLingBao,
		FuncOpenTypeHuaLingLingTi,
		FuncOpenTypeCangJingGe,
		FuncOpenTypeMountUnreal,
		FuncOpenTypeWingShengJi,
		FuncOpenTypeShenFaUnreal,
		FuncOpenTypeShenFaEquip,
		FuncOpenTypeShenFaSkill,
		FuncOpenTypeLingYuUnreal,
		FuncOpenTypeLingYuEquip,
		FuncOpenTypeLingYuSkill,
		FuncOpenTypeAnQiPeiYang,
		FuncOpenTypeAnQiShengJi,
		FuncOpenTypeDingZhiTitle,
		FuncOpenTypeJingCaiHuoDong,
		FuncOpenTypeXunHuanCycleSingle,
		FuncOpenTypeXunHuanShouChongFanYuanBao,
		FuncOpenTypeXunHuanShouChongFanBangYuan,
		FuncOpenTypeXunHuanChongZhiPaiHang,
		FuncOpenTypeXunHuanJinJieFanHuan,
		FuncOpenTypeXunHuanHuanLeChouJiang,
		FuncOpenTypeXunHuanMoJin,
		FuncOpenTypeXunHuanGuaGuaLe,
		FuncOpenTypeXunHuanZhiZhunLunPan,
		FuncOpenTypeXunHuanQuanMinQiangGou,
		FuncOpenTypeXunHuanJuHuaSuan,
		FuncOpenTypeXunHuanXianShiLiBao,
		FuncOpenTypeXunHuanMaiYiSongYi,
		FuncOpenTypeXunHuanZhaoCaiJinBao,
		FuncOpenTypeXunHuanHuangJinKuangGong,
		FuncOpenTypeXunHuanShenMiBaoXiang,
		FuncOpenTypeXunHuanYiXiaoBoDa,
		FuncOpenTypeAllianceDaily,
		FuncOpenTypeAllianceBoss,
		FuncOpenTypeEquipBaoKu,
		FuncOpenTypeGemTaoZhuang,
		FuncOpenTypeYuanGodQiangHuaTaoZhuang,
		FuncOpenTypeYuanGodOpenLightTaoZhuang,
		FuncOpenTypeYuanGodUpstarTaoZhuang,
		FuncOpenTypeAllianceAltar,
		FuncOpenTypeGoldEquipExtend,
		FuncOpenTypeShenQi,
		FuncOpenTypeShenQiSmelt,
		FuncOpenTypeShenQiQiLing,
		FuncOpenTypeShenQiZhuLing,
		FuncOpenTypeShenQiTaoZhuang,
		FuncOpenTypeShenQiResolve,
		FuncOpenTypeShenQiXunBao,
		FuncOpenTypeYingLingPu,
		FuncOpenTypeZhiZhunYingLingPu,
		FuncOpenTypeShengHen,
		FuncOpenTypeShengHenXunBao,
		FuncOpenTypeTuLongEquip,
		FuncOpenTypeTuLongQiangHua,
		FuncOpenTypeTuLongRongHe,
		FuncOpenTypeTuLongZhuanHua,
		FuncOpenTypeTuLongSkill,
		FuncOpenTypeTuLongXunBao,
		FuncOpenTypeMingGe,
		FuncOpenTypeMingGong,
		FuncOpenTypeMingPan,
		FuncOpenTypeMingGeHeCheng,
		FuncOpenTypeMingGeXunBao,
		FuncOpenTypeZhenFa,
		FuncOpenTypeZhenQi,
		FuncOpenTypeZhenHuo,
		FuncOpenTypeZhenFaSuit,
		FuncOpenTypeXunHuanXunHuanChongZhi,
		FuncOpenTypeXunHuanJiFenShangCheng,
		FuncOpenTypeXunHuanShengJieHaoLi,
		FuncOpenTypeXunHuanXingYuXinYuan,
		FuncOpenTypeXunHuanBaiBeiFanLi,
		FuncOpenTypeXunHuanChaoZhiHuanGou,
		FuncOpenTypeXunHuanZaJinZhu,
		FuncOpenTypeXunHuanMingRenPu,
		FuncOpenTypeXunHuanMeiLiPaiHang,
		FuncOpenTypeXunHuanBiaoBaiPaiHang,
		FuncOpenTypeDongFang,
		FuncOpenTypeBaoBao,
		FuncOpenTypeBaoBaoLearn,
		FuncOpenTypeBaoBaoToy,
		FuncOpenTypeBaoBaoZhuanShi,
		FuncOpenTypeTrad,
		FuncOpenTypeDingQing,
		FuncOpenTypeMarryJiNian,
		FuncOpenTypeMarryMenu,
		FuncOpenTypeRankCharm,
		FuncOpenTypeRankBiaoBai,
		FuncOpenTypeLingTongWeaponEquipUpgrade,
		FuncOpenTypeLingTongWeaponUnrealLabel,
		FuncOpenTypeLingTongWeaponExpert,
		FuncOpenTypeLingTongWeaponAdvencedRew,
		FuncOpenTypeLingTongWingEquipUpgrade,
		FuncOpenTypeLingTongWingUnrealLabel,
		FuncOpenTypeLingTongWingExpert,
		FuncOpenTypeLingTongWingAdvencedRew,
		FuncOpenTypeLingTongShenFaEquipUpgrade,
		FuncOpenTypeLingTongShenFaUnrealLabel,
		FuncOpenTypeLingTongShenFaExpert,
		FuncOpenTypeLingTongShenFaAdvencedRew,
		FuncOpenTypeLingTongLingYuEquipUpgrade,
		FuncOpenTypeLingTongLingYuUnrealLabel,
		FuncOpenTypeLingTongLingYuExpert,
		FuncOpenTypeLingTongLingYuAdvencedRew,
		FuncOpenTypeLingTongFaBaoEquipUpgrade,
		FuncOpenTypeLingTongFaBaoTongLingLabel,
		FuncOpenTypeLingTongFaBaoExpert,
		FuncOpenTypeLingTongFaBaoAdvencedRew,
		FuncOpenTypeLingTongXianTiEquipUpgrade,
		FuncOpenTypeLingTongXianTiUnrealLabel,
		FuncOpenTypeLingTongXianTiExpert,
		FuncOpenTypeLingTongXianTiAdvencedRew,
		FuncOpenTypeLingTongMountEquipUpgrade,
		FuncOpenTypeLingTongMountExpert,
		FuncOpenTypeLingTongMountAdvencedRew,
		FuncOpenTypeWingRuneUpgrade,
		FuncOpenTypeWingAdvencedLabel,
		FuncOpenTypeWingUnrealLabel,
		FuncOpenTypeShenFaEquipUpgrade,
		FuncOpenTypeShenFaUnrealLabel,
		FuncOpenTypeShenFaExpert,
		FuncOpenTypeLingYuEquipUpgrade,
		FuncOpenTypeLingYuUnrealLabel,
		FuncOpenTypeLingYuExpert,
		FuncOpenTypeAnQiEquipUpgrade,
		FuncOpenTypeAnQiExpert,
		FuncOpenTypeAnQiAdvencedRew,
		FuncOpenTypeFaBaoEquipUpgrade,
		FuncOpenTypeFaBaoUnrealLabel,
		FuncOpenTypeXianTiEquipUpgrade,
		FuncOpenTypeXianTiUnrealLabel,
		FuncOpenTypeXianTiExpert,
		FuncOpenTypeShiHunFanEquipUpgrade,
		FuncOpenTypeShiHunFanExpert,
		FuncOpenTypeShiHunFanAdvencedRew,
		FuncOpenTypeTianMoEquipUpgrade,
		FuncOpenTypeTianMoExpert,
		FuncOpenTypeTianMoAdvencedRew,
		FuncOpenTypeWingExpert,
		FuncOpenTypeWingAdvencedRew,
		FuncOpenTypeShenFaAdvencedRew,
		FuncOpenTypeLingYuAdvencedRew,
		FuncOpenTypeFaBaoExpert,
		FuncOpenTypeFaBaoAdvencedRew,
		FuncOpenTypeXianTiAdvencedRew,
		FuncOpenTypeMountJinJie,
		FuncOpenTypeMountHuanHua,
		FuncOpenTypeMountEquipQiangHua,
		FuncOpenTypeMountDaRen,
		FuncOpenTypeMountJinJieRew,
		FuncOpenTypeXunHuanShenMiBox,
		FuncOpenTypeXunHuanRankCost,
		FuncOpenTypeFuQiFuBen,
		FuncOpenTypeMarryHunJie,
		FuncOpenTypeMarryAiQingShu,
		FuncOpenTypeCycleSingleChargeTwo,
		FuncOpenTypeSingleChargeMaxRew,
		FuncOpenTypeBangYuanFanLi,
		FuncOpenTypeYuanBaoFanLi,
		FuncOpenTypeBaoKuCritDay,
		FuncOpenTypeMarryBiaoBai,
		FuncOpenTypeAdvancedBlessDayMax,
		FuncOpenTypeNiChongWoSongTwo,
		FuncOpenTypeXianTaoDaHui,
		FuncOpenTypeShenYuZhiZhan,
		FuncOpenTypeLongGongTanBao,
		FuncOpenTypePrivateCustomer,
		FuncOpenTypeFashionShopXunHuan,
		FuncOpenTypeMyHouse,
		FuncOpenTypeMergeXunHuanBaoJi,
		FuncOpenTypeMergeXunHuanAdvancedReturn,
		FuncOpenTypeMergeXunHuanBlessDay,
		FuncOpenTypeMergeHeFuQiangGou,
		FuncOpenTypeHouseExtended,
		FuncOpenTypeXunHuanMount,
		FuncOpenTypeXunHuanWing,
		FuncOpenTypeXunHuanAnQi,
		FuncOpenTypeXunHuanXianTi,
		FuncOpenTypeXunHuanLingYu,
		FuncOpenTypeXunHuanShenFa,
		FuncOpenTypeXunHuanFaBao,
		FuncOpenTypeXunHuanLingQi,
		FuncOpenTypeXunHuanShiHunFan,
		FuncOpenTypeXunHuanLingYi,
		FuncOpenTypeXunHuanTianMoTi,
		FuncOpenTypeXunHuanLingBing,
		FuncOpenTypeXunHuanLingBao,
		FuncOpenTypeXunHuanLingTi,
		FuncOpenTypeXunHuanLingTongYu,
		FuncOpenTypeXunHuanLingShen,
		FuncOpenTypeShenZhuMount,
		FuncOpenTypeShenZhuWing,
		FuncOpenTypeShenZhuAnQi,
		FuncOpenTypeShenZhuLingYu,
		FuncOpenTypeShenZhuFaBao,
		FuncOpenTypeShenZhuXianTi,
		FuncOpenTypeShenZhuShenFa,
		FuncOpenTypeShenZhuLingTongMount,
		FuncOpenTypeShenZhuShiHunFan,
		FuncOpenTypeShenZhuLingTongWing,
		FuncOpenTypeShenZhuTianMoTi,
		FuncOpenTypeShenZhuLingTongWeapon,
		FuncOpenTypeShenZhuLingTongFaBao,
		FuncOpenTypeShenZhuLingTongXianTi,
		FuncOpenTypeShenZhuLingTongLingYu,
		FuncOpenTypeShenZhuLingTongShenFa,
		FuncOpenTypeTongLingMount,
		FuncOpenTypeTongLingWing,
		FuncOpenTypeTongLingAnQi,
		FuncOpenTypeTongLingLingYu,
		FuncOpenTypeTongLingFaBao,
		FuncOpenTypeTongLingXianTi,
		FuncOpenTypeTongLingShenFa,
		FuncOpenTypeTongLingLingTongMount,
		FuncOpenTypeTongLingShiHunFan,
		FuncOpenTypeTongLingLingTongWing,
		FuncOpenTypeTongLingTianMoTi,
		FuncOpenTypeTongLingLingTongWeapon,
		FuncOpenTypeTongLingLingTongFaBao,
		FuncOpenTypeTongLingLingTongXianTi,
		FuncOpenTypeTongLingLingTongLingYu,
		FuncOpenTypeTongLingLingTongShenFa,
		FuncOpenTypeJueXingXianTi,
		FuncOpenTypeJueXingShenFa,
		FuncOpenTypeJueXingLingTongMount,
		FuncOpenTypeJueXingShiHunFan,
		FuncOpenTypeJueXingLingTongWing,
		FuncOpenTypeJueXingTianMoTi,
		FuncOpenTypeJueXingLingTongWeapon,
		FuncOpenTypeJueXingLingTongFaBao,
		FuncOpenTypeJueXingLingTongXianTi,
		FuncOpenTypeJueXingLingTongLingYu,
		FuncOpenTypeJueXingLingTongShenFa,
		FuncOpenTypeJueXingWing,
		FuncOpenTypeJueXingMount,
		FuncOpenTypeJueXingAnQi,
		FuncOpenTypeJueXingLingYu,
		FuncOpenTypeJueXingFaBao,
		FuncOpenTypeYuXiZhiZhan,
		FuncOpenTypeChildrenDay,
		FuncOpenTypeChildrenDayBinFenTangGuo,
		FuncOpenTypeChildrenDayHuanLeQiQiu,
		FuncOpenTypeChildrenDayLiuYiLiBao,
		FuncOpenTypeChildrenDayChongZhiFangSong,
		FuncOpenTypeDuanWuZongQingDuanWu,
		FuncOpenTypeDuanWuDuanWuMingRen,
		FuncOpenTypeDuanWuSaiLongZhou,
		FuncOpenTypeDuanWuWuWeiZongHe,
		FuncOpenTypeDuanWuChongZhiFangSong,
		FuncOpenTypeDuanWuQiYuDao,
		FuncOpenTypeZhouKa,
		FuncOpenTypeFuShi,
		FuncOpenType618KuangHuan,
		FuncOpenTypeJueZhan618,
		FuncOpenTypeNianZhongDaZu,
		FuncOpenTypeQiangHongBao,
		FuncOpenTypeNianZhongChongZhiFangSong,
		FuncOpenType618MuBiao,
		FuncOpenTypeTitleUpstar,
		FuncOpenTypeLingTongUpstar,
		FuncOpenTypeLingTongEquip,
		FuncOpenTypeLingTongEquipStrengthen,
		FuncOpenTypeLingTongEquipAdvanced,
		FuncOpenTypeShenZhuLingTongEquip,
		FuncOpenTypeLingTongEquipWuXingLingZhu,
		FuncOpenTypeQiXueQiang,
		FuncOpenTypeShaLuZhiDu,
		FuncOpenTypeXiaLingYing,
		FuncOpenTypeQingLiangYiXia,
		FuncOpenTypeHaiLuoZhiSheng,
		FuncOpenTypeHaiDaoTanXian,
		FuncOpenTypeShaTanWaBao,
		FuncOpenTypeChongZhiFangSong,
		FuncOpenTypeChuangShiPreview,
		FuncOpenTypeJieYi,
		FuncOpenTypeXiongDiWeiMing,
		FuncOpenTypeWeiMingUpLev,
		FuncOpenTypeXiongDiToken,
		FuncOpenTypeTokenUpLev,
		FuncOpenTypeBiWuDaHui,
		FuncOpenTypeHuanLeJingCai,
		FuncOpenTypeWuLianDianFeng,
		FuncOpenTypeZhanLeiGu,
		FuncOpenTypeYiQuanZhiBa,
		FuncOpenTypeTianFengMuBiao,
		FuncOpenTypeWuLianChongZhiFangSong,
		FuncOpenTypeZhanLiLiBao,
		FuncOpenTypeLeiTaiZhuLi,
		FuncOpenTypeWushuangWeapon,
		FuncOpenTypeMaterialBaoKu,
		FuncOpenTypeShengJieHaoLi,
		FuncOpenTypeXinYuXinYuan,
		FuncOpenTypeXuanHuanChongZhi,
		FuncOpenTypeChaoJiHuanGou,
		FuncOpenTypeForgeTotalEntra,
		FuncOpenTypeShenQiRenZhu,
		FuncOpenTypeCycleSingleChargeThree,
		FuncOpenTypeGodCasting,
		FuncOpenTypeForgeSoul,
		FuncOpenTypeCastingSpirit,
		FuncOpenTypeGodCastingInherit,
		FuncOpenTypeZhenXiBoss,
		FuncOpneTypeShenQiRenZhuJieRi,
		FuncOpenTypeArenaPvpActivity,
		FuncOpenTypeChuangShiZhiZhan,
		FuncOpenTypeDingShiBoss,
		FuncOpenTypeShengShouBoss,
		FuncOpenTypeBossEntra,
		FuncOpenTypeNewFirstCharge,
		FuncOpenTypeSanJieMiBao,
		FuncOpenTypeXianRenZhiLu,
		FuncOpenTypeLianJinLu,
		FuncOpenTypeXiuXianDianJi,
		FuncOpenTypeTongTianTa,
		FuncOpenTypeYunYinXianRen,
		FuncOpenTypeSanJieJinBang,
		FuncOpenTypeMergeInvest,
		FuncOpenTypeZhuanshengChongci,
		FuncOpenTypeXianZunCard,
		FuncOpenTypeRing,
		FuncOpenTypeRingStrengthen,
		FuncOpenTypeRingAdvance,
		FuncOpenTypeRingFuse,
		FuncOpenTypeRingJingLing,
		FuncOpenTypeRingXunBao,
		FuncOpenTypeShangGuZhiLing,
		FuncOpenTypeShangGuZhiLingUpLevel,
		FuncOpenTypeShangGuZhiLingLingWen,
		FuncOpenTypeShangGuZhiLingUpRank,
		FuncOpenTypeShangGuZhiLingLingLian,
		FuncOpenTypeShangGuZhiLingMiLing:
		return true
	}
	return false
}

var (
	funcOpenTypeMap = map[FuncOpenType]string{
		FuncOpenTypeMain:                   "主界面",
		FuncOpenTypeFunc:                   "功能面板",
		FuncOpenTypeQuest:                  "任务对话面板",
		FuncOpenTypeInventory:              "背包",
		FuncOpenTypeAlliance:               "仙盟",
		FuncOpenTypeRole:                   "角色",
		FuncOpenTypeSheZhi:                 "设置",
		FuncOpenTypeShangCheng:             "商城",
		FuncOpenTypesFriend:                "好友",
		FuncOpenTypeTeam:                   "组队",
		FuncOpenTypeSkill:                  "技能（职业技能标签）",
		FuncOpenTypeJueXue:                 "绝学",
		FuncOpenTypeXinFa:                  "心法",
		FuncOpenTypeMount:                  "坐骑（主标签与升阶标签）",
		FuncOpenTypeMountUpgrade:           "坐骑升级（吃幻化丹）",
		FuncOpenTypeMountCul:               "坐骑培养（吃培养丹）",
		FuncOpenTypeMountEquipment:         "坐骑装备",
		FuncOpenTypeEquipmentStrength:      "锻造（装备强化标签）",
		FuncOpenTypeEquimentStar:           "装备升星",
		FuncOpenTypeEquimentXiLian:         "装备洗炼",
		FuncOpenTypeEquimentKaiGuang:       "装备开光",
		FuncOpenTypeBaGuaXiangQian:         "八卦镶嵌",
		FuncOpenTypeSoul:                   "帝魂（主标签，包含激活与佩戴）",
		FuncOpenTypeSoulUpgrade:            "帝魂升级",
		FuncOpenTypeSoulAwaken:             "帝魂顿悟",
		FuncOpenTypeSoulStrength:           "帝魂强化",
		FuncOpenTypeBodyShield:             "护体盾（主标签与升阶标签）",
		FuncOpenTypeBodyShieldCul:          "护盾培养（标签页）",
		FuncOpenTypeDanEat:                 "仙丹服用",
		FuncOpenTypeDanAlchemy:             "仙丹炼成",
		FuncOpenTypeWeapon:                 "兵魂（激活与佩戴标签）",
		FuncOpenTypeWeaponStar:             "兵魂升星",
		FuncOpenTypeWeaponAwaken:           "兵魂觉醒",
		FuncOpenTypeGemDigged:              "宝石（挖矿）",
		FuncOpenTypeGemEmbedded:            "宝石（镶嵌）",
		FuncOpenTypeWing:                   "战翼（主标签与进阶）",
		FuncOpenTypeWingUpgrade:            "战翼（吃幻化丹）",
		FuncOpenTypeTitle:                  "外观（称号标签）",
		FuncOpenTypeFashion:                "外观（时装标签）",
		FuncOpenTypeJingJie:                "境界（主标签）",
		FuncOpenTypeSynthesis:              "合成(宝石)",
		FuncOpenTypeActivity:               "活动",
		FuncOpenTypeYueXiaQingYuan:         "月下情缘",
		FuncOpenTypeJiuXiaoChengZhan:       "九霄城战",
		FuncOpenTypeQiangLongYi:            "抢龙椅",
		FuncOpenTypeTianJieTa:              "天劫塔",
		FuncOpenTypeTuMo:                   "屠魔任务",
		FuncOpenTypeXinaFu:                 "秘境仙府",
		FuncOpenTypeDiHunYiJi:              "帝陵遗迹",
		FuncOpenTypePaiHangBang:            "排行榜",
		FuncOpenTypeResFind:                "资源找回",
		FuncOpenTypeTinaJiPai:              "天机牌",
		FuncOpenTypeZhuoYaoJi:              "捉妖记",
		FuncOpenTypeFourGod:                "四神遗迹",
		FuncOpenTypeArena:                  "3V3",
		FuncOpenTypePersonalTransportation: "个人镖车",
		FuncOpenTypeAllianceTransportation: "仙盟镖车",
		FuncOpenTypeDragon:                 "神龙现世",
		FuncOpenTypeWorldBoss:              "世界BOSS",
		FuncOpenTypeShenfa:                 "身法",
		FuncOpenTypeLingYu:                 "领域",
		FuncOpenTypeMarry:                  "结婚",
		FuncOpenTypeShield:                 "神盾尖刺",
		FuncOpenTypeFeather:                "护体仙羽",
		FuncOpenTypeYuanGodGold:            "元神金装",
		FuncOpenTypeYuanGodStrengthen:      "元神金装·强化",
		FuncOpenTypeYuanGodRecasting:       "元神金装·重铸",
		FuncOpenTypeYuanGodEat:             "元神金装·吞噬",
		FuncOpenTypeLingChiFighting:        "灵池争夺",
		FuncOpenTypeShenfaAdvanced:         "身法进阶",
		FuncOpenTypeLingYuAdvanced:         "领域进阶",
		FuncOpenTypeExpFuBen:               "经验副本",
		FuncOpenTypeSilverFuBen:            "银两副本",
		FuncOpenTypeCangLongQiJu:           "苍龙棋局",
		FuncOpenTypeDaBaoTower:             "打宝塔",
		FuncOpenTypeMyBoss:                 "个人BOSS",
		FuncOpenTypeBoss:                   "BOSS",
		FuncOpenTypeFirstCharge:            "首冲奖励",
		FuncOpenTypeWelfareHall:            "福利大厅",
		FuncOpenTypeWelfareLogin:           "登录奖励",
		FuncOpenTypeWelfareUplevel:         "升级礼包",
		FuncOpenTypeWelfareOnline:          "在线礼包",
		FuncOpenTypeWelfareSign:            "签到(暂略",
		FuncOpenTypeWelfareGiftBag:         "激活码兑换",
		FuncOpenTypeOpenActivity:           "开服活动",
		FuncOpenTypeInvestLevel:            "投资计划",
		FuncOpenTypeInvestSevenDay:         "7日投资",
		FuncOpenTypeFeedbackCharge:         "充值返利",
		FuncOpenTypeFeedbackCost:           "消费返利",
		FuncOpenTypeRankCharge:             "充值排行",
		FuncOpenTypeRankCost:               "消费排行",
		FuncOpenTypeRankSevenDay:           "7日冲刺",
		FuncOpenTypeRankMount:              "坐骑排行",
		FuncOpenTypeRankWing:               "战翼排行",
		FuncOpenTypeRankBodyShield:         "护盾排行",
		FuncOpenTypeRankLingyu:             "领域排行",
		FuncOpenTypeRankShenfa:             "身法排行",
		FuncOpenTypeRankFeather:            "护体仙羽排行",
		FuncOpenTypeRankShield:             "盾刺排行",
		FuncOpenTypeRankAnQi:               "暗器排行",
		FuncOpenTypeRankFaBao:              "法宝排行",
		FuncOpenTypeRankXianTi:             "仙体排行",
		FuncOpenTypeAnQi:                   "暗器",
		FuncOpenTypeAnQiAdvanced:           "暗器升阶",
		FuncOpenTypeTianShu:                "天书",
		FuncOpenTypeCoupleFuBen:            "双修副本",
		FuncOpenTypeCrossWorldBoss:         "跨服世界boss",
		FuncOpenTypeBloodPool:              "血池",
		FuncOpenTypeZhuanSheng:             "转生",
		FuncOpenTypeFirstChargeCycleDay:    "每日首充",
		FuncOpenTypeVIPPluss:               "至尊会员",
		//FuncOpenTypeFashionStar:              "时装升星",
		FuncOpenTypeCrossTuLong:                 "跨服屠龙",
		FuncOpenTypeSinglePlay:                  "单人玩法",
		FuncOpenTypeContendPK:                   "竞争pk",
		FuncOpenTypeMerge:                       "合服庆典",
		FuncOpenTypeMergeActivity:               "合服乐翻天",
		FuncOpenTypeBindGoldBowl:                "绑元聚宝盆",
		FuncOpenTypeGoldBowl:                    "元宝聚宝盆",
		FuncOpenTypeLuckyDrew:                   "幸运大转盘",
		FuncOpenTypeZhaKuang:                    "炸矿",
		FuncOpenTypeAdvancedBless:               "升阶狂欢日",
		FuncOpenTypeMarryPre:                    "结婚预告",
		FuncOpenTypeCrossLianYu:                 "无间炼狱",
		FuncOpenTypeMassacre:                    "戮仙刃",
		FuncOpenTypeGoldEquiBag:                 "转生大礼包",
		FuncOpenTypeOpenActivityCombine:         "开服乐翻天",
		FuncOpenTypeGodSiege:                    "神兽攻城",
		FuncOpenTypeAnQiJiGuan:                  "暗器机关",
		FuncOpenTypeWingRune:                    "战翼符文",
		FuncOpenTypeMountSkill:                  "坐骑技能",
		FuncOpenTypeAnQiSkill:                   "暗器技能",
		FuncOpenTypeWingSkill:                   "战翼技能",
		FuncOpenTypeYuanGodUpgrade:              "元神金装.强化",
		FuncOpenTypeYuanGodKaiGuang:             "元神金装.开光",
		FuncOpenTypeLongFengChengXiangOpen:      "龙凤呈祥开服",
		FuncOpenTypeLongFengChengXiangMerge:     "龙凤呈祥合服",
		FuncOpenTypeLuckyTurnplate:              "幸运转盘",
		FuncOpenTypeHuHuShengFeng:               "虎虎生风",
		FuncOpenTypeHeiHuoShop:                  "黑货商店",
		FuncOpenTypeFeiHuangTengDa:              "飞黄腾达",
		FuncOpenTypeSingleCharge:                "单笔充值",
		FuncOpenTypeFaBao:                       "法宝",
		FuncOpenTypeFaBaoAdvanced:               "法宝进阶",
		FuncOpenTypeFaBaoUnreal:                 "法宝幻化",
		FuncOpenTypeFaBaoTongLing:               "法宝通灵",
		FuncOpenTypeFaBaoSuit:                   "法宝配饰",
		FuncOpenTypeFaBaoSkill:                  "法宝技能",
		FuncOpenTypeXianTi:                      "仙体",
		FuncOpenTypeXianTiAdvanced:              "仙体进阶",
		FuncOpenTypeXianTiUnreal:                "仙体幻化",
		FuncOpenTypeXianTiSkill:                 "仙体技能",
		FuncOpenTypeXianTiLingYu:                "仙体灵玉",
		FuncOpenTypeLiveness:                    "活跃度系统",
		FuncOpenTypeXueDu:                       "血盾",
		FuncOpenTypeFuBen:                       "副本玩法",
		FuncOpenTypeAdvanced:                    "升阶",
		FuncOpenTypeMingRiKaiQi:                 "明日开启",
		FuncOpenTypeExpBox:                      "经验魔瓶",
		FuncOpenTypeCaiLiaoFuBen:                "材料副本",
		FuncOpenTypeCaiLiaoMount:                "材料副本.坐骑",
		FuncOpenTypeCaiLiaoWing:                 "材料副本.战翼",
		FuncOpenTypeCaiLiaoShenFa:               "材料副本.身法",
		FuncOpenTypeCaiLiaoFaBao:                "材料副本.法宝",
		FuncOpenTypeCaiLiaoXianTi:               "材料副本.仙体",
		FuncOpenTypeBossHuanJing:                "boss幻境",
		FuncOpenTypeOutlandBoss:                 "外域boss",
		FuncOpenTypeYiZheTeHui:                  "一折特惠",
		FuncOpenTypeSystemJinJie:                "系统进阶活",
		FuncOpenTypeSystemJinJieMount:           "系统进阶活动(坐骑)",
		FuncOpenTypeSystemJinJieWing:            "系统进阶活动(战翼)",
		FuncOpenTypeSystemJinJieAnQi:            "系统进阶活动(暗器)",
		FuncOpenTypeSystemJinJieBodyShiled:      "系统进阶活动(护体盾)",
		FuncOpenTypeSystemJinJieLingYu:          "系统进阶活动(领域)",
		FuncOpenTypeSystemJinJieShenFa:          "系统进阶活动(身法)",
		FuncOpenTypeSystemJinJieFaBao:           "系统进阶活动(法宝)",
		FuncOpenTypeQiTianLeiChong:              "七天累充",
		FuncOpenTypeXinaShiXuYuan:               "限时许愿",
		FuncOpenTypeYuanBaoSongBuTing:           "元宝送不停",
		FuncOpenTypeNiChongWoSong:               "你充我送",
		FuncOpenTypeShengJieSongJiang:           "升阶送奖励",
		FuncOpenTypeYaoHu:                       "妖狐来袭",
		FuncOpenTypeGoldLaBa:                    "元宝拉霸",
		FuncOpenTypeDailyQuest:                  "日环任务",
		FuncOpenTypeBaGuaMiJing:                 "八卦秘境",
		FuncOpenTypeChristmasParty:              "圣诞嘉年华",
		FuncOpenTypeChristmasXueRen:             "圣诞雪人",
		FuncOpenTypeChristmasTree:               "圣诞树",
		FuncOpenTypeChristmasGift:               "圣诞豪礼",
		FuncOpenTypeChristmasShop:               "圣诞商店",
		FuncOpenTypeChristmasGoldCar:            "驯鹿元宝车",
		FuncOpenTypeXianFu:                      "仙府",
		FuncOpenTypeQuiz:                        "仙尊答题",
		FuncOpenTypeTeamCopy:                    "组队副本",
		FuncOpenTypeNewYearParty:                "元旦嘉年华",
		FuncOpenTypeNewYearJinZhuSongLi:         "金猪送礼",
		FuncOpenTypeNewYearJinZhuNaoCheng:       "金猪闹城",
		FuncOpenTypeNewYearDaPaoDaQiu:           "大炮打气球",
		FuncOpenTypeNewYearKuangHuan:            "元旦狂欢",
		FuncOpenTypeNewYearZhaoCaiJinBao:        "招财进宝",
		FuncOpenTypeChallengeQuest:              "挑战任务",
		FuncOpenTypeDenseWat:                    "金银密窟",
		FuncOpenTypeDianXing:                    "点星",
		FuncOpenTypeDianXingJieFeng:             "点星解封",
		FuncOpenTypeTeamCopySilver:              "组队副本.银两",
		FuncOpenTypeTeamCopyXingChen:            "组队副本.星尘",
		FuncOpenTypeTeamCopyXueMo:               "组队副本.血魔",
		FuncOpenTypeTeamCopyZhuangShengEquip:    "组队副本.转生装备",
		FuncOpenTypeTeamCopyWeapon:              "组队副本.兵魂",
		FuncOpenTypeCaiLiaoTianMo:               "材料副本.天魔",
		FuncOpenTypeWardrobe:                    "衣橱套装",
		FuncOpenTypeFaBaoShengJi:                "法宝升级",
		FuncOpenTypeXianTiShengJi:               "仙体升级",
		FuncOpenTypeLingYuShengJi:               "领域升级",
		FuncOpenTypeShenfaShengJi:               "身法升级",
		FuncOpenTypeTianMo:                      "天魔体",
		FuncOpenTypeTianMoAdvanced:              "天魔体进阶",
		FuncOpenTypeTianMoUnreal:                "天魔体幻化",
		FuncOpenTypeTianMoUplevel:               "天魔体升级",
		FuncOpenTypeTianMoEquipment:             "天魔体装备",
		FuncOpenTypeTianMoSkill:                 "天魔体技能",
		FuncOpenTypeShiHunFan:                   "噬魂幡",
		FuncOpenTypeShiHunFanAdvanced:           "噬魂幡进阶",
		FuncOpenTypeShiHunFanUnreal:             "噬魂幡幻化",
		FuncOpenTypeShiHunFanShengJi:            "噬魂幡升级",
		FuncOpenTypeShiHunFanEquip:              "噬魂幡装备",
		FuncOpenTypeShiHunFanSkill:              "噬魂幡技能",
		FuncOpenTypeContinueCharge:              "连续充值",
		FuncOpenTypeMadeResource:                "经验炼制",
		FuncOpenTypeCycleSingleCharge:           "每日单笔",
		FuncOpenTypeHuanLeDaLie:                 "欢乐打猎",
		FuncOpenTypeZhuanZhuanLe:                "转转乐",
		FuncOpenTypeXiaoLiFeiBiao:               "小李飞镖",
		FuncOpenTypeHuanLeLieQiang:              "欢乐猎枪",
		FuncOpenTypeGuaGuaLe:                    "刮刮乐",
		FuncOpenTypeKaiFuZhaKuang:               "开服炸矿",
		FuncOpenTypeXiangYunZhuanPan:            "幸运转盘",
		FuncOpenTypeBossFirstKill:               "BOSS首杀",
		FuncOpenTypeKanJiaGift:                  "砍价大礼包",
		FuncOpenTypeChaoZhiTaoCan:               "超值套餐",
		FuncOpenTypeYangJinJi:                   "养鸡生金蛋",
		FuncOpenTypeLingQiChongCi:               "灵骑冲刺",
		FuncOpenTypeShiHunChongCi:               "噬魂冲刺",
		FuncOpenTypeLingYi:                      "灵翼冲刺",
		FuncOpenTypeChongCi:                     "天魔冲刺",
		FuncOpenTypeLingBingChongCi:             "灵兵冲刺",
		FuncOpenTypeLingBaoChongCi:              "灵宝冲刺",
		FuncOpenTypeLingTiChongCi:               "灵体冲刺",
		FuncOpenTypeLingYuChongCi:               "灵域冲刺",
		FuncOpenTypeLingShenChongCi:             "灵身冲刺",
		FuncOpenTypeLingQiJinJie:                "灵骑进阶",
		FuncOpenTypeShiHunJinJie:                "噬魂进阶",
		FuncOpenTypeLingYiJinJie:                "灵翼进阶",
		FuncOpenTypeLingMoJinJie:                "天魔进阶",
		FuncOpenTypeLingBingJinJie:              "灵兵进阶",
		FuncOpenTypeLingBaoJinJie:               "灵宝进阶",
		FuncOpenTypeLingTiJinJie:                "灵体进阶",
		FuncOpenTypeLingYuJinJie:                "灵域进阶",
		FuncOpenTypeLingShenJinJie:              "灵身进阶",
		FuncOpenTypeFashionShop:                 "时装商店",
		FuncOpenTypeZhuWeiChengZhan:             "助威城战",
		FuncOpenTypeHeFuQiangGou:                "合服抢购",
		FuncOpenTypeCrazyBox:                    "疯狂宝箱",
		FuncOpenTypeRankLevel:                   "等级排行",
		FuncOpenTypeChongZhiSongFangZi:          "充值送房子",
		FuncOpenTypeZaJinDan:                    "砸金蛋",
		FuncOpenTypeMingRenPu:                   "名人谱",
		FuncOpenTypeShouChongFanGold:            "首充返元宝",
		FuncOpenTypeShouChongFanBindGold:        "首充返绑元",
		FuncOpenTypeLingTongYiJiMenu:            "灵童一级菜单",
		FuncOpenTypeLingTong:                    "灵童",
		FuncOpenTypeLingTongUpgrade:             "灵童升级",
		FuncOpenTypeLingTongPeiYang:             "灵童培养",
		FuncOpenTypeLingTongHuanHua:             "灵童幻化",
		FuncOpenTypeLingTongWeapon:              "灵兵",
		FuncOpenTypeLingTongWeaponAdvanced:      "灵兵进阶",
		FuncOpenTypeLingTongWeaponUnreal:        "灵兵幻化",
		FuncOpenTypeLingTongWeaponUpgrade:       "灵兵升级",
		FuncOpenTypeLingTongWeaponEquip:         "灵兵装备",
		FuncOpenTypeLingTongWeaponSkill:         "灵兵技能",
		FuncOpenTypeLingTongMount:               "灵骑",
		FuncOpenTypeLingTongMountAdvanced:       "灵骑进阶",
		FuncOpenTypeLingTongMountUnreal:         "灵骑幻化",
		FuncOpenTypeLingTongMountUpgrade:        "灵骑升级",
		FuncOpenTypeLingTongMountEquip:          "灵骑装备",
		FuncOpenTypeLingTongMountSkill:          "灵骑技能",
		FuncOpenTypeLingTongMountPeiYang:        "灵骑技能",
		FuncOpenTypeLingTongWing:                "灵翼",
		FuncOpenTypeLingTongWingAdvanced:        "灵翼进阶",
		FuncOpenTypeLingTongWingUnreal:          "灵翼幻化",
		FuncOpenTypeLingTongWingUpgrade:         "灵翼升级",
		FuncOpenTypeLingTongWingEquip:           "灵翼装备",
		FuncOpenTypeLingTongWingSkill:           "灵翼技能",
		FuncOpenTypeLingTongShenFa:              "灵身",
		FuncOpenTypeLingTongShenFaAdvanced:      "灵身进阶",
		FuncOpenTypeLingTongShenFaUnreal:        "灵身幻化",
		FuncOpenTypeLingTongShenFaUpgrade:       "灵身升级",
		FuncOpenTypeLingTongShenFaEquip:         "灵身装备",
		FuncOpenTypeLingTongShenFaSkill:         "灵身技能",
		FuncOpenTypeLingTongLingYu:              "灵域",
		FuncOpenTypeLingTongLingYuAdvanced:      "灵身进阶",
		FuncOpenTypeLingTongLingYuUnreal:        "灵身幻化",
		FuncOpenTypeLingTongLingYuUpgrade:       "灵身升级",
		FuncOpenTypeLingTongLingYuEquip:         "灵身装备",
		FuncOpenTypeLingTongLingYuSkill:         "灵身技能",
		FuncOpenTypeLingTongFaBao:               "灵宝",
		FuncOpenTypeLingTongFaBaoAdvanced:       "灵宝进阶",
		FuncOpenTypeLingTongFaBaoUnreal:         "灵宝幻化",
		FuncOpenTypeLingTongFaBaoUpgrade:        "灵宝升级",
		FuncOpenTypeLingTongFaBaoEquip:          "灵宝装备",
		FuncOpenTypeLingTongFaBaoSkill:          "灵宝技能",
		FuncOpenTypeLingTngFaBaoTongLing:        "灵宝通灵",
		FuncOpenTypeLingTongXianTi:              "灵体",
		FuncOpenTypeLingTongXianTiAdvanced:      "灵宝进阶",
		FuncOpenTypeLingTongXianTiUnreal:        "灵宝幻化",
		FuncOpenTypeLingTongXianTiUpgrade:       "灵宝升级",
		FuncOpenTypeLingTongXianTiEquip:         "灵宝装备",
		FuncOpenTypeLingTongXianTiSkill:         "灵宝技能",
		FuncOpenTypeShouChongFanBei:             "首冲翻倍",
		FuncOpenTypeAdvancedRew:                 "升阶狂欢",
		FuncOpenTypeAdvancedReturn:              "升阶返还",
		FuncOpenTypeAdvancedCrit:                "暴击日",
		FuncOpenTypeAdvancedBlessDay:            "祝福日",
		FuncOpenTypeCaiLiaoLingTong:             "材料副本.灵童",
		FuncOpenTypeCaiLiaoLingBing:             "材料副本.灵兵",
		FuncOpenTypeCaiLiaoLingYu:               "材料副本.灵域",
		FuncOpenTypeTeamCopyLingTong:            "组队副本.灵童",
		FuncOpenTypeTeamCopyStrength:            "组队副本.强化",
		FuncOpenTypeTeamCopyUpstar:              "组队副本.升星",
		FuncOpenTypeShenMoWar:                   "神魔战场",
		FuncOpenTypeFeiSheng:                    "飞升",
		FuncOpenTypeKaiFuMuBiao:                 "开发目标",
		FuncOpenTypeHongBao:                     "红包",
		FuncOpenTypeTelent:                      "天赋",
		FuncOpenTypeHuaLingMount:                "坐骑化灵",
		FuncOpenTypeHuaLingWing:                 "战翼化灵",
		FuncOpenTypeHuaLingShenFa:               "身法化灵",
		FuncOpenTypeHuaLingAnQi:                 "暗器化灵",
		FuncOpenTypeHuaLingFaBao:                "法宝化灵",
		FuncOpenTypeHuaLingXianTi:               "仙体化灵",
		FuncOpenTypeHuaLingLingYu:               "领域化灵",
		FuncOpenTypeHuaLingShiHunFan:            "噬魂幡化灵",
		FuncOpenTypeHuaLingTianMoTi:             "天魔体化灵",
		FuncOpenTypeHuaLingLingBing:             "灵兵化灵",
		FuncOpenTypeHuaLingLingQi:               "灵骑化灵",
		FuncOpenTypeHuaLingLingWing:             "灵翼化灵",
		FuncOpenTypeHuaLingLingShen:             "灵身化灵",
		FuncOpenTypeHuaLingLingArea:             "灵域化灵",
		FuncOpenTypeHuaLingLingBao:              "灵宝化灵",
		FuncOpenTypeHuaLingLingTi:               "灵体化灵",
		FuncOpenTypeCangJingGe:                  "藏经阁",
		FuncOpenTypeMountUnreal:                 "坐骑幻化",
		FuncOpenTypeWingShengJi:                 "战翼升级",
		FuncOpenTypeShenFaUnreal:                "身法幻化",
		FuncOpenTypeShenFaEquip:                 "身法装备",
		FuncOpenTypeShenFaSkill:                 "身法装备",
		FuncOpenTypeLingYuUnreal:                "领域幻化",
		FuncOpenTypeLingYuEquip:                 "领域装备",
		FuncOpenTypeLingYuSkill:                 "领域技能",
		FuncOpenTypeAnQiPeiYang:                 "暗器培养",
		FuncOpenTypeAnQiShengJi:                 "暗器升级",
		FuncOpenTypeDingZhiTitle:                "定制称号",
		FuncOpenTypeJingCaiHuoDong:              "精彩活动",
		FuncOpenTypeXunHuanCycleSingle:          "每日单笔",
		FuncOpenTypeXunHuanShouChongFanYuanBao:  "首充返元宝",
		FuncOpenTypeXunHuanShouChongFanBangYuan: "首充返绑元",
		FuncOpenTypeXunHuanChongZhiPaiHang:      "充值排行",
		FuncOpenTypeXunHuanJinJieFanHuan:        "升阶返还",
		FuncOpenTypeXunHuanHuanLeChouJiang:      "欢乐抽签",
		FuncOpenTypeXunHuanMoJin:                "摸金",
		FuncOpenTypeXunHuanGuaGuaLe:             "刮刮乐",
		FuncOpenTypeXunHuanZhiZhunLunPan:        "至尊轮盘",
		FuncOpenTypeXunHuanQuanMinQiangGou:      "全民抢购",
		FuncOpenTypeXunHuanJuHuaSuan:            "聚划算",
		FuncOpenTypeXunHuanXianShiLiBao:         "限时礼包",
		FuncOpenTypeXunHuanMaiYiSongYi:          "买一送一",
		FuncOpenTypeXunHuanZhaoCaiJinBao:        "招财进宝",
		FuncOpenTypeXunHuanHuangJinKuangGong:    "黄金矿工",
		FuncOpenTypeXunHuanShenMiBaoXiang:       "神秘宝箱",
		FuncOpenTypeXunHuanYiXiaoBoDa:           "以小博大",
		FuncOpenTypeAllianceDaily:               "仙盟日常",
		FuncOpenTypeAllianceBoss:                "仙盟Boss",
		FuncOpenTypeEquipBaoKu:                  "装备宝库",
		FuncOpenTypeGemTaoZhuang:                "宝石（套装）",
		FuncOpenTypeYuanGodQiangHuaTaoZhuang:    "强化（套装）",
		FuncOpenTypeYuanGodOpenLightTaoZhuang:   "开光（套装）",
		FuncOpenTypeYuanGodUpstarTaoZhuang:      "升星（套装）",
		FuncOpenTypeAllianceAltar:               "仙盟圣坛",
		FuncOpenTypeGoldEquipExtend:             "装备·继承",
		FuncOpenTypeShenQi:                      "神器",
		FuncOpenTypeShenQiSmelt:                 "神器淬炼",
		FuncOpenTypeShenQiQiLing:                "神器器灵",
		FuncOpenTypeShenQiZhuLing:               "神器注灵",
		FuncOpenTypeShenQiTaoZhuang:             "神器套装",
		FuncOpenTypeShenQiResolve:               "神器分解",
		FuncOpenTypeShenQiXunBao:                "神器寻宝",
		FuncOpenTypeYingLingPu:                  "英灵谱",
		FuncOpenTypeZhiZhunYingLingPu:           "至尊英灵谱",
		FuncOpenTypeShengHen:                    "圣痕",
		FuncOpenTypeShengHenXunBao:              "圣痕寻宝",
		FuncOpenTypeTuLongEquip:                 "屠龙装备",
		FuncOpenTypeTuLongQiangHua:              "屠龙装备-强化",
		FuncOpenTypeTuLongRongHe:                "屠龙装备-融合",
		FuncOpenTypeTuLongZhuanHua:              "屠龙装备-转化",
		FuncOpenTypeTuLongSkill:                 "屠龙装备-技能",
		FuncOpenTypeTuLongXunBao:                "屠龙装备-寻宝",
		FuncOpenTypeMingGe:                      "命格",
		FuncOpenTypeMingGong:                    "命宫",
		FuncOpenTypeMingPan:                     "命盘",
		FuncOpenTypeMingGeHeCheng:               "命格合成",
		FuncOpenTypeMingGeXunBao:                "命格寻宝",
		FuncOpenTypeZhenFa:                      "阵法",
		FuncOpenTypeZhenQi:                      "阵旗",
		FuncOpenTypeZhenHuo:                     "阵法仙火",
		FuncOpenTypeZhenFaSuit:                  "阵法套装",
		FuncOpenTypeXunHuanXunHuanChongZhi:      "循环充值",
		FuncOpenTypeXunHuanJiFenShangCheng:      "积分商城",
		FuncOpenTypeXunHuanShengJieHaoLi:        "升阶豪礼",
		FuncOpenTypeXunHuanXingYuXinYuan:        "星语心愿",
		FuncOpenTypeXunHuanBaiBeiFanLi:          "百倍返利",
		FuncOpenTypeXunHuanChaoZhiHuanGou:       "超值换购",
		FuncOpenTypeXunHuanZaJinZhu:             "砸金猪",
		FuncOpenTypeXunHuanMingRenPu:            "名人谱",
		FuncOpenTypeXunHuanMeiLiPaiHang:         "魅力排行",
		FuncOpenTypeXunHuanBiaoBaiPaiHang:       "表白排行",
		FuncOpenTypeDongFang:                    "洞房",
		FuncOpenTypeBaoBao:                      "宝宝",
		FuncOpenTypeBaoBaoLearn:                 "四书五经",
		FuncOpenTypeBaoBaoToy:                   "玩具",
		FuncOpenTypeBaoBaoZhuanShi:              "投胎转世",
		FuncOpenTypeTrad:                        "交易市场",
		FuncOpenTypeDingQing:                    "定期信物",
		FuncOpenTypeMarryJiNian:                 "结婚纪念",
		FuncOpenTypeMarryMenu:                   "结婚",
		FuncOpenTypeRankCharm:                   "魅力排行",
		FuncOpenTypeRankBiaoBai:                 "表白排行",
		FuncOpenTypeLingTongWeaponEquipUpgrade:  "灵兵装备强化",
		FuncOpenTypeLingTongWeaponUnrealLabel:   "灵兵幻化",
		FuncOpenTypeLingTongWeaponExpert:        "灵兵达人",
		FuncOpenTypeLingTongWeaponAdvencedRew:   "灵兵升阶奖励",
		FuncOpenTypeLingTongWingEquipUpgrade:    "灵翼装备强化",
		FuncOpenTypeLingTongWingUnrealLabel:     "灵翼幻化",
		FuncOpenTypeLingTongWingExpert:          "灵翼达人",
		FuncOpenTypeLingTongWingAdvencedRew:     "灵翼升阶奖励",
		FuncOpenTypeLingTongShenFaEquipUpgrade:  "灵身装备强化",
		FuncOpenTypeLingTongShenFaUnrealLabel:   "灵身幻化",
		FuncOpenTypeLingTongShenFaExpert:        "灵身达人",
		FuncOpenTypeLingTongShenFaAdvencedRew:   "灵身升阶奖励",
		FuncOpenTypeLingTongLingYuEquipUpgrade:  "灵域装备强化",
		FuncOpenTypeLingTongLingYuUnrealLabel:   "灵域幻化",
		FuncOpenTypeLingTongLingYuExpert:        "灵域达人",
		FuncOpenTypeLingTongLingYuAdvencedRew:   "灵域升阶奖励",
		FuncOpenTypeLingTongFaBaoEquipUpgrade:   "灵宝装备强化",
		FuncOpenTypeLingTongFaBaoTongLingLabel:  "灵宝通灵",
		FuncOpenTypeLingTongFaBaoExpert:         "灵宝达人",
		FuncOpenTypeLingTongFaBaoAdvencedRew:    "灵宝升阶奖励",
		FuncOpenTypeLingTongXianTiEquipUpgrade:  "灵体装备强化",
		FuncOpenTypeLingTongXianTiUnrealLabel:   "灵体幻化",
		FuncOpenTypeLingTongXianTiExpert:        "灵体达人",
		FuncOpenTypeLingTongXianTiAdvencedRew:   "灵体升阶奖励",
		FuncOpenTypeLingTongMountEquipUpgrade:   "灵骑装备强化",
		FuncOpenTypeLingTongMountExpert:         "灵骑达人",
		FuncOpenTypeLingTongMountAdvencedRew:    "灵骑升阶奖励",
		FuncOpenTypeWingRuneUpgrade:             "战翼符文强化",
		FuncOpenTypeWingAdvencedLabel:           "战翼进阶标签页",
		FuncOpenTypeWingUnrealLabel:             "战翼幻化标签页",
		FuncOpenTypeShenFaEquipUpgrade:          "身法装备强化",
		FuncOpenTypeShenFaUnrealLabel:           "身法幻化标签页",
		FuncOpenTypeShenFaExpert:                "身法达人",
		FuncOpenTypeLingYuEquipUpgrade:          "领域装备强化",
		FuncOpenTypeLingYuUnrealLabel:           "领域幻化标签页",
		FuncOpenTypeLingYuExpert:                "领域达人",
		FuncOpenTypeAnQiEquipUpgrade:            "暗器机关强化",
		FuncOpenTypeAnQiExpert:                  "暗器达人",
		FuncOpenTypeAnQiAdvencedRew:             "暗器升阶奖励",
		FuncOpenTypeFaBaoEquipUpgrade:           "法宝配饰强化",
		FuncOpenTypeFaBaoUnrealLabel:            "法宝幻化标签页",
		FuncOpenTypeXianTiEquipUpgrade:          "仙体灵玉强化",
		FuncOpenTypeXianTiUnrealLabel:           "仙体幻化丹",
		FuncOpenTypeXianTiExpert:                "仙体达人",
		FuncOpenTypeShiHunFanEquipUpgrade:       "噬魂幡装备强化",
		FuncOpenTypeShiHunFanExpert:             "噬魂幡达人",
		FuncOpenTypeShiHunFanAdvencedRew:        "噬魂幡升阶奖励",
		FuncOpenTypeTianMoEquipUpgrade:          "天魔体装备强化",
		FuncOpenTypeTianMoExpert:                "天魔体达人",
		FuncOpenTypeTianMoAdvencedRew:           "天魔体升阶奖励",
		FuncOpenTypeWingExpert:                  "战翼达人",
		FuncOpenTypeWingAdvencedRew:             "战翼升阶奖励",
		FuncOpenTypeShenFaAdvencedRew:           "身法升阶奖励",
		FuncOpenTypeLingYuAdvencedRew:           "领域升阶奖励",
		FuncOpenTypeFaBaoExpert:                 "法宝达人",
		FuncOpenTypeFaBaoAdvencedRew:            "法宝升阶奖励",
		FuncOpenTypeXianTiAdvencedRew:           "仙体升阶奖励",
		FuncOpenTypeMountJinJie:                 "坐骑进阶标签页",
		FuncOpenTypeMountHuanHua:                "坐骑幻化标签页",
		FuncOpenTypeMountEquipQiangHua:          "坐骑装备强化",
		FuncOpenTypeMountDaRen:                  "坐骑达人",
		FuncOpenTypeMountJinJieRew:              "坐骑升阶奖励",
		FuncOpenTypeXunHuanShenMiBox:            "循环-神秘宝箱",
		FuncOpenTypeXunHuanRankCost:             "消费排行",
		FuncOpenTypeFuQiFuBen:                   "夫妻副本",
		FuncOpenTypeMarryHunJie:                 "结婚婚戒",
		FuncOpenTypeMarryAiQingShu:              "结婚爱情树",
		FuncOpenTypeCycleSingleChargeTwo:        "每日单笔",
		FuncOpenTypeSingleChargeMaxRew:          "单笔充值(最近档次)",
		FuncOpenTypeBangYuanFanLi:               "绑元返利",
		FuncOpenTypeYuanBaoFanLi:                "元宝返利",
		FuncOpenTypeBaoKuCritDay:                "宝库暴击",
		FuncOpenTypeMarryBiaoBai:                "表白",
		FuncOpenTypeAdvancedBlessDayMax:         "祝福日",
		FuncOpenTypeNiChongWoSongTwo:            "你充我送",
		FuncOpenTypeXianTaoDaHui:                "仙桃大会",
		FuncOpenTypeShenYuZhiZhan:               "神域之战",
		FuncOpenTypeLongGongTanBao:              "龙宫探宝",
		FuncOpenTypePrivateCustomer:             "专属客服",
		FuncOpenTypeFashionShopXunHuan:          "时装商店",
		FuncOpenTypeMyHouse:                     "我的房子",
		FuncOpenTypeMergeXunHuanBaoJi:           "暴击日",
		FuncOpenTypeMergeXunHuanAdvancedReturn:  "升阶返",
		FuncOpenTypeMergeXunHuanBlessDay:        "祝福日",
		FuncOpenTypeMergeHeFuQiangGou:           "合服抢购(合服循环)",
		FuncOpenTypeHouseExtended:               "房产活动",
		FuncOpenTypeXunHuanMount:                "坐骑进阶",
		FuncOpenTypeXunHuanWing:                 "战翼进阶",
		FuncOpenTypeXunHuanAnQi:                 "暗器进阶",
		FuncOpenTypeXunHuanXianTi:               "仙体进阶",
		FuncOpenTypeXunHuanLingYu:               "领域进阶",
		FuncOpenTypeXunHuanShenFa:               "身法进阶",
		FuncOpenTypeXunHuanFaBao:                "法宝进阶",
		FuncOpenTypeXunHuanLingQi:               "灵骑进阶",
		FuncOpenTypeXunHuanShiHunFan:            "噬魂进阶",
		FuncOpenTypeXunHuanLingYi:               "灵翼进阶",
		FuncOpenTypeXunHuanTianMoTi:             "天魔进阶",
		FuncOpenTypeXunHuanLingBing:             "灵兵进阶",
		FuncOpenTypeXunHuanLingBao:              "灵宝进阶",
		FuncOpenTypeXunHuanLingTi:               "灵体进阶",
		FuncOpenTypeXunHuanLingTongYu:           "灵域进阶",
		FuncOpenTypeXunHuanLingShen:             "灵身进阶",
		FuncOpenTypeShenZhuMount:                "坐骑装备神铸",
		FuncOpenTypeShenZhuWing:                 "战翼装备神铸",
		FuncOpenTypeShenZhuAnQi:                 "暗器装备神铸",
		FuncOpenTypeShenZhuLingYu:               "领域装备神铸",
		FuncOpenTypeShenZhuFaBao:                "法宝装备神铸",
		FuncOpenTypeShenZhuXianTi:               "仙体装备神铸",
		FuncOpenTypeShenZhuShenFa:               "身法装备神铸",
		FuncOpenTypeShenZhuLingTongMount:        "灵骑装备神铸",
		FuncOpenTypeShenZhuShiHunFan:            "噬魂幡装备神铸",
		FuncOpenTypeShenZhuLingTongWing:         "灵翼装备神铸",
		FuncOpenTypeShenZhuTianMoTi:             "天魔体装备神铸",
		FuncOpenTypeShenZhuLingTongWeapon:       "灵兵装备神铸",
		FuncOpenTypeShenZhuLingTongFaBao:        "灵宝装备神铸",
		FuncOpenTypeShenZhuLingTongXianTi:       "灵体装备神铸",
		FuncOpenTypeShenZhuLingTongLingYu:       "灵域装备神铸",
		FuncOpenTypeShenZhuLingTongShenFa:       "灵身装备神铸",
		FuncOpenTypeTongLingMount:               "坐骑通灵",
		FuncOpenTypeTongLingWing:                "战翼通灵",
		FuncOpenTypeTongLingAnQi:                "暗器通灵",
		FuncOpenTypeTongLingLingYu:              "领域通灵",
		FuncOpenTypeTongLingFaBao:               "法宝通灵",
		FuncOpenTypeTongLingXianTi:              "仙体通灵",
		FuncOpenTypeTongLingShenFa:              "身法通灵",
		FuncOpenTypeTongLingLingTongMount:       "灵骑通灵",
		FuncOpenTypeTongLingShiHunFan:           "噬魂幡通灵",
		FuncOpenTypeTongLingLingTongWing:        "灵翼通灵",
		FuncOpenTypeTongLingTianMoTi:            "天魔体通灵",
		FuncOpenTypeTongLingLingTongWeapon:      "灵兵通灵",
		FuncOpenTypeTongLingLingTongFaBao:       "灵宝通灵",
		FuncOpenTypeTongLingLingTongXianTi:      "灵体通灵",
		FuncOpenTypeTongLingLingTongLingYu:      "灵域通灵",
		FuncOpenTypeTongLingLingTongShenFa:      "灵身通灵",
		FuncOpenTypeJueXingXianTi:               "仙体觉醒",
		FuncOpenTypeJueXingShenFa:               "身法觉醒",
		FuncOpenTypeJueXingLingTongMount:        "灵骑觉醒",
		FuncOpenTypeJueXingShiHunFan:            "噬魂幡觉醒",
		FuncOpenTypeJueXingLingTongWing:         "灵翼觉醒",
		FuncOpenTypeJueXingTianMoTi:             "天魔体觉醒",
		FuncOpenTypeJueXingLingTongWeapon:       "灵兵觉醒",
		FuncOpenTypeJueXingLingTongFaBao:        "灵宝觉醒",
		FuncOpenTypeJueXingLingTongXianTi:       "灵体觉醒",
		FuncOpenTypeJueXingLingTongLingYu:       "灵域觉醒",
		FuncOpenTypeJueXingLingTongShenFa:       "灵身觉醒",
		FuncOpenTypeJueXingWing:                 "战翼觉醒",
		FuncOpenTypeJueXingMount:                "坐骑觉醒",
		FuncOpenTypeJueXingAnQi:                 "暗器觉醒",
		FuncOpenTypeJueXingLingYu:               "领域觉醒",
		FuncOpenTypeJueXingFaBao:                "法宝觉醒",
		FuncOpenTypeYuXiZhiZhan:                 "玉玺之战",
		FuncOpenTypeChildrenDay:                 "六一狂欢",
		FuncOpenTypeChildrenDayBinFenTangGuo:    "缤纷糖果",
		FuncOpenTypeChildrenDayHuanLeQiQiu:      "欢乐气球",
		FuncOpenTypeChildrenDayLiuYiLiBao:       "六一礼包",
		FuncOpenTypeChildrenDayChongZhiFangSong: "充值放送",
		FuncOpenTypeDuanWuZongQingDuanWu:        "粽情端午",
		FuncOpenTypeDuanWuDuanWuMingRen:         "端午名人",
		FuncOpenTypeDuanWuSaiLongZhou:           "赛龙舟",
		FuncOpenTypeDuanWuWuWeiZongHe:           "五味粽盒",
		FuncOpenTypeDuanWuChongZhiFangSong:      "充值放送",
		FuncOpenTypeDuanWuQiYuDao:               "奇遇岛",
		FuncOpenTypeZhouKa:                      "周卡",
		FuncOpenTypeFuShi:                       "八卦符石",
		FuncOpenType618KuangHuan:                "618狂欢",
		FuncOpenTypeJueZhan618:                  "决战618",
		FuncOpenTypeNianZhongDaZu:               "年中大促",
		FuncOpenTypeQiangHongBao:                "抢红包",
		FuncOpenTypeNianZhongChongZhiFangSong:   "充值放送",
		FuncOpenType618MuBiao:                   "618目标",
		FuncOpenTypeTitleUpstar:                 "称号升星",
		FuncOpenTypeLingTongUpstar:              "灵童升星",
		FuncOpenTypeLingTongEquip:               "灵童装备",
		FuncOpenTypeLingTongEquipStrengthen:     "灵童装备强化",
		FuncOpenTypeLingTongEquipAdvanced:       "灵童装备升阶",
		FuncOpenTypeShenZhuLingTongEquip:        "灵童装备灵锻",
		FuncOpenTypeLingTongEquipWuXingLingZhu:  "灵童装备五行宝珠",
		FuncOpenTypeQiXueQiang:                  "泣血枪",
		FuncOpenTypeShaLuZhiDu:                  "杀戮之都",
		FuncOpenTypeXiaLingYing:                 "夏令营",
		FuncOpenTypeQingLiangYiXia:              "清凉一夏",
		FuncOpenTypeHaiLuoZhiSheng:              "海螺之声",
		FuncOpenTypeHaiDaoTanXian:               "海岛探险",
		FuncOpenTypeShaTanWaBao:                 "沙滩挖宝",
		FuncOpenTypeChongZhiFangSong:            "充值放送",
		FuncOpenTypeChuangShiPreview:            "创世之战预告",
		FuncOpenTypeJieYi:                       "结义",
		FuncOpenTypeXiongDiWeiMing:              "兄弟威名",
		FuncOpenTypeWeiMingUpLev:                "威名升级",
		FuncOpenTypeXiongDiToken:                "兄弟信物",
		FuncOpenTypeTokenUpLev:                  "信物强化",
		FuncOpenTypeBiWuDaHui:                   "比武大会",
		FuncOpenTypeHuanLeJingCai:               "欢乐竞猜",
		FuncOpenTypeWuLianDianFeng:              "武炼巅峰",
		FuncOpenTypeZhanLeiGu:                   "战鼓擂",
		FuncOpenTypeYiQuanZhiBa:                 "一拳制霸",
		FuncOpenTypeTianFengMuBiao:              "巅峰目标",
		FuncOpenTypeWuLianChongZhiFangSong:      "武炼巅峰之充值放送",
		FuncOpenTypeZhanLiLiBao:                 "战力礼包",
		FuncOpenTypeLeiTaiZhuLi:                 "擂台助力",
		FuncOpenTypeWushuangWeapon:              "无双神器",
		FuncOpenTypeMaterialBaoKu:               "材料宝库",
		FuncOpenTypeShengJieHaoLi:               "升阶豪礼",
		FuncOpenTypeXinYuXinYuan:                "星语心愿",
		FuncOpenTypeXuanHuanChongZhi:            "循环充值",
		FuncOpenTypeChaoJiHuanGou:               "超值换购",
		FuncOpenTypeForgeTotalEntra:             "锻造总入口",
		FuncOpenTypeBossEntra:                   "BOSS总入口",
		FuncOpenTypeShenQiRenZhu:                "神器认主",
		FuncOpenTypeCycleSingleChargeThree:      "每日单笔3",
		FuncOpenTypeGodCasting:                  "神铸",
		FuncOpenTypeForgeSoul:                   "锻魂",
		FuncOpenTypeCastingSpirit:               "铸灵",
		FuncOpenTypeGodCastingInherit:           "神铸继承",
		FuncOpenTypeZhenXiBoss:                  "珍稀BOSS",
		FuncOpneTypeShenQiRenZhuJieRi:           "神器认主节日",
		FuncOpenTypeArenaPvpActivity:            "比武助力活动",
		FuncOpenTypeChuangShiZhiZhan:            "创世之战",
		FuncOpenTypeDingShiBoss:                 "定时boss",
		FuncOpenTypeShengShouBoss:               "圣兽密境",
		FuncOpenTypeNewFirstCharge:              "首充翻倍",
		FuncOpenTypeSanJieMiBao:                 "三界秘宝",
		FuncOpenTypeXianRenZhiLu:                "仙人指路",
		FuncOpenTypeLianJinLu:                   "炼金炉",
		FuncOpenTypeXiuXianDianJi:               "修仙典籍",
		FuncOpenTypeTongTianTa:                  "通天塔",
		FuncOpenTypeYunYinXianRen:               "云隐仙人",
		FuncOpenTypeSanJieJinBang:               "三界金榜",
		FuncOpenTypeMergeInvest:                 "合服投资",
		FuncOpenTypeZhuanshengChongci:           "转生冲刺",
		FuncOpenTypeXianZunCard:                 "仙尊特权卡",
		FuncOpenTypeRing:                        "特戒",
		FuncOpenTypeRingStrengthen:              "特戒强化",
		FuncOpenTypeRingAdvance:                 "特戒进阶",
		FuncOpenTypeRingFuse:                    "特戒融合",
		FuncOpenTypeRingJingLing:                "特戒净灵",
		FuncOpenTypeRingXunBao:                  "特戒寻宝",
		FuncOpenTypeShangGuZhiLing:              "上古之灵",
		FuncOpenTypeShangGuZhiLingUpLevel:       "上古之灵升级",
		FuncOpenTypeShangGuZhiLingLingWen:       "上古之灵灵纹",
		FuncOpenTypeShangGuZhiLingUpRank:        "上古之灵进阶",
		FuncOpenTypeShangGuZhiLingLingLian:      "上古之灵灵炼",
		FuncOpenTypeShangGuZhiLingMiLing:        "上古之灵觅灵",
	}
)

func (ft FuncOpenType) String() string {
	return funcOpenTypeMap[ft]
}
