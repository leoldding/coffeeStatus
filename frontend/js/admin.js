logout = document.getElementById("logout");
submit = document.getElementById("submit");

/* check to see if cookie exists */
window.onload = function() {
    fetch("../backend/checkCookie", {
        method: "GET",
    }).then((response) => {
        if (response.status == 401) {
            /* redirect to login page if cookie is not valid */
            window.location = "../html/login.html";
        }
    }).catch((error) => {
        console.log(error)
    })
}

logout.addEventListener('click', function() {
    fetch("../backend/logout", {
        method: "GET",
    }).then((response) => {
        /* redirect to main page */
        window.location = "../index.html";
    }).catch((error) => {
        console.log(error)
    })
})

submit.addEventListener('click', function() {
    /* retrieve values from page */
    status = document.getElementById("statuses").value;
    substatus = document.getElementById("substatus").value;

    /* adjust value if no substatus exists */
    if (substatus == "") {
        substatus = "none";
    }

    /* reset text box */
    document.getElementById("substatus").value = "";

    if (substatus.length > 30) {
        alert("Message is too long.")
    } else {
        let update = {
            "status": status,
            "substatus": substatus,
        }

        /* update values in backend */
        fetch("../backend/updateStatus", {
            method: "POST",
            body: JSON.stringify(update)
        }).catch((error) => {
            console.log(error)
        })
    }
})