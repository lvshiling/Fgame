import request from '@/utils/request'

export function getGoldChange({ platformId, serverId, startMoney, endMoney, startTime, endTime, goldType }) {
    return request({
        url: '/report/goldchange',
        method: 'post',
        data: {
            platformId, serverId, startMoney, endMoney, startTime, endTime, goldType
        }
    })
}

export function getNewGoldChange({ platformId, serverId, startMoney, endMoney, startTime, endTime, goldType }) {
    return request({
        url: '/report/newgoldchange',
        method: 'post',
        data: {
            platformId, serverId, startMoney, endMoney, startTime, endTime, goldType
        }
    })
}

export function getNewBindGold({ platformId, serverId, startMoney, endMoney, startTime, endTime, goldType }) {
    return request({
        url: '/report/newbindgold',
        method: 'post',
        data: {
            platformId, serverId, startMoney, endMoney, startTime, endTime, goldType
        }
    })
}

export function getGoldChangeType({goldType}) {
    return request({
        url: '/report/goldchangetype',
        method: 'post',
        data: {
            goldType
        }
    })
}