import {useCallback, useRef } from "react";

import {VanillaJsonEditor} from "./VanillaJsonEditor/VanillaJsonEditor";
import {useData} from "./hooks/useData";
import {renderMenuFactory} from "./menu/renderMenu";

import "./css/override-jse.css";

export function JsonEditor({Id, newEditor, closeEditor, onShowFileUploader, onCloseFileUploader}) {
    const refEditor = useRef(null)

    const {data, setData, getData, postData} = useData()

     const onRenderMenu = renderMenuFactory({
        getData,
        onShowFileUploader,
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

    return <VanillaJsonEditor {...{Id, refEditor, content: data, onChange, onRenderMenu}} />;
}
