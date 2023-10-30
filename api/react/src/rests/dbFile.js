export async function restDbFilePost(file) {
    const body = new FormData();

    body.append('upload.bolt.db', file);

    const options = {
        body,
        method: 'POST',
    }

    const resp = await fetch("/api/db/file", options)

    return await resp.json()
}

export async function resetDbFileGet() {
    const resp = await fetch("/api/db/file")
    const blob = await resp.blob()

    const filename = resp.headers.get('Content-Disposition').replace('attachment; filename=', '')

    return {blob, filename}
}