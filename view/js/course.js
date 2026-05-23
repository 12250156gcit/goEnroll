window.onload = function () {
  fetch("/course")
    .then(response => {
      if (!response.ok) throw new Error(response.statusText);
      return response.text();
    })
    .then(data => showCourses(data))
    .catch(e => handleCourseError(e));
};

var selectedRow = null;
var cid = null;

function handleCourseError(e) {
  if (String(e) === "Error: Unauthorized") {
    alert("User not logged in!");
    window.open("index.html", "_self");
  } else {
    alert(e);
  }
}

function getCourseFormData() {
  return {
    cid: document.getElementById("cid").value.trim(),
    coursename: document.getElementById("cname").value.trim()
  };
}

function resetCourseForm() {
  document.getElementById("cid").value = "";
  document.getElementById("cname").value = "";
}

function newCourseRow(table, course) {
  var row = table.insertRow(table.length);
  var td = [];
  for (var i = 0; i < table.rows[0].cells.length; i++) td[i] = row.insertCell(i);
  td[0].innerHTML = course.cid;
  td[1].innerHTML = course.coursename;
  td[2].innerHTML = '<input type="button" onclick="deleteCourse(this)" value="delete" id="button-1" />';
  td[3].innerHTML = '<input type="button" onclick="updateCourse(this)" value="edit" id="button-2" />';
}

function showCourses(data) {
  var courses = JSON.parse(data || "[]");
  var table = document.getElementById("myTable");
  courses.forEach(course => newCourseRow(table, course));
}

function showCourse(data) {
  var course = JSON.parse(data);
  newCourseRow(document.getElementById("myTable"), course);
}

function addCourse() {
  var data = getCourseFormData();
  if (data.cid === "" || data.coursename === "") {
    alert("Course ID and course name are required");
    return;
  }
  fetch("/course", {
    method: "POST",
    body: JSON.stringify(data),
    headers: {"Content-type": "application/json; charset=UTF-8"}
  }).then(response => {
    if (response.ok) {
      return fetch("/course/" + encodeURIComponent(data.cid))
        .then(response => response.text())
        .then(data => showCourse(data));
    }
    throw new Error(response.statusText);
  }).then(() => resetCourseForm())
    .catch(e => handleCourseError(e));
}

function updateCourse(input) {
  selectedRow = input.parentElement.parentElement;
  document.getElementById("cid").value = selectedRow.cells[0].innerHTML;
  document.getElementById("cname").value = selectedRow.cells[1].innerHTML;
  cid = selectedRow.cells[0].innerHTML;
  var btn = document.getElementById("button-add");
  btn.innerHTML = "Update";
  btn.setAttribute("onclick", "updateCourseAPI(cid)");
}

function updateCourseAPI(oldCid) {
  var data = getCourseFormData();
  fetch("/course/" + encodeURIComponent(oldCid), {
    method: "PUT",
    body: JSON.stringify(data),
    headers: {"Content-type": "application/json; charset=UTF-8"}
  }).then(response => {
    if (response.ok) {
      selectedRow.cells[0].innerHTML = data.cid;
      selectedRow.cells[1].innerHTML = data.coursename;
      var btn = document.getElementById("button-add");
      btn.innerHTML = "Add";
      btn.setAttribute("onclick", "addCourse()");
      selectedRow = null;
      resetCourseForm();
    } else {
      throw new Error(response.statusText);
    }
  }).catch(e => handleCourseError(e));
}

function deleteCourse(input) {
  if (confirm("Are you sure you want to DELETE this?")) {
    selectedRow = input.parentElement.parentElement;
    cid = selectedRow.cells[0].innerHTML;
    fetch("/course/" + encodeURIComponent(cid), { method: "DELETE" })
      .then(response => {
        if (response.ok) {
          var rowIndex = selectedRow.rowIndex;
          if (rowIndex > 0) document.getElementById("myTable").deleteRow(rowIndex);
          selectedRow = null;
        } else {
          throw new Error(response.statusText);
        }
      }).catch(e => handleCourseError(e));
  }
}