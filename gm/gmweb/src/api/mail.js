import request from '@/utils/request'

export function getApplyList({ pageIndex, mailState, title, playerId }) {
    let mymilestate = parseInt(mailState)
    return request({
        url: '/manage/applymaillist',
        data: {
            pageIndex: pageIndex,
            mailState: mymilestate,
            title: title,
            playerId: playerId
        },
        method: 'post'
    })
}

export function addmail({ mailType, serverId, title, content, playerlist, proplist, freezTime, effectDays, roleStartTime, roleEndTime, minLevel, maxLevel, sdkType, centerPlatformId, platformId, channelId, bindFlag, remark }) {
    return request({
        url: '/manage/addmail',
        data: {
            mailType, serverId, title, content, playerlist, proplist, freezTime, effectDays, roleStartTime, roleEndTime, minLevel, maxLevel, sdkType, centerPlatformId, platformId, channelId, bindFlag, remark
        },
        method: 'post'
    })
}

export function updatemail({ id, mailType, serverId, title, content, playerlist, proplist, freezTime, effectDays, roleStartTime, roleEndTime, minLevel, maxLevel, sdkType, centerPlatformId, bindFlag, remark }) {
    return request({
        url: '/manage/updatemail',
        data: {
            id, mailType, serverId, title, content, playerlist, proplist, freezTime, effectDays, roleStartTime, roleEndTime, minLevel, maxLevel, sdkType, centerPlatformId, bindFlag, remark
        },
        method: 'post'
    })
}

export function deletemail({ id }) {
    return request({
        url: '/manage/deletemail',
        data: {
            id
        },
        method: 'post'
    })
}

export function getApproveList({ pageIndex, mailState, title, playerId }) {
    let mymilestate = parseInt(mailState)
    return request({
        url: '/manage/approvemaillist',
        data: {
            pageIndex: pageIndex,
            mailState: mymilestate,
            title: title,
            playerId: playerId
        },
        method: 'post'
    })
}

export function approveMail({ id, mailState, approveReason }) {
    return request({
        url: '/manage/approvemail',
        data: {
            id, mailState, approveReason
        },
        method: 'post'
    })
}

export function approveMailMultiple({ id, mailState, approveReason }) {
    return request({
        url: '/manage/approvemailarray',
        data: {
            id, mailState, approveReason
        },
        method: 'post'
    })
}

export function sendMail({ id }) {
    return request({
        url: '/manage/sendmail',
        data: {
            id
        },
        method: 'post'
    })
}