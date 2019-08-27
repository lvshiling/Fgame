export function Queue() {
    this.items = [];
    //初始化队列方法
    if (typeof Queue.prototype.push != "function") {
        //入队
        Queue.prototype.enqueue = function () {
            var len = arguments.length;
            if (len == 0) {
                return;
            }
            for (var i = 0; i < len; i++) {
                this.items.push(arguments[i])
            }
        }
        //出队
        Queue.prototype.dequeue = function () {
            var result = this.items.shift();
            return typeof result != 'undefined' ? result : false;
        }
        //返回队首元素
        Queue.prototype.front = function () {
            return this.items[items.length - 1];
        }
        //队列是否为空
        Queue.prototype.isEmpty = function () {
            return this.items.length == 0;
        }
        //返回队列长度
        Queue.prototype.size = function () {
            return this.items.length;
        }
        //清空队列
        Queue.prototype.clear = function () {
            this.items = [];
        }
        //返回队列
        Queue.prototype.show = function () {
            return this.items;
        }
    }
}