import {useCallback, useState} from "react";

const MAX_EDITORS = 3

function getNewEditor(Id) {
    return {
        Id
    }
}

function getNextEditorId(editorList) {
    return editorList.length + 1
}

function editorListInsert(editorList, Id, newEditor) {
    const newEditorList = [...editorList];

    const index = newEditorList.findIndex(editor => editor.Id === Id);
    if (index !== -1) {
        newEditorList.splice(index + 1, 0, newEditor)
    }

    return newEditorList
}

function editorListRemove(editorList, Id) {
    const newEditorList = [...editorList];

    const index = newEditorList.findIndex(editor => editor.Id === Id);
    if (index !== -1) {
        newEditorList.splice(index, 1)
    }

    return newEditorList
}

export function useJsonEditorManager() {
    const [editorList, setEditorList] = useState([getNewEditor(1)])

    const newEditorCb = useCallback((Id) => {
        if (editorList.length < MAX_EDITORS) {
            const nextId = getNextEditorId(editorList)
            const newEditor = getNewEditor(nextId)
            const newEditorList = editorListInsert(editorList, Id, newEditor)

            setEditorList(newEditorList)
        }
    }, [editorList])

    const closeEditorCb = useCallback((Id) => {
        if (editorList.length > 1) {
            const newEditorList = editorListRemove(editorList, Id)
            setEditorList(newEditorList)
        }
    }, [editorList])


    return {
        editorList,
        newEditor: newEditorCb,
        closeEditor: closeEditorCb,
    };
}