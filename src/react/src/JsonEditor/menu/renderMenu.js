import { faClose, faDownload, faPlus, faUpload, faRotateRight } from "assets/icons";

export const renderMenuFactory =
  ({
    getData,
    postData,
    onDownloadFile,
    onShowFileUploader,
    Id,
    newEditor,
    closeEditor,
    editor,
  }) =>
  (items, context) => {
    const separator = { type: "separator" };

    const refreshButton = {
      type: "button",
      icon: faRotateRight,
      title: "Refresh",
      onClick: getData,
    };

    const downloadButton = {
      type: "button",
      icon: faDownload,
      title: "Download",
      onClick: () => onDownloadFile(),
    };

    const uploadButton = {
      type: "button",
      icon: faUpload,
      title: "Upload",
      onClick: () => onShowFileUploader(),
    };

    const newEditorButton = {
      type: "button",
      icon: faPlus,
      title: "New Editor",
      onClick: () => {
        newEditor(Id);
      },
    };

    const closeEditorButton = {
      type: "button",
      icon: faClose,
      title: "Close Editor",
      onClick: () => {
        closeEditor(Id);
      },
    };

    return [
      newEditorButton,
      closeEditorButton,
      separator,
      refreshButton,
      separator,
      downloadButton,
      uploadButton,
      separator,
      ...items,
    ];
  };
