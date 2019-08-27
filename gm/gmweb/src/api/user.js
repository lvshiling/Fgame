import request from '@/utils/request'

export function getUserList(userName, privilege, pageIndex) {
    return request({
        url: '/user/get_list',
        data: {
            userName,
            privilege,
            pageIndex,
        },
        method: 'post'
    })
}

export function saveUserInfo({ userId, password, userName, privilegeid, channelId, platformId }) {
    return request({
        url: '/user/saveuser',
        data: {
            userId,
            password,
            userName,
            privilegeid,
            channelId,
            platformId
        },
        method: 'post'
    })
}

export function deleteUserInfo({ userId }) {
    return request({
        url: '/user/deleteuser',
        data: {
            userId
        },
        method: 'post'
    })
}

export function changePassword({ userId, password }) {
    return request({
        url: '/user/changepwd',
        data: {
            userId,
            password
        },
        method: 'post'
    })
}

export function childPrivilege(){
    return request({
        url: '/user/childprivilege',
        method: 'post'
    })
}