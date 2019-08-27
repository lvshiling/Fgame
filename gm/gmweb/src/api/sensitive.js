import request from '@/utils/request'

export function getSensitive() {
    return request({
        url: '/sensitive/get',
        method: 'post'
    })
}

export function saveUserInfo({ content }) {
    return request({
        url: '/sensitive/add',
        data: {
            content
        },
        method: 'post'
    })
}