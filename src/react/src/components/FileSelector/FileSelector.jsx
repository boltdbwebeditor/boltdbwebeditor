import {toast} from "react-toastify";
import {useState} from "react";

const toastId = "file-selector"

const toastOptions = {
    toastId,
    draggable: true,
    position: "top-center",
    autoClose: false,
    closeButton: false,
    closeOnClick: false,
    hideProgressBar: true,
}

export function fileSelector() {
    return new Promise((resolve, reject) => {
        const onClose = file => {
            toast.done(toastId)
            resolve(file)
        }

        if (toast.isActive(toastId)) {
            resolve(null)
        } else {
            toast(<FileSelector onClose={onClose}/>, toastOptions)
        }
    });
}

export function FileSelector({onClose}) {
    const [file, setFile] = useState(null);

    const onChange = (event) => {
        setFile(event.target.files[0]);
    };

    const onOk = () => onClose(file)
    const onCancel = () => onClose(null)
    // {
    //     if (file) {
    //         // Create a new FormData object
    //         const formData = new FormData();
    //         formData.append('file', file);
    //
    //         // Make the API request
    //         fetch('/api/db/upload', {
    //             method: 'POST',
    //             body: formData,
    //         })
    //             .then((response) => response.json())
    //             .then((data) => {
    //                 // Handle the response from the server
    //                 console.log("=====", data);
    //                 // props.onClose();
    //             })
    //             .catch((error) => {
    //                 // Handle any errors
    //                 console.error(error);
    //             });
    //     }
    // };

    return (
        <div>
            <input type="file" onChange={onChange}/>
            <div>
                <button onClick={onOk} disabled={file === null}>
                    OK
                </button>
                <button onClick={onCancel}>
                    Cancel
                </button>
            </div>
        </div>
    )
}
