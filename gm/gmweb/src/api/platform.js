import request from '@/utils/request'

export function getPlatformList({ platformName, channelId, pageIndex }) {
    let mychannelid = parseInt(channelId)
    return request({
        url: '/platform/list',
        data: {
            platformName: platformName,
            channelId: mychannelid,
            pageIndex: pageIndex
        },
        method: 'post'
    })
}

export function getAllPlatformList() {
    return request({
        url: '/platform/all',
        method: 'post'
    })
}

export function setPlatformAdd({ platformName, channelId, centerPlatformId,sdkType,signKey }) {
    return request({
        url: '/platform/add',
        data: {
            platformName,
            channelId,
            centerPlatformId,
            sdkType,
            signKey
        },
        method: 'post'
    })
}

export function setPlatformUpdate({ platformId, platformName, channelId, centerPlatformId,sdkType,signKey }) {
    return request({
        url: '/platform/update',
        data: {
            platformId,
            platformName,
            channelId,
            centerPlatformId,
            sdkType,
            signKey
        },
        method: 'post'
    })
}

export function setPlatformDelete({ platformId }) {
    return request({
        url: '/platform/delete',
        data: {
            platformId
        },
        method: 'post'
    })
}

export function refreshSdk() {
    return request({
        url: '/platform/refreshsdk',
        method: 'post'
    })
}