function edit(obj) {
    var thisid = obj.id.toString().substring(4);
    var ipobj = document.getElementById("ip" + thisid);
    var portobj = document.getElementById("port" + thisid);
    var oldip = ipobj.innerHTML
    var oldport = portobj.innerHTML

    // 变成文本框
    if (ipobj.childNodes[0].value) {
        ipobj.innerHTML = ipobj.childNodes[0].value;
    } else {
        ipobj.innerHTML = "<input type='text' name='ip' value='" + oldip + "' />";
    }
    if (portobj.childNodes[0].value) {
        portobj.innerHTML = o.childNodes[0].value;
    } else {
        portobj.innerHTML = "<input type='text' name='port' value='" + oldport + "' />";
    }
    // 为编辑完成按钮添加onclick事件
    document.getElementById("done" + thisid).setAttribute("onclick", "editdone(this,form" + thisid + ")")
    // 编辑状态下无法删除指纹，将删除按钮onclick事件删除
    document.getElementById("del" + thisid).onclick = null
}

// 编辑指纹，提交表单前的准备
function editdone(obj) {
    var thisid = obj.id.toString().substring(4);
    form = document.getElementById("form"+thisid)
    formData = new FormData(form);
    let port = formData.get("port")
    let ip = formData.get("ip")

    reg = /^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[0-9])\.((1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.){2}(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)$/
    // 判断ip是否合法
    if (reg.test(ip)) {
        // 判断端口号是否为数字
        if (!isNaN(parseFloat(port)) && isFinite(port)) {
            // 判断端口号是否在合法范围
            if (parseInt(port) > 0 && parseInt(port) < 65536) {
                // 将id = op thisid 的input的value变成edit
                document.getElementById("op" + thisid).value = "edit"
                // 提交表单
                form.submit()
            }
        }
        errorbut = document.getElementById("error" + thisid)
        errorbut.innerText = "端口非法"+errorbut.innerText
        errorbut.style.display = "inline"
    } else {
        errorbut = document.getElementById("error" + thisid)
        errorbut.innerText = "ip非法"+errorbut.innerText
        errorbut.style.display = "inline"
    }

}

// 删除指纹，提交表单前的准备
function del(obj) {
    var thisid = obj.id.toString().substring(3)
    form = document.getElementById("form"+thisid)
    document.getElementById("op" + thisid).value = "del"
    // 提交表单
    form.submit()
}

