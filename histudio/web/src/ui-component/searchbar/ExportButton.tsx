//material-ui
import CloudDownload from '@mui/icons-material/CloudDownload';
import { useIntl } from 'react-intl';
import { IconButton, Tooltip } from '@mui/material';

type PropType = {
    onClick?: () => void;
};

const ExportButton = ({ onClick }: PropType) => {
    return (
        <Tooltip title={useIntl().formatMessage({ id: 'general.export' })}>
            <IconButton size="large" onClick={onClick}>
                <CloudDownload />
            </IconButton>
        </Tooltip>
    );
};

export default ExportButton;