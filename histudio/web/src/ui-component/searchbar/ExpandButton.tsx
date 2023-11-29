//material-ui
import ExpandCircleDownTwoToneIcon from '@mui/icons-material/ExpandCircleDownTwoTone';
import { useIntl } from 'react-intl';
import {
    IconButton
    , Tooltip
} from '@mui/material';

type PropType = {
    onClick?: () => void,
    transformValue?: string
}

const ExpandButton = ({ onClick, transformValue }: PropType) => {
    return (
        <Tooltip title={useIntl().formatMessage({ id: "searchbar.expand" })}>
            <IconButton size="large" onClick={onClick} sx={{ transform: transformValue ? transformValue : 'unset' }}>
                <ExpandCircleDownTwoToneIcon />
            </IconButton>
        </Tooltip>
    )
}

export default ExpandButton;