import {
    Button,
    Dialog,
    Grid,
    IconButton,
    InputLabel,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Tooltip,
    Typography,
    useTheme
} from '@mui/material';
import CancelTwoToneIcon from '@mui/icons-material/CancelTwoTone';
import DownloadTwoToneIcon from '@mui/icons-material/DownloadTwoTone';
import AttachFileTwoToneIcon from '@mui/icons-material/AttachFileTwoTone';
import { useIntl } from 'react-intl';
import React from 'react';
import { gridSpacing } from 'store/constant';
import * as XLSX from 'xlsx';
import { dispatch } from 'store';
import { openSnackbar } from 'store/slices/snackbar';
import { defaultErrorMessage } from 'constant/general';

type PropType = {
    id: string;
    isOpen: boolean;
    setIsOpen: () => void;
    confirmFunction: (data?: any) => void;
    // this will be exact checking which means totally equals, not include
    checkHeader?: string[];
    importTemplatePath?: string;
    importTemplateFileName?: string;
    valueFields?: string[];
};

// TODO handle multi file
const ImportDialog = React.memo(
    ({
        id,
        isOpen,
        setIsOpen,
        confirmFunction,
        checkHeader = [],
        importTemplatePath = '',
        importTemplateFileName = '',
        valueFields = []
    }: PropType) => {
        const intl = useIntl();
        const theme = useTheme();

        const [uploadedFiles, setUploadedFiles] = React.useState<any>(null);
        const [uploadedFileData, setUploadedFileData] = React.useState<any>(null);
        const [uploadedFileDataHeader, setUploadedFileDataHeader] = React.useState<any>(null);

        const [isHeaderCorrect, setIsHeaderCorrect] = React.useState<boolean>(true);

        const toggleModal = () => {
            setIsOpen();
            setUploadedFiles(null);
            setUploadedFileData(null);
            setUploadedFileDataHeader(null);
        };

        const readFile = (file: File) => {
            const reader = new FileReader();

            reader.onload = (e) => {
                if (e.target?.result) {
                    const binaryString = e.target.result as string;
                    const workbook = XLSX.read(binaryString, { type: 'binary' });

                    const sheetName = workbook.SheetNames[0];
                    const worksheet = workbook.Sheets[sheetName];

                    const excelArray = XLSX.utils.sheet_to_json(worksheet, { header: 1 });
                    const excelHeaders = excelArray[0] as string[];

                    if (checkHeader && checkHeader.length > 0) {
                        const headerCorrect = checkHeader.every((header) => excelHeaders.includes(header));
                        setIsHeaderCorrect(headerCorrect);
                        if (headerCorrect) {
                            setUploadedFileDataHeader(excelHeaders);
                            setUploadedFileData(excelArray);
                        } else {
                            setUploadedFiles(null);
                            setUploadedFileData(null);
                            setUploadedFileDataHeader(null);
                        }
                    } else {
                        setUploadedFileDataHeader(excelHeaders);
                        setUploadedFileData(excelArray);
                    }
                }
            };

            reader.readAsBinaryString(file);
        };

        React.useEffect(() => {
            if (!isHeaderCorrect) {
                dispatch(
                    openSnackbar({
                        open: true,
                        message: intl.formatMessage({ id: 'general.upload-file-error' }) || defaultErrorMessage,
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
        }, [isHeaderCorrect]);

        const handleFileUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
            const file = e.target.files?.[0];
            if (e.target.files && e.target.files.length > 0) {
                setUploadedFiles([...e.target.files]);
            }

            if (file) {
                readFile(file);
            }
            e.target.value = '';
        };

        function handleFileRemove(filename: any): void {
            // handling multi files upload
            if (uploadedFiles.length > 1) {
                setUploadedFiles([...uploadedFiles.filter((uploadedFile: any) => uploadedFile.name !== filename)]);
            }
            // handling single file upload
            else {
                setUploadedFiles(null);
                setUploadedFileData(null);
                setUploadedFileDataHeader(null);
            }
        }

        function handleFileDrop(e: React.DragEvent<HTMLLabelElement>): void {
            e.preventDefault();
            const file = e.dataTransfer.files[0];
            if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
                setUploadedFiles(e.dataTransfer.files);
            }
            if (file) {
                readFile(file);
            }
        }

        function handleData(uploadedFileData: any[]): any[] {
            let formattedDataList: any[] = [];

            if (valueFields.length > 0) {
                uploadedFileData.slice(1).forEach((row: any[]) => {
                    const formattedData: any = {};

                    valueFields.forEach((field, index) => {
                        formattedData[field] = row[index];
                    });

                    formattedDataList.push(formattedData);
                });
            } else {
                formattedDataList = uploadedFileData.slice(1); // Skip the header row
            }

            return formattedDataList;
        }

        return (
            <Dialog
                id={id}
                className="hideBackdrop"
                maxWidth="sm"
                fullWidth
                onClose={() => toggleModal()}
                open={isOpen}
                sx={{ '& .MuiDialog-paper': { p: '1.5rem 2rem' }, backgroundColor: '#f5f5f5' }}
            >
                {/* Header */}
                <Grid container spacing={gridSpacing}>
                    <Grid item xs={6} textAlign="left" sx={{ alignSelf: 'center', display: 'flex', alignItems: 'center' }}>
                        <Typography variant="h3">
                            {intl.formatMessage({ id: 'general.import' })}
                            {importTemplatePath !== '' && importTemplateFileName !== '' && (
                                <Tooltip title="general.download-template">
                                    <a href={importTemplatePath} download={importTemplateFileName}>
                                        <IconButton>
                                            <DownloadTwoToneIcon />
                                        </IconButton>
                                    </a>
                                </Tooltip>
                            )}
                        </Typography>
                    </Grid>
                    <Grid item xs={6} textAlign="right">
                        <IconButton onClick={() => toggleModal()}>
                            <CancelTwoToneIcon />
                        </IconButton>
                    </Grid>
                </Grid>
                {/* Body */}
                <Grid container spacing={gridSpacing} marginY="1rem">
                    <Grid item sm={12} textAlign="center" padding="1rem">
                        <input
                            accept=".csv,.xlsx,.xls,.xlsm,.xlsb,.ods"
                            id="upload-input"
                            type="file"
                            onChange={handleFileUpload}
                            disabled={uploadedFiles !== null}
                            style={{ display: 'none' }}
                        />
                        <InputLabel
                            disabled={uploadedFiles !== null}
                            htmlFor="upload-input"
                            onDragOver={(e) => e.preventDefault()}
                            onDrop={(e) => handleFileDrop(e)}
                        >
                            <Button variant="text" component="span" fullWidth sx={{ padding: '2rem' }} className="inputFileContainer">
                                <AttachFileTwoToneIcon />
                                {intl.formatMessage({ id: 'general.upload-file-content' })}
                            </Button>
                        </InputLabel>
                        {uploadedFiles &&
                            Array.from(uploadedFiles).map((file: any, index: number) => {
                                return (
                                    <Grid
                                        container
                                        spacing={gridSpacing}
                                        marginY="1rem"
                                        key={index}
                                        alignItems="center"
                                        className="generalContainer"
                                    >
                                        <Grid item xs={6}>
                                            <Typography variant="body1" textAlign="left">
                                                {file.name}
                                            </Typography>
                                        </Grid>
                                        <Grid item xs={6} textAlign="right">
                                            <IconButton onClick={() => handleFileRemove(file.name)}>
                                                <CancelTwoToneIcon />
                                            </IconButton>
                                        </Grid>
                                    </Grid>
                                );
                            })}
                        {uploadedFileData && (
                            <TableContainer style={{ maxHeight: '30vmax', width: '100%' }}>
                                <Table>
                                    <TableHead sx={{ bgcolor: theme.palette.primary.main }}>
                                        <TableRow>
                                            {uploadedFileDataHeader.map((cell: any, cellIndex: any) => (
                                                <TableCell key={cellIndex}>{cell}</TableCell>
                                            ))}
                                        </TableRow>
                                    </TableHead>
                                    <TableBody style={{ overflow: 'scroll', width: '100%', height: '100%' }}>
                                        {uploadedFileData.slice(1).map((row: any, rowIndex: any) => (
                                            <TableRow key={rowIndex}>
                                                {row.map((cell: any, cellIndex: any) => (
                                                    <TableCell key={cellIndex}>{cell || '-'}</TableCell>
                                                ))}
                                            </TableRow>
                                        ))}
                                    </TableBody>
                                </Table>
                            </TableContainer>
                        )}
                    </Grid>
                </Grid>
                {/* Footer */}
                <Grid container spacing={gridSpacing}>
                    <Grid item xs={12} sm={6} textAlign="center">
                        <Button
                            disabled={uploadedFiles === null && !isHeaderCorrect}
                            variant="contained"
                            onClick={() => confirmFunction(handleData(uploadedFileData))}
                            sx={{ alignSelf: 'end' }}
                        >
                            {intl.formatMessage({ id: 'general.confirm' })}
                        </Button>
                    </Grid>
                    <Grid item xs={12} sm={6} textAlign="center">
                        <Button variant="outlined" onClick={toggleModal} sx={{ alignSelf: 'end' }}>
                            {intl.formatMessage({ id: 'general.cancel' })}
                        </Button>
                    </Grid>
                </Grid>
            </Dialog>
        );
    }
);
export default ImportDialog;
