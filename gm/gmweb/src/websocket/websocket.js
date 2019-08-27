import basemessage from '@/proto/basic_pb'
import messages from '@/proto/login_pb'
import commonpb from '@/proto/common_pb'
import messagetype from '@/proto/messagetype_pb'
import { getToken } from '@/utils/auth'

export function newWebSocket(url, queueDelegate, scoketEvent) {
  const wsuri = url
  const websock = new WebSocket(wsuri)

  let pingTimer

  websock.onopen = function() {
    // 打开
    console.log('WebSocket连接')
    sendAuth(websock)
    queueDelegate.begin()
  }
  websock.onmessage = function(e) {
    var fileReader = new FileReader()
    fileReader.onload = function(progressEvent) {
      var arrayBuffer = this.result // arrayBuffer即为blob对应的arrayBuffer
      // console.log(arrayBuffer);
      const rec = basemessage.Message.deserializeBinary(arrayBuffer)
      // console.log(rec);
      // 处理并加入队列
      if (!rec) {
        return
      }
      const msgType = rec.getMessagetype()
      if (!msgType) {
        return
      }
      if (msgType == messagetype.QiPaiMessageType.GCLOGINTYPE) { // 登陆后的操作
        pingTimer = setInterval(() => {
          sendPing(websock)
        }, 5000)
        return
      }
      if (msgType == messagetype.QiPaiMessageType.GCPINGTYPE) { // ping后的回复
        return
      }
      const queue = queueDelegate.getQuery()
      if (queue) {
        console.log('放入队列,messageType:' + msgType)
        queue.enqueue(rec)
      }
    }
    fileReader.readAsArrayBuffer(e.data)
  }
  websock.onclose = function() {
    // 关闭
    console.log('WebSocket关闭')
    if (pingTimer) {
      clearInterval(pingTimer)
    }
    queueDelegate.stop()
    websock.close()
    if (scoketEvent && scoketEvent.onclose && typeof scoketEvent.onclose === 'function') {
      scoketEvent.onclose()
    }
  }
  websock.onerror = function() {
    // 失败
    console.log('WebSocket连接失败')
    if (pingTimer) {
      clearInterval(pingTimer)
    }
    queueDelegate.stop()
    if (scoketEvent && scoketEvent.onerror && typeof scoketEvent.onerror === 'function') {
      scoketEvent.onerror()
    }
  }

  return websock
}

// 发送登陆
function sendAuth(websock) {
  console.log('发送验证')
  const login = new messages.CGLogin()
  const mytoken = getToken()
  login.setPlayerid(1)
  login.setToken(mytoken)
  const mess = new basemessage.Message()
  mess.setMessagetype(messagetype.QiPaiMessageType.CGLOGINTYPE)
  mess.setExtension(messages.cglogin, login)
  // console.log(mess);
  websock.send(mess.serializeBinary())
}

// 发送ping
function sendPing(websock) {
  const ping = new messages.CGPing()
  const msg = new basemessage.Message()
  msg.setMessagetype(messagetype.QiPaiMessageType.CGPINGTYPE)
  msg.setExtension(messages.cgping, ping)
  // console.log('发送ping')
  // console.log(msg)
  websock.send(msg.serializeBinary())
}

// 处理接受信息
// function handlerReceive(message, queue, timer, scoket) {
//     if (!message) {
//         return
//     }
//     let msgType = message.getMessagetype()
//     if (!msgType) {
//         return
//     }
//     if (msgType == messagetype.GCLOGINTYPE) {//登陆后的操作
//         handlerAuth(timer)
//         return
//     }
//     if (msgType == messagetype.GCPINGTYPE) { //ping后的回复
//         return
//     }

//     if (queue) {
//         queue.enqueue(message)
//     }
// }
