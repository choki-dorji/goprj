let selectedRow = null;

function getFormData() {
    return {
        cid: document.getElementById("cid").value,
        coursename: document.getElementById("cname").value
    };
}

function resetform() {
    document.getElementById("cid").value = "";
    document.getElementById("cname").value = "";
}

function showCourse(data) {
    const course = JSON.parse(data);
    newRow(course);
}

function newRow(course) {
    const table = document.getElementById("myTable").getElementsByTagName("tbody")[0];
    const row = table.insertRow();

    row.innerHTML = `
      <td class="px-6 py-4 whitespace-nowrap">${course.cid}</td>
      <td class="px-6 py-4 whitespace-nowrap">${course.coursename}</td>
      <td class="px-6 py-4 text-center">
        <button onclick="updateCourse(this)" class="text-blue-600 hover:text-blue-800 font-medium underline">Edit</button>
      </td>
      <td class="px-6 py-4 text-center">
        <button onclick="deleteCourse(this)" class="text-gray-500 hover:text-red-600 font-medium">Delete</button>
      </td>
    `;
}


function addCourse() {
    let data = getFormData();
    fetch("/course", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json"}
    }).then(res => {
        if (res.ok) {
            fetch("/course/" + data.cid)
                .then(response => response.text())
                .then(data => showCourse(data));
        } else {
            alert("Course add failed.");
        }
    });
    resetform();
}

function deleteCourse(r) {
    if (confirm("Are you sure?")) {
        selectedRow = r.parentElement.parentElement;
        let cid = selectedRow.cells[0].innerHTML;

        fetch("/course/" + cid, {
            method: "DELETE",
            headers: {"Content-type": "application/json"}
        });

        document.getElementById("myTable").deleteRow(selectedRow.rowIndex);
    }
}

function updateCourse(r) {
    selectedRow = r.parentElement.parentElement;
    document.getElementById("cid").value = selectedRow.cells[0].innerHTML;
    document.getElementById("cname").value = selectedRow.cells[1].innerHTML;

    var btn = document.getElementById("button-add");
    btn.innerHTML = "Update";
    btn.setAttribute("onclick", "update()");
}

function update() {
    var newData = getFormData();
    fetch("/course/" + newData.cid, {
        method: "PUT",
        body: JSON.stringify(newData),
        headers: {"Content-type": "application/json"}
    }).then(res => {
        if (res.ok) {
            selectedRow.cells[0].innerHTML = newData.cid;
            selectedRow.cells[1].innerHTML = newData.coursename;

            var btn = document.getElementById("button-add");
            btn.innerHTML = "Add";
            btn.setAttribute("onclick", "addCourse()");
            selectedRow = null;
            resetform();
        } else {
            alert("Update failed");
        }
    });
}

function showCourses(data) {
    const courses = JSON.parse(data);
    console.log("course ", courses);
    courses.forEach(course => newRow(course));
}

window.onload = () => {
    fetch("/courses")
        .then(response => response.text())
        .then(data => showCourses(data));
};
