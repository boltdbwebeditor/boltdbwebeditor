import {useCallback, useState} from "react";
import {toast} from 'react-toastify';

import {restGetDb, restPostDb} from "rests/db";
import {fileSelector} from "components/FileSelector/FileSelector";
import {restPostDbFile} from "../../rests/dbFile";

export function useData() {
    const [data, setData] = useState({json: {}})

    const getData = useCallback(async () => {
        try {
            const jsonData = await restGetDb()
            const data = {json: jsonData}
            setData(data)
            toast.success("Download database successfully")
        } catch (e) {
            toast.error(e.errorMsg || "Failed to download database")
        }
    }, [])

    const postData = useCallback(async (data) => {
        try {
            const ret = await restPostDb(data)
            toast.success("Upload database successfully")
            return ret
        } catch (e) {
            toast.error("Failed to upload database")
        }
    }, [])

    const uploadDB = useCallback(async () => {
        const file = await fileSelector()
        if (file) {
            try {
                const payload = await restPostDbFile(file)
                setData({json: payload.data})
            } catch (e) {

            }
        }
    })

    return {data, setData, getData, postData, uploadDB};
}