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
    var table = document.getElementById("myTable");
    var row = table.insertRow(table.length);
    var td = [];

    for (let i = 0; i < table.rows[0].cells.length; i++) {
        td[i] = row.insertCell(i);
    }

    td[0].innerHTML = course.cid;
    td[1].innerHTML = course.coursename;
    td[2].innerHTML = '<input type="button" onclick="deleteCourse(this)" value="delete">';
    td[3].innerHTML = '<input type="button" onclick="updateCourse(this)" value="edit">';
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
    courses.forEach(course => newRow(course));
}

window.onload = () => {
    fetch("/courses")
        .then(response => response.text())
        .then(data => showCourses(data));
};
