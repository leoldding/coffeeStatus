/* retrieving elements that have changes */
yes = document.getElementById("yes-container");
enroute = document.getElementById("en-route-container");
no = document.getElementById("no-container");
solo = document.getElementById("solo-container");
soloText = document.getElementById("solo-text");

/* update element values when website loads in  */
window.onload = function() {
    /* retrieve current status from backend */
    fetch("backend/loadStatus", {
        method: "GET",
    }).then((response) => {
        response.text().then(function (data) {
            status = JSON.parse(data).status; /* status value from backend */

            /* changes for the "Yes" and solo containers based on status value */
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
                    child.style.color = "silver";
                }
            }

            /* changes for the "En Route" and solo containers based on status value */
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
                    child.style.color = "silver";
                }
            }

            /* changes for the "No" and solo containers based on status value */
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
                    child.style.color = "silver";
                }
            }
        });
    }).catch((error)=>{
        console.log(error)
    });
};