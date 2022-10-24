formUsername = document.getElementById("username");
formPassword = document.getElementById("password");
loginButton = document.getElementById("login-button");

loginButton.addEventListener('click',function()  {

    /* retrieve credentials from input fields */
    let credentials = {
        "username": formUsername.value.trim(),
        "password": formPassword.value.trim(),
    };

    /* clear input fields */
    formUsername.value = "";
    formPassword.value = "";

    /* check if either input is empty */
    if (credentials.username == "" || credentials.password == "") {
        alert("Username or password is empty!");
        return
    }

    /* check if special characters are in either input */
    if (containsSpecialChars(credentials.username) || containsSpecialChars(credentials.password)) {
        alert("Username or password contains illegal characters!");
        return
    }

    /* attempt to login with inputs */
    fetch("../backend/login", {
        method: "POST",
        body: JSON.stringify(credentials)
    }).then((response) => {
        response.text().then(function (data) {
            /* parse status from backend */
            status = JSON.parse(data).status
            if (status == "Valid Credentials") {
                /* redirect if validation is successful */
                window.location = "../html/admin.html";
            }
            else {
                /* display validation error */
                alert(status);
            }
        });
    }).catch((error) => {
        console.log(error)
    });
})

function containsSpecialChars(str) {
    const specialChars = /[`!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]/;
    return specialChars.test(str);
}