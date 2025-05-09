window.onload = function () {
  fetch("/student")
    .then(response => response.json())
    .then(data => showStudents(data));
};

function addStudent() {
  const data = getFormData();
  if (isNaN(data.stdid) || data.email === "" || data.fname === "") {
    alert("Please fill all required fields correctly.");
    return;
  }

  fetch("/student", {
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  })
    .then(response => response.ok ? fetch("/student") : Promise.reject("Add failed"))
    .then(res => res.json())
    .then(data => {
      showStudents(data);
      resetform();
    })
    .catch(e => alert("Error: " + e));
}

function showStudents(students) {
  const table = document.getElementById("myTable").getElementsByTagName("tbody")[0];
  table.innerHTML = "";
  students.forEach(student => newRow(student));
}

function newRow(student) {
  const table = document.getElementById("myTable").getElementsByTagName("tbody")[0];
  const row = table.insertRow();
  row.innerHTML = `
    <td class="px-6 py-4 whitespace-nowrap">${student.stdid}</td>
    <td class="px-6 py-4 whitespace-nowrap">${student.fname}</td>
    <td class="px-6 py-4 whitespace-nowrap">${student.lname}</td>
    <td class="px-6 py-4 whitespace-nowrap">${student.email}</td>
    <td class="px-6 py-4 text-center">
      <button onclick="updateStudent(this)" class="text-blue-600 hover:text-blue-800 font-medium underline">Edit</button>
    </td>
    <td class="px-6 py-4 text-center">
      <button onclick="deleteStudent(this)" class="text-gray-500 hover:text-red-600 font-medium">Delete</button>
    </td>
  `;
}

function resetform() {
  document.getElementById("sid").value = "";
  document.getElementById("fname").value = "";
  document.getElementById("lname").value = "";
  document.getElementById("email").value = "";
}

function getFormData() {
  return {
    stdid: parseInt(document.getElementById("sid").value),
    fname: document.getElementById("fname").value,
    lname: document.getElementById("lname").value,
    email: document.getElementById("email").value
  };
}

let selectedRow = null;

function updateStudent(button) {
  selectedRow = button.parentElement.parentElement;
  document.getElementById("sid").value = selectedRow.cells[0].innerText;
  document.getElementById("fname").value = selectedRow.cells[1].innerText;
  document.getElementById("lname").value = selectedRow.cells[2].innerText;
  document.getElementById("email").value = selectedRow.cells[3].innerText;

  const btn = document.getElementById("button-add");
  btn.innerText = "Update";
  btn.setAttribute("onclick", "update()");
}

function update() {
  const data = getFormData();
  const sid = selectedRow.cells[0].innerText;

  fetch("/student/" + sid, {
    method: "PUT",
    body: JSON.stringify(data),
    headers: { "Content-type": "application/json; charset=UTF-8" }
  })
    .then(res => {
      if (res.ok) {
        selectedRow.cells[0].innerText = data.stdid;
        selectedRow.cells[1].innerText = data.fname;
        selectedRow.cells[2].innerText = data.lname;
        selectedRow.cells[3].innerText = data.email;

        const btn = document.getElementById("button-add");
        btn.innerText = "Add";
        btn.setAttribute("onclick", "addStudent()");
        resetform();
        selectedRow = null;
      } else {
        alert("Update failed");
      }
    });
}

function deleteStudent(button) {
  if (confirm("Are you sure you want to delete this student?")) {
    const row = button.parentElement.parentElement;
    const sid = row.cells[0].innerText;

    fetch("/student/" + sid, {
      method: "DELETE",
      headers: { "Content-type": "application/json; charset=UTF-8" }
    }).then(() => row.remove());
  }
}
