import request from '@/utils/request'

export function getPlayerMail({ pageIndex, serverId, playerId, beginTime, endTime }) {
    let myserverId = parseInt(serverId)
    return request({
        url: '/player/playermail',
        data: {
            playerId: playerId,
            pageIndex: pageIndex,
            begin: beginTime,
            end: endTime,
            serverId: myserverId,
        },
        method: 'post'
    })
}

export function getPlayerLog({ pageIndex, serverId, playerId, beginTime, endTime, tableName }) {
    let myserverId = parseInt(serverId)
    return request({
        url: '/player/mongolog',
        data: {
            playerId: playerId,
            pageIndex: pageIndex,
            begin: beginTime,
            end: endTime,
            serverId: myserverId,
            tableName: tableName,
        },
        method: 'post'
    })
}

export function getPlayerItemChange({ pageIndex, serverId, playerId, itemId, beginTime, endTime }) {
    let myserverId = parseInt(serverId)
    let myItemid = parseInt(itemId)
    return request({
        url: '/player/itemchange',
        data: {
            playerId: playerId,
            pageIndex: pageIndex,
            begin: beginTime,
            end: endTime,
            serverId: myserverId,
            itemId: myItemid,
        },
        method: 'post'
    })
}