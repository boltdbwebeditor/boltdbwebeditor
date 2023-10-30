import {useCallback, useState} from "react";

const MAX_EDITORS = 3

let currentId = 0;

function getNextEditorId() {
    return ++currentId;
}

function getNewEditor() {
    const nextId = getNextEditorId()
    return {
        Id: nextId
    }
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
    const [editorList, setEditorList] = useState([getNewEditor()])

    const newEditorCb = useCallback((Id) => {
        if (editorList.length < MAX_EDITORS) {
            const newEditor = getNewEditor()
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