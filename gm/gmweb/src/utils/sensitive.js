/**
* @description
* 构造敏感词map,即一个字符串的二叉树
* @private
* @returns
*/
export function makeSensitiveMap(sensitiveWordList) {
    // 构造根节点
    sensitiveWordList.sort(compare)
    const result = new Map();
    for (let j = 0; j < sensitiveWordList.length; j++) {
        let word = sensitiveWordList[j];
        if (word.length == 0) {
            continue
        }
        let map = result;
        for (let i = 0; i < word.length; i++) {
            // 依次获取字
            const char = word.charAt(i);
            // 判断是否存在
            if (map.get(char)) {
                // 获取下一层节点
                map = map.get(char);
                // map.set('laster', true);
            } else {
                // 将当前节点设置为非结尾节点
                // if (map.get('laster') === true) {
                //     map.set('laster', false);
                // }
                const item = new Map();
                // 新增节点默认为结尾节点
                if (i == word.length - 1) { //最后一个节点的时候
                    item.set('laster', true);
                } else {
                    item.set('laster', false);
                }
                map.set(char, item);
                map = map.get(char);
            }
        }
    }
    return result;
}

export function replaceSensitiveWord(sensitiveMap, txt, beginSp, endSp) {
    let replaceArray = getSensitiveWordArrayIndex(sensitiveMap, txt)
    if (!replaceArray || replaceArray.length == 0) {
        return txt
    }

    let result = ''
    let lastStartIndex = 0
    for (let i = 0, len = replaceArray.length; i < len; i++) {
        let replaceItem = replaceArray[i]
        result += txt.substring(lastStartIndex, replaceItem.start)
        lastStartIndex = replaceItem.end
        result += beginSp + txt.substring(replaceItem.start, replaceItem.end) + endSp

        if (i == len - 1) {
            result += txt.substring(replaceItem.end)
        }

    }
    return result
}

function compare(x, y) {
    if (x.length < y.length) {
        return -1;
    } else if (x.length > y.length) {
        return 1;
    } else {
        return 0;
    }
}

/**
* @description
* 检查敏感词是否存在
* @private
* @param {any} txt
* @param {any} index
* @returns
*/
function checkSensitiveWord(sensitiveMap, txt, index) {
    let currentMap = sensitiveMap;
    let flag = false;
    let wordNum = 0;//记录过滤
    let tempNum = 0;
    for (let i = index; i < txt.length; i++) {
        const word = txt.charAt(i);
        currentMap = currentMap.get(word);
        if (currentMap) {
            tempNum++;
            if (currentMap.get('laster') === true) {
                // 表示已到词的结尾
                flag = true;
                wordNum = tempNum
            }
        } else {
            break;
        }
    }
    return { flag, wordNum };
}


//返回需要替换的起止位置数组
function getSensitiveWordArrayIndex(sensitiveMap, txt) {
    let senIndexArray = []
    for (let i = 0; i < txt.length; i++) {
        let rst = checkSensitiveWord(sensitiveMap, txt, i)
        if (rst && rst.flag) {
            senIndexArray.push({ start: i, end: i + rst.wordNum })
        }
    }

    let rst = []
    if (senIndexArray.length == 0) {
        return rst
    }
    let map = new Map()
    for (let j = 0, len = senIndexArray.length; j < len; j++) {
        let item = senIndexArray[j]
        if (item) {
            for (let x = item.start; x < item.end; x++) {
                map.set(x, true)
            }
        }
    }
    let startIndex = 0
    let isStart = false
    for (let rsi = 0; rsi < txt.length; rsi++) {
        if (map.get(rsi)) {
            if (!isStart) {
                startIndex = rsi
                isStart = true
                continue;
            }
            if (rsi == txt.length - 1) {
                rst.push({ start: startIndex, end: rsi + 1 })
            }
        } else {
            if (isStart) {
                rst.push({ start: startIndex, end: rsi })
            }
            isStart = false
        }
    }

    return rst
}
