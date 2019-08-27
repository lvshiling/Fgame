import request from '@/utils/request'

export function getPlayerList({ pageIndex, playerName, serverId, ordercol, ordertype }) {
    let myserverId = parseInt(serverId)
    return request({
        url: '/player/list',
        data: {
            playerName: playerName,
            serverId: myserverId,
            pageIndex: pageIndex,
            ordertype: ordertype,
            ordercol: ordercol,
        },
        method: 'post'
    })
}

export function getFengJinPlayerList({ pageIndex, playerName, centerPlatformId, centerServerId, reason }) {
    return request({
        url: '/player/fengjinlist',
        data: {
            pageIndex: pageIndex,
            playerName: playerName,
            centerPlatformId: centerPlatformId,
            centerServerId: centerServerId,
            reason: reason,
        },
        method: 'post'
    })
}

export function getJinYanPlayerList({ pageIndex, playerName, centerPlatformId, centerServerId, reason }) {
    return request({
        url: '/player/jinyanlist',
        data: {
            pageIndex: pageIndex,
            playerName: playerName,
            centerPlatformId: centerPlatformId,
            centerServerId: centerServerId,
            reason: reason,
        },
        method: 'post'
    })
}


export function forbidPlayer({ centerPlatformId, centerServerId, reason, playerId, forbidTime }) {
    return request({
        url: '/player/forbid',
        data: {
            reason: reason,
            playerId: playerId,
            centerPlatformId: centerPlatformId,
            centerServerId: centerServerId,
            forbidTime: forbidTime
        },
        method: 'post'
    })
}

export function unForbidPlayer({ centerPlatformId, centerServerId, playerId }) {
    return request({
        url: '/player/unforbid',
        data: {
            playerId: playerId,
            centerPlatformId: centerPlatformId,
            centerServerId: centerServerId
        },
        method: 'post'
    })
}

export function forbidChatPlayer({ centerPlatformId, centerServerId, reason, playerId, forbidTime }) {
    return request({
        url: '/player/forbidchat',
        data: {
            reason: reason,
            playerId: playerId,
            centerPlatformId: centerPlatformId,
            centerServerId: centerServerId,
            forbidTime: forbidTime
        },
        method: 'post'
    })
}



export function unForbidChatPlayer({ centerPlatformId, centerServerId, playerId }) {
    return request({
        url: '/player/unforbidchat',
        data: {
            playerId: playerId,
            centerPlatformId: centerPlatformId,
            centerServerId: centerServerId
        },
        method: 'post'
    })
}



export function ignoreChatPlayer({ centerPlatformId, centerServerId, reason, playerId, forbidTime }) {
    return request({
        url: '/player/ignorechat',
        data: {
            reason: reason,
            playerId: playerId,
            centerPlatformId: centerPlatformId,
            centerServerId: centerServerId,
            forbidTime: forbidTime
        },
        method: 'post'
    })
}



export function unIgnoreChatPlayer({ centerPlatformId, centerServerId, playerId }) {
    return request({
        url: '/player/unIgnorechat',
        data: {
            playerId: playerId,
            centerPlatformId: centerPlatformId,
            centerServerId: centerServerId
        },
        method: 'post'
    })
}

export function getJinMoPlayerList({ pageIndex, playerName, centerPlatformId, centerServerId, reason }) {
    return request({
        url: '/player/ignoreList',
        data: {
            pageIndex: pageIndex,
            playerName: playerName,
            centerPlatformId: centerPlatformId,
            centerServerId: centerServerId,
            reason: reason,
        },
        method: 'post'
    })
}

export function getPlayerInfo({ centerPlatformId, centerServerId, playerId,ip }) {
    return request({
        url: '/player/playerinfo',
        data: {
            centerPlatformId, centerServerId, playerId,ip
        },
        method: 'post'
    })
}

export function kickOutPlayer({ centerPlatformId, centerServerId, reason, playerId }) {
    return request({
        url: '/player/kickout',
        data: {
            reason: reason,
            playerId: playerId,
            centerPlatformId: centerPlatformId,
            centerServerId: centerServerId
        },
        method: 'post'
    })
}

export function getPlayerLevelStatic({serverId}){
    let myserverId = parseInt(serverId)
    return request({
        url: '/player/levelcount',
        data: {
            serverId: myserverId
        },
        method: 'post'
    })
}

export function getPlayerLevelStaticExport({serverId}){
    let myserverId = parseInt(serverId)
    return request({
        url: '/player/levelcountexport',
        data: {
            serverId: myserverId
        },
        method: 'post',
        responseType:'blob'
    })
}

export function getPlayerQuestStatic({serverId}){
    let myserverId = parseInt(serverId)
    return request({
        url: '/player/questcount',
        data: {
            serverId: myserverId
        },
        method: 'post'
    })
}

export function getPlayerQuestStaticExport({serverId}){
    let myserverId = parseInt(serverId)
    return request({
        url: '/player/questcountexport',
        data: {
            serverId: myserverId
        },
        method: 'post',
        responseType:'blob'
    })
}