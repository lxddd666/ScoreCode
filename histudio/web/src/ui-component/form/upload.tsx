import React from 'react';
import { Upload, Spin } from 'antd';
import { useIntl } from 'react-intl';
import { useState } from 'react';
import DeleteIcon from '@mui/icons-material/Delete';
import Dialog from '@mui/material/Dialog';
import { TransitionProps } from '@mui/material/transitions';
import Slide from '@mui/material/Slide';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { Button } from '@mui/material';

interface UploadElementParam {
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
    headers: { [key: string]: string };
    data: { [key: string]: string };
    value: string;
    fileRemove: (key: string, id: string) => void;
}

const Transition = React.forwardRef(function Transition(
    props: TransitionProps & {
        children: React.ReactElement<any, any>;
    },
    ref: React.Ref<unknown>
) {
    return <Slide direction="down" ref={ref} {...props} />;
});

function UploadElement({
    onChange = (e: React.ChangeEvent<HTMLInputElement>) => {},
    headers = {},
    data = {},
    value = '',
    fileRemove = (key: string, id: string) => {}
}: UploadElementParam) {
    const [isUploading, setIsUploading] = useState<boolean>(false);
    const [open, setOpen] = useState<boolean>(false);

    const intl = useIntl();

    const uploadButton = (
        <div>
            +<div style={{ marginTop: 8 }}>{intl.formatMessage({ id: 'geneal.upload' })}</div>
        </div>
    );

    const handleClose = () => {
        setOpen(false);
    };

    const handleOk = () => {
        if (typeof fileRemove !== 'undefined') {
            fileRemove("", "");
        }
        setOpen(false);
    };

    const uploadImage = async (e: any) => {
        setIsUploading(true);
        await onChange(e);
        setIsUploading(false);
    };

    const removeImage = (e: any) => {
        e.stopPropagation();
        setOpen(true);
    };

    return (
        <>
            <Upload
                name="file"
                listType="picture-card"
                className="avatar-uploader-general"
                showUploadList={false}
                beforeUpload={() => false}
                onChange={(e: any) => uploadImage(e)}
                headers={headers}
                data={data}
                onPreview={() => {}}
            >
                {value ? <img src={value} alt="avatar" style={{ width: '100%' }} /> : uploadButton}

                {isUploading && <Spin style={{ position: 'absolute' }} />}
                {value && (
                    <DeleteIcon
                        onClick={(e: any) => removeImage(e)}
                        style={{ position: 'absolute', top: 0, right: 0, color: 'red', fontSize: '20px', marginRight:3, marginTop: 2 }}
                    />
                )}
            </Upload>

            <Dialog
                open={open}
                TransitionComponent={Transition}
                keepMounted
                onClose={handleClose}
                aria-describedby="alert-dialog-slide-description"
                className='general-dialog'
            >
                <DialogTitle>{intl.formatMessage({ id: 'general.notice' })}</DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-slide-image">
                        {intl.formatMessage({ id: 'general.remove.alert' })}
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}>{intl.formatMessage({ id: 'general.cancel' })}</Button>
                    <Button onClick={handleOk}>{intl.formatMessage({ id: 'general.ok' })}</Button>
                </DialogActions>
            </Dialog>
        </>
    );
}

export default UploadElement;
