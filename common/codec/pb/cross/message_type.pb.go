// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message_type.proto

package cross

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MessageType int32

const (
	MessageType_SI_LOGIN_TYPE                                 MessageType = 30001
	MessageType_IS_LOGIN_TYPE                                 MessageType = 30002
	MessageType_SI_PLAYER_DATA_TYPE                           MessageType = 30003
	MessageType_IS_PLAYER_GET_DROP_ITEM_TYPE                  MessageType = 30004
	MessageType_SI_PLAYER_GET_DROP_ITEM_TYPE                  MessageType = 30005
	MessageType_SI_PLAYER_SYSTEM_BATTLE_PROPERTY_CHANGED_TYPE MessageType = 30006
	MessageType_SI_PLAYER_BASIC_PROPERTY_CHANGED_TYPE         MessageType = 30007
	MessageType_SI_PLAYER_SHOW_DATA_CHANGED_TYPE              MessageType = 30008
	MessageType_IS_PLAYER_EXIT_CROSS_TYPE                     MessageType = 30009
	MessageType_SI_PLAYER_EXIT_CROSS_TYPE                     MessageType = 30010
	MessageType_IS_XUE_CHI_SYNC_TYPE                          MessageType = 30011
	MessageType_SI_XUE_CHI_SYNC_TYPE                          MessageType = 30012
	MessageType_SI_XUE_CHI_ADD_TYPE                           MessageType = 30013
	MessageType_IS_PLAYER_RELIVE_SYNC_TYPE                    MessageType = 30014
	MessageType_SI_PLAYER_RELIVE_SYNC_TYPE                    MessageType = 30015
	MessageType_IS_PLAYER_RELIVE_TYPE                         MessageType = 30016
	MessageType_SI_PLAYER_RELIVE_TYPE                         MessageType = 30017
	MessageType_SI_HEARTBEAT_TYPE                             MessageType = 30018
	MessageType_IS_HEARTBEAT_TYPE                             MessageType = 30019
	MessageType_SI_PLAYER_BATTLE_DATA_CHANGED_TYPE            MessageType = 30020
	MessageType_SI_PLAYER_TEAM_CHANGED_TYPE                   MessageType = 30021
	MessageType_SI_PLAYER_ALLIANCE_CHANGED_TYPE               MessageType = 30022
	MessageType_IS_PLAYER_KILL_BIOLOGY_TYPE                   MessageType = 30023
	MessageType_SI_PLYAER_KILL_BIOLOGY_TYPE                   MessageType = 30024
	MessageType_IS_PLYAER_MOUNT_SYNC_TYPE                     MessageType = 30025
	MessageType_SI_LING_TONG_DATA_INIT_TYPE                   MessageType = 30026
	MessageType_SI_LING_TONG_DATA_CHANGED_TYPE                MessageType = 30027
	MessageType_SI_LING_TONG_DATA_REMOVE_TYPE                 MessageType = 30028
	MessageType_IS_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE       MessageType = 30029
	MessageType_SI_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE       MessageType = 30030
	MessageType_SI_BUFF_ADD_TYPE                              MessageType = 30031
	MessageType_SI_BUFF_REMOVE_TYPE                           MessageType = 30032
	MessageType_IS_PLAYER_ACTIVITY_RANK_DATA_CHANGED_TYPE     MessageType = 30033
	MessageType_SI_PLAYER_ACTIVITY_RANK_DATA_CHANGED_TYPE     MessageType = 30034
	MessageType_SI_BUFF_UPDATE_TYPE                           MessageType = 30035
	MessageType_SI_PLAYER_JIEYI_CHANGED_TYPE                  MessageType = 30036
	MessageType_IS_PLAYER_BOSS_RELIVE_SYNC_TYPE               MessageType = 30037
	MessageType_IS_PLAYER_ACTIVITY_TICKREW_DATA_CHANGED_TYPE  MessageType = 30038
	MessageType_SI_PLAYER_TESHU_SKILL_RESET_TYPE              MessageType = 30039
	MessageType_SI_ARENA_MATCH_TYPE                           MessageType = 33000
	MessageType_IS_ARENA_MATCH_TYPE                           MessageType = 33001
	MessageType_IS_ARENA_MATCH_RESULT_TYPE                    MessageType = 33002
	MessageType_SI_ARENA_STOP_MATCH_TYPE                      MessageType = 33003
	MessageType_IS_ARENA_STOP_MATCH_TYPE                      MessageType = 33004
	MessageType_IS_ARENA_WIN_TYPE                             MessageType = 33005
	MessageType_SI_ARENA_WIN_TYPE                             MessageType = 33006
	MessageType_SI_PLAYER_ARENA_DATA_CHANGED_TYPE             MessageType = 33007
	MessageType_IS_ARENA_RELIVE_TYPE                          MessageType = 33008
	MessageType_SI_ARENA_RELIVE_TYPE                          MessageType = 33009
	MessageType_IS_ARENA_COLLECT_EXP_TREE_TYPE                MessageType = 33010
	MessageType_SI_ARENA_COLLECT_EXP_TREE_TYPE                MessageType = 33011
	MessageType_IS_ARENA_COLLECT_BOX_TYPE                     MessageType = 33023
	MessageType_SI_ARENA_COLLECT_BOX_TYPE                     MessageType = 33024
	MessageType_IS_ARENA_GIVE_UP_TYPE                         MessageType = 33025
	MessageType_SI_ARENA_GIVE_UP_TYPE                         MessageType = 33026
	MessageType_IS_ARENA_RESET_RELIVE_TIMES_TYPE              MessageType = 33027
	MessageType_IS_TULONG_KILL_BOSS_TYPE                      MessageType = 33101
	MessageType_SI_TULONG_KILL_BOSS_TYPE                      MessageType = 33102
	MessageType_IS_TULONG_ATTEND_TYPE                         MessageType = 33103
	MessageType_SI_TULONG_ATTEND_TYPE                         MessageType = 33104
	MessageType_IS_COLLECT_FINISH_TYPE                        MessageType = 33201
	MessageType_SI_COLLECT_FINISH_TYPE                        MessageType = 33202
	MessageType_IS_COLLECT_MIZANG_FINISH_TYPE                 MessageType = 33203
	MessageType_SI_COLLECT_MIZANG_FINISH_TYPE                 MessageType = 33204
	MessageType_SI_LIANYU_ATTEND_TYPE                         MessageType = 33301
	MessageType_IS_LIANYU_ATTEND_TYPE                         MessageType = 33302
	MessageType_IS_LIANYU_LINEUP_SUCCESS_TYPE                 MessageType = 33303
	MessageType_SI_LIANYU_LINEUP_SUCCESS_TYPE                 MessageType = 33304
	MessageType_SI_LIANYU_CANCLE_LINEUP_TYPE                  MessageType = 33305
	MessageType_IS_LIANYU_CANCLE_LINEUP_TYPE                  MessageType = 33306
	MessageType_IS_LIANYU_FINISH_LINEUP_CANCLE_TYPE           MessageType = 33307
	MessageType_SI_LIANYU_FINISH_LINEUP_CANCLE_TYPE           MessageType = 33308
	MessageType_IS_MASSACRE_DROP_TYPE                         MessageType = 33401
	MessageType_SI_MASSACRE_DROP_TYPE                         MessageType = 33402
	MessageType_SI_GODSIEGE_ATTEND_TYPE                       MessageType = 35301
	MessageType_IS_GODSIEGE_ATTEND_TYPE                       MessageType = 35302
	MessageType_IS_GODSIEGE_LINEUP_SUCCESS_TYPE               MessageType = 35303
	MessageType_SI_GODSIEGE_LINEUP_SUCCESS_TYPE               MessageType = 35304
	MessageType_SI_GODSIEGE_CANCLE_LINEUP_TYPE                MessageType = 35305
	MessageType_IS_GODSIEGE_CANCLE_LINEUP_TYPE                MessageType = 35306
	MessageType_IS_GODSIEGE_FINISH_LINEUP_CANCLE_TYPE         MessageType = 35307
	MessageType_SI_GODSIEGE_FINISH_LINEUP_CANCLE_TYPE         MessageType = 35308
	MessageType_SI_TEAMCOPY_START_BATTLE_TYPE                 MessageType = 33601
	MessageType_IS_TEAMCOPY_START_BATTLE_TYPE                 MessageType = 33602
	MessageType_IS_TEAMCOPY_BATTLE_RESULT_TYPE                MessageType = 33603
	MessageType_SI_TEAMCOPY_BATTLE_RESULT_TYPE                MessageType = 33604
	MessageType_IS_DENSEWAT_SYNC_TYPE                         MessageType = 33701
	MessageType_SI_DENSEWAT_SYNC_TYPE                         MessageType = 33702
	MessageType_SI_SHENMO_ATTEND_TYPE                         MessageType = 33801
	MessageType_IS_SHENMO_ATTEND_TYPE                         MessageType = 33802
	MessageType_IS_SHENMO_LINEUP_SUCCESS_TYPE                 MessageType = 33803
	MessageType_SI_SHENMO_LINEUP_SUCCESS_TYPE                 MessageType = 33804
	MessageType_SI_SHENMO_CANCLE_LINEUP_TYPE                  MessageType = 33805
	MessageType_IS_SHENMO_CANCLE_LINEUP_TYPE                  MessageType = 33806
	MessageType_IS_SHENMO_FINISH_LINEUP_CANCLE_TYPE           MessageType = 33807
	MessageType_SI_SHENMO_FINISH_LINEUP_CANCLE_TYPE           MessageType = 33808
	MessageType_IS_PLAYER_GONGXUN_ADD_TYPE                    MessageType = 33809
	MessageType_SI_PLAEYR_GONGXUN_ADD_TYPE                    MessageType = 33810
	MessageType_IS_PLAYER_GONGXUN_SUB_TYPE                    MessageType = 33811
	MessageType_SI_PLAEYR_GONGXUN_SUB_TYPE                    MessageType = 33812
	MessageType_SI_PLAYER_GONGXUN_CHANGED_TYPE                MessageType = 33813
	MessageType_IS_PLAYER_GONGXUN_CHANGED_TYPE                MessageType = 33814
	MessageType_IS_SHENMO_KILLNUM_CHANGED_TYPE                MessageType = 33815
	MessageType_SI_SHENMO_KILLNUM_CHANGED_TYPE                MessageType = 33816
	MessageType_IS_QIXUE_DROP_TYPE                            MessageType = 33901
	MessageType_SI_QIXUE_DROP_TYPE                            MessageType = 33902
	MessageType_IS_ARENAPVP_ATTEND_TYPE                       MessageType = 34001
	MessageType_SI_ARENAPVP_ATTEND_TYPE                       MessageType = 34002
	MessageType_SI_ARENAPVP_PLAYER_DATA_CHANGED_TYPE          MessageType = 34003
	MessageType_IS_ARENAPVP_RELIVE_TYPE                       MessageType = 34004
	MessageType_SI_ARENAPVP_RELIVE_TYPE                       MessageType = 34005
	MessageType_IS_ARENAPVP_RESET_RELIVETIMES_TYPE            MessageType = 34006
	MessageType_IS_ARENAPVP_ATTEND_SUCCESS_TYPE               MessageType = 34007
	MessageType_IS_ARENAPVP_RESULT_ELECTION_TYPE              MessageType = 34008
	MessageType_IS_ARENAPVP_RESULT_BATTLE_TYPE                MessageType = 34009
	MessageType_IS_ARENAPVP_ELECTION_LUCKY_REW_TYPE           MessageType = 34010
	MessageType_IS_LINEUP_ATTEND_TYPE                         MessageType = 34101
	MessageType_SI_LINEUP_ATTEND_TYPE                         MessageType = 34102
	MessageType_IS_LINEUP_CANCEL_TYPE                         MessageType = 34103
	MessageType_SI_LINEUP_CANCEL_TYPE                         MessageType = 34104
	MessageType_IS_LINEUP_SCENE_FINISH_TO_CANCEL_TYPE         MessageType = 34105
	MessageType_IS_LINEUP_SUCCESS_TYPE                        MessageType = 34106
	MessageType_SI_LINEUP_SUCCESS_TYPE                        MessageType = 34107
	MessageType_IS_CHUANGSHI_ENTER_CITY_TYPE                  MessageType = 34201
	MessageType_SI_CHUANGSHI_ENTER_CITY_TYPE                  MessageType = 34202
	MessageType_IS_CHUANGSHI_KILL_PLAYER_TYPE                 MessageType = 34203
	MessageType_IS_CHUANGSHI_SCENE_FINISH_TYPE                MessageType = 34204
	MessageType_SI_CHUANGSHI_SCENE_FINISH_TYPE                MessageType = 34205
	MessageType_IS_SHENGWEI_DROP_TYPE                         MessageType = 34301
	MessageType_SI_SHENGWEI_DROP_TYPE                         MessageType = 34302
)

