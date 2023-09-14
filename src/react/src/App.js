import {JsonEditorManager} from "./JsonEditorManager/JsonEditorManager";
import {FileUploadDialog} from "./FileUploadDialog/FileUploadDialog";
import {Fragment, useState } from "react";
import Modal from './UI/Modal'

function App() {
  const [showFileUpload, setShowFileUpload] = useState(true);

   const openUploadDialog = () => {
		console.log("handleFileUploadCLick:", showFileUpload)
        setShowFileUpload(true);
    };

	const closeUploadDialog = () => {
		setShowFileUpload(false);
	}

  return <Fragment>
        {showFileUpload && 
            <Modal onClose={closeUploadDialog}>
				<h2>Choose your DB file</h2>
				<FileUploadDialog onClose={closeUploadDialog} />
            </Modal>
        }
      <JsonEditorManager onShowFileUploader={openUploadDialog}/>
  </Fragment>
}

export default App;
