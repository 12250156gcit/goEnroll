fetch("/check-auth", { method: "GET", credentials: "include" })
    .then(response => {
        if (!response.ok) {
            alert("User not logged in !!")
            window.open("index.html", "_self")
        }
    })
    .catch(() => {
        alert("User not logged in !!")
        window.open("index.html", "_self")
    })