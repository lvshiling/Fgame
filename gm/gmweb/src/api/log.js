import request from '@/utils/request'

export function getChatLog({ tableName, beginTime, endTime, platformId, serverType, serverId, pageIndex, playerId, chatContent,chatType }) {
    return request({
        url: '/log/getchatlog',
        method: 'post',
        data: {
            tableName,
            beginTime,
            endTime,
            platformId,
            serverType,
            serverId,
            pageIndex,
            playerId,
            chatContent,
            chatType
        }
    })
}

export function getLog({ tableName, beginTime, endTime, platformId, serverType, serverId, pageIndex, playerId, allianceId }) {
    return request({
        url: '/log/get',
        method: 'post',
        data: {
            tableName,
            beginTime,
            endTime,
            platformId,
            serverType,
            serverId,
            pageIndex,
            playerId,
            allianceId
        }
    })
}

export function getLogMeta(tableName, logType) {
    return request({
        url: '/log/meta',
        method: 'post',
        data: {
            tableName: tableName,
            logType: logType
        }
    })
}


export function getLogMetaMsgList(logType) {
    return request({
        url: '/log/metamsglist',
        method: 'post',
        data: {
            logType: logType
        }
    })
}

export function getPlayerStats({ beginTime, endTime, pageIndex }) {
    return request({
        url: '/log/playerstats',
        method: 'post',
        data: {
            beginTime,
            endTime,
            pageIndex
        }
    })
}