import ReactQuill from 'react-quill';
import React from 'react';
import 'react-quill/dist/quill.snow.css';
import { openSnackbar } from 'store/slices/snackbar';
import Request from 'utils/request';
import Loading from 'ui-component/Loading';
import { useState } from 'react';
import { Quill } from 'react-quill';
import Modal from '@mui/material/Modal';
import Box from '@mui/material/Box';
import CloseIcon from '@mui/icons-material/Close';
import { gridSpacing } from 'store/constant';
import { Grid, Button } from '@mui/material';
import { useIntl } from 'react-intl';
import AnimateButton from 'ui-component/extended/AnimateButton';
import { TextField } from '@mui/material';

interface MyEditorParam {
    value?: string;
    formik?: any;
    keyId?: string;
    dispatch?: any;
}

const styleModal = {
    bgcolor: 'background.paper',
    boxShadow: 24,
    p: 3,
    width: '100%'
};

const Embed = Quill.import('blots/block/embed');

class CustomEmbed extends Embed {
    static create(data: any) {
        const node = super.create();
        node.innerHTML = data.html;
        return node;
    }

    static value(node: any) {
        return {
            html: node.innerHTML
        };
    }
}

CustomEmbed.blotName = 'customEmbed';
CustomEmbed.tagName = 'div';
Quill.register(CustomEmbed);

