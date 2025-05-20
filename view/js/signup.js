function signUp() {
    var _data = {
        firstname : document.getElementById("fname").value,
        lastname : document.getElementById("lname").value,
        email : document.getElementById("email").value,
        password : document.getElementById("pw1").value,
        password2 : document.getElementById("pw2").value
    }
    if (_data.password != _data.password2) {
        alert("Passwords do not match");
        return;
    }
    fetch("/signup", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(_data)
    })
    .then(response => {
        if (response.status == 201) {
           window.open("/index.html", "_self");
        }
    })
}