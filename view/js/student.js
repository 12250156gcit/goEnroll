var editingSid = null

window.addEventListener('load', function () {
  loadStudents()
  var btn = document.getElementById('button-add')
  if (btn) {
    btn.onclick = submitStudent
  }
})

function loadStudents() {
  fetch('/student')
    .then(response => response.json())
    .then(data => showStudents(data))
    .catch(err => {
      console.error('Unable to load students:', err)
      alert('Unable to load students. Check the console for details.')
    })
}

function showStudents(students) {
  var table = document.getElementById('myTable')
  students.forEach(stud => {
    newRow(table, stud)
  })
}

// helper function that add data to row
function newRow(table, student) {
  var row = table.insertRow(table.length)
  var td = []
  for (var i = 0; i < table.rows[0].cells.length; i++) {
    td[i] = row.insertCell(i)
  }
  // insert data in the td cells
  td[0].innerHTML = student.stdid
  td[1].innerHTML = student.fname
  td[2].innerHTML = student.lname
  td[3].innerHTML = student.email

  var deleteBtn = document.createElement('input')
  deleteBtn.type = 'button'
  deleteBtn.value = 'delete'
  deleteBtn.addEventListener('click', function () {
    deleteStudent(student.stdid, row)
  })
  td[4].appendChild(deleteBtn)

  var editBtn = document.createElement('input')
  editBtn.type = 'button'
  editBtn.value = 'edit'
  editBtn.addEventListener('click', function () {
    updateStudent(row)
  })
  td[5].appendChild(editBtn)
}

function showStudent(data) {
  const student = JSON.parse(data)
  var table = document.getElementById('myTable')
  newRow(table, student)
}

function deleteStudent(sid, selectedRow) {
  fetch('/student/' + sid, {
    method: 'DELETE'
  })
    .then(response => {
      if (response.ok) {
        selectedRow.remove()
      } else {
        return response.text().then(text => {
          throw new Error(text || response.statusText)
        })
      }
    })
    .catch(err => {
      console.error('Delete failed:', err)
      alert('Unable to delete student: ' + err.message)
    })
}

function UpdateApiRequest(sid) {
  var data = {
    stdid: parseInt(document.getElementById('sid').value),
    fname: document.getElementById('fname').value,
    lname: document.getElementById('lname').value,
    email: document.getElementById('email').value
  }

  if (isNaN(data.stdid)) {
    alert('enter valid std ID')
    return
  }

  fetch('/student/' + sid, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: {'Content-type': 'application/json; charset=UTF-8'}
  })
    .then(response => {
      if (response.ok) {
        window.location.reload()
      } else {
        return response.text().then(text => {
          throw new Error(text || response.statusText)
        })
      }
    })
    .catch(err => {
      console.error('Update failed:', err)
      alert('Unable to update student: ' + err.message)
    })
}

// helper function to reset the field
function resetForm(){
  document.getElementById("sid").value = ""
  document.getElementById("fname").value = ""
  document.getElementById("lname").value = ""
  document.getElementById("email").value = ""
  editingSid = null
  var btn = document.getElementById('button-add')
  if (btn) {
    btn.innerHTML = 'Add'
  }
}
function addStudent () {
  //create a js object form store form data
  var data = {
    stdid : parseInt(document.getElementById("sid").value),
    fname : document.getElementById("fname").value,
    lname : document.getElementById("lname").value,
    email : document.getElementById("email").value
  }
  //form validation
  var sid= data.stdid
  if (isNaN(sid)){
    alert("enter valaid std ID")
    return
  } else if (data.email == ""){
     alert("enter valid email")
     return
  } // fname

  //call post API
  fetch('/student/add', {
    method: "POST",
    body: JSON.stringify(data),
    headers: {"Content-type": "application/json; charset=UTF-8"}
  }).then(response1 => {
    // check the response form fetch is resolved or rejected
    if (response1.ok) {
      fetch('/student/'+ data.stdid)
      .then(response2 => response2.text())
      .then(data => showStudent(data))
    } else {
      throw new Error(response1.statusText)
    }
  }).catch(e => {
    if (e.message == '401') {
      alert('User Not Logged in')
      window.open('index.html', '_self')
    } else if (e.message == '400') {
      alert('Bad Request')
    } else {
      alert('internal server error')
    }
  });
  resetForm();
}

function submitStudent(event) {
  event.preventDefault && event.preventDefault()
  if (editingSid) {
    UpdateApiRequest(editingSid)
  } else {
    addStudent()
  }
}

function updateStudent(input) {
  var selectedRow
  if (input && input.nodeName === 'TR') {
    selectedRow = input
  } else if (input && input.parentElement && input.parentElement.parentElement) {
    selectedRow = input.parentElement.parentElement
  } else {
    console.error('updateStudent: could not determine row from input', input)
    return
  }
  document.getElementById('sid').value = selectedRow.cells[0].innerHTML
  document.getElementById('fname').value = selectedRow.cells[1].innerHTML
  document.getElementById('lname').value = selectedRow.cells[2].innerHTML
  document.getElementById('email').value = selectedRow.cells[3].innerHTML
  editingSid = selectedRow.cells[0].innerHTML

  // change button value to update
  var btn = document.getElementById('button-add')
  btn.innerHTML = 'Update'
}