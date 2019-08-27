var PrivilegeType = {
    PrivilegeLevelSuperChannel: 1,
    PrivilegeLevelChannel: 2,
    PrivilegeLevelPlatform: 3,
    PrivilegeLevelKeFu: 4,
    PrivilegeLevelMinitor: 5,
    PrivilegeLevelSuperChannelService: 6,
    PrivilegeLevelCommonKeFu : 7,
    PrivilegeLevelGaoJiKeFu: 8,
    PrivilegeLevelNeiGua: 9,
    PrivilegeLevelGs: 10,
    PrivilegeLevelAdmin: 999
}

export var PrivilegeType

var PrivilegeTypeMap = [
    { key: PrivilegeType.PrivilegeLevelSuperChannel, name: "总渠道管理员" },
    { key: PrivilegeType.PrivilegeLevelChannel, name: "渠道管理员" },
    { key: PrivilegeType.PrivilegeLevelPlatform, name: "平台管理员" },
    { key: PrivilegeType.PrivilegeLevelKeFu, name: "平台客服" },
    { key: PrivilegeType.PrivilegeLevelMinitor, name: "监控" },
    { key: PrivilegeType.PrivilegeLevelAdmin, name: "超级管理员" },
    { key: PrivilegeType.PrivilegeLevelSuperChannelService, name: "总渠道客服" },
    { key: PrivilegeType.PrivilegeLevelCommonKeFu, name: "普通客服" },
    { key: PrivilegeType.PrivilegeLevelGaoJiKeFu, name: "高级客服" },
    { key: PrivilegeType.PrivilegeLevelNeiGua, name: "内挂管理" },
    { key: PrivilegeType.PrivilegeLevelGs, name: "GS管理" }
]

export var PrivilegeTypeMap