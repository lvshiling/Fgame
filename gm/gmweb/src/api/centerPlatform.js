import request from '@/utils/request'

export function getCenterPlatformList(centerPlatformName, pageIndex) {
    return request({
        url: '/center/platform/list',
        data: {
            centerPlatformName,
            pageIndex
        },
        method: 'post'
    })
}

export function getCenterPlatformAdd({centerPlatformName,sdkType}) {
    let mysdkType = parseInt(sdkType)
    return request({
        url: '/center/platform/add',
        data: {
            centerPlatformName:centerPlatformName,
            sdkType:mysdkType
        },
        method: 'post'
    })
}

export function getCenterPlatformUpdate({centerPlatformId,centerPlatformName,sdkType}) {
    let mysdkType = parseInt(sdkType)
    return request({
        url: '/center/platform/update',
        data: {
            centerPlatformId:centerPlatformId,
            centerPlatformName:centerPlatformName,
            sdkType:mysdkType
        },
        method: 'post'
    })
}

export function getCenterPlatformDelete({centerPlatformId}) {
    return request({
        url: '/center/platform/delete',
        data: {
            centerPlatformId
        },
        method: 'post'
    })
}


export function getCenterPlatformMarrySetList(centerPlatformName, pageIndex) {
    return request({
        url: '/center/platform/marrylist',
        data: {
            centerPlatformName,
            pageIndex
        },
        method: 'post'
    })
}

export function updateCenterPlatformMarrySetFlag({centerPlatformId, kindType}) {
    return request({
        url: '/center/platform/marryflag',
        data: {
            centerPlatformId,
            kindType
        },
        method: 'post'
    })
}

export function updateCenterPlatformMarrySetContent({centerPlatformId, marrySet}) {
    return request({
        url: '/center/platform/marrycontent',
        data: {
            centerPlatformId,
            marrySet
        },
        method: 'post'
    })
}

export function getCenterPlatformSetList(centerPlatformName, pageIndex) {
    return request({
        url: '/center/platform/settinglist',
        data: {
            centerPlatformName,
            pageIndex
        },
        method: 'post'
    })
}

export function updateCenterPlatformSaveSetting({centerPlatformId, setting}) {
    return request({
        url: '/center/platform/savesetting',
        data: {
            centerPlatformId,
            setting
        },
        method: 'post'
    })
}

export function getCenterPlatformMetaSetting() {
    return request({
        url: '/center/platform/setting',
        data: {},
        method: 'post'
    })
}

export function getCenterPlatformMarrySetLogList({platformId,pageIndex}) {
    return request({
        url: '/center/platform/marryloglist',
        data: {
            platformId,
            pageIndex
        },
        method: 'post'
    })
}

export function updateCenterPlatformMarrySetSend({centerPlatformId, centerServerId,id,kindType}) {
    return request({
        url: '/center/platform/marrylogsend',
        data: {
            centerPlatformId, centerServerId,id,kindType
        },
        method: 'post'
    })
}