//material-ui
import LoginTwoToneIcon from '@mui/icons-material/LoginTwoTone';
import { IconButton, Tooltip } from '@mui/material';
import { useIntl } from 'react-intl';

type PropType = {
    onClick: (ids?: number[]) => void;
    tooltipTitle: string;
};

const LoginUserButton = ({ onClick, tooltipTitle }: PropType) => {
    const intl = useIntl();
    return (
        <Tooltip title={intl.formatMessage({ id: tooltipTitle })}>
            <IconButton size="large" onClick={() => onClick()}>
                <LoginTwoToneIcon />
            </IconButton>
        </Tooltip>
    );
};

export default LoginUserButton;
