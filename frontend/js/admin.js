window.onload = function() {
    fetch("../backend/checkCookie", {
        method: "GET",
    }).then((response) => {
        if (response.status == 401) {
            window.location = "../html/login.html";
        }
    }).catch((error) => {
        console.log(error)
    })
}