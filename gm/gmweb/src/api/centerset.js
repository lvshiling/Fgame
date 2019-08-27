import request from '@/utils/request'

export function getClientVersion() {
    return request({
        url: '/center/set/clientverionget',
        data: {
        },
        method: 'post'
    })
}


export function setClientVersion({androidVersion,iosVersion}) {
    return request({
        url: '/center/set/clientverionset',
        data: {
            iosVersion,androidVersion
        },
        method: 'post'
    })
}

export function getPlatformServerConfig() {
    return request({
        url: '/center/set/platformserverconfigget',
        data: {
        },
        method: 'post'
    })
}

export function setPlatformServerConfig({tradeServerIp}) {
    return request({
        url: '/center/set/platformserverconfigset',
        data: {
            tradeServerIp
        },
        method: 'post'
    })
}