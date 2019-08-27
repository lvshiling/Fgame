import request from '@/utils/request'

export function getChannelList(channelName, pageIndex) {
    return request({
        url: '/channel/list',
        data: {
            channelName,
            pageIndex
        },
        method: 'post'
    })
}

export function getChannelAdd({channelName}) {
    return request({
        url: '/channel/add',
        data: {
            channelName
        },
        method: 'post'
    })
}

export function getChannelUpdate({channelId,channelName}) {
    return request({
        url: '/channel/update',
        data: {
            channelId,
            channelName
        },
        method: 'post'
    })
}

export function getChannelDelete({channelId}) {
    return request({
        url: '/channel/delete',
        data: {
            channelId
        },
        method: 'post'
    })
}

export function getAllChannel() {
    return request({
        url: '/channel/all',
        method: 'post'
    })
}