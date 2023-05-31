function validateForm(email, pw) {
    if (email == "" || pw == "" ) {
        window.alert("输入不可留空！")
        return false
    }else {
        return true
    }
    // 多加一个非法字符检测
}

function login(form) {
    formData = new FormData(form);
    let email = formData.get("email")
    let pw = formData.get("pw")
    let xsrf = formData.get("_xsrf")

    if (validateForm(email, pw) == true) {
        var xmlhttp;
        if (window.XMLHttpRequest) {
            xmlhttp = new XMLHttpRequest();
        }
        else {
            xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
        }
        data = "_xsrf="+xsrf+"&email="+email+"&pw="+pw
        xmlhttp.open("POST", "/login", true);
        xmlhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xmlhttp.send(data);


        xmlhttp.onreadystatechange = function () {
            if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                result = xmlhttp.responseText
                if (result == "OK") {
                    jumpAfter5()
                }else {
                    document.getElementById("loginResult").innerHTML = result;
                }
            }
        }

    }

}

function jumpAfter5() {
    var signupResult = document.getElementById("loginResult");
    var time = 5;
    timer();
    setInterval(timer, 1000)

    function timer() {
        if (time == 0) {
            location.href = '/msgBox'

        } else {
            signupResult.innerHTML = "登录成功！将在" + time + '秒之后跳转到用户主页面...';
            time--;
        }
    }
}