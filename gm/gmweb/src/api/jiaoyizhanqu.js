import request from '@/utils/request'

export function addJiaoYiZhanQu({ serverId, zhanquName, platformId }) {
    let myserverId = parseInt(serverId);
    let myPlatformId = parseInt(platformId);
    return request({
        url: '/center/zhanqu/add',
        data: {
            serverId: myserverId,
            zhanquName: zhanquName,
            platformId: myPlatformId
        },
        method: 'post'
    })
}


export function updateJiaoYiZhanQu({ id, serverId, zhanquName, platformId }) {
    let myId = parseInt(id);
    let myserverId = parseInt(serverId);
    let myPlatformId = parseInt(platformId);
    return request({
        url: '/center/zhanqu/update',
        data: {
            id: myId,
            serverId: myserverId,
            zhanquName: zhanquName,
            platformId: myPlatformId
        },
        method: 'post'
    })
}

export function deleteJiaoYiZhanQu({ id }) {
    let myId = parseInt(id);
    return request({
        url: '/center/zhanqu/delete',
        data: {
            id: myId
        },
        method: 'post'
    })
}


export function getJiaoYiZhanQuList({ platformId, pageIndex }) {
    let myplatformId = parseInt(platformId);
    return request({
        url: '/center/zhanqu/list',
        data: {
            platformId: myplatformId,
            pageIndex: pageIndex
        },
        method: 'post'
    })
}

export function getJiaoYiZhanQuListAll({ platformId }) {
    let myplatformId = parseInt(platformId);
    return request({
        url: '/center/zhanqu/alllist',
        data: {
            platformId: myplatformId
        },
        method: 'post'
    })
}

