import {useCallback, useEffect, useState} from "react";
import { toast } from 'react-toastify';

import {restGetDb, restPostDb} from "rests/db";

export function useData() {
    const [data, setData] = useState({json: {}})

    const getData = useCallback(async () => {
        try {
            const jsonData = await restGetDb()
            const data = {json: jsonData}
            setData(data)
            toast.success("Download database successfully")
        } catch (e) {
            toast.error("Failed to download database")
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

    useEffect(() => {
        getData()
    }, []);

    return {data, setData, getData, postData};
}