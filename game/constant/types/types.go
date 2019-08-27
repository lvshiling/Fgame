package types

type ConstantType int32

const (
	ConstantTypeBagTotalNum                      ConstantType = 1   //背包总数量
	ConstantTypeBagDefaultOpenNum                             = 2   //默认开启的背包数量
	ConstantTypeGoldForSingleSlot                             = 3   //开启单个背包所需的元宝
	ConstantTypeSlotNumForSingleOpen                          = 4   //一次开启多少格背包
	ConstantTypePVP                                           = 5   //pvp伤害系数
	ConstantTypeCrit                                          = 6   //暴击伤害系数
	ConstantTypeBlock                                         = 7   //格挡系数
	ConstantTypeHarmPercentMax                                = 8   //增伤万分比上限
	ConstantTypeCuthurtPercentMax                             = 9   //减伤万分比上限
	ConstantTypeInitalHit                                     = 10  //初始命中率
	ConstantTypeExitBattle                                    = 11  //脱离战斗时间（毫秒）
	ConstantTypeGemBagTotalNum                                = 12  //宝石背包格子数
	ConstantTypeBornQuestId                                   = 13  //出生接受任务的ID
	ConstantTypeReliveItemId                                  = 14  //复活所需物品id
	ConstantTypeReliveNoNeedItemsBeforeLevel                  = 15  //多少级之前复活无需物品
	ConstantTypeReliveTimesClearTime                          = 16  //复活次数清空时间(毫秒)
	ConstantTypeFirstReliveItemNum                            = 17  //第一次复活所需要的物品数量
	ConstantTypeItemNumAddEveryRelive                         = 18  //每次累计死亡次数增加的数量
	ConstantTypeInitMoveSpeed                                 = 19  //初始移动速度
	ConstantTypeSkillLimit                                    = 20  //技能与人物等级比例（等级*该值=人物技能上限）
	ConstantTypeChangeSceneCostSilver                         = 21  //跨地图传送需要银两
	ConstantTypeChangeSceneCostItem                           = 22  //跨地图传送需要物品
	ConstantTypeChangeSceneCostItemNum                        = 23  //跨地图传送需要物品数量
	ConstantTypeGiftId                                        = 24  //送礼物的物品id
	ConstantTypeEmailSaveDay                                  = 25  //邮件保存时效性(天)
	ConstantTypeEmailNumLimit                                 = 26  //邮件数量上限
	ConstantTypeTuMoTaskBarDefaultNum                         = 27  //屠魔任务栏默认开启个数
	ConstantTypeTuMoTaskBarOpenThirdPlayerLevel               = 28  //屠魔任务栏第三栏开启所需人物等级
	ConstantTypeTuMoTaskBarOpenFourthVipLevel                 = 29  //屠魔任务栏第四栏所需VIP等级
	ConstantTypeTuMoTaskInitialNum                            = 30  //屠魔任务初始次数
	ConstantTypeTuMoTaskVipAddBuyNum                          = 31  //屠魔任务Vip额外购买次数
	ConstantTypeTuMoTaskBuyNumCostGold                        = 32  //屠魔任务每次购买消耗元宝
	ConstantTypeTuMoTaskBuyNumVipLimit                        = 33  //屠魔任务额外购买所需VIP等级
	ConstantTypeLatestFriendLimit                             = 34  //最新联系人列表人数上限
	ConstantTypeFriendLimit                                   = 35  //好友列表人数上限
	ConstantTypeBlackLimit                                    = 36  //黑名单列表人数上限
	ConstantTypeSoulRuinsChallengeNum                         = 37  //帝陵遗迹每日默认闯关次数
	ConstantTypeSoulRuinsBuyNum                               = 38  //帝陵遗迹每日VIP可额外购买的次数
	ConstantTypeSoulRuinsBuyNumCostGold                       = 39  //购买1次次数所需要消耗的元宝
	ConstantTypeXianLingFuDonateNum                           = 40  //仙灵符每次捐献数量
	ConstantTypeXianfuExpChallengeNum                         = 41  //经验副本每日进入最大次数
	ConstantTypeXianfuSilverChallengeNum                      = 42  //银两副本每日进入最大次数
	ConstantTypeToDefineName1                                 = 43  //拾取物品的最大范围(半径)
	ConstantTypeToDefineName2                                 = 44  //每次拾取物品的最大数量
	ConstantTypeAllianceCreateMinLevel                        = 45  //创建仙盟的最低等级
	ConstantTypeAllianceCreateNeedItem                        = 46  //创建仙盟消耗的物品ID
	ConstantTypeAllianceCreateNeedItemCount                   = 47  //创建仙盟消耗的物品数量
	ConstantTypeAllianceJoinCoolingTime                       = 48  //申请加入联盟的冷却时间(毫秒)
	ConstantTypeAllianceGetItemTime                           = 49  //拾取物品的间隔时间（毫秒）
	ConstantTypeAllianceHuangGongOccupyTime                   = 50  //占领皇宫的读条时间(毫秒)
	ConstantTypeAllianceHuangGongOccupyFlagTime               = 51  //占领旗子的读条时间(毫秒)
	ConstantTypeAllianceItemHuFu                              = 52  //虎符对应的物品表ID
	ConstantTypeAlliancePropertyPerHuFu                       = 53  //一个虎符增加的属性(万分比)
	ConstantTypeAlliancelogLimit                              = 54  //仙盟日志数量的保存上限
	ConstantTypeConvertToQianKunBaDaiId                       = 55  //消耗腰牌兑换的物品id
	ConstantTypeConvertNeedNum                                = 56  //兑换一个八宝乾坤袋需要多少腰牌
	ConstantTypeConvertLimit                                  = 57  //兑换八宝乾坤袋的最大数量
	ConstantTypeSecretCardNum                                 = 58  //天机牌每日翻牌次数
	ConstantTypeSecretCardCostGold                            = 59  //天机牌一键完成消耗的元宝数量(每次)
	ConstantTypeTuMoTaskFinishCostGold                        = 60  //屠魔任务一键完成消耗元宝(每次)
	ConstantTypeSoulRuinsFinishCostGold                       = 61  //帝陵副本一键完成消耗的元宝数量（每次）
	ConstantTypeXianfuExpFinishCostGold                       = 62  //经验副本一键完成消耗的元宝数量(每次)
	ConstantTypeXianfuSilverFinishCostGold                    = 63  //银两副本一键完成消耗的元宝数量(每次)
	ConstantTypePersonalTransportationTimes                   = 64  //每日押送个人镖车的次数
	ConstantTypeAllianceTransportationTimes                   = 65  //每日押送仙盟镖车的次数
	ConstantTypeRobTransportationTimes                        = 66  //每日劫个人镖车的次数
	ConstantTypeDistressSignalCD                              = 67  //仙盟镖车穿云箭按钮冷却时间
	ConstantTypeRealTopRank                                   = 68  //天劫塔最大排名人数
	ConstantTypeBiaoCheBreakoffAttackTime                     = 69  //镖车脱离战斗的时间（毫秒）
	ConstantTypeGoldEquipStrengthenItemUseMax                 = 70  //一次强化最多可放置的材料数量
	ConstantTypeGoldEquipChongzhuItemUseNeed                  = 71  // 一次重铸需要放置的材料数量
	ConstantTypeGoldEquipChongzhuQualityNeed                  = 72  // 可进行重铸的最低品质
	ConstantTypeSoulStrengthenLevelLimit                      = 73  //每多少人物等级开放1级帝魂强化等级上限(帝魂强化等级上限=人物等级/该值,向下取整)
	ConstantTypeRealmInvitePairCdTime                         = 74  //天劫塔助战按钮邀请成功的冷却时间(毫秒)
	ConstantTypeChessRefreshTime                              = 75  //银两棋局和元宝棋局的刷新时间(小时)
	ConstantTypeChessSilverAttendTimesMax                     = 76  //银两棋局破解次数上限
	ConstantTypeChessChangedUseGold                           = 77  //换一批消耗的元宝
	ConstantTypeChessLogAddTimeMin                            = 78  //苍龙棋局日志特殊处理时间间隔下限(毫秒)
	ConstantTypeChessLogAddTimeMax                            = 79  //苍龙棋局日志特殊处理时间间隔上限(毫秒)
	ConstantTypeChessLogMaxNum                                = 80  //苍龙棋局日志数量的保存上限
	ConstantTypeDepotMaxNum                                   = 81  //仓库总数量
	ConstantTypeDepotDefaultOpenNum                           = 82  //仓库默认数量
	ConstantTypeDepotUnlockSlotNeedGold                       = 83  //开启单个仓库所需的元宝
	ConstantTypeOneArenaKunSolt                               = 85  //鲲背包总数量
	ConstantTypeEquipmentStrengthenLevelLimit                 = 86  //每多少人物等级开放1级装备强化等级上限(装备强化等级上限=人物等级/该值,向下取整)
	ConstantTypeProtectBuff                                   = 87  //保护状态的ID
	ConstantTypeKaiTianSpecialChessId                         = 88  // 	开天特殊掉落池id
	ConstantTypeKaiTianSpecialChessDropId1                    = 89  // 开天第一次掉落物品包id
	ConstantTypeKaiTianSpecialChessDropId2                    = 90  // 开天第二次掉落物品包id
	ConstantTypeKaiTianSpecialChessDropId3                    = 91  // 开天第三次掉落物品包id
	ConstantTypeYiJianSpecialChessId                          = 92  // 奕剑特殊掉落池id
	ConstantTypeYiJianSpecialChessDropId1                     = 93  // 奕剑第一次掉落物品包id
	ConstantTypeYiJianSpecialChessDropId2                     = 94  // 奕剑第二次掉落物品包id
	ConstantTypeYiJianSpecialChessDropId3                     = 95  // 奕剑第三次掉落物品包id
	ConstantTypePoYueSpecialChessId                           = 96  // 破月特殊掉落池id
	ConstantTypePoYueSpecialChessDropId1                      = 97  // 破月第一次掉落物品包id
	ConstantTypePoYueSpecialChessDropId2                      = 98  // 破月第二次掉落物品包id
	ConstantTypePoYueSpecialChessDropId3                      = 99  // 破月第三次掉落物品包id
	ConstantTypeGemStoneGive                                  = 100 //功能开启时直接给予玩家的原石数量
	ConstantTypeWingTrialTime                                 = 101 //战翼试用时间(毫秒)
	ConstantTypeOnlinDrewSilver                               = 104 //在线奖励银两抽奖掉落包id
	ConstantTypeOnlinDrewGold                                 = 105 //在线奖励元宝抽奖掉落包id
	ConstantTypeOneArenaKunLimit                              = 106 //1v1鱼的存储上限
	ConstantTypeMajorDefaultMaxNum                            = 107 //双修副本每日进入最大次数
	ConstantTypeRecoverHpUpperLimit                           = 108 //恢复生命值的百分比上限(百分比)
	ConstantTypeRecoverHpLowerLimit                           = 109 //恢复生命值的百分比下限(百分比)
	ConstantTypeSupplementrHpFixed                            = 110 //补充生命的固定值
	ConstantTypeSupplementrHpPercent                          = 111 //补充生命的百分比(百分比)
	ConstantTypeRecoverIntervalTime                           = 112 //补血间隔时间(毫秒)
	ConstantTypeBigShengMingPingId                            = 113 //大型生命瓶的物品id
	ConstantTypeSmallShengMingPingId                          = 114 //小型生命瓶的物品id
	ConstantTypeMajorCdTime                                   = 115 //双修副本战按钮邀请成功的冷却时间(毫秒)
	ConstantTypeXueChiInitBlood                               = 116 //血池功能开启后初始血量
	ConstantTypeXueChiInitBloodLine                           = 117 //第一次进入游戏默认的血量线(百分比)
	ConstantTypeResurrectionDanLimit                          = 118 //复活消耗复活丹上限
	ConstantTypeCollectDistance                               = 119 //采集距离
	ConstantTypeCollectTime                                   = 120 //通用采集物采集时间(毫秒)
	ConstantTypePkValueClearTime                              = 121 //清除1点红名值的在线时间(毫秒)
	ConstantTypeFriendInviteCdTime                            = 125 //加好友CD时间
	ConstantTypeReliveProtectBuff                             = 126 //复活保护buff
	ConstantTypeChessBindGoldAttendTimesMax                   = 127 //绑元棋局破解次数上限
	ConstantTypeChessChangedUseBindGold                       = 128 //绑元棋局换一批消耗的绑元
	ConstantTypePKProtectLevel                                = 135 //PK保护等级
	ConstantTypeChatWorldVipLevel                             = 139 //聊天vip等级(世界)
	ConstantTypeChatWorldLevel                                = 140 //聊天等级(世界)
	ConstantTypeTuMoImmediateFinishCostGold                   = 141 //屠魔任务接取后点击直接完成消耗的绑元数量(每次)
	ConstantTypeChatPrivateVipLevel                           = 142 //聊天vip等级(私聊)
	ConstantTypeChatPrivateLevel                              = 143 //聊天等级(私聊)
	ConstantTypeChatAllianceVipLevel                          = 144 //聊天vip等级(帮会)
	ConstantTypeChatAllianceLevel                             = 145 //聊天等级(帮会)
	ConstantTypeNewGift                                       = 147 //新手礼包
	ConstantTypePKProtectBuff                                 = 148 //pk保护buff
	ConstantTypeTowerInitTime                                 = 153 //打宝塔初始时间
	ConstantTypeBagDropSqMinPer                               = 154 //背包中杀气掉落下限(万分比)
	ConstantTypeBagDropSqMaxPer                               = 155 //背包中杀气掉落上限(万分比)
	ConstantTypePlayerDropSqCd                                = 156 //杀气掉落CD(毫秒)
	ConstantTypeDropSqProtectedTime                           = 157 //背包中杀气掉落保护时间(毫秒)
	ConstantTypeDropSqExistTime                               = 158 //背包中杀气掉落存活时间(毫秒)
	ConstantTypeDropSqMaxStack                                = 159 //被击杀后背包中杀气掉落最大堆数
	ConstantTypeZhiFeiFu                                      = 160 //直飞符id
	ConstantTypeMoonloveBuff                                  = 163 //月下情缘赏月buff
	ConstantTypeDaBaoBuff                                     = 164 //打宝塔buff
	ConstantTypeBagDropSqXiShu                                = 165 //被击杀后背包中杀气被系统回收比例（万分比）
	ConstantTypeDropSqMinStack                                = 166 //被击杀后背包中杀气掉落最小堆数
	ConstantTypeDropSqRate                                    = 167 //背包中杀气掉落万分比
	ConstantTypeTianJiPaiFinishVipLimit                       = 168 //天机牌任务不显示一键完成的VIP等级(小于该配置等级不显示)
	ConstantTypeFoeLimit                                      = 169 //仇人列表上限
	ConstantTypeTrackItem                                     = 170 //追踪符物品id
	ConstantTypeTowerDummyLogMinTime                          = 171 //打宝塔虚假记录时间下限（秒）
	ConstantTypeTowerDummyLogMaxTime                          = 172 //打宝塔虚假记录时间上限（秒）
	ConstantTypeHuHuRankItem                                  = 173 //运营活动-虎虎生风排行要求消耗的物品id
	ConstantTypeAllianceMemberDiedNoticeCD                    = 175 //仙盟成员死亡信息推送CD
	ConstantTypePiLaoInitNum                                  = 176 //疲劳值每天自然恢复上限
	ConstantTypePiLaoPrice                                    = 177 //疲劳值价格（元宝）
	ConstantTypeAllianceMemberResuceCD                        = 178 //仙盟成员救援请求CD
	ConstantTypePiLaoBuyNum                                   = 179 //一次性只能购买多少点疲劳值
	ConstantTypeXueDunBloodLimit                              = 180 //血炼值上限值
	ConstantTypeDailyFinishBindGold                           = 183 //单个日环任务一键完成所需绑元
	ConstantTypeDailyCommitBindGold                           = 184 //日环任务领取双倍奖励所需绑元
	ConstantTypeXianFuExpSaoDangNeedLevel                     = 181 //经验副本扫荡等级下限
	ConstantTypeXianFuSilverSaoDangNeedLevel                  = 182 //银两副本扫荡等级下限
	ConstantTypeAllianceDepotInitItemId1                      = 185 //仙盟仓库物品1
	ConstantTypeAllianceDepotInitItemId2                      = 186 //仙盟仓库物品2
	ConstantTypeAllianceDepotInitItemId3                      = 187 //仙盟仓库物品3
	ConstantTypeAllianceDepotInitItemId4                      = 188 //仙盟仓库物品4
	ConstantTypeDailyQuestNumLimit                            = 189 //日环任务数量上限
	ConstantTypeBaGuaMiJingInviteCd                           = 190 //八卦秘境助战按钮邀请成功的冷却时间(毫秒)
	ConstantTypeOutlandBossDropRecordsCount                   = 193 //外域boss掉落记录条数
	ConstantTypeOutlandBossDropRecordsAddTimeMin              = 194 //外域boss掉落虚假记录时间下限（毫秒）
	ConstantTypeOutlandBossDropRecordsAddTimeMax              = 195 //外域boss掉落虚假记录时间上限（毫秒）
	ConstantTypeOutlandBossZhuoQiLimit                        = 196 //外域boss浊气值上限
	ConstantTypeLaBaDummyLogAddTimeMin                        = 197 //元宝拉霸虚假记录时间下限（毫秒）
	ConstantTypeLaBaDummyLogAddTimeMax                        = 198 //元宝拉霸虚假记录时间上限（毫秒）
	ConstantTypeBackToZhuChengLevel                           = 199 //退出场景将回到主城复活点（等级）
	ConstantTypeBackToZhuChengMapId                           = 200 //退出场景将回到主城复活点（地图ID）
	ConstantTypeAllianceJoinBatchCD                           = 201 //仙盟一键加入CD（毫秒）
	ConstantTypeAddFriendInviteBatchCD                        = 202 //一键添加好友CD (毫秒)
	ConstantTypeAddFriendInviteBatchLimit                     = 203 //一键添加好友上限人数
	ConstantTypeDrewDummyLogAddTimeMin                        = 206 //运营活动通用虚假记录时间下限（毫秒）
	ConstantTypeDrewDummyLogAddTimeMax                        = 207 //运营活动通用虚假记录时间上限（毫秒）
	ConstantTypeFeiShengSuccessBuffId                         = 211 //飞升成功的buffid
	ConstantTypeFeiShengFaildBuffId                           = 212 //飞升失败的buffid
	ConstantTypeOpenEquipAutoStrengthLevel                    = 213 //开启装备自动强化的等级
	ConstantType3V3RushCdTime                                 = 214 //3V3催促按钮cd(毫秒)
	ConstantTypeKaiFuMuBiaoMaxOpenDay                         = 215 //开服目标持续到开服第几天的0点，配置为0则永久持续
	ConstantTypeFeiShengSanGongNPC                            = 216 //散功NPCid
	ConstantTypeKaiFuHuoDongChongZhi                          = 217 //开服活动重置
	ConstantTypeHongBaoDayCount                               = 218 //每个角色每天最多领取多少个红包
	ConstantTypeChatAwardDayCount                             = 219 //世界和仙盟频道前N句可获得银两奖励
	ConstantTypeChatAwardSilver                               = 220 //世界和仙盟频道发言奖励，每句发言给予的银两数额
	ConstantTypeHongBaoKeepTime                               = 221 //红包存活时间(毫秒)
	ConstantTypeMainQuestId                                   = 222 //主线任务ID，在完成此主线前左侧任务栏为特殊任务栏显示
	ConstantTypeMiBaoDepotSlotMax                             = 223 //装备宝库，秘宝仓库最大格数
	ConstantTypeEquipBaoKuLogMaxNum                           = 224 //装备宝库，抽奖记录条数
	ConstantTypeEquipBaoKuLogAddTimeMin                       = 225 //装备宝库，假记录时间min
	ConstantTypeEquipBaoKuLogAddTimeMax                       = 226 //装备宝库，假记录时间max
	ConstantTypeEquipBaoKuDummyDrop                           = 227 //装备宝库-假数据掉落包
	ConstantTypeAllianceDailyFinishBindGold                   = 228 //直接完成一个仙盟日常任务消耗的绑元
	ConstantTypeAllianceDailyCount                            = 229 //仙盟日常任务数量上限
	ConstantTypeMingGeSlotNum                                 = 230 //命格背包格子总数
	ConstantTypeShenQiSlotNum                                 = 231 //神器背包格子总数
	ConstantTypeTuLongEquipSlotNum                            = 232 //屠龙装备背包格子总数
	ConstantTypeQiLingSlotNum                                 = 233 //器灵背包格子总数
	ConstantTypeYingLingPuSlotNum                             = 234 //英灵谱背包格子总数
	ConstantTypeTuLongEquipRongHeNum                          = 235 //屠龙装备融合数量
	ConstantTypeTuLongEquipZhuanHuaNum                        = 236 //屠龙装备转化数量
	ConstantTypeTuLongEquipRongHeQuality                      = 237 //屠龙装备融合最低品质
	ConstantTypeVipLiBaoResetTime                             = 241 //VIP礼包重置时间
	ConstantTypeCoupleDefaultMaxNum                           = 242 //夫妻副本每日进入最大次数
	ConstantTypeCoupleFuBenCdTime                             = 243 //夫妻副本战按钮邀请成功的冷却时间(毫秒)
	ConstantTypeMarryDevelopLogMaxNum                         = 244 //表白日志数量的保存上限
	ConstantTypeAllianceDouShenLevelLimit                     = 245 //斗神领域生效等级
	ConstantTypeTradeItemPersonLimit                          = 246 //交易市场个人限制数量
	ConstantTypeTradeItemLimit                                = 247 //交易市场总限制数量
	ConstantTypeTuMoFinishVipLimit                            = 248 //屠魔任务不显示一键完成的VIP等级(小于该配置等级不显示)
	ConstantTypeMergeXunHuanKeepDay                           = 252 // 合服循环活动持续时间（天数）
	ConstantTypeEnterPvpBuff                                  = 254 //进入PVP状态后玩家获得的状态id
	ConstantTypeAllianceDailyCommitBindGold                   = 255 //仙盟日常任务领取双倍奖励所需绑元
	ConstantTypeDropOwnerBuffId                               = 256 //掉落归属buff
	ConstantTypeFeiShengLimit                                 = 257 //飞升次数限制
	ConstantTypeCodeExpireTime                                = 261 //兑换码过期时间
	ConstantTypeCodeExchangeLimit                             = 262 //每天s限制兑换额度
	ConstantTypeCreateNewHighAllianceUseGold                  = 263 //新版高级仙盟直接花费元宝创建仙盟所需的元宝数额
	ConstantTypeMaterialDepotSlotMax                          = 264 //材料宝库，材料仓库最大格数
	ConstantTypeMaterialBaoKuDummyDrop                        = 265 //材料宝库-假数据掉落包
	ConstantTypeWushuangEssence                               = 266 //无双神器吞噬道具
	ConstantTypeZhenXiReliveTimes                             = 267 //珍惜boss
	ConstantTypeShengShouItemId                               = 268 //圣兽boss物品id
	ConstantTypeShengShouItemNum                              = 269 //圣兽boss物品数量
	ConstantTypeShengShouReliveTimes                          = 270 //圣兽boss复活次数
)

