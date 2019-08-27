import request from '@/utils/request'

export function getCenterUserList({ userId, sdkType, platformUserId, userName, pageIndex }) {
    let mysdktype = parseInt(sdkType)
    return request({
        url: '/center/centeruser/querylist',
        data: {
            userId: userId, sdkType: mysdktype, userName: userName, pageIndex: pageIndex, platformUserId: platformUserId
        },
        method: 'post'
    })
}

export function getCenterNeiGuaUserList({ userId, sdkType, platformUserId, userName, pageIndex }) {
    let mysdktype = parseInt(sdkType)
    return request({
        url: '/center/centeruser/neigualist',
        data: {
            userId: userId, sdkType: mysdktype, userName: userName, pageIndex: pageIndex, platformUserId: platformUserId
        },
        method: 'post'
    })
}

export function getCenterUserManageList({ userId, sdkType, userName, platformUserId, pageIndex }) {
    let mysdktype = parseInt(sdkType)
    return request({
        url: '/center/centeruser/list',
        data: {
            userId: userId, sdkType: mysdktype, userName: userName, pageIndex: pageIndex, platformUserId: platformUserId
        },
        method: 'post'
    })
}

export function updateCenterUserGm({ id, gm,name,password }) {
    return request({
        url: '/center/centeruser/updategm',
        data: {
            id, gm,name,password
        },
        method: 'post'
    })
}

export function updateCenterUserName({ id, gm,name,password }) {
    return request({
        url: '/center/centeruser/updateusername',
        data: {
            id, gm,name,password
        },
        method: 'post'
    })
}

export function updateCenterForbid({ userId, forbid,forbidTime,reason }) {
    return request({
        url: '/center/centeruser/updateforbid',
        data: {
            userId, forbid,forbidTime,reason
        },
        method: 'post'
    })
}

export function updateCenterIpForbid({ ip, forbid,forbidTime,reason,centerPlatformId,centerServerId,playerId }) {
    return request({
        url: '/center/centeruser/updateipforbid',
        data: {
            ip, forbid,forbidTime,reason,centerPlatformId,centerServerId,playerId
        },
        method: 'post'
    })
}

export function getCenterIpState({ip}){
    return request({
        url: '/center/centeruser/getipstate',
        data: {
            ip
        },
        method: 'post'
    })
}

export function updateCenterIpUnForbid({ip}){
    return request({
        url: '/center/centeruser/updateipunforbid',
        data: {
            ip
        },
        method: 'post'
    })
}

export function getCenterUserInfo({ id }) {
    return request({
        url: '/center/centeruser/userinfo',
        data: {
            id
        },
        method: 'post'
    })
}
