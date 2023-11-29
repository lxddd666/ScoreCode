//material-ui
import AddCircleOutlineTwoToneIcon from '@mui/icons-material/AddCircleOutlineTwoTone';
import { IconButton, Tooltip } from '@mui/material';

type PropType = {
    onClick?: () => void;
    tooltipTitle: string;
};

const AddButton = ({ onClick, tooltipTitle }: PropType) => {
    return (
        <Tooltip title={tooltipTitle}>
            <IconButton size="large" onClick={onClick}>
                <AddCircleOutlineTwoToneIcon />
            </IconButton>
        </Tooltip>
    );
};

export default AddButton;
