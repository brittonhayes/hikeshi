const halfmoon = require("halfmoon");
require("@fortawesome/fontawesome-free/js/all.js");
window.addEventListener('DOMContentLoaded', () => {
    console.log("loaded")
    halfmoon.onDOMContentLoaded();
})

const sidebarToggle = document.getElementById('sidebarToggle');
sidebarToggle.addEventListener('click', () => {
    sidebarToggle.classList.toggle("rotate-180");
    halfmoon.toggleSidebar();
})

const darkmodeToggle = document.getElementById('darkmodeToggle');
darkmodeToggle.addEventListener('click', () => {
    darkmodeToggle.classList.toggle("active");
    halfmoon.toggleDarkMode();
})

