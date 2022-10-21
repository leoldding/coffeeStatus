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
                for (const child of yes.children) {
                    child.style.color = "black";
                }
            } else {
                yes.style.backgroundColor = "gray";
                for (const child of yes.children) {
                    child.style.color = "white";
                }
            }

            if (status == "enroute") {
                enroute.style.backgroundColor = "#ffff56";
                for (const child of enroute.children) {
                    child.style.color = "black";
                }
            } else {
                enroute.style.backgroundColor = "gray";
                for (const child of enroute.children) {
                    child.style.color = "white";
                }
            }

            if (status == "no") {
                no.style.backgroundColor = "#ff3434";
                for (const child of no.children) {
                    child.style.color = "black";
                }
            } else {
                no.style.backgroundColor = "gray";
                for (const child of no.children) {
                    child.style.color = "white";
                }
            }
        });
    }).catch((error)=>{
        console.log(error)
    });
};