function signUp(){
    //retrive from data
    var _data = {
        firstname: document.getElementById("fname").value,
        lastname: document.getElementById("lname").value, 
        email: document.getElementById("email").value, 
        password: document.getElementById("pw1").value,
        pw: document.getElementById("pw2").value 
    }
    if(_data.password !== _data.pw){
        alert("Password doesn't match")
        return
    }
    fetch("/signup",{
        method: "POST",
        body: JSON.stringify(_data),
        headers: {"Content-Type": "application/json; charset=UTF-8"},
        credentials: "include"
    }).then(async response => {
        if(response.status == 201){
            window.open("index.html", "_self")
            return
        }
        const body = await response.text()
        throw new Error(body || response.statusText)
    }).catch(e => alert(e));
}