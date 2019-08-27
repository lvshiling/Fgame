

export function uuid(len, radix) {
    var chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'.split('');
    var uuid = [], i;
    radix = radix || chars.length;

    if (len) {
        // Compact form
        for (i = 0; i < len; i++) uuid[i] = chars[0 | Math.random() * radix];
    } else {
        // rfc4122, version 4 form
        var r;

        // rfc4122 requires these characters
        uuid[8] = uuid[13] = uuid[18] = uuid[23] = '-';
        uuid[14] = '4';

        // Fill in random data.  At i==19 set the high bits of clock sequence as
        // per rfc4122, sec. 4.1.5
        for (i = 0; i < 36; i++) {
            if (!uuid[i]) {
                r = 0 | Math.random() * 16;
                uuid[i] = chars[(i == 19) ? (r & 0x3) | 0x8 : r];
            }
        }
    }

    return uuid.join('');
}

export function newUUID() {
    return uuid(32, 16)
}

//校验物品是否正确
export function checkItemContent(content) {
    if (content===undefined ||content == '' ) {
        return true;
    }
    let itemArray = content.split(',');
    for (let i = 0; i < itemArray.length; i++) {
        let item = itemArray[i];
        let skuArray = item.split(':');
        if (skuArray.length != 2) {
            return false;
        }
        for (let j = 0; j < 2; j++) {
            let value = skuArray[j];
            if(!checkRate(value)){
                return false;
            }
        }
    }

    return true;
}



/**
 * 判断字符串是否为数字
 * @param nubmer
 * @returns {boolean}
 */
function checkRate(nubmer) {
    //判断正整数/[1−9]+[0−9]∗]∗/
    var re = /^[0-9]+.?[0-9]*/;//
    if (!re.test(nubmer)) {
        return false;
    }
    return true;
}