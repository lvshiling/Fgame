import request from '@/utils/request'

export function addRedeem({ giftBagName, giftBagDesc, giftBagContent, redeemNum,redeemUseNum, redeemPlayerUseNum, redeemServerUseNum, sdkTypes, sendType, startTime, endTime, minPlayerLevel, minVipLevel }) {
    let myredeemNum = parseInt(redeemNum)
    let myredeemPlayerUseNum = parseInt(redeemPlayerUseNum)
    let myredeemServerUseNum = parseInt(redeemServerUseNum)
    let mysendType = parseInt(sendType)
    let mystartTime = parseInt(startTime)
    let myendTime = parseInt(endTime)
    let myminPlayerLevel = parseInt(minPlayerLevel)
    let myminVipLevel = parseInt(minVipLevel)
    let myredeemUseNum = parseInt(redeemUseNum)
    return request({
        url: '/center/redeem/add',
        data: {
            giftBagName: giftBagName,
            giftBagDesc: giftBagDesc,
            giftBagContent: giftBagContent,
            redeemNum: myredeemNum,
            redeemUseNum:myredeemUseNum,
            redeemPlayerUseNum: myredeemPlayerUseNum,
            redeemServerUseNum: myredeemServerUseNum,
            sdkTypes: sdkTypes,
            sendType: mysendType,
            startTime: mystartTime,
            endTime: myendTime,
            minPlayerLevel: myminPlayerLevel,
            minVipLevel: myminVipLevel
        },
        method: 'post'
    })
}

export function deleteRedeem({ id }) {
    let myid = parseInt(id)
    return request({
        url: '/center/redeem/delete',
        data: {
            id: myid
        },
        method: 'post'
    })
}

export function codeRedeem({ id }) {
    let myid = parseInt(id)
    return request({
        url: '/center/redeem/code',
        data: {
            id: myid
        },
        method: 'post'
    })
}

export function getRedeemCodeList({ id }) {
    let myid = parseInt(id)
    return request({
        url: '/center/redeem/codelist',
        data: {
            id: myid
        },
        method: 'post'
    })
}

export function getRedeemCodeListExport({ id }) {
    let myid = parseInt(id)
    return request({
        url: '/center/redeem/codelistexport',
        data: {
            id: myid
        },
        method: 'post',
        responseType:'blob'
    })
}


export function getRedeemList({ pageIndex, name, sdkType }) {
    let mysdkType = parseInt(sdkType)
    return request({
        url: '/center/redeem/list',
        data: {
            sdkType: mysdkType,
            name: name,
            pageIndex: pageIndex,
        },
        method: 'post'
    })
}
