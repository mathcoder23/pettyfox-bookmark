let fs = require('fs');
let path = require('path');
let readline = require('readline');

let filepath = 'web-version-record.md'
let versionPath = path.join('..','src','version.js')


let dateFormat = (date)=>{
    return format(date,'YYYY-mm-dd')
}

function readFileToArr(fReadName,callback){
    var fRead = fs.createReadStream(fReadName);
    var objReadline = readline.createInterface({
        input:fRead
    });
    var arr = new Array();
    objReadline.on('line',function (line) {
        arr.push(line);
        //console.log('line:'+ line);
    });
    objReadline.on('close',function () {
        // console.log(arr);
        callback(arr);
    });
}
readFileToArr(filepath,(arr)=>{
    let date = dateFormat(new Date())
    let version = "1.0.0"
    let v = version.split('.')
    if(arr[0]){
        version = arr[0].split(' ')[2]
        v = version.split('.')
        v[2] = parseInt(v[2])+1
    }

    let cv = v.join('.')
    let newText = [
        '# 版本 '+cv,
        '### 日期 ' +date,
        '### 特性',
        '',
        '',
        '---',
        ''
    ]
    fs.writeFile(filepath, newText.join('\n')+"\n"+arr.join('\n'), function (error) {
        if (error) {
            console.log('写入失败')
        } else {
            console.log('新建版本号:'+cv)
        }
    })
})
let format =  (date,fmt) =>{
    let ret;
    let opt = {
        "Y+": date.getFullYear().toString(),        // 年
        "m+": (date.getMonth() + 1).toString(),     // 月
        "d+": date.getDate().toString(),            // 日
        "H+": date.getHours().toString(),           // 时
        "M+": date.getMinutes().toString(),         // 分
        "S+": date.getSeconds().toString()          // 秒
        // 有其他格式化字符需求可以继续添加，必须转化成字符串
    };
    for (let k in opt) {
        ret = new RegExp("(" + k + ")").exec(fmt);
        if (ret) {
            fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
        };
    };
    return fmt;
}
