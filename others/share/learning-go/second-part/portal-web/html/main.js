window.addEventListener("load", function (evt) {
    var uri = "http://localhost:8088/parallel-sum"
    var output = document.getElementById("output");

    var print = function (message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };


    document.getElementById("send").onclick = function (evt) {
        if (!ws) {
            return false
        }

        var msg = {hi: nameTxt.value}

        req = JSON.stringify(msg)
        print('<span style="color: blue;">Sent request: </span>' + req);
        ws.send(JSON.stringify(msg));

        return false;
    };
})
