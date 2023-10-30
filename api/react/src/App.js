import {JsonEditorManager} from "./JsonEditorManager/JsonEditorManager";
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function App() {
  return (
      <>
        <JsonEditorManager/>
        <ToastContainer />
      </>
  );
}

export default App;
