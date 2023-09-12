import {VanillaJsonEditor} from "./VanillaJsonEditor/VanillaJsonEditor";
import {useData} from "./hooks/useData";
import {renderMenuFactory} from "./menu/renderMenu";

import "./css/override-jse.css";
import {useRef} from "react";

export function JsonEditor() {
    const refEditor = useRef(null)

    const [data, getData, postData] = useData()

    const onRenderMenu = renderMenuFactory({
        getData,
        postData,
        editor: refEditor.current
    })

    return (
        <VanillaJsonEditor
            refEditor={refEditor}
            content={{json: data}}
            onRenderMenu={onRenderMenu}
        />
    );
}
