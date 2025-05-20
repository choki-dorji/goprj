// include.js
function includeHTML() {
  document.querySelectorAll('[data-include]').forEach(el => {
    const file = el.getAttribute("data-include");
    fetch(file)
      .then(response => {
        if (response.ok) return response.text();
        throw new Error(`Could not fetch ${file}`);
      })
      .then(data => {
        el.innerHTML = data;
      })
      .catch(err => {
        el.innerHTML = `<div class="text-red-600 text-sm">${err.message}</div>`;
      });
  });
}

window.addEventListener("DOMContentLoaded", includeHTML);
