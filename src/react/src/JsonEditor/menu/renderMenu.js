import {faDownload, faUpload} from "assets/icons";
import {toJSONContent} from "vanilla-jsoneditor";

export const renderMenuFactory = ({getData, postData, editor}) => (items, context) => {
    const separator = {type: 'separator'}

    const downloadButton = {
        type: "button",
        icon: faDownload,
        title: "Download",
        onClick: getData,
    }

    const uploadButton = {
        type: "button",
        icon: faUpload,
        title: "Upload",
        onClick: () => {
            try {
                const isValid = editor.validate() === null
                const data = toJSONContent(editor.get()).json
                postData(data)
            } catch {
            }
        },
    }

    // pop space item
    items.pop();

    return [...items, separator, downloadButton, uploadButton]
}