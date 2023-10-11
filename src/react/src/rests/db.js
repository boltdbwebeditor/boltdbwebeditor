function NormalizeError(respJson) {
    const ret = {
        respJson
    }

    if (respJson.error && respJson.error.Op && respJson.error.Path) {
        ret.normalizeError = `failed to ${respJson.error.Op} ${respJson.error.Path}`
    }

    return ret
}


export async function restGetDb() {
        const resp = await fetch("/api/db")
        const respJson = await resp.json()

        if (!resp.ok) {
            throw NormalizeError(respJson)
        }

        return respJson
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