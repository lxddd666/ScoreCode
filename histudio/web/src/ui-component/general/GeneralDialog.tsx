import { Button, Dialog, Grid, IconButton, Typography } from '@mui/material';
import CancelTwoToneIcon from '@mui/icons-material/CancelTwoTone';
import InfoTwoToneIcon from '@mui/icons-material/InfoTwoTone';
import CheckCircleTwoToneIcon from '@mui/icons-material/CheckCircleTwoTone';
import WarningTwoToneIcon from '@mui/icons-material/WarningTwoTone';
import ErrorTwoToneIcon from '@mui/icons-material/ErrorTwoTone';
import { useIntl } from 'react-intl';
import React from 'react';
import { gridSpacing } from 'store/constant';

type PropType = {
    id: string;
    type: 'info' | 'success' | 'warning' | 'error';
    title?: string;
    content: React.ReactNode;
    isOpen: boolean;
    setIsOpen: () => void;
    confirmFunction: () => void;
};

const GeneralDialog = React.memo(({ id, type, title, content, isOpen, setIsOpen, confirmFunction }: PropType) => {
    const intl = useIntl();
    const [primaryButtonColor, setPrimaryButtonColor] = React.useState<'info' | 'success' | 'warning' | 'error' | 'primary'>('primary');
    React.useEffect(() => {
        setPrimaryButtonColor(type);
    }, []);

    const handleModalOpen = () => {
        setIsOpen();
    };

    const handleTitle = () => {
        if (title) {
            return title;
        } else {
            switch (type) {
                case 'success':
                    return intl.formatMessage({ id: 'general.success' });
                case 'warning':
                    return intl.formatMessage({ id: 'general.warning' });
                case 'error':
                    return intl.formatMessage({ id: 'general.error' });
                case 'info':
                    return intl.formatMessage({ id: 'general.info' });
                default:
                    return '';
            }
        }
    };

    const handleTypeIcon = (type: 'info' | 'success' | 'warning' | 'error') => {
        switch (type) {
            case 'success':
                return <CheckCircleTwoToneIcon color="success" />;
            case 'warning':
                return <WarningTwoToneIcon color="warning" />;
            case 'error':
                return <ErrorTwoToneIcon color="error" />;
            case 'info':
                return <InfoTwoToneIcon color="info" />;
            default:
                return <InfoTwoToneIcon color="info" />;
        }
    };

    return (
        <Dialog
            id={id}
            className="hideBackdrop"
            maxWidth="sm"
            fullWidth
            onClose={() => handleModalOpen()}
            open={isOpen}
            sx={{ '& .MuiDialog-paper': { p: '1.5rem 2rem' }, backgroundColor: '#f5f5f5' }}
        >
            {/* Header */}
            <Grid container spacing={gridSpacing}>
                <Grid item xs={6} textAlign="left" sx={{ alignSelf: 'center', display: 'flex', alignItems: 'center' }}>
                    {handleTypeIcon(type)}
                    <Typography sx={{ fontWeight: '800' }}>{handleTitle()}</Typography>
                </Grid>
                <Grid item xs={6} textAlign="right">
                    <IconButton onClick={() => handleModalOpen()} sx={{ alignSelf: 'end' }}>
                        <CancelTwoToneIcon />
                    </IconButton>
                </Grid>
            </Grid>
            {/* Body */}
            <Grid container spacing={gridSpacing} marginY="2rem">
                <Grid item sm={12} textAlign="center" padding="1rem">
                    <Typography>{content}</Typography>
                </Grid>
            </Grid>
            {/* Footer */}
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12} sm={6} textAlign="center">
                    <Button color={primaryButtonColor} variant="contained" onClick={confirmFunction} sx={{ alignSelf: 'end' }}>
                        {intl.formatMessage({ id: 'general.confirm' })}
                    </Button>
                </Grid>
                {type === 'warning' && (
                    <Grid item xs={12} sm={6} textAlign="center">
                        <Button color="secondary" variant="outlined" onClick={() => handleModalOpen()} sx={{ alignSelf: 'end' }}>
                            {intl.formatMessage({ id: 'general.cancel' })}
                        </Button>
                    </Grid>
                )}
            </Grid>
        </Dialog>
    );
});
export default GeneralDialog;
