import {useCallback, useEffect, useRef} from "react";

import {VanillaJsonEditor} from "./VanillaJsonEditor/VanillaJsonEditor";
import {useData} from "./hooks/useData";
import {renderMenuFactory} from "./menu/renderMenu";

import "./css/override-jse.css";

export function JsonEditor({Id, newEditor, closeEditor}) {
    const refEditor = useRef(null)

    const {data, setData, getData, postData, uploadDB, downloadDb} = useData()

    const onRenderMenu = renderMenuFactory({
        getData,
        postData,
        uploadDB,
        downloadDb,
        Id,
        newEditor,
        closeEditor,
        editor: refEditor.current
    })

    const onChange = useCallback(
        (updatedContent, previousContent, { contentErrors, patchResult }) => {
            setData(updatedContent)
        }, []
    )

    useEffect(() => {
        getData()
    }, []);

    return (
        <VanillaJsonEditor {...{
            Id,
            refEditor,
            content: data,
            onChange,
            onRenderMenu,
            escapeControlCharacters: true,
            escapeUnicodeCharacters: true,
        }} />
    );
}
