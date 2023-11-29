//material-ui
import RestartAltTwoToneIcon from '@mui/icons-material/RestartAltTwoTone';
import { useIntl } from 'react-intl';
import {
    IconButton
    , Tooltip
} from '@mui/material';

type PropType = {
    onClick?: () => void
}

const ResetButton = ({ onClick }: PropType) => {
    return (
        <Tooltip title={useIntl().formatMessage({ id: "searchbar.reset" })}>
            <IconButton size="large" onClick={onClick}>
                <RestartAltTwoToneIcon />
            </IconButton>
        </Tooltip>
    )
}

export default ResetButton;