var MessageType_name = map[int32]string{
	30001: "SI_LOGIN_TYPE",
	30002: "IS_LOGIN_TYPE",
	30003: "SI_PLAYER_DATA_TYPE",
	30004: "IS_PLAYER_GET_DROP_ITEM_TYPE",
	30005: "SI_PLAYER_GET_DROP_ITEM_TYPE",
	30006: "SI_PLAYER_SYSTEM_BATTLE_PROPERTY_CHANGED_TYPE",
	30007: "SI_PLAYER_BASIC_PROPERTY_CHANGED_TYPE",
	30008: "SI_PLAYER_SHOW_DATA_CHANGED_TYPE",
	30009: "IS_PLAYER_EXIT_CROSS_TYPE",
	30010: "SI_PLAYER_EXIT_CROSS_TYPE",
	30011: "IS_XUE_CHI_SYNC_TYPE",
	30012: "SI_XUE_CHI_SYNC_TYPE",
	30013: "SI_XUE_CHI_ADD_TYPE",
	30014: "IS_PLAYER_RELIVE_SYNC_TYPE",
	30015: "SI_PLAYER_RELIVE_SYNC_TYPE",
	30016: "IS_PLAYER_RELIVE_TYPE",
	30017: "SI_PLAYER_RELIVE_TYPE",
	30018: "SI_HEARTBEAT_TYPE",
	30019: "IS_HEARTBEAT_TYPE",
	30020: "SI_PLAYER_BATTLE_DATA_CHANGED_TYPE",
	30021: "SI_PLAYER_TEAM_CHANGED_TYPE",
	30022: "SI_PLAYER_ALLIANCE_CHANGED_TYPE",
	30023: "IS_PLAYER_KILL_BIOLOGY_TYPE",
	30024: "SI_PLYAER_KILL_BIOLOGY_TYPE",
	30025: "IS_PLYAER_MOUNT_SYNC_TYPE",
	30026: "SI_LING_TONG_DATA_INIT_TYPE",
	30027: "SI_LING_TONG_DATA_CHANGED_TYPE",
	30028: "SI_LING_TONG_DATA_REMOVE_TYPE",
	30029: "IS_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE",
	30030: "SI_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE",
	30031: "SI_BUFF_ADD_TYPE",
	30032: "SI_BUFF_REMOVE_TYPE",
	30033: "IS_PLAYER_ACTIVITY_RANK_DATA_CHANGED_TYPE",
	30034: "SI_PLAYER_ACTIVITY_RANK_DATA_CHANGED_TYPE",
	30035: "SI_BUFF_UPDATE_TYPE",
	30036: "SI_PLAYER_JIEYI_CHANGED_TYPE",
	30037: "IS_PLAYER_BOSS_RELIVE_SYNC_TYPE",
	30038: "IS_PLAYER_ACTIVITY_TICKREW_DATA_CHANGED_TYPE",
	30039: "SI_PLAYER_TESHU_SKILL_RESET_TYPE",
	33000: "SI_ARENA_MATCH_TYPE",
	33001: "IS_ARENA_MATCH_TYPE",
	33002: "IS_ARENA_MATCH_RESULT_TYPE",
	33003: "SI_ARENA_STOP_MATCH_TYPE",
	33004: "IS_ARENA_STOP_MATCH_TYPE",
	33005: "IS_ARENA_WIN_TYPE",
	33006: "SI_ARENA_WIN_TYPE",
	33007: "SI_PLAYER_ARENA_DATA_CHANGED_TYPE",
	33008: "IS_ARENA_RELIVE_TYPE",
	33009: "SI_ARENA_RELIVE_TYPE",
	33010: "IS_ARENA_COLLECT_EXP_TREE_TYPE",
	33011: "SI_ARENA_COLLECT_EXP_TREE_TYPE",
	33023: "IS_ARENA_COLLECT_BOX_TYPE",
	33024: "SI_ARENA_COLLECT_BOX_TYPE",
	33025: "IS_ARENA_GIVE_UP_TYPE",
	33026: "SI_ARENA_GIVE_UP_TYPE",
	33027: "IS_ARENA_RESET_RELIVE_TIMES_TYPE",
	33101: "IS_TULONG_KILL_BOSS_TYPE",
	33102: "SI_TULONG_KILL_BOSS_TYPE",
	33103: "IS_TULONG_ATTEND_TYPE",
	33104: "SI_TULONG_ATTEND_TYPE",
	33201: "IS_COLLECT_FINISH_TYPE",
	33202: "SI_COLLECT_FINISH_TYPE",
	33203: "IS_COLLECT_MIZANG_FINISH_TYPE",
	33204: "SI_COLLECT_MIZANG_FINISH_TYPE",
	33301: "SI_LIANYU_ATTEND_TYPE",
	33302: "IS_LIANYU_ATTEND_TYPE",
	33303: "IS_LIANYU_LINEUP_SUCCESS_TYPE",
	33304: "SI_LIANYU_LINEUP_SUCCESS_TYPE",
	33305: "SI_LIANYU_CANCLE_LINEUP_TYPE",
	33306: "IS_LIANYU_CANCLE_LINEUP_TYPE",
	33307: "IS_LIANYU_FINISH_LINEUP_CANCLE_TYPE",
	33308: "SI_LIANYU_FINISH_LINEUP_CANCLE_TYPE",
	33401: "IS_MASSACRE_DROP_TYPE",
	33402: "SI_MASSACRE_DROP_TYPE",
	35301: "SI_GODSIEGE_ATTEND_TYPE",
	35302: "IS_GODSIEGE_ATTEND_TYPE",
	35303: "IS_GODSIEGE_LINEUP_SUCCESS_TYPE",
	35304: "SI_GODSIEGE_LINEUP_SUCCESS_TYPE",
	35305: "SI_GODSIEGE_CANCLE_LINEUP_TYPE",
	35306: "IS_GODSIEGE_CANCLE_LINEUP_TYPE",
	35307: "IS_GODSIEGE_FINISH_LINEUP_CANCLE_TYPE",
	35308: "SI_GODSIEGE_FINISH_LINEUP_CANCLE_TYPE",
	33601: "SI_TEAMCOPY_START_BATTLE_TYPE",
	33602: "IS_TEAMCOPY_START_BATTLE_TYPE",
	33603: "IS_TEAMCOPY_BATTLE_RESULT_TYPE",
	33604: "SI_TEAMCOPY_BATTLE_RESULT_TYPE",
	33701: "IS_DENSEWAT_SYNC_TYPE",
	33702: "SI_DENSEWAT_SYNC_TYPE",
	33801: "SI_SHENMO_ATTEND_TYPE",
	33802: "IS_SHENMO_ATTEND_TYPE",
	33803: "IS_SHENMO_LINEUP_SUCCESS_TYPE",
	33804: "SI_SHENMO_LINEUP_SUCCESS_TYPE",
	33805: "SI_SHENMO_CANCLE_LINEUP_TYPE",
	33806: "IS_SHENMO_CANCLE_LINEUP_TYPE",
	33807: "IS_SHENMO_FINISH_LINEUP_CANCLE_TYPE",
	33808: "SI_SHENMO_FINISH_LINEUP_CANCLE_TYPE",
	33809: "IS_PLAYER_GONGXUN_ADD_TYPE",
	33810: "SI_PLAEYR_GONGXUN_ADD_TYPE",
	33811: "IS_PLAYER_GONGXUN_SUB_TYPE",
	33812: "SI_PLAEYR_GONGXUN_SUB_TYPE",
	33813: "SI_PLAYER_GONGXUN_CHANGED_TYPE",
	33814: "IS_PLAYER_GONGXUN_CHANGED_TYPE",
	33815: "IS_SHENMO_KILLNUM_CHANGED_TYPE",
	33816: "SI_SHENMO_KILLNUM_CHANGED_TYPE",
	33901: "IS_QIXUE_DROP_TYPE",
	33902: "SI_QIXUE_DROP_TYPE",
	34001: "IS_ARENAPVP_ATTEND_TYPE",
	34002: "SI_ARENAPVP_ATTEND_TYPE",
	34003: "SI_ARENAPVP_PLAYER_DATA_CHANGED_TYPE",
	34004: "IS_ARENAPVP_RELIVE_TYPE",
	34005: "SI_ARENAPVP_RELIVE_TYPE",
	34006: "IS_ARENAPVP_RESET_RELIVETIMES_TYPE",
	34007: "IS_ARENAPVP_ATTEND_SUCCESS_TYPE",
	34008: "IS_ARENAPVP_RESULT_ELECTION_TYPE",
	34009: "IS_ARENAPVP_RESULT_BATTLE_TYPE",
	34010: "IS_ARENAPVP_ELECTION_LUCKY_REW_TYPE",
	34101: "IS_LINEUP_ATTEND_TYPE",
	34102: "SI_LINEUP_ATTEND_TYPE",
	34103: "IS_LINEUP_CANCEL_TYPE",
	34104: "SI_LINEUP_CANCEL_TYPE",
	34105: "IS_LINEUP_SCENE_FINISH_TO_CANCEL_TYPE",
	34106: "IS_LINEUP_SUCCESS_TYPE",
	34107: "SI_LINEUP_SUCCESS_TYPE",
	34201: "IS_CHUANGSHI_ENTER_CITY_TYPE",
	34202: "SI_CHUANGSHI_ENTER_CITY_TYPE",
	34203: "IS_CHUANGSHI_KILL_PLAYER_TYPE",
	34204: "IS_CHUANGSHI_SCENE_FINISH_TYPE",
	34205: "SI_CHUANGSHI_SCENE_FINISH_TYPE",
	34301: "IS_SHENGWEI_DROP_TYPE",
	34302: "SI_SHENGWEI_DROP_TYPE",
}
var MessageType_value = map[string]int32{
	"SI_LOGIN_TYPE":                                 30001,
	"IS_LOGIN_TYPE":                                 30002,
	"SI_PLAYER_DATA_TYPE":                           30003,
	"IS_PLAYER_GET_DROP_ITEM_TYPE":                  30004,
	"SI_PLAYER_GET_DROP_ITEM_TYPE":                  30005,
	"SI_PLAYER_SYSTEM_BATTLE_PROPERTY_CHANGED_TYPE": 30006,
	"SI_PLAYER_BASIC_PROPERTY_CHANGED_TYPE":         30007,
	"SI_PLAYER_SHOW_DATA_CHANGED_TYPE":              30008,
	"IS_PLAYER_EXIT_CROSS_TYPE":                     30009,
	"SI_PLAYER_EXIT_CROSS_TYPE":                     30010,
	"IS_XUE_CHI_SYNC_TYPE":                          30011,
	"SI_XUE_CHI_SYNC_TYPE":                          30012,
	"SI_XUE_CHI_ADD_TYPE":                           30013,
	"IS_PLAYER_RELIVE_SYNC_TYPE":                    30014,
	"SI_PLAYER_RELIVE_SYNC_TYPE":                    30015,
	"IS_PLAYER_RELIVE_TYPE":                         30016,
	"SI_PLAYER_RELIVE_TYPE":                         30017,
	"SI_HEARTBEAT_TYPE":                             30018,
	"IS_HEARTBEAT_TYPE":                             30019,
	"SI_PLAYER_BATTLE_DATA_CHANGED_TYPE":            30020,
	"SI_PLAYER_TEAM_CHANGED_TYPE":                   30021,
	"SI_PLAYER_ALLIANCE_CHANGED_TYPE":               30022,
	"IS_PLAYER_KILL_BIOLOGY_TYPE":                   30023,
	"SI_PLYAER_KILL_BIOLOGY_TYPE":                   30024,
	"IS_PLYAER_MOUNT_SYNC_TYPE":                     30025,
	"SI_LING_TONG_DATA_INIT_TYPE":                   30026,
	"SI_LING_TONG_DATA_CHANGED_TYPE":                30027,
	"SI_LING_TONG_DATA_REMOVE_TYPE":                 30028,
	"IS_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE":       30029,
	"SI_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE":       30030,
	"SI_BUFF_ADD_TYPE":                              30031,
	"SI_BUFF_REMOVE_TYPE":                           30032,
	"IS_PLAYER_ACTIVITY_RANK_DATA_CHANGED_TYPE":     30033,
	"SI_PLAYER_ACTIVITY_RANK_DATA_CHANGED_TYPE":     30034,
	"SI_BUFF_UPDATE_TYPE":                           30035,
	"SI_PLAYER_JIEYI_CHANGED_TYPE":                  30036,
	"IS_PLAYER_BOSS_RELIVE_SYNC_TYPE":               30037,
	"IS_PLAYER_ACTIVITY_TICKREW_DATA_CHANGED_TYPE":  30038,
	"SI_PLAYER_TESHU_SKILL_RESET_TYPE":              30039,
	"SI_ARENA_MATCH_TYPE":                           33000,
	"IS_ARENA_MATCH_TYPE":                           33001,
	"IS_ARENA_MATCH_RESULT_TYPE":                    33002,
	"SI_ARENA_STOP_MATCH_TYPE":                      33003,
	"IS_ARENA_STOP_MATCH_TYPE":                      33004,
	"IS_ARENA_WIN_TYPE":                             33005,
	"SI_ARENA_WIN_TYPE":                             33006,
	"SI_PLAYER_ARENA_DATA_CHANGED_TYPE":             33007,
	"IS_ARENA_RELIVE_TYPE":                          33008,
	"SI_ARENA_RELIVE_TYPE":                          33009,
	"IS_ARENA_COLLECT_EXP_TREE_TYPE":                33010,
	"SI_ARENA_COLLECT_EXP_TREE_TYPE":                33011,
	"IS_ARENA_COLLECT_BOX_TYPE":                     33023,
	"SI_ARENA_COLLECT_BOX_TYPE":                     33024,
	"IS_ARENA_GIVE_UP_TYPE":                         33025,
	"SI_ARENA_GIVE_UP_TYPE":                         33026,
	"IS_ARENA_RESET_RELIVE_TIMES_TYPE":              33027,
	"IS_TULONG_KILL_BOSS_TYPE":                      33101,
	"SI_TULONG_KILL_BOSS_TYPE":                      33102,
	"IS_TULONG_ATTEND_TYPE":                         33103,
	"SI_TULONG_ATTEND_TYPE":                         33104,
	"IS_COLLECT_FINISH_TYPE":                        33201,
	"SI_COLLECT_FINISH_TYPE":                        33202,
	"IS_COLLECT_MIZANG_FINISH_TYPE":                 33203,
	"SI_COLLECT_MIZANG_FINISH_TYPE":                 33204,
	"SI_LIANYU_ATTEND_TYPE":                         33301,
	"IS_LIANYU_ATTEND_TYPE":                         33302,
	"IS_LIANYU_LINEUP_SUCCESS_TYPE":                 33303,
	"SI_LIANYU_LINEUP_SUCCESS_TYPE":                 33304,
	"SI_LIANYU_CANCLE_LINEUP_TYPE":                  33305,
	"IS_LIANYU_CANCLE_LINEUP_TYPE":                  33306,
	"IS_LIANYU_FINISH_LINEUP_CANCLE_TYPE":           33307,
	"SI_LIANYU_FINISH_LINEUP_CANCLE_TYPE":           33308,
	"IS_MASSACRE_DROP_TYPE":                         33401,
	"SI_MASSACRE_DROP_TYPE":                         33402,
	"SI_GODSIEGE_ATTEND_TYPE":                       35301,
	"IS_GODSIEGE_ATTEND_TYPE":                       35302,
	"IS_GODSIEGE_LINEUP_SUCCESS_TYPE":               35303,
	"SI_GODSIEGE_LINEUP_SUCCESS_TYPE":               35304,
	"SI_GODSIEGE_CANCLE_LINEUP_TYPE":                35305,
	"IS_GODSIEGE_CANCLE_LINEUP_TYPE":                35306,
	"IS_GODSIEGE_FINISH_LINEUP_CANCLE_TYPE":         35307,
	"SI_GODSIEGE_FINISH_LINEUP_CANCLE_TYPE":         35308,
	"SI_TEAMCOPY_START_BATTLE_TYPE":                 33601,
	"IS_TEAMCOPY_START_BATTLE_TYPE":                 33602,
	"IS_TEAMCOPY_BATTLE_RESULT_TYPE":                33603,
	"SI_TEAMCOPY_BATTLE_RESULT_TYPE":                33604,
	"IS_DENSEWAT_SYNC_TYPE":                         33701,
	"SI_DENSEWAT_SYNC_TYPE":                         33702,
	"SI_SHENMO_ATTEND_TYPE":                         33801,
	"IS_SHENMO_ATTEND_TYPE":                         33802,
	"IS_SHENMO_LINEUP_SUCCESS_TYPE":                 33803,
	"SI_SHENMO_LINEUP_SUCCESS_TYPE":                 33804,
	"SI_SHENMO_CANCLE_LINEUP_TYPE":                  33805,
	"IS_SHENMO_CANCLE_LINEUP_TYPE":                  33806,
	"IS_SHENMO_FINISH_LINEUP_CANCLE_TYPE":           33807,
	"SI_SHENMO_FINISH_LINEUP_CANCLE_TYPE":           33808,
	"IS_PLAYER_GONGXUN_ADD_TYPE":                    33809,
	"SI_PLAEYR_GONGXUN_ADD_TYPE":                    33810,
	"IS_PLAYER_GONGXUN_SUB_TYPE":                    33811,
	"SI_PLAEYR_GONGXUN_SUB_TYPE":                    33812,
	"SI_PLAYER_GONGXUN_CHANGED_TYPE":                33813,
	"IS_PLAYER_GONGXUN_CHANGED_TYPE":                33814,
	"IS_SHENMO_KILLNUM_CHANGED_TYPE":                33815,
	"SI_SHENMO_KILLNUM_CHANGED_TYPE":                33816,
	"IS_QIXUE_DROP_TYPE":                            33901,
	"SI_QIXUE_DROP_TYPE":                            33902,
	"IS_ARENAPVP_ATTEND_TYPE":                       34001,
	"SI_ARENAPVP_ATTEND_TYPE":                       34002,
	"SI_ARENAPVP_PLAYER_DATA_CHANGED_TYPE":          34003,
	"IS_ARENAPVP_RELIVE_TYPE":                       34004,
	"SI_ARENAPVP_RELIVE_TYPE":                       34005,
	"IS_ARENAPVP_RESET_RELIVETIMES_TYPE":            34006,
	"IS_ARENAPVP_ATTEND_SUCCESS_TYPE":               34007,
	"IS_ARENAPVP_RESULT_ELECTION_TYPE":              34008,
	"IS_ARENAPVP_RESULT_BATTLE_TYPE":                34009,
	"IS_ARENAPVP_ELECTION_LUCKY_REW_TYPE":           34010,
	"IS_LINEUP_ATTEND_TYPE":                         34101,
	"SI_LINEUP_ATTEND_TYPE":                         34102,
	"IS_LINEUP_CANCEL_TYPE":                         34103,
	"SI_LINEUP_CANCEL_TYPE":                         34104,
	"IS_LINEUP_SCENE_FINISH_TO_CANCEL_TYPE":         34105,
	"IS_LINEUP_SUCCESS_TYPE":                        34106,
	"SI_LINEUP_SUCCESS_TYPE":                        34107,
	"IS_CHUANGSHI_ENTER_CITY_TYPE":                  34201,
	"SI_CHUANGSHI_ENTER_CITY_TYPE":                  34202,
	"IS_CHUANGSHI_KILL_PLAYER_TYPE":                 34203,
	"IS_CHUANGSHI_SCENE_FINISH_TYPE":                34204,
	"SI_CHUANGSHI_SCENE_FINISH_TYPE":                34205,
	"IS_SHENGWEI_DROP_TYPE":                         34301,
	"SI_SHENGWEI_DROP_TYPE":                         34302,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

func init() {
	proto.RegisterEnum("cross.MessageType", MessageType_name, MessageType_value)
}

func init() { proto.RegisterFile("message_type.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 1467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x97, 0xc9, 0x8f, 0xdc, 0xc4,
	0x17, 0xc7, 0x65, 0x4b, 0xbf, 0x4b, 0xfd, 0x84, 0x14, 0x2a, 0x90, 0x7d, 0x5f, 0x08, 0x09, 0x04,
	0x24, 0xf8, 0x0b, 0xdc, 0x9e, 0x8a, 0xbb, 0x18, 0xb7, 0x6d, 0x5c, 0xe5, 0xcc, 0x34, 0x97, 0x12,
	0x42, 0x23, 0x4e, 0x28, 0x51, 0x26, 0x97, 0xdc, 0xe8, 0x99, 0xd1, 0xd0, 0x3d, 0xcd, 0xde, 0x64,
	0x21, 0xc0, 0x11, 0xee, 0x59, 0xc9, 0xc2, 0x9a, 0x84, 0x9d, 0x84, 0xec, 0x01, 0xee, 0x40, 0xe2,
	0x2c, 0x6c, 0x27, 0x90, 0x58, 0x64, 0xd7, 0xe2, 0x72, 0xbb, 0xba, 0xb9, 0xfa, 0xfb, 0xa9, 0x57,
	0xf5, 0x5e, 0xbd, 0xf7, 0xea, 0x19, 0xc0, 0x67, 0x26, 0x26, 0x27, 0x9f, 0x7c, 0x7a, 0x82, 0xed,
	0xdc, 0xb5, 0x7d, 0xe2, 0xa1, 0xed, 0x3b, 0xb6, 0xed, 0xdc, 0x06, 0xff, 0xf7, 0xd4, 0x8e, 0x6d,
	0x93, 0x93, 0x9b, 0xf6, 0xad, 0x05, 0xff, 0x6f, 0x70, 0x95, 0xee, 0xda, 0x3e, 0x01, 0xe7, 0x83,
	0xbb, 0x08, 0x66, 0x7e, 0xe8, 0xe1, 0x80, 0xd1, 0x66, 0x84, 0xe6, 0x1d, 0x48, 0xad, 0xec, 0x23,
	0x26, 0xfa, 0xc7, 0x83, 0xa9, 0x05, 0x17, 0x83, 0xf9, 0x04, 0xb3, 0xc8, 0x77, 0x9a, 0x28, 0x66,
	0x23, 0x0e, 0x75, 0xb8, 0x74, 0x28, 0xb5, 0xe0, 0x1a, 0xb0, 0x0c, 0x13, 0x29, 0x79, 0x88, 0xb2,
	0x91, 0x38, 0x8c, 0x18, 0xa6, 0xa8, 0xc1, 0x99, 0xc3, 0x9c, 0x29, 0x96, 0x1b, 0x98, 0x23, 0xa9,
	0x05, 0x1f, 0x05, 0x9b, 0x0b, 0x86, 0x34, 0x49, 0x26, 0xd6, 0x1c, 0x4a, 0x7d, 0xc4, 0xa2, 0x38,
	0x8c, 0x50, 0x4c, 0x9b, 0xcc, 0xad, 0x3b, 0x81, 0x87, 0x46, 0xf8, 0xa2, 0xa3, 0xa9, 0x05, 0x1f,
	0x00, 0xeb, 0x8b, 0x45, 0x35, 0x87, 0x60, 0x77, 0x00, 0xfc, 0x6e, 0x6a, 0xc1, 0xfb, 0xc0, 0x2a,
	0x6d, 0x87, 0x7a, 0x38, 0xc6, 0x3d, 0x29, 0x71, 0xc7, 0x52, 0x0b, 0xae, 0x04, 0x8b, 0x0b, 0x8f,
	0xd0, 0x38, 0xa6, 0xcc, 0x8d, 0x43, 0x42, 0x38, 0x70, 0x9c, 0x03, 0x85, 0xa1, 0x7e, 0xe0, 0x44,
	0x6a, 0xc1, 0x25, 0xe0, 0x1e, 0x4c, 0xd8, 0x78, 0x82, 0x98, 0x5b, 0xc7, 0x8c, 0x34, 0x03, 0x97,
	0x6b, 0x27, 0xb9, 0x46, 0xb0, 0x41, 0x7b, 0x4f, 0x85, 0x59, 0x6a, 0xce, 0x88, 0x38, 0xd4, 0xfb,
	0xa9, 0x05, 0x57, 0x81, 0x25, 0xc5, 0xa1, 0x62, 0xe4, 0xe3, 0xad, 0x48, 0x5b, 0xfc, 0x01, 0x27,
	0x8a, 0x53, 0x55, 0x88, 0x0f, 0x53, 0x0b, 0x2e, 0x05, 0xf7, 0x56, 0x6c, 0xe4, 0xe2, 0x47, 0x5c,
	0xac, 0x2c, 0xcf, 0xc5, 0x8f, 0x53, 0x0b, 0x2e, 0x04, 0x77, 0x13, 0xcc, 0xea, 0xc8, 0x89, 0x69,
	0x0d, 0x39, 0x94, 0x0b, 0xa7, 0xb8, 0x80, 0x49, 0xbf, 0x70, 0x3a, 0xb5, 0xe0, 0xfd, 0x60, 0x8d,
	0x7e, 0x33, 0xf9, 0x3d, 0x56, 0xc3, 0x7d, 0x26, 0xb5, 0xe0, 0x6a, 0xb0, 0xb4, 0x20, 0x29, 0x72,
	0x1a, 0x65, 0xe4, 0x93, 0xd4, 0x82, 0xeb, 0xc1, 0xca, 0x02, 0x71, 0x7c, 0x1f, 0x3b, 0x81, 0x8b,
	0xca, 0xd8, 0xa7, 0xdc, 0x52, 0xe1, 0xdf, 0x28, 0xf6, 0x7d, 0x56, 0xc3, 0xa1, 0x1f, 0x7a, 0x4d,
	0x8e, 0x7c, 0xa6, 0x6d, 0xd6, 0x74, 0x8c, 0xc8, 0xe7, 0xda, 0xf5, 0xe7, 0x48, 0x23, 0x4c, 0x02,
	0xaa, 0x85, 0xf1, 0x0b, 0x65, 0xc3, 0xc7, 0x81, 0xc7, 0x68, 0x18, 0x78, 0xdc, 0x2d, 0x1c, 0x60,
	0xe1, 0xfd, 0x97, 0xa9, 0x05, 0xd7, 0x81, 0x15, 0x55, 0xa4, 0x74, 0xde, 0xaf, 0x52, 0x0b, 0xae,
	0x05, 0xcb, 0xab, 0x54, 0x8c, 0x1a, 0xa1, 0x0c, 0xfd, 0xd7, 0xa9, 0x05, 0x37, 0x83, 0x0d, 0x85,
	0x53, 0x8e, 0x4b, 0xf1, 0x56, 0x4c, 0x9b, 0x2c, 0x1a, 0x35, 0xd8, 0x3c, 0xcb, 0x71, 0x2d, 0x54,
	0x43, 0xf1, 0x73, 0xa9, 0x05, 0x17, 0x80, 0x79, 0x04, 0xb3, 0x5a, 0xb2, 0x65, 0x4b, 0x91, 0x6e,
	0xdf, 0xa8, 0x4c, 0xcc, 0xbf, 0xeb, 0x07, 0x3a, 0x9f, 0x5a, 0xf0, 0x61, 0xb0, 0xd1, 0x70, 0xa0,
	0xd8, 0x09, 0x4c, 0x7b, 0x5c, 0xe0, 0x0b, 0x0c, 0x47, 0x1a, 0xb0, 0xe0, 0x62, 0x79, 0xf3, 0x24,
	0x1a, 0x71, 0xa8, 0xd8, 0xfc, 0x52, 0x7f, 0x27, 0x79, 0x0c, 0xa3, 0x26, 0x2e, 0x2f, 0xbf, 0xcc,
	0xb3, 0xa5, 0x38, 0x60, 0x2d, 0x2b, 0xcc, 0x4a, 0x35, 0x5c, 0x49, 0x2d, 0xf8, 0x08, 0x78, 0xd0,
	0xe0, 0x07, 0xc5, 0xee, 0x68, 0x8c, 0x4c, 0xad, 0xe1, 0x6a, 0x7f, 0x0b, 0xa1, 0x88, 0xd4, 0x13,
	0x46, 0xf2, 0x2c, 0x8a, 0x11, 0x41, 0xe2, 0xfe, 0xaf, 0x29, 0x0f, 0x9c, 0x18, 0x05, 0x0e, 0x6b,
	0x38, 0xd4, 0xad, 0x73, 0xe9, 0x7a, 0xcb, 0xce, 0x24, 0x4c, 0xaa, 0xd2, 0x8d, 0x96, 0x2d, 0x6a,
	0x5c, 0x97, 0x62, 0x44, 0x12, 0x5f, 0xd8, 0x4d, 0x5b, 0x36, 0x5c, 0x01, 0x16, 0x29, 0xbb, 0x84,
	0x86, 0x91, 0x6e, 0xe1, 0x26, 0xd7, 0x95, 0x85, 0x7e, 0xfd, 0x56, 0xcb, 0x16, 0xe5, 0xca, 0xf5,
	0x31, 0xd9, 0xe0, 0x6f, 0x73, 0x41, 0x19, 0x56, 0xc2, 0x9d, 0x96, 0x0d, 0x37, 0x80, 0xd5, 0xda,
	0xe5, 0xe5, 0x7a, 0x35, 0x34, 0x3f, 0xb7, 0x6c, 0xd1, 0xf3, 0x38, 0xa1, 0xb7, 0x8f, 0x5f, 0xb8,
	0xa6, 0xac, 0xeb, 0xda, 0xaf, 0x2d, 0x3b, 0x2b, 0x15, 0xb5, 0xce, 0x0d, 0x7d, 0x1f, 0xb9, 0x94,
	0xa1, 0xf1, 0x88, 0xd1, 0x18, 0x09, 0xea, 0x37, 0x4e, 0x29, 0x0b, 0x66, 0xea, 0xf7, 0x96, 0x2d,
	0x4a, 0xb7, 0x4c, 0xd5, 0xc2, 0x71, 0x0e, 0xfc, 0xc3, 0x81, 0x8a, 0x19, 0x05, 0x3c, 0x3b, 0x65,
	0x8b, 0x16, 0xc9, 0x01, 0x2f, 0x3b, 0x67, 0x12, 0x71, 0xb1, 0xc5, 0x45, 0xb5, 0xba, 0x24, 0x4e,
	0x4d, 0xd9, 0x59, 0x6a, 0x68, 0xfe, 0x67, 0xd9, 0x20, 0x3d, 0xc5, 0x0d, 0x24, 0xde, 0x86, 0xe9,
	0x29, 0x79, 0x45, 0x34, 0xf1, 0xb3, 0x8a, 0xe7, 0x1d, 0x48, 0xbd, 0x1d, 0x67, 0xa7, 0xe4, 0x15,
	0x9b, 0xf5, 0x73, 0xea, 0x84, 0x42, 0x77, 0x28, 0x45, 0x81, 0x2c, 0x5b, 0x75, 0x42, 0x83, 0x78,
	0x7e, 0xca, 0x86, 0xcb, 0xc0, 0x02, 0x4c, 0x94, 0xdb, 0x5b, 0x70, 0x80, 0x89, 0x48, 0x8d, 0x03,
	0xd3, 0xb9, 0x4a, 0xb0, 0x51, 0x3d, 0x38, 0x6d, 0x67, 0xad, 0x4a, 0x5b, 0xdb, 0xc0, 0x4f, 0x38,
	0x81, 0x57, 0x82, 0x0e, 0x71, 0x48, 0x33, 0x61, 0x80, 0x0e, 0x4f, 0xcb, 0x23, 0x66, 0x2d, 0xbc,
	0x99, 0x94, 0x8e, 0xb8, 0x7b, 0x46, 0x3a, 0x67, 0x10, 0xf7, 0xcc, 0xc8, 0x33, 0x08, 0xd1, 0xc7,
	0x01, 0x4a, 0x22, 0x46, 0x12, 0xd7, 0x45, 0x32, 0x3c, 0x7b, 0x67, 0x6c, 0xd5, 0x53, 0x07, 0x42,
	0xfb, 0x66, 0x6c, 0xd1, 0x45, 0x04, 0xe4, 0x3a, 0x81, 0xeb, 0x23, 0xc9, 0xe6, 0xcc, 0xeb, 0x9c,
	0x29, 0x76, 0x33, 0x30, 0xfb, 0x67, 0x6c, 0xb8, 0x11, 0xac, 0x2d, 0x18, 0xe1, 0xa8, 0x60, 0xc4,
	0x8a, 0x1c, 0x7d, 0x83, 0xa3, 0xc5, 0x96, 0x83, 0xd1, 0x37, 0x55, 0x10, 0x1a, 0x0e, 0x21, 0x8e,
	0x1b, 0x23, 0x3e, 0x2b, 0xe5, 0xe2, 0x1f, 0x33, 0x32, 0x7c, 0x06, 0xf1, 0xcf, 0x19, 0x1b, 0x2e,
	0x07, 0x0b, 0x09, 0x66, 0x5e, 0x38, 0x42, 0x30, 0xf2, 0x50, 0x29, 0x80, 0x3f, 0xf4, 0x72, 0x19,
	0x13, 0xb3, 0xfc, 0x63, 0xcf, 0x16, 0x7d, 0x53, 0xc9, 0xa6, 0xe0, 0xfd, 0xc4, 0x31, 0x7d, 0x13,
	0x13, 0x76, 0xbd, 0x27, 0x2b, 0x56, 0x61, 0x86, 0x08, 0xde, 0xe8, 0xc9, 0xea, 0x1f, 0x46, 0xa5,
	0x3d, 0x3b, 0x1b, 0xf3, 0x74, 0x6a, 0x70, 0xf8, 0x6e, 0x72, 0x58, 0xdf, 0x78, 0x30, 0x7c, 0xab,
	0x27, 0xd3, 0x25, 0x9b, 0x3a, 0xdc, 0x30, 0x6a, 0x32, 0x42, 0x9d, 0x98, 0xca, 0x71, 0x85, 0x4f,
	0x3f, 0xb3, 0x32, 0xf1, 0x86, 0x40, 0xa7, 0x66, 0xa5, 0x27, 0x0a, 0x12, 0xb2, 0xde, 0xc0, 0x4f,
	0xcf, 0xca, 0xa8, 0x0c, 0xa3, 0xce, 0xcc, 0xca, 0x0c, 0x18, 0x41, 0x01, 0x41, 0x63, 0x8e, 0x3e,
	0x7e, 0xbc, 0xfd, 0x9c, 0xcc, 0x00, 0x83, 0xf8, 0x8e, 0x12, 0x49, 0x1d, 0x05, 0x8d, 0xb0, 0x74,
	0xc1, 0x9d, 0xb6, 0x34, 0x6b, 0x10, 0xe7, 0xda, 0xd2, 0x49, 0x21, 0x9a, 0x2e, 0xb5, 0xdb, 0x96,
	0xe1, 0x1a, 0x02, 0x3d, 0xdf, 0x96, 0xd5, 0x25, 0x20, 0xc3, 0x8d, 0xbe, 0xd0, 0x96, 0xd5, 0x35,
	0x98, 0x79, 0xb1, 0x2d, 0xab, 0x4b, 0x30, 0x83, 0xaf, 0xf1, 0xa5, 0xb6, 0xac, 0xae, 0xff, 0x44,
	0x5f, 0x6e, 0xdb, 0xe5, 0x41, 0xda, 0x0b, 0x03, 0x6f, 0x3c, 0x09, 0x8a, 0xd9, 0xe7, 0x15, 0x4e,
	0xf0, 0x27, 0x0f, 0x35, 0x0d, 0xc4, 0xab, 0x83, 0x6c, 0x90, 0xa4, 0xc6, 0x89, 0xde, 0x20, 0x1b,
	0x8a, 0x78, 0xad, 0x2d, 0x33, 0xa1, 0xcf, 0x46, 0xe9, 0x55, 0xdd, 0xdd, 0x96, 0x59, 0x35, 0x8c,
	0xda, 0xa3, 0x28, 0xe1, 0x7e, 0xf6, 0x66, 0x04, 0x49, 0xdf, 0x14, 0xbd, 0x57, 0xed, 0x38, 0x8c,
	0xda, 0xd7, 0xb6, 0xe1, 0x22, 0x00, 0x31, 0x61, 0x8f, 0xe3, 0xec, 0x2f, 0xa4, 0xe8, 0x2e, 0xb7,
	0xb9, 0x42, 0x70, 0x45, 0xb9, 0xd3, 0x96, 0x8d, 0x25, 0x7f, 0xfb, 0xa2, 0xad, 0x51, 0x29, 0xb5,
	0x2e, 0x74, 0x64, 0x5b, 0x32, 0xca, 0x17, 0x3b, 0x36, 0xdc, 0x04, 0xd6, 0xe9, 0xb2, 0xfe, 0x97,
	0x59, 0x3a, 0xdd, 0xa5, 0x4e, 0x65, 0x27, 0x7d, 0x98, 0xb8, 0x5c, 0xdd, 0x49, 0x97, 0xaf, 0x74,
	0xec, 0xec, 0xa7, 0xa4, 0xbc, 0xba, 0x78, 0xa6, 0xb5, 0x57, 0xfa, 0x6a, 0x47, 0xf6, 0xc2, 0xfe,
	0x23, 0x97, 0x52, 0xfd, 0x5a, 0xa7, 0xf4, 0xe8, 0x0b, 0x83, 0x59, 0x1d, 0xa3, 0xec, 0xfd, 0xc3,
	0xa1, 0x98, 0xa2, 0xbe, 0xed, 0x94, 0x86, 0x1c, 0x8d, 0xd3, 0x5b, 0xc8, 0x77, 0x1d, 0x99, 0xf0,
	0x8a, 0x52, 0x66, 0xfc, 0xc4, 0x1d, 0x6d, 0xb2, 0x6c, 0x2a, 0xcd, 0xd1, 0xef, 0x3b, 0xc5, 0x43,
	0x99, 0xa7, 0xb8, 0x1e, 0xd0, 0x23, 0x73, 0xc5, 0x13, 0x5b, 0x11, 0x8f, 0xce, 0xf5, 0xad, 0xcc,
	0x8a, 0x03, 0xf9, 0xe2, 0x17, 0xb9, 0x7f, 0xa5, 0x2e, 0x1e, 0x9b, 0x93, 0x5d, 0x58, 0x56, 0xbd,
	0x8b, 0x02, 0xd5, 0x5c, 0x69, 0x58, 0x82, 0x8f, 0xcf, 0xc9, 0x61, 0xc3, 0xd4, 0x22, 0x4e, 0xcc,
	0xc9, 0x61, 0xc3, 0xa4, 0x9e, 0x9c, 0x93, 0xcd, 0xc1, 0xad, 0x27, 0x4e, 0xe0, 0x91, 0x3a, 0x66,
	0x28, 0xa0, 0x28, 0x66, 0x6e, 0x3e, 0x9e, 0xe7, 0xcf, 0x73, 0x57, 0x36, 0x99, 0xc1, 0xcc, 0xfe,
	0xae, 0x1a, 0x5a, 0x14, 0x93, 0x4f, 0x53, 0x72, 0x78, 0xcf, 0x1f, 0xe6, 0xae, 0xbc, 0x9a, 0x02,
	0x2a, 0x3b, 0x96, 0xbf, 0xc9, 0x5d, 0x59, 0x3b, 0xc3, 0xa8, 0xb7, 0xba, 0x7a, 0x83, 0xf5, 0xc6,
	0x10, 0xd6, 0x8a, 0xe4, 0xaf, 0xae, 0xde, 0x9a, 0xfb, 0xc4, 0xbf, 0xbb, 0xf6, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xf3, 0x7a, 0xc2, 0xcb, 0xb6, 0x11, 0x00, 0x00,
}