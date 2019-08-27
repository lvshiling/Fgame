import request from '@/utils/request'

export function getCenterNoticeList({ pageIndex }) {
    return request({
        url: '/center/notice/list',
        data: {
            pageIndex
        },
        method: 'post'
    })
}


export function getCenterNoticeLoginAdd({ platformId,content }) {
    return request({
        url: '/center/notice/add',
        data: {
            platformId,content
        },
        method: 'post'
    })
}

export function getCenterNoticeLoginUpdate({ id, content }) {
    return request({
        url: '/center/notice/update',
        data: {
            id, content
        },
        method: 'post'
    })
}

export function getCenterNoticeLoginDelete({ id }) {
    return request({
        url: '/center/notice/delete',
        data: {
            id
        },
        method: 'post'
    })
}

export function getCenterDefaultNotice() {
    return request({
        url: '/center/notice/defaultinfo',
        method: 'post'
    })
}

export function updateCenterDefaultNotice({content}) {
    return request({
        url: '/center/notice/defaultadd',
        data: {
             content
        },
        method: 'post'
    })
}

export function refreshNotice() {
    return request({
        url: '/center/notice/refresh',
        method: 'post'
    })
}