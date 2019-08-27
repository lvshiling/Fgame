import request from '@/utils/request'

export function getCenterTradeItemList({ centerPlatformId,centerServerId,startTime,endTime,tradeId,playerId,level,state,pageIndex }) {
    let myPlatformId = parseInt(centerPlatformId)
    let myServerId = parseInt(centerServerId)
    let myLevel = parseInt(level)
    let myState = parseInt(state)-1
    if(state === undefined || state === ""){
        myState = -1
    }
    return request({
        url: '/report/tradeitem',
        data: {
            platformId:myPlatformId,
            serverId:myServerId,
            startTime:startTime,
            endTime:endTime,
            tradeId:tradeId,
            playerId:playerId,
            level : myLevel,
            state : myState,
            pageIndex : pageIndex
        },
        method: 'post'
    })
}