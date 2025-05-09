function login(){
    var _data = {
        email : document.getElementById("email").value,
        password : document.getElementById("pw").value
    }
    fetch("/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(_data)
    })
    .then(response => {
        if (response.ok) {
            window.open("student.html", "_self");
        } else {
            throw new Error(response.statusText);
        }
    }).catch(e => {
        if (e.message == "Unauthorized") {
            alert("Invalid email or password");
            return
        }
    })
}

function logout() {
    fetch('/logout')
    .then(response => {
    if (response.ok) {
    window.open("index.html", "_self")
    } else {
    throw new Error(response.statusText)
    }
    }).catch (e => {
    alert(e)
    })
    }
    