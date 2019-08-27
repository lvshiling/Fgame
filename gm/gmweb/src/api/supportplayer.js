import request from '@/utils/request'

export function getPlayerList({ pageIndex, playerName, serverId, ordercol, ordertype }) {
    let myserverId = parseInt(serverId)
    return request({
        url: '/manage/supportplayerlist',
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

export function privilegeCharge({ channelId, platformId, serverId, playerId, gold, reason, playerName, num }) {
    let myserverid = parseInt(serverId)
    let myplayerId = playerId
    let mygold = parseInt(gold)
    let mychannelid = parseInt(channelId)
    let myplatformId = parseInt(platformId)
    let myNum = parseInt(num)
    return request({
        url: '/manage/privilegecharge',
        data: {
            channelId: mychannelid,
            platformId: myplatformId,
            serverId: myserverid,
            playerId: myplayerId,
            gold: mygold,
            reason: reason,
            playerName: playerName,
            num: myNum
        },
        method: 'post'
    })
}

export function privilegeChargeMulity({ channelId, platformId, serverId, playerId, gold, reason, playerName, allServer, num }) {
    let myserverid = parseInt(serverId)
    let myplayerId = playerId
    let mygold = parseInt(gold)
    let mychannelid = parseInt(channelId)
    let myplatformId = parseInt(platformId)
    let myNum = parseInt(num)
    return request({
        url: '/manage/privilegechargemulity',
        data: {
            channelId: mychannelid,
            platformId: myplatformId,
            serverId: myserverid,
            playerId: myplayerId,
            gold: mygold,
            reason: reason,
            playerName: playerName,
            allServer: allServer,
            num
        },
        method: 'post'
    })
}

export function privilegeSet({ serverId, playerId, privilege, playerName }) {
    let myserverid = parseInt(serverId)
    let myplayerId = playerId
    let myprivilege = parseInt(privilege)
    return request({
        url: '/manage/privilegeset',
        data: {
            serverId: myserverid,
            playerId: myplayerId,
            privilege: myprivilege,
            playerName: playerName,
        },
        method: 'post'
    })
}

export function supportPlayerLog({ channelId, platformId, serverId, playerName, playerId, pageIndex }) {
    let myserverid = parseInt(serverId)
    let mychannelid = parseInt(channelId)
    let myplatformId = parseInt(platformId)
    return request({
        url: '/manage/supportplayerlog',
        data: {
            channelId: mychannelid,
            platformId: myplatformId,
            serverId: myserverid,
            playerName: playerName,
            playerId: playerId,
            pageIndex: pageIndex
        },
        method: 'post'
    })
}