function uploadFile() {
    const fileInput = document.getElementById('fileInput');
    const username = document.getElementById("username").value || "user";

    if (fileInput.files.length === 0) {
        alert("Please select a file");
        return;
    }

    const file = fileInput.files[0];

    file.arrayBuffer().then(buffer => {
        fetch("http://localhost:8080/upload", {
            method: "POST",
            body: buffer,
            headers: {
                "X-Filename": file.name,
                "Content-Type": "application/octet-stream",
                "Username": username
            }
        })
            .then(response => {
                if (response.ok) {
                    alert("File uploaded successfully!");
                } else {
                    alert("File upload failed.");
                }
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });
}

function downloadFile() {
    const fileName = document.getElementById('fileName').value.trim();
    const username = document.getElementById("username").value || "user";

    if (!fileName) {
        alert("Please enter a file name");
        return;
    }

    fetch(`http://localhost:8080/download?file_name=${encodeURIComponent(fileName)}`, {
        method: "GET",
        headers: {
            "Username": username
        }
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("File not found.");
            }
            return response.blob();
        })
        .then(blob => {
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement("a");
            a.href = url;
            a.download = fileName;
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
            window.URL.revokeObjectURL(url);
        })
        .catch(error => {
            console.error("Error:", error);
            alert("File download failed.");
        });
}
