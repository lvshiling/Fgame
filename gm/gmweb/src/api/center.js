import request from '@/utils/request'

export function getCenterPlatList() {
    return request({
        url: '/center/platform',
        method: 'post'
    })
}

export function getCenterGroupList(centerPlatformId) {
    return request({
        url: '/center/group',
        data: { centerPlatformId },
        method: 'post'
    })
}

export function getCenterServerList(centerPlatformId) {
    return request({
        url: '/center/server',
        data: { centerPlatformId },
        method: 'post'
    })
}

export function getAllCenterServerList(centerPlatformId) {
    return request({
        url: '/center/allserver',
        data: { centerPlatformId },
        method: 'post'
    })
}

export function getAllUserCenterServerList() {
    return request({
        url: '/center/alluserserver',
        method: 'post'
    })
}

export function getAllSdkType() {
    return request({
        url: '/center/sdktype',
        method: 'post'
    })
}