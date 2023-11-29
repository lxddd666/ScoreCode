import { SyntheticEvent } from 'react';
import { useIntl } from 'react-intl';
import { Link } from 'react-router-dom';

// material-ui
import { Alert, Avatar, Button, Fade, Grid, Grow, IconButton, Slide, SlideProps, Typography } from '@mui/material';
import MuiSnackbar from '@mui/material/Snackbar';

import CloseIcon from '@mui/icons-material/Close';
import NotificationImportantTwoToneIcon from '@mui/icons-material/NotificationImportantTwoTone';
import CampaignTwoToneIcon from '@mui/icons-material/CampaignTwoTone';

// assets
import { KeyedObject } from 'types';
import { useDispatch, useSelector } from 'store';
import { closeSnackbar } from 'store/slices/snackbar';
import Chip from './Chip';

// animation function
function TransitionSlideLeft(props: SlideProps) {
    return <Slide {...props} direction="left" />;
}

function TransitionSlideUp(props: SlideProps) {
    return <Slide {...props} direction="up" />;
}

function TransitionSlideRight(props: SlideProps) {
    return <Slide {...props} direction="right" />;
}

function TransitionSlideDown(props: SlideProps) {
    return <Slide {...props} direction="down" />;
}

function GrowTransition(props: SlideProps) {
    return <Grow {...props} />;
}

// animation options
const animation: KeyedObject = {
    SlideLeft: TransitionSlideLeft,
    SlideUp: TransitionSlideUp,
    SlideRight: TransitionSlideRight,
    SlideDown: TransitionSlideDown,
    Grow: GrowTransition,
    Fade
};

// ==============================|| SNACKBAR ||============================== //

const Snackbar = () => {
    const dispatch = useDispatch();
    const intl = useIntl();
    const snackbar = useSelector((state) => state.snackbar);
    const { actionButton, anchorOrigin, alert, close, message, open, transition, variant } = snackbar;

    const jsonObj: any = variant === 'notification' ? JSON.parse(message) : message;

    const handleClose = (event: SyntheticEvent | Event, reason?: string) => {
        if (reason === 'clickaway') {
            return;
        }
        dispatch(closeSnackbar());
    };

    function handleTag(tag: number) {
        switch (tag) {
            case 0:
                return <></>;
            case 1:
                return <Chip label={intl.formatMessage({ id: 'general.normal' })} size="medium" chipcolor="primary" />;
            case 2:
                return <Chip label={intl.formatMessage({ id: 'general.urgent' })} size="medium" chipcolor="error" />;
            case 3:
                return <Chip label={intl.formatMessage({ id: 'general.important' })} size="medium" chipcolor="warning" />;
            case 4:
                return <Chip label={intl.formatMessage({ id: 'general.remind' })} size="medium" chipcolor="success" />;
            case 5:
                return <Chip label={intl.formatMessage({ id: 'general.secondary' })} size="medium" chipcolor="" />;
            default:
                return <></>;
        }
    }

    return (
        <>
            {/* default snackbar */}
            {variant === 'default' && (
                <MuiSnackbar
                    anchorOrigin={anchorOrigin}
                    open={open}
                    autoHideDuration={6000}
                    onClose={handleClose}
                    message={message}
                    TransitionComponent={animation[transition]}
                    action={
                        <>
                            <Button color="secondary" size="small" onClick={handleClose}>
                                UNDO
                            </Button>
                            <IconButton size="small" aria-label="close" color="inherit" onClick={handleClose} sx={{ mt: 0.25 }}>
                                <CloseIcon fontSize="small" />
                            </IconButton>
                        </>
                    }
                />
            )}

            {/* alert snackbar */}
            {variant === 'alert' && (
                <MuiSnackbar
                    TransitionComponent={animation[transition]}
                    anchorOrigin={anchorOrigin}
                    open={open}
                    autoHideDuration={6000}
                    onClose={handleClose}
                >
                    <Alert
                        severity={alert.severity}
                        variant={alert.variant}
                        color={alert.color}
                        action={
                            <>
                                {actionButton !== false && (
                                    <Button size="small" onClick={handleClose} sx={{ color: 'background.paper' }}>
                                        UNDO
                                    </Button>
                                )}
                                {close !== false && (
                                    <IconButton sx={{ color: 'background.paper' }} size="small" aria-label="close" onClick={handleClose}>
                                        <CloseIcon fontSize="small" />
                                    </IconButton>
                                )}
                            </>
                        }
                        sx={{
                            ...(alert.variant === 'outlined' && {
                                bgcolor: 'background.paper'
                            })
                        }}
                    >
                        {message}
                    </Alert>
                </MuiSnackbar>
            )}

            {variant === 'notification' && (
                <MuiSnackbar
                    TransitionComponent={animation[transition]}
                    anchorOrigin={anchorOrigin}
                    open={open}
                    autoHideDuration={10000}
                    onClose={handleClose}
                >
                    <Alert
                        icon={false}
                        color={alert.color}
                        action={
                            <>
                                {close !== false && (
                                    <IconButton size="small" aria-label="close" onClick={handleClose}>
                                        <CloseIcon fontSize="small" />
                                    </IconButton>
                                )}
                            </>
                        }
                        onClick={handleClose}
                        sx={{ width: '70%' }}
                    >
                        <Grid container alignItems="center">
                            <Grid item xs={3} textAlign="center">
                                <Avatar
                                    src={jsonObj.senderAvatar}
                                    alt={jsonObj.type === 3 ? jsonObj.senderAvatar : undefined}
                                    sizes="medium"
                                >
                                    {jsonObj.type === 1 && <CampaignTwoToneIcon sx={{ fontSize: '3rem' }} />}
                                    {jsonObj.type === 2 && <NotificationImportantTwoToneIcon sx={{ fontSize: '3rem' }} />}
                                </Avatar>
                            </Grid>
                            <Grid item xs={9} container>
                                <Grid item xs={8}>
                                    <Typography variant="h4">{jsonObj.title}</Typography>
                                </Grid>
                                <Grid item xs={4}>
                                    {handleTag(jsonObj.tag)}
                                </Grid>
                                <Grid item xs={12}>
                                    {jsonObj.content.length <= 100 && (
                                        <Typography variant="body1" dangerouslySetInnerHTML={{ __html: jsonObj.content }} />
                                    )}
                                    {jsonObj.content.length > 100 && (
                                        <Typography
                                            variant="body1"
                                            dangerouslySetInnerHTML={{ __html: jsonObj.content.substring(0, 100) + `...` }}
                                        />
                                    )}
                                </Grid>
                                <Grid item xs={12} container alignItems="center">
                                    <Grid item xs={7}>
                                        <Typography variant="body2" sx={{ paddingTop: '0.5rem', paddingBottom: '1rem' }}>
                                            {jsonObj.createdAt}
                                        </Typography>
                                    </Grid>
                                    <Grid item xs={5} marginTop="-0.5rem">
                                        <Link to={`/home/message?type=${jsonObj.type}`} onClick={handleClose}>
                                            {intl.formatMessage({ id: 'general.view-details' })}
                                        </Link>
                                    </Grid>
                                </Grid>
                            </Grid>
                        </Grid>
                    </Alert>
                </MuiSnackbar>
            )}
        </>
    );
};

export default Snackbar;
