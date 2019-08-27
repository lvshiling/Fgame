import request from '@/utils/request'

export function notice({ channelId, platformId, serverId,content,intervalTime,beginTime,endTime}) {
    return request({
        url: '/notice/notice',
        data: {
            channelId, platformId, serverId,content,intervalTime,beginTime,endTime
        },
        method: 'post'
    })
}

export function noticeList({ successFlag, beginTime, endTime,pageIndex}) {
    return request({
        url: '/notice/list',
        data: {
            successFlag, beginTime, endTime,pageIndex
        },
        method: 'post'
    })
}