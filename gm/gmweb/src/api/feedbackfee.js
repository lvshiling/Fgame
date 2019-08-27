import request from '@/utils/request'

export function getGameFeedBackFeeList({ serverId, playerId, pageIndex, startTime, endTime, code }) {
    return request({
        url: '/feedbackfee/gamelist',
        data: {
            serverId, playerId, pageIndex, startTime, endTime, code
        },
        method: 'post'
    })
}


export function getCenterFeedBackFeeList({ platformId,serverId, playerId, pageIndex, startTime, endTime, code }) {
    let myPlatformid = parseInt(platformId);
    return request({
        url: '/feedbackfee/centerlist',
        data: {
            platformId:myPlatformid,
            serverId:serverId, 
            playerId:playerId, 
            pageIndex:pageIndex, 
            startTime:startTime, 
            endTime:endTime,
            code:code
        },
        method: 'post'
    })
}