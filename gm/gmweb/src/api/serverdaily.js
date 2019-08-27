import request from '@/utils/request'

export function getServerDailyStatic({centerPlatformId,startTime,endTime}) {
    let myCenterPlatformId = parseInt(centerPlatformId)
    return request({
        url: '/serverdaily/list',
        method: 'post',
        data: {
            centerPlatformId: myCenterPlatformId,
            startTime: startTime,
            endTime: endTime
        },
    })
}
