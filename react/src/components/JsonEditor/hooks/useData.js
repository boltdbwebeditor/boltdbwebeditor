import {useCallback, useState} from "react";
import {toast} from 'react-toastify';
import { saveAs } from 'file-saver';

import {fileSelector} from "components/FileSelector/FileSelector";
import {restDbJsonGet, restDbJsonPost} from "rests/dbJson";
import {resetDbFileGet, restDbFilePost} from "../../../rests/dbFile";

export function useData() {
    const [data, setData] = useState({json: {}})

    const getData = useCallback(async () => {
        try {
            const jsonData = await restDbJsonGet()
            const data = {json: jsonData}
            setData(data)
            toast.success("Download database successfully")
        } catch (e) {
            toast.error(e.errorMsg || "Failed to download database")
        }
    }, [])

    const postData = useCallback(async (data) => {
        try {
            const ret = await restDbJsonPost(data)
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
                const payload = await restDbFilePost(file)
                setData({json: payload.data})
            } catch (e) {

            }
        }
    })

    const downloadDb = useCallback(async () => {
        const {blob, filename} = await resetDbFileGet()
        saveAs(blob, filename)
    })

    return {data, setData, getData, postData, uploadDB, downloadDb};
}