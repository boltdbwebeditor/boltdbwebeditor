import {faClose, faDownload, faPlus, faUpload} from "assets/icons";

export const renderMenuFactory = ({getData, onShowFileUploader, Id, newEditor, closeEditor, editor}) => (items, context) => {
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
        onClick: () => onShowFileUploader(),
    }

    const newEditorButton = {
        type: "button",
        icon: faPlus,
        title: "New Editor",
        onClick: () => {
            newEditor(Id)
        },
    }

    const closeEditorButton = {
        type: "button",
        icon: faClose,
        title: "Close Editor",
        onClick: () => {
            closeEditor(Id)
        },
    }

    return [newEditorButton, closeEditorButton, separator, downloadButton, uploadButton, separator, ...items]
}