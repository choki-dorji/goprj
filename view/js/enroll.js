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
  const tbody = document.getElementById("myTable").getElementsByTagName("tbody")[0];
  const row = tbody.insertRow();

  row.innerHTML = `
    <td class="px-6 py-4 whitespace-nowrap">
      <button onclick="showStudentDetail(${enrolled.stdid})" class="text-indigo-600 hover:underline">${enrolled.stdid}</button>
    </td>
    <td class="px-6 py-4 whitespace-nowrap">
      <button onclick="showCourseDetail('${enrolled.cid}')" class="text-indigo-600 hover:underline">${enrolled.cid}</button>
    </td>
    <td class="px-6 py-4 whitespace-nowrap">${enrolled.date.split("T")[0]}</td>
    <td class="px-6 py-4 text-center">
      <button onclick="deleteEnroll(this)" class="text-gray-500 hover:text-red-600 font-medium">Delete</button>
    </td>
  `;

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


function showStudentDetail(id) {
  fetch(`/student/${id}`)
    .then(res => res.json())
    .then(student => {
      const html = `
        <p><strong>ID:</strong> ${student.stdid}</p>
        <p><strong>Name:</strong> ${student.fname} ${student.lname}</p>
        <p><strong>Email:</strong> ${student.email}</p>
      `;
      document.getElementById("studentContent").innerHTML = html;
      document.getElementById("studentDetail").classList.remove("hidden");
    });
}

function showCourseDetail(cid) {
  fetch(`/course/${cid}`)
    .then(res => res.json())
    .then(course => {
      const html = `
        <p><strong>Course ID:</strong> ${course.cid}</p>
        <p><strong>Course Name:</strong> ${course.coursename}</p>
      `;
      document.getElementById("courseContent").innerHTML = html;
      document.getElementById("courseDetail").classList.remove("hidden");
    });
}