const MyEditor = ({ value, formik, keyId, dispatch }: MyEditorParam) => {
    const intl = useIntl();
    const quillRef = React.useRef<ReactQuill | null>(null);
    const [isUpload, setIsUpload] = useState<boolean>(false);
    const [isShowPreview, setIsShowPreview] = useState<boolean>(false);
    const [imagePreview, setImagePreview] = useState<string>('');
    const [width, setWidth] = useState<string>('200');
    const [height, setHeight] = useState<string>('200');

    const ImageHandler = async (image: any, callback: any) => {
        const input = document.createElement('input');
        input.setAttribute('type', 'file');
        input.setAttribute('accept', 'image/*');
        input.click();

        input.onchange = async () => {
            const fileUpload = input.files?.[0];

            if (fileUpload) {
                const formData = new FormData();
                formData.append('type', '0');
                formData.append('file', fileUpload);
                setIsUpload(true);
                const response = await Request(
                    {
                        url: 'admin/upload/file',
                        method: 'POST',
                        param: formData,
                        header: {
                            'Content-Type': 'multipart/form-data',
                            uploadType: 'image'
                        }
                    },
                    dispatch,
                    intl
                );
                setIsUpload(false);
                if (response?.data?.code == 0) {
                    dispatch(
                        openSnackbar({
                            open: true,
                            message: response?.data?.message,
                            variant: 'alert',
                            alert: {
                                color: 'success'
                            },
                            close: false,
                            anchorOrigin: {
                                vertical: 'top',
                                horizontal: 'center'
                            }
                        })
                    );

                    setIsShowPreview(true);
                    setImagePreview(response?.data?.data?.fileUrl);

                    /*
                    const range = quillRef?.current?.getEditor()?.getSelection();
                    const imageElement = `<img src="${response?.data?.data?.fileUrl}" style="width: ${100}px;height: ${100}px;">`;
                    quillRef?.current?.getEditor()?.insertEmbed(range?.index || 0,  'customEmbed', { html : imageElement });
                    */
                } else {
                    dispatch(
                        openSnackbar({
                            open: true,
                            message: response?.data?.message,
                            variant: 'alert',
                            alert: {
                                color: 'error'
                            },
                            close: false,
                            anchorOrigin: {
                                vertical: 'top',
                                horizontal: 'center'
                            }
                        })
                    );
                }
            }
        };
    };

    React.useEffect(() => {
        const editor = quillRef?.current?.getEditor();
        editor?.getModule('toolbar').addHandler('image', ImageHandler);
    }, []);

    const handleChangeEditor = (content: string) => {
        formik?.setFieldValue(keyId, content);
    };

    const handleClose = () => {
        setIsShowPreview(false);
        setImagePreview('');
        setWidth('200');
        setHeight('200');
    };

    function insertImage() {
        const range = quillRef?.current?.getEditor()?.getSelection();
        const imageElement = `<img src="${imagePreview}" style="width: ${width}px;height: ${height}px;object-fit: contain;">`;
        quillRef?.current?.getEditor()?.insertEmbed(range?.index || 0, 'customEmbed', { html: imageElement });
        handleClose();
    }

    const changeValue = (key: string, value: string) => {
        switch (key) {
            case 'width':
                setWidth(value);
                break;
            case 'height':
                setHeight(value);
                break;
        }
    };
    return (
        <>
            {isUpload && <Loading isTransprent={true} isFixed={true} zIndex={1101}/>}
            <ReactQuill
                ref={quillRef}
                theme="snow"
                value={value}
                onChange={handleChangeEditor}
                modules={{
                    toolbar: [
                        [{ font: [] }, { header: [1, 2, 3, 4, 5, 6, false] }],
                        ['bold', 'italic', 'underline', 'strike'],
                        [{ color: [] }, { background: [] }],
                        [{ script: 'super' }, { script: 'sub' }],
                        [{ header: '1' }, { header: '2' }, 'blockquote', 'code-block'],
                        [{ list: 'ordered' }, { list: 'bullet' }, { indent: '-1' }, { indent: '+1' }],
                        ['direction', { align: [] }],
                        ['link', 'image', 'video', 'formula'],
                        ['clean']
                    ]
                }}
            />
            <Modal
                className=""
                open={isShowPreview}
                onClose={handleClose}
                aria-labelledby="modal-modal-title"
                aria-describedby="modal-modal-description"
            >
                <div className="general-modal-background">
                    <div className="general-modal-dialog">
                        <Box sx={styleModal}>
                            <div className="general-modal-header">
                                <div className="general-modal-title">{intl.formatMessage({ id: 'general.setting' })}</div>
                                <div className="general-modal-close" onClick={handleClose}>
                                    <CloseIcon />
                                </div>
                            </div>

                            <div className="general-modal-body mail-modal-body">
                                <Grid container spacing={gridSpacing} style={{ display: 'flex', justifyContent: 'center' }}>
                                    <Grid item xs={12} sm={12} sx={{ display: 'flex', justifyContent: 'center' }}>
                                        <img src={imagePreview} width={150} height={150} style={{ objectFit: 'contain' }} />
                                    </Grid>
                                    <Grid item xs={12} sm={12}>
                                        <TextField
                                            fullWidth
                                            id={`width`}
                                            name={`height`}
                                            label={intl.formatMessage({ id: 'general.width' })}
                                            value={width}
                                            onChange={(e) => changeValue('width', e?.target?.value)}
                                        />
                                    </Grid>
                                    <Grid item xs={12} sm={12}>
                                        <TextField
                                            fullWidth
                                            id={`width`}
                                            name={`height`}
                                            label={intl.formatMessage({ id: 'general.height' })}
                                            value={height}
                                            onChange={(e) => changeValue('height', e?.target?.value)}
                                        />
                                    </Grid>
                                </Grid>
                            </div>

                            <div className="general-modal-footer">
                                <AnimateButton>
                                    <Button
                                        variant="contained"
                                        style={{ display: 'flex', marginRight: 5 }}
                                        color="primary"
                                        onClick={handleClose}
                                    >
                                        {intl.formatMessage({ id: 'general.cancel' })}
                                    </Button>
                                </AnimateButton>
                                <AnimateButton>
                                    <Button variant="contained" style={{ display: 'flex' }} color="secondary" onClick={() => insertImage()}>
                                        {intl.formatMessage({ id: 'general.ok' })}
                                    </Button>
                                </AnimateButton>
                            </div>
                        </Box>
                    </div>
                </div>
            </Modal>
        </>
    );
};

export default MyEditor;
