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