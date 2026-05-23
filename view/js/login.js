function login(){
    var data={
        email: document.getElementById("email").value,
        password: document.getElementById("pw").value
    }

    fetch("/login", {
         method: "POST",
        body: JSON.stringify(data),
        headers: {"content-Type": "application/json; charset=UTF-8"},
        credentials: "include"
    }).then(response => {
        if(response.ok){
            window.open("student.html", "_self")
        } else {
            throw new Error(response.statusText)
        }
    }).catch(e => alert(e))
}

function logout(){
    fetch("/logout", {credentials: "include"})
    .then(response => {
        if(response.ok) {
            window.open("index.html", "_self")
        } else{
            throw new Error(response.statusText)
        }
    }).catch(e => alert(e))
}