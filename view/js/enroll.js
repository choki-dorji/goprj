function getStudents(data) {
    const students = JSON.parse(data);
    var select = document.getElementById("sid");
    students.forEach(stud => {
        let option = document.createElement("option");
        option.textContent = stud.stdid;
        option.value = stud.stdid;
        select.appendChild(option);
    });
}

function getCourses(data) {
    const courses = JSON.parse(data);
    var select = document.getElementById("cid");
    courses.forEach(course => {
        let option = document.createElement("option");
        option.textContent = course.cid;
        option.value = course.cid;
        select.appendChild(option);
    });
}

function addEnroll() {
    const stdid = parseInt(document.getElementById("sid").value);
    const cid = document.getElementById("cid").value;
    fetch('/enroll', {
        method: "POST",
        body: JSON.stringify({ stdid: stdid, cid: cid }),
        headers: { "Content-type": "application/json; charset=UTF-8" }
    }).then(response => {
        if (response.ok) {
            fetch(`/enroll/${stdid}/${cid}`)
                .then(res => res.text())
                .then(data => getEnrolled(data));
        } else {
            throw new Error(response.statusText);
        }
    }).catch(e => {
        if (e == "Error: Forbidden") {
            alert("Duplicate entry!");
        }
    });
    resetFields();
}

function getEnrolled(data) {
    const enrolled = JSON.parse(data);
    showTable(enrolled);
}

function showTable(enrolled) {
    const table = document.getElementById("myTable");
    const row = table.insertRow(-1);
    row.insertCell(0).innerText = enrolled.stdid;
    row.insertCell(1).innerText = enrolled.cid;
    row.insertCell(2).innerText = enrolled.date.split("T")[0];
    row.insertCell(3).innerHTML = `<input type="button" onclick="deleteEnroll(this)" value="Delete">`;
}

function resetFields() {
    document.getElementById("sid").value = "";
    document.getElementById("cid").value = "";
}

window.onload = () => {
    fetch('/student').then(res => res.text()).then(data => getStudents(data));
    fetch('/courses').then(res => res.text()).then(data => getCourses(data));
    fetch('/enrolls').then(res => res.text()).then(data => {
        const all = JSON.parse(data);
        all.forEach(e => showTable(e));
    });
};
