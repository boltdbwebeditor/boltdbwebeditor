import {faClose, faDisk, faDownload, faRefresh, faSplit, faUpload} from "assets/icons";
import {toJSONContent} from "vanilla-jsoneditor";

export const renderMenuFactory = ({getData, postData, uploadDB, Id, newEditor, closeEditor, editor}) => (items, context) => {
    const separator = {type: 'separator'}

    const splitViewButton = {
        type: "button",
        icon: faSplit,
        title: "Split View",
        onClick: () => {
            newEditor(Id)
        },
    }

    const closeViewButton = {
        type: "button",
        icon: faClose,
        title: "Close View",
        onClick: () => {
            closeEditor(Id)
        },
    }

    const uploadButton = {
        type: "button",
        icon: faUpload,
        title: "Upload",
        onClick: uploadDB,
    }

    const downloadButton = {
        type: "button",
        icon: faDownload,
        title: "Download",
        onClick: getData,
    }

    const saveButton = {
        type: "button",
        icon: faDisk,
        title: "Save",
        onClick: () => {
            try {
                const isValid = editor.validate() === null
                const data = toJSONContent(editor.get()).json
                postData(data)
            } catch {
            }
        },
    }

    const reloadButton = {
        type: "button",
        icon: faRefresh,
        title: "Reload",
        onClick: getData,
    }

    return [
        splitViewButton, closeViewButton, separator,
        uploadButton, downloadButton, separator,
        saveButton, reloadButton, separator,
        ...items
    ]
}
