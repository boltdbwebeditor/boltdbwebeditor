function NormalizeError(respJson) {
    const ret = {
        respJson
    }

    if (respJson.error && respJson.error.Op && respJson.error.Path) {
        ret.normalizeError = `failed to ${respJson.error.Op} ${respJson.error.Path}`
    }

    return ret
}

export async function restGetDbFile() {
        const resp = await fetch("/api/db/file")
        const respJson = await resp.json()

        if (!resp.ok) {
            throw NormalizeError(respJson)
        }

        return respJson
}

export async function restPostDbFile(file) {
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