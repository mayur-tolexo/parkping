async function loadData() {
    const res = await fetch("/api/v1/scan/history");
    const data = await res.json();

    const container = document.getElementById("results");
    container.innerHTML = "";

    data.forEach(item => {
        container.innerHTML += `
      <div class="card">
        <p>Vehicle: ${item.vehicleNumber}</p>
        <p>FASTag: ${item.qrToken}</p>
      </div>`;
    });
}

loadData();
