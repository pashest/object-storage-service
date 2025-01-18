async function uploadFileInChunks() {
    const fileInput = document.getElementById('fileInput');
    const file = fileInput.files[0];

    if (!file) {
        alert("Please select a file first.");
        return;
    }

    const chunksNumber = 6;
    let chunkSize = Math.floor(file.size / 6);
    let extraSize = file.size % 6;

    for (let i = 0; i < chunksNumber; i++) {
        if (i === (chunksNumber -1)) {
            chunkSize += extraSize
        }
        const start = i * chunkSize;
        const end = Math.min(file.size, start + chunkSize);
        const chunk = file.slice(start, end);

        const formData = new FormData();
        formData.append("chunk", chunk);
        formData.append("chunkIndex", i);
        formData.append("chunkSize", chunkSize);
        formData.append("fileName", file.name);

        await fetch("http://localhost:8080/upload-chunk", {
            method: "POST",
            body: formData,
        });
    }
}

async function downloadFile() {
    const fileName = document.getElementById('fileName');

    if (!fileName) {
        alert("Please type file name");
        return;
    }

    const url = `http://localhost:8080/download?file_name=${fileName}`;
    const response = await fetch(url);

    if (!response.ok) {
        console.error('Error downloading file');
        return;
    }

    const blob = await response.blob();
    const link = document.createElement('a');
    link.href = URL.createObjectURL(blob);
    link.download = 'downloaded_file';
    link.click();
}
