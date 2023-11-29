//material-ui
import UploadTwoToneIcon from '@mui/icons-material/UploadTwoTone';
import { IconButton, Tooltip } from '@mui/material';

type PropType = {
    onClick?: () => void;
    tooltipTitle: string;
};

const ImportButton = ({ onClick, tooltipTitle }: PropType) => {
    return (
        <Tooltip title={tooltipTitle}>
            <IconButton size="large" onClick={onClick}>
                <UploadTwoToneIcon />
            </IconButton>
        </Tooltip>
    );
};

export default ImportButton;
