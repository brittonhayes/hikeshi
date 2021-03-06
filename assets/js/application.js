const halfmoon = require("halfmoon");
require("@fortawesome/fontawesome-free/js/all.js");
window.addEventListener('DOMContentLoaded', () => {
    console.log("loaded")
    halfmoon.onDOMContentLoaded();
})

const sidebarToggle = document.getElementById('sidebarToggle');
const sidebarToggleIcon = document.getElementById('sidebar-toggle-icon');
sidebarToggle.addEventListener('click', () => {
    sidebarToggleIcon.classList.toggle("fa-chevron-left");
    sidebarToggleIcon.classList.toggle("fa-chevron-right");
    halfmoon.toggleSidebar();
})

const darkmodeToggle = document.getElementById('darkmodeToggle');
darkmodeToggle.addEventListener('click', () => {
    darkmodeToggle.classList.toggle("active");
    halfmoon.toggleDarkMode();
})

