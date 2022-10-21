yes = document.getElementById("yes-container");
enroute = document.getElementById("en-route-container");
no = document.getElementById("no-container");
solo = document.getElementById("solo-container");
soloText = document.getElementById("solo-text");

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

                solo.style.backgroundColor = "#4cff4c";
                soloText.textContent = "Yes";
                for (const child of solo.children) {
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

                solo.style.backgroundColor = "#ffff56";
                soloText.textContent = "En Route";
                for (const child of solo.children) {
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

                solo.style.backgroundColor = "#ff3434";
                soloText.textContent = "No";
                for (const child of solo.children) {
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