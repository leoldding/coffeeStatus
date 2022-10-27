logout = document.getElementById("logout");

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
        window.location = "../index.html";
    }).catch((error) => {
        console.log(error)
    })
})