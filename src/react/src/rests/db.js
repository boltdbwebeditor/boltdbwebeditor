export async function restGetDb() {
    const resp = await fetch("/api/db")
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