import {useCallback, useEffect, useState} from "react";

import {restGetDb, restPostDb} from "rests/db";

export function useData() {
    const [data, setData] = useState()

    const getData = useCallback(async () => {
        const data = await restGetDb()
        setData(data)
    }, [])

    const postData = useCallback(async (data) => {
        return restPostDb(data)
    }, [])

    useEffect(() => {
        getData()
    }, []);

    return [data, getData, postData];
}