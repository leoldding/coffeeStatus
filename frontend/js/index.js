/* retrieving elements that have changes */
yes = document.getElementById("yes-container");
yesSub = document.getElementById("yes-substatus");
enroute = document.getElementById("en-route-container");
enrouteSub = document.getElementById("en-route-substatus");
no = document.getElementById("no-container");
noSub = document.getElementById("no-substatus");
solo = document.getElementById("solo-container");
soloText = document.getElementById("solo-text");
soloSub = document.getElementById("solo-substatus");

/* update element values when website loads in  */
window.onload = function() {
    /* retrieve current status from backend */
    fetch("backend/loadStatus", {
        method: "GET",
    }).then((response) => {
        response.text().then(function (data) {
            if (response.status == 200) {
                status = JSON.parse(data).status; /* status value from backend */
                substatus = JSON.parse(data).substatus; /* substatus value from backend */
            } else {
                /* default values in case of server error */
                status = "no"
                substatus = "none"
            }

            /* changes for the "Yes" and solo containers based on status value */
            if (status == "yes") {
                yes.style.backgroundColor = "#4cff4c";
                for (const child of yes.children) {
                    child.style.color = "black";
                }
                if (substatus != "none") {
                    yesSub.textContent = substatus;
                }

                solo.style.backgroundColor = "#4cff4c";
                soloText.textContent = "Yes";
                for (const child of solo.children) {
                    child.style.color = "black";
                }
                if (substatus != "none") {
                    soloSub.textContent = substatus;
                } else {
                    soloSub.textContent = "";
                }
            } else {
                yes.style.backgroundColor = "gray";
                for (const child of yes.children) {
                    child.style.color = "silver";
                }
                yesSub.textContent = "";
            }

            /* changes for the "En Route" and solo containers based on status value */
            if (status == "enroute") {
                enroute.style.backgroundColor = "#ffff56";
                for (const child of enroute.children) {
                    child.style.color = "black";
                }
                if (substatus != "none") {
                    enrouteSub.textContent = substatus;
                }

                solo.style.backgroundColor = "#ffff56";
                soloText.textContent = "En Route";
                for (const child of solo.children) {
                    child.style.color = "black";
                }
                if (substatus != "none") {
                    soloSub.textContent = substatus;
                } else {
                    soloSub.textContent = "";
                }
            } else {
                enroute.style.backgroundColor = "gray";
                for (const child of enroute.children) {
                    child.style.color = "silver";
                }
                enrouteSub.textContent = "";
            }

            /* changes for the "No" and solo containers based on status value */
            if (status == "no") {
                no.style.backgroundColor = "#ff3434";
                for (const child of no.children) {
                    child.style.color = "black";
                }
                if (substatus != "none") {
                    noSub.textContent = substatus
                }

                solo.style.backgroundColor = "#ff3434";
                soloText.textContent = "No";
                for (const child of solo.children) {
                    child.style.color = "black";
                }
                if (substatus != "none") {
                    soloSub.textContent = substatus
                } else {
                    soloSub.textContent = "";
                }
            } else {
                no.style.backgroundColor = "gray";
                for (const child of no.children) {
                    child.style.color = "silver";
                }
                noSub.textContent = "";
            }
        });
    }).catch((error)=>{
        console.log(error)
    });
};