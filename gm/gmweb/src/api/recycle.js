import request from '@/utils/request'

export function getRecycleList({ serverId }) {
    let myserverId = parseInt(serverId)
    return request({
        url: '/recycle/list',
        data: {
            serverId: myserverId,
        },
        method: 'post'
    })
}

export function setRecycleGold({ serverId, gold }) {
    let myserverId = parseInt(serverId)
    let mygold = parseInt(gold)
    return request({
        url: '/recycle/recyclegold',
        data: {
            serverId: myserverId,
            gold: mygold,
        },
        method: 'post'
    })
}