import request from '@/utils/request'

export function getCenterServerList({ centerServerName, platformId, serverType, pageIndex }) {
    let servertypeid = parseInt(serverType);
    if (serverType === undefined || serverType === "") {
        servertypeid = -1
    }
    let queryPlatformId = parseInt(platformId)

    return request({
        url: '/center/server/list',
        data: {
            centerServerName: centerServerName,
            pageIndex: pageIndex,
            platformId: queryPlatformId,
            serverType: servertypeid
        },
        method: 'post'
    })
}

export function getSimpleCenterServerList({ centerServerName, centerPlatformId, pageIndex }) {
    let queryPlatformId = parseInt(centerPlatformId)

    return request({
        url: '/center/server/simplelist',
        data: {
            centerServerName: centerServerName,
            pageIndex: pageIndex,
            centerPlatformId: queryPlatformId
        },
        method: 'post'
    })
}


export function updateCenterServerParentId(id, parentServerId) {
    return request({
        url: '/center/server/updateparent',
        data: {
            id, parentServerId
        },
        method: 'post'
    })
}

export function updateCenterServerParentIdArray(id, parentServerId) {
    return request({
        url: '/center/server/updateparentarray',
        data: {
            id, parentServerId
        },
        method: 'post'
    })
}

export function updateCenterServerJiaoYiZhanQu(id, jiaoYiZhanQuServerId) {
    let myjiaoYiZhanQuServerId = parseInt(jiaoYiZhanQuServerId);
    return request({
        url: '/center/server/updatejiaoyizhanqu',
        data: {
            id: id,
            jiaoYiZhanQuServerId: myjiaoYiZhanQuServerId
        },
        method: 'post'
    })
}

export function updateCenterServerJiaoYiZhanQuArray(id, jiaoYiZhanQuServerId) {
    let myjiaoYiZhanQuServerId = parseInt(jiaoYiZhanQuServerId);
    return request({
        url: '/center/server/updatejiaoyizhanquarray',
        data: {
            id: id,
            jiaoYiZhanQuServerId: myjiaoYiZhanQuServerId
        },
        method: 'post'
    })
}

export function updateCenterServerPingTaiFu(id, pingTaiFuServerId) {
    let mypingTaiFuServerId = parseInt(pingTaiFuServerId);
    return request({
        url: '/center/server/updatepingtaifu',
        data: {
            id: id,
            pingTaiFuServerId: mypingTaiFuServerId
        },
        method: 'post'
    })
}

export function updateCenterServerPingTaiFuArray(id, pingTaiFuServerId) {
    let mypingTaiFuServerId = parseInt(pingTaiFuServerId);
    return request({
        url: '/center/server/updatepingtaifuarray',
        data: {
            id: id,
            pingTaiFuServerId: mypingTaiFuServerId
        },
        method: 'post'
    })
}

export function updateCenterServerChengZhan(id, chengZhanServerId) {
    let mychengZhanServerId = parseInt(chengZhanServerId);
    return request({
        url: '/center/server/updatechengzhanfu',
        data: {
            id: id,
            chengZhanServerId: mychengZhanServerId
        },
        method: 'post'
    })
}

export function updateCenterServerChengZhanArray(id, chengZhanServerId) {
    let mychengZhanServerId = parseInt(chengZhanServerId);
    return request({
        url: '/center/server/updatechengzhanfuarray',
        data: {
            id: id,
            chengZhanServerId: mychengZhanServerId
        },
        method: 'post'
    })
}

export function updateCenterServerName({ id, serverName }) {
    return request({
        url: '/center/server/simpleupdate',
        data: {
            id, serverName
        },
        method: 'post'
    })
}

export function getCenterServerListByServerType(platformId, serverType) {
    let centerServerType = parseInt(serverType)
    let centerPlatformId = parseInt(platformId)
    return request({
        url: '/center/server/serverlisttype',
        data: {
            serverType: centerServerType,
            platformId: centerPlatformId,
        },
        method: 'post'
    })
}

export function getCenterServerListQuanPingTai() {
    
    return request({
        url: '/center/server/serverlisttype',
        data: {
            serverType: 4,
            platformId: 0,
        },
        method: 'post'
    })
}

export function getCenterServerListChengZhan() {
    
    return request({
        url: '/center/server/serverlisttype',
        data: {
            serverType: 5,
            platformId: 0,
        },
        method: 'post'
    })
}

export function getCenterServerListZhanQuList(platformId) {
    let centerPlatformId = parseInt(platformId)
    return request({
        url: '/center/server/serverlisttype',
        data: {
            serverType: 2,
            platformId: centerPlatformId,
        },
        method: 'post'
    })
}

export function getCenterServerAdd({ serverType, serverId, platformId, serverName, startTime, serverIp, serverPort, serverRemoteIp, serverRemotePort, serverDBIp, serverDBPort, serverDBName, serverDBUser, serverDBPassword, serverTag, serverStatus, parentServerId, preShow }) {
    return request({
        url: '/center/server/add',
        data: {
            serverType, serverId, platformId, serverName, startTime, serverIp, serverPort, serverRemoteIp, serverRemotePort, serverDBIp, serverDBPort, serverDBName, serverDBUser, serverDBPassword, serverTag, serverStatus, parentServerId, preShow
        },
        method: 'post'
    })
}

export function getCenterServerUpdate({ id, serverType, serverId, platformId, serverName, startTime, serverIp, serverPort, serverRemoteIp, serverRemotePort, serverDBIp, serverDBPort, serverDBName, serverDBUser, serverDBPassword, serverTag, serverStatus, parentServerId, preShow }) {
    return request({
        url: '/center/server/update',
        data: {
            id, serverType, serverId, platformId, serverName, startTime, serverIp, serverPort, serverRemoteIp, serverRemotePort, serverDBIp, serverDBPort, serverDBName, serverDBUser, serverDBPassword, serverTag, serverStatus, parentServerId, preShow
        },
        method: 'post'
    })
}

export function getCenterServerDelete({ id }) {
    return request({
        url: '/center/server/delete',
        data: {
            id
        },
        method: 'post'
    })
}

export function getCenterServerPing({ id,serverId,platformId }) {
    return request({
        url: '/center/server/ping',
        data: {
            id,serverId,platformId
        },
        method: 'post'
    })
}


export function refreshCenterServer({ serverId }) {
    return request({
        url: '/center/server/refresh',
        data: {
            serverId
        },
        method: 'post'
    })
}

export function getZhanQuCenterServerList({ centerPlatformId }) {
    let queryPlatformId = parseInt(centerPlatformId)

    return request({
        url: '/center/server/zhanqulist',
        data: {
            centerPlatformId: queryPlatformId
        },
        method: 'post'
    })
}


export function getZhanQuCenterServerListExport({ centerPlatformId }) {
    let queryPlatformId = parseInt(centerPlatformId)

    return request({
        url: '/center/server/zhanqulistexport',
        data: {
            centerPlatformId: queryPlatformId
        },
        method: 'post',
        responseType: 'blob'
    })
}

