import request from '@/utils/request'

export function getServerSupportPoolList({ pageIndex, serverId, centerPlatformId }) {
    let myserverId = parseInt(serverId)
    let mycenterPlatformId = parseInt(centerPlatformId)
    return request({
        url: '/manage/serversppoollist',
        data: {
            serverid: myserverId,
            pageIndex: pageIndex,
            centerPlatformId: mycenterPlatformId,
        },
        method: 'post'
    })
}

export function addServerSupportPool({ serverId, curGold, sdkType, centerPlatformId,percent }) {
    let myserverId = parseInt(serverId)
    let mygold = parseInt(curGold)
    let mysdkType = parseInt(sdkType)
    let mycenterPlatformId = parseInt(centerPlatformId)
    let mypercent = parseInt(percent)
    return request({
        url: '/manage/addserversppool',
        data: {
            serverid: myserverId,
            gold: mygold,
            sdkType: mysdkType,
            centerPlatformId: mycenterPlatformId,
            percent:mypercent,
        },
        method: 'post'
    })
}

export function updateServerSupportPool({ id, curGold,percent }) {
    let myid = parseInt(id)
    let mygold = parseInt(curGold)
    let mypercent = parseInt(percent)
    return request({
        url: '/manage/updateserversppool',
        data: {
            id: myid,
            gold: mygold,
            percent:mypercent,
        },
        method: 'post'
    })
}

export function deleteServerSupportPool({ id }) {
    let myid = parseInt(id)
    return request({
        url: '/manage/deleteserversppool',
        data: {
            id: myid,
        },
        method: 'post'
    })
}


export function getPlatformSupportPoolList({ pageIndex, centerPlatformId }) {
    let mycenterPlatformId = parseInt(centerPlatformId)
    return request({
        url: '/manage/platformpoollist',
        data: {
            pageIndex: pageIndex,
            centerPlatformId: mycenterPlatformId,
        },
        method: 'post'
    })
}

export function addPlatformSupportPoolSet({ gold, centerPlatformId,percent }) {
    let mygold = parseInt(gold)
    let mycenterPlatformId = parseInt(centerPlatformId)
    let mypercent = parseInt(percent)
    return request({
        url: '/manage/addplatformpool',
        data: {
            gold: mygold,
            centerPlatformId: mycenterPlatformId,
            percent:mypercent,
        },
        method: 'post'
    })
}

export function updatePlatformSupportPoolSet({ id, gold,percent }) {
    let myid = parseInt(id)
    let mygold = parseInt(gold)
    let mypercent = parseInt(percent)
    return request({
        url: '/manage/updateplatformpool',
        data: {
            id: myid,
            gold: mygold,
            percent:mypercent,
        },
        method: 'post'
    })
}

export function deletePlatformSupportPoolSet({ id }) {
    let myid = parseInt(id)
    return request({
        url: '/manage/deleteplatformpool',
        data: {
            id: myid,
        },
        method: 'post'
    })
}