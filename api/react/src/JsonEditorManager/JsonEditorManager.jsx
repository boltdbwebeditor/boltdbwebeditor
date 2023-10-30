import {JsonEditor} from "JsonEditor/JsonEditor.jsx";

import {useJsonEditorManager} from "./hooks/useJsonEditorManager";

export function JsonEditorManager() {
  const {editorList, newEditor, closeEditor} = useJsonEditorManager()

  return (
      <>
          {editorList.map(
              (editor) => {
                  return <JsonEditor key={editor.Id} Id={editor.Id} newEditor={newEditor} closeEditor={closeEditor}/>
              }
          )}
      </>
  );
}
