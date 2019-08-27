import request from '@/utils/request'

export function getAllianceList({ pageIndex, allianceName, serverId, ordercol, ordertype }) {
    let myserverId = parseInt(serverId)
    return request({
        url: '/alliance/list',
        data: {
            allianceName: allianceName,
            serverId: myserverId,
            pageIndex: pageIndex,
            ordertype: ordertype,
            ordercol: ordercol,
        },
        method: 'post'
    })
}

export function getServerState({ serverId }) {
    let myserverId = parseInt(serverId)
    return request({
        url: '/alliance/serverstate',
        data: {
            serverId: myserverId
        },
        method: 'post'
    })
}

export function registerServerState({ serverId,open }) {
    let myserverId = parseInt(serverId)
    let myopen = parseInt(open)
    return request({
        url: '/alliance/serverset',
        data: {
            serverId: myserverId,
            open:myopen
        },
        method: 'post'
    })
}

export function serverSetloglist({serverId,pageIndex}){
    let myserverId = parseInt(serverId)
    return request({
        url: '/alliance/serversetloglist',
        data: {
            serverId: myserverId,
            pageIndex:pageIndex
        },
        method: 'post'
    })
}

export function allianceGongGaoForm({ serverId,allianceId,gongGao }) {
    let myserverId = parseInt(serverId)
    let myallianceId=String(allianceId)
    return request({
        url: '/alliance/modifygonggao',
        data: {
            serverId: myserverId,
            gongGao: gongGao,
            allianceId:myallianceId
        },
        method: 'post'
    })
}

export function allianceDismissForm({serverId,allianceId}){
    let myserverId = parseInt(serverId)
    let myallianceId=String(allianceId)
    return request({
        url: '/alliance/dismiss',
        data: {
            serverId: myserverId,
            allianceId:myallianceId
        },
        method: 'post'
    })
}