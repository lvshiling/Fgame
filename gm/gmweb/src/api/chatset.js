import request from '@/utils/request'
import { parseTime } from '@/utils/index'

export function addChatSet({ centerPlatformId, centerServerArray, worldVip, worldPlayerLevel, pChatVip, pChatPlayerLevel, guildVip, guildPlayerLevel,sdkType,teamVip,teamPlayerLevel }) {
    let platid = parseInt(centerPlatformId)
    // let serverid = parseInt(centerServerId)
    let myskdType = parseInt(sdkType)
    return request({
        url: '/center/chatset/add',
        data: {
            centerPlatformId: platid,
            centerServerId: centerServerArray,
            worldVip: parseInt(worldVip),
            worldPlayerLevel: parseInt(worldPlayerLevel),
            pChatVip: parseInt(pChatVip),
            pChatPlayerLevel: parseInt(pChatPlayerLevel),
            guildVip: parseInt(guildVip),
            guildPlayerLevel: parseInt(guildPlayerLevel),
            sdkType:myskdType,
            teamVip:parseInt(teamVip),
            teamPlayerLevel:parseInt(teamPlayerLevel),
        },
        method: 'post'
    })
}


export function updateChatSet({ chatSetId, centerPlatformId, centerServerArray, worldVip, worldPlayerLevel, pChatVip, pChatPlayerLevel, guildVip, guildPlayerLevel,sdkType,teamVip,teamPlayerLevel }) {
    let id = parseInt(chatSetId)
    let platid = parseInt(centerPlatformId)
    // let serverid = parseInt(centerServerId)
    return request({
        url: '/center/chatset/update',
        data: {
            chatSetId: id,
            centerPlatformId: platid,
            centerServerId: centerServerArray,
            worldVip: parseInt(worldVip),
            worldPlayerLevel: parseInt(worldPlayerLevel),
            pChatVip: parseInt(pChatVip),
            pChatPlayerLevel: parseInt(pChatPlayerLevel),
            guildVip: parseInt(guildVip),
            guildPlayerLevel: parseInt(guildPlayerLevel),
            sdkType:parseInt(sdkType),
            teamVip:parseInt(teamVip),
            teamPlayerLevel:parseInt(teamPlayerLevel),
        },
        method: 'post'
    })
}

export function deleteChatSet({ chatSetId }) {
    let id = parseInt(chatSetId)
    return request({
        url: '/center/chatset/delete',
        data: { chatSetId: id },
        method: 'post'
    })
}

export function getChatSetList({ centerPlatformId, centerServerId, pageIndex }) {
    let platid = parseInt(centerPlatformId)
    let serverid = parseInt(centerServerId)
    return request({
        url: '/center/chatset/list',
        data: {
            centerPlatformId: platid,
            centerServerId: serverid,
            pageIndex: pageIndex
        },
        method: 'post'
    })
}

//平台聊天配置

export function addChatSetPlatform({ centerPlatformId, worldVip, worldPlayerLevel, pChatVip, pChatPlayerLevel, guildVip, guildPlayerLevel,teamVip,teamPlayerLevel }) {
    let platid = parseInt(centerPlatformId)
    return request({
        url: '/center/chatset/addplatform',
        data: {
            centerPlatformId: platid,
            worldVip: parseInt(worldVip),
            worldPlayerLevel: parseInt(worldPlayerLevel),
            pChatVip: parseInt(pChatVip),
            pChatPlayerLevel: parseInt(pChatPlayerLevel),
            guildVip: parseInt(guildVip),
            guildPlayerLevel: parseInt(guildPlayerLevel),
            teamVip:parseInt(teamVip),
            teamPlayerLevel:parseInt(teamPlayerLevel),
        },
        method: 'post'
    })
}


export function updateChatSetPlatform({ chatSetId, centerPlatformId, worldVip, worldPlayerLevel, pChatVip, pChatPlayerLevel, guildVip, guildPlayerLevel,teamVip,teamPlayerLevel }) {
    let id = parseInt(chatSetId)
    let platid = parseInt(centerPlatformId)
    // let serverid = parseInt(centerServerId)
    return request({
        url: '/center/chatset/updateplatform',
        data: {
            chatSetId: id,
            centerPlatformId: platid,
            worldVip: parseInt(worldVip),
            worldPlayerLevel: parseInt(worldPlayerLevel),
            pChatVip: parseInt(pChatVip),
            pChatPlayerLevel: parseInt(pChatPlayerLevel),
            guildVip: parseInt(guildVip),
            guildPlayerLevel: parseInt(guildPlayerLevel),
            teamVip:parseInt(teamVip),
            teamPlayerLevel:parseInt(teamPlayerLevel),
        },
        method: 'post'
    })
}

export function deleteChatSetPlatform({ chatSetId }) {
    let id = parseInt(chatSetId)
    return request({
        url: '/center/chatset/deleteplatform',
        data: { chatSetId: id },
        method: 'post'
    })
}

export function getChatSetListPlatform({ centerPlatformId, pageIndex }) {
    let platid = parseInt(centerPlatformId)
    return request({
        url: '/center/chatset/listplatform',
        data: {
            centerPlatformId: platid,
            pageIndex: pageIndex
        },
        method: 'post'
    })
}