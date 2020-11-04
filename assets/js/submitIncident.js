const halfmoon = require("halfmoon");
const submitIncident = document.getElementById('submitIncident')
submitIncident.addEventListener('click', () => {
    toastAlert("Submitted", "Thanks for your submission!", "primary", 3000)
})

function toastAlert(Title, Body, Type, Timeout) {
    halfmoon.initStickyAlert({
        content: Body,
        title: Title,
        alertType: `alert-${Type}`,
        hasDismissButton: false,
        timeShown: Timeout
    });
}
