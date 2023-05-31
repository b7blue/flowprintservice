function validateForm(uid, pw, pw1) {

    if (uid == "" || pw == "" || pw1 == "") {
        return "输入不可留空！"
    } else if (pw1 != pw) {
        return "两次输入的密码不一致，请重新确认！"
    } else if (uid.length>15){
        return "用户名长度非法"
    }else if (!isNaN(parseFloat(uid)) && isFinite(uid)){
        return "用户名不能是纯数字"
    }else {
        return "ok"
    }
    // 多加一个非法字符检测
}

function register(form) {
    formData = new FormData(form);
    let uid = formData.get("uid")
    let pw = formData.get("pw")
    let pw1 = formData.get("pw1")
    let xsrf = formData.get("_xsrf")
    validresult = validateForm(uid, pw, pw1)
    if (validresult == "ok") {
        var xmlhttp;
        if (window.XMLHttpRequest) {
            xmlhttp = new XMLHttpRequest();
        }
        else {
            xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
        }
        data = "_xsrf=" + xsrf + "&uid=" + uid + "&pw=" + pw 
        xmlhttp.open("POST", "/signup", true);
        xmlhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xmlhttp.send(data);


        xmlhttp.onreadystatechange = function () {
            if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                result = xmlhttp.responseText
                if (result == "OK") {
                    jumpAfter5()
                } else {
                    registerfail(result)
                }
            }
        }

    }else {
        registerfail(validresult)
    }

}

function registerfail(failreason) {
    document.getElementById("warningcard").style.display = "flex"
    document.getElementById("warningcontent").innerHTML = failreason
}

function jumpAfter5() {
    document.getElementById("okcard").style.display = "flex"
    signupResult = document.getElementById("okcontent")
    var time = 5;
    timer();
    setInterval(timer, 1000)

    function timer() {
        if (time == 0) {
            location.href = '/login'

        } else {
            signupResult.innerHTML = "将在" + time + '秒之后跳转到登陆页面';
            time--;
        }
    }
}


function getVariCode() {
    var xmlhttp;
    if (window.XMLHttpRequest) {
        // IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
        xmlhttp = new XMLHttpRequest();
    }
    else {
        // IE6, IE5 浏览器执行代码
        xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
    }
    var param = "_xsrf=" + document.forms["myForm"]["_xsrf"].value + "&email=" + document.forms["myForm"]["email"].value
    xmlhttp.open("POST", "/sendVariCode", true);
    xmlhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xmlhttp.send(param);


    xmlhttp.onreadystatechange = function () {
        if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
            window.alert(xmlhttp.responseText)
            // document.getElementById("getVariCodeResult").innerHTML = xmlhttp.responseText;
        }
    }

}

// function to_stu_info() {

//     document.getElementById("base_info_form").id = "base_info_form_folded"
//     document.getElementById("stu_info_form_folded").id = "stu_info_form"
// }
