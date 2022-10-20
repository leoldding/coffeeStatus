yes = document.getElementById("yes");
enroute = document.getElementById("enroute");
no = document.getElementById("no");

window.onload = function() {
    fetch("backend/loadstatus", {
        method: "GET",
    }).then((response) => {
        response.text().then(function (data) {
            status = JSON.parse(data).status;
            if (status == "yes") {
                yes.style.backgroundColor = "#4cff4c";
            } else {
                yes.style.backgroundColor = "#4f4f4fff"
            }
            if (status == "enroute") {
                enroute.style.backgroundColor = "#ffff56";
            } else {
                enroute.style.backgroundColor = "#4f4f4fff";
            }
            if (status == "no") {
                no.style.backgroundColor = "#ff3434";
            } else {
                no.style.backgroundColor = "#4f4f4fff"
            }
        });
    }).catch((error)=>{
        console.log(error)
    });
};