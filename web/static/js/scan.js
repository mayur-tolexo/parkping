const message = document.getElementById("message");
const retryBtn = document.getElementById("retry");
let scanner;

function startScanner() {
    message.innerText = "Initializing camera...";
    retryBtn.classList.add("hidden");

    scanner = new Html5Qrcode("reader");

    scanner.start(
        { facingMode: "environment" },
        { fps: 10, qrbox: 250 },
        onScanSuccess,
        onScanFailure
    );
}

function onScanSuccess(fastag) {
    scanner.stop().then(() => {
        message.innerText = `FASTag detected: ${fastag}`;
        lookupVehicle(fastag);
    });
}

function onScanFailure(error) {
    // Ignored to reduce spam
}

function lookupVehicle(fastag) {
    const token = localStorage.getItem("token");
    if (!token) {
        alert("Please login first");
        return;
    }

    fetch(`/api/v1/vehicle/lookup?fastag=${fastag}`, {
        headers: { "Authorization": "Bearer " + token }
    })
        .then(res => {
            if (!res.ok) throw new Error("Vehicle not found or scan failed");
            return res.json();
        })
        .then(data => {
            // Redirect to profile page with vehicle info
            localStorage.setItem("vehicleData", JSON.stringify(data));
            window.location.href = "/profile.html";
        })
        .catch(err => {
            message.innerText = err.message;
            retryBtn.classList.remove("hidden");
        });
}

// Start scanner on page load
startScanner();
