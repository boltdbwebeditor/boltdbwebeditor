import {useCallback, useEffect, useState} from "react";

import {restGetDb, restPostDb} from "rests/db";

export function useData() {
    const [data, setData] = useState({})

    const getData = useCallback(async () => {
        const jsonData = await restGetDb()
        const data = {json: jsonData}
        setData(data)
    }, [])

    const postData = useCallback(async (data) => {
        return restPostDb(data)
    }, [])

    useEffect(() => {
        getData()
    }, []);

    return {data, setData, getData, postData};
}