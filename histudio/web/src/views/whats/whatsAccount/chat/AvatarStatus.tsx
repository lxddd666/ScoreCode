// material-ui
import { useTheme } from '@mui/material/styles';

// assets
import FiberManualRecordIcon from '@mui/icons-material/FiberManualRecord';

// ==============================|| AVATAR STATUS ICONS ||============================== //

interface Props {
    status: string;
    mr?: number;
    ml?: number;
}

const AvatarStatus = ({ status, mr, ml }: Props) => {
    const theme = useTheme();
    switch (status) {
        case '1':
        case 'available':
            return (
                <FiberManualRecordIcon
                    sx={{
                        cursor: 'pointer',
                        color: theme.palette.success.dark,
                        verticalAlign: 'middle',
                        fontSize: '0.875rem',
                        mr
                    }}
                />
            );

        case 'away':
            return (
                <FiberManualRecordIcon
                    sx={{
                        cursor: 'pointer',
                        color: theme.palette.warning.dark,
                        verticalAlign: 'middle',
                        fontSize: '0.875rem',
                        mr
                    }}
                />
            );

        case 'do_not_disturb':
            return (
                <FiberManualRecordIcon
                    sx={{
                        cursor: 'pointer',
                        color: theme.palette.error.dark,
                        verticalAlign: 'middle',
                        fontSize: '0.875rem',
                        mr
                    }}
                />
            );

        case '2':
        case 'offline':
        case 'invisible':
            return (
                <FiberManualRecordIcon
                    sx={{
                        cursor: 'pointer',
                        color: theme.palette.grey[100],
                        verticalAlign: 'middle',
                        fontSize: '0.875rem',
                        mr
                    }}
                />
            );

        default:
            return null;
    }
};

export default AvatarStatus;
