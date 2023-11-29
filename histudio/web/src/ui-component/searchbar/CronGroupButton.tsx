//material-ui
import FeaturedPlayListOutlinedIcon from '@mui/icons-material/FeaturedPlayListOutlined';
import { IconButton, Tooltip } from '@mui/material';

type PropType = {
    onClick?: () => void;
    tooltipTitle: string;
};

const AddButton = ({ onClick, tooltipTitle }: PropType) => {
    return (
        <Tooltip title={tooltipTitle}>
            <IconButton size="large" onClick={onClick}>
                <FeaturedPlayListOutlinedIcon />
            </IconButton>
        </Tooltip>
    );
};

export default AddButton;
