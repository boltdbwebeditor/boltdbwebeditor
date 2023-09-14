export async function restGetDb() {
    const resp = await fetch("/api/db")
    if (resp.status !== 200) {
        return {}
    }
    return await resp.json()
}

export async function restPostDb(db) {
    const options = {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(db)
    }

    const resp = await fetch("/api/db", options)

    return await resp.json()
}

export async function restDownloadDb() {
    const resp = await fetch("/api/db/download")
    return await resp.blob()
}