const (
	ConstantTypeMin = ConstantTypeBagTotalNum
	ConstantTypeMax = ConstantTypeDropOwnerBuffId
)

var (
	constantTypeMap = map[ConstantType]string{
		ConstantTypeBagTotalNum:                      "背包总数量",
		ConstantTypeBagDefaultOpenNum:                "默认开启的背包数量",
		ConstantTypeGoldForSingleSlot:                "开启单个背包所需的元宝",
		ConstantTypeSlotNumForSingleOpen:             "一次开启多少格背包",
		ConstantTypePVP:                              "pvp伤害系数",
		ConstantTypeCrit:                             "暴击伤害系数",
		ConstantTypeBlock:                            "格挡系数",
		ConstantTypeHarmPercentMax:                   "增伤万分比上限",
		ConstantTypeCuthurtPercentMax:                "减伤万分比上限",
		ConstantTypeInitalHit:                        "初始命中率",
		ConstantTypeExitBattle:                       "脱离战斗时间（毫秒）",
		ConstantTypeGemBagTotalNum:                   "宝石背包格子数",
		ConstantTypeBornQuestId:                      "出生接受任务的ID",
		ConstantTypeReliveItemId:                     "复活所需物品id",
		ConstantTypeReliveNoNeedItemsBeforeLevel:     "多少级之前复活无需物品",
		ConstantTypeReliveTimesClearTime:             "复活次数清空时间(毫秒)",
		ConstantTypeFirstReliveItemNum:               "第一次复活所需要的物品数量",
		ConstantTypeItemNumAddEveryRelive:            "每次累计死亡次数增加的数量",
		ConstantTypeInitMoveSpeed:                    "初始移动速度",
		ConstantTypeSkillLimit:                       "技能与人物等级比例（等级*该值=人物技能上限）",
		ConstantTypeChangeSceneCostSilver:            "跨地图传送需要银两",
		ConstantTypeChangeSceneCostItem:              "跨地图传送需要物品",
		ConstantTypeChangeSceneCostItemNum:           "跨地图传送需要物品数量",
		ConstantTypeGiftId:                           "送礼物的物品id",
		ConstantTypeEmailSaveDay:                     "邮件保存时效性(天)",
		ConstantTypeEmailNumLimit:                    "邮件数量上限",
		ConstantTypeTuMoTaskBarDefaultNum:            "屠魔任务栏默认开启个数",
		ConstantTypeTuMoTaskBarOpenThirdPlayerLevel:  "屠魔任务栏第三栏开启所需人物等级",
		ConstantTypeTuMoTaskBarOpenFourthVipLevel:    "屠魔任务栏第四栏所需VIP等级",
		ConstantTypeTuMoTaskInitialNum:               "屠魔任务初始次数",
		ConstantTypeTuMoTaskVipAddBuyNum:             "屠魔任务Vip额外购买次数",
		ConstantTypeTuMoTaskBuyNumCostGold:           "屠魔任务每次购买消耗元宝",
		ConstantTypeTuMoTaskBuyNumVipLimit:           "屠魔任务额外购买所需VIP等级",
		ConstantTypeLatestFriendLimit:                "最新联系人列表人数上限",
		ConstantTypeFriendLimit:                      "好友列表人数上限",
		ConstantTypeBlackLimit:                       "黑名单列表人数上限",
		ConstantTypeSoulRuinsChallengeNum:            "帝陵遗迹每日默认闯关次数",
		ConstantTypeSoulRuinsBuyNum:                  "帝陵遗迹每日VIP可额外购买的次数",
		ConstantTypeSoulRuinsBuyNumCostGold:          "购买1次次数所需要消耗的元宝",
		ConstantTypeXianLingFuDonateNum:              "仙灵符每次捐献数量",
		ConstantTypeXianfuSilverChallengeNum:         "银两副本每日进入最大次数",
		ConstantTypeXianfuExpChallengeNum:            "经验副本每日进入最大次数",
		ConstantTypeAllianceCreateMinLevel:           "创建仙盟的最低等级",
		ConstantTypeAllianceCreateNeedItem:           "创建仙盟消耗的物品ID",
		ConstantTypeAllianceCreateNeedItemCount:      "创建仙盟消耗的物品数量",
		ConstantTypeAllianceJoinCoolingTime:          "申请加入联盟的冷却时间(毫秒)",
		ConstantTypeAllianceGetItemTime:              "拾取物品的间隔时间（毫秒）",
		ConstantTypeAllianceHuangGongOccupyTime:      "占领皇宫的读条时间(毫秒)",
		ConstantTypeAllianceHuangGongOccupyFlagTime:  "占领旗子的读条时间(毫秒)",
		ConstantTypeAllianceItemHuFu:                 "虎符对应的物品表ID",
		ConstantTypeAlliancePropertyPerHuFu:          "一个虎符增加的属性(万分比)",
		ConstantTypeAlliancelogLimit:                 "仙盟日志数量的保存上限",
		ConstantTypeConvertToQianKunBaDaiId:          "消耗腰牌兑换的物品id",
		ConstantTypeConvertNeedNum:                   "兑换一个八宝乾坤袋需要多少腰牌",
		ConstantTypeConvertLimit:                     "兑换八宝乾坤袋的最大数量",
		ConstantTypeSecretCardNum:                    "天机牌每日翻牌次数",
		ConstantTypeSecretCardCostGold:               "天机牌一键完成消耗的元宝数量(每次)",
		ConstantTypeTuMoTaskFinishCostGold:           "屠魔任务一键完成消耗元宝(每次)",
		ConstantTypeSoulRuinsFinishCostGold:          "帝陵副本一键完成消耗的元宝数量(每次)",
		ConstantTypeXianfuExpFinishCostGold:          "经验副本一键完成消耗的元宝数量(每次)",
		ConstantTypeXianfuSilverFinishCostGold:       "银两副本一键完成消耗的元宝数量(每次)",
		ConstantTypePersonalTransportationTimes:      "每日押送个人镖车的次数",
		ConstantTypeAllianceTransportationTimes:      "每日押送仙盟镖车的次数",
		ConstantTypeRobTransportationTimes:           "每日劫个人镖车的次数",
		ConstantTypeDistressSignalCD:                 "仙盟镖车穿云箭按钮冷却时间",
		ConstantTypeRealTopRank:                      "天劫塔最大排名人数",
		ConstantTypeBiaoCheBreakoffAttackTime:        "镖车脱离战斗的时间（毫秒）",
		ConstantTypeGoldEquipStrengthenItemUseMax:    "一次强化最多可放置的材料数量",
		ConstantTypeGoldEquipChongzhuItemUseNeed:     " 一次重铸需要放置的材料数量",
		ConstantTypeGoldEquipChongzhuQualityNeed:     " 可进行重铸的最低品质",
		ConstantTypeSoulStrengthenLevelLimit:         "每多少人物等级开放1级帝魂强化等级上限(帝魂强化等级上限=人物等级/该值,向下取整)",
		ConstantTypeRealmInvitePairCdTime:            "天劫塔助战按钮邀请成功的冷却时间(毫秒)",
		ConstantTypeChessRefreshTime:                 "银两棋局和元宝棋局的刷新时间(毫秒)",
		ConstantTypeChessSilverAttendTimesMax:        "银两棋局破解次数上限",
		ConstantTypeChessChangedUseGold:              "换一批消耗的元宝",
		ConstantTypeChessLogAddTimeMin:               "苍龙棋局日志特殊处理时间间隔下限(毫秒)",
		ConstantTypeChessLogAddTimeMax:               "苍龙棋局日志特殊处理时间间隔上限(毫秒)",
		ConstantTypeChessLogMaxNum:                   "苍龙棋局日志数量的保存上限",
		ConstantTypeDepotMaxNum:                      "仓库总数量",
		ConstantTypeDepotDefaultOpenNum:              "仓库默认数量",
		ConstantTypeDepotUnlockSlotNeedGold:          "开启单个仓库所需的元宝",
		ConstantTypeOneArenaKunSolt:                  "鲲背包总数量",
		ConstantTypeEquipmentStrengthenLevelLimit:    "每多少人物等级开放1级装备强化等级上限(装备强化等级上限=人物等级/该值,向下取整)",
		ConstantTypeProtectBuff:                      "保护buff",
		ConstantTypeGemStoneGive:                     "功能开启时直接给予玩家的原石数量",
		ConstantTypeWingTrialTime:                    "战翼试用时间(毫秒)",
		ConstantTypeOnlinDrewSilver:                  "在线奖励银两抽奖掉落包id",
		ConstantTypeOnlinDrewGold:                    "在线奖励元宝抽奖掉落包id",
		ConstantTypeOneArenaKunLimit:                 "1v1鱼的存储上限",
		ConstantTypeMajorDefaultMaxNum:               "双修副本每日进入最大次数",
		ConstantTypeRecoverHpUpperLimit:              "恢复生命值的百分比上限(百分比)",
		ConstantTypeRecoverHpLowerLimit:              "恢复生命值的百分比下限(百分比)",
		ConstantTypeSupplementrHpFixed:               "补充生命的固定值",
		ConstantTypeSupplementrHpPercent:             "补充生命的百分比(百分比)",
		ConstantTypeRecoverIntervalTime:              "补血间隔时间(毫秒)",
		ConstantTypeBigShengMingPingId:               "大型生命瓶的物品id",
		ConstantTypeSmallShengMingPingId:             "小型生命瓶的物品id",
		ConstantTypeMajorCdTime:                      "双修副本战按钮邀请成功的冷却时间(毫秒)",
		ConstantTypeXueChiInitBlood:                  "血池功能开启后初始血量",
		ConstantTypeXueChiInitBloodLine:              "第一次进入游戏默认的血量线(百分比)",
		ConstantTypeResurrectionDanLimit:             "复活消耗复活丹上限",
		ConstantTypeCollectDistance:                  "采集距离",
		ConstantTypeCollectTime:                      "通用采集物采集时间(毫秒)",
		ConstantTypePkValueClearTime:                 "清除1点红名值的在线时间(毫秒)",
		ConstantTypeFriendInviteCdTime:               "加好友CD时间",
		ConstantTypeReliveProtectBuff:                "复活保护buff",
		ConstantTypeChessBindGoldAttendTimesMax:      "绑元棋局破解次数上限",
		ConstantTypeChessChangedUseBindGold:          "绑元棋局换一批消耗的绑元",
		ConstantTypePKProtectLevel:                   "pk保护等级",
		ConstantTypeChatWorldVipLevel:                "聊天vip等级(世界)",
		ConstantTypeChatWorldLevel:                   "聊天等级(世界)",
		ConstantTypeTuMoImmediateFinishCostGold:      "屠魔任务接取后点击直接完成消耗的绑元数量(每次)",
		ConstantTypeChatPrivateVipLevel:              "聊天vip等级(私聊)",
		ConstantTypeChatPrivateLevel:                 "聊天等级(私聊)",
		ConstantTypeChatAllianceVipLevel:             "聊天vip等级(帮会)",
		ConstantTypeChatAllianceLevel:                "聊天等级(帮会)",
		ConstantTypeNewGift:                          "新手礼包id",
		ConstantTypePKProtectBuff:                    "pk保护Buff",
		ConstantTypeTowerInitTime:                    "打宝塔初始时间",
		ConstantTypeBagDropSqMinPer:                  "背包中杀气掉落下限(万分比)",
		ConstantTypeBagDropSqMaxPer:                  "背包中杀气掉落上限(万分比)",
		ConstantTypePlayerDropSqCd:                   "杀气掉落CD(毫秒)",
		ConstantTypeDropSqProtectedTime:              "背包中杀气掉落保护时间(毫秒)",
		ConstantTypeDropSqExistTime:                  "背包中杀气掉落存活时间(毫秒)",
		ConstantTypeDropSqMaxStack:                   "被击杀后背包中杀气掉落最大堆数",
		ConstantTypeMoonloveBuff:                     "月下情缘赏月buff",
		ConstantTypeDaBaoBuff:                        "打宝塔buff",
		ConstantTypeBagDropSqXiShu:                   "被击杀后背包中杀气被系统回收比例（万分比）",
		ConstantTypeDropSqMinStack:                   "被击杀后背包中杀气掉落最小堆数",
		ConstantTypeDropSqRate:                       "背包中杀气掉落万分比",
		ConstantTypeTowerDummyLogMinTime:             "打宝塔虚假记录时间下限（秒）",
		ConstantTypeTowerDummyLogMaxTime:             "打宝塔虚假记录时间上限（秒）",
		ConstantTypeHuHuRankItem:                     "运营活动-虎虎生风排行要求消耗的物品id",
		ConstantTypeFoeLimit:                         "仇人列表上限",
		ConstantTypeTrackItem:                        "追踪符物品id",
		ConstantTypeAllianceMemberDiedNoticeCD:       "仙盟成员死亡信息推送CD",
		ConstantTypeAllianceMemberResuceCD:           "仙盟成员救援请求CD",
		ConstantTypeZhiFeiFu:                         "打宝塔直飞符id",
		ConstantTypePiLaoInitNum:                     "疲劳值每天自然恢复上限",
		ConstantTypePiLaoPrice:                       "每点疲劳值价格（元宝）",
		ConstantTypePiLaoBuyNum:                      "一次性只能购买多少点疲劳值",
		ConstantTypeXueDunBloodLimit:                 "血炼值上限值",
		ConstantTypeDailyFinishBindGold:              "单个日环任务一键完成所需绑元",
		ConstantTypeDailyCommitBindGold:              "日环任务领取双倍奖励所需绑元",
		ConstantTypeXianFuExpSaoDangNeedLevel:        "经验副本扫荡等级下限",
		ConstantTypeXianFuSilverSaoDangNeedLevel:     "银两副本扫荡等级下限",
		ConstantTypeAllianceDepotInitItemId1:         "仙盟仓库物品1",
		ConstantTypeAllianceDepotInitItemId2:         "仙盟仓库物品2",
		ConstantTypeAllianceDepotInitItemId3:         "仙盟仓库物品3",
		ConstantTypeAllianceDepotInitItemId4:         "仙盟仓库物品4",
		ConstantTypeDailyQuestNumLimit:               "日环任务数量上限",
		ConstantTypeBaGuaMiJingInviteCd:              "八卦秘境助战按钮邀请成功的冷却时间(毫秒)",
		ConstantTypeOutlandBossDropRecordsCount:      "外域boss掉落记录条数",
		ConstantTypeOutlandBossDropRecordsAddTimeMin: "外域boss掉落虚假记录时间下限(毫秒)",
		ConstantTypeOutlandBossDropRecordsAddTimeMax: "外域boss掉落虚假记录时间上限(毫秒)",
		ConstantTypeOutlandBossZhuoQiLimit:           "外域boss浊气值上限",
		ConstantTypeLaBaDummyLogAddTimeMin:           "元宝拉霸虚假记录时间下限（毫秒）",
		ConstantTypeLaBaDummyLogAddTimeMax:           "元宝拉霸虚假记录时间上限（毫秒）",
		ConstantTypeBackToZhuChengLevel:              "退出场景将回到主城复活点（等级）",
		ConstantTypeBackToZhuChengMapId:              "退出场景将回到主城复活点（地图ID）",
		ConstantTypeAllianceJoinBatchCD:              "仙盟一键加入CD（毫秒）",
		ConstantTypeAddFriendInviteBatchCD:           "一键添加好友CD(毫秒)",
		ConstantTypeAddFriendInviteBatchLimit:        "一键添加好友上限人数",
		ConstantTypeDrewDummyLogAddTimeMin:           "元宝拉霸虚假记录时间下限（毫秒）",
		ConstantTypeDrewDummyLogAddTimeMax:           "元宝拉霸虚假记录时间上限（毫秒）",
		ConstantTypeFeiShengSuccessBuffId:            "飞升成功的buffId",
		ConstantTypeFeiShengFaildBuffId:              "飞升失败的buffId",
		ConstantTypeOpenEquipAutoStrengthLevel:       "开启装备自动强化的等级",
		ConstantType3V3RushCdTime:                    "3V3催促按钮cd(毫秒)",
		ConstantTypeKaiFuMuBiaoMaxOpenDay:            "开服目标持续到开服第几天的0点，配置为0则永久持续",
		ConstantTypeFeiShengSanGongNPC:               "散功NPCid",
		ConstantTypeKaiFuHuoDongChongZhi:             "开服活动重置",
		ConstantTypeHongBaoDayCount:                  "每个角色每天最多领取多少个红包",
		ConstantTypeChatAwardDayCount:                "世界和仙盟频道前N句可获得银两奖励",
		ConstantTypeChatAwardSilver:                  "世界和仙盟频道发言奖励，每句发言给予的银两数额",
		ConstantTypeHongBaoKeepTime:                  "红包存活时间(毫秒)",
		ConstantTypeMainQuestId:                      "主线任务ID,在完成此主线前左侧任务栏为特殊任务栏显示",
		ConstantTypeMiBaoDepotSlotMax:                "装备宝库，秘宝仓库最大格数",
		ConstantTypeEquipBaoKuLogMaxNum:              "装备宝库，抽奖记录条数",
		ConstantTypeEquipBaoKuLogAddTimeMin:          "装备宝库，假记录时间min(毫秒)",
		ConstantTypeEquipBaoKuLogAddTimeMax:          "装备宝库，假记录时间max(毫秒)",
		ConstantTypeEquipBaoKuDummyDrop:              "装备宝库-假数据掉落包",
		ConstantTypeAllianceDailyFinishBindGold:      "直接完成一个仙盟日常任务消耗的绑元",
		ConstantTypeAllianceDailyCount:               "仙盟日常任务数量上限",
		ConstantTypeMingGeSlotNum:                    "命格背包格子总数",
		ConstantTypeTuLongEquipSlotNum:               "屠龙装备背包格子总数",
		ConstantTypeTuLongEquipRongHeNum:             "屠龙装备融合数量",
		ConstantTypeTuLongEquipZhuanHuaNum:           "屠龙装备转化数量",
		ConstantTypeShenQiSlotNum:                    "神器背包格子总数",
		ConstantTypeQiLingSlotNum:                    "器灵背包格子总数",
		ConstantTypeYingLingPuSlotNum:                "英灵谱背包格子总数",
		ConstantTypeTuLongEquipRongHeQuality:         "屠龙装备融合最低品质",
		ConstantTypeVipLiBaoResetTime:                "VIP礼包重置时间",
		ConstantTypeCoupleDefaultMaxNum:              "夫妻副本每日进入最大次数",
		ConstantTypeCoupleFuBenCdTime:                "夫妻副本战按钮邀请成功的冷却时间(毫秒)",
		ConstantTypeMarryDevelopLogMaxNum:            "表白日志数量的保存上限",
		ConstantTypeAllianceDouShenLevelLimit:        "斗神领域生效等级",
		ConstantTypeTradeItemPersonLimit:             "交易市场个人限制数量",
		ConstantTypeTradeItemLimit:                   "交易市场总限制数量",
		ConstantTypeTuMoFinishVipLimit:               "屠魔任务不显示一键完成的VIP等级(小于该配置等级不显示)",
		ConstantTypeTianJiPaiFinishVipLimit:          "天机牌任务不显示一键完成的VIP等级(小于该配置等级不显示)",
		ConstantTypeMergeXunHuanKeepDay:              "合服循环活动持续时间（天数）",
		ConstantTypeEnterPvpBuff:                     "进入PVP状态后玩家获得的状态id",
		ConstantTypeAllianceDailyCommitBindGold:      "仙盟日常任务领取双倍奖励所需绑元",
		ConstantTypeDropOwnerBuffId:                  "掉落归属buff",
		ConstantTypeFeiShengLimit:                    "飞升次数限制",
		ConstantTypeCodeExpireTime:                   "兑换码过期时间",
		ConstantTypeCodeExchangeLimit:                "每天限制兑换额度",
		ConstantTypeCreateNewHighAllianceUseGold:     "新版高级仙盟直接花费元宝创建仙盟所需的元宝数额",
		ConstantTypeWushuangEssence:                  "无双神器吞噬道具",
		ConstantTypeZhenXiReliveTimes:                "珍稀boss复活次数",
		ConstantTypeShengShouItemId:                  "圣兽boss物品id",
		ConstantTypeShengShouItemNum:                 "圣兽boss物品数量",
		ConstantTypeShengShouReliveTimes:             "圣兽boss复活次数",
	}
)

func (ct ConstantType) String() string {
	return constantTypeMap[ct]
}
