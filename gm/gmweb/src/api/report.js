import request from '@/utils/request'

export function getOnLineReport({ centerPlatformId, beginTime, endTime }) {
    return request({
        url: '/report/online',
        data: {
            platformId: centerPlatformId,
            startTime: beginTime,
            endTime: endTime
        },
        timeout: 20000,
        method: 'post'
    })
}

export function getNgOnLineReport({ centerPlatformId, beginTime, endTime }) {
    return request({
        url: '/report/ngonline',
        data: {
            platformId: centerPlatformId,
            startTime: beginTime,
            endTime: endTime
        },
        timeout: 20000,
        method: 'post'
    })
}

export function getRecycleReport({ centerPlatformId, beginTime, endTime }) {
    return request({
        url: '/report/recycle',
        data: {
            platformId: centerPlatformId,
            startTime: beginTime,
            endTime: endTime
        },
        timeout: 20000,
        method: 'post'
    })
}

export function getLastOnLineReport() {
    return request({
        url: '/report/onlinetotal',
        timeout: 20000,
        method: 'post'
    })
}

export function getPlayerRetention({ serverId, beginTime, endTime }) {
    let myserverId = parseInt(serverId)
    return request({
        url: '/report/retention',
        data: {
            serverId: myserverId,
            startTime: beginTime,
            endTime: endTime
        },
        timeout: 20000,
        method: 'post'
    })
}