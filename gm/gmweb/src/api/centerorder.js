import request from '@/utils/request'

export function getCenterOrderList({ orderId, sdkOrderId, pageIndex, sdkType, startTime, endTime }) {
    return request({
        url: '/center/order/list',
        data: {
            orderId, sdkOrderId, pageIndex, sdkType, startTime, endTime
        },
        method: 'post'
    })
}

export function getGameOrderList({ pageIndex, ordercol, ordertype, serverId, startTime, endTime, minAmount, maxAmount, playerId, userId, orderId, sdkOrderId, playerName, sdkType }) {
    return request({
        url: '/center/order/gamelist',
        data: {
            pageIndex, ordercol, ordertype, serverId, startTime, endTime, minAmount, maxAmount, playerId, userId, orderId, sdkOrderId, playerName, sdkType
        },
        method: 'post'
    })
}

export function getGameOrderStatic({ serverId, sdkType }) {
    return request({
        url: '/center/order/gamestatic',
        data: {
            serverId, sdkType
        },
        method: 'post'
    })
}

export function getCenterOrderStatic({ sdkType }) {
    let mysdkType = parseInt(sdkType)
    return request({
        url: '/center/order/centerstatic',
        data: {
            sdkType: mysdkType
        },
        method: 'post'
    })
}

export function getCenterOrderDateStatic({ startTime, endTime, channelId, platformId, serverId }) {
    let myserverId = serverId
    let myplatformId = platformId
    let mychannelId = channelId
    for (let i = 0, len = serverId.length; i < len; i++) {
        if (serverId[i] == -1) {
            myserverId = []
            break;
        }
    }
    if (myserverId.length == 0) {
        for (let i = 0, len = platformId.length; i < len; i++) {
            if (platformId[i] == -1) {
                myplatformId = []
                break;
            }
        }
    }

    return request({
        url: '/center/order/centerdatestatic',
        data: {
            startTime: startTime,
            endTime: endTime,
            channelId: mychannelId,
            platformId: myplatformId,
            serverId: myserverId
        },
        method: 'post'
    })
}
/** 订单总量 充量汇总*/
export function getOrderDatePlatformStatic({ startTime, endTime, channelId, platformId }) {
    let channelIdList = [];
    let platformIdIdList = [];
    if(channelId != undefined){
        channelIdList.push(channelId)
    }
    if(platformId != undefined){
        platformIdIdList.push(platformId)
    }
    return request({
        url: '/center/order/centerdateplatformstatic',
        data:
        {
            startTime: startTime,
            endTime: endTime,
            channelId: channelIdList,
            platformId:platformIdIdList
        },
        method: 'post'
    })
}

export function getCenterOrderStaticTotal() {
    return request({
        url: '/center/order/centertotalstatic',
        method: 'post'
    })
}