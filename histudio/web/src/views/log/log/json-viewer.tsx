// material-ui
import { Grid, Typography, Divider } from '@mui/material';
import MainCard from '../../../ui-component/cards/MainCard';

import JsonView from 'react18-json-view';
import 'react18-json-view/src/style.css';
// If dark mode is needed, import `dark.css`.
// import 'react18-json-view/src/dark.css';

// project imports
// import { gridSpacing } from 'store/constant';

interface JsonViewerProps {
    title?: string;
    jsonString?: object;
}

const JsonViewer = (props: JsonViewerProps) => {
    return (
        <MainCard>
            <Grid container spacing={2}>
                {props.title && (
                    <>
                        <Grid item xs={12}>
                            <Typography variant="h4">{props.title}</Typography>
                        </Grid>
                        <Grid item xs={12}>
                            <Divider />
                        </Grid>
                    </>
                )}
                <Grid item xs={12}>
                    <MainCard
                        boxShadow
                        sx={{
                            '&:hover': {
                                transform: 'scale3d(1.01, 1.01, 1)',
                                transition: 'all .4s ease-in-out'
                            }
                        }}
                        style={{ border: '1px solid #e3e3e3', borderRadius: '4px', padding: '10px', transition: '0.4s' }}
                    >
                        <div>
                            <JsonView
                                // dark={true}
                                // theme="github"
                                collapsed={2}
                                // collapseObjectsAfterLength={12}
                                collapseStringsAfterLength={150}
                                collapseStringMode="address"
                                enableClipboard={true}
                                src={props.jsonString}
                            />
                        </div>
                    </MainCard>
                </Grid>
            </Grid>
        </MainCard>
    );
};

export default JsonViewer;
