import request from '@/utils/request'



export function getServerDoubleCharge({ serverId }) {
    let myserverId = parseInt(serverId)
    return request({
        url: '/singleserver/doublecharge',
        data: {
            serverId: myserverId
        },
        method: 'post'
    })
}

export function resetServerDoubleCharge({ serverId }) {
    let myserverId = parseInt(serverId)
    let myopen = parseInt(open)
    return request({
        url: '/singleserver/doublechargereset',
        data: {
            serverId: myserverId,
        },
        method: 'post'
    })
}

export function serverDoubleChargeLoglist({serverId,pageIndex}){
    let myserverId = parseInt(serverId)
    return request({
        url: '/singleserver/doublechargeloglist',
        data: {
            serverId: myserverId,
            pageIndex:pageIndex
        },
        method: 'post'
    })
}