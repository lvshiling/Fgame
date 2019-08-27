import { Queue } from "@/websocket/queue";
export function QueueDelegate() {
    this.query = new Queue()
    this.delegate = new Map()
    this.activity = true

    QueueDelegate.prototype.begin = function () {
        if (!this.query) {
            return
        }
        setTimeout(() => {
            // console.log("队列调用执行")
            if (!this.query.isEmpty()) {
                while (!this.query.isEmpty()) { //不为空的时候
                    let msg = this.query.dequeue()
                    if (!msg) {
                        continue
                    }
                    let msgtype = msg.getMessagetype()
                    if (!msgtype) {
                        continue
                    }
                    if (!this.delegate.has(msgtype)) {
                        continue
                    }
                    
                    let doFunc = this.delegate.get(msgtype)
                    if (!doFunc) {
                        continue
                    }
                    try {
                        doFunc(msg)
                    } catch (error) {
                        console.log(error)
                    }
                }
            }
            if (this.activity) {
                this.begin()
            }
        }, 100)
    }

    //注册执行队列
    QueueDelegate.prototype.register = function (msgtype, func) {
        this.delegate.set(msgtype, func)
    }

    QueueDelegate.prototype.getQuery = function () {
        return this.query
    }
    QueueDelegate.prototype.stop = function () {
        this.activity = false
    }
}