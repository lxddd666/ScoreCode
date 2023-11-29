//material-ui
import QrCode2TwoToneIcon from '@mui/icons-material/QrCode2TwoTone';
import { useIntl } from 'react-intl';
import {
    IconButton
    , Tooltip
} from '@mui/material';

type PropType = {
    onClick?: () => void,
}

const QrButton = ({ onClick }: PropType) => {
    return (
        <Tooltip title={useIntl().formatMessage({ id: "searchbar.qr" })}>
            <IconButton size="large" onClick={onClick}>
                <QrCode2TwoToneIcon />
            </IconButton>
        </Tooltip>
    )
}

export default QrButton;