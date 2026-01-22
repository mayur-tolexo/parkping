async function sendOTP() {
    const phone = document.getElementById("phone").value;

    await fetch("/api/v1/auth/send-otp", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ phone })
    });

    document.getElementById("otp-section").style.display = "block";
}

async function verifyOTP() {
    const phone = document.getElementById("phone").value;
    const otp = document.getElementById("otp").value;

    const res = await fetch("/api/v1/auth/verify-otp", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ phone, otp })
    });

    if (res.ok) {
        window.location.href = "/dashboard";
    } else {
        alert("Invalid OTP");
    }
}
