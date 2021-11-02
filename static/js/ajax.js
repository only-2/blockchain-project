function ajaxget(url, fnSucc, fnField) {
    if (window.XMLHttpRequest) {
        var oAjax = new XMLHttpRequest();
    }
    else {
        var oAjax = new ActiveXObject("Microsoft.XMLHTTP");
    }
    oAjax.open("GET", url, true);
    oAjax.send();
    oAjax.onreadystatechange = function () {
        if (oAjax.readyState == 4) {
            if (oAjax.status == 200) {
                fnSucc(oAjax.responseText);
            }
            else {
                if (fnfiled) {
                    fnField(oAjax.status);
                }
            }
        }
    };
}
function ajaxpost(url, obj, fnSucc1) {
    var httpRequest = new XMLHttpRequest();
    httpRequest.open('POST', url, true);
    httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    httpRequest.send(JSON.stringify(obj));

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            fnSucc1(httpRequest.responseText);
        }
    };
}
function ajaxjson(url, obj, fnSucc2) {
    var httpRequest = new XMLHttpRequest();
    httpRequest.open('POST', url, true);
    httpRequest.setRequestHeader("Content-type", "application/json"); var obj = { name: 'aman', age: 20 };
    httpRequest.send(JSON.stringify(obj));

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            var json = httpRequest.responseText;
            console.log(json);
            fnSucc2(httpRequest.responseText);
        }
    };
}